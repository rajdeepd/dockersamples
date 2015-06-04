package main

import (
	"encoding/json"
	"fmt"
	"github.com/rajdeepd/dockersamples/golang/dockersamples/sampleutils"
)

type ContainerInfo struct {
	Names []string
	Id    string
}

func main() {
	ListContainers()
}

func ListContainers() {
	_, body, err := sampleutils.SockRequest("GET", "/containers/json?all=1", nil)
	var respJSON *sampleutils.ResponseJSON
	if err = json.Unmarshal(body, &respJSON); err != nil {
		fmt.Printf("unable to unmarshal response body: %v", err)
	}
	sampleutils.PrettyPrint(respJSON)
}
