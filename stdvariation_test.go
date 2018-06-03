package main

import (
	"testing"
	"math"
)

var VarianceHb = 150077446423.81;
var VarianceLb = -150077446423.81;
var InvalidVariance = 0;
var ExceptedStdVariation = 387398.3046;


func TestStdVariationHb(t *testing.T){
	stdVariantion := stdVariation(VarianceHb)
	if math.Abs(ExceptedStdVariation-stdVariantion) > 1 {
		t.Errorf("Expected : %v Actual : %v ", exceptedStdVariation,stdVariantion)
	}
}

func TestStdVariationLb(t *testing.T){
	stdVariantion := stdVariation(VarianceLb)
	if math.Abs(ExceptedStdVariation-stdVariantion) > 1 {
		t.Errorf("Expected : %v Actual : %v ", exceptedStdVariation,stdVariantion)
	}
}

func TestStdVariationInvalid(t *testing.T){ // ?
	stdVariantion := stdVariation(0)
	if ! (math.Abs(ExceptedStdVariation-stdVariantion) > 1){
		t.Errorf("Expected : %v Actual : %v ", exceptedStdVariation,stdVariantion)
	}
}