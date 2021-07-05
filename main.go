package main

import (
	"flag"
	"log"
	"time"

	"github.com/kiyocy24/uuid-sim/sim"
)

func main() {
	var digit int
	flag.IntVar(&digit, "d", 5, "一致させる桁数")

	log.Printf("digit = %d\n", digit)
	sm := sim.NewSim(digit)

	log.Println("start")
	start := time.Now()
	c := sm.Run()
	end := time.Now()

	diff := end.Sub(start).Seconds()
	log.Printf("time = %fs, count = %d\n", diff, c)
}
