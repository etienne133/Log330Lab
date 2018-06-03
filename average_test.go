package main

import (
	"testing"
	"math"
)

var LowerBoundValues = []int{-100000,-100000,-100000,-100000, -100000,-100000,-100000,-100000,-100000,-100000}
var exceptedLowerBound = -100000.0;

var HigherBoundValues = []int{100000,100000,100000,100000,100000,100000,100000,100000,100000,100000}
var exceptedHigherBound = 100000.0;

var InvalidMean = 10.0;


func TestMeanLb(t *testing.T){
	mean := mean(LowerBoundValues)
	if math.Abs(mean-exceptedLowerBound) > .1 {
		t.Errorf("Expected : %v Actual : %v ", exceptedMean,mean)
	}
}


func TestMeanHb(t *testing.T){
	mean := mean(HigherBoundValues)
	if math.Abs(mean-exceptedHigherBound) > .1 {
		t.Errorf("Expected : %v Actual : %v ", exceptedMean,mean)
	}
}

func TestMeanInvalid(t *testing.T){
	mean := mean(HigherBoundValues)
	if !(math.Abs(mean-InvalidMean) > .1) {
		t.Errorf("Expected : %v Actual : %v ", exceptedMean,mean)
	}
}
