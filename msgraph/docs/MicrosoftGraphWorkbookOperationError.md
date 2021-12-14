# MicrosoftGraphWorkbookOperationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **NullableString** | The error code. | [optional] 
**InnerError** | Pointer to [**AnyOfmicrosoftGraphWorkbookOperationError**](anyOf&lt;microsoft.graph.workbookOperationError&gt;.md) |  | [optional] 
**Message** | Pointer to **NullableString** | The error message. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookOperationError

`func NewMicrosoftGraphWorkbookOperationError() *MicrosoftGraphWorkbookOperationError`

NewMicrosoftGraphWorkbookOperationError instantiates a new MicrosoftGraphWorkbookOperationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookOperationErrorWithDefaults

`func NewMicrosoftGraphWorkbookOperationErrorWithDefaults() *MicrosoftGraphWorkbookOperationError`

NewMicrosoftGraphWorkbookOperationErrorWithDefaults instantiates a new MicrosoftGraphWorkbookOperationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *MicrosoftGraphWorkbookOperationError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MicrosoftGraphWorkbookOperationError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MicrosoftGraphWorkbookOperationError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *MicrosoftGraphWorkbookOperationError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *MicrosoftGraphWorkbookOperationError) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *MicrosoftGraphWorkbookOperationError) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetInnerError

`func (o *MicrosoftGraphWorkbookOperationError) GetInnerError() AnyOfmicrosoftGraphWorkbookOperationError`

GetInnerError returns the InnerError field if non-nil, zero value otherwise.

### GetInnerErrorOk

`func (o *MicrosoftGraphWorkbookOperationError) GetInnerErrorOk() (*AnyOfmicrosoftGraphWorkbookOperationError, bool)`

GetInnerErrorOk returns a tuple with the InnerError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerError

`func (o *MicrosoftGraphWorkbookOperationError) SetInnerError(v AnyOfmicrosoftGraphWorkbookOperationError)`

SetInnerError sets InnerError field to given value.

### HasInnerError

`func (o *MicrosoftGraphWorkbookOperationError) HasInnerError() bool`

HasInnerError returns a boolean if a field has been set.

### SetInnerErrorNil

`func (o *MicrosoftGraphWorkbookOperationError) SetInnerErrorNil(b bool)`

 SetInnerErrorNil sets the value for InnerError to be an explicit nil

### UnsetInnerError
`func (o *MicrosoftGraphWorkbookOperationError) UnsetInnerError()`

UnsetInnerError ensures that no value is present for InnerError, not even an explicit nil
### GetMessage

`func (o *MicrosoftGraphWorkbookOperationError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MicrosoftGraphWorkbookOperationError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MicrosoftGraphWorkbookOperationError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MicrosoftGraphWorkbookOperationError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *MicrosoftGraphWorkbookOperationError) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *MicrosoftGraphWorkbookOperationError) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


