package app

func mapUrls() {
	router.GET("/ping", controllers.Ping)

	// router.GET("/users/search", controllers.SearchUser)
	router.GET("/users/:user_id", controllers.GetUser)
	router.POST("/users", controllers.CreateUser)
}
