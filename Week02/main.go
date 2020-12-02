package main

import (
	"Week002/api"
	"fmt"
)

func main() {
	//negative case
	video, err := api.QueryVideoInfoById(0)

	//posttive case
	//video, err := api.QueryVideoInfoById(1)
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		fmt.Println(*video)
	}
}