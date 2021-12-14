# MicrosoftGraphApplicationTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Categories** | Pointer to **[]string** | The list of categories for the application. Supported values can be: Collaboration, Business Management, Consumer, Content management, CRM, Data services, Developer services, E-commerce, Education, ERP, Finance, Health, Human resources, IT infrastructure, Mail, Management, Marketing, Media, Productivity, Project management, Telecommunications, Tools, Travel, and Web design &amp; hosting. | [optional] 
**Description** | Pointer to **NullableString** | A description of the application. | [optional] 
**DisplayName** | Pointer to **NullableString** | The name of the application. | [optional] 
**HomePageUrl** | Pointer to **NullableString** | The home page URL of the application. | [optional] 
**LogoUrl** | Pointer to **NullableString** | The URL to get the logo for this application. | [optional] 
**Publisher** | Pointer to **NullableString** | The name of the publisher for this application. | [optional] 
**SupportedProvisioningTypes** | Pointer to **[]string** | The list of provisioning modes supported by this application. The only valid value is sync. | [optional] 
**SupportedSingleSignOnModes** | Pointer to **[]string** | The list of single sign-on modes supported by this application. The supported values are oidc, password, saml, and notSupported. | [optional] 

## Methods

### NewMicrosoftGraphApplicationTemplate

`func NewMicrosoftGraphApplicationTemplate() *MicrosoftGraphApplicationTemplate`

NewMicrosoftGraphApplicationTemplate instantiates a new MicrosoftGraphApplicationTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphApplicationTemplateWithDefaults

`func NewMicrosoftGraphApplicationTemplateWithDefaults() *MicrosoftGraphApplicationTemplate`

NewMicrosoftGraphApplicationTemplateWithDefaults instantiates a new MicrosoftGraphApplicationTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphApplicationTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphApplicationTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphApplicationTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphApplicationTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCategories

`func (o *MicrosoftGraphApplicationTemplate) GetCategories() []*string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *MicrosoftGraphApplicationTemplate) GetCategoriesOk() (*[]*string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *MicrosoftGraphApplicationTemplate) SetCategories(v []*string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *MicrosoftGraphApplicationTemplate) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphApplicationTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphApplicationTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphApplicationTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphApplicationTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphApplicationTemplate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphApplicationTemplate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphApplicationTemplate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphApplicationTemplate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphApplicationTemplate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphApplicationTemplate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphApplicationTemplate) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphApplicationTemplate) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetHomePageUrl

`func (o *MicrosoftGraphApplicationTemplate) GetHomePageUrl() string`

GetHomePageUrl returns the HomePageUrl field if non-nil, zero value otherwise.

### GetHomePageUrlOk

`func (o *MicrosoftGraphApplicationTemplate) GetHomePageUrlOk() (*string, bool)`

GetHomePageUrlOk returns a tuple with the HomePageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePageUrl

`func (o *MicrosoftGraphApplicationTemplate) SetHomePageUrl(v string)`

SetHomePageUrl sets HomePageUrl field to given value.

### HasHomePageUrl

`func (o *MicrosoftGraphApplicationTemplate) HasHomePageUrl() bool`

HasHomePageUrl returns a boolean if a field has been set.

### SetHomePageUrlNil

`func (o *MicrosoftGraphApplicationTemplate) SetHomePageUrlNil(b bool)`

 SetHomePageUrlNil sets the value for HomePageUrl to be an explicit nil

### UnsetHomePageUrl
`func (o *MicrosoftGraphApplicationTemplate) UnsetHomePageUrl()`

UnsetHomePageUrl ensures that no value is present for HomePageUrl, not even an explicit nil
### GetLogoUrl

`func (o *MicrosoftGraphApplicationTemplate) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *MicrosoftGraphApplicationTemplate) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *MicrosoftGraphApplicationTemplate) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *MicrosoftGraphApplicationTemplate) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### SetLogoUrlNil

`func (o *MicrosoftGraphApplicationTemplate) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *MicrosoftGraphApplicationTemplate) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetPublisher

`func (o *MicrosoftGraphApplicationTemplate) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *MicrosoftGraphApplicationTemplate) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *MicrosoftGraphApplicationTemplate) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *MicrosoftGraphApplicationTemplate) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### SetPublisherNil

`func (o *MicrosoftGraphApplicationTemplate) SetPublisherNil(b bool)`

 SetPublisherNil sets the value for Publisher to be an explicit nil

### UnsetPublisher
`func (o *MicrosoftGraphApplicationTemplate) UnsetPublisher()`

UnsetPublisher ensures that no value is present for Publisher, not even an explicit nil
### GetSupportedProvisioningTypes

`func (o *MicrosoftGraphApplicationTemplate) GetSupportedProvisioningTypes() []*string`

GetSupportedProvisioningTypes returns the SupportedProvisioningTypes field if non-nil, zero value otherwise.

### GetSupportedProvisioningTypesOk

`func (o *MicrosoftGraphApplicationTemplate) GetSupportedProvisioningTypesOk() (*[]*string, bool)`

GetSupportedProvisioningTypesOk returns a tuple with the SupportedProvisioningTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedProvisioningTypes

`func (o *MicrosoftGraphApplicationTemplate) SetSupportedProvisioningTypes(v []*string)`

SetSupportedProvisioningTypes sets SupportedProvisioningTypes field to given value.

### HasSupportedProvisioningTypes

`func (o *MicrosoftGraphApplicationTemplate) HasSupportedProvisioningTypes() bool`

HasSupportedProvisioningTypes returns a boolean if a field has been set.

### GetSupportedSingleSignOnModes

`func (o *MicrosoftGraphApplicationTemplate) GetSupportedSingleSignOnModes() []*string`

GetSupportedSingleSignOnModes returns the SupportedSingleSignOnModes field if non-nil, zero value otherwise.

### GetSupportedSingleSignOnModesOk

`func (o *MicrosoftGraphApplicationTemplate) GetSupportedSingleSignOnModesOk() (*[]*string, bool)`

GetSupportedSingleSignOnModesOk returns a tuple with the SupportedSingleSignOnModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedSingleSignOnModes

`func (o *MicrosoftGraphApplicationTemplate) SetSupportedSingleSignOnModes(v []*string)`

SetSupportedSingleSignOnModes sets SupportedSingleSignOnModes field to given value.

### HasSupportedSingleSignOnModes

`func (o *MicrosoftGraphApplicationTemplate) HasSupportedSingleSignOnModes() bool`

HasSupportedSingleSignOnModes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


