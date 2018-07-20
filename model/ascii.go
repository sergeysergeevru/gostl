package model

import (
	"io"
	"fmt"
)

type AsciiStlReader struct {

}
func (a AsciiStlReader) ReadVertex (r io.Reader) (x,y,z StlFractionalType) {
	_, err := fmt.Fscanf(r, "  vertex %f %f %f\n", &x, &y, &z)
	if err != nil {
		panic(err)
	}
	return
}
func (a AsciiStlReader) ReadNormal(r io.Reader) (i,j,k StlFractionalType){
	_, err := fmt.Fscanf(r," facet normal %f %f %f\n  outer loop\n", &i, &j, &k)
	if err != nil {
		panic(err)
	}
	return
}