package acceptance

import (
	"fmt"
	"log"

	"github.com/imega-teleport/auth/config"
)

func getAPIEntryPoint(command string) string {
	return fmt.Sprintf("http://app:8080/api/v1/auth/%s", command)
}

func getDSN() string {
	host, err := config.GetConfigValue("TELEPORTDB_HOST")
	if err != nil {
		log.Fatalf("Env not exists %s", err)
	}
	port, err := config.GetConfigValue("TELEPORTDB_PORT")
	if err != nil {
		log.Fatalf("Env not exists %s", err)
	}
	user, err := config.GetConfigValue("TELEPORTDB_USER")
	if err != nil {
		log.Fatalf("Env not exists %s", err)
	}
	pass, err := config.GetConfigValue("TELEPORTDB_PASS")
	if err != nil {
		log.Fatalf("Env not exists %s", err)
	}
	name, err := config.GetConfigValue("TELEPORTDB_NAME")
	if err != nil {
		log.Fatalf("Env not exists %s", err)
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", user, pass, host, port, name)
}
