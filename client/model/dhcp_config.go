/*
 * VMware Workstation REST API
 *
 * vmrest 1.3.0 build-20800274
 *
 * API version: 1.3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// DHCP configuration
type DhcpConfig struct {
	Enabled bool   `json:"enabled"`
	Setting string `json:"setting"`
}
