package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/nibble-4bits/ondemand-go-bootcamp/adapter"
	"github.com/nibble-4bits/ondemand-go-bootcamp/usecase"

	"github.com/gin-gonic/gin"
)

// GetCommentByID handles the logic of getting a comment by ID from a data source and sending it to an http client
// as a JSON object.
//
// If an error occurs, it will send that instead.
func GetCommentByID(service usecase.CommentService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		comment, err := service.GetByID(id)
		switch {
		case errors.Is(err, adapter.ErrCommentNotFoundByID):
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, comment)
	}
}
