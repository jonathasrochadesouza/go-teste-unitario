package triangulo

/*
	Verifica se é um triangulo valido e qual trinagulo se refere.
	Recebe os lados do trinagulo como agumentos.
*/
func Triangulo(ladoA uint64, ladoB uint64, ladoC uint64) string {
	if ((ladoA + ladoB) <= ladoC) || ((ladoA + ladoC) <= ladoB) || ((ladoB + ladoC) <= ladoA) {
		return "inválido"
	} else if ladoA == ladoB && ladoA == ladoC {
		return "equilátero"
	} else if ladoA == ladoB || ladoA == ladoC || ladoB == ladoC {
		if ladoA < ladoB && ladoA < ladoC || ladoB < ladoA && ladoB < ladoC || ladoC < ladoA && ladoC < ladoB {
			return "isósceles"
		} else {
			return "inválido"
		}
	}

	return "escaleno"
}

/*
	Perímetro do triangulo. Recebe os lados do trinagulo como argumento.
*/
func PerimetroTriangulo(ladoA int, ladoB int, ladoC int) int {
	return ladoA + ladoB + ladoC
}
