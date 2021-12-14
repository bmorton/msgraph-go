# MicrosoftGraphIntuneBrand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactITEmailAddress** | Pointer to **NullableString** | Email address of the person/organization responsible for IT support. | [optional] 
**ContactITName** | Pointer to **NullableString** | Name of the person/organization responsible for IT support. | [optional] 
**ContactITNotes** | Pointer to **NullableString** | Text comments regarding the person/organization responsible for IT support. | [optional] 
**ContactITPhoneNumber** | Pointer to **NullableString** | Phone number of the person/organization responsible for IT support. | [optional] 
**DarkBackgroundLogo** | Pointer to [**AnyOfmicrosoftGraphMimeContent**](anyOf&lt;microsoft.graph.mimeContent&gt;.md) | Logo image displayed in Company Portal apps which have a dark background behind the logo. | [optional] 
**DisplayName** | Pointer to **NullableString** | Company/organization name that is displayed to end users. | [optional] 
**LightBackgroundLogo** | Pointer to [**AnyOfmicrosoftGraphMimeContent**](anyOf&lt;microsoft.graph.mimeContent&gt;.md) | Logo image displayed in Company Portal apps which have a light background behind the logo. | [optional] 
**OnlineSupportSiteName** | Pointer to **NullableString** | Display name of the company/organization’s IT helpdesk site. | [optional] 
**OnlineSupportSiteUrl** | Pointer to **NullableString** | URL to the company/organization’s IT helpdesk site. | [optional] 
**PrivacyUrl** | Pointer to **NullableString** | URL to the company/organization’s privacy policy. | [optional] 
**ShowDisplayNameNextToLogo** | Pointer to **bool** | Boolean that represents whether the administrator-supplied display name will be shown next to the logo image. | [optional] 
**ShowLogo** | Pointer to **bool** | Boolean that represents whether the administrator-supplied logo images are shown or not shown. | [optional] 
**ShowNameNextToLogo** | Pointer to **bool** | Boolean that represents whether the administrator-supplied display name will be shown next to the logo image. | [optional] 
**ThemeColor** | Pointer to [**AnyOfmicrosoftGraphRgbColor**](anyOf&lt;microsoft.graph.rgbColor&gt;.md) | Primary theme color used in the Company Portal applications and web portal. | [optional] 

## Methods

### NewMicrosoftGraphIntuneBrand

`func NewMicrosoftGraphIntuneBrand() *MicrosoftGraphIntuneBrand`

NewMicrosoftGraphIntuneBrand instantiates a new MicrosoftGraphIntuneBrand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphIntuneBrandWithDefaults

`func NewMicrosoftGraphIntuneBrandWithDefaults() *MicrosoftGraphIntuneBrand`

NewMicrosoftGraphIntuneBrandWithDefaults instantiates a new MicrosoftGraphIntuneBrand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactITEmailAddress

`func (o *MicrosoftGraphIntuneBrand) GetContactITEmailAddress() string`

GetContactITEmailAddress returns the ContactITEmailAddress field if non-nil, zero value otherwise.

### GetContactITEmailAddressOk

`func (o *MicrosoftGraphIntuneBrand) GetContactITEmailAddressOk() (*string, bool)`

GetContactITEmailAddressOk returns a tuple with the ContactITEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactITEmailAddress

`func (o *MicrosoftGraphIntuneBrand) SetContactITEmailAddress(v string)`

SetContactITEmailAddress sets ContactITEmailAddress field to given value.

### HasContactITEmailAddress

`func (o *MicrosoftGraphIntuneBrand) HasContactITEmailAddress() bool`

HasContactITEmailAddress returns a boolean if a field has been set.

### SetContactITEmailAddressNil

`func (o *MicrosoftGraphIntuneBrand) SetContactITEmailAddressNil(b bool)`

 SetContactITEmailAddressNil sets the value for ContactITEmailAddress to be an explicit nil

### UnsetContactITEmailAddress
`func (o *MicrosoftGraphIntuneBrand) UnsetContactITEmailAddress()`

UnsetContactITEmailAddress ensures that no value is present for ContactITEmailAddress, not even an explicit nil
### GetContactITName

`func (o *MicrosoftGraphIntuneBrand) GetContactITName() string`

GetContactITName returns the ContactITName field if non-nil, zero value otherwise.

### GetContactITNameOk

`func (o *MicrosoftGraphIntuneBrand) GetContactITNameOk() (*string, bool)`

GetContactITNameOk returns a tuple with the ContactITName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactITName

`func (o *MicrosoftGraphIntuneBrand) SetContactITName(v string)`

SetContactITName sets ContactITName field to given value.

### HasContactITName

`func (o *MicrosoftGraphIntuneBrand) HasContactITName() bool`

