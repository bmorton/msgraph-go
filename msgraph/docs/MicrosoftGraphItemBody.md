# MicrosoftGraphItemBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **NullableString** | The content of the item. | [optional] 
**ContentType** | Pointer to [**AnyOfmicrosoftGraphBodyType**](anyOf&lt;microsoft.graph.bodyType&gt;.md) | The type of the content. Possible values are text and html. | [optional] 

## Methods

### NewMicrosoftGraphItemBody

`func NewMicrosoftGraphItemBody() *MicrosoftGraphItemBody`

NewMicrosoftGraphItemBody instantiates a new MicrosoftGraphItemBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphItemBodyWithDefaults

`func NewMicrosoftGraphItemBodyWithDefaults() *MicrosoftGraphItemBody`

NewMicrosoftGraphItemBodyWithDefaults instantiates a new MicrosoftGraphItemBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *MicrosoftGraphItemBody) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MicrosoftGraphItemBody) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MicrosoftGraphItemBody) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *MicrosoftGraphItemBody) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *MicrosoftGraphItemBody) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *MicrosoftGraphItemBody) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetContentType

`func (o *MicrosoftGraphItemBody) GetContentType() AnyOfmicrosoftGraphBodyType`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *MicrosoftGraphItemBody) GetContentTypeOk() (*AnyOfmicrosoftGraphBodyType, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *MicrosoftGraphItemBody) SetContentType(v AnyOfmicrosoftGraphBodyType)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *MicrosoftGraphItemBody) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *MicrosoftGraphItemBody) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *MicrosoftGraphItemBody) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


