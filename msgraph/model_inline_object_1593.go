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

// InlineObject1593 struct for InlineObject1593
type InlineObject1593 struct {
	Name NullableString `json:"name,omitempty"`
	Reference AnyOfobject `json:"reference,omitempty"`
	Comment NullableString `json:"comment,omitempty"`
}

// NewInlineObject1593 instantiates a new InlineObject1593 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1593() *InlineObject1593 {
	this := InlineObject1593{}
	return &this
}

// NewInlineObject1593WithDefaults instantiates a new InlineObject1593 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1593WithDefaults() *InlineObject1593 {
	this := InlineObject1593{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1593) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1593) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject1593) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *InlineObject1593) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *InlineObject1593) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *InlineObject1593) UnsetName() {
	o.Name.Unset()
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1593) GetReference() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1593) GetReferenceOk() (*AnyOfobject, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *InlineObject1593) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given AnyOfobject and assigns it to the Reference field.
func (o *InlineObject1593) SetReference(v AnyOfobject) {
	o.Reference = v
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1593) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1593) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *InlineObject1593) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *InlineObject1593) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *InlineObject1593) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *InlineObject1593) UnsetComment() {
	o.Comment.Unset()
}

func (o InlineObject1593) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1593 struct {
	value *InlineObject1593
	isSet bool
}

func (v NullableInlineObject1593) Get() *InlineObject1593 {
	return v.value
}

func (v *NullableInlineObject1593) Set(val *InlineObject1593) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1593) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1593) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1593(val *InlineObject1593) *NullableInlineObject1593 {
	return &NullableInlineObject1593{value: val, isSet: true}
}

func (v NullableInlineObject1593) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1593) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


