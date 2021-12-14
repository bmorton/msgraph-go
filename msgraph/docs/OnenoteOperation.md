# OnenoteOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**AnyOfmicrosoftGraphOnenoteOperationError**](anyOf&lt;microsoft.graph.onenoteOperationError&gt;.md) | The error returned by the operation. | [optional] 
**PercentComplete** | Pointer to **NullableString** | The operation percent complete if the operation is still in running status. | [optional] 
**ResourceId** | Pointer to **NullableString** | The resource id. | [optional] 
**ResourceLocation** | Pointer to **NullableString** | The resource URI for the object. For example, the resource URI for a copied page or section. | [optional] 

## Methods

### NewOnenoteOperation

`func NewOnenoteOperation() *OnenoteOperation`

NewOnenoteOperation instantiates a new OnenoteOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnenoteOperationWithDefaults

`func NewOnenoteOperationWithDefaults() *OnenoteOperation`

NewOnenoteOperationWithDefaults instantiates a new OnenoteOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *OnenoteOperation) GetError() AnyOfmicrosoftGraphOnenoteOperationError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OnenoteOperation) GetErrorOk() (*AnyOfmicrosoftGraphOnenoteOperationError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OnenoteOperation) SetError(v AnyOfmicrosoftGraphOnenoteOperationError)`

SetError sets Error field to given value.

### HasError

`func (o *OnenoteOperation) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *OnenoteOperation) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *OnenoteOperation) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetPercentComplete

`func (o *OnenoteOperation) GetPercentComplete() string`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *OnenoteOperation) GetPercentCompleteOk() (*string, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *OnenoteOperation) SetPercentComplete(v string)`

SetPercentComplete sets PercentComplete field to given value.

### HasPercentComplete

`func (o *OnenoteOperation) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.

### SetPercentCompleteNil

`func (o *OnenoteOperation) SetPercentCompleteNil(b bool)`

 SetPercentCompleteNil sets the value for PercentComplete to be an explicit nil

### UnsetPercentComplete
`func (o *OnenoteOperation) UnsetPercentComplete()`

UnsetPercentComplete ensures that no value is present for PercentComplete, not even an explicit nil
### GetResourceId

`func (o *OnenoteOperation) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *OnenoteOperation) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *OnenoteOperation) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *OnenoteOperation) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *OnenoteOperation) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *OnenoteOperation) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetResourceLocation

`func (o *OnenoteOperation) GetResourceLocation() string`

GetResourceLocation returns the ResourceLocation field if non-nil, zero value otherwise.

### GetResourceLocationOk

`func (o *OnenoteOperation) GetResourceLocationOk() (*string, bool)`

GetResourceLocationOk returns a tuple with the ResourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLocation

`func (o *OnenoteOperation) SetResourceLocation(v string)`

SetResourceLocation sets ResourceLocation field to given value.

### HasResourceLocation

`func (o *OnenoteOperation) HasResourceLocation() bool`

HasResourceLocation returns a boolean if a field has been set.

### SetResourceLocationNil

`func (o *OnenoteOperation) SetResourceLocationNil(b bool)`

 SetResourceLocationNil sets the value for ResourceLocation to be an explicit nil

### UnsetResourceLocation
`func (o *OnenoteOperation) UnsetResourceLocation()`

UnsetResourceLocation ensures that no value is present for ResourceLocation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


