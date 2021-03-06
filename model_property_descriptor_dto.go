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

// PropertyDescriptorDTO struct for PropertyDescriptorDTO
type PropertyDescriptorDTO struct {
	// The name for the property.
	Name *string `json:"name,omitempty"`
	// The human readable name for the property.
	DisplayName *string `json:"displayName,omitempty"`
	// The description for the property. Used to relay additional details to a user or provide a mechanism of documenting intent.
	Description *string `json:"description,omitempty"`
	// The default value for the property.
	DefaultValue *string `json:"defaultValue,omitempty"`
	// Allowable values for the property. If empty then the allowed values are not constrained.
	AllowableValues *[]AllowableValueEntity `json:"allowableValues,omitempty"`
	// Whether the property is required.
	Required *bool `json:"required,omitempty"`
	// Whether the property is sensitive and protected whenever stored or represented.
	Sensitive *bool `json:"sensitive,omitempty"`
	// Whether the property is dynamic (user-defined).
	Dynamic *bool `json:"dynamic,omitempty"`
	// Whether the property supports expression language.
	SupportsEl *bool `json:"supportsEl,omitempty"`
	// Scope of the Expression Language evaluation for the property.
	ExpressionLanguageScope *string `json:"expressionLanguageScope,omitempty"`
	// If the property identifies a controller service this returns the fully qualified type.
	IdentifiesControllerService       *string    `json:"identifiesControllerService,omitempty"`
	IdentifiesControllerServiceBundle *BundleDTO `json:"identifiesControllerServiceBundle,omitempty"`
	// A list of dependencies that must be met in order for this Property to be relevant. If any of these dependencies is not met, the property described by this Property Descriptor is not relevant.
	Dependencies *[]PropertyDependencyDTO `json:"dependencies,omitempty"`
}

// NewPropertyDescriptorDTO instantiates a new PropertyDescriptorDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyDescriptorDTO() *PropertyDescriptorDTO {
	this := PropertyDescriptorDTO{}
	return &this
}

// NewPropertyDescriptorDTOWithDefaults instantiates a new PropertyDescriptorDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyDescriptorDTOWithDefaults() *PropertyDescriptorDTO {
	this := PropertyDescriptorDTO{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PropertyDescriptorDTO) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDescriptorDTO) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PropertyDescriptorDTO) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PropertyDescriptorDTO) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *PropertyDescriptorDTO) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDescriptorDTO) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PropertyDescriptorDTO) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *PropertyDescriptorDTO) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PropertyDescriptorDTO) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDescriptorDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PropertyDescriptorDTO) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PropertyDescriptorDTO) SetDescription(v string) {
	o.Description = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *PropertyDescriptorDTO) GetDefaultValue() string {
	if o == nil || o.DefaultValue == nil {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDescriptorDTO) GetDefaultValueOk() (*string, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *PropertyDescriptorDTO) HasDefaultValue() bool {
	if o != nil && o.DefaultValue != nil {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *PropertyDescriptorDTO) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetAllowableValues returns the AllowableValues field value if set, zero value otherwise.
func (o *PropertyDescriptorDTO) GetAllowableValues() []AllowableValueEntity {
	if o == nil || o.AllowableValues == nil {
		var ret []AllowableValueEntity
		return ret
	}
	return *o.AllowableValues
}

// GetAllowableValuesOk returns a tuple with the AllowableValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDescriptorDTO) GetAllowableValuesOk() (*[]AllowableValueEntity, bool) {
	if o == nil || o.AllowableValues == nil {
		return nil, false
	}
	return o.AllowableValues, true
}

// HasAllowableValues returns a boolean if a field has been set.
func (o *PropertyDescriptorDTO) HasAllowableValues() bool {
	if o != nil && o.AllowableValues != nil {
		return true
	}

	return false
}

// SetAllowableValues gets a reference to the given []AllowableValueEntity and assigns it to the AllowableValues field.
func (o *PropertyDescriptorDTO) SetAllowableValues(v []AllowableValueEntity) {
	o.AllowableValues = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *PropertyDescriptorDTO) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDescriptorDTO) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *PropertyDescriptorDTO) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *PropertyDescriptorDTO) SetRequired(v bool) {
	o.Required = &v
}

// GetSensitive returns the Sensitive field value if set, zero value otherwise.
func (o *PropertyDescriptorDTO) GetSensitive() bool {
	if o == nil || o.Sensitive == nil {
		var ret bool
		return ret
	}
	return *o.Sensitive
}

