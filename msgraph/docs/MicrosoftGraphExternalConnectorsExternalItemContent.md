# MicrosoftGraphExternalConnectorsExternalItemContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**AnyOfmicrosoftGraphExternalConnectorsExternalItemContentType**](anyOf&lt;microsoft.graph.externalConnectors.externalItemContentType&gt;.md) | The type of content in the value property. Possible values are: text, html, unknownFutureValue. | [optional] 
**Value** | Pointer to **NullableString** | The content for the externalItem. Required. | [optional] 

## Methods

### NewMicrosoftGraphExternalConnectorsExternalItemContent

`func NewMicrosoftGraphExternalConnectorsExternalItemContent() *MicrosoftGraphExternalConnectorsExternalItemContent`

NewMicrosoftGraphExternalConnectorsExternalItemContent instantiates a new MicrosoftGraphExternalConnectorsExternalItemContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphExternalConnectorsExternalItemContentWithDefaults

`func NewMicrosoftGraphExternalConnectorsExternalItemContentWithDefaults() *MicrosoftGraphExternalConnectorsExternalItemContent`

NewMicrosoftGraphExternalConnectorsExternalItemContentWithDefaults instantiates a new MicrosoftGraphExternalConnectorsExternalItemContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MicrosoftGraphExternalConnectorsExternalItemContent) GetType() AnyOfmicrosoftGraphExternalConnectorsExternalItemContentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphExternalConnectorsExternalItemContent) GetTypeOk() (*AnyOfmicrosoftGraphExternalConnectorsExternalItemContentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphExternalConnectorsExternalItemContent) SetType(v AnyOfmicrosoftGraphExternalConnectorsExternalItemContentType)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphExternalConnectorsExternalItemContent) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphExternalConnectorsExternalItemContent) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphExternalConnectorsExternalItemContent) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetValue

`func (o *MicrosoftGraphExternalConnectorsExternalItemContent) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MicrosoftGraphExternalConnectorsExternalItemContent) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MicrosoftGraphExternalConnectorsExternalItemContent) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MicrosoftGraphExternalConnectorsExternalItemContent) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *MicrosoftGraphExternalConnectorsExternalItemContent) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *MicrosoftGraphExternalConnectorsExternalItemContent) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


