package main

import (
	"errors"
	"fmt"
)

func getInfo(input string) (int, error) {
	infor := len(input)
	if infor == 0 {
		return 0, errors.New("String Vazia")
	}
	return infor, nil
}

func main() {
	valor := "Olá Mundo Gabriel!"
	tamanho, err := getInfo(valor)
	if err != nil {

		fmt.Println("Erro:", err)
	}
	fmt.Printf("Tamanho da String: %d\n", tamanho)
}
