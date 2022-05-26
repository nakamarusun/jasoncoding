package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"jasoncoding.com/backendv2/config"
)

type gRecaptchaRes struct {
	Success bool     `json:"success"`
	Action  string   `json:"action"`
	Score   float32  `json:"score"`
	Errors  []string `json:"error-codes"`
}

type getContact struct {
	Response string `json:"response"`
}

func GetIdentityGoog(c *gin.Context) {

	var reqBody getContact

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	params := fmt.Sprintf("?secret=%s&response=%s", config.Cfg.GetString("GCAPTCHA_SECRET"), reqBody.Response)

	// Create a request to Google's recaptcha confirmer
	res, err := http.Post("https://www.google.com/recaptcha/api/siteverify"+params, "", nil)

	// Handle error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	defer res.Body.Close()

	// Read body
	var body gRecaptchaRes
	json.NewDecoder(res.Body).Decode(&body)

	// Do some checking
	if !body.Success || body.Action != "getcontact" || body.Score < 0.5 {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Bot Detected.",
		})
		return
	}

	// Send the contact
	c.Data(http.StatusOK, "application/json", []byte(config.Cfg.GetString("CONTACT")))
}

func GetIdentity(c *gin.Context) {
	c.Data(http.StatusOK, "application/json", []byte(config.Cfg.GetString("CONTACT")))
}
