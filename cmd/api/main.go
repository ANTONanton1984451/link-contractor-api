package main

import (
	"fmt"
	"link-contractor-api/internal/app"
)

func main() {
	if err := app.Start(); err != nil {
		fmt.Println(err.Error())
	}
}
