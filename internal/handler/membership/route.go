package membership

func (h *Handler) RegisterRoutes() {
	r := h.Group("/membership")
	r.POST("/signup", h.SignUp)
	r.POST("/signin", h.SignIn)
}
