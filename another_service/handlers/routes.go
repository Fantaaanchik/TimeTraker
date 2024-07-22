package handlers

func (h Handler) AllRoutes() {
	h.Engine.GET("/info", h.UserInformation)
}
