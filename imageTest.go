package main

import (
  "fmt"
  "github.com/ehrudxo/GoMap/app/services"
  "github.com/ehrudxo/GoMap/app/controllers"
  "github.com/ehrudxo/GoMap/app/models"
)
//ninjasphere is very important
func main(){
  // Initialize the graphic context on an RGBA image
  controllers.InitDB();
  var toilets []models.Toilet
  controllers.Db.Where("geom && ST_MakeEnvelope(?,?,?,?,?)",177000,437000, 219000, 466000,2097).Find(&toilets);
  //services.PrintPoint(toilets);
  minP := models.Point{177000, 437000}
  maxP := models.Point{219000, 466000}
  bnd := &models.Bounds{models.Rectangle{minP,maxP}}
  filename:=services.DrawPoint( toilets, bnd, 1024, 768 )
  fmt.Println(filename);
}
// func DrawTest(){
//   dest := image.NewRGBA(image.Rect(0, 0, 297, 210.0))
//   gc := draw2dimg.NewGraphicContext(dest)
//   // Draw Android logo
//   // fn, err := android.Main(gc, "png")
//   // if err != nil {
//   //   fmt.Println("Drawing  failed");
//   //     // t.Errorf("Drawing %q failed: %v", fn, err)
//   //     return
//   // }
//   // // Save to png
//   // err = draw2dimg.SaveToPngFile(fn, dest)
//   // if err != nil {
//   //   fmt.Println("Saving failed",err);
//   //     // t.Errorf("Saving %q failed: %v", fn, err)
//   // }
//   // Set some properties
//   gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
//   gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
//   gc.SetLineWidth(5)
//
//   // Draw a closed shape
//   gc.MoveTo(10, 10) // should always be called first for a new path
//   gc.LineTo(100, 50)
//   gc.QuadCurveTo(100, 10, 10, 10)
//   gc.Close()
//   gc.FillStroke()
//
//   // Save to file
//   draw2dimg.SaveToPngFile("output/samples/hello.png", dest)
// }
