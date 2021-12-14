# MicrosoftGraphBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Album** | Pointer to [**AnyOfmicrosoftGraphAlbum**](anyOf&lt;microsoft.graph.album&gt;.md) | If the bundle is an [album][], then the album property is included | [optional] 
**ChildCount** | Pointer to **NullableInt32** | Number of children contained immediately within this container. | [optional] 

## Methods

### NewMicrosoftGraphBundle

`func NewMicrosoftGraphBundle() *MicrosoftGraphBundle`

NewMicrosoftGraphBundle instantiates a new MicrosoftGraphBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphBundleWithDefaults

`func NewMicrosoftGraphBundleWithDefaults() *MicrosoftGraphBundle`

NewMicrosoftGraphBundleWithDefaults instantiates a new MicrosoftGraphBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlbum

`func (o *MicrosoftGraphBundle) GetAlbum() AnyOfmicrosoftGraphAlbum`

GetAlbum returns the Album field if non-nil, zero value otherwise.

### GetAlbumOk

`func (o *MicrosoftGraphBundle) GetAlbumOk() (*AnyOfmicrosoftGraphAlbum, bool)`

GetAlbumOk returns a tuple with the Album field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbum

`func (o *MicrosoftGraphBundle) SetAlbum(v AnyOfmicrosoftGraphAlbum)`

SetAlbum sets Album field to given value.

### HasAlbum

`func (o *MicrosoftGraphBundle) HasAlbum() bool`

HasAlbum returns a boolean if a field has been set.

### SetAlbumNil

`func (o *MicrosoftGraphBundle) SetAlbumNil(b bool)`

 SetAlbumNil sets the value for Album to be an explicit nil

### UnsetAlbum
`func (o *MicrosoftGraphBundle) UnsetAlbum()`

UnsetAlbum ensures that no value is present for Album, not even an explicit nil
### GetChildCount

`func (o *MicrosoftGraphBundle) GetChildCount() int32`

GetChildCount returns the ChildCount field if non-nil, zero value otherwise.

### GetChildCountOk

`func (o *MicrosoftGraphBundle) GetChildCountOk() (*int32, bool)`

GetChildCountOk returns a tuple with the ChildCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildCount

`func (o *MicrosoftGraphBundle) SetChildCount(v int32)`

SetChildCount sets ChildCount field to given value.

### HasChildCount

`func (o *MicrosoftGraphBundle) HasChildCount() bool`

HasChildCount returns a boolean if a field has been set.

### SetChildCountNil

`func (o *MicrosoftGraphBundle) SetChildCountNil(b bool)`

 SetChildCountNil sets the value for ChildCount to be an explicit nil

### UnsetChildCount
`func (o *MicrosoftGraphBundle) UnsetChildCount()`

UnsetChildCount ensures that no value is present for ChildCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


