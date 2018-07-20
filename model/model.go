package model

import (
	"io"
	"fmt"
)

type TreePoint [3]Vertex

type Triangle struct {
	N Normal
	V TreePoint
}

type Normal struct{
	I, J, K StlFractionalType
	Reader FacetReader
}
func (n *Normal) GetFromLine(r io.Reader){
	n.I, n.J, n.K = n.Reader.ReadNormal(r)
}

type StlFractionalType float32

func (v StlFractionalType) getATplusV(t StlFractionalType,p StlFractionalType) StlFractionalType {
	return p*t+ v
}

func (v StlFractionalType) isBetween(v0, v1 StlFractionalType) bool {
	if v0 <= v && v <= v1 || v0 >= v && v >= v1 {
		return true
	}
	return false
}

type Vertex struct{
	X, Y, Z StlFractionalType
	Reader FacetReader
}

type FacetReader interface {
	ReadVertex(r io.Reader) (x,y,z StlFractionalType)
	ReadNormal(r io.Reader) (i,j,k StlFractionalType)
}

func (v *Vertex) GetFromLine(r io.Reader){
	v.X, v.Y, v.Z = v.Reader.ReadVertex(r)
}

type Model struct {

}

func (t TreePoint) GetIntersection(z StlFractionalType){
	/*
	x = a*t + x0 | a = x1 - x0
	y = b*t + y0 | b = y1 - y0
	z = c*t + z0 | z = z1 - z=
	 */
	 for i, v0 := range t {
	 	for j, v1 := range t {
	 		if i == j {
	 			break
			}
			a := v1.X - v0.X
			b := v1.Y - v0.Y
			c := v1.Z - v0.Z
			t := (z - v0.Z)/c
			x := v0.X.getATplusV(t, a)
			y := v0.X.getATplusV(t, b)
			fmt.Println(a,b,c, t,x,y)
			if x.isBetween(v0.X, v1.X) && y.isBetween(v0.Y, v1.Y) && z.isBetween(v0.Z, v1.Z){
				fmt.Println(x, y, z)
			}

		}
	 }

}