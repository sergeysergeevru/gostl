package model

import (
	"io"
	"encoding/binary"
)

type BinaryStlReader struct {

}

func (a BinaryStlReader) ReadNormal (r io.Reader) (x,y,z StlFractionalType) {
	x,y,z = a.getVector(r)
	return
}

func (a BinaryStlReader) ReadVertex (r io.Reader) (x,y,z StlFractionalType) {
	x,y,z = a.getVector(r)
	return
}

func (a BinaryStlReader) getVector(r io.Reader) (x,y,z StlFractionalType) {
	for _, v := range []*StlFractionalType{&x,&y,&z} {
		err :=  binary.Read(r, binary.LittleEndian, v)
		if err != nil {
			panic(err)
		}
	}
	return
}