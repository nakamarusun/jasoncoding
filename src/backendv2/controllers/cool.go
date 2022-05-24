package controllers

import (
	"io"

	"github.com/gin-gonic/gin"
	"jasoncoding.com/backendv2/cool"
)

func GetCoolChallenge(c *gin.Context) {
	res, err := cool.GenCaptcha()

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
	c.Header("Content-Type", "image/jpeg")

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
