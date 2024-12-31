package track

import "music_catalog/internal/middleware"

func (h *Handler) RegisterRoutes() {
	r := h.Group("/track")
	r.Use(middleware.AuthMiddleware())
	r.GET("/search", h.Search)
	r.POST("/upsert", h.Upsert)
}
