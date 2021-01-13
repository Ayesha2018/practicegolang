package main

import "testing"

func TestCalculate(t *testing.T){
	 if calculate(2) != 5{
		 t.Error("expcted is 4")
	 }
}