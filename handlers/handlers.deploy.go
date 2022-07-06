package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"bojago/models"
)

func ShowIndexPage(c *gin.Context) {
	account := models.HntCloud //default account
	queryParams := c.Request.URL.Query()
	queryAccountKey := "account"
	fmt.Println(queryParams)

	if queryParams.Has(queryAccountKey) {
		if acc, err := strconv.Atoi(queryParams[queryAccountKey][0]); err == nil {
			if acc == models.Fnd {
				account = models.Fnd
			}
		}
	}

	deployments := models.GetDeployments(account)

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses
		gin.H{
			"title":   "BojaGo",
			"payload": deployments,
		},
	)

}
