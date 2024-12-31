package track

import (
	"github.com/gin-gonic/gin"
	model "music_catalog/internal/model/spotify"
	"net/http"
	"strconv"
)

//go:generate mockgen -source=handler.go -destination=handler_mock.go -package=track
type Service interface {
	Search(query string, pageSize, pageIndex int) (*model.SearchResponse, error)
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

func (h *Handler) Search(c *gin.Context) {
	query := c.Query("q")
	pageSizeStr := c.Query("pageSize")
	pageIndexStr := c.Query("pageIndex")

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 10
	}

	pageIndex, err := strconv.Atoi(pageIndexStr)
	if err != nil {
		pageIndex = 1
	}

	response, err := h.s.Search(query, pageSize, pageIndex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
