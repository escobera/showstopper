package api

import (
	"github.com/escobera/showstopper/resource"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ShowApi struct {
	db gorm.DB
}

func (api *ShowApi) CreateShow(c *gin.Context) {
	var show resource.Show

	c.Bind(&show)

	api.db.Save(&show)
	c.JSON(201, show)
}
