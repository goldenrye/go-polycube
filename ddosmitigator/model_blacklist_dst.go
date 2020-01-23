/*
 * ddosmitigator API
 *
 * ddosmitigator API generated from ddosmitigator.yang
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type BlacklistDst struct {
	// Destination IP Address
	Ip string `json:"ip,omitempty"`
	// Dropped Packets
	DropPkts int32 `json:"drop-pkts,omitempty"`
}
