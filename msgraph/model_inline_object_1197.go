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

// InlineObject1197 struct for InlineObject1197
type InlineObject1197 struct {
	Name NullableString `json:"name,omitempty"`
	ParentReference AnyOfmicrosoftGraphItemReference `json:"parentReference,omitempty"`
}

// NewInlineObject1197 instantiates a new InlineObject1197 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1197() *InlineObject1197 {
	this := InlineObject1197{}
	return &this
}

// NewInlineObject1197WithDefaults instantiates a new InlineObject1197 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1197WithDefaults() *InlineObject1197 {
	this := InlineObject1197{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1197) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1197) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject1197) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *InlineObject1197) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *InlineObject1197) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *InlineObject1197) UnsetName() {
	o.Name.Unset()
}

// GetParentReference returns the ParentReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1197) GetParentReference() AnyOfmicrosoftGraphItemReference {
	if o == nil  {
		var ret AnyOfmicrosoftGraphItemReference
		return ret
	}
	return o.ParentReference
}

// GetParentReferenceOk returns a tuple with the ParentReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1197) GetParentReferenceOk() (*AnyOfmicrosoftGraphItemReference, bool) {
	if o == nil || o.ParentReference == nil {
		return nil, false
	}
	return &o.ParentReference, true
}

// HasParentReference returns a boolean if a field has been set.
func (o *InlineObject1197) HasParentReference() bool {
	if o != nil && o.ParentReference != nil {
		return true
	}

	return false
}

// SetParentReference gets a reference to the given AnyOfmicrosoftGraphItemReference and assigns it to the ParentReference field.
func (o *InlineObject1197) SetParentReference(v AnyOfmicrosoftGraphItemReference) {
	o.ParentReference = v
}

func (o InlineObject1197) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ParentReference != nil {
		toSerialize["parentReference"] = o.ParentReference
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1197 struct {
	value *InlineObject1197
	isSet bool
}

func (v NullableInlineObject1197) Get() *InlineObject1197 {
	return v.value
}

func (v *NullableInlineObject1197) Set(val *InlineObject1197) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1197) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1197) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1197(val *InlineObject1197) *NullableInlineObject1197 {
	return &NullableInlineObject1197{value: val, isSet: true}
}

func (v NullableInlineObject1197) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1197) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


