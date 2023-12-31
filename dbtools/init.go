package dbtools

import (
	"log"
	"os"
)

func Init() {
	mysql := GetMysqlInstance().InitMysqlPool()
	if !mysql {
		log.Println("init database pool failure...")
		os.Exit(1)
	}
}
