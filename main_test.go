package main

import (
	"testing"
)

func Calculate(a, b int) int {
	return a + b
}
func Div(a, b int ) int {
	return a / b

}

func TestCalculate(t *testing.T){
	a, b := 10, 20
	must := 30
	result := Calculate(a, b)

	t.Logf("hasil calculate a , b = %v", result)

	if result != must {
		t.Errorf("Salah !, Hasil yang diharapkan adalah : %v dan yang didapatkan adalah : %v", must, result)
	}

}

func TestDiv(t *testing.T) {
	a, b := 20, 10
	must := 2
	result := Div(a, b)

	t.Logf("hasil div a, b : %v", result)

	if result != must {
		t.Errorf("Salah !, Hasil yang diharapkan adalah : %v dan yang didapatkan adalah : %v", must, result)
	}
	
}