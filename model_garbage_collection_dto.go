/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.11.4
 * Contact: dev@nifi.apache.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nifi

// GarbageCollectionDto struct for GarbageCollectionDto
type GarbageCollectionDto struct {
	// The name of the garbage collector.
	Name string `json:"name,omitempty"`
	// The number of times garbage collection has run.
	CollectionCount int64 `json:"collectionCount,omitempty"`
	// The total amount of time spent garbage collecting.
	CollectionTime string `json:"collectionTime,omitempty"`
	// The total number of milliseconds spent garbage collecting.
	CollectionMillis int64 `json:"collectionMillis,omitempty"`
}