HasContactITName returns a boolean if a field has been set.

### SetContactITNameNil

`func (o *MicrosoftGraphIntuneBrand) SetContactITNameNil(b bool)`

 SetContactITNameNil sets the value for ContactITName to be an explicit nil

### UnsetContactITName
`func (o *MicrosoftGraphIntuneBrand) UnsetContactITName()`

UnsetContactITName ensures that no value is present for ContactITName, not even an explicit nil
### GetContactITNotes

`func (o *MicrosoftGraphIntuneBrand) GetContactITNotes() string`

GetContactITNotes returns the ContactITNotes field if non-nil, zero value otherwise.

### GetContactITNotesOk

`func (o *MicrosoftGraphIntuneBrand) GetContactITNotesOk() (*string, bool)`

GetContactITNotesOk returns a tuple with the ContactITNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactITNotes

`func (o *MicrosoftGraphIntuneBrand) SetContactITNotes(v string)`

SetContactITNotes sets ContactITNotes field to given value.

### HasContactITNotes

`func (o *MicrosoftGraphIntuneBrand) HasContactITNotes() bool`

HasContactITNotes returns a boolean if a field has been set.

### SetContactITNotesNil

`func (o *MicrosoftGraphIntuneBrand) SetContactITNotesNil(b bool)`

 SetContactITNotesNil sets the value for ContactITNotes to be an explicit nil

### UnsetContactITNotes
`func (o *MicrosoftGraphIntuneBrand) UnsetContactITNotes()`

UnsetContactITNotes ensures that no value is present for ContactITNotes, not even an explicit nil
### GetContactITPhoneNumber

`func (o *MicrosoftGraphIntuneBrand) GetContactITPhoneNumber() string`

GetContactITPhoneNumber returns the ContactITPhoneNumber field if non-nil, zero value otherwise.

### GetContactITPhoneNumberOk

`func (o *MicrosoftGraphIntuneBrand) GetContactITPhoneNumberOk() (*string, bool)`

GetContactITPhoneNumberOk returns a tuple with the ContactITPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactITPhoneNumber

`func (o *MicrosoftGraphIntuneBrand) SetContactITPhoneNumber(v string)`

SetContactITPhoneNumber sets ContactITPhoneNumber field to given value.

### HasContactITPhoneNumber

`func (o *MicrosoftGraphIntuneBrand) HasContactITPhoneNumber() bool`

HasContactITPhoneNumber returns a boolean if a field has been set.

### SetContactITPhoneNumberNil

`func (o *MicrosoftGraphIntuneBrand) SetContactITPhoneNumberNil(b bool)`

 SetContactITPhoneNumberNil sets the value for ContactITPhoneNumber to be an explicit nil

### UnsetContactITPhoneNumber
`func (o *MicrosoftGraphIntuneBrand) UnsetContactITPhoneNumber()`

UnsetContactITPhoneNumber ensures that no value is present for ContactITPhoneNumber, not even an explicit nil
### GetDarkBackgroundLogo

`func (o *MicrosoftGraphIntuneBrand) GetDarkBackgroundLogo() AnyOfmicrosoftGraphMimeContent`

GetDarkBackgroundLogo returns the DarkBackgroundLogo field if non-nil, zero value otherwise.

### GetDarkBackgroundLogoOk

`func (o *MicrosoftGraphIntuneBrand) GetDarkBackgroundLogoOk() (*AnyOfmicrosoftGraphMimeContent, bool)`

GetDarkBackgroundLogoOk returns a tuple with the DarkBackgroundLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkBackgroundLogo

`func (o *MicrosoftGraphIntuneBrand) SetDarkBackgroundLogo(v AnyOfmicrosoftGraphMimeContent)`

SetDarkBackgroundLogo sets DarkBackgroundLogo field to given value.

### HasDarkBackgroundLogo

`func (o *MicrosoftGraphIntuneBrand) HasDarkBackgroundLogo() bool`

HasDarkBackgroundLogo returns a boolean if a field has been set.

### SetDarkBackgroundLogoNil

`func (o *MicrosoftGraphIntuneBrand) SetDarkBackgroundLogoNil(b bool)`

 SetDarkBackgroundLogoNil sets the value for DarkBackgroundLogo to be an explicit nil

### UnsetDarkBackgroundLogo
`func (o *MicrosoftGraphIntuneBrand) UnsetDarkBackgroundLogo()`

UnsetDarkBackgroundLogo ensures that no value is present for DarkBackgroundLogo, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphIntuneBrand) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphIntuneBrand) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphIntuneBrand) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphIntuneBrand) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphIntuneBrand) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphIntuneBrand) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetLightBackgroundLogo

