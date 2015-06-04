package sampleutils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"net/url"
	"time"
)

var (
	DefaultUnixSocket = "/var/run/docker.sock"
)

type ResponseJSON []struct {
	Names   []string
	Id      string
	Created float64
}

type ResponseCreateContainer struct {
	Id       string
	Warnings string
}

type ContainerInfo struct {
	Names []string
	Id    string
}

/*func SockRequest(method, endpoint string, data interface{}) ([]byte, error) {
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
}*/

func SockRequest(method, endpoint string, data interface{}) (int, []byte, error) {
	jsonData := bytes.NewBuffer(nil)
	if err := json.NewEncoder(jsonData).Encode(data); err != nil {
		return -1, nil, err
	}

	res, body, err := sockRequestRaw(method, endpoint, jsonData, "application/json")
	if err != nil {
		b, _ := ioutil.ReadAll(body)
		return -1, b, err
	}
	var b []byte
	b, err = readBody(body)
	return res.StatusCode, b, err
}

func sockRequestRaw(method, endpoint string, data io.Reader, ct string) (*http.Response, io.ReadCloser, error) {
	c, err := sockConn(time.Duration(10 * time.Second))
	if err != nil {
		return nil, nil, fmt.Errorf("could not dial docker daemon: %v", err)
	}

	client := httputil.NewClientConn(c, nil)

	req, err := http.NewRequest(method, endpoint, data)
	if err != nil {
		client.Close()
		return nil, nil, fmt.Errorf("could not create new request: %v", err)
	}

	if ct != "" {
		req.Header.Set("Content-Type", ct)
	}

	resp, err := client.Do(req)
	if err != nil {
		client.Close()
		return nil, nil, fmt.Errorf("could not perform request: %v", err)
	}
	body := NewReadCloserWrapper(resp.Body, func() error {
		defer client.Close()
		return resp.Body.Close()
	})

	return resp, body, nil
}

func daemonHost() string {
	daemonUrlStr := "unix://" + DefaultUnixSocket
	if daemonHostVar := os.Getenv("DOCKER_HOST"); daemonHostVar != "" {
		daemonUrlStr = daemonHostVar
	}
	return daemonUrlStr
}

func sockConn(timeout time.Duration) (net.Conn, error) {
	daemon := daemonHost()
	daemonUrl, err := url.Parse(daemon)
	if err != nil {
		return nil, fmt.Errorf("could not parse url %q: %v", daemon, err)
	}

	var c net.Conn
	switch daemonUrl.Scheme {
	case "unix":
		return net.DialTimeout(daemonUrl.Scheme, daemonUrl.Path, timeout)
	case "tcp":
		return net.DialTimeout(daemonUrl.Scheme, daemonUrl.Host, timeout)
	default:
		return c, fmt.Errorf("unknown scheme %v (%s)", daemonUrl.Scheme, daemon)
	}
}

type readCloserWrapper struct {
	io.Reader
	closer func() error
}

func (r *readCloserWrapper) Close() error {
	return r.closer()
}

func readBody(b io.ReadCloser) ([]byte, error) {
	defer b.Close()
	return ioutil.ReadAll(b)
}

func NewReadCloserWrapper(r io.Reader, closer func() error) io.ReadCloser {
	return &readCloserWrapper{
		Reader: r,
		closer: closer,
	}
}
func toJson(object interface{}) string {
	response, err := json.MarshalIndent(object, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(response)
}

func PrettyPrint(object interface{}) {
	fmt.Println(toJson(object))
}

func GetContainerInfoFor(name string, inspectJSON []ContainerInfo) (ContainerInfo, error) {
	var noOfContainers = len(inspectJSON)
	for i := 0; i < noOfContainers; i++ {
		nameLocal := inspectJSON[i].Names[0]
		name_trimmed := nameLocal[1:]
		if name_trimmed == name {
			return inspectJSON[i], nil
		}
	}
	var value ContainerInfo
	return value, errors.New("No container information found for " + name)
}

func GetContainerId(name string) (id string) {
	containerId := ""
	status, body, err := SockRequest("GET", "/containers/json?all=1", nil)
	fmt.Printf("Status of the Request: %v\n", status)
	var inspectJSON []ContainerInfo

	if err = json.Unmarshal(body, &inspectJSON); err != nil {
		fmt.Printf("unable to unmarshal response body: %v", err)
	}

	container, err := GetContainerInfoFor(name, inspectJSON)

	if err == nil {
		containerId = container.Id
	}

	return containerId
}