// GetSensitiveOk returns a tuple with the Sensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDescriptorDTO) GetSensitiveOk() (*bool, bool) {
	if o == nil || o.Sensitive == nil {
		return nil, false
	}
	return o.Sensitive, true
}

// HasSensitive returns a boolean if a field has been set.
func (o *PropertyDescriptorDTO) HasSensitive() bool {
	if o != nil && o.Sensitive != nil {
		return true
	}

	return false
}

// SetSensitive gets a reference to the given bool and assigns it to the Sensitive field.
func (o *PropertyDescriptorDTO) SetSensitive(v bool) {
	o.Sensitive = &v
}

// GetDynamic returns the Dynamic field value if set, zero value otherwise.
func (o *PropertyDescriptorDTO) GetDynamic() bool {
	if o == nil || o.Dynamic == nil {
		var ret bool
		return ret
	}
	return *o.Dynamic
}

// GetDynamicOk returns a tuple with the Dynamic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDescriptorDTO) GetDynamicOk() (*bool, bool) {
	if o == nil || o.Dynamic == nil {
		return nil, false
	}
	return o.Dynamic, true
}

// HasDynamic returns a boolean if a field has been set.
func (o *PropertyDescriptorDTO) HasDynamic() bool {
	if o != nil && o.Dynamic != nil {
		return true
	}

	return false
}

// SetDynamic gets a reference to the given bool and assigns it to the Dynamic field.
func (o *PropertyDescriptorDTO) SetDynamic(v bool) {
	o.Dynamic = &v
}

// GetSupportsEl returns the SupportsEl field value if set, zero value otherwise.
func (o *PropertyDescriptorDTO) GetSupportsEl() bool {
	if o == nil || o.SupportsEl == nil {
		var ret bool
		return ret
	}
	return *o.SupportsEl
}

// GetSupportsElOk returns a tuple with the SupportsEl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDescriptorDTO) GetSupportsElOk() (*bool, bool) {
	if o == nil || o.SupportsEl == nil {
		return nil, false
	}
	return o.SupportsEl, true
}

// HasSupportsEl returns a boolean if a field has been set.
func (o *PropertyDescriptorDTO) HasSupportsEl() bool {
	if o != nil && o.SupportsEl != nil {
		return true
	}

	return false
}

// SetSupportsEl gets a reference to the given bool and assigns it to the SupportsEl field.
func (o *PropertyDescriptorDTO) SetSupportsEl(v bool) {
	o.SupportsEl = &v
}

// GetExpressionLanguageScope returns the ExpressionLanguageScope field value if set, zero value otherwise.
func (o *PropertyDescriptorDTO) GetExpressionLanguageScope() string {
	if o == nil || o.ExpressionLanguageScope == nil {
		var ret string
		return ret
	}
	return *o.ExpressionLanguageScope
}

// GetExpressionLanguageScopeOk returns a tuple with the ExpressionLanguageScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDescriptorDTO) GetExpressionLanguageScopeOk() (*string, bool) {
	if o == nil || o.ExpressionLanguageScope == nil {
		return nil, false
	}
	return o.ExpressionLanguageScope, true
}

// HasExpressionLanguageScope returns a boolean if a field has been set.
func (o *PropertyDescriptorDTO) HasExpressionLanguageScope() bool {
	if o != nil && o.ExpressionLanguageScope != nil {
		return true
	}

	return false
}

// SetExpressionLanguageScope gets a reference to the given string and assigns it to the ExpressionLanguageScope field.
func (o *PropertyDescriptorDTO) SetExpressionLanguageScope(v string) {
	o.ExpressionLanguageScope = &v
}

// GetIdentifiesControllerService returns the IdentifiesControllerService field value if set, zero value otherwise.
func (o *PropertyDescriptorDTO) GetIdentifiesControllerService() string {
	if o == nil || o.IdentifiesControllerService == nil {
		var ret string
		return ret
	}
	return *o.IdentifiesControllerService
}

// GetIdentifiesControllerServiceOk returns a tuple with the IdentifiesControllerService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDescriptorDTO) GetIdentifiesControllerServiceOk() (*string, bool) {
	if o == nil || o.IdentifiesControllerService == nil {
		return nil, false
	}
	return o.IdentifiesControllerService, true
}

