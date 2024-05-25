package Utilities

import (
	"fmt"
	"os"
)

func CheckDiretory(path string) error {
	dirPath := ".\\" + path

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {

		err := os.MkdirAll(dirPath, 0755) // 0755 é uma permissão comum para diretórios
		if err != nil {
			fmt.Printf("Erro ao criar o diretório: %v\n", err)
			return err
		}
	} else if err != nil {
		fmt.Printf("Erro ao verificar o diretório: %v\n", err)
		return err
	}
	return nil
}
