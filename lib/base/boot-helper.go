package base

import (
	"flag"
	"log"

	"bp_service/lib/structs"
)

func BootUp() *structs.Depend {

	dependencies := &structs.Depend{}

	configFile := flag.String("config", "config/local.yml", "Config file name")

	if len(*configFile) == 0 {
		log.Fatal("Config File is not provided", *configFile)
	}

	loadConfig(dependencies, *configFile)
	createConnections(dependencies)
	return dependencies
}
