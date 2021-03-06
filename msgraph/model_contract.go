/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
)

// Contract struct for Contract
type Contract struct {
	// Type of contract. Possible values are:  SyndicationPartner, BreadthPartner, ResellerPartner. See more in the table below.
	ContractType NullableString `json:"contractType,omitempty"`
	// The unique identifier for the customer tenant referenced by this partnership. Corresponds to the id property of the customer tenant's organization resource.
	CustomerId NullableString `json:"customerId,omitempty"`
	// A copy of the customer tenant's default domain name. The copy is made when the partnership with the customer is established. It is not automatically updated if the customer tenant's default domain name changes.
	DefaultDomainName NullableString `json:"defaultDomainName,omitempty"`
	// A copy of the customer tenant's display name. The copy is made when the partnership with the customer is established. It is not automatically updated if the customer tenant's display name changes.
	DisplayName NullableString `json:"displayName,omitempty"`
}

// NewContract instantiates a new Contract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContract() *Contract {
	this := Contract{}
	return &this
}

// NewContractWithDefaults instantiates a new Contract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractWithDefaults() *Contract {
	this := Contract{}
	return &this
}

// GetContractType returns the ContractType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Contract) GetContractType() string {
	if o == nil || o.ContractType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContractType.Get()
}

// GetContractTypeOk returns a tuple with the ContractType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Contract) GetContractTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContractType.Get(), o.ContractType.IsSet()
}

// HasContractType returns a boolean if a field has been set.
func (o *Contract) HasContractType() bool {
	if o != nil && o.ContractType.IsSet() {
		return true
	}

	return false
}

// SetContractType gets a reference to the given NullableString and assigns it to the ContractType field.
func (o *Contract) SetContractType(v string) {
	o.ContractType.Set(&v)
}
// SetContractTypeNil sets the value for ContractType to be an explicit nil
func (o *Contract) SetContractTypeNil() {
	o.ContractType.Set(nil)
}

// UnsetContractType ensures that no value is present for ContractType, not even an explicit nil
func (o *Contract) UnsetContractType() {
	o.ContractType.Unset()
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Contract) GetCustomerId() string {
	if o == nil || o.CustomerId.Get() == nil {
		var ret string
		return ret
	}
	return *o.CustomerId.Get()
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Contract) GetCustomerIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CustomerId.Get(), o.CustomerId.IsSet()
}

// HasCustomerId returns a boolean if a field has been set.
func (o *Contract) HasCustomerId() bool {
	if o != nil && o.CustomerId.IsSet() {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given NullableString and assigns it to the CustomerId field.
func (o *Contract) SetCustomerId(v string) {
	o.CustomerId.Set(&v)
}
// SetCustomerIdNil sets the value for CustomerId to be an explicit nil
func (o *Contract) SetCustomerIdNil() {
	o.CustomerId.Set(nil)
}

// UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
func (o *Contract) UnsetCustomerId() {
	o.CustomerId.Unset()
}

// GetDefaultDomainName returns the DefaultDomainName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Contract) GetDefaultDomainName() string {
	if o == nil || o.DefaultDomainName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DefaultDomainName.Get()
}

// GetDefaultDomainNameOk returns a tuple with the DefaultDomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Contract) GetDefaultDomainNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DefaultDomainName.Get(), o.DefaultDomainName.IsSet()
}

// HasDefaultDomainName returns a boolean if a field has been set.
func (o *Contract) HasDefaultDomainName() bool {
	if o != nil && o.DefaultDomainName.IsSet() {
		return true
	}

	return false
}

// SetDefaultDomainName gets a reference to the given NullableString and assigns it to the DefaultDomainName field.
func (o *Contract) SetDefaultDomainName(v string) {
	o.DefaultDomainName.Set(&v)
}
// SetDefaultDomainNameNil sets the value for DefaultDomainName to be an explicit nil
func (o *Contract) SetDefaultDomainNameNil() {
	o.DefaultDomainName.Set(nil)
}

// UnsetDefaultDomainName ensures that no value is present for DefaultDomainName, not even an explicit nil
func (o *Contract) UnsetDefaultDomainName() {
	o.DefaultDomainName.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Contract) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Contract) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Contract) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *Contract) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *Contract) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *Contract) UnsetDisplayName() {
	o.DisplayName.Unset()
}

func (o Contract) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContractType.IsSet() {
		toSerialize["contractType"] = o.ContractType.Get()
	}
	if o.CustomerId.IsSet() {
		toSerialize["customerId"] = o.CustomerId.Get()
	}
	if o.DefaultDomainName.IsSet() {
		toSerialize["defaultDomainName"] = o.DefaultDomainName.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableContract struct {
	value *Contract
	isSet bool
}

func (v NullableContract) Get() *Contract {
	return v.value
}

func (v *NullableContract) Set(val *Contract) {
	v.value = val
	v.isSet = true
}

func (v NullableContract) IsSet() bool {
	return v.isSet
}

func (v *NullableContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContract(val *Contract) *NullableContract {
	return &NullableContract{value: val, isSet: true}
}

func (v NullableContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


