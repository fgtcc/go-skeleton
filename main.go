package main

import (
	"go-skeleton/model"
	"go-skeleton/routes"
	"go-skeleton/utils"
	"go-skeleton/utils/log"
)

func init() {
	log.InitLogger()
}

func main() {
	welcomeStr := `
   ____            ____  _        _      _              
  / ___| ___      / ___|| | _____| | ___| |_ ___  _ __  
 | |  _ / _ \ ____\___ \| |/ / _ \ |/ _ \ __/ _ \| '_ \ 
 | |_| | (_) |_____|__) |   <  __/ |  __/ || (_) | | | |
  \____|\___/     |____/|_|\_\___|_|\___|\__\___/|_| |_|
	`

	log.Infof(welcomeStr)
	model.InitDb()
	defer model.CloseDb()

	log.Infof("serving is running at port%v", utils.HttpPort)
	r := routes.NewRouter()
	r.Run(utils.HttpPort)
}
