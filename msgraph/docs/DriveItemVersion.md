# DriveItemVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **NullableString** | The content stream for this version of the item. | [optional] 
**Size** | Pointer to **NullableInt64** | Indicates the size of the content stream for this version of the item. | [optional] 

## Methods

### NewDriveItemVersion

`func NewDriveItemVersion() *DriveItemVersion`

NewDriveItemVersion instantiates a new DriveItemVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriveItemVersionWithDefaults

`func NewDriveItemVersionWithDefaults() *DriveItemVersion`

NewDriveItemVersionWithDefaults instantiates a new DriveItemVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *DriveItemVersion) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DriveItemVersion) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *DriveItemVersion) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *DriveItemVersion) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *DriveItemVersion) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *DriveItemVersion) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetSize

`func (o *DriveItemVersion) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DriveItemVersion) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DriveItemVersion) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *DriveItemVersion) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *DriveItemVersion) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *DriveItemVersion) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


