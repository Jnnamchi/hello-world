package blink

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio"
)

var colorMapping = map[string]int{
	"color1": 16,
	"color2": 12,
}

func Blink(color string) {

	// Set pin according to desired color
	pin := rpio.Pin(colorMapping[color])

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println("Got error:", err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	// Set pin to output mode
	pin.Output()

	// Toggle once for 100 milliseconds
	pin.Toggle()
	time.Sleep(time.Millisecond * 200)
	pin.Toggle()
}
