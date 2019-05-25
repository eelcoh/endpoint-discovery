package endpointdiscovery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Discover the instances of an API's endpoint.
// APIs have instances (for tenants,shards, whatever)
// and versions
// returns a list of physicalendpoints
func Discover(api string, instance string, version string) EndpointInstances {

	var buf = new(bytes.Buffer)

	url := buildDiscoverURL(api, instance, version)

	req, err := http.NewRequest("GET", url, buf)
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

	// Fill the record with the data from the JSON
	var endpoints EndpointInstances

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&endpoints); err != nil {
		log.Println(err)
	}

	return endpoints
}

// helpers

func buildDiscoverURL(api string, instance string, version string) string {
	var buffer bytes.Buffer
	buffer.WriteString("http://")
	buffer.WriteString(endpointDiscoveryURL)
	buffer.WriteString("/")
	buffer.WriteString(api)
	buffer.WriteString("/")
	buffer.WriteString(instance)
	buffer.WriteString("/")
	buffer.WriteString(version)
	return buffer.String()
}

//
