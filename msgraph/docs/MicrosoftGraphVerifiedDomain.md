# MicrosoftGraphVerifiedDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to **NullableString** | For example, &#39;Email&#39;, &#39;OfficeCommunicationsOnline&#39;. | [optional] 
**IsDefault** | Pointer to **NullableBool** | true if this is the default domain associated with the tenant; otherwise, false. | [optional] 
**IsInitial** | Pointer to **NullableBool** | true if this is the initial domain associated with the tenant; otherwise, false | [optional] 
**Name** | Pointer to **NullableString** | The domain name; for example, &#39;contoso.onmicrosoft.com&#39; | [optional] 
**Type** | Pointer to **NullableString** | For example, &#39;Managed&#39;. | [optional] 

## Methods

### NewMicrosoftGraphVerifiedDomain

`func NewMicrosoftGraphVerifiedDomain() *MicrosoftGraphVerifiedDomain`

NewMicrosoftGraphVerifiedDomain instantiates a new MicrosoftGraphVerifiedDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphVerifiedDomainWithDefaults

`func NewMicrosoftGraphVerifiedDomainWithDefaults() *MicrosoftGraphVerifiedDomain`

NewMicrosoftGraphVerifiedDomainWithDefaults instantiates a new MicrosoftGraphVerifiedDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *MicrosoftGraphVerifiedDomain) GetCapabilities() string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *MicrosoftGraphVerifiedDomain) GetCapabilitiesOk() (*string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *MicrosoftGraphVerifiedDomain) SetCapabilities(v string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *MicrosoftGraphVerifiedDomain) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *MicrosoftGraphVerifiedDomain) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *MicrosoftGraphVerifiedDomain) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetIsDefault

`func (o *MicrosoftGraphVerifiedDomain) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *MicrosoftGraphVerifiedDomain) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *MicrosoftGraphVerifiedDomain) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *MicrosoftGraphVerifiedDomain) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### SetIsDefaultNil

`func (o *MicrosoftGraphVerifiedDomain) SetIsDefaultNil(b bool)`

 SetIsDefaultNil sets the value for IsDefault to be an explicit nil

### UnsetIsDefault
`func (o *MicrosoftGraphVerifiedDomain) UnsetIsDefault()`

UnsetIsDefault ensures that no value is present for IsDefault, not even an explicit nil
### GetIsInitial

`func (o *MicrosoftGraphVerifiedDomain) GetIsInitial() bool`

GetIsInitial returns the IsInitial field if non-nil, zero value otherwise.

### GetIsInitialOk

`func (o *MicrosoftGraphVerifiedDomain) GetIsInitialOk() (*bool, bool)`

GetIsInitialOk returns a tuple with the IsInitial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInitial

`func (o *MicrosoftGraphVerifiedDomain) SetIsInitial(v bool)`

SetIsInitial sets IsInitial field to given value.

### HasIsInitial

`func (o *MicrosoftGraphVerifiedDomain) HasIsInitial() bool`

HasIsInitial returns a boolean if a field has been set.

### SetIsInitialNil

`func (o *MicrosoftGraphVerifiedDomain) SetIsInitialNil(b bool)`

 SetIsInitialNil sets the value for IsInitial to be an explicit nil

### UnsetIsInitial
`func (o *MicrosoftGraphVerifiedDomain) UnsetIsInitial()`

UnsetIsInitial ensures that no value is present for IsInitial, not even an explicit nil
### GetName

`func (o *MicrosoftGraphVerifiedDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphVerifiedDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphVerifiedDomain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphVerifiedDomain) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphVerifiedDomain) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphVerifiedDomain) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *MicrosoftGraphVerifiedDomain) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphVerifiedDomain) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphVerifiedDomain) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphVerifiedDomain) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphVerifiedDomain) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphVerifiedDomain) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


