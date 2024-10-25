package train

import (
	"fmt"
	"time"
)

func expandTrainRouteName(abbrRouteName string) string {
	switch abbrRouteName {
	case "P":
		return "Purple"
	case "Y":
		return "Yellow"
	case "Blue":
		return "Blue"
	case "Pink":
		return "Pink"
	case "G":
		return "Green"
	case "Org":
		return "Orange"
	case "Brn":
		return "Brown"
	case "Red":
		return "Red"
	default:
		return "UNKNOWNLINE"
	}
}

func ConvertCTATime(timeString string) time.Time {
	//Golang time: Mon Jan 2 15:04:05 MST 2006
	//CTA time example: 20240415 10:46:42
	loc, nil := time.LoadLocation("America/Chicago")
	timeFormat := "20060102 15:04:05"
	parsedTime, err := time.ParseInLocation(timeFormat, timeString, loc)
	if err != nil {
		fmt.Printf("This shouldn't happen. Error parsing CTA time: %s\n", err.Error())
	}
	return parsedTime
}
