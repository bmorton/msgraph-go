# UserFlowLanguageConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | The language name to display. This property is read-only. | [optional] 
**IsEnabled** | Pointer to **bool** | Indicates whether the language is enabled within the user flow. | [optional] 
**DefaultPages** | Pointer to [**[]MicrosoftGraphUserFlowLanguagePage**](MicrosoftGraphUserFlowLanguagePage.md) | Collection of pages with the default content to display in a user flow for a specified language. This collection does not allow any kind of modification. | [optional] 
**OverridesPages** | Pointer to [**[]MicrosoftGraphUserFlowLanguagePage**](MicrosoftGraphUserFlowLanguagePage.md) | Collection of pages with the overrides messages to display in a user flow for a specified language. This collection only allows to modify the content of the page, any other modification is not allowed (creation or deletion of pages). | [optional] 

## Methods

### NewUserFlowLanguageConfiguration

`func NewUserFlowLanguageConfiguration() *UserFlowLanguageConfiguration`

NewUserFlowLanguageConfiguration instantiates a new UserFlowLanguageConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFlowLanguageConfigurationWithDefaults

`func NewUserFlowLanguageConfigurationWithDefaults() *UserFlowLanguageConfiguration`

NewUserFlowLanguageConfigurationWithDefaults instantiates a new UserFlowLanguageConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *UserFlowLanguageConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserFlowLanguageConfiguration) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserFlowLanguageConfiguration) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserFlowLanguageConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *UserFlowLanguageConfiguration) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *UserFlowLanguageConfiguration) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetIsEnabled

`func (o *UserFlowLanguageConfiguration) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *UserFlowLanguageConfiguration) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *UserFlowLanguageConfiguration) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *UserFlowLanguageConfiguration) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetDefaultPages

`func (o *UserFlowLanguageConfiguration) GetDefaultPages() []MicrosoftGraphUserFlowLanguagePage`

GetDefaultPages returns the DefaultPages field if non-nil, zero value otherwise.

### GetDefaultPagesOk

`func (o *UserFlowLanguageConfiguration) GetDefaultPagesOk() (*[]MicrosoftGraphUserFlowLanguagePage, bool)`

GetDefaultPagesOk returns a tuple with the DefaultPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPages

`func (o *UserFlowLanguageConfiguration) SetDefaultPages(v []MicrosoftGraphUserFlowLanguagePage)`

SetDefaultPages sets DefaultPages field to given value.

### HasDefaultPages

`func (o *UserFlowLanguageConfiguration) HasDefaultPages() bool`

HasDefaultPages returns a boolean if a field has been set.

### GetOverridesPages

`func (o *UserFlowLanguageConfiguration) GetOverridesPages() []MicrosoftGraphUserFlowLanguagePage`

GetOverridesPages returns the OverridesPages field if non-nil, zero value otherwise.

### GetOverridesPagesOk

`func (o *UserFlowLanguageConfiguration) GetOverridesPagesOk() (*[]MicrosoftGraphUserFlowLanguagePage, bool)`

GetOverridesPagesOk returns a tuple with the OverridesPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridesPages

`func (o *UserFlowLanguageConfiguration) SetOverridesPages(v []MicrosoftGraphUserFlowLanguagePage)`

SetOverridesPages sets OverridesPages field to given value.

### HasOverridesPages

`func (o *UserFlowLanguageConfiguration) HasOverridesPages() bool`

HasOverridesPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


