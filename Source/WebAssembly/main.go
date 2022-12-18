//go:build js && wasm

package main

import (
	"strconv"
	"syscall/js"

	"github.com/littletrainee/StringToFloat"
)

var (
	arithmetic string
	setanother bool
	first      float64 = 0
	second     float64 = 0
	point      bool
	setsecond  bool
)

// first parameter not use can ignore, second parameter is the incoming parameters can't ignore
func Print(_ js.Value, values []js.Value) any {
	var (
		value       js.Value = values[0]
		ResultLabel js.Value = js.Global().Get("document").Call("getElementById", "Result")
		Label       string
	)

	// get orgin label text
	Label = ResultLabel.Get("innerHTML").String()
	if value.String() == "." {
		point = true
	} else {
		// check Label value is "0"
		if Label == "0" || setanother {
			// empty the value
			Label = ""
			setanother = false
		}
		if point {
			if Label == "" {
				Label = "0"
			}
			Label = Label + "." + strconv.Itoa(value.Int())
			point = false
		} else {
			// Add value to Label
			temp := strconv.Itoa(value.Int())
			Label += temp
		}
		// Set ResultLabel's innerHTML value to Label
		ResultLabel.Set("innerHTML", Label)
	}
	return nil
}

func Arithmetic(_ js.Value, values []js.Value) any {
	var result string
	arithmetic = values[0].String()
	second = StringToFloat.GetFloat64(js.Global().Get("document").Call(
		"getElementById", "Result").Get("innerHTML").String())
	if !setsecond {
		first = second
		setsecond = true
	} else {
		result = calculate(first, second)
		if result == "DZ" {
			result = "Can't Divide by 0"
			js.Global().Get("document").Call("getElementById", "Result").Set("innerHTML", result)
		} else {
			first = StringToFloat.GetFloat64(result)
		}
	}
	setanother = true
	return nil
}

func calculate(first, second float64) string {
	switch arithmetic {
	case "+":
		return strconv.FormatFloat(first+second, 'f', -1, 64)
	case "-":
		return strconv.FormatFloat(first-second, 'f', -1, 64)
	case "*":
		if first == 0 {
			return strconv.FormatFloat(1*second, 'f', -1, 64)
		} else {
			return strconv.FormatFloat(first*second, 'f', -1, 64)
		}
	case "/":
		if second == 0 {
			return "DZ"
		} else {
			return strconv.FormatFloat(first/second, 'f', -1, 64)
		}
	default:
		return "error"
	}
}

func Equal(_ js.Value, values []js.Value) any {
	second = StringToFloat.GetFloat64(js.Global().Get("document").Call("getElementById", "Result").Get("innerHTML").String())
	var result string = ""
	result = calculate(first, second)

	if result == "DZ" {
		result = "Can't Divide by 0"
	}
	setsecond = false
	js.Global().Get("document").Call("getElementById", "Result").Set("innerHTML", result)

	return nil
}

func Clear(_ js.Value, values []js.Value) any {
	first = 0
	second = 0
	arithmetic = ""
	setanother = false
	point = false
	setsecond = false
	js.Global().Get("document").Call("getElementById", "Result").Set("innerHTML", 0)
	return nil
}

func PostiveOrNegativeAndPercent(_ js.Value, values []js.Value) any {
	var (
		value  string  = values[0].String()
		Result float64 = StringToFloat.GetFloat64(js.Global().Get("document").Call("getElementById", "Result").Get("innerHTML").String())
	)
	if value == "+/-" {
		Result *= -1
	} else if value == "%" {
		Result /= 100
	}
	js.Global().Get("document").Call("getElementById", "Result").Set("innerHTML", strconv.FormatFloat(Result, 'f', -1, 64))
	return nil
}

func main() {
	// if html want call golang function need set function name to FuncOf
	js.Global().Set("Print", js.FuncOf(Print))
	js.Global().Set("Arithmetic", js.FuncOf(Arithmetic))
	js.Global().Set("Equal", js.FuncOf(Equal))
	js.Global().Set("Clear", js.FuncOf(Clear))
	js.Global().Set("PostiveOrNegativeAndPercent", js.FuncOf(PostiveOrNegativeAndPercent))
	<-make(chan any)
}
