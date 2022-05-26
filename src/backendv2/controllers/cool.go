package controllers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"jasoncoding.com/backendv2/config"
	"jasoncoding.com/backendv2/cool"
)

type CoolJwtClaim struct {
	Cool    string   `json:"cool"`
	Action  string   `json:"action"`
	Choices []string `json:"choices"`
	jwt.StandardClaims
}

type Answer struct {
	Answer string `json:"answer"`
}

// Generate HMACSHA256 for answers
func getAnswerSignature(answer string) string {
	h := hmac.New(sha256.New, []byte(config.Cfg.GetString("JWT_SECRET")))
	h.Write([]byte(answer))
	return hex.EncodeToString(h.Sum(nil))
}

func GetCoolChallenge(c *gin.Context) {
	res, err := cool.GenCaptcha(3, 2)

	// TODO: JWT Secret rotation

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error creating captcha " + err.Error(),
		})
		return
	}

	// Close the reader at the end
	defer res.Close()

	// Set headers
	c.Status(http.StatusOK)
	c.Header("Content-Type", res.Format)

	// Generate jwt for xcool
	var coolTok strings.Builder

	// Questions
	var coolQ strings.Builder
	for _, x := range res.Challenge {
		coolTok.WriteString(x.Answer)
		coolQ.WriteString(x.Question + ",")
	}

	// Generate jwt
	curTime := time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CoolJwtClaim{
		Cool:    getAnswerSignature(coolTok.String()),
		Action:  c.Param("action"),
		Choices: res.Choices,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  curTime,
			ExpiresAt: curTime + 60, // In seconds
		},
	})

	tokenstr, err := token.SignedString([]byte(config.Cfg.GetString("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "JWT Error",
		})
		return
	}

	c.Header("X-Cool", tokenstr)
	c.Header("X-Cool-Challenge", coolQ.String())

	// Send captcha as stream
	c.Stream(func(w io.Writer) bool {
		for {
			buf := make([]byte, 2048)
			_, err := res.Reader.Read(buf)
			if err == io.EOF {
				break
			}
			w.Write(buf)
		}
		// False I think is for (Should gin keep the connection?)
		return false
	})
}

func VerifyChallenge(action string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if !strings.HasPrefix(authorization, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "no bearer",
			})
			return
		}

		// Get body early
		var answer Answer
		if err := c.ShouldBindJSON(&answer); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Parse token
		token, err := jwt.ParseWithClaims(strings.Replace(authorization, "Bearer ", "", 1), &CoolJwtClaim{}, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("signing method error")
			}

			// Second variable is whether the type assertion suceeded
			claims, ok := t.Claims.(*CoolJwtClaim)
			if !ok {
				return nil, errors.New("assertion failed")
			}

			if claims.Action != action {
				return nil, errors.New("action mismatch")
			}

			// Verify the answer
			if claims.Cool != getAnswerSignature(answer.Answer) {
				return nil, errors.New("wrong answer")
			}
			return []byte(config.Cfg.GetString("JWT_SECRET")), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		} else if !token.Valid {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "jwt not valid",
			})
			return
		}

		c.Next()
	}
}
