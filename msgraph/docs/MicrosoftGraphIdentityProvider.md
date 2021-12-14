# MicrosoftGraphIdentityProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ClientId** | Pointer to **NullableString** | The client ID for the application. This is the client ID obtained when registering the application with the identity provider. Required. Not nullable. | [optional] 
**ClientSecret** | Pointer to **NullableString** | The client secret for the application. This is the client secret obtained when registering the application with the identity provider. This is write-only. A read operation will return ****.  Required. Not nullable. | [optional] 
**Name** | Pointer to **NullableString** | The display name of the identity provider. Not nullable. | [optional] 
**Type** | Pointer to **NullableString** | The identity provider type is a required field. For B2B scenario: Google, Facebook. For B2C scenario: Microsoft, Google, Amazon, LinkedIn, Facebook, GitHub, Twitter, Weibo, QQ, WeChat, OpenIDConnect. Not nullable. | [optional] 

## Methods

### NewMicrosoftGraphIdentityProvider

`func NewMicrosoftGraphIdentityProvider() *MicrosoftGraphIdentityProvider`

NewMicrosoftGraphIdentityProvider instantiates a new MicrosoftGraphIdentityProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphIdentityProviderWithDefaults

`func NewMicrosoftGraphIdentityProviderWithDefaults() *MicrosoftGraphIdentityProvider`

NewMicrosoftGraphIdentityProviderWithDefaults instantiates a new MicrosoftGraphIdentityProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphIdentityProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphIdentityProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphIdentityProvider) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphIdentityProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClientId

`func (o *MicrosoftGraphIdentityProvider) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *MicrosoftGraphIdentityProvider) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *MicrosoftGraphIdentityProvider) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *MicrosoftGraphIdentityProvider) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *MicrosoftGraphIdentityProvider) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *MicrosoftGraphIdentityProvider) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetClientSecret

`func (o *MicrosoftGraphIdentityProvider) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *MicrosoftGraphIdentityProvider) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *MicrosoftGraphIdentityProvider) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *MicrosoftGraphIdentityProvider) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### SetClientSecretNil

`func (o *MicrosoftGraphIdentityProvider) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *MicrosoftGraphIdentityProvider) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
### GetName

`func (o *MicrosoftGraphIdentityProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphIdentityProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphIdentityProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphIdentityProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphIdentityProvider) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphIdentityProvider) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *MicrosoftGraphIdentityProvider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphIdentityProvider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphIdentityProvider) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphIdentityProvider) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphIdentityProvider) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphIdentityProvider) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


