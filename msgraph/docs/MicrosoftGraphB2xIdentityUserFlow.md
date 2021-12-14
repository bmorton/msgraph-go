# MicrosoftGraphB2xIdentityUserFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**UserFlowType** | Pointer to [**AnyOfmicrosoftGraphUserFlowType**](anyOf&lt;microsoft.graph.userFlowType&gt;.md) |  | [optional] 
**UserFlowTypeVersion** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) |  | [optional] 
**ApiConnectorConfiguration** | Pointer to [**AnyOfmicrosoftGraphUserFlowApiConnectorConfiguration**](anyOf&lt;microsoft.graph.userFlowApiConnectorConfiguration&gt;.md) | Configuration for enabling an API connector for use as part of the self-service sign-up user flow. You can only obtain the value of this object using Get userFlowApiConnectorConfiguration. | [optional] 
**IdentityProviders** | Pointer to [**[]MicrosoftGraphIdentityProvider**](MicrosoftGraphIdentityProvider.md) | The identity providers included in the user flow. | [optional] 
**Languages** | Pointer to [**[]MicrosoftGraphUserFlowLanguageConfiguration**](MicrosoftGraphUserFlowLanguageConfiguration.md) | The languages supported for customization within the user flow. Language customization is enabled by default in self-service sign-up user flow. You cannot create custom languages in self-service sign-up user flows. | [optional] 
**UserAttributeAssignments** | Pointer to [**[]MicrosoftGraphIdentityUserFlowAttributeAssignment**](MicrosoftGraphIdentityUserFlowAttributeAssignment.md) | The user attribute assignments included in the user flow. | [optional] 
**UserFlowIdentityProviders** | Pointer to [**[]MicrosoftGraphIdentityProviderBase**](MicrosoftGraphIdentityProviderBase.md) |  | [optional] 

## Methods

### NewMicrosoftGraphB2xIdentityUserFlow

`func NewMicrosoftGraphB2xIdentityUserFlow() *MicrosoftGraphB2xIdentityUserFlow`

NewMicrosoftGraphB2xIdentityUserFlow instantiates a new MicrosoftGraphB2xIdentityUserFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphB2xIdentityUserFlowWithDefaults

`func NewMicrosoftGraphB2xIdentityUserFlowWithDefaults() *MicrosoftGraphB2xIdentityUserFlow`

NewMicrosoftGraphB2xIdentityUserFlowWithDefaults instantiates a new MicrosoftGraphB2xIdentityUserFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphB2xIdentityUserFlow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphB2xIdentityUserFlow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphB2xIdentityUserFlow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphB2xIdentityUserFlow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserFlowType

`func (o *MicrosoftGraphB2xIdentityUserFlow) GetUserFlowType() AnyOfmicrosoftGraphUserFlowType`

GetUserFlowType returns the UserFlowType field if non-nil, zero value otherwise.

### GetUserFlowTypeOk

`func (o *MicrosoftGraphB2xIdentityUserFlow) GetUserFlowTypeOk() (*AnyOfmicrosoftGraphUserFlowType, bool)`

GetUserFlowTypeOk returns a tuple with the UserFlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFlowType

`func (o *MicrosoftGraphB2xIdentityUserFlow) SetUserFlowType(v AnyOfmicrosoftGraphUserFlowType)`

SetUserFlowType sets UserFlowType field to given value.

### HasUserFlowType

`func (o *MicrosoftGraphB2xIdentityUserFlow) HasUserFlowType() bool`

HasUserFlowType returns a boolean if a field has been set.

### SetUserFlowTypeNil

`func (o *MicrosoftGraphB2xIdentityUserFlow) SetUserFlowTypeNil(b bool)`

 SetUserFlowTypeNil sets the value for UserFlowType to be an explicit nil

### UnsetUserFlowType
`func (o *MicrosoftGraphB2xIdentityUserFlow) UnsetUserFlowType()`

UnsetUserFlowType ensures that no value is present for UserFlowType, not even an explicit nil
### GetUserFlowTypeVersion

`func (o *MicrosoftGraphB2xIdentityUserFlow) GetUserFlowTypeVersion() AnyOfnumberstringstring`

GetUserFlowTypeVersion returns the UserFlowTypeVersion field if non-nil, zero value otherwise.

### GetUserFlowTypeVersionOk

`func (o *MicrosoftGraphB2xIdentityUserFlow) GetUserFlowTypeVersionOk() (*AnyOfnumberstringstring, bool)`

GetUserFlowTypeVersionOk returns a tuple with the UserFlowTypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFlowTypeVersion

`func (o *MicrosoftGraphB2xIdentityUserFlow) SetUserFlowTypeVersion(v AnyOfnumberstringstring)`

SetUserFlowTypeVersion sets UserFlowTypeVersion field to given value.

### HasUserFlowTypeVersion

`func (o *MicrosoftGraphB2xIdentityUserFlow) HasUserFlowTypeVersion() bool`

HasUserFlowTypeVersion returns a boolean if a field has been set.

### SetUserFlowTypeVersionNil

`func (o *MicrosoftGraphB2xIdentityUserFlow) SetUserFlowTypeVersionNil(b bool)`

 SetUserFlowTypeVersionNil sets the value for UserFlowTypeVersion to be an explicit nil

### UnsetUserFlowTypeVersion
`func (o *MicrosoftGraphB2xIdentityUserFlow) UnsetUserFlowTypeVersion()`

