package main

import (
	"encoding/json"
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
		name := os.Args[1]
		getContainerState(name)
	} else {
		fmt.Printf("Please specify container name on the command line\n")
	}
}

func getContainerState(name string) (containerState string) {
	containerId := ""
	containerId = sampleutils.GetContainerId(name)
	if containerId != "" {
		_, body, err := sampleutils.SockRequest("GET", "/containers/"+containerId+"/json", nil)
		var inspectJSON struct {
			Id    string
			State string
		}
		if err = json.Unmarshal(body, &inspectJSON); err != nil {
			fmt.Printf("unable to unmarshal response body: %v", err)
		}
		containerState := string(inspectJSON.State)
		return containerState
	} else {
		fmt.Printf("Container doesn't exist")
		return ""
	}
}
