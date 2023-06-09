package main

import "fmt"

type celsius float64
type fahrenheit float64
type kelvin float64

func (c celsius) celsius2F(entry float64) fahrenheit {
	return fahrenheit(entry*1.8 + 32)
}

func (c celsius) celsius2Kelvin(entry float64) kelvin {
	return kelvin(entry + 273.5)
}

func (f fahrenheit) fahrenheit2Kelvin(entry float64) kelvin {
	return kelvin((entry + 459.67) * (5.0 / 9))
}

func (f fahrenheit) fahrenheit2Celsius(entry float64) celsius {
	return celsius((entry - 32) / 1.8)
}

func (k kelvin) kelvin2F(entry float64) fahrenheit {
	return fahrenheit((entry * (9.0 / 5)) - 459.67)
}

func (k kelvin) kelvin2Celsius(entry float64) celsius {
	return celsius(entry - 273.15)
}

func main() {

	var c1, c2 celsius
	var f1, f2 fahrenheit
	var k1, k2 kelvin

	fmt.Println("Working....")
	fmt.Println("1,000 kelvin to fahrenheit = ", k1.kelvin2F(1000))
	fmt.Println("1,000 kelvin to celsius = ", k2.kelvin2Celsius(1000))
	fmt.Println("1,000 fahrenheit to kelvin = ", f1.fahrenheit2Kelvin(1000))
	fmt.Println("1,000 fahrenheit to celsius = ", f2.fahrenheit2Celsius(1000))
	fmt.Println("1,000 celsius to kelvin = ", c1.celsius2Kelvin(1000))
	fmt.Println("1,000 celsius to fahrenheit = ", c2.celsius2F(1000))
}
