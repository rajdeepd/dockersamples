package main

import (
	"fmt"
	"github.com/rajdeepd/dockersamples/golang/dockersamples/sampleutils"
	"os"
)

type ContainerInfo struct {
	Names []string
	Id    string
}

func main() {
	if len(os.Args) > 1 {
		arg := os.Args[1]
		name := arg
		StartContainer(name)
	} else {
		fmt.Printf("Please specify container name on the command line\n")
	}
}

func StartContainer(name string) {
	containerId := sampleutils.GetContainerId(name)
	fmt.Printf("startContainer :Container Id : %v", containerId)
	if containerId == "" {
		fmt.Printf("Invalid Container name %v\n", name)
	} else {
		status, _, err := sampleutils.SockRequest("POST", "/containers/"+containerId+"/start", nil)
		if err != nil {
			fmt.Printf("\nerror %v\n", err)
		} else {
			fmt.Printf("\nStatus of the Start Request: %#v\n", status)
		}
	}
}
