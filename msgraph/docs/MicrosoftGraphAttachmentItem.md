# MicrosoftGraphAttachmentItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachmentType** | Pointer to [**AnyOfmicrosoftGraphAttachmentType**](anyOf&lt;microsoft.graph.attachmentType&gt;.md) | The type of attachment. Possible values are: file, item, reference. Required. | [optional] 
**ContentType** | Pointer to **NullableString** | The nature of the data in the attachment. Optional. | [optional] 
**IsInline** | Pointer to **NullableBool** | true if the attachment is an inline attachment; otherwise, false. Optional. | [optional] 
**Name** | Pointer to **NullableString** | The display name of the attachment. This can be a descriptive string and does not have to be the actual file name. Required. | [optional] 
**Size** | Pointer to **NullableInt64** | The length of the attachment in bytes. Required. | [optional] 

## Methods

### NewMicrosoftGraphAttachmentItem

`func NewMicrosoftGraphAttachmentItem() *MicrosoftGraphAttachmentItem`

NewMicrosoftGraphAttachmentItem instantiates a new MicrosoftGraphAttachmentItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAttachmentItemWithDefaults

`func NewMicrosoftGraphAttachmentItemWithDefaults() *MicrosoftGraphAttachmentItem`

NewMicrosoftGraphAttachmentItemWithDefaults instantiates a new MicrosoftGraphAttachmentItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachmentType

`func (o *MicrosoftGraphAttachmentItem) GetAttachmentType() AnyOfmicrosoftGraphAttachmentType`

GetAttachmentType returns the AttachmentType field if non-nil, zero value otherwise.

### GetAttachmentTypeOk

`func (o *MicrosoftGraphAttachmentItem) GetAttachmentTypeOk() (*AnyOfmicrosoftGraphAttachmentType, bool)`

GetAttachmentTypeOk returns a tuple with the AttachmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentType

`func (o *MicrosoftGraphAttachmentItem) SetAttachmentType(v AnyOfmicrosoftGraphAttachmentType)`

SetAttachmentType sets AttachmentType field to given value.

### HasAttachmentType

`func (o *MicrosoftGraphAttachmentItem) HasAttachmentType() bool`

HasAttachmentType returns a boolean if a field has been set.

### SetAttachmentTypeNil

`func (o *MicrosoftGraphAttachmentItem) SetAttachmentTypeNil(b bool)`

 SetAttachmentTypeNil sets the value for AttachmentType to be an explicit nil

### UnsetAttachmentType
`func (o *MicrosoftGraphAttachmentItem) UnsetAttachmentType()`

UnsetAttachmentType ensures that no value is present for AttachmentType, not even an explicit nil
### GetContentType

`func (o *MicrosoftGraphAttachmentItem) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *MicrosoftGraphAttachmentItem) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *MicrosoftGraphAttachmentItem) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *MicrosoftGraphAttachmentItem) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *MicrosoftGraphAttachmentItem) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *MicrosoftGraphAttachmentItem) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetIsInline

`func (o *MicrosoftGraphAttachmentItem) GetIsInline() bool`

GetIsInline returns the IsInline field if non-nil, zero value otherwise.

### GetIsInlineOk

`func (o *MicrosoftGraphAttachmentItem) GetIsInlineOk() (*bool, bool)`

GetIsInlineOk returns a tuple with the IsInline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInline

`func (o *MicrosoftGraphAttachmentItem) SetIsInline(v bool)`

SetIsInline sets IsInline field to given value.

### HasIsInline

`func (o *MicrosoftGraphAttachmentItem) HasIsInline() bool`

HasIsInline returns a boolean if a field has been set.

### SetIsInlineNil

`func (o *MicrosoftGraphAttachmentItem) SetIsInlineNil(b bool)`

 SetIsInlineNil sets the value for IsInline to be an explicit nil

### UnsetIsInline
`func (o *MicrosoftGraphAttachmentItem) UnsetIsInline()`

UnsetIsInline ensures that no value is present for IsInline, not even an explicit nil
### GetName

`func (o *MicrosoftGraphAttachmentItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphAttachmentItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphAttachmentItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphAttachmentItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphAttachmentItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphAttachmentItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSize

`func (o *MicrosoftGraphAttachmentItem) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MicrosoftGraphAttachmentItem) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MicrosoftGraphAttachmentItem) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *MicrosoftGraphAttachmentItem) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *MicrosoftGraphAttachmentItem) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *MicrosoftGraphAttachmentItem) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


