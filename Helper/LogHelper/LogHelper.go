package LogHelper

import (
	"fmt"
	"log"
)

func LogInformation(message string) {
	log.Println(message)
}

func LogFatal(message string) {
	log.Fatalln("[Error] " + message)
}
func FormatErrorMessage(location string, err interface{}) string {
	return fmt.Sprintf("%s, %s", location, fmt.Sprint(err))
}

func LogFatalAndFormatErrorMessage(location string, err interface{}) {
	LogFatal(FormatErrorMessage(location, err))
}
