package main

import (
	"log"
	"time"
)

func main() {

	timeNow := time.Now()
	christmas := time.Date(time.Now().Year(), time.December, 25, 00, 00, 00, 00, time.UTC)

	diff := christmas.Sub(timeNow)
	if diff < 0 && diff.Hours() > -24 {
		log.Println("Today is Christmas")
		//hours left today
		//	christmas = christmas.AddDate(0, 0, 1)
		//	diff = christmas.Sub(timeNow)
		//	log.Printf("today is christmas %v left", diff)
		//diff = 0
		return
	} else if diff <= -24 {
		christmas := time.Date(time.Now().Year()+1, time.December, 25, 00, 00, 00, 00, time.UTC)
		diff = christmas.Sub(timeNow)
		log.Printf("time for christmas %v", diff)
	}

	finalOutput(diff)

	//log.Println(start)
}

func finalOutput(diff time.Duration) {
	log.Printf("Milliseconds to christmas: %v", diff.Milliseconds())
	log.Printf("Seconds to christmas: %v", int64(diff.Seconds()))
	log.Printf("Minutes to christmas: %v", int64(diff.Minutes()))
	log.Printf("Hours to christmas: %v", int64(diff.Hours()))
	log.Printf("Days to christmas: %v", int64(diff.Hours()/24.0))
}
