# MicrosoftGraphRecipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | Pointer to [**AnyOfmicrosoftGraphEmailAddress**](anyOf&lt;microsoft.graph.emailAddress&gt;.md) | The recipient&#39;s email address. | [optional] 

## Methods

### NewMicrosoftGraphRecipient

`func NewMicrosoftGraphRecipient() *MicrosoftGraphRecipient`

NewMicrosoftGraphRecipient instantiates a new MicrosoftGraphRecipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphRecipientWithDefaults

`func NewMicrosoftGraphRecipientWithDefaults() *MicrosoftGraphRecipient`

NewMicrosoftGraphRecipientWithDefaults instantiates a new MicrosoftGraphRecipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *MicrosoftGraphRecipient) GetEmailAddress() AnyOfmicrosoftGraphEmailAddress`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *MicrosoftGraphRecipient) GetEmailAddressOk() (*AnyOfmicrosoftGraphEmailAddress, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *MicrosoftGraphRecipient) SetEmailAddress(v AnyOfmicrosoftGraphEmailAddress)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *MicrosoftGraphRecipient) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *MicrosoftGraphRecipient) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *MicrosoftGraphRecipient) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


