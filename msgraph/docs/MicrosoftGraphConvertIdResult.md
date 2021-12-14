# MicrosoftGraphConvertIdResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorDetails** | Pointer to [**AnyOfmicrosoftGraphGenericError**](anyOf&lt;microsoft.graph.genericError&gt;.md) | An error object indicating the reason for the conversion failure. This value is not present if the conversion succeeded. | [optional] 
**SourceId** | Pointer to **NullableString** | The identifier that was converted. This value is the original, un-converted identifier. | [optional] 
**TargetId** | Pointer to **NullableString** | The converted identifier. This value is not present if the conversion failed. | [optional] 

## Methods

### NewMicrosoftGraphConvertIdResult

`func NewMicrosoftGraphConvertIdResult() *MicrosoftGraphConvertIdResult`

NewMicrosoftGraphConvertIdResult instantiates a new MicrosoftGraphConvertIdResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphConvertIdResultWithDefaults

`func NewMicrosoftGraphConvertIdResultWithDefaults() *MicrosoftGraphConvertIdResult`

NewMicrosoftGraphConvertIdResultWithDefaults instantiates a new MicrosoftGraphConvertIdResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorDetails

`func (o *MicrosoftGraphConvertIdResult) GetErrorDetails() AnyOfmicrosoftGraphGenericError`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *MicrosoftGraphConvertIdResult) GetErrorDetailsOk() (*AnyOfmicrosoftGraphGenericError, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *MicrosoftGraphConvertIdResult) SetErrorDetails(v AnyOfmicrosoftGraphGenericError)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *MicrosoftGraphConvertIdResult) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.

### SetErrorDetailsNil

`func (o *MicrosoftGraphConvertIdResult) SetErrorDetailsNil(b bool)`

 SetErrorDetailsNil sets the value for ErrorDetails to be an explicit nil

### UnsetErrorDetails
`func (o *MicrosoftGraphConvertIdResult) UnsetErrorDetails()`

UnsetErrorDetails ensures that no value is present for ErrorDetails, not even an explicit nil
### GetSourceId

`func (o *MicrosoftGraphConvertIdResult) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *MicrosoftGraphConvertIdResult) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *MicrosoftGraphConvertIdResult) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *MicrosoftGraphConvertIdResult) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *MicrosoftGraphConvertIdResult) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *MicrosoftGraphConvertIdResult) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetTargetId

`func (o *MicrosoftGraphConvertIdResult) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *MicrosoftGraphConvertIdResult) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *MicrosoftGraphConvertIdResult) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *MicrosoftGraphConvertIdResult) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *MicrosoftGraphConvertIdResult) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *MicrosoftGraphConvertIdResult) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


