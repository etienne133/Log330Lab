package main

import (
	"testing"
	"math"
	"fmt"
)

func TestCalculateInTermsOfXHb(t *testing.T){

	y :=CalculateInTermsOfX(31290832901283.0,3123212231.0,3213221312321.0)
	if math.Abs(y- 1.003557213484165e+22) > 0.00001 /*Incertitude car comparaison de float*/  {
		t.Errorf("Expected : %v Actual : %v ", 1.003557213484165e+22,y)
	}else{
		fmt.Print("lb test passed ")
	}
}

func TestCalculateInTermsOfXLb(t *testing.T){

	y:=CalculateInTermsOfX(-3421233123312312,-2312312312312312,-12312312312312)
	if math.Abs(y- 2.8469911352790087e+28 ) > 0.00001 /*Incertitude car comparaison de float*/  {
		t.Errorf("Expected : %v Actual : %v ", 2.8469911352790087e+28 ,y)
	}else{
		fmt.Print("lb test passed ")
	}
}



