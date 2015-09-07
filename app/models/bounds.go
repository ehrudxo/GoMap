package models

type Bounds struct {
  Rectangle
}

func (b *Bounds) String() string {
	return b.Min.String() + "-" + b.Max.String()
}

func (b *Bounds) GetWidth() int {
	return b.Dx()
}

func (b *Bounds) GetHeight() int {
	return b.Dy()
}
