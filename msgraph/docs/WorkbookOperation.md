# WorkbookOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**AnyOfmicrosoftGraphWorkbookOperationError**](anyOf&lt;microsoft.graph.workbookOperationError&gt;.md) | The error returned by the operation. | [optional] 
**ResourceLocation** | Pointer to **NullableString** | The resource URI for the result. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphWorkbookOperationStatus**](anyOf&lt;microsoft.graph.workbookOperationStatus&gt;.md) | The current status of the operation. Possible values are: NotStarted, Running, Completed, Failed. | [optional] 

## Methods

### NewWorkbookOperation

`func NewWorkbookOperation() *WorkbookOperation`

NewWorkbookOperation instantiates a new WorkbookOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkbookOperationWithDefaults

`func NewWorkbookOperationWithDefaults() *WorkbookOperation`

NewWorkbookOperationWithDefaults instantiates a new WorkbookOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *WorkbookOperation) GetError() AnyOfmicrosoftGraphWorkbookOperationError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WorkbookOperation) GetErrorOk() (*AnyOfmicrosoftGraphWorkbookOperationError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WorkbookOperation) SetError(v AnyOfmicrosoftGraphWorkbookOperationError)`

SetError sets Error field to given value.

### HasError

`func (o *WorkbookOperation) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *WorkbookOperation) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *WorkbookOperation) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetResourceLocation

`func (o *WorkbookOperation) GetResourceLocation() string`

GetResourceLocation returns the ResourceLocation field if non-nil, zero value otherwise.

### GetResourceLocationOk

`func (o *WorkbookOperation) GetResourceLocationOk() (*string, bool)`

GetResourceLocationOk returns a tuple with the ResourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLocation

`func (o *WorkbookOperation) SetResourceLocation(v string)`

SetResourceLocation sets ResourceLocation field to given value.

### HasResourceLocation

`func (o *WorkbookOperation) HasResourceLocation() bool`

HasResourceLocation returns a boolean if a field has been set.

### SetResourceLocationNil

`func (o *WorkbookOperation) SetResourceLocationNil(b bool)`

 SetResourceLocationNil sets the value for ResourceLocation to be an explicit nil

### UnsetResourceLocation
`func (o *WorkbookOperation) UnsetResourceLocation()`

UnsetResourceLocation ensures that no value is present for ResourceLocation, not even an explicit nil
### GetStatus

`func (o *WorkbookOperation) GetStatus() AnyOfmicrosoftGraphWorkbookOperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkbookOperation) GetStatusOk() (*AnyOfmicrosoftGraphWorkbookOperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkbookOperation) SetStatus(v AnyOfmicrosoftGraphWorkbookOperationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkbookOperation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *WorkbookOperation) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *WorkbookOperation) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


