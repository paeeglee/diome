package main

import "testing"

func ShouldSumCorrect(t *testing.T) {
	value := Soma(2, 2)
	result := 4.

	if value != result {
		t.Error("experado: ", result, "recebido: ", value)
	}
}

func ShouldDecreaseCorrect(t *testing.T) {
	value := Subtração(2, 2)
	result := 0.

	if value != result {
		t.Error("experado: ", result, "recebido: ", value)
	}
}

func ShouldDivideCorrect(t *testing.T) {
	value := Divisao(2, 2)
	result := 4.

	if value != result {
		t.Error("experado: ", result, "recebido: ", value)
	}
}

func ShouldMultiplyCorrect(t *testing.T) {
	value := Multiplicacao(2, 2)
	result := 4.

	if value != result {
		t.Error("experado: ", result, "recebido: ", value)
	}
}
