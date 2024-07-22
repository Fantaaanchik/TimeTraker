package controllers

func (uc UserController) AllUserRoutes() {

	uc.Engine.POST("/create_user", uc.CreateUser)
	uc.Engine.GET("/get_all_users", uc.GetAllUsers)
	uc.Engine.GET("/get_user_by_id/:id", uc.GetUserByID)
	uc.Engine.PUT("/update_user/:id", uc.UpdateUser)
	uc.Engine.DELETE("/delete_user/:id", uc.DeleteUser)

}
