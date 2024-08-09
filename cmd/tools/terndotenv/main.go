package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os/exec"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cmd := exec.Command("tern", "migrate", "--migrations", "./internal/store/pgstore/migrations", "--config", "./internal/store/pgstore/migrations/tern.conf")

	output, err := cmd.CombinedOutput() // Captura stdout e stderr
	if err != nil {
		fmt.Printf("Erro ao executar o comando: %v\n", err)
		fmt.Printf("Saída do comando: %s\n", output)
		panic(err)
	}
}
