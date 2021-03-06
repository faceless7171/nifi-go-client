# TenantDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the component. | [optional] 
**VersionedComponentId** | Pointer to **string** | The ID of the corresponding component that is under version control | [optional] 
**ParentGroupId** | Pointer to **string** | The id of parent process group of this component if applicable. | [optional] 
**Position** | Pointer to [**PositionDTO**](PositionDTO.md) |  | [optional] 
**Identity** | Pointer to **string** | The identity of the tenant. | [optional] 
**Configurable** | Pointer to **bool** | Whether this tenant is configurable. | [optional] 

## Methods

### NewTenantDTO

`func NewTenantDTO() *TenantDTO`

NewTenantDTO instantiates a new TenantDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantDTOWithDefaults

`func NewTenantDTOWithDefaults() *TenantDTO`

NewTenantDTOWithDefaults instantiates a new TenantDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersionedComponentId

`func (o *TenantDTO) GetVersionedComponentId() string`

GetVersionedComponentId returns the VersionedComponentId field if non-nil, zero value otherwise.

### GetVersionedComponentIdOk

`func (o *TenantDTO) GetVersionedComponentIdOk() (*string, bool)`

GetVersionedComponentIdOk returns a tuple with the VersionedComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedComponentId

`func (o *TenantDTO) SetVersionedComponentId(v string)`

SetVersionedComponentId sets VersionedComponentId field to given value.

### HasVersionedComponentId

`func (o *TenantDTO) HasVersionedComponentId() bool`

HasVersionedComponentId returns a boolean if a field has been set.

### GetParentGroupId

`func (o *TenantDTO) GetParentGroupId() string`

GetParentGroupId returns the ParentGroupId field if non-nil, zero value otherwise.

### GetParentGroupIdOk

`func (o *TenantDTO) GetParentGroupIdOk() (*string, bool)`

GetParentGroupIdOk returns a tuple with the ParentGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroupId

`func (o *TenantDTO) SetParentGroupId(v string)`

SetParentGroupId sets ParentGroupId field to given value.

### HasParentGroupId

`func (o *TenantDTO) HasParentGroupId() bool`

HasParentGroupId returns a boolean if a field has been set.

### GetPosition

`func (o *TenantDTO) GetPosition() PositionDTO`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *TenantDTO) GetPositionOk() (*PositionDTO, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *TenantDTO) SetPosition(v PositionDTO)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *TenantDTO) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetIdentity

`func (o *TenantDTO) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *TenantDTO) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *TenantDTO) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *TenantDTO) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetConfigurable

`func (o *TenantDTO) GetConfigurable() bool`

GetConfigurable returns the Configurable field if non-nil, zero value otherwise.

### GetConfigurableOk

`func (o *TenantDTO) GetConfigurableOk() (*bool, bool)`

GetConfigurableOk returns a tuple with the Configurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurable

`func (o *TenantDTO) SetConfigurable(v bool)`

SetConfigurable sets Configurable field to given value.

### HasConfigurable

`func (o *TenantDTO) HasConfigurable() bool`

HasConfigurable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


