package logger

import (
	"log"
	"os"
	"time"
)

var (
	errorLogFile, _ = os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
)

func init() {
	log.SetOutput(errorLogFile)
}

func LogToFile(message string) {
	log.Printf("%s | %s\n", message, time.Now().Format(time.RFC3339))
}

func LogError(message string) {
	LogToFile("[ERROR] " + message)
}

func LogInfo(message string) {
	LogToFile("[INFO] " + message)
}
