package main

import (
	"testing"
	"math"
	"fmt"
)

func TestCalculateInTermsOfYHb(t *testing.T){

	x,_ :=CalculateInTermsOfY(31290832901283.0,3123212231.0,3213221312321.0)
	if math.Abs(x- -8989.978750170309) > 0.00001 /*Incertitude car comparaison de float*/  {
		t.Errorf("Expected : %v Actual : %v ", 1.003557213484165e+22,x)
	}else{
		fmt.Print("lb test passed ")
	}
}

func TestCalculateInTermsOfYLb(t *testing.T){

	x,_:=CalculateInTermsOfY(-3421233123312312,-2312312312312312,-12312312312312)
	if math.Abs(x- -1.4742475715103898 ) > 0.00001 /*Incertitude car comparaison de float*/  {
		t.Errorf("Expected : %v Actual : %v ", 2.8469911352790087e+28 ,x)
	}else{
		fmt.Print("lb test passed ")
	}
}

func TestYInvalid(t *testing.T){

	_,err := CalculateInTermsOfY(0,0,0) // division by 0

	if err == nil {
		t.Errorf("No error hence test case failed")
	}

}

