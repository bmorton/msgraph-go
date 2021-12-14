# MicrosoftGraphExternalConnectorsSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**BaseType** | Pointer to **string** | Must be set to microsoft.graph.externalConnector.externalItem. Required. | [optional] 
**Properties** | Pointer to [**[]AnyOfmicrosoftGraphExternalConnectorsProperty**](AnyOfmicrosoftGraphExternalConnectorsProperty.md) | The properties defined for the items in the connection. The minimum number of properties is one, the maximum is 128. | [optional] 

## Methods

### NewMicrosoftGraphExternalConnectorsSchema

`func NewMicrosoftGraphExternalConnectorsSchema() *MicrosoftGraphExternalConnectorsSchema`

NewMicrosoftGraphExternalConnectorsSchema instantiates a new MicrosoftGraphExternalConnectorsSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphExternalConnectorsSchemaWithDefaults

`func NewMicrosoftGraphExternalConnectorsSchemaWithDefaults() *MicrosoftGraphExternalConnectorsSchema`

NewMicrosoftGraphExternalConnectorsSchemaWithDefaults instantiates a new MicrosoftGraphExternalConnectorsSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphExternalConnectorsSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphExternalConnectorsSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphExternalConnectorsSchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphExternalConnectorsSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBaseType

`func (o *MicrosoftGraphExternalConnectorsSchema) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *MicrosoftGraphExternalConnectorsSchema) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *MicrosoftGraphExternalConnectorsSchema) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *MicrosoftGraphExternalConnectorsSchema) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetProperties

`func (o *MicrosoftGraphExternalConnectorsSchema) GetProperties() []*AnyOfmicrosoftGraphExternalConnectorsProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *MicrosoftGraphExternalConnectorsSchema) GetPropertiesOk() (*[]*AnyOfmicrosoftGraphExternalConnectorsProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *MicrosoftGraphExternalConnectorsSchema) SetProperties(v []*AnyOfmicrosoftGraphExternalConnectorsProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *MicrosoftGraphExternalConnectorsSchema) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


