package service

import (
	"fmt"

	"github.com/escobera/showstopper/api"
	"github.com/escobera/showstopper/resource"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Config struct {
	SvcHost    string
	DbUser     string
	DbPassword string
	DbHost     string
	DbName     string
}

type ShowStopper struct {
}

func (s *ShowStopper) getDb(cfg Config) (gorm.DB, error) {
	connectionString := fmt.Sprintf("user=%s dbname=%s sslmode=disable", cfg.DbUser, cfg.DbName)

	return gorm.Open("postgres", connectionString)
}

func (s *ShowStopper) Migrate(cfg Config) error {
	db, err := s.getDb(cfg)
	if err != nil {
		fmt.Println(err)
	}
	db.LogMode(true)
	defer db.Close()

	db.AutoMigrate(&resource.Show{})
	return nil
}

func (s *ShowStopper) Run(cfg Config) error {
	db, err := s.getDb(cfg)
	if err != nil {
		return err
	}
	defer db.Close()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello world")
	})

	showApi := api.ShowApi{db: db}

	r.POST("/shows", showApi.CreateShow)

	// r.GET("/todo", todoResource.GetAllTodos)
	// r.GET("/todo/:id", todoResource.GetTodo)
	// r.POST("/todo", todoResource.CreateTodo)
	// r.PUT("/todo/:id", todoResource.UpdateTodo)
	// r.PATCH("/todo/:id", todoResource.PatchTodo)
	// r.DELETE("/todo/:id", todoResource.DeleteTodo)

	r.Run(cfg.SvcHost)

	return nil
}
