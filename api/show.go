package api

import (
	"strconv"

	"github.com/escobera/showstopper/resource"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ShowAPI struct {
	Db gorm.DB
}

func (api *ShowAPI) CreateShow(c *gin.Context) {
	showJSON := resource.ShowJSON{}

	c.Bind(&showJSON)

	api.Db.Save(&showJSON.Show)

	c.JSON(201, showJSON)
}

func (api *ShowAPI) UpdateShow(c *gin.Context) {
	showJSON := resource.ShowJSON{}

	api.Db.First(&showJSON.Show, c.Params.ByName("id"))

	c.Bind(&showJSON)

	showID, _ := strconv.Atoi(c.Params.ByName("id"))
	showJSON.Show.ID = uint32(showID)
	api.Db.Save(&showJSON.Show)
	c.Data(204, gin.MIMEHTML, nil)
}

func (api *ShowAPI) IndexShows(c *gin.Context) {
	shows := []resource.Show{}
	api.Db.Find(&shows)

	c.JSON(200, map[string]interface{}{"shows": shows})
}
