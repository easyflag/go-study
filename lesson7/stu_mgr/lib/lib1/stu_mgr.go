package lib1

import (
	"encoding/json"
	"fmt"

	"github.com/go-study/lesson7/stu_mgr/lib/lib2"
)

//StuMgr xxx
type StuMgr struct {
	list []lib2.Student
}

//StudentManage xxx
var StudentManage StuMgr

//Init xxx
func init() {
	StudentManage.list = make([]lib2.Student, 0, 1)
}

//AddStudent xxx
func (s *StuMgr) AddStudent() {
	stu := new(lib2.Student)
	stu.AddInfo()
	s.list = append(s.list, *stu)
}

//AddStudentV2 xxx
func (s *StuMgr) AddStudentV2(name string, age uint8, sex string, grade string, score uint8) {
	stu := new(lib2.Student)
	stu.Name = name
	stu.Age = age
	stu.Sex = sex
	stu.Grade = grade
	stu.Score = score
	s.list = append(s.list, *stu)
}

//ShowStudent xxx
func (s *StuMgr) ShowStudent() {
	str, err := json.Marshal(s.list)
	if err != nil {
		fmt.Println("json marshal failed")
	}
	fmt.Printf("%s\n", string(str))
	fmt.Printf("%#v\n", *s)
}

//DeleteStudent xxx
func (s *StuMgr) DeleteStudent() {
	var stu string
	var deleteFlag bool
	fmt.Println("please input name")
	fmt.Scanf("%s\n", &stu)
	for i, v := range s.list {
		if stu == v.Name {
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
