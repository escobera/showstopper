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
	var showJSON resource.ShowJSON

	c.Bind(&showJSON)

	show := showJSON.Show
	api.Db.Save(&show)

	c.JSON(201, show)
}

func (api *ShowAPI) UpdateShow(c *gin.Context) {
	showJSON := resource.ShowJSON{}

	c.Bind(&showJSON)

	show := showJSON.Show

	showID, _ := strconv.Atoi(c.Params.ByName("id"))
	show.ID = uint32(showID)
	api.Db.Save(&show)

	c.Data(204, gin.MIMEHTML, nil)
}

func (api *ShowAPI) IndexShows(c *gin.Context) {
	shows := make([]resource.Show, 0)

	api.Db.Find(&shows)
	c.JSON(200, map[string]interface{}{"shows": shows})
}
