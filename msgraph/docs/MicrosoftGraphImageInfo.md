# MicrosoftGraphImageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddImageQuery** | Pointer to **NullableBool** | Optional; parameter used to indicate the server is able to render image dynamically in response to parameterization. For example – a high contrast image | [optional] 
**AlternateText** | Pointer to **NullableString** | Optional; alt-text accessible content for the image | [optional] 
**AlternativeText** | Pointer to **NullableString** |  | [optional] 
**IconUrl** | Pointer to **NullableString** | Optional; URI that points to an icon which represents the application used to generate the activity | [optional] 

## Methods

### NewMicrosoftGraphImageInfo

`func NewMicrosoftGraphImageInfo() *MicrosoftGraphImageInfo`

NewMicrosoftGraphImageInfo instantiates a new MicrosoftGraphImageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphImageInfoWithDefaults

`func NewMicrosoftGraphImageInfoWithDefaults() *MicrosoftGraphImageInfo`

NewMicrosoftGraphImageInfoWithDefaults instantiates a new MicrosoftGraphImageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddImageQuery

`func (o *MicrosoftGraphImageInfo) GetAddImageQuery() bool`

GetAddImageQuery returns the AddImageQuery field if non-nil, zero value otherwise.

### GetAddImageQueryOk

`func (o *MicrosoftGraphImageInfo) GetAddImageQueryOk() (*bool, bool)`

GetAddImageQueryOk returns a tuple with the AddImageQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddImageQuery

`func (o *MicrosoftGraphImageInfo) SetAddImageQuery(v bool)`

SetAddImageQuery sets AddImageQuery field to given value.

### HasAddImageQuery

`func (o *MicrosoftGraphImageInfo) HasAddImageQuery() bool`

HasAddImageQuery returns a boolean if a field has been set.

### SetAddImageQueryNil

`func (o *MicrosoftGraphImageInfo) SetAddImageQueryNil(b bool)`

 SetAddImageQueryNil sets the value for AddImageQuery to be an explicit nil

### UnsetAddImageQuery
`func (o *MicrosoftGraphImageInfo) UnsetAddImageQuery()`

UnsetAddImageQuery ensures that no value is present for AddImageQuery, not even an explicit nil
### GetAlternateText

`func (o *MicrosoftGraphImageInfo) GetAlternateText() string`

GetAlternateText returns the AlternateText field if non-nil, zero value otherwise.

### GetAlternateTextOk

`func (o *MicrosoftGraphImageInfo) GetAlternateTextOk() (*string, bool)`

GetAlternateTextOk returns a tuple with the AlternateText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateText

`func (o *MicrosoftGraphImageInfo) SetAlternateText(v string)`

SetAlternateText sets AlternateText field to given value.

### HasAlternateText

`func (o *MicrosoftGraphImageInfo) HasAlternateText() bool`

HasAlternateText returns a boolean if a field has been set.

### SetAlternateTextNil

`func (o *MicrosoftGraphImageInfo) SetAlternateTextNil(b bool)`

 SetAlternateTextNil sets the value for AlternateText to be an explicit nil

### UnsetAlternateText
`func (o *MicrosoftGraphImageInfo) UnsetAlternateText()`

UnsetAlternateText ensures that no value is present for AlternateText, not even an explicit nil
### GetAlternativeText

`func (o *MicrosoftGraphImageInfo) GetAlternativeText() string`

GetAlternativeText returns the AlternativeText field if non-nil, zero value otherwise.

### GetAlternativeTextOk

`func (o *MicrosoftGraphImageInfo) GetAlternativeTextOk() (*string, bool)`

GetAlternativeTextOk returns a tuple with the AlternativeText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeText

`func (o *MicrosoftGraphImageInfo) SetAlternativeText(v string)`

SetAlternativeText sets AlternativeText field to given value.

### HasAlternativeText

`func (o *MicrosoftGraphImageInfo) HasAlternativeText() bool`

HasAlternativeText returns a boolean if a field has been set.

### SetAlternativeTextNil

`func (o *MicrosoftGraphImageInfo) SetAlternativeTextNil(b bool)`

 SetAlternativeTextNil sets the value for AlternativeText to be an explicit nil

### UnsetAlternativeText
`func (o *MicrosoftGraphImageInfo) UnsetAlternativeText()`

UnsetAlternativeText ensures that no value is present for AlternativeText, not even an explicit nil
### GetIconUrl

`func (o *MicrosoftGraphImageInfo) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *MicrosoftGraphImageInfo) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *MicrosoftGraphImageInfo) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *MicrosoftGraphImageInfo) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *MicrosoftGraphImageInfo) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *MicrosoftGraphImageInfo) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


