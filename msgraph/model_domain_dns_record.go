/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// DomainDnsRecord struct for DomainDnsRecord
type DomainDnsRecord struct {
	// If false, this record must be configured by the customer at the DNS host for Microsoft Online Services to operate correctly with the domain.
	IsOptional *bool `json:"isOptional,omitempty"`
	// Value used when configuring the name of the DNS record at the DNS host.
	Label *string `json:"label,omitempty"`
	// Indicates what type of DNS record this entity represents.The value can be one of the following: CName, Mx, Srv, TxtKey
	RecordType NullableString `json:"recordType,omitempty"`
	// Microsoft Online Service or feature that has a dependency on this DNS record.Can be one of the following values: null, Email, Sharepoint, EmailInternalRelayOnly, OfficeCommunicationsOnline, SharePointDefaultDomain, FullRedelegation, SharePointPublic, OrgIdAuthentication, Yammer, Intune
	SupportedService *string `json:"supportedService,omitempty"`
	// Value to use when configuring the time-to-live (ttl) property of the DNS record at the DNS host. Not nullable
	Ttl *int32 `json:"ttl,omitempty"`
}

// NewDomainDnsRecord instantiates a new DomainDnsRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainDnsRecord() *DomainDnsRecord {
	this := DomainDnsRecord{}
	return &this
}

// NewDomainDnsRecordWithDefaults instantiates a new DomainDnsRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainDnsRecordWithDefaults() *DomainDnsRecord {
	this := DomainDnsRecord{}
	return &this
}

// GetIsOptional returns the IsOptional field value if set, zero value otherwise.
func (o *DomainDnsRecord) GetIsOptional() bool {
	if o == nil || o.IsOptional == nil {
		var ret bool
		return ret
	}
	return *o.IsOptional
}

// GetIsOptionalOk returns a tuple with the IsOptional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDnsRecord) GetIsOptionalOk() (*bool, bool) {
	if o == nil || o.IsOptional == nil {
		return nil, false
	}
	return o.IsOptional, true
}

// HasIsOptional returns a boolean if a field has been set.
func (o *DomainDnsRecord) HasIsOptional() bool {
	if o != nil && o.IsOptional != nil {
		return true
	}

	return false
}

// SetIsOptional gets a reference to the given bool and assigns it to the IsOptional field.
func (o *DomainDnsRecord) SetIsOptional(v bool) {
	o.IsOptional = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *DomainDnsRecord) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDnsRecord) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *DomainDnsRecord) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *DomainDnsRecord) SetLabel(v string) {
	o.Label = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DomainDnsRecord) GetRecordType() string {
	if o == nil || o.RecordType.Get() == nil {
		var ret string
		return ret
	}
	return *o.RecordType.Get()
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DomainDnsRecord) GetRecordTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RecordType.Get(), o.RecordType.IsSet()
}

// HasRecordType returns a boolean if a field has been set.
func (o *DomainDnsRecord) HasRecordType() bool {
	if o != nil && o.RecordType.IsSet() {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given NullableString and assigns it to the RecordType field.
func (o *DomainDnsRecord) SetRecordType(v string) {
	o.RecordType.Set(&v)
}
// SetRecordTypeNil sets the value for RecordType to be an explicit nil
func (o *DomainDnsRecord) SetRecordTypeNil() {
	o.RecordType.Set(nil)
}

// UnsetRecordType ensures that no value is present for RecordType, not even an explicit nil
func (o *DomainDnsRecord) UnsetRecordType() {
	o.RecordType.Unset()
}

// GetSupportedService returns the SupportedService field value if set, zero value otherwise.
func (o *DomainDnsRecord) GetSupportedService() string {
	if o == nil || o.SupportedService == nil {
		var ret string
		return ret
	}
	return *o.SupportedService
}

// GetSupportedServiceOk returns a tuple with the SupportedService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDnsRecord) GetSupportedServiceOk() (*string, bool) {
	if o == nil || o.SupportedService == nil {
		return nil, false
	}
	return o.SupportedService, true
}

// HasSupportedService returns a boolean if a field has been set.
func (o *DomainDnsRecord) HasSupportedService() bool {
	if o != nil && o.SupportedService != nil {
		return true
	}

	return false
}

// SetSupportedService gets a reference to the given string and assigns it to the SupportedService field.
func (o *DomainDnsRecord) SetSupportedService(v string) {
	o.SupportedService = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *DomainDnsRecord) GetTtl() int32 {
	if o == nil || o.Ttl == nil {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainDnsRecord) GetTtlOk() (*int32, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *DomainDnsRecord) HasTtl() bool {
	if o != nil && o.Ttl != nil {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *DomainDnsRecord) SetTtl(v int32) {
	o.Ttl = &v
}

func (o DomainDnsRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsOptional != nil {
		toSerialize["isOptional"] = o.IsOptional
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.RecordType.IsSet() {
		toSerialize["recordType"] = o.RecordType.Get()
	}
	if o.SupportedService != nil {
		toSerialize["supportedService"] = o.SupportedService
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	return json.Marshal(toSerialize)
}

type NullableDomainDnsRecord struct {
	value *DomainDnsRecord
	isSet bool
}

func (v NullableDomainDnsRecord) Get() *DomainDnsRecord {
	return v.value
}

func (v *NullableDomainDnsRecord) Set(val *DomainDnsRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainDnsRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainDnsRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainDnsRecord(val *DomainDnsRecord) *NullableDomainDnsRecord {
	return &NullableDomainDnsRecord{value: val, isSet: true}
}

func (v NullableDomainDnsRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainDnsRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


