package route

import (
	"boilerplate/backend-refactor/config"
	"boilerplate/backend-refactor/handler"
	"boilerplate/backend-refactor/middleware"
	repo2 "boilerplate/backend-refactor/repo"
	"boilerplate/backend-refactor/service"
)

type Service struct {
	*config.App
}

func NewService() *Service {
	s := Service{
		config.NewApp(),
	}

	db := s.GetDB()
	repo := repo2.NewRepo(db)

	userService := service.NewUser(repo)
	user := handler.NewUser(userService)
	migrate := handler.NewMigration(db)

	router := s.Router
	v1 := router.Group("/v1")

	// user
	v1.POST("/login", user.Login)

	// migration
	router.POST("/migrate", middleware.CheckAdmin(), migrate.Migrate)
	return &s
}
