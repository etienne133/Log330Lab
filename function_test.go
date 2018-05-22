package main

import (
"testing"
"math"
)
/*
10 données : 186, 699, 132, 272, 291, 331, 199, 1890, 788, 1601
● Moyenne : 638,9
● Variance = σ = 391 417,8778 2
● Écart type = σ = 625,63
*/

var values = []int{186, 699,132,272, 291,331,199,1890,788,1601}

var exceptedMean float64 = 638.9;
var exceptedVariance float64 = 391417.87782;
var exceptedDistance float64 = 3522760.9;
var exceptedStdVariation float64 = 625.63;


func TestMean(t *testing.T){
	mean := mean(values)
	if math.Abs(mean-exceptedMean) > 1 {
		t.Errorf("Expected : %v Actual : %v ", exceptedMean,mean)
	}
}

func TestDistancePowerTwo(t *testing.T){
	distance := distancePowerTwo(exceptedMean,values)
	if math.Abs(distance-exceptedDistance) > 1 {
		t.Errorf("Expected : %v Actual : %v ", exceptedDistance,distance)
	}
}

func TestVariance(t *testing.T){
	variance := variance(exceptedDistance,values)
	if math.Abs(exceptedVariance-variance) > 1 {
		t.Errorf("Expected : %v Actual : %v ", exceptedVariance,variance)
	}
}
func TestStdVariation(t *testing.T){
	stdVariantion := stdVariation(exceptedVariance)
	if math.Abs(exceptedStdVariation-stdVariantion) > 1 {
		t.Errorf("Expected : %v Actual : %v ", exceptedStdVariation,stdVariantion)
	}
}