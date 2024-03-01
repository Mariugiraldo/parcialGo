package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {

	totalTicketsChan := make(chan int)
	passengerCountChan := make(chan map[string]int)
	averageChan := make(chan float64)

	go func() {
		total, _ := tickets.GetTotalTickets("Brazil")
		totalTicketsChan <- total
	}()
	go func() {
		passengerCount := tickets.GetPassengersByTimeOfDay()
		passengerCountChan <- passengerCount
	}()

	go func() {
		averange, _ := tickets.GetAverageDestination("Brazil")
		averageChan <- averange
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ocurrió un error:", r)
		}
	}()

	totalTickets := <-totalTicketsChan
	fmt.Printf("Total de tickets para Japan: %d\n", totalTickets)

	passengerCount := <-passengerCountChan
	fmt.Println("Cantidad de pasajeros por franja horaria:")
	for key, value := range passengerCount {
		fmt.Printf("%s: %d\n", key, value)
	}

	average := <-averageChan
	fmt.Printf("Porcentaje de personas que viajan a Brazil: %.2f%%\n", average)

	close(totalTicketsChan)
	close(passengerCountChan)
	close(averageChan)

}
