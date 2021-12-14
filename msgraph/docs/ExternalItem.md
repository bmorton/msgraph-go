# ExternalItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acl** | Pointer to [**[]AnyOfmicrosoftGraphExternalConnectorsAcl**](AnyOfmicrosoftGraphExternalConnectorsAcl.md) | An array of access control entries. Each entry specifies the access granted to a user or group. Required. | [optional] 
**Content** | Pointer to [**AnyOfmicrosoftGraphExternalConnectorsExternalItemContent**](anyOf&lt;microsoft.graph.externalConnectors.externalItemContent&gt;.md) | A plain-text  representation of the contents of the item. The text in this property is full-text indexed. Optional. | [optional] 
**Properties** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | A property bag with the properties of the item. The properties MUST conform to the schema defined for the externalConnection. Required. | [optional] 

## Methods

### NewExternalItem

`func NewExternalItem() *ExternalItem`

NewExternalItem instantiates a new ExternalItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalItemWithDefaults

`func NewExternalItemWithDefaults() *ExternalItem`

NewExternalItemWithDefaults instantiates a new ExternalItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcl

`func (o *ExternalItem) GetAcl() []*AnyOfmicrosoftGraphExternalConnectorsAcl`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *ExternalItem) GetAclOk() (*[]*AnyOfmicrosoftGraphExternalConnectorsAcl, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *ExternalItem) SetAcl(v []*AnyOfmicrosoftGraphExternalConnectorsAcl)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *ExternalItem) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetContent

`func (o *ExternalItem) GetContent() AnyOfmicrosoftGraphExternalConnectorsExternalItemContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ExternalItem) GetContentOk() (*AnyOfmicrosoftGraphExternalConnectorsExternalItemContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ExternalItem) SetContent(v AnyOfmicrosoftGraphExternalConnectorsExternalItemContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *ExternalItem) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *ExternalItem) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *ExternalItem) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetProperties

`func (o *ExternalItem) GetProperties() AnyOfobject`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ExternalItem) GetPropertiesOk() (*AnyOfobject, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ExternalItem) SetProperties(v AnyOfobject)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ExternalItem) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *ExternalItem) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *ExternalItem) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


