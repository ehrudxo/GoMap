package controllers

import (
	"github.com/revel/revel"
	"github.com/ehrudxo/revel_test1/app/models"
  // 
)

type GoMap struct {
	*revel.Controller
}

func (c GoMap) Index() revel.Result {
  maps:=&models.Seoul{}
  Db.First(&maps)
	// // Db.NewRecord(user)
	// // Db.Create(&user)
	return c.Render(maps)
}
