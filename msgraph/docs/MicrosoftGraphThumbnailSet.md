# MicrosoftGraphThumbnailSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Large** | Pointer to [**AnyOfmicrosoftGraphThumbnail**](anyOf&lt;microsoft.graph.thumbnail&gt;.md) | A 1920x1920 scaled thumbnail. | [optional] 
**Medium** | Pointer to [**AnyOfmicrosoftGraphThumbnail**](anyOf&lt;microsoft.graph.thumbnail&gt;.md) | A 176x176 scaled thumbnail. | [optional] 
**Small** | Pointer to [**AnyOfmicrosoftGraphThumbnail**](anyOf&lt;microsoft.graph.thumbnail&gt;.md) | A 48x48 cropped thumbnail. | [optional] 
**Source** | Pointer to [**AnyOfmicrosoftGraphThumbnail**](anyOf&lt;microsoft.graph.thumbnail&gt;.md) | A custom thumbnail image or the original image used to generate other thumbnails. | [optional] 

## Methods

### NewMicrosoftGraphThumbnailSet

`func NewMicrosoftGraphThumbnailSet() *MicrosoftGraphThumbnailSet`

NewMicrosoftGraphThumbnailSet instantiates a new MicrosoftGraphThumbnailSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphThumbnailSetWithDefaults

`func NewMicrosoftGraphThumbnailSetWithDefaults() *MicrosoftGraphThumbnailSet`

NewMicrosoftGraphThumbnailSetWithDefaults instantiates a new MicrosoftGraphThumbnailSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphThumbnailSet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphThumbnailSet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphThumbnailSet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphThumbnailSet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLarge

`func (o *MicrosoftGraphThumbnailSet) GetLarge() AnyOfmicrosoftGraphThumbnail`

GetLarge returns the Large field if non-nil, zero value otherwise.

### GetLargeOk

`func (o *MicrosoftGraphThumbnailSet) GetLargeOk() (*AnyOfmicrosoftGraphThumbnail, bool)`

GetLargeOk returns a tuple with the Large field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLarge

`func (o *MicrosoftGraphThumbnailSet) SetLarge(v AnyOfmicrosoftGraphThumbnail)`

SetLarge sets Large field to given value.

### HasLarge

`func (o *MicrosoftGraphThumbnailSet) HasLarge() bool`

HasLarge returns a boolean if a field has been set.

### SetLargeNil

`func (o *MicrosoftGraphThumbnailSet) SetLargeNil(b bool)`

 SetLargeNil sets the value for Large to be an explicit nil

### UnsetLarge
`func (o *MicrosoftGraphThumbnailSet) UnsetLarge()`

UnsetLarge ensures that no value is present for Large, not even an explicit nil
### GetMedium

`func (o *MicrosoftGraphThumbnailSet) GetMedium() AnyOfmicrosoftGraphThumbnail`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *MicrosoftGraphThumbnailSet) GetMediumOk() (*AnyOfmicrosoftGraphThumbnail, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *MicrosoftGraphThumbnailSet) SetMedium(v AnyOfmicrosoftGraphThumbnail)`

SetMedium sets Medium field to given value.

### HasMedium

`func (o *MicrosoftGraphThumbnailSet) HasMedium() bool`

HasMedium returns a boolean if a field has been set.

### SetMediumNil

`func (o *MicrosoftGraphThumbnailSet) SetMediumNil(b bool)`

 SetMediumNil sets the value for Medium to be an explicit nil

### UnsetMedium
`func (o *MicrosoftGraphThumbnailSet) UnsetMedium()`

UnsetMedium ensures that no value is present for Medium, not even an explicit nil
### GetSmall

`func (o *MicrosoftGraphThumbnailSet) GetSmall() AnyOfmicrosoftGraphThumbnail`

GetSmall returns the Small field if non-nil, zero value otherwise.

### GetSmallOk

`func (o *MicrosoftGraphThumbnailSet) GetSmallOk() (*AnyOfmicrosoftGraphThumbnail, bool)`

GetSmallOk returns a tuple with the Small field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmall

`func (o *MicrosoftGraphThumbnailSet) SetSmall(v AnyOfmicrosoftGraphThumbnail)`

SetSmall sets Small field to given value.

### HasSmall

`func (o *MicrosoftGraphThumbnailSet) HasSmall() bool`

HasSmall returns a boolean if a field has been set.

### SetSmallNil

`func (o *MicrosoftGraphThumbnailSet) SetSmallNil(b bool)`

 SetSmallNil sets the value for Small to be an explicit nil

### UnsetSmall
`func (o *MicrosoftGraphThumbnailSet) UnsetSmall()`

UnsetSmall ensures that no value is present for Small, not even an explicit nil
### GetSource

`func (o *MicrosoftGraphThumbnailSet) GetSource() AnyOfmicrosoftGraphThumbnail`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MicrosoftGraphThumbnailSet) GetSourceOk() (*AnyOfmicrosoftGraphThumbnail, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MicrosoftGraphThumbnailSet) SetSource(v AnyOfmicrosoftGraphThumbnail)`

SetSource sets Source field to given value.

### HasSource

`func (o *MicrosoftGraphThumbnailSet) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *MicrosoftGraphThumbnailSet) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *MicrosoftGraphThumbnailSet) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


