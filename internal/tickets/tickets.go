package tickets

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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
		return nil
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

func GetTotalTickets(destination string) (int, error) {
	var ticketList = loadTickets()

	var total int

	for _, ticket := range ticketList {
		if strings.EqualFold(ticket.DestinationCountry, destination) {
			total++
		}

	}
	fmt.Println("Total de tickets para " + destination,  total)


	return total, nil
}

// ejemplo 2
func GetMornings(time string) (int, error) {
	return 1, nil
}

// ejemplo 3
func AverageDestination(destination string, total int) (int, error) {
	return 1, nil
}
