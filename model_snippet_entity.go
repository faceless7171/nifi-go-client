/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.13.2
 * Contact: dev@nifi.apache.org
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nifi

import (
	"encoding/json"
)

// SnippetEntity struct for SnippetEntity
type SnippetEntity struct {
	Snippet *SnippetDTO `json:"snippet,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged *bool `json:"disconnectedNodeAcknowledged,omitempty"`
}

// NewSnippetEntity instantiates a new SnippetEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnippetEntity() *SnippetEntity {
	this := SnippetEntity{}
	return &this
}

// NewSnippetEntityWithDefaults instantiates a new SnippetEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnippetEntityWithDefaults() *SnippetEntity {
	this := SnippetEntity{}
	return &this
}

// GetSnippet returns the Snippet field value if set, zero value otherwise.
func (o *SnippetEntity) GetSnippet() SnippetDTO {
	if o == nil || o.Snippet == nil {
		var ret SnippetDTO
		return ret
	}
	return *o.Snippet
}

// GetSnippetOk returns a tuple with the Snippet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnippetEntity) GetSnippetOk() (*SnippetDTO, bool) {
	if o == nil || o.Snippet == nil {
		return nil, false
	}
	return o.Snippet, true
}

// HasSnippet returns a boolean if a field has been set.
func (o *SnippetEntity) HasSnippet() bool {
	if o != nil && o.Snippet != nil {
		return true
	}

	return false
}

// SetSnippet gets a reference to the given SnippetDTO and assigns it to the Snippet field.
func (o *SnippetEntity) SetSnippet(v SnippetDTO) {
	o.Snippet = &v
}

// GetDisconnectedNodeAcknowledged returns the DisconnectedNodeAcknowledged field value if set, zero value otherwise.
func (o *SnippetEntity) GetDisconnectedNodeAcknowledged() bool {
	if o == nil || o.DisconnectedNodeAcknowledged == nil {
		var ret bool
		return ret
	}
	return *o.DisconnectedNodeAcknowledged
}

// GetDisconnectedNodeAcknowledgedOk returns a tuple with the DisconnectedNodeAcknowledged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnippetEntity) GetDisconnectedNodeAcknowledgedOk() (*bool, bool) {
	if o == nil || o.DisconnectedNodeAcknowledged == nil {
		return nil, false
	}
	return o.DisconnectedNodeAcknowledged, true
}

// HasDisconnectedNodeAcknowledged returns a boolean if a field has been set.
func (o *SnippetEntity) HasDisconnectedNodeAcknowledged() bool {
	if o != nil && o.DisconnectedNodeAcknowledged != nil {
		return true
	}

	return false
}

// SetDisconnectedNodeAcknowledged gets a reference to the given bool and assigns it to the DisconnectedNodeAcknowledged field.
func (o *SnippetEntity) SetDisconnectedNodeAcknowledged(v bool) {
	o.DisconnectedNodeAcknowledged = &v
}

func (o SnippetEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Snippet != nil {
		toSerialize["snippet"] = o.Snippet
	}
	if o.DisconnectedNodeAcknowledged != nil {
		toSerialize["disconnectedNodeAcknowledged"] = o.DisconnectedNodeAcknowledged
	}
	return json.Marshal(toSerialize)
}

type NullableSnippetEntity struct {
	value *SnippetEntity
	isSet bool
}

func (v NullableSnippetEntity) Get() *SnippetEntity {
	return v.value
}

func (v *NullableSnippetEntity) Set(val *SnippetEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableSnippetEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableSnippetEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnippetEntity(val *SnippetEntity) *NullableSnippetEntity {
	return &NullableSnippetEntity{value: val, isSet: true}
}

func (v NullableSnippetEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnippetEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
