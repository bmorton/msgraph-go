# MicrosoftGraphIdentityContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ConditionalAccess** | Pointer to [**AnyOfmicrosoftGraphConditionalAccessRoot**](anyOf&lt;microsoft.graph.conditionalAccessRoot&gt;.md) | the entry point for the Conditional Access (CA) object model. | [optional] 
**ApiConnectors** | Pointer to [**[]MicrosoftGraphIdentityApiConnector**](MicrosoftGraphIdentityApiConnector.md) | Represents entry point for API connectors. | [optional] 
**B2xUserFlows** | Pointer to [**[]MicrosoftGraphB2xIdentityUserFlow**](MicrosoftGraphB2xIdentityUserFlow.md) | Represents entry point for B2X/self-service sign-up identity userflows. | [optional] 
**IdentityProviders** | Pointer to [**[]MicrosoftGraphIdentityProviderBase**](MicrosoftGraphIdentityProviderBase.md) | Represents entry point for identity provider base. | [optional] 
**UserFlowAttributes** | Pointer to [**[]MicrosoftGraphIdentityUserFlowAttribute**](MicrosoftGraphIdentityUserFlowAttribute.md) | Represents entry point for identity userflow attributes. | [optional] 

## Methods

### NewMicrosoftGraphIdentityContainer

`func NewMicrosoftGraphIdentityContainer() *MicrosoftGraphIdentityContainer`

NewMicrosoftGraphIdentityContainer instantiates a new MicrosoftGraphIdentityContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphIdentityContainerWithDefaults

`func NewMicrosoftGraphIdentityContainerWithDefaults() *MicrosoftGraphIdentityContainer`

NewMicrosoftGraphIdentityContainerWithDefaults instantiates a new MicrosoftGraphIdentityContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphIdentityContainer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphIdentityContainer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphIdentityContainer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphIdentityContainer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConditionalAccess

`func (o *MicrosoftGraphIdentityContainer) GetConditionalAccess() AnyOfmicrosoftGraphConditionalAccessRoot`

GetConditionalAccess returns the ConditionalAccess field if non-nil, zero value otherwise.

### GetConditionalAccessOk

`func (o *MicrosoftGraphIdentityContainer) GetConditionalAccessOk() (*AnyOfmicrosoftGraphConditionalAccessRoot, bool)`

GetConditionalAccessOk returns a tuple with the ConditionalAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalAccess

`func (o *MicrosoftGraphIdentityContainer) SetConditionalAccess(v AnyOfmicrosoftGraphConditionalAccessRoot)`

SetConditionalAccess sets ConditionalAccess field to given value.

### HasConditionalAccess

`func (o *MicrosoftGraphIdentityContainer) HasConditionalAccess() bool`

HasConditionalAccess returns a boolean if a field has been set.

### SetConditionalAccessNil

`func (o *MicrosoftGraphIdentityContainer) SetConditionalAccessNil(b bool)`

 SetConditionalAccessNil sets the value for ConditionalAccess to be an explicit nil

### UnsetConditionalAccess
`func (o *MicrosoftGraphIdentityContainer) UnsetConditionalAccess()`

UnsetConditionalAccess ensures that no value is present for ConditionalAccess, not even an explicit nil
### GetApiConnectors

`func (o *MicrosoftGraphIdentityContainer) GetApiConnectors() []MicrosoftGraphIdentityApiConnector`

GetApiConnectors returns the ApiConnectors field if non-nil, zero value otherwise.

### GetApiConnectorsOk

`func (o *MicrosoftGraphIdentityContainer) GetApiConnectorsOk() (*[]MicrosoftGraphIdentityApiConnector, bool)`

GetApiConnectorsOk returns a tuple with the ApiConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiConnectors

`func (o *MicrosoftGraphIdentityContainer) SetApiConnectors(v []MicrosoftGraphIdentityApiConnector)`

SetApiConnectors sets ApiConnectors field to given value.

### HasApiConnectors

`func (o *MicrosoftGraphIdentityContainer) HasApiConnectors() bool`

HasApiConnectors returns a boolean if a field has been set.

### GetB2xUserFlows

`func (o *MicrosoftGraphIdentityContainer) GetB2xUserFlows() []MicrosoftGraphB2xIdentityUserFlow`

GetB2xUserFlows returns the B2xUserFlows field if non-nil, zero value otherwise.

### GetB2xUserFlowsOk

`func (o *MicrosoftGraphIdentityContainer) GetB2xUserFlowsOk() (*[]MicrosoftGraphB2xIdentityUserFlow, bool)`

GetB2xUserFlowsOk returns a tuple with the B2xUserFlows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB2xUserFlows

`func (o *MicrosoftGraphIdentityContainer) SetB2xUserFlows(v []MicrosoftGraphB2xIdentityUserFlow)`

SetB2xUserFlows sets B2xUserFlows field to given value.

### HasB2xUserFlows

`func (o *MicrosoftGraphIdentityContainer) HasB2xUserFlows() bool`

HasB2xUserFlows returns a boolean if a field has been set.

### GetIdentityProviders

`func (o *MicrosoftGraphIdentityContainer) GetIdentityProviders() []MicrosoftGraphIdentityProviderBase`

GetIdentityProviders returns the IdentityProviders field if non-nil, zero value otherwise.

### GetIdentityProvidersOk

`func (o *MicrosoftGraphIdentityContainer) GetIdentityProvidersOk() (*[]MicrosoftGraphIdentityProviderBase, bool)`

GetIdentityProvidersOk returns a tuple with the IdentityProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviders

`func (o *MicrosoftGraphIdentityContainer) SetIdentityProviders(v []MicrosoftGraphIdentityProviderBase)`

SetIdentityProviders sets IdentityProviders field to given value.

### HasIdentityProviders

`func (o *MicrosoftGraphIdentityContainer) HasIdentityProviders() bool`

HasIdentityProviders returns a boolean if a field has been set.

### GetUserFlowAttributes

`func (o *MicrosoftGraphIdentityContainer) GetUserFlowAttributes() []MicrosoftGraphIdentityUserFlowAttribute`

GetUserFlowAttributes returns the UserFlowAttributes field if non-nil, zero value otherwise.

### GetUserFlowAttributesOk

`func (o *MicrosoftGraphIdentityContainer) GetUserFlowAttributesOk() (*[]MicrosoftGraphIdentityUserFlowAttribute, bool)`

GetUserFlowAttributesOk returns a tuple with the UserFlowAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFlowAttributes

`func (o *MicrosoftGraphIdentityContainer) SetUserFlowAttributes(v []MicrosoftGraphIdentityUserFlowAttribute)`

SetUserFlowAttributes sets UserFlowAttributes field to given value.

### HasUserFlowAttributes

`func (o *MicrosoftGraphIdentityContainer) HasUserFlowAttributes() bool`

HasUserFlowAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


