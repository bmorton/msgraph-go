# ProfilePhoto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Height** | Pointer to **NullableInt32** | The height of the photo. Read-only. | [optional] 
**Width** | Pointer to **NullableInt32** | The width of the photo. Read-only. | [optional] 

## Methods

### NewProfilePhoto

`func NewProfilePhoto() *ProfilePhoto`

NewProfilePhoto instantiates a new ProfilePhoto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfilePhotoWithDefaults

`func NewProfilePhotoWithDefaults() *ProfilePhoto`

NewProfilePhotoWithDefaults instantiates a new ProfilePhoto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeight

`func (o *ProfilePhoto) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *ProfilePhoto) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *ProfilePhoto) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *ProfilePhoto) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeightNil

`func (o *ProfilePhoto) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *ProfilePhoto) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetWidth

`func (o *ProfilePhoto) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *ProfilePhoto) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *ProfilePhoto) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *ProfilePhoto) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidthNil

`func (o *ProfilePhoto) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *ProfilePhoto) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


