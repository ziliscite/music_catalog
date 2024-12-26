package membership

import (
	"github.com/gin-gonic/gin"
	"music_catalog/internal/models/membership"
	"net/http"
)

type Service interface {
	SignUp(request *membership.SignUpRequest) error
}

type Handler struct {
	*gin.Engine
	s Service
}

func NewHandler(en *gin.Engine, s Service) *Handler {
	return &Handler{
		en,
		s,
	}
}

func (h *Handler) SignUp(c *gin.Context) {
	var request membership.SignUpRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.s.SignUp(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
