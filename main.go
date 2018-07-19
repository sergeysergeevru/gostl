package main

import (
	"fmt"
	"io"
	"encoding/binary"
	"os"
	"bufio"
)

type Triangle struct {
	N Normal
	V [3]Vertex
}

type Normal struct{
	I, J, K float64
	Reader FacetReader
}
func (n *Normal) GetFromLine(r io.Reader){
	n.I, n.J, n.K = n.Reader.ReadNormal(r)
}

type Vertex struct{
	X, Y, Z float64
	Reader FacetReader
}

type FacetReader interface {
	ReadVertex(r io.Reader) (x,y,z float64)
	ReadNormal(r io.Reader) (i,j,k float64)
}

func (v *Vertex) GetFromLine(r io.Reader){
	v.X, v.Y, v.Z = v.Reader.ReadVertex(r)
}

type AsciiStlReader struct {

}
func (a AsciiStlReader) ReadVertex (r io.Reader) (x,y,z float64) {
	_, err := fmt.Fscanf(r, "  vertex %f %f %f\n", &x, &y, &z)
	if err != nil {
		panic(err)
	}
	return
}
func (a AsciiStlReader) ReadNormal(r io.Reader) (i,j,k float64){
	_, err := fmt.Fscanf(r," facet normal %f %f %f\n  outer loop\n", &i, &j, &k)
	if err != nil {
		panic(err)
	}
	return
}


type BinaryStlReader struct {

}

func (a BinaryStlReader) ReadNormal (r io.Reader) (x,y,z float64) {
	x1,y1,z1 := a.getVector(r)
	fmt.Println(x,y,z)
	return float64(x1),float64(y1),float64(z1)
}

func (a BinaryStlReader) ReadVertex (r io.Reader) (x,y,z float64) {
	x1,y1,z1 := a.getVector(r)
	fmt.Println(x,y,z)
	return float64(x1),float64(y1),float64(z1)
}

func (a BinaryStlReader) getVector(r io.Reader) (x,y,z float32) {
	//buf := make([]byte, 12)
	//n , err := r.Read(buf)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("read ", n)
	//fmt.Println(hex.Dump(buf))
	//binary.Rea
	//bits := binary.LittleEndian.Uint32(buf[0:4])
	//x = math.Float64frombits(uint64(bits))
	err :=  binary.Read(r, binary.LittleEndian, &x)
	if err != nil {
		panic(err)
	}
	err =  binary.Read(r, binary.LittleEndian, &y)
	if err != nil {
		panic(err)
	}
	err =  binary.Read(r, binary.LittleEndian, &z)
	if err != nil {
		panic(err)
	}
	fmt.Println(x,y,z)
	return
}

type Model struct {

}

func main(){
	//readAscii()
	readBinary()
}

func readAscii(){
	file, err := os.Open("gopher_ascii.stl")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	line, isPrefix, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}
	t := Triangle{
		N: Normal{Reader:AsciiStlReader{}},
		V:[3]Vertex{Vertex{Reader:AsciiStlReader{}},Vertex{Reader:AsciiStlReader{}}, Vertex{Reader:AsciiStlReader{}}},
	}
	//line, _, _ = reader.ReadLine()
	t.N.GetFromLine(reader)
	t.V[0].GetFromLine(reader)
	t.V[1].GetFromLine(reader)
	t.V[2].GetFromLine(reader)
	fmt.Printf("%s %t \n",line, isPrefix)
	fmt.Println(t)
}

func readBinary(){
	file, err := os.Open("gopher.stl")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 80)
	reader := bufio.NewReader(file)
	c, err := file.Read(buf)
	var size  uint32
	binary.Read(reader, binary.LittleEndian, &size)
	fmt.Println("size ", size)
	fmt.Println("read rr", c,err, string(buf))
	t := Triangle{
		N: Normal{Reader:BinaryStlReader{}},
		V:[3]Vertex{Vertex{Reader:BinaryStlReader{}},Vertex{Reader:BinaryStlReader{}}, Vertex{Reader:BinaryStlReader{}}},
	}
	t.N.GetFromLine(reader)
	t.V[0].GetFromLine(reader)
	t.V[1].GetFromLine(reader)
	t.V[2].GetFromLine(reader)
	fmt.Println(t)
}
