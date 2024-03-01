package tickets

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Ticket struct {
	Id                 string
	Name               string
	Email              string
	DestinationCountry string
	FlightTime         string
	Price              int
}

func loadTickets() []Ticket {
	data, err := os.ReadFile("./tickets.csv")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	var ticketList []Ticket

	for _, line := range lines {
		fields := strings.Split(line, ",")
		if len(fields) == 6 {

			i, err := strconv.Atoi(fields[5])
			if err != nil {
				fmt.Println("Error al convertir el precio:", err)
				return nil
			}
			var ticket = Ticket{
				Id:                 fields[0],
				Name:               fields[1],
				Email:              fields[2],
				DestinationCountry: fields[3],
				FlightTime:         fields[4],
				Price:              i,
			}
			ticketList = append(ticketList, ticket)
		}

	}

	return ticketList

}

// Requerimiento 1: Una función que calcule cuántas personas viajan a un país determinado.
func GetTotalTickets(destination string) (int, error) {
	var ticketList = loadTickets()

	var total int

	for _, ticket := range ticketList {
		if strings.EqualFold(ticket.DestinationCountry, destination) {
			total++
		}

	}
	fmt.Println("Total de tickets para "+destination, total)

	return total, nil
}

// Requerimiento 2: Una o varias funciones que calculen cuántas personas viajan en madrugada (0 → 6), mañana (7 → 12), tarde (13 → 19), y noche (20 → 23).
func GetPassengersByTimeOfDay() map[string]int {

	ticketList := loadTickets()

	var madrugadaCount, mañanaCount, tardeCount, nocheCount int

	for _, ticket := range ticketList {
		flightTime, err := time.Parse("15:04", ticket.FlightTime)
		if err != nil {
			fmt.Println("Error al analizar la hora de vuelo:", err)
			continue
		}

		hour := flightTime.Hour()
		if hour >= 0 && hour < 6 {
			madrugadaCount++
		} else if hour >= 6 && hour < 12 {
			mañanaCount++
		} else if hour >= 12 && hour < 20 {
			tardeCount++
		} else {
			nocheCount++
		}
	}

	passengerCountByTimeOfDay := map[string]int{
		"Madrugada": madrugadaCount,
		"Mañana":    mañanaCount,
		"Tarde":     tardeCount,
		"Noche":     nocheCount,
	}

	fmt.Println("Cantidad de pasajeros por franja horaria:")
	for key, value := range passengerCountByTimeOfDay {
		fmt.Printf("%s: %d\n", key, value)
	}

	return passengerCountByTimeOfDay
}

// Requerimiento 3: Calcular el porcentaje de personas que viajan a un país determinado en un día.
func GetAverageDestination(destination string) (float64, error) {
	ticketList := loadTickets()

	DestinationTickets := 0
	TotalTickets := 0

	for _, ticket := range ticketList {
		if strings.EqualFold(ticket.DestinationCountry, destination) {
			DestinationTickets++
		}
		TotalTickets++
	}

	if TotalTickets == 0 {
		return 0, fmt.Errorf("No se encontraron tickets")
	}

	average := (float64(DestinationTickets) / float64(TotalTickets)) * 100

	fmt.Printf("Porcentaje de personas que viajan a %s: %.2f%%\n", destination, average)

	return average, nil
}
