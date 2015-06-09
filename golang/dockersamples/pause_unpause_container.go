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
	if len(os.Args) > 2 {
		name := os.Args[1]
		action := os.Args[2]
		if action == "pause" {
			pauseContainer(name)
		} else if action == "unpause" {
			unpauseContainer(name)
		} else {
			fmt.Printf("Please specify a valid action : 'pause' or 'unpause'")
		}
	} else {
		fmt.Printf("Please specify container name and action on the command line\n")
		fmt.Printf("go run dockersamples/pause_unpause_container.go [container_name] [pause|unpause]\n")
	}
}

func pauseContainer(name string) {
	containerId := sampleutils.GetContainerId(name)
	if containerId == "" {
		fmt.Printf("Invalid Container name %v\n", name)
	} else {
		_, body, err := sampleutils.SockRequest("POST", "/containers/"+containerId+"/pause", nil)
		if err != nil {
			fmt.Printf("\nerror %v\n", err)
		} else {
			fmt.Printf("body: %#v\n", body)
		}
	}
}

func unpauseContainer(name string) {
	containerId := sampleutils.GetContainerId(name)
	if containerId == "" {
		fmt.Printf("Invalid Container name %v\n", name)
	} else {
		_, body, err := sampleutils.SockRequest("POST", "/containers/"+containerId+"/unpause", nil)
		if err != nil {
			fmt.Printf("\nerror %v\n", err)
		} else {
			fmt.Printf("body: %#v\n", body)
		}
	}
}
