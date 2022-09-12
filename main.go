package main

import (
	"fmt"
	"time"
)

const (
	ID         = "Asia/Jakarta"
	UTC        = "UTC"
	timeFormat = "2022-09-12T06:05:32.167Z"
)

func main() {

	timeNow := time.Now().UTC()
	fmt.Println("timeNow: ", timeNow.Format(time.RFC3339Nano))

	timeExample := timeNow.Format(time.RFC3339Nano)

	timeID := ConvertToID(timeExample)
	timeUTC := ConvertToUTC(timeID)

	fmt.Println(timeID)
	fmt.Println(timeUTC)
}

func ConvertToID(timeString string) string {
	var outputTime string

	timeFormatID := "2006-01-02T15:04:05.999Z07:00"

	fmt.Println("timeString: ", timeString)

	location, _ := time.LoadLocation(ID)

	parse, _ := time.Parse(timeFormatID, timeString)

	parseLocal := parse.In(location)
	fmt.Println("parseLocal: ", parseLocal)

	outputTime = parseLocal.Format(timeFormatID)

	return outputTime
}

func ConvertToUTC(timeString string) string {
	var outputTime string

	timeFormatID := time.RFC3339Nano

	fmt.Println("timeString: ", timeString)

	location, _ := time.LoadLocation(UTC)

	parse, _ := time.Parse(timeFormatID, timeString)

	parseLocal := parse.In(location)
	fmt.Println("parseLocal: ", parseLocal)

	outputTime = parseLocal.Format(timeFormatID)

	return outputTime
}
