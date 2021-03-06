/*
 * Manages devices through Devicelink. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

type DeviceInfoModelStream struct {
	Key int32 `json:"key,omitempty"`
	ReadOnly bool `json:"read-only,omitempty"`
	Time int32 `json:"time,omitempty"`
	LastValue map[string]interface{} `json:"last-value,omitempty"`
}
