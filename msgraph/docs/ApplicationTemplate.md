# ApplicationTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to **[]string** | The list of categories for the application. Supported values can be: Collaboration, Business Management, Consumer, Content management, CRM, Data services, Developer services, E-commerce, Education, ERP, Finance, Health, Human resources, IT infrastructure, Mail, Management, Marketing, Media, Productivity, Project management, Telecommunications, Tools, Travel, and Web design &amp; hosting. | [optional] 
**Description** | Pointer to **NullableString** | A description of the application. | [optional] 
**DisplayName** | Pointer to **NullableString** | The name of the application. | [optional] 
**HomePageUrl** | Pointer to **NullableString** | The home page URL of the application. | [optional] 
**LogoUrl** | Pointer to **NullableString** | The URL to get the logo for this application. | [optional] 
**Publisher** | Pointer to **NullableString** | The name of the publisher for this application. | [optional] 
**SupportedProvisioningTypes** | Pointer to **[]string** | The list of provisioning modes supported by this application. The only valid value is sync. | [optional] 
**SupportedSingleSignOnModes** | Pointer to **[]string** | The list of single sign-on modes supported by this application. The supported values are oidc, password, saml, and notSupported. | [optional] 

## Methods

### NewApplicationTemplate

`func NewApplicationTemplate() *ApplicationTemplate`

NewApplicationTemplate instantiates a new ApplicationTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationTemplateWithDefaults

`func NewApplicationTemplateWithDefaults() *ApplicationTemplate`

NewApplicationTemplateWithDefaults instantiates a new ApplicationTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *ApplicationTemplate) GetCategories() []*string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ApplicationTemplate) GetCategoriesOk() (*[]*string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ApplicationTemplate) SetCategories(v []*string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ApplicationTemplate) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApplicationTemplate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApplicationTemplate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *ApplicationTemplate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApplicationTemplate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApplicationTemplate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ApplicationTemplate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *ApplicationTemplate) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ApplicationTemplate) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetHomePageUrl

`func (o *ApplicationTemplate) GetHomePageUrl() string`

GetHomePageUrl returns the HomePageUrl field if non-nil, zero value otherwise.

### GetHomePageUrlOk

`func (o *ApplicationTemplate) GetHomePageUrlOk() (*string, bool)`

GetHomePageUrlOk returns a tuple with the HomePageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePageUrl

`func (o *ApplicationTemplate) SetHomePageUrl(v string)`

SetHomePageUrl sets HomePageUrl field to given value.

### HasHomePageUrl

`func (o *ApplicationTemplate) HasHomePageUrl() bool`

HasHomePageUrl returns a boolean if a field has been set.

### SetHomePageUrlNil

`func (o *ApplicationTemplate) SetHomePageUrlNil(b bool)`

 SetHomePageUrlNil sets the value for HomePageUrl to be an explicit nil

### UnsetHomePageUrl
`func (o *ApplicationTemplate) UnsetHomePageUrl()`

UnsetHomePageUrl ensures that no value is present for HomePageUrl, not even an explicit nil
### GetLogoUrl

`func (o *ApplicationTemplate) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *ApplicationTemplate) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *ApplicationTemplate) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *ApplicationTemplate) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### SetLogoUrlNil

`func (o *ApplicationTemplate) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *ApplicationTemplate) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetPublisher

`func (o *ApplicationTemplate) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *ApplicationTemplate) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *ApplicationTemplate) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *ApplicationTemplate) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### SetPublisherNil

`func (o *ApplicationTemplate) SetPublisherNil(b bool)`

 SetPublisherNil sets the value for Publisher to be an explicit nil

### UnsetPublisher
`func (o *ApplicationTemplate) UnsetPublisher()`

UnsetPublisher ensures that no value is present for Publisher, not even an explicit nil
### GetSupportedProvisioningTypes

`func (o *ApplicationTemplate) GetSupportedProvisioningTypes() []*string`

GetSupportedProvisioningTypes returns the SupportedProvisioningTypes field if non-nil, zero value otherwise.

### GetSupportedProvisioningTypesOk

`func (o *ApplicationTemplate) GetSupportedProvisioningTypesOk() (*[]*string, bool)`

GetSupportedProvisioningTypesOk returns a tuple with the SupportedProvisioningTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedProvisioningTypes

`func (o *ApplicationTemplate) SetSupportedProvisioningTypes(v []*string)`

SetSupportedProvisioningTypes sets SupportedProvisioningTypes field to given value.

### HasSupportedProvisioningTypes

`func (o *ApplicationTemplate) HasSupportedProvisioningTypes() bool`

HasSupportedProvisioningTypes returns a boolean if a field has been set.

### GetSupportedSingleSignOnModes

`func (o *ApplicationTemplate) GetSupportedSingleSignOnModes() []*string`

GetSupportedSingleSignOnModes returns the SupportedSingleSignOnModes field if non-nil, zero value otherwise.

### GetSupportedSingleSignOnModesOk

`func (o *ApplicationTemplate) GetSupportedSingleSignOnModesOk() (*[]*string, bool)`

GetSupportedSingleSignOnModesOk returns a tuple with the SupportedSingleSignOnModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedSingleSignOnModes

`func (o *ApplicationTemplate) SetSupportedSingleSignOnModes(v []*string)`

SetSupportedSingleSignOnModes sets SupportedSingleSignOnModes field to given value.

### HasSupportedSingleSignOnModes

`func (o *ApplicationTemplate) HasSupportedSingleSignOnModes() bool`

HasSupportedSingleSignOnModes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


