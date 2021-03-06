/*
 * Manages devices through Devicelink. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

type DevicePropertiesStatus struct {
	Topic DevicePropertiesStatusTopic `json:"topic,omitempty"`
	Predicate DevicePropertiesStatusPredicate `json:"predicate,omitempty"`
	Actions DevicePropertiesStatusActions `json:"actions,omitempty"`
	Status string `json:"status,omitempty"`
}
