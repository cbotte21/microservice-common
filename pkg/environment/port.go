package environment

import (
	"log"
	"strconv"
)

func GetPort() int {
	port, err := strconv.Atoi(GetEnvVariable("port"))
	if err != nil {
		log.Fatalf("Could not parse port variable!")
	}
	return port
}
