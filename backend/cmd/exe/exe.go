package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	err := os.Chdir("backend/cmd")
	if err != nil {
		fmt.Println("erreur lors du changement de répertoire:", err)
		return
	}

	cmd := exec.Command("go", "run", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("erreur lors de l'exécution de 'go run .':", err)
	}
}
