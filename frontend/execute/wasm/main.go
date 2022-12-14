package main

import (
	"fmt"
	"syscall/js"
)

func listTeam(this js.Value, args []js.Value) interface{} {
	return fmt.Sprintf("Listing team %v", args[0])
}

func main() {
	done := make(chan struct{}, 0)
	js.Global().Set("listTeam", js.FuncOf(listTeam))
	<-done
}
