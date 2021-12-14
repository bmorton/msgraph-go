# GroupSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | Display name of this group of settings, which comes from the associated template. | [optional] 
**TemplateId** | Pointer to **NullableString** | Unique identifier for the template used to create this group of settings. Read-only. | [optional] 
**Values** | Pointer to [**[]MicrosoftGraphSettingValue**](MicrosoftGraphSettingValue.md) | Collection of name value pairs. Must contain and set all the settings defined in the template. | [optional] 

## Methods

### NewGroupSetting

`func NewGroupSetting() *GroupSetting`

NewGroupSetting instantiates a new GroupSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupSettingWithDefaults

`func NewGroupSettingWithDefaults() *GroupSetting`

NewGroupSettingWithDefaults instantiates a new GroupSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *GroupSetting) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GroupSetting) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GroupSetting) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GroupSetting) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *GroupSetting) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *GroupSetting) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetTemplateId

`func (o *GroupSetting) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *GroupSetting) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *GroupSetting) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *GroupSetting) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *GroupSetting) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *GroupSetting) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetValues

`func (o *GroupSetting) GetValues() []MicrosoftGraphSettingValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GroupSetting) GetValuesOk() (*[]MicrosoftGraphSettingValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GroupSetting) SetValues(v []MicrosoftGraphSettingValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *GroupSetting) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


