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

// CollectionOfDeviceComplianceScheduledActionForRule struct for CollectionOfDeviceComplianceScheduledActionForRule
type CollectionOfDeviceComplianceScheduledActionForRule struct {
	Value *[]MicrosoftGraphDeviceComplianceScheduledActionForRule `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfDeviceComplianceScheduledActionForRule instantiates a new CollectionOfDeviceComplianceScheduledActionForRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfDeviceComplianceScheduledActionForRule() *CollectionOfDeviceComplianceScheduledActionForRule {
	this := CollectionOfDeviceComplianceScheduledActionForRule{}
	return &this
}

// NewCollectionOfDeviceComplianceScheduledActionForRuleWithDefaults instantiates a new CollectionOfDeviceComplianceScheduledActionForRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfDeviceComplianceScheduledActionForRuleWithDefaults() *CollectionOfDeviceComplianceScheduledActionForRule {
	this := CollectionOfDeviceComplianceScheduledActionForRule{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfDeviceComplianceScheduledActionForRule) GetValue() []MicrosoftGraphDeviceComplianceScheduledActionForRule {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphDeviceComplianceScheduledActionForRule
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfDeviceComplianceScheduledActionForRule) GetValueOk() (*[]MicrosoftGraphDeviceComplianceScheduledActionForRule, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfDeviceComplianceScheduledActionForRule) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphDeviceComplianceScheduledActionForRule and assigns it to the Value field.
func (o *CollectionOfDeviceComplianceScheduledActionForRule) SetValue(v []MicrosoftGraphDeviceComplianceScheduledActionForRule) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfDeviceComplianceScheduledActionForRule) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfDeviceComplianceScheduledActionForRule) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfDeviceComplianceScheduledActionForRule) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfDeviceComplianceScheduledActionForRule) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfDeviceComplianceScheduledActionForRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfDeviceComplianceScheduledActionForRule struct {
	value *CollectionOfDeviceComplianceScheduledActionForRule
	isSet bool
}

func (v NullableCollectionOfDeviceComplianceScheduledActionForRule) Get() *CollectionOfDeviceComplianceScheduledActionForRule {
	return v.value
}

func (v *NullableCollectionOfDeviceComplianceScheduledActionForRule) Set(val *CollectionOfDeviceComplianceScheduledActionForRule) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfDeviceComplianceScheduledActionForRule) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfDeviceComplianceScheduledActionForRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfDeviceComplianceScheduledActionForRule(val *CollectionOfDeviceComplianceScheduledActionForRule) *NullableCollectionOfDeviceComplianceScheduledActionForRule {
	return &NullableCollectionOfDeviceComplianceScheduledActionForRule{value: val, isSet: true}
}

func (v NullableCollectionOfDeviceComplianceScheduledActionForRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfDeviceComplianceScheduledActionForRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


