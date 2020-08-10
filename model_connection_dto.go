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

// ConnectionDto struct for ConnectionDto
type ConnectionDto struct {
	// The id of the component.
	Id string `json:"id,omitempty"`
	// The ID of the corresponding component that is under version control
	VersionedComponentId string `json:"versionedComponentId,omitempty"`
	// The id of parent process group of this component if applicable.
	ParentGroupId string         `json:"parentGroupId,omitempty"`
	Position      PositionDto    `json:"position,omitempty"`
	Source        ConnectableDto `json:"source,omitempty"`
	Destination   ConnectableDto `json:"destination,omitempty"`
	// The name of the connection.
	Name string `json:"name,omitempty"`
	// The index of the bend point where to place the connection label.
	LabelIndex int32 `json:"labelIndex,omitempty"`
	// The z index of the connection.
	GetzIndex int64 `json:"getzIndex,omitempty"`
	// The selected relationship that comprise the connection.
	SelectedRelationships []string `json:"selectedRelationships,omitempty"`
	// The relationships that the source of the connection currently supports.
	AvailableRelationships []string `json:"availableRelationships,omitempty"`
	// The object count threshold for determining when back pressure is applied. Updating this value is a passive change in the sense that it won't impact whether existing files over the limit are affected but it does help feeder processors to stop pushing too much into this work queue.
	BackPressureObjectThreshold int64 `json:"backPressureObjectThreshold,omitempty"`
	// The object data size threshold for determining when back pressure is applied. Updating this value is a passive change in the sense that it won't impact whether existing files over the limit are affected but it does help feeder processors to stop pushing too much into this work queue.
	BackPressureDataSizeThreshold string `json:"backPressureDataSizeThreshold,omitempty"`
	// The amount of time a flow file may be in the flow before it will be automatically aged out of the flow. Once a flow file reaches this age it will be terminated from the flow the next time a processor attempts to start work on it.
	FlowFileExpiration string `json:"flowFileExpiration,omitempty"`
	// The comparators used to prioritize the queue.
	Prioritizers []string `json:"prioritizers,omitempty"`
	// The bend points on the connection.
	Bends []PositionDto `json:"bends,omitempty"`
	// How to load balance the data in this Connection across the nodes in the cluster.
	LoadBalanceStrategy string `json:"loadBalanceStrategy,omitempty"`
	// The FlowFile Attribute to use for determining which node a FlowFile will go to if the Load Balancing Strategy is set to PARTITION_BY_ATTRIBUTE
	LoadBalancePartitionAttribute string `json:"loadBalancePartitionAttribute,omitempty"`
	// Whether or not data should be compressed when being transferred between nodes in the cluster.
	LoadBalanceCompression string `json:"loadBalanceCompression,omitempty"`
	// The current status of the Connection's Load Balancing Activities. Status can indicate that Load Balancing is not configured for the connection, that Load Balancing is configured but inactive (not currently transferring data to another node), or that Load Balancing is configured and actively transferring data to another node.
	LoadBalanceStatus string `json:"loadBalanceStatus,omitempty"`
}
