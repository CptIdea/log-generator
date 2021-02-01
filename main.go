package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"strconv"
)

var count int
var file string
var clear bool

func init() {
	flag.IntVar(&count,"c",1000,"count of logs")
	flag.StringVar(&file,"f","log.log","Log file path")
	flag.BoolVar(&clear,"clear",false,"clear file")

	flag.Parse()
}

func main() {
	LogFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	if clear{
		LogFile.Truncate(0)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(LogFile)

	for i := 0;i<count ; i++ {
		log.Info("Test " + strconv.Itoa(i))
		fmt.Println("Test " + strconv.Itoa(i))
		rand.Seed(int64(i))
		if rand.Intn(5) == 3 {
			i++
			log.Warn("Warn " + strconv.Itoa(i))
			fmt.Println("Warn " + strconv.Itoa(i))
		}
	}
}
