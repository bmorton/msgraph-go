# MicrosoftGraphContract

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DeletedDateTime** | Pointer to **NullableTime** |  | [optional] 
**ContractType** | Pointer to **NullableString** | Type of contract. Possible values are:  SyndicationPartner, BreadthPartner, ResellerPartner. See more in the table below. | [optional] 
**CustomerId** | Pointer to **NullableString** | The unique identifier for the customer tenant referenced by this partnership. Corresponds to the id property of the customer tenant&#39;s organization resource. | [optional] 
**DefaultDomainName** | Pointer to **NullableString** | A copy of the customer tenant&#39;s default domain name. The copy is made when the partnership with the customer is established. It is not automatically updated if the customer tenant&#39;s default domain name changes. | [optional] 
**DisplayName** | Pointer to **NullableString** | A copy of the customer tenant&#39;s display name. The copy is made when the partnership with the customer is established. It is not automatically updated if the customer tenant&#39;s display name changes. | [optional] 

## Methods

### NewMicrosoftGraphContract

`func NewMicrosoftGraphContract() *MicrosoftGraphContract`

NewMicrosoftGraphContract instantiates a new MicrosoftGraphContract object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphContractWithDefaults

`func NewMicrosoftGraphContractWithDefaults() *MicrosoftGraphContract`

NewMicrosoftGraphContractWithDefaults instantiates a new MicrosoftGraphContract object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphContract) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphContract) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphContract) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphContract) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeletedDateTime

`func (o *MicrosoftGraphContract) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *MicrosoftGraphContract) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *MicrosoftGraphContract) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *MicrosoftGraphContract) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

### SetDeletedDateTimeNil

`func (o *MicrosoftGraphContract) SetDeletedDateTimeNil(b bool)`

 SetDeletedDateTimeNil sets the value for DeletedDateTime to be an explicit nil

### UnsetDeletedDateTime
`func (o *MicrosoftGraphContract) UnsetDeletedDateTime()`

UnsetDeletedDateTime ensures that no value is present for DeletedDateTime, not even an explicit nil
### GetContractType

`func (o *MicrosoftGraphContract) GetContractType() string`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *MicrosoftGraphContract) GetContractTypeOk() (*string, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *MicrosoftGraphContract) SetContractType(v string)`

SetContractType sets ContractType field to given value.

### HasContractType

`func (o *MicrosoftGraphContract) HasContractType() bool`

HasContractType returns a boolean if a field has been set.

### SetContractTypeNil

`func (o *MicrosoftGraphContract) SetContractTypeNil(b bool)`

 SetContractTypeNil sets the value for ContractType to be an explicit nil

### UnsetContractType
`func (o *MicrosoftGraphContract) UnsetContractType()`

UnsetContractType ensures that no value is present for ContractType, not even an explicit nil
### GetCustomerId

`func (o *MicrosoftGraphContract) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *MicrosoftGraphContract) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *MicrosoftGraphContract) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *MicrosoftGraphContract) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *MicrosoftGraphContract) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *MicrosoftGraphContract) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetDefaultDomainName

`func (o *MicrosoftGraphContract) GetDefaultDomainName() string`

GetDefaultDomainName returns the DefaultDomainName field if non-nil, zero value otherwise.

### GetDefaultDomainNameOk

`func (o *MicrosoftGraphContract) GetDefaultDomainNameOk() (*string, bool)`

GetDefaultDomainNameOk returns a tuple with the DefaultDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDomainName

`func (o *MicrosoftGraphContract) SetDefaultDomainName(v string)`

SetDefaultDomainName sets DefaultDomainName field to given value.

### HasDefaultDomainName

`func (o *MicrosoftGraphContract) HasDefaultDomainName() bool`

HasDefaultDomainName returns a boolean if a field has been set.

### SetDefaultDomainNameNil

`func (o *MicrosoftGraphContract) SetDefaultDomainNameNil(b bool)`

 SetDefaultDomainNameNil sets the value for DefaultDomainName to be an explicit nil

### UnsetDefaultDomainName
`func (o *MicrosoftGraphContract) UnsetDefaultDomainName()`

UnsetDefaultDomainName ensures that no value is present for DefaultDomainName, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphContract) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphContract) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphContract) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphContract) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphContract) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphContract) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


