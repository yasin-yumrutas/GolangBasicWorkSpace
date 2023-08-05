package controller

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yasin-yumrutas/config"
)

func GoogleLogin(c *gin.Context) {
	googleConfig := config.SetupConfig()

	url := googleConfig.AuthCodeURL("randomclient")

	// Redirect to google login page
	c.Redirect(http.StatusSeeOther, url)
}

func GoogleCallBack(c *gin.Context) {
	state := c.Query("state")
	if state != "randomclient" {
		c.String(http.StatusBadRequest, "states dont match")
		return
	}

	// Code
	code := c.Query("code")

	// Configuration
	googleConfig := config.SetupConfig()

	// Exchange code for token
	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		c.String(http.StatusInternalServerError, "Code-Token Exchange failed")
		return
	}

	// Fetch user data from Google API
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		c.String(http.StatusInternalServerError, "User Data fetch Failed")
		return
	}
	defer response.Body.Close()

	// Parse response
	userData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "JSON Parsing Failed")
		return
	}

	// Send response to client
	c.String(http.StatusOK, string(userData))
}
