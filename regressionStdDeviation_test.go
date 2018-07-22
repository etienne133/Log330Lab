package main

import (
	"testing"
	"math"
	"fmt"
)

func TestStdRegressionHb(t *testing.T){

	HighBoundValue := [][]float64{
		{546546543.0,123123123.0},
		{246546543.0,323123123.0},
		{346546543.0,423123123.0},
		{146546543.0,523123123.0},
		{346546543.0,523123123.0}}

	var averageX,averageY float64;
	for i := 0; i < len(HighBoundValue); i++ {
		averageX += HighBoundValue[i][0]
		averageY += HighBoundValue[i][1]
	}

	b1,_ :=  GenerateB1(HighBoundValue, averageX, averageY)
	b0,_ :=  GenerateB0(b1, averageX, averageY)

	fmt.Println("B1 :", b1)
	fmt.Println("B0 :", b0)

	varianceWithRegression,_ := VarianceWithRegression(HighBoundValue, b0,b1)


	if math.Abs(varianceWithRegression- 1.0078294690643595e+17 ) > 0.001 /*Incertitude car comparaison de float*/  {
		t.Errorf("Expected : %v Actual : %v ", 1.0078294690643595e+17,varianceWithRegression)
	}else{
		fmt.Print("lb test passed ")
	}
}



func TestStdRegressionLb(t *testing.T){

	LowBoundValue := [][]float64{
		{-546546543.0,-123123123.0},
		{-246546543.0,-323123123.0},
		{-346546543.0,-423123123.0},
		{-146546543.0,-523123123.0},
		{-346546543.0,-523123123.0}}

		var averageX,averageY float64;
	for i := 0; i < len(LowBoundValue); i++ {
		averageX += LowBoundValue[i][0]
		averageY += LowBoundValue[i][1]
	}

	b1,_ :=  GenerateB1(LowBoundValue, averageX, averageY)
	b0,_ :=  GenerateB0(b1, averageX, averageY)

	fmt.Println("B1 :", b1)
	fmt.Println("B0 :", b0)

	varianceWithRegression,_ := VarianceWithRegression(LowBoundValue, b0,b1)


	if math.Abs(varianceWithRegression- 1.0078294690643595e+17 ) > 0.001 /*Incertitude car comparaison de float*/  {
		t.Errorf("Expected : %v Actual : %v ", 1.0078294690643595e+17,varianceWithRegression)
	}else{
		fmt.Print("lb test passed ")
	}
}




//Test a correlation where the denominator equals 0 (invalid division)

func TestStdRegressionInvalid(t *testing.T){

	invalidValues := [][]float64{{0.0,0.0}}
	_,err := VarianceWithRegression(invalidValues, 0,0)

	if err == nil {
		t.Errorf("No error hence test case failed")
	}
}


