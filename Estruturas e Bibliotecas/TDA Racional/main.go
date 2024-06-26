// A tarefa aqui neste problema é ler uma expressão matemática na forma de dois números Racionais (numerador / denominador) e apresentar o resultado da operação. Cada operando ou operador é separado por um espaço em branco. A sequência de cada linha que contém a expressão a ser lida é: número, caractere, número, caractere, número, caractere, número. A resposta deverá ser apresentada e posteriormente simplificada. Deverá então ser apresentado o sinal de igualdade e em seguida a resposta simplificada. No caso de não ser possível uma simplificação, deve ser apresentada a mesma resposta após o sinal de igualdade.

// Considerando N1 e D1 como numerador e denominador da primeira fração, segue a orientação de como deverá ser realizada cada uma das operações:
// Soma: (N1*D2 + N2*D1) / (D1*D2)
// Subtração: (N1*D2 - N2*D1) / (D1*D2)
// Multiplicação: (N1*N2) / (D1*D2)
// Divisão: (N1/D1) / (N2/D2), ou seja (N1*D2)/(N2*D1)
// Entrada
// A entrada contem vários casos de teste. A primeira linha de cada caso de teste contem um inteiro N (1 ≤ N ≤ 1*104), indicando a quantidade de casos de teste que devem ser lidos logo a seguir. Cada caso de teste contém um valor racional X (1 ≤ X ≤ 1000), uma operação (-, +, * ou /) e outro valor racional Y (1 ≤ Y ≤ 1000).

// Saída
// A saída consiste em um valor racional, seguido de um sinal de igualdade e outro valor racional, que é a simplificação do primeiro valor. No caso do primeiro valor não poder ser simplificado, o mesmo deve ser repetido após o sinal de igualdade.

package main

import (
	"fmt"
)

type Fraction struct {
	Numerator   int
	Denominator int
}

func mdc(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	if b == 0 {
		return a
	}
	return mdc(b, a%b)
}

func simplify(a, b int) (int, int) {
	divisor := mdc(a, b)

	// Codigo para evitar bug do GO
	if divisor == 0 {
		return a, b
	}
	return (a / divisor), (b / divisor)
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i <= n; i++ {
		var n1, d1, n2, d2 int
		var op string
		fmt.Scanf("%d / %d %s %d / %d", &n1, &d1, &op, &n2, &d2)

		var result Fraction

		switch op {
		case "+":
			result = Fraction{
				Numerator:   n1*d2 + n2*d1,
				Denominator: d1 * d2,
			}
		case "-":
			result = Fraction{
				Numerator:   n1*d2 - n2*d1,
				Denominator: d1 * d2,
			}
		case "*":
			result = Fraction{
				Numerator:   n1 * n2,
				Denominator: d1 * d2,
			}
		case "/":
			result = Fraction{
				Numerator:   n1 * d2,
				Denominator: n2 * d1,
			}
		}

		numerator, denominator := simplify(result.Numerator, result.Denominator)

		// Codigo para evitar bug do GO
		if result.Denominator != 0 {
			fmt.Printf("%v/%v = %v/%v\n", result.Numerator, result.Denominator, numerator, denominator)
		}
	}
}
