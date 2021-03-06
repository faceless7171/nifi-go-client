# NodeConnectionStatusSnapshotDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | Pointer to **string** | The unique ID that identifies the node | [optional] 
**Address** | Pointer to **string** | The API address of the node | [optional] 
**ApiPort** | Pointer to **int32** | The API port used to communicate with the node | [optional] 
**StatusSnapshot** | Pointer to [**ConnectionStatusSnapshotDTO**](ConnectionStatusSnapshotDTO.md) |  | [optional] 

## Methods

### NewNodeConnectionStatusSnapshotDTO

`func NewNodeConnectionStatusSnapshotDTO() *NodeConnectionStatusSnapshotDTO`

NewNodeConnectionStatusSnapshotDTO instantiates a new NodeConnectionStatusSnapshotDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeConnectionStatusSnapshotDTOWithDefaults

`func NewNodeConnectionStatusSnapshotDTOWithDefaults() *NodeConnectionStatusSnapshotDTO`

NewNodeConnectionStatusSnapshotDTOWithDefaults instantiates a new NodeConnectionStatusSnapshotDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *NodeConnectionStatusSnapshotDTO) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *NodeConnectionStatusSnapshotDTO) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *NodeConnectionStatusSnapshotDTO) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *NodeConnectionStatusSnapshotDTO) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetAddress

`func (o *NodeConnectionStatusSnapshotDTO) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NodeConnectionStatusSnapshotDTO) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NodeConnectionStatusSnapshotDTO) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NodeConnectionStatusSnapshotDTO) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetApiPort

`func (o *NodeConnectionStatusSnapshotDTO) GetApiPort() int32`

GetApiPort returns the ApiPort field if non-nil, zero value otherwise.

### GetApiPortOk

`func (o *NodeConnectionStatusSnapshotDTO) GetApiPortOk() (*int32, bool)`

GetApiPortOk returns a tuple with the ApiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPort

`func (o *NodeConnectionStatusSnapshotDTO) SetApiPort(v int32)`

SetApiPort sets ApiPort field to given value.

### HasApiPort

`func (o *NodeConnectionStatusSnapshotDTO) HasApiPort() bool`

HasApiPort returns a boolean if a field has been set.

### GetStatusSnapshot

`func (o *NodeConnectionStatusSnapshotDTO) GetStatusSnapshot() ConnectionStatusSnapshotDTO`

GetStatusSnapshot returns the StatusSnapshot field if non-nil, zero value otherwise.

### GetStatusSnapshotOk

`func (o *NodeConnectionStatusSnapshotDTO) GetStatusSnapshotOk() (*ConnectionStatusSnapshotDTO, bool)`

GetStatusSnapshotOk returns a tuple with the StatusSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusSnapshot

`func (o *NodeConnectionStatusSnapshotDTO) SetStatusSnapshot(v ConnectionStatusSnapshotDTO)`

SetStatusSnapshot sets StatusSnapshot field to given value.

### HasStatusSnapshot

`func (o *NodeConnectionStatusSnapshotDTO) HasStatusSnapshot() bool`

HasStatusSnapshot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


