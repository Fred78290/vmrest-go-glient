/*
 * VMware Workstation REST API
 *
 * vmrest 1.3.0 build-20800274
 *
 * API version: 1.3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// DNS configuration
type DnsConfig struct {
	Hostname   string   `json:"hostname,omitempty"`
	Domainname string   `json:"domainname,omitempty"`
	Server     []string `json:"server,omitempty"`
	Search     []string `json:"search,omitempty"`
}
