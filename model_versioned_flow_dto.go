/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.12.0-SNAPSHOT
 * Contact: dev@nifi.apache.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nifi

// VersionedFlowDto struct for VersionedFlowDto
type VersionedFlowDto struct {
	// The ID of the registry that the flow is tracked to
	RegistryId string `json:"registryId,omitempty"`
	// The ID of the bucket where the flow is stored
	BucketId string `json:"bucketId,omitempty"`
	// The ID of the flow
	FlowId string `json:"flowId,omitempty"`
	// The name of the flow
	FlowName string `json:"flowName,omitempty"`
	// A description of the flow
	Description string `json:"description,omitempty"`
	// Comments for the changeset
	Comments string `json:"comments,omitempty"`
	// The action being performed
	Action string `json:"action,omitempty"`
}
