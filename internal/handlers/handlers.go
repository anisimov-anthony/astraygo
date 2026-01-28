package handlers

import (
	"net/http"
	"strconv"

	"github.com/anisimov-anthony/astraygo/internal/service"
	"github.com/gin-gonic/gin"
)

type AstrayHandler struct {
	service *service.AstrayService
}

func InitAstrayHandler(service *service.AstrayService) *AstrayHandler {
	return &AstrayHandler{service: service}
}

func (h *AstrayHandler) GetObjects(c *gin.Context) {
	statusQuery := c.Query("status")

	var status *bool
	if statusQuery == "" {
		status = nil
	} else if statusQuery == "true" {
		val := true
		status = &val
	} else if statusQuery == "false" {
		val := false
		status = &val
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status must be 'true' or 'false'"})
		return
	}

	objects, err := h.service.GetObjects(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot get objects"})
		return
	}

	c.JSON(http.StatusOK, objects)
}

func (h *AstrayHandler) GetObjectByID(c *gin.Context) {
	id, error := strconv.ParseInt(c.Param("id"), 10, 64)
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorect id"})
		return
	}

	object, err := h.service.GetObjectByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "cannot get object"})
		return
	}

	c.JSON(http.StatusOK, object)
}

func (h *AstrayHandler) PostObject(c *gin.Context) {
	var updatedObject service.ObjectInfo
	if err := c.BindJSON(&updatedObject); err != nil {
		return
	}

	updated, err := h.service.UpdateObjectLocation(&updatedObject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot update object"})
		return
	}

	c.JSON(http.StatusCreated, updated)
}
