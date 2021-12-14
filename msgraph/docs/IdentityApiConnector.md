# IdentityApiConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationConfiguration** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | The object which describes the authentication configuration details for calling the API. Basic and PKCS 12 client certificate are supported. | [optional] 
**DisplayName** | Pointer to **NullableString** | The name of the API connector. | [optional] 
**TargetUrl** | Pointer to **NullableString** | The URL of the API endpoint to call. | [optional] 

## Methods

### NewIdentityApiConnector

`func NewIdentityApiConnector() *IdentityApiConnector`

NewIdentityApiConnector instantiates a new IdentityApiConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityApiConnectorWithDefaults

`func NewIdentityApiConnectorWithDefaults() *IdentityApiConnector`

NewIdentityApiConnectorWithDefaults instantiates a new IdentityApiConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationConfiguration

`func (o *IdentityApiConnector) GetAuthenticationConfiguration() AnyOfobject`

GetAuthenticationConfiguration returns the AuthenticationConfiguration field if non-nil, zero value otherwise.

### GetAuthenticationConfigurationOk

`func (o *IdentityApiConnector) GetAuthenticationConfigurationOk() (*AnyOfobject, bool)`

GetAuthenticationConfigurationOk returns a tuple with the AuthenticationConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationConfiguration

`func (o *IdentityApiConnector) SetAuthenticationConfiguration(v AnyOfobject)`

SetAuthenticationConfiguration sets AuthenticationConfiguration field to given value.

### HasAuthenticationConfiguration

`func (o *IdentityApiConnector) HasAuthenticationConfiguration() bool`

HasAuthenticationConfiguration returns a boolean if a field has been set.

### SetAuthenticationConfigurationNil

`func (o *IdentityApiConnector) SetAuthenticationConfigurationNil(b bool)`

 SetAuthenticationConfigurationNil sets the value for AuthenticationConfiguration to be an explicit nil

### UnsetAuthenticationConfiguration
`func (o *IdentityApiConnector) UnsetAuthenticationConfiguration()`

UnsetAuthenticationConfiguration ensures that no value is present for AuthenticationConfiguration, not even an explicit nil
### GetDisplayName

`func (o *IdentityApiConnector) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *IdentityApiConnector) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *IdentityApiConnector) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *IdentityApiConnector) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *IdentityApiConnector) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *IdentityApiConnector) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetTargetUrl

`func (o *IdentityApiConnector) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *IdentityApiConnector) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *IdentityApiConnector) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.

### HasTargetUrl

`func (o *IdentityApiConnector) HasTargetUrl() bool`

HasTargetUrl returns a boolean if a field has been set.

### SetTargetUrlNil

`func (o *IdentityApiConnector) SetTargetUrlNil(b bool)`

 SetTargetUrlNil sets the value for TargetUrl to be an explicit nil

### UnsetTargetUrl
`func (o *IdentityApiConnector) UnsetTargetUrl()`

UnsetTargetUrl ensures that no value is present for TargetUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


