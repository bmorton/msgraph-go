# MicrosoftGraphPrintDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ContentType** | Pointer to **NullableString** | The document&#39;s content (MIME) type. Read-only. | [optional] 
**DisplayName** | Pointer to **NullableString** | The document&#39;s name. Read-only. | [optional] 
**Size** | Pointer to **int64** | The document&#39;s size in bytes. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphPrintDocument

`func NewMicrosoftGraphPrintDocument() *MicrosoftGraphPrintDocument`

NewMicrosoftGraphPrintDocument instantiates a new MicrosoftGraphPrintDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPrintDocumentWithDefaults

`func NewMicrosoftGraphPrintDocumentWithDefaults() *MicrosoftGraphPrintDocument`

NewMicrosoftGraphPrintDocumentWithDefaults instantiates a new MicrosoftGraphPrintDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphPrintDocument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPrintDocument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphPrintDocument) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphPrintDocument) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContentType

`func (o *MicrosoftGraphPrintDocument) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *MicrosoftGraphPrintDocument) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *MicrosoftGraphPrintDocument) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *MicrosoftGraphPrintDocument) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *MicrosoftGraphPrintDocument) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *MicrosoftGraphPrintDocument) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphPrintDocument) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphPrintDocument) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphPrintDocument) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphPrintDocument) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphPrintDocument) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphPrintDocument) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetSize

`func (o *MicrosoftGraphPrintDocument) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MicrosoftGraphPrintDocument) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MicrosoftGraphPrintDocument) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *MicrosoftGraphPrintDocument) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


