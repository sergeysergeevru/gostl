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
	outFile, err := os.Create("out.cnc")
	defer outFile.Close()
	if err != nil {
		panic(err)
	}
	step := model.StlFractionalType(0.05)
	layers := make(map[model.StlFractionalType][]*model.PerimeterLineSegment)
	for i:= uint32(0); i < size; i++ {
		t.N.GetFromLine(reader)
		t.V[0].GetFromLine(reader)
		t.V[1].GetFromLine(reader)
		t.V[2].GetFromLine(reader)
		var byteCount uint16
		binary.Read(reader, binary.LittleEndian, &byteCount)
		fmt.Println(byteCount)
		//t.V.GetIntersection(4)
		//fmt.Println(t.GetZRange())
		segments := t.GetPerimeterSegments(step)
		for _,v := range segments {
			layers[model.StlFractionalType(v.N)*step] = append(layers[model.StlFractionalType(v.N)*step], v.Segment)
			//outFile.WriteString(fmt.Sprintf("G0 Z%0.3f F3000.000 \n", v.Z))
			//outFile.WriteString(fmt.Sprintf("G0 X%.3f Y%.3f \n", v.V[0].X,v.V[0].Y))
			//outFile.WriteString(fmt.Sprintf("G1 X%.3f Y%.3f \n", v.V[1].X,v.V[1].Y))
		}
		//fmt.Println(t)
	}
	for z, v := range layers {
		outFile.WriteString(fmt.Sprintf("G1 Z%0.3f F3000.000 \n", z))
		for _, item := range v {
			outFile.WriteString(fmt.Sprintf("G1 X%.3f Y%.3f \n", item.V[0].X, item.V[0].Y))
			outFile.WriteString(fmt.Sprintf("G1 X%.3f Y%.3f \n", item.V[1].X, item.V[1].Y))
		}
	}
}
