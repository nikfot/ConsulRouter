package configuration

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

//readConfiguration reads the configuration from the yaml file
func readConfiguration() (consulServers []string, port string, host string) {

	var cfg Config

	readFile(&cfg)
	log.Println(time.Now(), ": Loading configuration from application.yml")

	consulServers = cfg.ConsulServersSettings.ConsulServers
	port = cfg.RouterAPISettings.Port
	host = cfg.RouterAPISettings.Host

	log.Printf("%+v", cfg)
	log.Println("Configuration loaded...")
	return consulServers, port, host
}

//readFile reads the configuration file
func readFile(cfg *Config) {
	f, err := os.Open("application.yml")
	if err != nil {
		processError(err)
	}

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

//processError prints out errors during loading configuration
func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}
