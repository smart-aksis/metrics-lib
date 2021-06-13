package logger

import (
	"log"
	"os"
)

var (
	Warn *log.Logger
	Info    *log.Logger
	Error   *log.Logger
	exit = make(chan bool)
)

func init() {
	//now:=time.Now()
	//patternLog:=fmt.Sprintf("%d_%d_%d_logs.txt", now.Year(), now.Month(), now.Day())
	//file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//File
	//io.MultiWriter(os.Stdout, file)

	Info = log.New(os.Stdout, " INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(os.Stdout, " WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

}



