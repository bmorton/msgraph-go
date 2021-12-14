# ProvisionedIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Details of the identity. | [optional] 
**IdentityType** | Pointer to **NullableString** | Type of identity that has been provisioned, such as &#39;user&#39; or &#39;group&#39;. | [optional] 

## Methods

### NewProvisionedIdentity

`func NewProvisionedIdentity() *ProvisionedIdentity`

NewProvisionedIdentity instantiates a new ProvisionedIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionedIdentityWithDefaults

`func NewProvisionedIdentityWithDefaults() *ProvisionedIdentity`

NewProvisionedIdentityWithDefaults instantiates a new ProvisionedIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *ProvisionedIdentity) GetDetails() AnyOfobject`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ProvisionedIdentity) GetDetailsOk() (*AnyOfobject, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ProvisionedIdentity) SetDetails(v AnyOfobject)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ProvisionedIdentity) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetailsNil

`func (o *ProvisionedIdentity) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *ProvisionedIdentity) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil
### GetIdentityType

`func (o *ProvisionedIdentity) GetIdentityType() string`

GetIdentityType returns the IdentityType field if non-nil, zero value otherwise.

### GetIdentityTypeOk

`func (o *ProvisionedIdentity) GetIdentityTypeOk() (*string, bool)`

GetIdentityTypeOk returns a tuple with the IdentityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityType

`func (o *ProvisionedIdentity) SetIdentityType(v string)`

SetIdentityType sets IdentityType field to given value.

### HasIdentityType

`func (o *ProvisionedIdentity) HasIdentityType() bool`

HasIdentityType returns a boolean if a field has been set.

### SetIdentityTypeNil

`func (o *ProvisionedIdentity) SetIdentityTypeNil(b bool)`

 SetIdentityTypeNil sets the value for IdentityType to be an explicit nil

### UnsetIdentityType
`func (o *ProvisionedIdentity) UnsetIdentityType()`

UnsetIdentityType ensures that no value is present for IdentityType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


