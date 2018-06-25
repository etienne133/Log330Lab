package main

import (
	"testing"
	"math"
	"fmt"
)

func TestB0Hb(t *testing.T){


	b0,_:=GenerateB0(2138732347.0,12312312312.0, 3213123123.0)
	if math.Abs(b0- -2.633274060482763e+19 ) > 0.00001 /*Incertitude car comparaison de float*/  {
		t.Errorf("Expected : %v Actual : %v ", -2.633274060482763e+19 ,b0)
	}else{
		fmt.Print("lb test passed ")
	}
}

func TestB0Lb(t *testing.T){


	b1,_:=GenerateB0(-231123123123,-326546543, -383123123)
	if math.Abs(b1- -7.547245686356214e+19 ) > 0.00001 /*Incertitude car comparaison de float*/  {
		t.Errorf("Expected : %v Actual : %v ", -7.547245686356214e+19 ,b1)
	}else{
		fmt.Print("lb test passed ")
	}
}


func TestB0Invalid(t *testing.T){

	invalidValues := [][]float64{{0.0,0.0},{0.0,0.0},{0.0,0.0},{0.0,0.0}}
	_,err := GenerateB1(invalidValues,0,0) // division by 0

	if err == nil {
		t.Errorf("No error hence test case failed")
	}

}

