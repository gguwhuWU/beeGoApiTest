package models

import (
	"errors"
)

var (
	StudentList map[string]*Student
)

type Student struct {
	StudentId   string
	StudentName string
	ClassName   string
}

func init() {
	StudentList = make(map[string]*Student)
	StudentList["s201700001"] = &Student{"s201700001", "Sharon", "apple"}
	StudentList["s201700002"] = &Student{"s201700002", "Summer", "apple"}
	StudentList["70915"] = &Student{"70915", "fff", "agggpple"}
	StudentList["1"] = &Student{"1", "jojo", "agggpple"}
}

func GetStudent(studentId string) (student *Student, err error) {
	if v, ok := StudentList[studentId]; ok {
		return v, nil
	}
	return nil, errors.New("studentId Not Exist")
}

func GetAllStudents() map[string]*Student {
	return StudentList
}

func AddNewStudent(student *Student) string {
	StudentList[student.StudentId] = student

	return "1111" + student.StudentId
}

func DeleteStudent(sId string) {
	delete(StudentList, sId)
}
