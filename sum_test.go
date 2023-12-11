package main

import "testing"

func TestSum(t *testing.T) {

	total := sum(2, 13)

	if total != 15 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 15)
	}
}