`func (o *MicrosoftGraphIntuneBrand) GetLightBackgroundLogo() AnyOfmicrosoftGraphMimeContent`

GetLightBackgroundLogo returns the LightBackgroundLogo field if non-nil, zero value otherwise.

### GetLightBackgroundLogoOk

`func (o *MicrosoftGraphIntuneBrand) GetLightBackgroundLogoOk() (*AnyOfmicrosoftGraphMimeContent, bool)`

GetLightBackgroundLogoOk returns a tuple with the LightBackgroundLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLightBackgroundLogo

`func (o *MicrosoftGraphIntuneBrand) SetLightBackgroundLogo(v AnyOfmicrosoftGraphMimeContent)`

SetLightBackgroundLogo sets LightBackgroundLogo field to given value.

### HasLightBackgroundLogo

`func (o *MicrosoftGraphIntuneBrand) HasLightBackgroundLogo() bool`

HasLightBackgroundLogo returns a boolean if a field has been set.

### SetLightBackgroundLogoNil

`func (o *MicrosoftGraphIntuneBrand) SetLightBackgroundLogoNil(b bool)`

 SetLightBackgroundLogoNil sets the value for LightBackgroundLogo to be an explicit nil

### UnsetLightBackgroundLogo
`func (o *MicrosoftGraphIntuneBrand) UnsetLightBackgroundLogo()`

UnsetLightBackgroundLogo ensures that no value is present for LightBackgroundLogo, not even an explicit nil
### GetOnlineSupportSiteName

`func (o *MicrosoftGraphIntuneBrand) GetOnlineSupportSiteName() string`

GetOnlineSupportSiteName returns the OnlineSupportSiteName field if non-nil, zero value otherwise.

### GetOnlineSupportSiteNameOk

`func (o *MicrosoftGraphIntuneBrand) GetOnlineSupportSiteNameOk() (*string, bool)`

GetOnlineSupportSiteNameOk returns a tuple with the OnlineSupportSiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineSupportSiteName

`func (o *MicrosoftGraphIntuneBrand) SetOnlineSupportSiteName(v string)`

SetOnlineSupportSiteName sets OnlineSupportSiteName field to given value.

### HasOnlineSupportSiteName

`func (o *MicrosoftGraphIntuneBrand) HasOnlineSupportSiteName() bool`

HasOnlineSupportSiteName returns a boolean if a field has been set.

### SetOnlineSupportSiteNameNil

`func (o *MicrosoftGraphIntuneBrand) SetOnlineSupportSiteNameNil(b bool)`

 SetOnlineSupportSiteNameNil sets the value for OnlineSupportSiteName to be an explicit nil

### UnsetOnlineSupportSiteName
`func (o *MicrosoftGraphIntuneBrand) UnsetOnlineSupportSiteName()`

UnsetOnlineSupportSiteName ensures that no value is present for OnlineSupportSiteName, not even an explicit nil
### GetOnlineSupportSiteUrl

`func (o *MicrosoftGraphIntuneBrand) GetOnlineSupportSiteUrl() string`

GetOnlineSupportSiteUrl returns the OnlineSupportSiteUrl field if non-nil, zero value otherwise.

### GetOnlineSupportSiteUrlOk

`func (o *MicrosoftGraphIntuneBrand) GetOnlineSupportSiteUrlOk() (*string, bool)`

GetOnlineSupportSiteUrlOk returns a tuple with the OnlineSupportSiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineSupportSiteUrl

`func (o *MicrosoftGraphIntuneBrand) SetOnlineSupportSiteUrl(v string)`

SetOnlineSupportSiteUrl sets OnlineSupportSiteUrl field to given value.

### HasOnlineSupportSiteUrl

`func (o *MicrosoftGraphIntuneBrand) HasOnlineSupportSiteUrl() bool`

HasOnlineSupportSiteUrl returns a boolean if a field has been set.

### SetOnlineSupportSiteUrlNil

`func (o *MicrosoftGraphIntuneBrand) SetOnlineSupportSiteUrlNil(b bool)`

 SetOnlineSupportSiteUrlNil sets the value for OnlineSupportSiteUrl to be an explicit nil

### UnsetOnlineSupportSiteUrl
`func (o *MicrosoftGraphIntuneBrand) UnsetOnlineSupportSiteUrl()`

UnsetOnlineSupportSiteUrl ensures that no value is present for OnlineSupportSiteUrl, not even an explicit nil
### GetPrivacyUrl

`func (o *MicrosoftGraphIntuneBrand) GetPrivacyUrl() string`

GetPrivacyUrl returns the PrivacyUrl field if non-nil, zero value otherwise.

### GetPrivacyUrlOk

`func (o *MicrosoftGraphIntuneBrand) GetPrivacyUrlOk() (*string, bool)`

