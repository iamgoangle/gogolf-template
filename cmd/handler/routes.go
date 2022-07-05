package handler

func (h *Handler) initRoutes() error {
	h.Server.GET("/healthcheck", healthCheckHandler)

	v1 := h.Server.Group("/v1")
	v1.POST("/users", h.CreateUserHandler)

	return nil
}