// HasIdentifiesControllerService returns a boolean if a field has been set.
func (o *PropertyDescriptorDTO) HasIdentifiesControllerService() bool {
	if o != nil && o.IdentifiesControllerService != nil {
		return true
	}

	return false
}

// SetIdentifiesControllerService gets a reference to the given string and assigns it to the IdentifiesControllerService field.
func (o *PropertyDescriptorDTO) SetIdentifiesControllerService(v string) {
	o.IdentifiesControllerService = &v
}

// GetIdentifiesControllerServiceBundle returns the IdentifiesControllerServiceBundle field value if set, zero value otherwise.
func (o *PropertyDescriptorDTO) GetIdentifiesControllerServiceBundle() BundleDTO {
	if o == nil || o.IdentifiesControllerServiceBundle == nil {
		var ret BundleDTO
		return ret
	}
	return *o.IdentifiesControllerServiceBundle
}

// GetIdentifiesControllerServiceBundleOk returns a tuple with the IdentifiesControllerServiceBundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDescriptorDTO) GetIdentifiesControllerServiceBundleOk() (*BundleDTO, bool) {
	if o == nil || o.IdentifiesControllerServiceBundle == nil {
		return nil, false
	}
	return o.IdentifiesControllerServiceBundle, true
}

// HasIdentifiesControllerServiceBundle returns a boolean if a field has been set.
func (o *PropertyDescriptorDTO) HasIdentifiesControllerServiceBundle() bool {
	if o != nil && o.IdentifiesControllerServiceBundle != nil {
		return true
	}

	return false
}

// SetIdentifiesControllerServiceBundle gets a reference to the given BundleDTO and assigns it to the IdentifiesControllerServiceBundle field.
func (o *PropertyDescriptorDTO) SetIdentifiesControllerServiceBundle(v BundleDTO) {
	o.IdentifiesControllerServiceBundle = &v
}

// GetDependencies returns the Dependencies field value if set, zero value otherwise.
func (o *PropertyDescriptorDTO) GetDependencies() []PropertyDependencyDTO {
	if o == nil || o.Dependencies == nil {
		var ret []PropertyDependencyDTO
		return ret
	}
	return *o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDescriptorDTO) GetDependenciesOk() (*[]PropertyDependencyDTO, bool) {
	if o == nil || o.Dependencies == nil {
		return nil, false
	}
	return o.Dependencies, true
}

// HasDependencies returns a boolean if a field has been set.
func (o *PropertyDescriptorDTO) HasDependencies() bool {
	if o != nil && o.Dependencies != nil {
		return true
	}

	return false
}

// SetDependencies gets a reference to the given []PropertyDependencyDTO and assigns it to the Dependencies field.
func (o *PropertyDescriptorDTO) SetDependencies(v []PropertyDependencyDTO) {
	o.Dependencies = &v
}

func (o PropertyDescriptorDTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if o.AllowableValues != nil {
		toSerialize["allowableValues"] = o.AllowableValues
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.Sensitive != nil {
		toSerialize["sensitive"] = o.Sensitive
	}
	if o.Dynamic != nil {
		toSerialize["dynamic"] = o.Dynamic
	}
	if o.SupportsEl != nil {
		toSerialize["supportsEl"] = o.SupportsEl
	}
	if o.ExpressionLanguageScope != nil {
		toSerialize["expressionLanguageScope"] = o.ExpressionLanguageScope
	}
	if o.IdentifiesControllerService != nil {
		toSerialize["identifiesControllerService"] = o.IdentifiesControllerService
	}
	if o.IdentifiesControllerServiceBundle != nil {
		toSerialize["identifiesControllerServiceBundle"] = o.IdentifiesControllerServiceBundle
	}
	if o.Dependencies != nil {
		toSerialize["dependencies"] = o.Dependencies
	}
	return json.Marshal(toSerialize)
}

type NullablePropertyDescriptorDTO struct {
	value *PropertyDescriptorDTO
	isSet bool
}

func (v NullablePropertyDescriptorDTO) Get() *PropertyDescriptorDTO {
	return v.value
}

func (v *NullablePropertyDescriptorDTO) Set(val *PropertyDescriptorDTO) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyDescriptorDTO) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyDescriptorDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyDescriptorDTO(val *PropertyDescriptorDTO) *NullablePropertyDescriptorDTO {
	return &NullablePropertyDescriptorDTO{value: val, isSet: true}
}

func (v NullablePropertyDescriptorDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyDescriptorDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
