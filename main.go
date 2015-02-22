package main

import (
	"fmt"
	"github.com/d2one/go-i3status/i3status"
)

func main() {
	fmt.Println(`{"version":1,"click_events": true}`)
	fmt.Println("[")
	b := i3status.NewBar()

	b.Add(i3status.NewDateWidget())
	b.Add(i3status.NewImapWidget())
	// b.Add(i3status.NewTimerWidget())
	// b.Add(i3status.NewPowerWidget())
	// b.Add(i3status.NewOnOffWidget())
	// b.Add(i3status.NewWlanWidget())
	// b.Add(i3status.NewWeatherWidget())
	//b.Add(i3status.NewI3statusWidget())
	//b.Add(i3status.NewEchoWidget())

	for {
		m := <-b.Output
		fmt.Println(m + ",")
	}

}
