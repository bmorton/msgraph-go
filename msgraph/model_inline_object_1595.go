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

// InlineObject1595 struct for InlineObject1595
type InlineObject1595 struct {
	Criteria AnyOfmicrosoftGraphWorkbookFilterCriteria `json:"criteria,omitempty"`
}

// NewInlineObject1595 instantiates a new InlineObject1595 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1595() *InlineObject1595 {
	this := InlineObject1595{}
	return &this
}

// NewInlineObject1595WithDefaults instantiates a new InlineObject1595 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1595WithDefaults() *InlineObject1595 {
	this := InlineObject1595{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1595) GetCriteria() AnyOfmicrosoftGraphWorkbookFilterCriteria {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookFilterCriteria
		return ret
	}
	return o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1595) GetCriteriaOk() (*AnyOfmicrosoftGraphWorkbookFilterCriteria, bool) {
	if o == nil || o.Criteria == nil {
		return nil, false
	}
	return &o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *InlineObject1595) HasCriteria() bool {
	if o != nil && o.Criteria != nil {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given AnyOfmicrosoftGraphWorkbookFilterCriteria and assigns it to the Criteria field.
func (o *InlineObject1595) SetCriteria(v AnyOfmicrosoftGraphWorkbookFilterCriteria) {
	o.Criteria = v
}

func (o InlineObject1595) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Criteria != nil {
		toSerialize["criteria"] = o.Criteria
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1595 struct {
	value *InlineObject1595
	isSet bool
}

func (v NullableInlineObject1595) Get() *InlineObject1595 {
	return v.value
}

func (v *NullableInlineObject1595) Set(val *InlineObject1595) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1595) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1595) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1595(val *InlineObject1595) *NullableInlineObject1595 {
	return &NullableInlineObject1595{value: val, isSet: true}
}

func (v NullableInlineObject1595) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1595) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


