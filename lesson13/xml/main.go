package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

func main() {
	//students := new(Students)
	/*for i := 0; i < 10; i++ {
		stu := &Student{
			Person: Person{
				Name: fmt.Sprintf("stu%d", rand.Intn(100)),
				Age:  uint8(18 + rand.Intn(3)),
				Sex:  "male",
			},
			Class: "1554",
			Score: uint8(30 + rand.Intn(50)),
		}
		students.Student = append(students.Student, stu)
	}*/

	//write(students)

	c := new(CursorConfig)
	read(c)
}

func write(c *CursorConfig) {
	data, err := xml.Marshal(c)
	if err != nil {
		fmt.Printf("marshal failed,error:%v", err)
	}
	err = ioutil.WriteFile("./Cursor.xml", data, 0755)
	if err != nil {
		fmt.Printf("write failed,error:%v", err)
	}
}

func read(c *CursorConfig) {
	data, err := ioutil.ReadFile("./Cursor.xml")
	if err != nil {
		fmt.Printf("read failed,error:%v", err)
	}
	err = xml.Unmarshal(data, c)
	if err != nil {
		fmt.Printf("Unmarshal failed,error:%v", err)
	}

	/*for _, v := range students.Student {
		fmt.Printf("%#v\n", *v)
	}*/

}

//Students xxx
type Students struct {
	Student []*Student
}

//Student xxx
type Student struct {
	Person
	Class string
	Score uint8
}

//Person xxx
type Person struct {
	Name string
	Age  uint8
	Sex  string
}

//CursorConfig xxx
type CursorConfig struct {
	StandardCursor StandardCursor
	CustomCursor   CustomCursor
}

//StandardCursor xxx
type StandardCursor struct {
	Cursor []Cursor
}

//CustomCursor xxx
type CustomCursor struct {
	Cursor []Cursor
}

//Cursor xxx
type Cursor struct {
	HotSpotY string `xml:"HotSpotY,attr"`
	HotSpotX string `xml:"HotSpotX,attr"`
	Path     string `xml:"Path,attr"`
	Name     string `xml:"Name,attr"`
}
