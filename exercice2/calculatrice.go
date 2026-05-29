package main

import (
	"errors"
	"fmt"
)

func operer(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division par zéro impossible")
		}
		return a / b, nil
	default:
		return 0, errors.New("opération inconnue : " + op)
	}
}

func creerOperation(op string) func(float64, float64) float64 {
	return func(a, b float64) float64 {
		resultat, _ := operer(a, b, op)
		return resultat
	}
}

func main() {
	for {
		var a, b float64
		var op string

		fmt.Scan(&a, &b, &op)

		if op == "quit" {
			break
		}

		_, err := operer(a, b, op)
		if err != nil {
			fmt.Println(err)
		} else {
			calcul := creerOperation(op)
			resultat := calcul(a, b)
			fmt.Printf("%.2f %s %.2f = %.2f\n", a, op, b, resultat)
		}
	}
}
