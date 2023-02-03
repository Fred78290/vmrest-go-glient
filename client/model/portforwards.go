/*
 * VMware Workstation REST API
 *
 * vmrest 1.3.0 build-20800274
 *
 * API version: 1.3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// The list of port forwarding
type Portforwards struct {
	Num int `json:"num"`
	// The list of port forwardings
	PortForwardings []Portforward `json:"port_forwardings"`
}
