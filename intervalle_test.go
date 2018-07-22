package main

import (
	"testing"
	"math"
	"fmt"
)

func TestInvervalleHb(t *testing.T){

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

	varianceWithRegression,_ := VarianceWithRegression(HighBoundValue, b0,b1)

	intervalle,_ := Intervalle(1119,math.Sqrt(varianceWithRegression),1.86,HighBoundValue,averageX)

	if math.Abs(intervalle- 7.254300201324446e+08  ) > 0.001 /*Incertitude car comparaison de float*/  {
		t.Errorf("Expected : %v Actual : %v ", 1.0078294690643595e+17,intervalle)
	}else{
		fmt.Print("lb test passed ")
	}
}



func TestInvervalleLb(t *testing.T){

	HighBoundValue := [][]float64{
		{-546546543.0,-123123123.0},
		{-246546543.0,-323123123.0},
		{-346546543.0,-423123123.0},
		{-146546543.0,-523123123.0},
		{-346546543.0,-523123123.0}}

	var averageX,averageY float64;
	for i := 0; i < len(HighBoundValue); i++ {
		averageX += HighBoundValue[i][0]
		averageY += HighBoundValue[i][1]
	}

	b1,_ :=  GenerateB1(HighBoundValue, averageX, averageY)
	b0,_ :=  GenerateB0(b1, averageX, averageY)


	varianceWithRegression,_ := VarianceWithRegression(HighBoundValue, b0,b1)

	intervalle,_ := Intervalle(1119,math.Sqrt(varianceWithRegression),1.86,HighBoundValue,averageX)

	if math.Abs(intervalle- 7.254302239094731e+08  ) > 0.001 /*Incertitude car comparaison de float*/  {
		t.Errorf("Expected : %v Actual : %v ", 1.0078294690643595e+17,intervalle)
	}else{
		fmt.Print("lb test passed ")
	}
}





func TestInvervalleInvalid(t *testing.T){

	invalidValues := [][]float64{}// division by zero
	_,err := Intervalle(1119,0,1.86,invalidValues,0)

	if err == nil {
		t.Errorf("No error hence test case failed")
	}
}





