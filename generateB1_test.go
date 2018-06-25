package main

import (
	"testing"
	"math"
	"fmt"
)

func TestB1Hb(t *testing.T){
	values := [][]float64{
		{546546543.0,123123123.0},
		{246546543.0,323123123.0},
		{346546543.0,423123123.0},
		{146546543.0,523123123.0},
		{346546543.0,523123123.0}}

	b1,_:=GenerateB1(values,-326546543, -383123123)
	if math.Abs(b1- 1.0934399665752959) > 0.00001 /*Incertitude car comparaison de float*/  {
		t.Errorf("Expected : %v Actual : %v ", 1.0934399665752959,b1)
	}else{
		fmt.Print("lb test passed ")
	}
}

func TestB1Lb(t *testing.T){
	values := [][]float64{
		{-546546543.0,-123123123.0},
		{-246546543.0,-323123123.0},
		{-346546543.0,-423123123.0},
		{-146546543.0,-523123123.0},
		{-346546543.0,-523123123.0}}
	b1,_:=GenerateB1(values,-326546543, -383123123)
	if math.Abs(b1- -0.8409090909090909) > 0.00001 /*Incertitude car comparaison de float*/  {
		t.Errorf("Expected : %v Actual : %v ", -0.8409090909090909,b1)
	}else{
		fmt.Print("lb test passed ")
	}
}


