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

// FunnelDTO struct for FunnelDTO
type FunnelDTO struct {
	// The id of the component.
	Id *string `json:"id,omitempty"`
	// The ID of the corresponding component that is under version control
	VersionedComponentId *string `json:"versionedComponentId,omitempty"`
	// The id of parent process group of this component if applicable.
	ParentGroupId *string      `json:"parentGroupId,omitempty"`
	Position      *PositionDTO `json:"position,omitempty"`
}

// NewFunnelDTO instantiates a new FunnelDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFunnelDTO() *FunnelDTO {
	this := FunnelDTO{}
	return &this
}

// NewFunnelDTOWithDefaults instantiates a new FunnelDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFunnelDTOWithDefaults() *FunnelDTO {
	this := FunnelDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FunnelDTO) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelDTO) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FunnelDTO) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FunnelDTO) SetId(v string) {
	o.Id = &v
}

// GetVersionedComponentId returns the VersionedComponentId field value if set, zero value otherwise.
func (o *FunnelDTO) GetVersionedComponentId() string {
	if o == nil || o.VersionedComponentId == nil {
		var ret string
		return ret
	}
	return *o.VersionedComponentId
}

// GetVersionedComponentIdOk returns a tuple with the VersionedComponentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelDTO) GetVersionedComponentIdOk() (*string, bool) {
	if o == nil || o.VersionedComponentId == nil {
		return nil, false
	}
	return o.VersionedComponentId, true
}

// HasVersionedComponentId returns a boolean if a field has been set.
func (o *FunnelDTO) HasVersionedComponentId() bool {
	if o != nil && o.VersionedComponentId != nil {
		return true
	}

	return false
}

// SetVersionedComponentId gets a reference to the given string and assigns it to the VersionedComponentId field.
func (o *FunnelDTO) SetVersionedComponentId(v string) {
	o.VersionedComponentId = &v
}

// GetParentGroupId returns the ParentGroupId field value if set, zero value otherwise.
func (o *FunnelDTO) GetParentGroupId() string {
	if o == nil || o.ParentGroupId == nil {
		var ret string
		return ret
	}
	return *o.ParentGroupId
}

// GetParentGroupIdOk returns a tuple with the ParentGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelDTO) GetParentGroupIdOk() (*string, bool) {
	if o == nil || o.ParentGroupId == nil {
		return nil, false
	}
	return o.ParentGroupId, true
}

// HasParentGroupId returns a boolean if a field has been set.
func (o *FunnelDTO) HasParentGroupId() bool {
	if o != nil && o.ParentGroupId != nil {
		return true
	}

	return false
}

// SetParentGroupId gets a reference to the given string and assigns it to the ParentGroupId field.
func (o *FunnelDTO) SetParentGroupId(v string) {
	o.ParentGroupId = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *FunnelDTO) GetPosition() PositionDTO {
	if o == nil || o.Position == nil {
		var ret PositionDTO
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelDTO) GetPositionOk() (*PositionDTO, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *FunnelDTO) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given PositionDTO and assigns it to the Position field.
func (o *FunnelDTO) SetPosition(v PositionDTO) {
	o.Position = &v
}

func (o FunnelDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.VersionedComponentId != nil {
		toSerialize["versionedComponentId"] = o.VersionedComponentId
	}
	if o.ParentGroupId != nil {
		toSerialize["parentGroupId"] = o.ParentGroupId
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	return json.Marshal(toSerialize)
}

type NullableFunnelDTO struct {
	value *FunnelDTO
	isSet bool
}

func (v NullableFunnelDTO) Get() *FunnelDTO {
	return v.value
}

func (v *NullableFunnelDTO) Set(val *FunnelDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableFunnelDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableFunnelDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFunnelDTO(val *FunnelDTO) *NullableFunnelDTO {
	return &NullableFunnelDTO{value: val, isSet: true}
}

func (v NullableFunnelDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFunnelDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
