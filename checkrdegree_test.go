package main

import (
	"testing"
	"fmt"
)


func TestRDegreeLb(t *testing.T){
	rDegree,_:=checkRDegree(-1.0)
	if rDegree !="Très forte à parfaite" {
		t.Errorf("Expected : %v Actual : %v ", "Très forte à parfaite",rDegree)
	}else{
		fmt.Print("lb test passed ")
	}
}

func TestRDegreeHb(t *testing.T){
	rDegree,_:=checkRDegree(1.0)
	if rDegree !="Très forte à parfaite" {
		t.Errorf("Expected : %v Actual : %v ", "Très forte à parfaite",rDegree)
	}else{
		fmt.Print("lb test passed ")
	}
}

func TestRDegreeInvalid(t *testing.T){

	_,err:=checkRDegree(1000)
	if err == nil {
		t.Errorf("No error hence test case failed")
	}

}
