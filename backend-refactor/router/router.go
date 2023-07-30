package router

import (
	"boilerplate/backend-refactor/config"
	"boilerplate/backend-refactor/handler"
	"boilerplate/backend-refactor/middleware"
	repo2 "boilerplate/backend-refactor/repo"
	"boilerplate/backend-refactor/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Service struct {
	db     *config.App
	router *gin.Engine
}

func NewService() *Service {
	db := &gorm.DB{}
	s := Service{
		db:     config.NewApp(db),
		router: gin.Default(),
	}

	repo := repo2.NewRepo(s.db.GetDB())

	userService := service.NewUser(repo)
	user := handler.NewUser(userService)
	migrate := handler.NewMigration(s.db.GetDB())

	router := s.router
	v1 := router.Group("/v1")

	// user
	v1.POST("/login", user.Login)

	// migration
	router.POST("/migrate", middleware.CheckAdmin(), migrate.Migrate)

	router.Run(":8000")
	return &s
}
