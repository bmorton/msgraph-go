# MicrosoftGraphParentalControlSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountriesBlockedForMinors** | Pointer to **[]string** | Specifies the two-letter ISO country codes. Access to the application will be blocked for minors from the countries specified in this list. | [optional] 
**LegalAgeGroupRule** | Pointer to **NullableString** | Specifies the legal age group rule that applies to users of the app. Can be set to one of the following values: ValueDescriptionAllowDefault. Enforces the legal minimum. This means parental consent is required for minors in the European Union and Korea.RequireConsentForPrivacyServicesEnforces the user to specify date of birth to comply with COPPA rules. RequireConsentForMinorsRequires parental consent for ages below 18, regardless of country minor rules.RequireConsentForKidsRequires parental consent for ages below 14, regardless of country minor rules.BlockMinorsBlocks minors from using the app. | [optional] 

## Methods

### NewMicrosoftGraphParentalControlSettings

`func NewMicrosoftGraphParentalControlSettings() *MicrosoftGraphParentalControlSettings`

NewMicrosoftGraphParentalControlSettings instantiates a new MicrosoftGraphParentalControlSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphParentalControlSettingsWithDefaults

`func NewMicrosoftGraphParentalControlSettingsWithDefaults() *MicrosoftGraphParentalControlSettings`

NewMicrosoftGraphParentalControlSettingsWithDefaults instantiates a new MicrosoftGraphParentalControlSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountriesBlockedForMinors

`func (o *MicrosoftGraphParentalControlSettings) GetCountriesBlockedForMinors() []*string`

GetCountriesBlockedForMinors returns the CountriesBlockedForMinors field if non-nil, zero value otherwise.

### GetCountriesBlockedForMinorsOk

`func (o *MicrosoftGraphParentalControlSettings) GetCountriesBlockedForMinorsOk() (*[]*string, bool)`

GetCountriesBlockedForMinorsOk returns a tuple with the CountriesBlockedForMinors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountriesBlockedForMinors

`func (o *MicrosoftGraphParentalControlSettings) SetCountriesBlockedForMinors(v []*string)`

SetCountriesBlockedForMinors sets CountriesBlockedForMinors field to given value.

### HasCountriesBlockedForMinors

`func (o *MicrosoftGraphParentalControlSettings) HasCountriesBlockedForMinors() bool`

HasCountriesBlockedForMinors returns a boolean if a field has been set.

### GetLegalAgeGroupRule

`func (o *MicrosoftGraphParentalControlSettings) GetLegalAgeGroupRule() string`

GetLegalAgeGroupRule returns the LegalAgeGroupRule field if non-nil, zero value otherwise.

### GetLegalAgeGroupRuleOk

`func (o *MicrosoftGraphParentalControlSettings) GetLegalAgeGroupRuleOk() (*string, bool)`

GetLegalAgeGroupRuleOk returns a tuple with the LegalAgeGroupRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalAgeGroupRule

`func (o *MicrosoftGraphParentalControlSettings) SetLegalAgeGroupRule(v string)`

SetLegalAgeGroupRule sets LegalAgeGroupRule field to given value.

### HasLegalAgeGroupRule

`func (o *MicrosoftGraphParentalControlSettings) HasLegalAgeGroupRule() bool`

HasLegalAgeGroupRule returns a boolean if a field has been set.

### SetLegalAgeGroupRuleNil

`func (o *MicrosoftGraphParentalControlSettings) SetLegalAgeGroupRuleNil(b bool)`

 SetLegalAgeGroupRuleNil sets the value for LegalAgeGroupRule to be an explicit nil

### UnsetLegalAgeGroupRule
`func (o *MicrosoftGraphParentalControlSettings) UnsetLegalAgeGroupRule()`

UnsetLegalAgeGroupRule ensures that no value is present for LegalAgeGroupRule, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


