package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"bojago/models"
)

func ShowIndexPage(c *gin.Context) {
	deployments := models.GetDeployments(models.Fnd)

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses
		gin.H{
			"title":   "Home Page",
			"payload": deployments,
		},
	)

}
