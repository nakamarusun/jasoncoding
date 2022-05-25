package controllers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"jasoncoding.com/backendv2/cool"
)

type CoolJwtClaim struct {
	Cool   string `json:"cool"`
	Action string `json:"action"`
	jwt.StandardClaims
}

var secret = []byte("dashdjkashdjkashdjask")

func GetCoolChallenge(c *gin.Context) {
	res, err := cool.GenCaptcha(3, 2)

	// TODO: JWT Secret rotation

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Server error creating captcha",
		})
		return
	}

	// Close the reader at the end
	defer res.Close()

	// Set headers
	c.Status(200)
	c.Header("Content-Type", res.Format)

	// Generate jwt for xcool
	var coolTok strings.Builder

	// Questions
	var coolQ strings.Builder
	for _, x := range res.Challenge {
		coolTok.WriteString(x.Answer)
		coolQ.WriteString(x.Question + ",")
	}

	// Generate HMACSHA256 for answers
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(coolTok.String()))

	// Generate jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CoolJwtClaim{
		Cool:   hex.EncodeToString(h.Sum(nil)),
		Action: "//TODO",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 60_000,
			Issuer:    "cool",
		},
	})

	tokenstr, err := token.SignedString(secret)
	if err != nil {
		c.JSON(500, gin.H{
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
		return false
	})
}

func VerifyChallenge(c *gin.Context) {

}