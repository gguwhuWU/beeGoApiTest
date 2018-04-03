package models

import (
	"errors"
	"strings"
)

var (
	ClassList map[string]*Class
)

type Class struct {
	ClassId   string
	ClassName string
}

func init() {
	ClassList = make(map[string]*Class)
	ClassList["1"] = &Class{"1", "a"}
	ClassList["2"] = &Class{"2", "b"}
	ClassList["3"] = &Class{"3", "c"}
	ClassList["4"] = &Class{"4", "d"}
}

func GetClassById(classId string) (class *Class, err error) {
	if v, ok := ClassList[classId]; ok {
		return v, nil
	}
	return nil, errors.New("classId Not Exist")
}

func GetClassByName(className string) (class *Class, err error) {

	for key, value := range ClassList {

		if strings.EqualFold(value.ClassName, className) {
			return ClassList[key], nil
		}
	}

	return nil, errors.New("className Not Exist")
}

func GetAllClasses() map[string]*Class {
	return ClassList
}

func AddNewClass(class *Class) string {
	ClassList[class.ClassId] = class

	return "00000" + class.ClassId
}
