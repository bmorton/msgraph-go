# MicrosoftGraphSharingInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **NullableString** | The email address provided for the recipient of the sharing invitation. Read-only. | [optional] 
**InvitedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Provides information about who sent the invitation that created this permission, if that information is available. Read-only. | [optional] 
**RedeemedBy** | Pointer to **NullableString** |  | [optional] 
**SignInRequired** | Pointer to **NullableBool** | If true the recipient of the invitation needs to sign in in order to access the shared item. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphSharingInvitation

`func NewMicrosoftGraphSharingInvitation() *MicrosoftGraphSharingInvitation`

NewMicrosoftGraphSharingInvitation instantiates a new MicrosoftGraphSharingInvitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSharingInvitationWithDefaults

`func NewMicrosoftGraphSharingInvitationWithDefaults() *MicrosoftGraphSharingInvitation`

NewMicrosoftGraphSharingInvitationWithDefaults instantiates a new MicrosoftGraphSharingInvitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *MicrosoftGraphSharingInvitation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MicrosoftGraphSharingInvitation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MicrosoftGraphSharingInvitation) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MicrosoftGraphSharingInvitation) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *MicrosoftGraphSharingInvitation) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *MicrosoftGraphSharingInvitation) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetInvitedBy

`func (o *MicrosoftGraphSharingInvitation) GetInvitedBy() AnyOfmicrosoftGraphIdentitySet`

GetInvitedBy returns the InvitedBy field if non-nil, zero value otherwise.

### GetInvitedByOk

`func (o *MicrosoftGraphSharingInvitation) GetInvitedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetInvitedByOk returns a tuple with the InvitedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedBy

`func (o *MicrosoftGraphSharingInvitation) SetInvitedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetInvitedBy sets InvitedBy field to given value.

### HasInvitedBy

`func (o *MicrosoftGraphSharingInvitation) HasInvitedBy() bool`

HasInvitedBy returns a boolean if a field has been set.

### SetInvitedByNil

`func (o *MicrosoftGraphSharingInvitation) SetInvitedByNil(b bool)`

 SetInvitedByNil sets the value for InvitedBy to be an explicit nil

### UnsetInvitedBy
`func (o *MicrosoftGraphSharingInvitation) UnsetInvitedBy()`

UnsetInvitedBy ensures that no value is present for InvitedBy, not even an explicit nil
### GetRedeemedBy

`func (o *MicrosoftGraphSharingInvitation) GetRedeemedBy() string`

GetRedeemedBy returns the RedeemedBy field if non-nil, zero value otherwise.

### GetRedeemedByOk

`func (o *MicrosoftGraphSharingInvitation) GetRedeemedByOk() (*string, bool)`

GetRedeemedByOk returns a tuple with the RedeemedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeemedBy

`func (o *MicrosoftGraphSharingInvitation) SetRedeemedBy(v string)`

SetRedeemedBy sets RedeemedBy field to given value.

### HasRedeemedBy

`func (o *MicrosoftGraphSharingInvitation) HasRedeemedBy() bool`

HasRedeemedBy returns a boolean if a field has been set.

### SetRedeemedByNil

`func (o *MicrosoftGraphSharingInvitation) SetRedeemedByNil(b bool)`

 SetRedeemedByNil sets the value for RedeemedBy to be an explicit nil

### UnsetRedeemedBy
`func (o *MicrosoftGraphSharingInvitation) UnsetRedeemedBy()`

UnsetRedeemedBy ensures that no value is present for RedeemedBy, not even an explicit nil
### GetSignInRequired

`func (o *MicrosoftGraphSharingInvitation) GetSignInRequired() bool`

GetSignInRequired returns the SignInRequired field if non-nil, zero value otherwise.

### GetSignInRequiredOk

`func (o *MicrosoftGraphSharingInvitation) GetSignInRequiredOk() (*bool, bool)`

GetSignInRequiredOk returns a tuple with the SignInRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInRequired

`func (o *MicrosoftGraphSharingInvitation) SetSignInRequired(v bool)`

SetSignInRequired sets SignInRequired field to given value.

### HasSignInRequired

`func (o *MicrosoftGraphSharingInvitation) HasSignInRequired() bool`

HasSignInRequired returns a boolean if a field has been set.

### SetSignInRequiredNil

`func (o *MicrosoftGraphSharingInvitation) SetSignInRequiredNil(b bool)`

 SetSignInRequiredNil sets the value for SignInRequired to be an explicit nil

### UnsetSignInRequired
`func (o *MicrosoftGraphSharingInvitation) UnsetSignInRequired()`

UnsetSignInRequired ensures that no value is present for SignInRequired, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


