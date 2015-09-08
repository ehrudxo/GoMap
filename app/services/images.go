package services

import (
  "image"
  "image/color"
  "fmt"
  "strconv"
  "github.com/revel/revel"
  "github.com/ninjasphere/draw2d/draw2dkit"
  "github.com/ninjasphere/draw2d/draw2dimg"
  "github.com/ehrudxo/GoMap/app/models"
)

func PrintPoint(toilets []models.Toilet) {
  fmt.Println(len(toilets));
  for idx := range toilets{
    fmt.Println("toilet :", toilets[idx]);
  }
}
func DrawPoint( toilets []models.Toilet, b *models.Bounds, imgWidth int, imgHeight int) string{
  // fmt.Println(imgWidth,imgHeight);
  filename := strconv.Itoa(b.Min.X) + "_" + strconv.Itoa(b.Min.Y)
  filename += "_" + strconv.Itoa(b.Max.X) + "_" + strconv.Itoa(b.Max.Y)
  filename += "_" +strconv.Itoa(imgWidth) + "_" + strconv.Itoa(imgHeight) +".png"
  HRes := b.GetWidth()/imgWidth
  VRes := b.GetHeight()/imgHeight
  dest := image.NewRGBA(image.Rect(0, 0, imgWidth,imgHeight))
  gc := draw2dimg.NewGraphicContext(dest)
  gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})

  for _,toilet := range toilets{
    dx := (toilet.Geom.Lng-float64(b.Min.X))/float64(HRes)//x
    dy := float64(imgHeight) - (toilet.Geom.Lat-float64(b.Min.Y))/float64(VRes)//y
    // fmt.Println("toilet :", toilet,dx,dy);
    draw2dkit.Circle(gc, dx, dy, 3)
    gc.FillStroke()
  }
  gc.Close()
  draw2dimg.SaveToPngFile(revel.BasePath+"/output/wms/"+filename, dest)
  return filename
}