GetPrivacyUrlOk returns a tuple with the PrivacyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyUrl

`func (o *MicrosoftGraphIntuneBrand) SetPrivacyUrl(v string)`

SetPrivacyUrl sets PrivacyUrl field to given value.

### HasPrivacyUrl

`func (o *MicrosoftGraphIntuneBrand) HasPrivacyUrl() bool`

HasPrivacyUrl returns a boolean if a field has been set.

### SetPrivacyUrlNil

`func (o *MicrosoftGraphIntuneBrand) SetPrivacyUrlNil(b bool)`

 SetPrivacyUrlNil sets the value for PrivacyUrl to be an explicit nil

### UnsetPrivacyUrl
`func (o *MicrosoftGraphIntuneBrand) UnsetPrivacyUrl()`

UnsetPrivacyUrl ensures that no value is present for PrivacyUrl, not even an explicit nil
### GetShowDisplayNameNextToLogo

`func (o *MicrosoftGraphIntuneBrand) GetShowDisplayNameNextToLogo() bool`

GetShowDisplayNameNextToLogo returns the ShowDisplayNameNextToLogo field if non-nil, zero value otherwise.

### GetShowDisplayNameNextToLogoOk

`func (o *MicrosoftGraphIntuneBrand) GetShowDisplayNameNextToLogoOk() (*bool, bool)`

GetShowDisplayNameNextToLogoOk returns a tuple with the ShowDisplayNameNextToLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDisplayNameNextToLogo

`func (o *MicrosoftGraphIntuneBrand) SetShowDisplayNameNextToLogo(v bool)`

SetShowDisplayNameNextToLogo sets ShowDisplayNameNextToLogo field to given value.

### HasShowDisplayNameNextToLogo

`func (o *MicrosoftGraphIntuneBrand) HasShowDisplayNameNextToLogo() bool`

HasShowDisplayNameNextToLogo returns a boolean if a field has been set.

### GetShowLogo

`func (o *MicrosoftGraphIntuneBrand) GetShowLogo() bool`

GetShowLogo returns the ShowLogo field if non-nil, zero value otherwise.

### GetShowLogoOk

`func (o *MicrosoftGraphIntuneBrand) GetShowLogoOk() (*bool, bool)`

GetShowLogoOk returns a tuple with the ShowLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLogo

`func (o *MicrosoftGraphIntuneBrand) SetShowLogo(v bool)`

SetShowLogo sets ShowLogo field to given value.

### HasShowLogo

`func (o *MicrosoftGraphIntuneBrand) HasShowLogo() bool`

HasShowLogo returns a boolean if a field has been set.

### GetShowNameNextToLogo

`func (o *MicrosoftGraphIntuneBrand) GetShowNameNextToLogo() bool`

GetShowNameNextToLogo returns the ShowNameNextToLogo field if non-nil, zero value otherwise.

### GetShowNameNextToLogoOk

`func (o *MicrosoftGraphIntuneBrand) GetShowNameNextToLogoOk() (*bool, bool)`

GetShowNameNextToLogoOk returns a tuple with the ShowNameNextToLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNameNextToLogo

`func (o *MicrosoftGraphIntuneBrand) SetShowNameNextToLogo(v bool)`

SetShowNameNextToLogo sets ShowNameNextToLogo field to given value.

### HasShowNameNextToLogo

`func (o *MicrosoftGraphIntuneBrand) HasShowNameNextToLogo() bool`

HasShowNameNextToLogo returns a boolean if a field has been set.

### GetThemeColor

`func (o *MicrosoftGraphIntuneBrand) GetThemeColor() AnyOfmicrosoftGraphRgbColor`

GetThemeColor returns the ThemeColor field if non-nil, zero value otherwise.

### GetThemeColorOk

`func (o *MicrosoftGraphIntuneBrand) GetThemeColorOk() (*AnyOfmicrosoftGraphRgbColor, bool)`

GetThemeColorOk returns a tuple with the ThemeColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeColor

`func (o *MicrosoftGraphIntuneBrand) SetThemeColor(v AnyOfmicrosoftGraphRgbColor)`

SetThemeColor sets ThemeColor field to given value.

### HasThemeColor

`func (o *MicrosoftGraphIntuneBrand) HasThemeColor() bool`

HasThemeColor returns a boolean if a field has been set.

### SetThemeColorNil

`func (o *MicrosoftGraphIntuneBrand) SetThemeColorNil(b bool)`

 SetThemeColorNil sets the value for ThemeColor to be an explicit nil

### UnsetThemeColor
`func (o *MicrosoftGraphIntuneBrand) UnsetThemeColor()`

UnsetThemeColor ensures that no value is present for ThemeColor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


