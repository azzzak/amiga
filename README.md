# amiga [![GoDoc](https://godoc.org/github.com/azzzak/amiga?status.svg)](https://godoc.org/github.com/azzzak/amiga)

Golang connector for Asterisk Manager Interface (AMI).

## Install

`go get -u github.com/azzzak/amiga`

## Usage

```go
package main

import (
	"fmt"

	"github.com/azzzak/amiga"
	"github.com/azzzak/amiga/action"
	"github.com/azzzak/amiga/event"
)

func main() {
	// catch connection event
	onConnect := amiga.OnConnect(func(msg string) {
		fmt.Printf("connect %s\n", msg)
	})

	// catch errors
	errHandler := amiga.OnError(func(resp amiga.ErrResp) {
		fmt.Printf("error %v\n", resp)
	})

	// catch all events
	onEvent := amiga.OnEvent(func(m map[string]string) {
		fmt.Printf("event %v\n", m)
	})

	// new amiga instance, you may omit any handler, the order doesn't matter
	ami := amiga.New(errHandler, onConnect, onEvent)

	// register handler for specific event
	ami.RegisterHandler(event.FullyBooted, func(m map[string]string) {
		fmt.Printf("booted %v\n", m)

		// amiga.Populate helps you get fields of interest
		var status, uptime string
		err := amiga.Populate(m, map[string]*string{
			"Status": &status,
			"Uptime": &uptime,
		})
		if err != nil {
			// error shows some fields are missing
		}
		fmt.Println(status, uptime)
	})

	// connect to asterisk
	err := ami.Connect(amiga.Options{
		Host:     "localhost", // default: "localhost"
		Port:     5038,        // default: 5038
		Username: "guest",
		Secret:   "guest",
	})
	if err != nil {
		fmt.Println("error:", err)
	}

	// send action
	res, err := ami.Action(action.QueueStatus, map[string]string{
		"ActionID": "7a67-b519-09fa",
	})
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(res)

	ch := make(chan int)
	<-ch

	// send Logoff action, close connection
	ami.Close()
}
```