# GroupSettingTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Description of the template. | [optional] 
**DisplayName** | Pointer to **NullableString** | Display name of the template. | [optional] 
**Values** | Pointer to [**[]MicrosoftGraphSettingTemplateValue**](MicrosoftGraphSettingTemplateValue.md) | Collection of settingTemplateValues that list the set of available settings, defaults and types that make up this template. | [optional] 

## Methods

### NewGroupSettingTemplate

`func NewGroupSettingTemplate() *GroupSettingTemplate`

NewGroupSettingTemplate instantiates a new GroupSettingTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupSettingTemplateWithDefaults

`func NewGroupSettingTemplateWithDefaults() *GroupSettingTemplate`

NewGroupSettingTemplateWithDefaults instantiates a new GroupSettingTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GroupSettingTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupSettingTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupSettingTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupSettingTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GroupSettingTemplate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GroupSettingTemplate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *GroupSettingTemplate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GroupSettingTemplate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GroupSettingTemplate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GroupSettingTemplate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *GroupSettingTemplate) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *GroupSettingTemplate) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetValues

`func (o *GroupSettingTemplate) GetValues() []MicrosoftGraphSettingTemplateValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GroupSettingTemplate) GetValuesOk() (*[]MicrosoftGraphSettingTemplateValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GroupSettingTemplate) SetValues(v []MicrosoftGraphSettingTemplateValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *GroupSettingTemplate) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


