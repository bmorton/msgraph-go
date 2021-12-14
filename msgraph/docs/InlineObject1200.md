# InlineObject1200

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequireSignIn** | Pointer to **NullableBool** |  | [optional] [default to false]
**Roles** | Pointer to **[]string** |  | [optional] 
**SendInvitation** | Pointer to **NullableBool** |  | [optional] [default to false]
**Message** | Pointer to **NullableString** |  | [optional] 
**Recipients** | Pointer to [**[]MicrosoftGraphDriveRecipient**](MicrosoftGraphDriveRecipient.md) |  | [optional] 
**ExpirationDateTime** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInlineObject1200

`func NewInlineObject1200() *InlineObject1200`

NewInlineObject1200 instantiates a new InlineObject1200 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject1200WithDefaults

`func NewInlineObject1200WithDefaults() *InlineObject1200`

NewInlineObject1200WithDefaults instantiates a new InlineObject1200 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequireSignIn

`func (o *InlineObject1200) GetRequireSignIn() bool`

GetRequireSignIn returns the RequireSignIn field if non-nil, zero value otherwise.

### GetRequireSignInOk

`func (o *InlineObject1200) GetRequireSignInOk() (*bool, bool)`

GetRequireSignInOk returns a tuple with the RequireSignIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSignIn

`func (o *InlineObject1200) SetRequireSignIn(v bool)`

SetRequireSignIn sets RequireSignIn field to given value.

### HasRequireSignIn

`func (o *InlineObject1200) HasRequireSignIn() bool`

HasRequireSignIn returns a boolean if a field has been set.

### SetRequireSignInNil

`func (o *InlineObject1200) SetRequireSignInNil(b bool)`

 SetRequireSignInNil sets the value for RequireSignIn to be an explicit nil

### UnsetRequireSignIn
`func (o *InlineObject1200) UnsetRequireSignIn()`

UnsetRequireSignIn ensures that no value is present for RequireSignIn, not even an explicit nil
### GetRoles

`func (o *InlineObject1200) GetRoles() []*string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *InlineObject1200) GetRolesOk() (*[]*string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *InlineObject1200) SetRoles(v []*string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *InlineObject1200) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSendInvitation

`func (o *InlineObject1200) GetSendInvitation() bool`

GetSendInvitation returns the SendInvitation field if non-nil, zero value otherwise.

### GetSendInvitationOk

`func (o *InlineObject1200) GetSendInvitationOk() (*bool, bool)`

GetSendInvitationOk returns a tuple with the SendInvitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendInvitation

`func (o *InlineObject1200) SetSendInvitation(v bool)`

SetSendInvitation sets SendInvitation field to given value.

### HasSendInvitation

`func (o *InlineObject1200) HasSendInvitation() bool`

HasSendInvitation returns a boolean if a field has been set.

### SetSendInvitationNil

`func (o *InlineObject1200) SetSendInvitationNil(b bool)`

 SetSendInvitationNil sets the value for SendInvitation to be an explicit nil

### UnsetSendInvitation
`func (o *InlineObject1200) UnsetSendInvitation()`

UnsetSendInvitation ensures that no value is present for SendInvitation, not even an explicit nil
### GetMessage

`func (o *InlineObject1200) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InlineObject1200) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InlineObject1200) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InlineObject1200) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *InlineObject1200) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *InlineObject1200) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetRecipients

`func (o *InlineObject1200) GetRecipients() []MicrosoftGraphDriveRecipient`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *InlineObject1200) GetRecipientsOk() (*[]MicrosoftGraphDriveRecipient, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *InlineObject1200) SetRecipients(v []MicrosoftGraphDriveRecipient)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *InlineObject1200) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetExpirationDateTime

`func (o *InlineObject1200) GetExpirationDateTime() string`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *InlineObject1200) GetExpirationDateTimeOk() (*string, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *InlineObject1200) SetExpirationDateTime(v string)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *InlineObject1200) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### SetExpirationDateTimeNil

`func (o *InlineObject1200) SetExpirationDateTimeNil(b bool)`

 SetExpirationDateTimeNil sets the value for ExpirationDateTime to be an explicit nil

### UnsetExpirationDateTime
`func (o *InlineObject1200) UnsetExpirationDateTime()`

UnsetExpirationDateTime ensures that no value is present for ExpirationDateTime, not even an explicit nil
### GetPassword

`func (o *InlineObject1200) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InlineObject1200) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InlineObject1200) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *InlineObject1200) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *InlineObject1200) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *InlineObject1200) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


