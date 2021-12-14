# MicrosoftGraphActionResultPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**AnyOfmicrosoftGraphPublicError**](anyOf&lt;microsoft.graph.publicError&gt;.md) | The error that occurred, if any, during the course of the bulk operation. | [optional] 

## Methods

### NewMicrosoftGraphActionResultPart

`func NewMicrosoftGraphActionResultPart() *MicrosoftGraphActionResultPart`

NewMicrosoftGraphActionResultPart instantiates a new MicrosoftGraphActionResultPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphActionResultPartWithDefaults

`func NewMicrosoftGraphActionResultPartWithDefaults() *MicrosoftGraphActionResultPart`

NewMicrosoftGraphActionResultPartWithDefaults instantiates a new MicrosoftGraphActionResultPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *MicrosoftGraphActionResultPart) GetError() AnyOfmicrosoftGraphPublicError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MicrosoftGraphActionResultPart) GetErrorOk() (*AnyOfmicrosoftGraphPublicError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MicrosoftGraphActionResultPart) SetError(v AnyOfmicrosoftGraphPublicError)`

SetError sets Error field to given value.

### HasError

`func (o *MicrosoftGraphActionResultPart) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *MicrosoftGraphActionResultPart) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *MicrosoftGraphActionResultPart) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


