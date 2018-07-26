package model

import (
	"io"
	"math"
)

type TreePoint [3]Vertex

type Triangle struct {
	N Normal
	V TreePoint
}

func (t Triangle) GetZRange() (zMin, zMax StlFractionalType){
	zMin, zMax = t.V[0].Z, t.V[0].Z
	for i:=1; i<len(t.V); i++ {
		if zMax < t.V[i].Z {
			zMax = t.V[i].Z
		}
		if zMin > t.V[i].Z {
			zMin = t.V[i].Z
		}
	}
	return
}

type LayerSegment struct {
	N int
	Segment *PerimeterLineSegment
}

func (t Triangle) GetPerimeterSegments(step StlFractionalType) []*LayerSegment {
	zMin, zMax := t.GetZRange()
	tMin := math.Ceil(float64(zMin / step))
	tMax := math.Floor(float64(zMax / step))
	//fmt.Println("min max,", tMin,tMax)
	var segments []*LayerSegment
	for i := tMin; i <= tMax; i++ {
		//fmt.Println(i)
		segments = append(segments, &LayerSegment{int(i),t.V.GetIntersection(StlFractionalType(i)*step)})
	}
	return segments
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
	//fmt.Println("<=>", v0, v, v1)
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

/*
	Get intersection between z slicing face and triangle
 */

type PerimeterPoint struct {
	X StlFractionalType
	Y StlFractionalType
}
type PerimeterLineSegment struct {
	V [2]PerimeterPoint
	Z StlFractionalType
}

func (p PerimeterPoint) IsEqual(x,y StlFractionalType) bool {
	if p.X == x && p.Y == y {
		return true
	}
	return false
}

func (t TreePoint) GetIntersection(z StlFractionalType) *PerimeterLineSegment{
	/*
	x = a*t + x0 | a = x1 - x0
	y = b*t + y0 | b = y1 - y0
	z = c*t + z0 | z = z1 - z0
	 */
	 foundPoint := 0
	 lineSegment := PerimeterLineSegment{Z:z}
MainLoop:
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
			y := v0.Y.getATplusV(t, b)
			if x.isBetween(v0.X, v1.X) && y.isBetween(v0.Y, v1.Y) && z.isBetween(v0.Z, v1.Z){
				switch foundPoint {
				case 0:
					lineSegment.V[0] = PerimeterPoint{x,y}
					foundPoint++
				case 1:
					if !lineSegment.V[0].IsEqual(x, y) {
						lineSegment.V[1] = PerimeterPoint{x,y}
						foundPoint++
					}
				}
				if foundPoint == 2 {
					break MainLoop
				}
			}
		}
	 }
	 if foundPoint == 2 {
	 	return &lineSegment
	 }
	 return nil
}