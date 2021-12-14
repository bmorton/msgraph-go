# MicrosoftGraphListInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentTypesEnabled** | Pointer to **NullableBool** | If true, indicates that content types are enabled for this list. | [optional] 
**Hidden** | Pointer to **NullableBool** | If true, indicates that the list is not normally visible in the SharePoint user experience. | [optional] 
**Template** | Pointer to **NullableString** | An enumerated value that represents the base list template used in creating the list. Possible values include documentLibrary, genericList, task, survey, announcements, contacts, and more. | [optional] 

## Methods

### NewMicrosoftGraphListInfo

`func NewMicrosoftGraphListInfo() *MicrosoftGraphListInfo`

NewMicrosoftGraphListInfo instantiates a new MicrosoftGraphListInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphListInfoWithDefaults

`func NewMicrosoftGraphListInfoWithDefaults() *MicrosoftGraphListInfo`

NewMicrosoftGraphListInfoWithDefaults instantiates a new MicrosoftGraphListInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentTypesEnabled

`func (o *MicrosoftGraphListInfo) GetContentTypesEnabled() bool`

GetContentTypesEnabled returns the ContentTypesEnabled field if non-nil, zero value otherwise.

### GetContentTypesEnabledOk

`func (o *MicrosoftGraphListInfo) GetContentTypesEnabledOk() (*bool, bool)`

GetContentTypesEnabledOk returns a tuple with the ContentTypesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypesEnabled

`func (o *MicrosoftGraphListInfo) SetContentTypesEnabled(v bool)`

SetContentTypesEnabled sets ContentTypesEnabled field to given value.

### HasContentTypesEnabled

`func (o *MicrosoftGraphListInfo) HasContentTypesEnabled() bool`

HasContentTypesEnabled returns a boolean if a field has been set.

### SetContentTypesEnabledNil

`func (o *MicrosoftGraphListInfo) SetContentTypesEnabledNil(b bool)`

 SetContentTypesEnabledNil sets the value for ContentTypesEnabled to be an explicit nil

### UnsetContentTypesEnabled
`func (o *MicrosoftGraphListInfo) UnsetContentTypesEnabled()`

UnsetContentTypesEnabled ensures that no value is present for ContentTypesEnabled, not even an explicit nil
### GetHidden

`func (o *MicrosoftGraphListInfo) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *MicrosoftGraphListInfo) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *MicrosoftGraphListInfo) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *MicrosoftGraphListInfo) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### SetHiddenNil

`func (o *MicrosoftGraphListInfo) SetHiddenNil(b bool)`

 SetHiddenNil sets the value for Hidden to be an explicit nil

### UnsetHidden
`func (o *MicrosoftGraphListInfo) UnsetHidden()`

UnsetHidden ensures that no value is present for Hidden, not even an explicit nil
### GetTemplate

`func (o *MicrosoftGraphListInfo) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *MicrosoftGraphListInfo) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *MicrosoftGraphListInfo) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *MicrosoftGraphListInfo) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *MicrosoftGraphListInfo) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *MicrosoftGraphListInfo) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


