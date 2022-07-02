package handler

func (h *Handler) initRoutes() error {
	h.e.GET("/healthcheck", healthCheckHandler)

	v1 := h.e.Group("/v1")
	v1.POST("/users", createUserHandler)

	return nil
}
