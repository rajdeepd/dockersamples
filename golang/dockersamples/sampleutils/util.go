package sampleutils

import (
    "encoding/json"
    "fmt"
    "bytes"
    "io/ioutil"
    "net"
    "net/http"
    "net/http/httputil"
    "path/filepath"
    "time"
)
type ResponseJSON []struct {
       Names []string
       Id string
       Created float64
}

type ResponseCreateContainer struct {
       Id string
       Warnings string
}

type ContainerInfo struct {
	Names []string
	Id string
}

func SockRequest(method, endpoint string, data interface{}) ([]byte, error) {
	sock := filepath.Join("/", "var", "run", "docker.sock")
	c, err := net.DialTimeout("unix", sock, time.Duration(10*time.Second))
	if err != nil {
		return nil, fmt.Errorf("could not dial docker sock at %s: %v", sock, err)
	}

	client := httputil.NewClientConn(c, nil)
	defer client.Close()

	jsonData := bytes.NewBuffer(nil)
	if err := json.NewEncoder(jsonData).Encode(data); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, endpoint, jsonData)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, fmt.Errorf("could not create new request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not perform request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return body, fmt.Errorf("received status != 200 OK: %s", resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}

func toJson(object interface {}) string {
    response, err := json.MarshalIndent(object, "", "  ")
    if err != nil {
        panic(err)
    }

    return string(response)
}

func PrettyPrint(object interface {})() {
    fmt.Println(toJson(object))
}
