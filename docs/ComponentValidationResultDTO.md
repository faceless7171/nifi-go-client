# ComponentValidationResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessGroupId** | **string** | The UUID of the Process Group that this component is in | [optional] 
**Id** | **string** | The UUID of this component | [optional] 
**ReferenceType** | **string** | The type of this component | [optional] 
**Name** | **string** | The name of this component. | [optional] 
**State** | **string** | The scheduled state of a processor or reporting task referencing a controller service. If this component is another controller service, this field represents the controller service state. | [optional] 
**ActiveThreadCount** | **int32** | The number of active threads for the referencing component. | [optional] 
**ValidationErrors** | **[]string** | The validation errors for the component. | [optional] 
**CurrentlyValid** | **bool** | Whether or not the component is currently valid | [optional] 
**ResultsValid** | **bool** | Whether or not the component will be valid if the Parameter Context is changed | [optional] 
**ResultantValidationErrors** | **[]string** | The validation errors that will apply to the component if the Parameter Context is changed | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

