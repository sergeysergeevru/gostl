package main

import (
	"fmt"
	"encoding/binary"
	"os"
	"bufio"
	"github.com/sergeysergeevru/gostl/model"
)

func main() {
	//readAscii()
	readBinary()
}

func readAscii() {
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
	t := model.Triangle{
		N: model.Normal{Reader: model.AsciiStlReader{}},
		V: [3]model.Vertex{model.Vertex{Reader: model.AsciiStlReader{}}, model.Vertex{Reader: model.AsciiStlReader{}}, model.Vertex{Reader: model.AsciiStlReader{}}},
	}
	//line, _, _ = reader.ReadLine()
	t.N.GetFromLine(reader)
	t.V[0].GetFromLine(reader)
	t.V[1].GetFromLine(reader)
	t.V[2].GetFromLine(reader)
	fmt.Printf("%s %t \n", line, isPrefix)
	fmt.Println(t)
}

func readBinary() {
	file, err := os.Open("gopher.stl")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 80)
	reader := bufio.NewReader(file)
	c, err := file.Read(buf)
	var size uint32
	binary.Read(reader, binary.LittleEndian, &size)
	fmt.Println("size ", size)
	fmt.Println("read rr", c, err, string(buf))
	t := model.Triangle{
		N: model.Normal{Reader: model.BinaryStlReader{}},
		V: [3]model.Vertex{model.Vertex{Reader: model.BinaryStlReader{}}, model.Vertex{Reader: model.BinaryStlReader{}}, model.Vertex{Reader: model.BinaryStlReader{}}},
	}
	t.N.GetFromLine(reader)
	t.V[0].GetFromLine(reader)
	t.V[1].GetFromLine(reader)
	t.V[2].GetFromLine(reader)
	t.V.GetIntersection(3.926045)
	fmt.Println(t)
}
