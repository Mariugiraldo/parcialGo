package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {
	res, err := os.ReadFile("./tickets.csv")
	if err != nil {
		fmt.Println("Error en la lectura del archivo")
		return

	}
	data := strings.Split(string(res), ",")

	fmt.Println(data)

	total, err := tickets.GetTotalTickets("Grecia")
	if err != nil {
		fmt.Println("Error al obtener el total de tickets:", err)
		return

	}
	fmt.Println("Total de tickets para Grecia:", total)

}
