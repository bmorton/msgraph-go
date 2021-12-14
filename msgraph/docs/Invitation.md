# Invitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvitedUserDisplayName** | Pointer to **NullableString** | The display name of the user being invited. | [optional] 
**InvitedUserEmailAddress** | Pointer to **string** | The email address of the user being invited. Required. The following special characters are not permitted in the email address:Tilde (~)Exclamation point (!)Number sign (#)Dollar sign ($)Percent (%)Circumflex (^)Ampersand (&amp;)Asterisk (*)Parentheses (( ))Plus sign (+)Equal sign (&#x3D;)Brackets ([ ])Braces ({ })Backslash (/)Slash mark (/)Pipe (/|)Semicolon (;)Colon (:)Quotation marks (&#39;)Angle brackets (&lt; &gt;)Question mark (?)Comma (,)However, the following exceptions apply:A period (.) or a hyphen (-) is permitted anywhere in the user name, except at the beginning or end of the name.An underscore (_) is permitted anywhere in the user name. This includes at the beginning or end of the name. | [optional] 
**InvitedUserMessageInfo** | Pointer to [**AnyOfmicrosoftGraphInvitedUserMessageInfo**](anyOf&lt;microsoft.graph.invitedUserMessageInfo&gt;.md) | Additional configuration for the message being sent to the invited user, including customizing message text, language and cc recipient list. | [optional] 
**InvitedUserType** | Pointer to **NullableString** | The userType of the user being invited. By default, this is Guest. You can invite as Member if you are a company administrator. | [optional] 
**InviteRedeemUrl** | Pointer to **NullableString** | The URL the user can use to redeem their invitation. Read-only. | [optional] 
**InviteRedirectUrl** | Pointer to **string** | The URL the user should be redirected to once the invitation is redeemed. Required. | [optional] 
**SendInvitationMessage** | Pointer to **NullableBool** | Indicates whether an email should be sent to the user being invited. The default is false. | [optional] 
**Status** | Pointer to **NullableString** | The status of the invitation. Possible values are: PendingAcceptance, Completed, InProgress, and Error. | [optional] 
**InvitedUser** | Pointer to [**AnyOfmicrosoftGraphUser**](anyOf&lt;microsoft.graph.user&gt;.md) | The user created as part of the invitation creation. Read-Only | [optional] 

## Methods

### NewInvitation

`func NewInvitation() *Invitation`

NewInvitation instantiates a new Invitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationWithDefaults

`func NewInvitationWithDefaults() *Invitation`

NewInvitationWithDefaults instantiates a new Invitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvitedUserDisplayName

`func (o *Invitation) GetInvitedUserDisplayName() string`

GetInvitedUserDisplayName returns the InvitedUserDisplayName field if non-nil, zero value otherwise.

### GetInvitedUserDisplayNameOk

`func (o *Invitation) GetInvitedUserDisplayNameOk() (*string, bool)`

GetInvitedUserDisplayNameOk returns a tuple with the InvitedUserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedUserDisplayName

`func (o *Invitation) SetInvitedUserDisplayName(v string)`

SetInvitedUserDisplayName sets InvitedUserDisplayName field to given value.

### HasInvitedUserDisplayName

`func (o *Invitation) HasInvitedUserDisplayName() bool`

HasInvitedUserDisplayName returns a boolean if a field has been set.

### SetInvitedUserDisplayNameNil

`func (o *Invitation) SetInvitedUserDisplayNameNil(b bool)`

 SetInvitedUserDisplayNameNil sets the value for InvitedUserDisplayName to be an explicit nil

### UnsetInvitedUserDisplayName
`func (o *Invitation) UnsetInvitedUserDisplayName()`

UnsetInvitedUserDisplayName ensures that no value is present for InvitedUserDisplayName, not even an explicit nil
### GetInvitedUserEmailAddress

`func (o *Invitation) GetInvitedUserEmailAddress() string`

GetInvitedUserEmailAddress returns the InvitedUserEmailAddress field if non-nil, zero value otherwise.

### GetInvitedUserEmailAddressOk

`func (o *Invitation) GetInvitedUserEmailAddressOk() (*string, bool)`

GetInvitedUserEmailAddressOk returns a tuple with the InvitedUserEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedUserEmailAddress

`func (o *Invitation) SetInvitedUserEmailAddress(v string)`

SetInvitedUserEmailAddress sets InvitedUserEmailAddress field to given value.

### HasInvitedUserEmailAddress

`func (o *Invitation) HasInvitedUserEmailAddress() bool`

HasInvitedUserEmailAddress returns a boolean if a field has been set.

### GetInvitedUserMessageInfo

`func (o *Invitation) GetInvitedUserMessageInfo() AnyOfmicrosoftGraphInvitedUserMessageInfo`

GetInvitedUserMessageInfo returns the InvitedUserMessageInfo field if non-nil, zero value otherwise.

### GetInvitedUserMessageInfoOk

`func (o *Invitation) GetInvitedUserMessageInfoOk() (*AnyOfmicrosoftGraphInvitedUserMessageInfo, bool)`

GetInvitedUserMessageInfoOk returns a tuple with the InvitedUserMessageInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedUserMessageInfo

`func (o *Invitation) SetInvitedUserMessageInfo(v AnyOfmicrosoftGraphInvitedUserMessageInfo)`

SetInvitedUserMessageInfo sets InvitedUserMessageInfo field to given value.

### HasInvitedUserMessageInfo

`func (o *Invitation) HasInvitedUserMessageInfo() bool`

HasInvitedUserMessageInfo returns a boolean if a field has been set.

### SetInvitedUserMessageInfoNil

`func (o *Invitation) SetInvitedUserMessageInfoNil(b bool)`

 SetInvitedUserMessageInfoNil sets the value for InvitedUserMessageInfo to be an explicit nil

### UnsetInvitedUserMessageInfo
`func (o *Invitation) UnsetInvitedUserMessageInfo()`

UnsetInvitedUserMessageInfo ensures that no value is present for InvitedUserMessageInfo, not even an explicit nil
### GetInvitedUserType

`func (o *Invitation) GetInvitedUserType() string`

GetInvitedUserType returns the InvitedUserType field if non-nil, zero value otherwise.

### GetInvitedUserTypeOk

`func (o *Invitation) GetInvitedUserTypeOk() (*string, bool)`

GetInvitedUserTypeOk returns a tuple with the InvitedUserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedUserType

`func (o *Invitation) SetInvitedUserType(v string)`

SetInvitedUserType sets InvitedUserType field to given value.

### HasInvitedUserType

`func (o *Invitation) HasInvitedUserType() bool`

HasInvitedUserType returns a boolean if a field has been set.

### SetInvitedUserTypeNil

`func (o *Invitation) SetInvitedUserTypeNil(b bool)`

 SetInvitedUserTypeNil sets the value for InvitedUserType to be an explicit nil

### UnsetInvitedUserType
`func (o *Invitation) UnsetInvitedUserType()`

UnsetInvitedUserType ensures that no value is present for InvitedUserType, not even an explicit nil
### GetInviteRedeemUrl

`func (o *Invitation) GetInviteRedeemUrl() string`

GetInviteRedeemUrl returns the InviteRedeemUrl field if non-nil, zero value otherwise.

### GetInviteRedeemUrlOk

`func (o *Invitation) GetInviteRedeemUrlOk() (*string, bool)`

GetInviteRedeemUrlOk returns a tuple with the InviteRedeemUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteRedeemUrl

`func (o *Invitation) SetInviteRedeemUrl(v string)`

SetInviteRedeemUrl sets InviteRedeemUrl field to given value.

### HasInviteRedeemUrl

`func (o *Invitation) HasInviteRedeemUrl() bool`

HasInviteRedeemUrl returns a boolean if a field has been set.

### SetInviteRedeemUrlNil

`func (o *Invitation) SetInviteRedeemUrlNil(b bool)`

 SetInviteRedeemUrlNil sets the value for InviteRedeemUrl to be an explicit nil

### UnsetInviteRedeemUrl
`func (o *Invitation) UnsetInviteRedeemUrl()`

UnsetInviteRedeemUrl ensures that no value is present for InviteRedeemUrl, not even an explicit nil
### GetInviteRedirectUrl

`func (o *Invitation) GetInviteRedirectUrl() string`

GetInviteRedirectUrl returns the InviteRedirectUrl field if non-nil, zero value otherwise.

### GetInviteRedirectUrlOk

`func (o *Invitation) GetInviteRedirectUrlOk() (*string, bool)`

GetInviteRedirectUrlOk returns a tuple with the InviteRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteRedirectUrl

`func (o *Invitation) SetInviteRedirectUrl(v string)`

SetInviteRedirectUrl sets InviteRedirectUrl field to given value.

### HasInviteRedirectUrl

`func (o *Invitation) HasInviteRedirectUrl() bool`

HasInviteRedirectUrl returns a boolean if a field has been set.

### GetSendInvitationMessage

`func (o *Invitation) GetSendInvitationMessage() bool`

GetSendInvitationMessage returns the SendInvitationMessage field if non-nil, zero value otherwise.

### GetSendInvitationMessageOk

`func (o *Invitation) GetSendInvitationMessageOk() (*bool, bool)`

GetSendInvitationMessageOk returns a tuple with the SendInvitationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendInvitationMessage

`func (o *Invitation) SetSendInvitationMessage(v bool)`

SetSendInvitationMessage sets SendInvitationMessage field to given value.

### HasSendInvitationMessage

`func (o *Invitation) HasSendInvitationMessage() bool`

HasSendInvitationMessage returns a boolean if a field has been set.

### SetSendInvitationMessageNil

`func (o *Invitation) SetSendInvitationMessageNil(b bool)`

 SetSendInvitationMessageNil sets the value for SendInvitationMessage to be an explicit nil

### UnsetSendInvitationMessage
`func (o *Invitation) UnsetSendInvitationMessage()`

UnsetSendInvitationMessage ensures that no value is present for SendInvitationMessage, not even an explicit nil
### GetStatus

`func (o *Invitation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Invitation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Invitation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Invitation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Invitation) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Invitation) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetInvitedUser

`func (o *Invitation) GetInvitedUser() AnyOfmicrosoftGraphUser`

GetInvitedUser returns the InvitedUser field if non-nil, zero value otherwise.

### GetInvitedUserOk

`func (o *Invitation) GetInvitedUserOk() (*AnyOfmicrosoftGraphUser, bool)`

GetInvitedUserOk returns a tuple with the InvitedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedUser

`func (o *Invitation) SetInvitedUser(v AnyOfmicrosoftGraphUser)`

SetInvitedUser sets InvitedUser field to given value.

### HasInvitedUser

`func (o *Invitation) HasInvitedUser() bool`

HasInvitedUser returns a boolean if a field has been set.

### SetInvitedUserNil

`func (o *Invitation) SetInvitedUserNil(b bool)`

 SetInvitedUserNil sets the value for InvitedUser to be an explicit nil

### UnsetInvitedUser
`func (o *Invitation) UnsetInvitedUser()`

UnsetInvitedUser ensures that no value is present for InvitedUser, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


