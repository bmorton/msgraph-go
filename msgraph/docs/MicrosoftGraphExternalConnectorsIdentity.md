# MicrosoftGraphExternalConnectorsIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Type** | Pointer to [**AnyOfmicrosoftGraphExternalConnectorsIdentityType**](anyOf&lt;microsoft.graph.externalConnectors.identityType&gt;.md) | The type of identity. Possible values are: user or group for Azure AD identities and externalgroup for groups in an external system. | [optional] 

## Methods

### NewMicrosoftGraphExternalConnectorsIdentity

`func NewMicrosoftGraphExternalConnectorsIdentity() *MicrosoftGraphExternalConnectorsIdentity`

NewMicrosoftGraphExternalConnectorsIdentity instantiates a new MicrosoftGraphExternalConnectorsIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphExternalConnectorsIdentityWithDefaults

`func NewMicrosoftGraphExternalConnectorsIdentityWithDefaults() *MicrosoftGraphExternalConnectorsIdentity`

NewMicrosoftGraphExternalConnectorsIdentityWithDefaults instantiates a new MicrosoftGraphExternalConnectorsIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphExternalConnectorsIdentity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphExternalConnectorsIdentity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphExternalConnectorsIdentity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphExternalConnectorsIdentity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *MicrosoftGraphExternalConnectorsIdentity) GetType() AnyOfmicrosoftGraphExternalConnectorsIdentityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphExternalConnectorsIdentity) GetTypeOk() (*AnyOfmicrosoftGraphExternalConnectorsIdentityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphExternalConnectorsIdentity) SetType(v AnyOfmicrosoftGraphExternalConnectorsIdentityType)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphExternalConnectorsIdentity) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphExternalConnectorsIdentity) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphExternalConnectorsIdentity) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


