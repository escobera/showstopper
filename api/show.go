package api

import (
	"strconv"

	"github.com/escobera/showstopper/resource"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ShowApi struct {
	Db gorm.DB
}

func (api *ShowApi) CreateShow(c *gin.Context) {
	var show resource.Show

	c.Bind(&show)

	api.Db.Save(&show)
	c.JSON(201, show)
}

func (api *ShowApi) UpdateShow(c *gin.Context) {
	showJSON := resource.ShowJSON{}

	c.Bind(&showJSON)

	show := showJSON.Show

	showID, _ := strconv.Atoi(c.Params.ByName("id"))
	show.ID = uint32(showID)
	api.Db.Save(&show)

	c.Data(204, gin.MIMEHTML, nil)
}

func (api *ShowApi) IndexShows(c *gin.Context) {
	shows := make([]resource.Show, 0)

	api.Db.Find(&shows)
	c.JSON(200, map[string]interface{}{"shows": shows})
}
