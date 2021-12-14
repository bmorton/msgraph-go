# MicrosoftGraphResultInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | The result code. | [optional] 
**Message** | Pointer to **NullableString** | The message. | [optional] 
**Subcode** | Pointer to **int32** | The result sub-code. | [optional] 

## Methods

### NewMicrosoftGraphResultInfo

`func NewMicrosoftGraphResultInfo() *MicrosoftGraphResultInfo`

NewMicrosoftGraphResultInfo instantiates a new MicrosoftGraphResultInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphResultInfoWithDefaults

`func NewMicrosoftGraphResultInfoWithDefaults() *MicrosoftGraphResultInfo`

NewMicrosoftGraphResultInfoWithDefaults instantiates a new MicrosoftGraphResultInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *MicrosoftGraphResultInfo) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MicrosoftGraphResultInfo) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MicrosoftGraphResultInfo) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *MicrosoftGraphResultInfo) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *MicrosoftGraphResultInfo) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MicrosoftGraphResultInfo) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MicrosoftGraphResultInfo) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MicrosoftGraphResultInfo) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *MicrosoftGraphResultInfo) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *MicrosoftGraphResultInfo) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetSubcode

`func (o *MicrosoftGraphResultInfo) GetSubcode() int32`

GetSubcode returns the Subcode field if non-nil, zero value otherwise.

### GetSubcodeOk

`func (o *MicrosoftGraphResultInfo) GetSubcodeOk() (*int32, bool)`

GetSubcodeOk returns a tuple with the Subcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubcode

`func (o *MicrosoftGraphResultInfo) SetSubcode(v int32)`

SetSubcode sets Subcode field to given value.

### HasSubcode

`func (o *MicrosoftGraphResultInfo) HasSubcode() bool`

HasSubcode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


