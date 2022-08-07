package math

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

func TestMedia(t *testing.T) {
	t.Parallel()
	valorEsperado := 10.0
	valor := Media(10.0, 10.7)
	// t.Logf(erroPadrao, valor, valorEssperado)
	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}
