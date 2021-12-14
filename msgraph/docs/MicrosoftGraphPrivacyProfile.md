# MicrosoftGraphPrivacyProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactEmail** | Pointer to **NullableString** | A valid smtp email address for the privacy statement contact. Not required. | [optional] 
**StatementUrl** | Pointer to **NullableString** | A valid URL format that begins with http:// or https://. Maximum length is 255 characters. The URL that directs to the company&#39;s privacy statement. Not required. | [optional] 

## Methods

### NewMicrosoftGraphPrivacyProfile

`func NewMicrosoftGraphPrivacyProfile() *MicrosoftGraphPrivacyProfile`

NewMicrosoftGraphPrivacyProfile instantiates a new MicrosoftGraphPrivacyProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPrivacyProfileWithDefaults

`func NewMicrosoftGraphPrivacyProfileWithDefaults() *MicrosoftGraphPrivacyProfile`

NewMicrosoftGraphPrivacyProfileWithDefaults instantiates a new MicrosoftGraphPrivacyProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactEmail

`func (o *MicrosoftGraphPrivacyProfile) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *MicrosoftGraphPrivacyProfile) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *MicrosoftGraphPrivacyProfile) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.

### HasContactEmail

`func (o *MicrosoftGraphPrivacyProfile) HasContactEmail() bool`

HasContactEmail returns a boolean if a field has been set.

### SetContactEmailNil

`func (o *MicrosoftGraphPrivacyProfile) SetContactEmailNil(b bool)`

 SetContactEmailNil sets the value for ContactEmail to be an explicit nil

### UnsetContactEmail
`func (o *MicrosoftGraphPrivacyProfile) UnsetContactEmail()`

UnsetContactEmail ensures that no value is present for ContactEmail, not even an explicit nil
### GetStatementUrl

`func (o *MicrosoftGraphPrivacyProfile) GetStatementUrl() string`

GetStatementUrl returns the StatementUrl field if non-nil, zero value otherwise.

### GetStatementUrlOk

`func (o *MicrosoftGraphPrivacyProfile) GetStatementUrlOk() (*string, bool)`

GetStatementUrlOk returns a tuple with the StatementUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementUrl

`func (o *MicrosoftGraphPrivacyProfile) SetStatementUrl(v string)`

SetStatementUrl sets StatementUrl field to given value.

### HasStatementUrl

`func (o *MicrosoftGraphPrivacyProfile) HasStatementUrl() bool`

HasStatementUrl returns a boolean if a field has been set.

### SetStatementUrlNil

`func (o *MicrosoftGraphPrivacyProfile) SetStatementUrlNil(b bool)`

 SetStatementUrlNil sets the value for StatementUrl to be an explicit nil

### UnsetStatementUrl
`func (o *MicrosoftGraphPrivacyProfile) UnsetStatementUrl()`

UnsetStatementUrl ensures that no value is present for StatementUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


