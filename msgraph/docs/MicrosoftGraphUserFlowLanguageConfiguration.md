# MicrosoftGraphUserFlowLanguageConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DisplayName** | Pointer to **NullableString** | The language name to display. This property is read-only. | [optional] 
**IsEnabled** | Pointer to **bool** | Indicates whether the language is enabled within the user flow. | [optional] 
**DefaultPages** | Pointer to [**[]MicrosoftGraphUserFlowLanguagePage**](MicrosoftGraphUserFlowLanguagePage.md) | Collection of pages with the default content to display in a user flow for a specified language. This collection does not allow any kind of modification. | [optional] 
**OverridesPages** | Pointer to [**[]MicrosoftGraphUserFlowLanguagePage**](MicrosoftGraphUserFlowLanguagePage.md) | Collection of pages with the overrides messages to display in a user flow for a specified language. This collection only allows to modify the content of the page, any other modification is not allowed (creation or deletion of pages). | [optional] 

## Methods

### NewMicrosoftGraphUserFlowLanguageConfiguration

`func NewMicrosoftGraphUserFlowLanguageConfiguration() *MicrosoftGraphUserFlowLanguageConfiguration`

NewMicrosoftGraphUserFlowLanguageConfiguration instantiates a new MicrosoftGraphUserFlowLanguageConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphUserFlowLanguageConfigurationWithDefaults

`func NewMicrosoftGraphUserFlowLanguageConfigurationWithDefaults() *MicrosoftGraphUserFlowLanguageConfiguration`

NewMicrosoftGraphUserFlowLanguageConfigurationWithDefaults instantiates a new MicrosoftGraphUserFlowLanguageConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphUserFlowLanguageConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphUserFlowLanguageConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphUserFlowLanguageConfiguration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphUserFlowLanguageConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *MicrosoftGraphUserFlowLanguageConfiguration) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphUserFlowLanguageConfiguration) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphUserFlowLanguageConfiguration) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphUserFlowLanguageConfiguration) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphUserFlowLanguageConfiguration) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphUserFlowLanguageConfiguration) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetIsEnabled

`func (o *MicrosoftGraphUserFlowLanguageConfiguration) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *MicrosoftGraphUserFlowLanguageConfiguration) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *MicrosoftGraphUserFlowLanguageConfiguration) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *MicrosoftGraphUserFlowLanguageConfiguration) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetDefaultPages

`func (o *MicrosoftGraphUserFlowLanguageConfiguration) GetDefaultPages() []MicrosoftGraphUserFlowLanguagePage`

GetDefaultPages returns the DefaultPages field if non-nil, zero value otherwise.

### GetDefaultPagesOk

`func (o *MicrosoftGraphUserFlowLanguageConfiguration) GetDefaultPagesOk() (*[]MicrosoftGraphUserFlowLanguagePage, bool)`

GetDefaultPagesOk returns a tuple with the DefaultPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPages

`func (o *MicrosoftGraphUserFlowLanguageConfiguration) SetDefaultPages(v []MicrosoftGraphUserFlowLanguagePage)`

SetDefaultPages sets DefaultPages field to given value.

### HasDefaultPages

`func (o *MicrosoftGraphUserFlowLanguageConfiguration) HasDefaultPages() bool`

HasDefaultPages returns a boolean if a field has been set.

### GetOverridesPages

`func (o *MicrosoftGraphUserFlowLanguageConfiguration) GetOverridesPages() []MicrosoftGraphUserFlowLanguagePage`

GetOverridesPages returns the OverridesPages field if non-nil, zero value otherwise.

### GetOverridesPagesOk

`func (o *MicrosoftGraphUserFlowLanguageConfiguration) GetOverridesPagesOk() (*[]MicrosoftGraphUserFlowLanguagePage, bool)`

GetOverridesPagesOk returns a tuple with the OverridesPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridesPages

`func (o *MicrosoftGraphUserFlowLanguageConfiguration) SetOverridesPages(v []MicrosoftGraphUserFlowLanguagePage)`

SetOverridesPages sets OverridesPages field to given value.

### HasOverridesPages

`func (o *MicrosoftGraphUserFlowLanguageConfiguration) HasOverridesPages() bool`

HasOverridesPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


