package environment

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var inited = false

func initEnv() {
	if !inited {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("could not load .env file, please create one!")
		}
		inited = true
	}
}

func GetEnvVariable(name string) string {
	initEnv()

	envVar, uriPresent := os.LookupEnv(name)
	if !uriPresent {
		log.Fatalf("could not find {" + name + "} environment variable")
	}
	return envVar
}
