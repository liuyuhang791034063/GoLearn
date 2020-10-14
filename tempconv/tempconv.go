package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type Kelvins float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC 	  Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9.5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CToK(c Celsius) Kelvins {
	return Kelvins(c - AbsoluteZeroC)
}

func KToC(k Kelvins) Celsius {
	return Celsius(k + Kelvins(AbsoluteZeroC))
}

type celsiusFlag struct { Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	_, err := fmt.Sscanf(s, "%f%s", &value, &unit)
	if err != nil {

	}
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = KToC(Kelvins(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}


