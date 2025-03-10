package infra

import (
	"fmt"
	"os"
)

func NewLogConnector() *os.File {
	q, err := os.OpenFile("logs/log.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("error log init")
	}
	return q

}


