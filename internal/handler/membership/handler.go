package membership

import (
	"errors"
	"github.com/gin-gonic/gin"
	"music_catalog/internal/models/membership"
	"net/http"
)

//go:generate mockgen -source=handler.go -destination=handler_mock.go -package=membership
type Service interface {
	SignUp(request *membership.SignUpRequest) error
	SignIn(request *membership.SignInRequest) (string, error)
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

	err := h.s.SignUp(&request)
	if errors.Is(err, membership.ErrUserAlreadyExists) {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *Handler) SignIn(c *gin.Context) {
	var request membership.SignInRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := h.s.SignIn(&request)
	if errors.Is(err, membership.ErrUserNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if errors.Is(err, membership.ErrInvalidCredentials) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, membership.SignInResponse{AccessToken: accessToken})
}
