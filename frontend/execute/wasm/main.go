package main

import (
	"syscall/js"
)

func listTeams(this js.Value, args []js.Value) interface{} {
	return "Listing teams to be implemented"
}

func main() {
	done := make(chan struct{}, 0)
	js.Global().Set("listTeams", js.FuncOf(listTeams))
	<-done
}
