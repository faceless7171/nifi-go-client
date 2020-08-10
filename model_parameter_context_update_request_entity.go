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

// ParameterContextUpdateRequestEntity struct for ParameterContextUpdateRequestEntity
type ParameterContextUpdateRequestEntity struct {
	ParameterContextRevision RevisionDto                      `json:"parameterContextRevision,omitempty"`
	Request                  ParameterContextUpdateRequestDto `json:"request,omitempty"`
}
