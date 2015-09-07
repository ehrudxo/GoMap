package controllers

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	_ "net/http"
	"github.com/revel/revel"
	"github.com/ehrudxo/revel_test1/app/models"
	"github.com/ehrudxo/revel_test1/app/services"
)
const epsg = 2097

type GoMap struct {
	*revel.Controller
}

func (c GoMap) Index() revel.Result {
  // maps:=&models.Seoul{}
  // Db.First(&maps)
	maps:="wow"
	return c.Render(maps)
}
//look at the Custom Result More.

func (c GoMap) Toilet() revel.Result {
	var bbox string
	c.Params.Bind(&bbox, "BBOX")
	var bboxSlice = strings.Split(bbox,",");
	fmt.Println(bbox);
  var toilets []models.Toilet
	//177000,437000,219000,466000
	minx,_ := strconv.Atoi(bboxSlice[0])
	miny,_ := strconv.Atoi(bboxSlice[1])
	maxx,_ := strconv.Atoi(bboxSlice[2])
	maxy,_ := strconv.Atoi(bboxSlice[3])
  Db.Where("geom && ST_MakeEnvelope(?,?,?,?,?)",minx,miny,maxx,maxy,epsg).Find(&toilets);
	fmt.Println();
	// minP := models.Point{177000, 437000}
	// maxP := models.Point{219000, 466000}
	// bnd := &models.Bounds{models.Rectangle{minP,maxP}}
	fmt.Println(minx,miny,maxx,maxy);
	minP := models.Point{minx,miny}
  maxP := models.Point{maxx,maxy}
  bnd := &models.Bounds{models.Rectangle{minP,maxP}}
  filename := services.DrawPoint( toilets, bnd, 1024, 768 )
	fmt.Println(revel.BasePath,filename);
	file, _ := os.Open(revel.BasePath+"/output/wms/"+filename)
	// defer file.Close()
	fileInfo, err:= file.Stat()
	fmt.Println(fileInfo.Size(), err) //the err is nil
	// return c.Render(filename)
	return c.RenderFile(file, revel.Inline )//Not an attachment. But Who disconnect the file.
}
