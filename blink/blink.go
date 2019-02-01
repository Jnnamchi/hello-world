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

// InitalizePinAddress opens the memory range for GPIO access
func InitalizePinAddress() {

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println("Got error:", err)
		os.Exit(1)
	}
}

// Blink triggers the LED mapped at the specified pin (depending on the
// desired color) to shine for 200 Milliseconds.
func Blink(color string) {

	// Set pin according to desired color
	pin := rpio.Pin(colorMapping[color])

	// Set pin to output mode
	pin.Output()

	// Toggle once for 100 milliseconds
	pin.High()
	time.Sleep(time.Millisecond * 200)
	pin.Low()
}
