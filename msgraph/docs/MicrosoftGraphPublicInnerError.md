# MicrosoftGraphPublicInnerError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **NullableString** | The error code. | [optional] 
**Details** | Pointer to [**[]AnyOfmicrosoftGraphPublicErrorDetail**](AnyOfmicrosoftGraphPublicErrorDetail.md) | A collection of error details. | [optional] 
**Message** | Pointer to **NullableString** | The error message. | [optional] 
**Target** | Pointer to **NullableString** | The target of the error. | [optional] 

## Methods

### NewMicrosoftGraphPublicInnerError

`func NewMicrosoftGraphPublicInnerError() *MicrosoftGraphPublicInnerError`

NewMicrosoftGraphPublicInnerError instantiates a new MicrosoftGraphPublicInnerError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPublicInnerErrorWithDefaults

`func NewMicrosoftGraphPublicInnerErrorWithDefaults() *MicrosoftGraphPublicInnerError`

NewMicrosoftGraphPublicInnerErrorWithDefaults instantiates a new MicrosoftGraphPublicInnerError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *MicrosoftGraphPublicInnerError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MicrosoftGraphPublicInnerError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MicrosoftGraphPublicInnerError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *MicrosoftGraphPublicInnerError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *MicrosoftGraphPublicInnerError) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *MicrosoftGraphPublicInnerError) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetDetails

`func (o *MicrosoftGraphPublicInnerError) GetDetails() []*AnyOfmicrosoftGraphPublicErrorDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *MicrosoftGraphPublicInnerError) GetDetailsOk() (*[]*AnyOfmicrosoftGraphPublicErrorDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *MicrosoftGraphPublicInnerError) SetDetails(v []*AnyOfmicrosoftGraphPublicErrorDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *MicrosoftGraphPublicInnerError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetMessage

`func (o *MicrosoftGraphPublicInnerError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MicrosoftGraphPublicInnerError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MicrosoftGraphPublicInnerError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MicrosoftGraphPublicInnerError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *MicrosoftGraphPublicInnerError) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *MicrosoftGraphPublicInnerError) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetTarget

`func (o *MicrosoftGraphPublicInnerError) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *MicrosoftGraphPublicInnerError) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *MicrosoftGraphPublicInnerError) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *MicrosoftGraphPublicInnerError) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *MicrosoftGraphPublicInnerError) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *MicrosoftGraphPublicInnerError) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


