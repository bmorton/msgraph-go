# MicrosoftGraphExternalConnectorsAcl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | Pointer to [**AnyOfmicrosoftGraphExternalConnectorsAccessType**](anyOf&lt;microsoft.graph.externalConnectors.accessType&gt;.md) | The access granted to the identity. Possible values are: grant, deny, unknownFutureValue. | [optional] 
**Type** | Pointer to [**AnyOfmicrosoftGraphExternalConnectorsAclType**](anyOf&lt;microsoft.graph.externalConnectors.aclType&gt;.md) | The type of identity. Possible values are: user, group, everyone, everyoneExceptGuests, externalGroup, unknownFutureValue. | [optional] 
**Value** | Pointer to **string** | The unique identifer of the identity. In case of Azure Active Directory identities, value is set to the object identifier of the user, group or tenant for types user, group and everyone (and everyoneExceptGuests) respectively. In case of external groups value is set to the ID of the externalGroup | [optional] 

## Methods

### NewMicrosoftGraphExternalConnectorsAcl

`func NewMicrosoftGraphExternalConnectorsAcl() *MicrosoftGraphExternalConnectorsAcl`

NewMicrosoftGraphExternalConnectorsAcl instantiates a new MicrosoftGraphExternalConnectorsAcl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphExternalConnectorsAclWithDefaults

`func NewMicrosoftGraphExternalConnectorsAclWithDefaults() *MicrosoftGraphExternalConnectorsAcl`

NewMicrosoftGraphExternalConnectorsAclWithDefaults instantiates a new MicrosoftGraphExternalConnectorsAcl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *MicrosoftGraphExternalConnectorsAcl) GetAccessType() AnyOfmicrosoftGraphExternalConnectorsAccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *MicrosoftGraphExternalConnectorsAcl) GetAccessTypeOk() (*AnyOfmicrosoftGraphExternalConnectorsAccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *MicrosoftGraphExternalConnectorsAcl) SetAccessType(v AnyOfmicrosoftGraphExternalConnectorsAccessType)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *MicrosoftGraphExternalConnectorsAcl) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### SetAccessTypeNil

`func (o *MicrosoftGraphExternalConnectorsAcl) SetAccessTypeNil(b bool)`

 SetAccessTypeNil sets the value for AccessType to be an explicit nil

### UnsetAccessType
`func (o *MicrosoftGraphExternalConnectorsAcl) UnsetAccessType()`

UnsetAccessType ensures that no value is present for AccessType, not even an explicit nil
### GetType

`func (o *MicrosoftGraphExternalConnectorsAcl) GetType() AnyOfmicrosoftGraphExternalConnectorsAclType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphExternalConnectorsAcl) GetTypeOk() (*AnyOfmicrosoftGraphExternalConnectorsAclType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphExternalConnectorsAcl) SetType(v AnyOfmicrosoftGraphExternalConnectorsAclType)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphExternalConnectorsAcl) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphExternalConnectorsAcl) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphExternalConnectorsAcl) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetValue

`func (o *MicrosoftGraphExternalConnectorsAcl) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MicrosoftGraphExternalConnectorsAcl) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MicrosoftGraphExternalConnectorsAcl) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MicrosoftGraphExternalConnectorsAcl) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


