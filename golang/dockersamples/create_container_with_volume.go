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
		arg := os.Args[1]
		name := arg
		CreateContainerWithVolume(name)
	} else {
		fmt.Printf("Please specify container name on the command line\n")
	}
}

func CreateContainerWithVolume(name string) {
	config := map[string]interface{}{
		"Image":     "busybox",
		"OpenStdin": true,
		"Volumes":   map[string]struct{}{"/tmp": {}},
	}
	_, body, err := sampleutils.SockRequest("POST", "/containers/create?name="+name, config)
	if err != nil {
		fmt.Printf("Error %v\n", err)

	} else {
		var resp *sampleutils.ResponseCreateContainer
		if err = json.Unmarshal(body, &resp); err != nil {
			fmt.Printf("unable to unmarshal response body: %v\n", err)
		}
		sampleutils.PrettyPrint(resp)
	}

}
