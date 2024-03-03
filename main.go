package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {

	totalTicketsChan := make(chan int)
	passengerCountChan := make(chan map[string]int)
	averageChan := make(chan float64)

	destination := "Utrecht"
	go func() {
		total, _ := tickets.GetTotalTickets(destination)
		totalTicketsChan <- total
	}()
	go func() {
		passengerCount := tickets.GetPassengersByTimeOfDay()
		passengerCountChan <- passengerCount
	}()

	go func() {
		averange, _ := tickets.GetAverageDestination(destination)
		averageChan <- averange
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ocurrió un error:", r)
		}
	}()

	totalTickets := <-totalTicketsChan
	fmt.Printf("Total de tickets para: %d\n", totalTickets)

	passengerCount := <-passengerCountChan
	fmt.Println("Cantidad de pasajeros por franja horaria:")
	for key, value := range passengerCount {
		fmt.Printf("%s: %d\n", key, value)
	}

	average := <-averageChan
	fmt.Printf("Porcentaje de personas que viajan a %s: %.2f%%\n", destination, average)

	close(totalTicketsChan)
	close(passengerCountChan)
	close(averageChan)

}
