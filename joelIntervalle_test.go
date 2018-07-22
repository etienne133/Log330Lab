package main

import (
	"testing"
	"math"
)

func TestCTJoel1(t *testing.T){
	testValue := [][]float64{{1.3434,3.423543},{2.4332,4.111},{2.2222,6.666},{1.111,9.9999},{3.333,3.3333}}

	var avgX,avgY float64
	for i := 0; i < len(testValue); i++ {
		avgX += testValue[i][0]
		avgY += testValue[i][1]
	}

	b1,_ :=  GenerateB1(testValue, avgX, avgY)
	b0,_ :=  GenerateB0(b1, avgX, avgY)

	varianceWithRegression,_ := VarianceWithRegression(testValue, b0,b1)

	interval,_ := Intervalle(1119,math.Sqrt(varianceWithRegression),1.86,testValue,avgX)

	if toFixed(interval, 4) != 507.0060 {
		t.Errorf("Expected 507.0059 but got %v instead", interval)
	}
}

func TestCTJoel2(t *testing.T){
	testValue := [][]float64{{99.9999,111.11111},{22222.4332,3333.111},{44444.2222,5555.666},
		{55555.111,666.9999},{5.333,1.3333}}

	var avgX,avgY float64;
	for i := 0; i < len(testValue); i++ {
		avgX += testValue[i][0]
		avgY += testValue[i][1]
	}

	b1,_ :=  GenerateB1(testValue, avgX, avgY)
	b0,_ :=  GenerateB0(b1, avgX, avgY)

	varianceWithRegression,_ := VarianceWithRegression(testValue, b0,b1)

	interval,_ := Intervalle(1119,math.Sqrt(varianceWithRegression),1.86,testValue,avgX)


	if toFixed(interval, 4) != 5130.3624 {
		t.Errorf("Expected 5130.3624but got %v instead", interval)
	}
}

func TestCTJoel3(t *testing.T){
	testValue := [][]float64{{1,1},{1,1}}

	_,err := Intervalle(1119,0,1.86,testValue,1)


	if err == nil {
		t.Errorf("No division by zero hence the test case failed ")
	}
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num * output)) / output
}