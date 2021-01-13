package main

import "testing"

func TestCalculate(t *testing.T){
	 if calculate(2) != 4{
		 t.Error("expcted is 4")
	 }
}