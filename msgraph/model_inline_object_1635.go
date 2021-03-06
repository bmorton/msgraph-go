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

// InlineObject1635 struct for InlineObject1635
type InlineObject1635 struct {
	Criteria AnyOfmicrosoftGraphWorkbookFilterCriteria `json:"criteria,omitempty"`
}

// NewInlineObject1635 instantiates a new InlineObject1635 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1635() *InlineObject1635 {
	this := InlineObject1635{}
	return &this
}

// NewInlineObject1635WithDefaults instantiates a new InlineObject1635 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1635WithDefaults() *InlineObject1635 {
	this := InlineObject1635{}
	return &this
}

// GetCriteria returns the Criteria field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1635) GetCriteria() AnyOfmicrosoftGraphWorkbookFilterCriteria {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookFilterCriteria
		return ret
	}
	return o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1635) GetCriteriaOk() (*AnyOfmicrosoftGraphWorkbookFilterCriteria, bool) {
	if o == nil || o.Criteria == nil {
		return nil, false
	}
	return &o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *InlineObject1635) HasCriteria() bool {
	if o != nil && o.Criteria != nil {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given AnyOfmicrosoftGraphWorkbookFilterCriteria and assigns it to the Criteria field.
func (o *InlineObject1635) SetCriteria(v AnyOfmicrosoftGraphWorkbookFilterCriteria) {
	o.Criteria = v
}

func (o InlineObject1635) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Criteria != nil {
		toSerialize["criteria"] = o.Criteria
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1635 struct {
	value *InlineObject1635
	isSet bool
}

func (v NullableInlineObject1635) Get() *InlineObject1635 {
	return v.value
}

func (v *NullableInlineObject1635) Set(val *InlineObject1635) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1635) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1635) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1635(val *InlineObject1635) *NullableInlineObject1635 {
	return &NullableInlineObject1635{value: val, isSet: true}
}

func (v NullableInlineObject1635) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1635) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


