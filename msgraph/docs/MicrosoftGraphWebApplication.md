# MicrosoftGraphWebApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HomePageUrl** | Pointer to **NullableString** | Home page or landing page of the application. | [optional] 
**ImplicitGrantSettings** | Pointer to [**AnyOfmicrosoftGraphImplicitGrantSettings**](anyOf&lt;microsoft.graph.implicitGrantSettings&gt;.md) | Specifies whether this web application can request tokens using the OAuth 2.0 implicit flow. | [optional] 
**LogoutUrl** | Pointer to **NullableString** | Specifies the URL that will be used by Microsoft&#39;s authorization service to logout an user using front-channel, back-channel or SAML logout protocols. | [optional] 
**RedirectUris** | Pointer to **[]string** | Specifies the URLs where user tokens are sent for sign-in, or the redirect URIs where OAuth 2.0 authorization codes and access tokens are sent. | [optional] 

## Methods

### NewMicrosoftGraphWebApplication

`func NewMicrosoftGraphWebApplication() *MicrosoftGraphWebApplication`

NewMicrosoftGraphWebApplication instantiates a new MicrosoftGraphWebApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWebApplicationWithDefaults

`func NewMicrosoftGraphWebApplicationWithDefaults() *MicrosoftGraphWebApplication`

NewMicrosoftGraphWebApplicationWithDefaults instantiates a new MicrosoftGraphWebApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHomePageUrl

`func (o *MicrosoftGraphWebApplication) GetHomePageUrl() string`

GetHomePageUrl returns the HomePageUrl field if non-nil, zero value otherwise.

### GetHomePageUrlOk

`func (o *MicrosoftGraphWebApplication) GetHomePageUrlOk() (*string, bool)`

GetHomePageUrlOk returns a tuple with the HomePageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePageUrl

`func (o *MicrosoftGraphWebApplication) SetHomePageUrl(v string)`

SetHomePageUrl sets HomePageUrl field to given value.

### HasHomePageUrl

`func (o *MicrosoftGraphWebApplication) HasHomePageUrl() bool`

HasHomePageUrl returns a boolean if a field has been set.

### SetHomePageUrlNil

`func (o *MicrosoftGraphWebApplication) SetHomePageUrlNil(b bool)`

 SetHomePageUrlNil sets the value for HomePageUrl to be an explicit nil

### UnsetHomePageUrl
`func (o *MicrosoftGraphWebApplication) UnsetHomePageUrl()`

UnsetHomePageUrl ensures that no value is present for HomePageUrl, not even an explicit nil
### GetImplicitGrantSettings

`func (o *MicrosoftGraphWebApplication) GetImplicitGrantSettings() AnyOfmicrosoftGraphImplicitGrantSettings`

GetImplicitGrantSettings returns the ImplicitGrantSettings field if non-nil, zero value otherwise.

### GetImplicitGrantSettingsOk

`func (o *MicrosoftGraphWebApplication) GetImplicitGrantSettingsOk() (*AnyOfmicrosoftGraphImplicitGrantSettings, bool)`

GetImplicitGrantSettingsOk returns a tuple with the ImplicitGrantSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplicitGrantSettings

`func (o *MicrosoftGraphWebApplication) SetImplicitGrantSettings(v AnyOfmicrosoftGraphImplicitGrantSettings)`

SetImplicitGrantSettings sets ImplicitGrantSettings field to given value.

### HasImplicitGrantSettings

`func (o *MicrosoftGraphWebApplication) HasImplicitGrantSettings() bool`

HasImplicitGrantSettings returns a boolean if a field has been set.

### SetImplicitGrantSettingsNil

`func (o *MicrosoftGraphWebApplication) SetImplicitGrantSettingsNil(b bool)`

 SetImplicitGrantSettingsNil sets the value for ImplicitGrantSettings to be an explicit nil

### UnsetImplicitGrantSettings
`func (o *MicrosoftGraphWebApplication) UnsetImplicitGrantSettings()`

UnsetImplicitGrantSettings ensures that no value is present for ImplicitGrantSettings, not even an explicit nil
### GetLogoutUrl

`func (o *MicrosoftGraphWebApplication) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *MicrosoftGraphWebApplication) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *MicrosoftGraphWebApplication) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *MicrosoftGraphWebApplication) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### SetLogoutUrlNil

`func (o *MicrosoftGraphWebApplication) SetLogoutUrlNil(b bool)`

 SetLogoutUrlNil sets the value for LogoutUrl to be an explicit nil

### UnsetLogoutUrl
`func (o *MicrosoftGraphWebApplication) UnsetLogoutUrl()`

UnsetLogoutUrl ensures that no value is present for LogoutUrl, not even an explicit nil
### GetRedirectUris

`func (o *MicrosoftGraphWebApplication) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *MicrosoftGraphWebApplication) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *MicrosoftGraphWebApplication) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *MicrosoftGraphWebApplication) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


