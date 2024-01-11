package main

import "testing"

func SumTest(t *testing.T) {

	total := Sum(2, 2)

	if total != 4 {
		t.Errorf("Resultado inválido. Resultado: %d. Esperado: %d.", total, 4)
	}

}