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

// UserGroupsEntity struct for UserGroupsEntity
type UserGroupsEntity struct {
	UserGroups *[]UserGroupEntity `json:"userGroups,omitempty"`
}

// NewUserGroupsEntity instantiates a new UserGroupsEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroupsEntity() *UserGroupsEntity {
	this := UserGroupsEntity{}
	return &this
}

// NewUserGroupsEntityWithDefaults instantiates a new UserGroupsEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupsEntityWithDefaults() *UserGroupsEntity {
	this := UserGroupsEntity{}
	return &this
}

// GetUserGroups returns the UserGroups field value if set, zero value otherwise.
func (o *UserGroupsEntity) GetUserGroups() []UserGroupEntity {
	if o == nil || o.UserGroups == nil {
		var ret []UserGroupEntity
		return ret
	}
	return *o.UserGroups
}

// GetUserGroupsOk returns a tuple with the UserGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupsEntity) GetUserGroupsOk() (*[]UserGroupEntity, bool) {
	if o == nil || o.UserGroups == nil {
		return nil, false
	}
	return o.UserGroups, true
}

// HasUserGroups returns a boolean if a field has been set.
func (o *UserGroupsEntity) HasUserGroups() bool {
	if o != nil && o.UserGroups != nil {
		return true
	}

	return false
}

// SetUserGroups gets a reference to the given []UserGroupEntity and assigns it to the UserGroups field.
func (o *UserGroupsEntity) SetUserGroups(v []UserGroupEntity) {
	o.UserGroups = &v
}

func (o UserGroupsEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserGroups != nil {
		toSerialize["userGroups"] = o.UserGroups
	}
	return json.Marshal(toSerialize)
}

type NullableUserGroupsEntity struct {
	value *UserGroupsEntity
	isSet bool
}

func (v NullableUserGroupsEntity) Get() *UserGroupsEntity {
	return v.value
}

func (v *NullableUserGroupsEntity) Set(val *UserGroupsEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroupsEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroupsEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroupsEntity(val *UserGroupsEntity) *NullableUserGroupsEntity {
	return &NullableUserGroupsEntity{value: val, isSet: true}
}

func (v NullableUserGroupsEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroupsEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
