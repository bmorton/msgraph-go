# MicrosoftGraphRelatedContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessConsent** | Pointer to **NullableBool** | Indicates whether the user has been consented to access student data. | [optional] 
**DisplayName** | Pointer to **string** | Name of the contact. Required. | [optional] 
**EmailAddress** | Pointer to **string** | Primary email address of the contact. | [optional] 
**MobilePhone** | Pointer to **NullableString** | Mobile phone number of the contact. | [optional] 
**Relationship** | Pointer to [**AnyOfmicrosoftGraphContactRelationship**](anyOf&lt;microsoft.graph.contactRelationship&gt;.md) | Relationship to the user. Possible values are parent, relative, aide, doctor, guardian, child, other, unknownFutureValue. | [optional] 

## Methods

### NewMicrosoftGraphRelatedContact

`func NewMicrosoftGraphRelatedContact() *MicrosoftGraphRelatedContact`

NewMicrosoftGraphRelatedContact instantiates a new MicrosoftGraphRelatedContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphRelatedContactWithDefaults

`func NewMicrosoftGraphRelatedContactWithDefaults() *MicrosoftGraphRelatedContact`

NewMicrosoftGraphRelatedContactWithDefaults instantiates a new MicrosoftGraphRelatedContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessConsent

`func (o *MicrosoftGraphRelatedContact) GetAccessConsent() bool`

GetAccessConsent returns the AccessConsent field if non-nil, zero value otherwise.

### GetAccessConsentOk

`func (o *MicrosoftGraphRelatedContact) GetAccessConsentOk() (*bool, bool)`

GetAccessConsentOk returns a tuple with the AccessConsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessConsent

`func (o *MicrosoftGraphRelatedContact) SetAccessConsent(v bool)`

SetAccessConsent sets AccessConsent field to given value.

### HasAccessConsent

`func (o *MicrosoftGraphRelatedContact) HasAccessConsent() bool`

HasAccessConsent returns a boolean if a field has been set.

### SetAccessConsentNil

`func (o *MicrosoftGraphRelatedContact) SetAccessConsentNil(b bool)`

 SetAccessConsentNil sets the value for AccessConsent to be an explicit nil

### UnsetAccessConsent
`func (o *MicrosoftGraphRelatedContact) UnsetAccessConsent()`

UnsetAccessConsent ensures that no value is present for AccessConsent, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphRelatedContact) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphRelatedContact) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphRelatedContact) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphRelatedContact) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmailAddress

`func (o *MicrosoftGraphRelatedContact) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *MicrosoftGraphRelatedContact) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *MicrosoftGraphRelatedContact) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *MicrosoftGraphRelatedContact) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetMobilePhone

`func (o *MicrosoftGraphRelatedContact) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *MicrosoftGraphRelatedContact) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *MicrosoftGraphRelatedContact) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *MicrosoftGraphRelatedContact) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### SetMobilePhoneNil

`func (o *MicrosoftGraphRelatedContact) SetMobilePhoneNil(b bool)`

 SetMobilePhoneNil sets the value for MobilePhone to be an explicit nil

### UnsetMobilePhone
`func (o *MicrosoftGraphRelatedContact) UnsetMobilePhone()`

UnsetMobilePhone ensures that no value is present for MobilePhone, not even an explicit nil
### GetRelationship

`func (o *MicrosoftGraphRelatedContact) GetRelationship() AnyOfmicrosoftGraphContactRelationship`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *MicrosoftGraphRelatedContact) GetRelationshipOk() (*AnyOfmicrosoftGraphContactRelationship, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *MicrosoftGraphRelatedContact) SetRelationship(v AnyOfmicrosoftGraphContactRelationship)`

SetRelationship sets Relationship field to given value.

### HasRelationship

`func (o *MicrosoftGraphRelatedContact) HasRelationship() bool`

HasRelationship returns a boolean if a field has been set.

### SetRelationshipNil

`func (o *MicrosoftGraphRelatedContact) SetRelationshipNil(b bool)`

 SetRelationshipNil sets the value for Relationship to be an explicit nil

### UnsetRelationship
`func (o *MicrosoftGraphRelatedContact) UnsetRelationship()`

UnsetRelationship ensures that no value is present for Relationship, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


