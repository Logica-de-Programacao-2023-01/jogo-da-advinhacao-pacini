package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	jogadas := [][]int{}
	jogarNovamente := true

	for jogarNovamente {
		numero := rand.Intn(100) + 1
		tentativas := []int{}
		acertou := false

		fmt.Println("--- Jogo de Adivinhação ---")
		fmt.Println("Tente adivinhar o número entre 1 e 100.")

		for !acertou {
			var valor int
			fmt.Print("Digite um número: ")
			fmt.Scanln(&valor)
			tentativas = append(tentativas, valor)

			if valor < numero {
				fmt.Println("O número é maior que:", valor)
			} else if valor > numero {
				fmt.Println("O número é menor que:", valor)
			} else {
				acertou = true
				fmt.Println("Parabéns, você acertou!")
			}
		}

		jogadas = append(jogadas, tentativas)
		fmt.Println("Número de tentativas:", len(tentativas))

		var jogarNovamenteStr string
		fmt.Print("Deseja jogar novamente? (S/N): ")
		fmt.Scanln(&jogarNovamenteStr)

		if jogarNovamenteStr != "S" && jogarNovamenteStr != "s" {
			jogarNovamente = false
		}
	}

	fmt.Println("\n --Histórico de Jogadas-- ")
	for i, jogada := range jogadas {
		fmt.Printf("Jogada %d ( Número de tentativas: %d)\n", i+1, len(jogada))
	}
}
