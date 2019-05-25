// Package endpointdiscovery contains utility functions for working with ed.
package endpointdiscovery

import "os"

var endpointDiscoveryURL string
var ip string

func init() {
	endpointDiscoveryURL = os.Getenv("ENDPOINT_DISCOVERY_SERVICE_HOST")
	ip = os.Getenv("MY_POD_IP")
}

// Path ...
type API struct {
	API      string      `json:"api"`
	Instance string      `json:"instance"`
	Version  string      `json:"version"`
	Endpoint []*Endpoint `json:"endpoints"`
}
type Endpoint struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

type Service struct {
	Hostname string `json:"hostname"`
	Version  string `json:"version"`
	IP       string `json:"ip"`
	Port     int    `json:"port"`
}

// Manifest ....
type Manifest struct {
	APIs    []*API  `json:"endpoints"`
	Service Service `json:"service"`
}

// Endpoint is the structure for the output
type EndpointInstance struct {
	Instance string `json:"instance"`
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
	Port     int    `json:"port"`
	Method   string `json:"method"`
	Path     string `json:"path"`
}

// Endpoint is the structure for the output
type EndpointInstances struct {
	Endpoints []*EndpointInstance `json:"endpoints"`
}
