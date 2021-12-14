# MicrosoftGraphFreeBusyError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **NullableString** | Describes the error. | [optional] 
**ResponseCode** | Pointer to **NullableString** | The response code from querying for the availability of the user, distribution list, or resource. | [optional] 

## Methods

### NewMicrosoftGraphFreeBusyError

`func NewMicrosoftGraphFreeBusyError() *MicrosoftGraphFreeBusyError`

NewMicrosoftGraphFreeBusyError instantiates a new MicrosoftGraphFreeBusyError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphFreeBusyErrorWithDefaults

`func NewMicrosoftGraphFreeBusyErrorWithDefaults() *MicrosoftGraphFreeBusyError`

NewMicrosoftGraphFreeBusyErrorWithDefaults instantiates a new MicrosoftGraphFreeBusyError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *MicrosoftGraphFreeBusyError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MicrosoftGraphFreeBusyError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MicrosoftGraphFreeBusyError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MicrosoftGraphFreeBusyError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *MicrosoftGraphFreeBusyError) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *MicrosoftGraphFreeBusyError) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetResponseCode

`func (o *MicrosoftGraphFreeBusyError) GetResponseCode() string`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *MicrosoftGraphFreeBusyError) GetResponseCodeOk() (*string, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *MicrosoftGraphFreeBusyError) SetResponseCode(v string)`

SetResponseCode sets ResponseCode field to given value.

### HasResponseCode

`func (o *MicrosoftGraphFreeBusyError) HasResponseCode() bool`

HasResponseCode returns a boolean if a field has been set.

### SetResponseCodeNil

`func (o *MicrosoftGraphFreeBusyError) SetResponseCodeNil(b bool)`

 SetResponseCodeNil sets the value for ResponseCode to be an explicit nil

### UnsetResponseCode
`func (o *MicrosoftGraphFreeBusyError) UnsetResponseCode()`

UnsetResponseCode ensures that no value is present for ResponseCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


