package controllers

import (
	"github.com/revel/revel"
	"github.com/ehrudxo/revel_test1/app/models"
)

type GoMap struct {
	*revel.Controller
}

func (c GoMap) Index() revel.Result {
  maps:=&models.Seoul{}
  Db.First(&maps)
	return c.Render(maps)
}
func (c GoMap) Toilet() revel.Result {
  var maps []models.Toilet
  Db.Where("geom && ST_MakeEnvelope(?,?,?,?,?)",177000,437000, 219000, 466000,2097).Find(&maps);
	return c.RenderJson(maps)
}
