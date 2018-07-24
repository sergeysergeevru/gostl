package gcode

type GcodeMethods interface {
	MoveTo(x,y float32)
	MoveFromTo(x1,y1,x2,y2 float32)
}