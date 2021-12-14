# Contract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractType** | Pointer to **NullableString** | Type of contract. Possible values are:  SyndicationPartner, BreadthPartner, ResellerPartner. See more in the table below. | [optional] 
**CustomerId** | Pointer to **NullableString** | The unique identifier for the customer tenant referenced by this partnership. Corresponds to the id property of the customer tenant&#39;s organization resource. | [optional] 
**DefaultDomainName** | Pointer to **NullableString** | A copy of the customer tenant&#39;s default domain name. The copy is made when the partnership with the customer is established. It is not automatically updated if the customer tenant&#39;s default domain name changes. | [optional] 
**DisplayName** | Pointer to **NullableString** | A copy of the customer tenant&#39;s display name. The copy is made when the partnership with the customer is established. It is not automatically updated if the customer tenant&#39;s display name changes. | [optional] 

## Methods

### NewContract

`func NewContract() *Contract`

NewContract instantiates a new Contract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractWithDefaults

`func NewContractWithDefaults() *Contract`

NewContractWithDefaults instantiates a new Contract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractType

`func (o *Contract) GetContractType() string`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *Contract) GetContractTypeOk() (*string, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *Contract) SetContractType(v string)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *Contract) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### SetContractTypeNil

`func (o *Contract) SetContractTypeNil(b bool)`

 SetContractTypeNil sets the value for ContractType to be an explicit nil

### UnsetContractType
`func (o *Contract) UnsetContractType()`

UnsetContractType ensures that no value is present for ContractType, not even an explicit nil
### GetCustomerId

`func (o *Contract) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *Contract) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *Contract) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *Contract) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *Contract) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *Contract) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetDefaultDomainName

`func (o *Contract) GetDefaultDomainName() string`

GetDefaultDomainName returns the DefaultDomainName field if non-nil, zero value otherwise.

### GetDefaultDomainNameOk

`func (o *Contract) GetDefaultDomainNameOk() (*string, bool)`

GetDefaultDomainNameOk returns a tuple with the DefaultDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDomainName

`func (o *Contract) SetDefaultDomainName(v string)`

SetDefaultDomainName sets DefaultDomainName field to given value.

### HasDefaultDomainName

`func (o *Contract) HasDefaultDomainName() bool`

HasDefaultDomainName returns a boolean if a field has been set.

### SetDefaultDomainNameNil

`func (o *Contract) SetDefaultDomainNameNil(b bool)`

 SetDefaultDomainNameNil sets the value for DefaultDomainName to be an explicit nil

### UnsetDefaultDomainName
`func (o *Contract) UnsetDefaultDomainName()`

UnsetDefaultDomainName ensures that no value is present for DefaultDomainName, not even an explicit nil
### GetDisplayName

`func (o *Contract) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Contract) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Contract) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Contract) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *Contract) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *Contract) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


