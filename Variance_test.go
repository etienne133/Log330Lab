package main

import (
	"testing"
	"math"
)

var VarianceValuesHb = []int{
	166320,
	699300,
	19300,
	929810,
	780960,
	36980,
	709360,
}

var exceptedDistancePower2Hb = 900464678542.86;
var exceptedAverageHb = 477432.8571;
var exceptedVarianceHb = 150077446423.81;

var VarianceValuesLb = []int{
	-166320,
	-699300,
	-19300,
	-929810,
	-780960,
	-36980,
	-709360,
}

var exceptedDistancePower2Lb = 900464678542.86;
var exceptedAverageLb = -477432.8571;
var exceptedVarianceLb = -150077446423.81;





//test lowet bound

func TestDistancePowerTwolb(t *testing.T){
	distance := distancePowerTwo(exceptedAverageLb,VarianceValuesLb)
	if math.Abs(distance-exceptedDistancePower2Lb) > 1 {
		t.Errorf("Expected : %v Actual : %v ", exceptedDistancePower2Lb,distance)
	}
}

func TestVariancelb(t *testing.T){
	variance := variance(exceptedDistancePower2Lb,VarianceValuesLb)
	if math.Abs(exceptedVarianceLb-variance) > .1 {
		t.Errorf("Expected : %v Actual : %v ", exceptedVarianceLb,variance)
	}
}

//test higher bound
func TestDistancePowerTwoHb(t *testing.T){
	distance := distancePowerTwo(exceptedAverageHb,VarianceValuesHb)
	if math.Abs(distance-exceptedDistancePower2Hb) > 1 {
		t.Errorf("Expected : %v Actual : %v ", exceptedDistancePower2Hb,distance)
	}
}

func TestVarianceHb(t *testing.T){
	variance := variance(exceptedDistancePower2Hb,VarianceValuesHb)
	if math.Abs(exceptedVarianceHb-variance) > .1 {
		t.Errorf("Expected : %v Actual : %v ", exceptedVarianceHb,variance)
	}
}

//test invalid
func TestDistancePowerTwoInvalid(t *testing.T){
	distance := distancePowerTwo(0,VarianceValuesHb)
	if !(math.Abs(distance-exceptedDistance) > .1) {
		t.Errorf("Expected : %v Actual : %v ", exceptedDistance,distance)
	}
}

func TestVarianceInvalid(t *testing.T){
	variance := variance(0,VarianceValuesHb)
	if !(math.Abs(exceptedVariance-variance) > .1) {
		t.Errorf("Expected : %v Actual : %v ", exceptedVariance,variance)
	}
}