UnsetUserFlowTypeVersion ensures that no value is present for UserFlowTypeVersion, not even an explicit nil
### GetApiConnectorConfiguration

`func (o *MicrosoftGraphB2xIdentityUserFlow) GetApiConnectorConfiguration() AnyOfmicrosoftGraphUserFlowApiConnectorConfiguration`

GetApiConnectorConfiguration returns the ApiConnectorConfiguration field if non-nil, zero value otherwise.

### GetApiConnectorConfigurationOk

`func (o *MicrosoftGraphB2xIdentityUserFlow) GetApiConnectorConfigurationOk() (*AnyOfmicrosoftGraphUserFlowApiConnectorConfiguration, bool)`

GetApiConnectorConfigurationOk returns a tuple with the ApiConnectorConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiConnectorConfiguration

`func (o *MicrosoftGraphB2xIdentityUserFlow) SetApiConnectorConfiguration(v AnyOfmicrosoftGraphUserFlowApiConnectorConfiguration)`

SetApiConnectorConfiguration sets ApiConnectorConfiguration field to given value.

### HasApiConnectorConfiguration

`func (o *MicrosoftGraphB2xIdentityUserFlow) HasApiConnectorConfiguration() bool`

HasApiConnectorConfiguration returns a boolean if a field has been set.

### SetApiConnectorConfigurationNil

`func (o *MicrosoftGraphB2xIdentityUserFlow) SetApiConnectorConfigurationNil(b bool)`

 SetApiConnectorConfigurationNil sets the value for ApiConnectorConfiguration to be an explicit nil

### UnsetApiConnectorConfiguration
`func (o *MicrosoftGraphB2xIdentityUserFlow) UnsetApiConnectorConfiguration()`

UnsetApiConnectorConfiguration ensures that no value is present for ApiConnectorConfiguration, not even an explicit nil
### GetIdentityProviders

`func (o *MicrosoftGraphB2xIdentityUserFlow) GetIdentityProviders() []MicrosoftGraphIdentityProvider`

GetIdentityProviders returns the IdentityProviders field if non-nil, zero value otherwise.

### GetIdentityProvidersOk

`func (o *MicrosoftGraphB2xIdentityUserFlow) GetIdentityProvidersOk() (*[]MicrosoftGraphIdentityProvider, bool)`

GetIdentityProvidersOk returns a tuple with the IdentityProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviders

`func (o *MicrosoftGraphB2xIdentityUserFlow) SetIdentityProviders(v []MicrosoftGraphIdentityProvider)`

SetIdentityProviders sets IdentityProviders field to given value.

### HasIdentityProviders

`func (o *MicrosoftGraphB2xIdentityUserFlow) HasIdentityProviders() bool`

HasIdentityProviders returns a boolean if a field has been set.

### GetLanguages

`func (o *MicrosoftGraphB2xIdentityUserFlow) GetLanguages() []MicrosoftGraphUserFlowLanguageConfiguration`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *MicrosoftGraphB2xIdentityUserFlow) GetLanguagesOk() (*[]MicrosoftGraphUserFlowLanguageConfiguration, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *MicrosoftGraphB2xIdentityUserFlow) SetLanguages(v []MicrosoftGraphUserFlowLanguageConfiguration)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *MicrosoftGraphB2xIdentityUserFlow) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetUserAttributeAssignments

`func (o *MicrosoftGraphB2xIdentityUserFlow) GetUserAttributeAssignments() []MicrosoftGraphIdentityUserFlowAttributeAssignment`

GetUserAttributeAssignments returns the UserAttributeAssignments field if non-nil, zero value otherwise.

### GetUserAttributeAssignmentsOk

`func (o *MicrosoftGraphB2xIdentityUserFlow) GetUserAttributeAssignmentsOk() (*[]MicrosoftGraphIdentityUserFlowAttributeAssignment, bool)`

GetUserAttributeAssignmentsOk returns a tuple with the UserAttributeAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributeAssignments

`func (o *MicrosoftGraphB2xIdentityUserFlow) SetUserAttributeAssignments(v []MicrosoftGraphIdentityUserFlowAttributeAssignment)`

SetUserAttributeAssignments sets UserAttributeAssignments field to given value.

### HasUserAttributeAssignments

`func (o *MicrosoftGraphB2xIdentityUserFlow) HasUserAttributeAssignments() bool`

HasUserAttributeAssignments returns a boolean if a field has been set.

### GetUserFlowIdentityProviders

`func (o *MicrosoftGraphB2xIdentityUserFlow) GetUserFlowIdentityProviders() []MicrosoftGraphIdentityProviderBase`

GetUserFlowIdentityProviders returns the UserFlowIdentityProviders field if non-nil, zero value otherwise.

### GetUserFlowIdentityProvidersOk

`func (o *MicrosoftGraphB2xIdentityUserFlow) GetUserFlowIdentityProvidersOk() (*[]MicrosoftGraphIdentityProviderBase, bool)`

GetUserFlowIdentityProvidersOk returns a tuple with the UserFlowIdentityProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFlowIdentityProviders

`func (o *MicrosoftGraphB2xIdentityUserFlow) SetUserFlowIdentityProviders(v []MicrosoftGraphIdentityProviderBase)`

SetUserFlowIdentityProviders sets UserFlowIdentityProviders field to given value.

### HasUserFlowIdentityProviders

`func (o *MicrosoftGraphB2xIdentityUserFlow) HasUserFlowIdentityProviders() bool`

HasUserFlowIdentityProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


