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
		ListProcesses(name)
	} else {
		fmt.Printf("Please specify container name on the command line\n")
	}
}

func ListProcesses(name string) {
	containerId := sampleutils.GetContainerId(name)
	if containerId == "" {
		fmt.Printf("Invalid Container name %v\n", name)
	} else {
		_, body, err := sampleutils.SockRequest("POST", "/containers/"+containerId+"/top", nil)
		if err != nil {
			fmt.Printf("error %v\n", err)
		} else {
			fmt.Printf("body: %#v\n", body)
		}
	}
}
