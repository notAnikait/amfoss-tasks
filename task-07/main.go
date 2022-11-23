package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

var Count = 0

func main() {

	var result = js.Global().Get("document").Call("getElementById", "int")
	var increase = js.Global().Get("document").Call("getElementById", "inc")
	var decrease = js.Global().Get("document").Call("getElementById", "dec")
	var reset = js.Global().Get("document").Call("getElementById", "res")

	increase.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		increase.Set("disabled", true)
		decrease.Set("disabled", true)
		reset.Set("disabled", true)

		Count++
		fmt.Printf("change to %v\n", Count)
		result.Set("innerHTML", strconv.Itoa(Count))

		increase.Set("disabled", false)
		decrease.Set("disabled", false)
		reset.Set("disabled", false)

		return Count
	}))

	decrease.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		increase.Set("disabled", true)
		decrease.Set("disabled", true)
		reset.Set("disabled", true)

		Count--
		fmt.Printf("change to %v\n", Count)
		result.Set("innerHTML", strconv.Itoa(Count))

		increase.Set("disabled", false)
		decrease.Set("disabled", false)
		reset.Set("disabled", false)

		return Count
	}))

	reset.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		increase.Set("disabled", true)
		decrease.Set("disabled", true)
		reset.Set("disabled", true)

		Count = 0
		fmt.Printf("change to %v\n", Count)
		result.Set("innerHTML", strconv.Itoa(Count))

		increase.Set("disabled", false)
		decrease.Set("disabled", false)
		reset.Set("disabled", false)

		return Count
	}))

	<-make(chan bool)
}
