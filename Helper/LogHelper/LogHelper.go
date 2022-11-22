package LogHelper

import (
	"fmt"
	"log"
)

func LogInformation(message string) {
	log.Print(message)
}

func FormatErrorMessage(location string, err interface{}) string {
	return fmt.Sprintln(location, fmt.Sprint(err))
}
