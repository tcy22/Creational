package main

import "fmt"

type Color interface {
	Use()
}

type Red struct{}
func (r Red) Use() {
	fmt.Println("Use Red color")
}

type Green struct{}
func (g Green) Use() {
	fmt.Println("Use Green color")
}

type Yellow struct{}
func (y Yellow) Use() {
	fmt.Println("Use Yellow color")
}



type BrushPen interface {
	DrawPicture()
}

type BigBrushPen struct {
	Color
}
func (bbp BigBrushPen) DrawPicture() {
	fmt.Println("Draw picture with big brush pen")
	bbp.Use()
}

type SmallBrushPen struct {
	Color
}
func (sbp SmallBrushPen) DrawPicture() {
	fmt.Println("Draw picture with small brush pen")
	sbp.Use()
}


func NewBrushPen(t string, color Color) BrushPen {
	switch t {
	case "BIG":
		return BigBrushPen{
			Color: color,
		}
	case "SMALL":
		return SmallBrushPen{
			Color: color,
		}
	default:
		return nil
	}
}

func main() {
	var tColor Color
	tColor = Red{}
	tBrushPen := NewBrushPen("BIG", tColor)
	tBrushPen.DrawPicture()
	fmt.Println("=========================")
	tColor = Green{}
	tBrushPen = NewBrushPen("SMALL", tColor)
	tBrushPen.DrawPicture()
	fmt.Println("=========================")
	tColor = Yellow{}
	tBrushPen = NewBrushPen("BIG", tColor)
	tBrushPen.DrawPicture()
}