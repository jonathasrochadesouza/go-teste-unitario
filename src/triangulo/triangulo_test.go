package triangulo

import (
	"testing"
)

const mensagemErroTriangulo = "Esperado o triangulo %v, mas, o triangulo encontrado é %v."
const mensagemErroCalc = "O valor esperado é %v, mas, o valor recebido foi %v."

/*
Teste referente ao tipo do triangulo passado como argumento. Utilizado dataset para realizar trecho/bloco de teste.
*/
func TestTriangulo(t *testing.T) {
	t.Parallel()

	testes := []struct {
		triangulo string
		esperado  string
	}{
		{Triangulo(10, 10, 10), "equilátero"},
		{Triangulo(10, 10, 5), "isósceles"},
		{Triangulo(9, 10, 11), "escaleno"},
		{Triangulo(10, 10, 21), "inválido"},
	}

	for _, teste := range testes {
		t.Logf("Teste realizado: %v", teste)

		if teste.triangulo != teste.esperado {
			t.Errorf(mensagemErroTriangulo, teste.esperado, teste.triangulo)
		}
	}
}

/*
Teste referente ao perímetro do triangulo passado como argumento.
*/
func TestPerimetroTriangulo(t *testing.T) {
	t.Parallel()

	valor := PerimetroTriangulo(10, 15, 12)
	valorEsperado := 37

	if valor != valorEsperado {
		t.Errorf(mensagemErroCalc, valorEsperado, valor)
	}
}
