package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"syscall/js"
	"time"
)

const (
	backPortOffset = 8
)

func backendUrl(inUrl js.Value) string {
	urlParsed, _ := url.Parse(inUrl.String())
	// fmt.Println(urlParsed.Scheme, urlParsed.Hostname(), urlParsed.Port())
	var backPort string = ""
	if urlParsed.Port() != "" {
		if portNum, err := strconv.Atoi(urlParsed.Port()); err == nil {
			backPort = fmt.Sprintf(":%v", portNum+backPortOffset)
		}
	}
	return strings.Join(
		[]string{
			urlParsed.Scheme,
			"://",
			urlParsed.Hostname(),
			backPort,
		},
		"",
	)
}

func listAllTeams(this js.Value, args []js.Value) interface{} {
	cli := http.Client{Timeout: time.Duration(1) * time.Second}
	fmt.Println(backendUrl(args[0]))
	response, err := cli.Get(
		fmt.Sprintf(
			"%s/api/teams/",
			backendUrl(args[0]),
		),
	)
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return body
}

func listTeamMembers(this js.Value, args []js.Value) interface{} {
	cli := http.Client{Timeout: time.Duration(1) * time.Second}
	backendUrl := args[0]
	response, err := cli.Get(
		fmt.Sprintf(
			"%s/api/members/",
			backendUrl,
		),
	)
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return body
}

func listTeamOrders(this js.Value, args []js.Value) interface{} {
	cli := http.Client{Timeout: time.Duration(1) * time.Second}
	backendUrl := args[0]
	response, err := cli.Get(
		fmt.Sprintf(
			"%s/api/orders/",
			backendUrl,
		),
	)
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return body
}

func main() {
	done := make(chan struct{}, 0)
	js.Global().Set("listAllTeams", js.FuncOf(listAllTeams))
	js.Global().Set("listTeamMembers", js.FuncOf(listTeamMembers))
	js.Global().Set("listTeamOrders", js.FuncOf(listTeamOrders))
	<-done
}
