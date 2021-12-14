# MicrosoftGraphOperationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **NullableString** | Operation error code. | [optional] 
**Message** | Pointer to **NullableString** | Operation error message. | [optional] 

## Methods

### NewMicrosoftGraphOperationError

`func NewMicrosoftGraphOperationError() *MicrosoftGraphOperationError`

NewMicrosoftGraphOperationError instantiates a new MicrosoftGraphOperationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphOperationErrorWithDefaults

`func NewMicrosoftGraphOperationErrorWithDefaults() *MicrosoftGraphOperationError`

NewMicrosoftGraphOperationErrorWithDefaults instantiates a new MicrosoftGraphOperationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *MicrosoftGraphOperationError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MicrosoftGraphOperationError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MicrosoftGraphOperationError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *MicrosoftGraphOperationError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *MicrosoftGraphOperationError) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *MicrosoftGraphOperationError) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetMessage

`func (o *MicrosoftGraphOperationError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MicrosoftGraphOperationError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MicrosoftGraphOperationError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MicrosoftGraphOperationError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *MicrosoftGraphOperationError) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *MicrosoftGraphOperationError) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


