# B2xIdentityUserFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiConnectorConfiguration** | Pointer to [**AnyOfmicrosoftGraphUserFlowApiConnectorConfiguration**](anyOf&lt;microsoft.graph.userFlowApiConnectorConfiguration&gt;.md) | Configuration for enabling an API connector for use as part of the self-service sign-up user flow. You can only obtain the value of this object using Get userFlowApiConnectorConfiguration. | [optional] 
**IdentityProviders** | Pointer to [**[]MicrosoftGraphIdentityProvider**](MicrosoftGraphIdentityProvider.md) | The identity providers included in the user flow. | [optional] 
**Languages** | Pointer to [**[]MicrosoftGraphUserFlowLanguageConfiguration**](MicrosoftGraphUserFlowLanguageConfiguration.md) | The languages supported for customization within the user flow. Language customization is enabled by default in self-service sign-up user flow. You cannot create custom languages in self-service sign-up user flows. | [optional] 
**UserAttributeAssignments** | Pointer to [**[]MicrosoftGraphIdentityUserFlowAttributeAssignment**](MicrosoftGraphIdentityUserFlowAttributeAssignment.md) | The user attribute assignments included in the user flow. | [optional] 
**UserFlowIdentityProviders** | Pointer to [**[]MicrosoftGraphIdentityProviderBase**](MicrosoftGraphIdentityProviderBase.md) |  | [optional] 

## Methods

### NewB2xIdentityUserFlow

`func NewB2xIdentityUserFlow() *B2xIdentityUserFlow`

NewB2xIdentityUserFlow instantiates a new B2xIdentityUserFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewB2xIdentityUserFlowWithDefaults

`func NewB2xIdentityUserFlowWithDefaults() *B2xIdentityUserFlow`

NewB2xIdentityUserFlowWithDefaults instantiates a new B2xIdentityUserFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiConnectorConfiguration

`func (o *B2xIdentityUserFlow) GetApiConnectorConfiguration() AnyOfmicrosoftGraphUserFlowApiConnectorConfiguration`

GetApiConnectorConfiguration returns the ApiConnectorConfiguration field if non-nil, zero value otherwise.

### GetApiConnectorConfigurationOk

`func (o *B2xIdentityUserFlow) GetApiConnectorConfigurationOk() (*AnyOfmicrosoftGraphUserFlowApiConnectorConfiguration, bool)`

GetApiConnectorConfigurationOk returns a tuple with the ApiConnectorConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiConnectorConfiguration

`func (o *B2xIdentityUserFlow) SetApiConnectorConfiguration(v AnyOfmicrosoftGraphUserFlowApiConnectorConfiguration)`

SetApiConnectorConfiguration sets ApiConnectorConfiguration field to given value.

### HasApiConnectorConfiguration

`func (o *B2xIdentityUserFlow) HasApiConnectorConfiguration() bool`

HasApiConnectorConfiguration returns a boolean if a field has been set.

### SetApiConnectorConfigurationNil

`func (o *B2xIdentityUserFlow) SetApiConnectorConfigurationNil(b bool)`

 SetApiConnectorConfigurationNil sets the value for ApiConnectorConfiguration to be an explicit nil

### UnsetApiConnectorConfiguration
`func (o *B2xIdentityUserFlow) UnsetApiConnectorConfiguration()`

UnsetApiConnectorConfiguration ensures that no value is present for ApiConnectorConfiguration, not even an explicit nil
### GetIdentityProviders

`func (o *B2xIdentityUserFlow) GetIdentityProviders() []MicrosoftGraphIdentityProvider`

GetIdentityProviders returns the IdentityProviders field if non-nil, zero value otherwise.

### GetIdentityProvidersOk

`func (o *B2xIdentityUserFlow) GetIdentityProvidersOk() (*[]MicrosoftGraphIdentityProvider, bool)`

GetIdentityProvidersOk returns a tuple with the IdentityProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviders

`func (o *B2xIdentityUserFlow) SetIdentityProviders(v []MicrosoftGraphIdentityProvider)`

SetIdentityProviders sets IdentityProviders field to given value.

### HasIdentityProviders

`func (o *B2xIdentityUserFlow) HasIdentityProviders() bool`

HasIdentityProviders returns a boolean if a field has been set.

### GetLanguages

`func (o *B2xIdentityUserFlow) GetLanguages() []MicrosoftGraphUserFlowLanguageConfiguration`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *B2xIdentityUserFlow) GetLanguagesOk() (*[]MicrosoftGraphUserFlowLanguageConfiguration, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *B2xIdentityUserFlow) SetLanguages(v []MicrosoftGraphUserFlowLanguageConfiguration)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *B2xIdentityUserFlow) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetUserAttributeAssignments

`func (o *B2xIdentityUserFlow) GetUserAttributeAssignments() []MicrosoftGraphIdentityUserFlowAttributeAssignment`

GetUserAttributeAssignments returns the UserAttributeAssignments field if non-nil, zero value otherwise.

### GetUserAttributeAssignmentsOk

`func (o *B2xIdentityUserFlow) GetUserAttributeAssignmentsOk() (*[]MicrosoftGraphIdentityUserFlowAttributeAssignment, bool)`

GetUserAttributeAssignmentsOk returns a tuple with the UserAttributeAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributeAssignments

`func (o *B2xIdentityUserFlow) SetUserAttributeAssignments(v []MicrosoftGraphIdentityUserFlowAttributeAssignment)`

SetUserAttributeAssignments sets UserAttributeAssignments field to given value.

### HasUserAttributeAssignments

`func (o *B2xIdentityUserFlow) HasUserAttributeAssignments() bool`

HasUserAttributeAssignments returns a boolean if a field has been set.

### GetUserFlowIdentityProviders

`func (o *B2xIdentityUserFlow) GetUserFlowIdentityProviders() []MicrosoftGraphIdentityProviderBase`

GetUserFlowIdentityProviders returns the UserFlowIdentityProviders field if non-nil, zero value otherwise.

### GetUserFlowIdentityProvidersOk

`func (o *B2xIdentityUserFlow) GetUserFlowIdentityProvidersOk() (*[]MicrosoftGraphIdentityProviderBase, bool)`

GetUserFlowIdentityProvidersOk returns a tuple with the UserFlowIdentityProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFlowIdentityProviders

`func (o *B2xIdentityUserFlow) SetUserFlowIdentityProviders(v []MicrosoftGraphIdentityProviderBase)`

SetUserFlowIdentityProviders sets UserFlowIdentityProviders field to given value.

### HasUserFlowIdentityProviders

`func (o *B2xIdentityUserFlow) HasUserFlowIdentityProviders() bool`

HasUserFlowIdentityProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


