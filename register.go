package endpointdiscovery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Register the endpoints of a the APIs a service implements,
func Register(apis []*API, svc Service) {

	var buf = new(bytes.Buffer)

	instance, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	url := buildRegisterURL(instance)

	payload := &Manifest{apis, svc}
	json.NewEncoder(buf).Encode(&payload)

	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		panic(err)
	}

	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

// helpers

func buildRegisterURL(service Service) string {
	var buffer bytes.Buffer
	buffer.WriteString("http://")
	buffer.WriteString(endpointDiscoveryURL)
	buffer.WriteString("/")
	buffer.WriteString(service.Hostname)
	return buffer.String()
}

//
