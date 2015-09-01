package controllers

import (
	"github.com/revel/revel"
	"github.com/jinzhu/gorm"
	"map1/app/models"
)

type App struct {
	*revel.Controller
	Tx *gorm.DB
}

func (c App) Index() revel.Result {
	user := &models.User{Username: "keen2"}
	c.Tx.NewRecord(user)
	c.Tx.Create(user)
	return c.Render(user)
}
