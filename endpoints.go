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
type Endpoint struct {
	API    string `json:"api"`
	Method string `json:"method"`
	Path   string `json:"path"`
}

// Manifest ....
type Manifest struct {
	Endpoints []*Endpoint `json:"endpoints"`
	Hostname  string      `json:"hostname"`
	Version   string      `json:"version"`
	IP        string      `json:"ip"`
}

// Endpoint is the structure for the output
type EndpointInstance struct {
	Instance string `json:"instance"`
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
	API      string `json:"api"`
	Method   string `json:"method"`
	Path     string `json:"path"`
}
