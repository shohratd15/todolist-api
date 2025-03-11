package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func InitLogger(){
	
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	Log.SetLevel(logrus.InfoLevel)

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Log.Fatal("Failed to open log file", err)
	}

	Log.SetOutput(file)
}