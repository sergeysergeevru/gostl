package reprap

import "io"

var state struct {
	Position *Position
	Writer io.Writer
	CommentsEnabled bool
}

type Position struct {
	X,Y,Z float32
}

func init()  {
	state.Position = &Position{X:0,Y:0,Z:0}
	state.CommentsEnabled = false
}