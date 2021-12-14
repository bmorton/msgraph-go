# Identity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**AnyOfmicrosoftGraphExternalConnectorsIdentityType**](anyOf&lt;microsoft.graph.externalConnectors.identityType&gt;.md) | The type of identity. Possible values are: user or group for Azure AD identities and externalgroup for groups in an external system. | [optional] 

## Methods

### NewIdentity

`func NewIdentity() *Identity`

NewIdentity instantiates a new Identity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityWithDefaults

`func NewIdentityWithDefaults() *Identity`

NewIdentityWithDefaults instantiates a new Identity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Identity) GetType() AnyOfmicrosoftGraphExternalConnectorsIdentityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Identity) GetTypeOk() (*AnyOfmicrosoftGraphExternalConnectorsIdentityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Identity) SetType(v AnyOfmicrosoftGraphExternalConnectorsIdentityType)`

SetType sets Type field to given value.

### HasType

`func (o *Identity) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Identity) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Identity) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


