package track

import (
	"github.com/gin-gonic/gin"
	model "music_catalog/internal/model/spotify"
	"music_catalog/internal/model/usertrack"
	"net/http"
	"strconv"
)

//go:generate mockgen -source=handler.go -destination=handler_mock.go -package=track
type Service interface {
	Search(query string, pageSize, pageIndex int, userId uint) (*model.SearchResponse, error)
	Upsert(userId uint, request *usertrack.LikeRequest) (bool, error)
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
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 10
	}

	pageIndexStr := c.Query("pageIndex")
	pageIndex, err := strconv.Atoi(pageIndexStr)
	if err != nil {
		pageIndex = 1
	}

	response, err := h.s.Search(query, pageSize, pageIndex, c.GetUint("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) Upsert(c *gin.Context) {
	var request usertrack.LikeRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := h.s.Upsert(c.GetUint("userID"), &request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !created {
		c.Status(http.StatusOK)
		return
	}

	c.Status(http.StatusCreated)
}
