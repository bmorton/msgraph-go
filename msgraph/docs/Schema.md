# Schema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseType** | Pointer to **string** | Must be set to microsoft.graph.externalConnector.externalItem. Required. | [optional] 
**Properties** | Pointer to [**[]AnyOfmicrosoftGraphExternalConnectorsProperty**](AnyOfmicrosoftGraphExternalConnectorsProperty.md) | The properties defined for the items in the connection. The minimum number of properties is one, the maximum is 128. | [optional] 

## Methods

### NewSchema

`func NewSchema() *Schema`

NewSchema instantiates a new Schema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaWithDefaults

`func NewSchemaWithDefaults() *Schema`

NewSchemaWithDefaults instantiates a new Schema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseType

`func (o *Schema) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *Schema) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *Schema) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *Schema) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetProperties

`func (o *Schema) GetProperties() []*AnyOfmicrosoftGraphExternalConnectorsProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Schema) GetPropertiesOk() (*[]*AnyOfmicrosoftGraphExternalConnectorsProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Schema) SetProperties(v []*AnyOfmicrosoftGraphExternalConnectorsProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Schema) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


