package main

import (
	"testing"
)


func TestStudentToTabValide(t *testing.T){
	students:= []student{
		{"abc",[6]float64{1,2,3,4,5,6},11.0},
	}

	tab,_ := StudentToTab(1,students)

	if students[0].note != tab[0][1] {
		t.Errorf("Expected : %v Actual : %v ", 1,students[0].note)
	}
}

func TestStudentToTabInvalide(t *testing.T){
	students:= []student{
		{"abc",[6]float64{1,2,3,4,5,6},11.0},
	}

	_,err := StudentToTab(-1,students)

	if err == nil {
		t.Errorf("No error hence test case failed")
	}

}