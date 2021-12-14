# IdentityContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionalAccess** | Pointer to [**AnyOfmicrosoftGraphConditionalAccessRoot**](anyOf&lt;microsoft.graph.conditionalAccessRoot&gt;.md) | the entry point for the Conditional Access (CA) object model. | [optional] 
**ApiConnectors** | Pointer to [**[]MicrosoftGraphIdentityApiConnector**](MicrosoftGraphIdentityApiConnector.md) | Represents entry point for API connectors. | [optional] 
**B2xUserFlows** | Pointer to [**[]MicrosoftGraphB2xIdentityUserFlow**](MicrosoftGraphB2xIdentityUserFlow.md) | Represents entry point for B2X/self-service sign-up identity userflows. | [optional] 
**IdentityProviders** | Pointer to [**[]MicrosoftGraphIdentityProviderBase**](MicrosoftGraphIdentityProviderBase.md) | Represents entry point for identity provider base. | [optional] 
**UserFlowAttributes** | Pointer to [**[]MicrosoftGraphIdentityUserFlowAttribute**](MicrosoftGraphIdentityUserFlowAttribute.md) | Represents entry point for identity userflow attributes. | [optional] 

## Methods

### NewIdentityContainer

`func NewIdentityContainer() *IdentityContainer`

NewIdentityContainer instantiates a new IdentityContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityContainerWithDefaults

`func NewIdentityContainerWithDefaults() *IdentityContainer`

NewIdentityContainerWithDefaults instantiates a new IdentityContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionalAccess

`func (o *IdentityContainer) GetConditionalAccess() AnyOfmicrosoftGraphConditionalAccessRoot`

GetConditionalAccess returns the ConditionalAccess field if non-nil, zero value otherwise.

### GetConditionalAccessOk

`func (o *IdentityContainer) GetConditionalAccessOk() (*AnyOfmicrosoftGraphConditionalAccessRoot, bool)`

GetConditionalAccessOk returns a tuple with the ConditionalAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalAccess

`func (o *IdentityContainer) SetConditionalAccess(v AnyOfmicrosoftGraphConditionalAccessRoot)`

SetConditionalAccess sets ConditionalAccess field to given value.

### HasConditionalAccess

`func (o *IdentityContainer) HasConditionalAccess() bool`

HasConditionalAccess returns a boolean if a field has been set.

### SetConditionalAccessNil

`func (o *IdentityContainer) SetConditionalAccessNil(b bool)`

 SetConditionalAccessNil sets the value for ConditionalAccess to be an explicit nil

### UnsetConditionalAccess
`func (o *IdentityContainer) UnsetConditionalAccess()`

UnsetConditionalAccess ensures that no value is present for ConditionalAccess, not even an explicit nil
### GetApiConnectors

`func (o *IdentityContainer) GetApiConnectors() []MicrosoftGraphIdentityApiConnector`

GetApiConnectors returns the ApiConnectors field if non-nil, zero value otherwise.

### GetApiConnectorsOk

`func (o *IdentityContainer) GetApiConnectorsOk() (*[]MicrosoftGraphIdentityApiConnector, bool)`

GetApiConnectorsOk returns a tuple with the ApiConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiConnectors

`func (o *IdentityContainer) SetApiConnectors(v []MicrosoftGraphIdentityApiConnector)`

SetApiConnectors sets ApiConnectors field to given value.

### HasApiConnectors

`func (o *IdentityContainer) HasApiConnectors() bool`

HasApiConnectors returns a boolean if a field has been set.

### GetB2xUserFlows

`func (o *IdentityContainer) GetB2xUserFlows() []MicrosoftGraphB2xIdentityUserFlow`

GetB2xUserFlows returns the B2xUserFlows field if non-nil, zero value otherwise.

### GetB2xUserFlowsOk

`func (o *IdentityContainer) GetB2xUserFlowsOk() (*[]MicrosoftGraphB2xIdentityUserFlow, bool)`

GetB2xUserFlowsOk returns a tuple with the B2xUserFlows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB2xUserFlows

`func (o *IdentityContainer) SetB2xUserFlows(v []MicrosoftGraphB2xIdentityUserFlow)`

SetB2xUserFlows sets B2xUserFlows field to given value.

### HasB2xUserFlows

`func (o *IdentityContainer) HasB2xUserFlows() bool`

HasB2xUserFlows returns a boolean if a field has been set.

### GetIdentityProviders

`func (o *IdentityContainer) GetIdentityProviders() []MicrosoftGraphIdentityProviderBase`

GetIdentityProviders returns the IdentityProviders field if non-nil, zero value otherwise.

### GetIdentityProvidersOk

`func (o *IdentityContainer) GetIdentityProvidersOk() (*[]MicrosoftGraphIdentityProviderBase, bool)`

GetIdentityProvidersOk returns a tuple with the IdentityProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviders

`func (o *IdentityContainer) SetIdentityProviders(v []MicrosoftGraphIdentityProviderBase)`

SetIdentityProviders sets IdentityProviders field to given value.

### HasIdentityProviders

`func (o *IdentityContainer) HasIdentityProviders() bool`

HasIdentityProviders returns a boolean if a field has been set.

### GetUserFlowAttributes

`func (o *IdentityContainer) GetUserFlowAttributes() []MicrosoftGraphIdentityUserFlowAttribute`

GetUserFlowAttributes returns the UserFlowAttributes field if non-nil, zero value otherwise.

### GetUserFlowAttributesOk

`func (o *IdentityContainer) GetUserFlowAttributesOk() (*[]MicrosoftGraphIdentityUserFlowAttribute, bool)`

GetUserFlowAttributesOk returns a tuple with the UserFlowAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFlowAttributes

`func (o *IdentityContainer) SetUserFlowAttributes(v []MicrosoftGraphIdentityUserFlowAttribute)`

SetUserFlowAttributes sets UserFlowAttributes field to given value.

### HasUserFlowAttributes

`func (o *IdentityContainer) HasUserFlowAttributes() bool`

HasUserFlowAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


