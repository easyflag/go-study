package main

import (
	"fmt"
)

//Person sss
type person struct {
	name string
	Age  uint8
	Sex  string
}

//Student sss
type student struct {
	Grade string
	Score uint8
	person
}

func (s *student) addInfo() {
	fmt.Println("please input name")
	fmt.Scanf("%s\n", &s.name)
	fmt.Println("please input age")
	fmt.Scanf("%d\n", &s.Age)
	fmt.Println("please input sex")
	fmt.Scanf("%s\n", &s.Sex)
	fmt.Println("please input grade")
	fmt.Scanf("%s\n", &s.Grade)
	fmt.Println("please input score")
	fmt.Scanf("%d\n", &s.Score)
}

func (s *student) showInfo() {
	fmt.Printf("%v\n", *s)
}

type studentList struct {
	list []student
}

func (s *studentList) init() {
	s.list = make([]student, 0, 1)
}

func (s *studentList) addStudent() {
	stu := new(student)
	stu.addInfo()
	s.list = append(s.list, *stu)
}

func (s *studentList) addStudentV2(name string, age uint8, sex string, grade string, score uint8) {
	stu := new(student)
	stu.name = name
	stu.Age = age
	stu.Sex = sex
	stu.Grade = grade
	stu.Score = score
	s.list = append(s.list, *stu)
}

func (s *studentList) showStudent() {
	fmt.Printf("%v\n", *s)
}

func (s *studentList) deleteStudent() {
	var stu string
	var deleteFlag bool
	fmt.Println("please input name")
	fmt.Scanf("%s\n", &stu)
	for i, v := range s.list {
		if stu == v.name {
			s.list = append(s.list[:i], s.list[i+1:]...)
			deleteFlag = true
			break
		}
	}
	if deleteFlag != true {
		fmt.Println("can't find this student")
	} else {
		fmt.Println("delete ok")
	}
}

func main() {
	stuList := new(studentList)
	stuList.init()
	stuList.addStudentV2("Jack", 12, "male", "Three", 99)
	stuList.addStudentV2("Manve", 12, "male", "Three", 99)
	stuList.addStudentV2("tyyr", 12, "male", "Three", 99)
	stuList.showStudent()
	stuList.deleteStudent()
	stuList.showStudent()
	stuList.deleteStudent()
	stuList.showStudent()
	stuList.deleteStudent()
	stuList.showStudent()
}
