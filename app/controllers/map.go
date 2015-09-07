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
	var bbox,srs,widthStr,heightStr string
	srsNo:=2097
	c.Params.Bind(&bbox, "BBOX")
	c.Params.Bind(&srs, "SRS")
	c.Params.Bind(&widthStr, "WIDTH")
	c.Params.Bind(&heightStr, "HEIGHT")
	splited:=strings.Split(srs,"EPSG:")
	if(len(splited)>1){
		srsNo,_ = strconv.Atoi(splited[1])
	}else{
		srsNo,_ = strconv.Atoi(splited[0])
	}
	var bboxSlice = strings.Split(bbox,",");
	var toilets []models.Toilet

	minx,err := strconv.ParseFloat(bboxSlice[0],64)
	miny,err := strconv.ParseFloat(bboxSlice[1],64)
	maxx,err := strconv.ParseFloat(bboxSlice[2],64)
	maxy,err := strconv.ParseFloat(bboxSlice[3],64)
	width,err  := strconv.Atoi(widthStr)
	height,err := strconv.Atoi(heightStr)

	selectQuery:="Gid,Gu_nm,Hnr_nam,Mtc_at,Masterno,Slaveno,Neadres_nm,Wc_nam,Wc_gbn,Hd_wc_yno,Creat_de,Po_fe_nm,"
	selectQuery+="ST_Transform(Geom,3857) as Geom"
	whereQuery:="geom && ST_Transform(ST_MakeEnvelope(?,?,?,?,?),2097)"
  Db.Where(whereQuery,minx,miny,maxx,maxy,srsNo).Select(selectQuery).Find(&toilets)

	minP := models.Point{int(minx),int(miny)}
  maxP := models.Point{int(maxx),int(maxy)}
  bnd := &models.Bounds{models.Rectangle{minP,maxP}}
  filename := services.DrawPoint( toilets, bnd, width,height )
 // 	fmt.Println(revel.BasePath,filename);
	file, _ := os.Open(revel.BasePath+"/output/wms/"+filename)

	if err != nil {
		fmt.Println(err) //the err is nil
		return nil
	}
	defer func(){
		os.Remove(revel.BasePath+"/output/wms/"+filename)
	}()
	return c.RenderFile(file, revel.Inline )//Not an attachment. But Who disconnect the file.
}
