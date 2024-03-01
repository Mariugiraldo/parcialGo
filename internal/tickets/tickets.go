package tickets

import (
	"os"
	"strings"
)

type Ticket struct {
	id                 string
	name               string
	Email              string
	destinationCountry string
	flightTime         string
	price              int
}

// ejemplo 1*/
func GetTotalTickets(destination string) (int, error) {
	data, err := os.ReadFile("./tickets.csv")
	if err != nil {
		return 0, err
	}

	lines := strings.Split(string(data), "\n")
	var total int

	for _, line := range lines {
		fields := strings.Split(line, ",")
		if len(fields) > 3 {

			ticket := Ticket{
				id:                 fields[0],
				name:               fields[1],
				Email:              fields[2],
				destinationCountry: fields[3],
				flightTime:         fields[4],
				price:              0,
			}

			if strings.EqualFold(ticket.destinationCountry, destination) {
				total++
			}
		}
	}

	return total, nil
}

/*// ejemplo 2
func GetMornings(time string) (int, error) {}*/
//ejemplo 3
/*func AverageDestination(destination string, total int) (int, error) {}*/
