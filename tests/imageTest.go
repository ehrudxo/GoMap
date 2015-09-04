package main

import (
    "image"
    "fmt"
    "github.com/ninjasphere/draw2d/draw2dimg"
    "github.com/ninjasphere/draw2d/samples/android"
)
//ninjasphere is very important
func main(){
    // Initialize the graphic context on an RGBA image
    dest := image.NewRGBA(image.Rect(0, 0, 297, 210.0))
    gc := draw2dimg.NewGraphicContext(dest)
    // Draw Android logo
    fn, err := android.Main(gc, "png")
    if err != nil {
      fmt.Println("Drawing  failed");
        // t.Errorf("Drawing %q failed: %v", fn, err)
        return
    }
    // Save to png
    err = draw2dimg.SaveToPngFile(fn, dest)
    if err != nil {
      fmt.Println("Saving failed",err);
        // t.Errorf("Saving %q failed: %v", fn, err)
    }
}
