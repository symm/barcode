package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/holoplot/go-evdev"
)

// KEYCODES is the map of hex found in the InputEvent.Code field, and
// its corresponding char (string) representation
// [source: Vojtech Pavlik (author of the Linux Input Drivers project),
// via linuxquestions.org user bricedebrignaisplage]
var KEYCODES = map[byte]string{
	0x02: "1",
	0x03: "2",
	0x04: "3",
	0x05: "4",
	0x06: "5",
	0x07: "6",
	0x08: "7",
	0x09: "8",
	0x0a: "9",
	0x0b: "0",
	0x0c: "-",
	0x10: "q",
	0x11: "w",
	0x12: "e",
	0x13: "r",
	0x14: "t",
	0x15: "y",
	0x16: "u",
	0x17: "i",
	0x18: "o",
	0x19: "p",
	0x1e: "a",
	0x1f: "s",
	0x20: "d",
	0x21: "f",
	0x22: "g",
	0x23: "h",
	0x24: "j",
	0x25: "k",
	0x26: "l",
	0x2A: "",
	0x2c: "z",
	0x2d: "x",
	0x2e: "c",
	0x2f: "v",
	0x30: "b",
	0x31: "n",
	0x32: "m",
	0x35: "_",
}

const KEY_ENTER = 0x1c
const KEY_SHIFT = 0x2a

func main() {

	d, err := evdev.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Cannot read %s: %v\n", os.Args[1], err)
		return
	}

	err = d.Grab()

	if err != nil {
		log.Fatal(err)
	}

	barcode := ""
	enableUppercase := false
	for {
		e, err := d.ReadOne()
		if err != nil {
			fmt.Printf("Error reading from device: %v\n", err)
			return
		}

		if e.Value != 1 {
			// Only process keyup events
			continue
		}

		if e.Code == KEY_SHIFT {
			enableUppercase = true
			continue
		}

		if e.Type == evdev.EV_KEY {
			character := KEYCODES[byte(e.Code)]

			if enableUppercase == true {
				character = strings.ToUpper(KEYCODES[byte(e.Code)])
				enableUppercase = false
			}

			barcode = barcode + character
		}

		if e.Code == KEY_ENTER {
			fmt.Println(barcode)
			barcode = ""
			os.Exit(0)
		}
	}
}
