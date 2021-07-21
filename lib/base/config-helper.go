package base

import (
	"bp_service/lib/structs"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func loadConfig(dependencies *structs.Depend, fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error loading config file", err)
	}
	defer f.Close()

	var config structs.Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Error loading config file", err)
	}
	(*dependencies).Config = config
}
