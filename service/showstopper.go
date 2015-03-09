package service

import (
	"fmt"

	"github.com/escobera/showstopper/api"
	"github.com/escobera/showstopper/resource"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/tommy351/gin-cors"
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
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cfg.DbUser, cfg.DbPassword, cfg.DbName)

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
	r.Use(cors.Middleware(cors.Options{}))

	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello world")
	})

	ShowAPI := api.ShowAPI{Db: db}

	r.OPTIONS("/*cors", func(c *gin.Context) {})
	r.GET("/shows", ShowAPI.IndexShows)
	r.POST("/shows", ShowAPI.CreateShow)
	r.PUT("/shows/:id", ShowAPI.UpdateShow)
	r.DELETE("/shows/:id", ShowAPI.DeleteShow)

	r.Run(cfg.SvcHost)

	return nil
}
