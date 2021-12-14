# ExternalConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to [**AnyOfmicrosoftGraphExternalConnectorsConfiguration**](anyOf&lt;microsoft.graph.externalConnectors.configuration&gt;.md) | Specifies additional application IDs that are allowed to manage the connection and to index content in the connection. Optional. | [optional] 
**Description** | Pointer to **NullableString** | Description of the connection displayed in the Microsoft 365 admin center. Optional. | [optional] 
**Name** | Pointer to **NullableString** | The display name of the connection to be displayed in the Microsoft 365 admin center. Maximum length of 128 characters. Required. | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphExternalConnectorsConnectionState**](anyOf&lt;microsoft.graph.externalConnectors.connectionState&gt;.md) | Indicates the current state of the connection. Possible values are: draft, ready, obsolete, limitExceeded, unknownFutureValue. | [optional] 
**Groups** | Pointer to [**[]MicrosoftGraphExternalConnectorsExternalGroup**](MicrosoftGraphExternalConnectorsExternalGroup.md) | Read-only. Nullable. | [optional] 
**Items** | Pointer to [**[]MicrosoftGraphExternalConnectorsExternalItem**](MicrosoftGraphExternalConnectorsExternalItem.md) | Read-only. Nullable. | [optional] 
**Operations** | Pointer to [**[]MicrosoftGraphExternalConnectorsConnectionOperation**](MicrosoftGraphExternalConnectorsConnectionOperation.md) | Read-only. Nullable. | [optional] 
**Schema** | Pointer to [**AnyOfmicrosoftGraphExternalConnectorsSchema**](anyOf&lt;microsoft.graph.externalConnectors.schema&gt;.md) | Read-only. Nullable. | [optional] 

## Methods

### NewExternalConnection

`func NewExternalConnection() *ExternalConnection`

NewExternalConnection instantiates a new ExternalConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalConnectionWithDefaults

`func NewExternalConnectionWithDefaults() *ExternalConnection`

NewExternalConnectionWithDefaults instantiates a new ExternalConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *ExternalConnection) GetConfiguration() AnyOfmicrosoftGraphExternalConnectorsConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *ExternalConnection) GetConfigurationOk() (*AnyOfmicrosoftGraphExternalConnectorsConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *ExternalConnection) SetConfiguration(v AnyOfmicrosoftGraphExternalConnectorsConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *ExternalConnection) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *ExternalConnection) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *ExternalConnection) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetDescription

`func (o *ExternalConnection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExternalConnection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExternalConnection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExternalConnection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ExternalConnection) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ExternalConnection) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *ExternalConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalConnection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExternalConnection) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ExternalConnection) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ExternalConnection) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetState

`func (o *ExternalConnection) GetState() AnyOfmicrosoftGraphExternalConnectorsConnectionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ExternalConnection) GetStateOk() (*AnyOfmicrosoftGraphExternalConnectorsConnectionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ExternalConnection) SetState(v AnyOfmicrosoftGraphExternalConnectorsConnectionState)`

SetState sets State field to given value.

### HasState

`func (o *ExternalConnection) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *ExternalConnection) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *ExternalConnection) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetGroups

`func (o *ExternalConnection) GetGroups() []MicrosoftGraphExternalConnectorsExternalGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ExternalConnection) GetGroupsOk() (*[]MicrosoftGraphExternalConnectorsExternalGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ExternalConnection) SetGroups(v []MicrosoftGraphExternalConnectorsExternalGroup)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ExternalConnection) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetItems

`func (o *ExternalConnection) GetItems() []MicrosoftGraphExternalConnectorsExternalItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ExternalConnection) GetItemsOk() (*[]MicrosoftGraphExternalConnectorsExternalItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ExternalConnection) SetItems(v []MicrosoftGraphExternalConnectorsExternalItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *ExternalConnection) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOperations

`func (o *ExternalConnection) GetOperations() []MicrosoftGraphExternalConnectorsConnectionOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *ExternalConnection) GetOperationsOk() (*[]MicrosoftGraphExternalConnectorsConnectionOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *ExternalConnection) SetOperations(v []MicrosoftGraphExternalConnectorsConnectionOperation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *ExternalConnection) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetSchema

`func (o *ExternalConnection) GetSchema() AnyOfmicrosoftGraphExternalConnectorsSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *ExternalConnection) GetSchemaOk() (*AnyOfmicrosoftGraphExternalConnectorsSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *ExternalConnection) SetSchema(v AnyOfmicrosoftGraphExternalConnectorsSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *ExternalConnection) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *ExternalConnection) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *ExternalConnection) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


