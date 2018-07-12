package main

import (
"testing"
"math"
)

func TestJoelCT11(t *testing.T){

	testValue := [][]float64{{10.0,1.0},{1.0,10.0}}
	correlation := Correlation(testValue)

	if absolute(correlation) != 1 {
		t.Errorf("Expected 1 but got %v instead", correlation)
	}
}

func TestJoelCT2(t *testing.T){
	testValue := [][]float64{{164560,54646},{65346,23352352},{32423432, 52542}}
	correlation := Correlation(testValue)

	if absolute(correlation) < 0.5 || absolute(correlation) > 0.6 {
		t.Errorf("Expected 0.5023 but got %v instead", correlation)
	}
}

func TestJoelCT3(t *testing.T){
	testValue := [][]float64{{0,0},{0,0}}
	correlation := Correlation(testValue)

	if correlation!=1000  {
		t.Errorf("Expected NaN but got %v instead", correlation)
	}

	}


func absolute(value float64)float64{
	return math.Abs(value)
}