package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	loggerEntry *logrus.Entry
}

func NewHandler(
	logger *logrus.Logger,
) *Handler {
	loggerEntry := logger.WithField("service_name", "http_handler")

	return &Handler{
		loggerEntry: loggerEntry,
	}
}

func (h *Handler) Post(c *gin.Context) {
	var req *postRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.loggerEntry.Errorf("incorrect json in request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"name": req.Name})
}
