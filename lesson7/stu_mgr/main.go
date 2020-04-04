package main

import (
	lib1 "github.com/go-study/lesson7/stu_mgr/lib/lib1"
)

func main() {
	lib1.StudentManage.AddStudentV2("Jack", 12, "male", "Three", 99)
	lib1.StudentManage.AddStudentV2("Mave", 12, "male", "Three", 99)
	lib1.StudentManage.AddStudentV2("Noma", 12, "male", "Three", 99)
	lib1.StudentManage.ShowStudent()
	lib1.StudentManage.DeleteStudent()
	lib1.StudentManage.ShowStudent()
	lib1.StudentManage.DeleteStudent()
	lib1.StudentManage.ShowStudent()
	lib1.StudentManage.DeleteStudent()
	lib1.StudentManage.ShowStudent()
	lib1.StudentManage.DeleteStudent()
}
