package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/svadr/gofbposter-service/internal/models"
)

func (app *Application) createPost(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.Writer.Header().Set("Allow", http.MethodPost)
		app.clientError(c.Writer, http.StatusMethodNotAllowed)
		return
	}

	var pReq models.CreatePostRequest

	if err := c.ShouldBindJSON(&pReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inputs"})
		return
	}

	posts, err := app.posts.Insert(pReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}
