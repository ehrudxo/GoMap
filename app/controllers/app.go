package controllers

import (
	"github.com/revel/revel"
	"github.com/ehrudxo/GoMap/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	user := &models.User{Seq:0,Username: "keen2" ,Description:"keen is keen" ,Password:"1qaz" ,Mail:"ehrudxo@hotmail.com"}
	// Db.NewRecord(user)
	// Db.Create(&user)
	return c.Render(user)
}

func (c App) GetUser() revel.Result {
	user := &models.User{Seq:0,Username: "keen2" ,Description:"keen is keen" ,Password:"1qaz" ,Mail:"ehrudxo@hotmail.com"}
	// Db.NewRecord(user)
	// Db.Create(&user)
	return c.RenderJson(user)
}
