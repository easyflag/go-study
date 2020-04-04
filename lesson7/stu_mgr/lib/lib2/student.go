package lib2

import "fmt"

//Person xxx
type person struct {
	Name string
	Age  uint8
	Sex  string
}

//Student xxx
type Student struct {
	Grade string
	Score uint8
	person
}

//AddInfo xxx
func (s *Student) AddInfo() {
	fmt.Println("please input name")
	fmt.Scanf("%s\n", &s.Name)
	fmt.Println("please input age")
	fmt.Scanf("%d\n", &s.Age)
	fmt.Println("please input sex")
	fmt.Scanf("%s\n", &s.Sex)
	fmt.Println("please input grade")
	fmt.Scanf("%s\n", &s.Grade)
	fmt.Println("please input score")
	fmt.Scanf("%d\n", &s.Score)
}

//ShowInfo xxx
func (s *Student) ShowInfo() {
	fmt.Printf("%v\n", *s)
}
