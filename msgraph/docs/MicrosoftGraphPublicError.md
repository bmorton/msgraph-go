# MicrosoftGraphPublicError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **NullableString** | Represents the error code. | [optional] 
**Details** | Pointer to [**[]AnyOfmicrosoftGraphPublicErrorDetail**](AnyOfmicrosoftGraphPublicErrorDetail.md) | Details of the error. | [optional] 
**InnerError** | Pointer to [**AnyOfmicrosoftGraphPublicInnerError**](anyOf&lt;microsoft.graph.publicInnerError&gt;.md) | Details of the inner error. | [optional] 
**Message** | Pointer to **NullableString** | A non-localized message for the developer. | [optional] 
**Target** | Pointer to **NullableString** | The target of the error. | [optional] 

## Methods

### NewMicrosoftGraphPublicError

`func NewMicrosoftGraphPublicError() *MicrosoftGraphPublicError`

NewMicrosoftGraphPublicError instantiates a new MicrosoftGraphPublicError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPublicErrorWithDefaults

`func NewMicrosoftGraphPublicErrorWithDefaults() *MicrosoftGraphPublicError`

NewMicrosoftGraphPublicErrorWithDefaults instantiates a new MicrosoftGraphPublicError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *MicrosoftGraphPublicError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MicrosoftGraphPublicError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MicrosoftGraphPublicError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *MicrosoftGraphPublicError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *MicrosoftGraphPublicError) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *MicrosoftGraphPublicError) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetDetails

`func (o *MicrosoftGraphPublicError) GetDetails() []*AnyOfmicrosoftGraphPublicErrorDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *MicrosoftGraphPublicError) GetDetailsOk() (*[]*AnyOfmicrosoftGraphPublicErrorDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *MicrosoftGraphPublicError) SetDetails(v []*AnyOfmicrosoftGraphPublicErrorDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *MicrosoftGraphPublicError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetInnerError

`func (o *MicrosoftGraphPublicError) GetInnerError() AnyOfmicrosoftGraphPublicInnerError`

GetInnerError returns the InnerError field if non-nil, zero value otherwise.

### GetInnerErrorOk

`func (o *MicrosoftGraphPublicError) GetInnerErrorOk() (*AnyOfmicrosoftGraphPublicInnerError, bool)`

GetInnerErrorOk returns a tuple with the InnerError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerError

`func (o *MicrosoftGraphPublicError) SetInnerError(v AnyOfmicrosoftGraphPublicInnerError)`

SetInnerError sets InnerError field to given value.

### HasInnerError

`func (o *MicrosoftGraphPublicError) HasInnerError() bool`

HasInnerError returns a boolean if a field has been set.

### SetInnerErrorNil

`func (o *MicrosoftGraphPublicError) SetInnerErrorNil(b bool)`

 SetInnerErrorNil sets the value for InnerError to be an explicit nil

### UnsetInnerError
`func (o *MicrosoftGraphPublicError) UnsetInnerError()`

UnsetInnerError ensures that no value is present for InnerError, not even an explicit nil
### GetMessage

`func (o *MicrosoftGraphPublicError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MicrosoftGraphPublicError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MicrosoftGraphPublicError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MicrosoftGraphPublicError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *MicrosoftGraphPublicError) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *MicrosoftGraphPublicError) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetTarget

`func (o *MicrosoftGraphPublicError) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *MicrosoftGraphPublicError) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *MicrosoftGraphPublicError) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *MicrosoftGraphPublicError) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *MicrosoftGraphPublicError) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *MicrosoftGraphPublicError) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


