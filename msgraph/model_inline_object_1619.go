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

// InlineObject1619 struct for InlineObject1619
type InlineObject1619 struct {
	Name NullableString `json:"name,omitempty"`
	Formula NullableString `json:"formula,omitempty"`
	Comment NullableString `json:"comment,omitempty"`
}

// NewInlineObject1619 instantiates a new InlineObject1619 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1619() *InlineObject1619 {
	this := InlineObject1619{}
	return &this
}

// NewInlineObject1619WithDefaults instantiates a new InlineObject1619 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1619WithDefaults() *InlineObject1619 {
	this := InlineObject1619{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1619) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1619) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject1619) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *InlineObject1619) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *InlineObject1619) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *InlineObject1619) UnsetName() {
	o.Name.Unset()
}

// GetFormula returns the Formula field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1619) GetFormula() string {
	if o == nil || o.Formula.Get() == nil {
		var ret string
		return ret
	}
	return *o.Formula.Get()
}

// GetFormulaOk returns a tuple with the Formula field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1619) GetFormulaOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Formula.Get(), o.Formula.IsSet()
}

// HasFormula returns a boolean if a field has been set.
func (o *InlineObject1619) HasFormula() bool {
	if o != nil && o.Formula.IsSet() {
		return true
	}

	return false
}

// SetFormula gets a reference to the given NullableString and assigns it to the Formula field.
func (o *InlineObject1619) SetFormula(v string) {
	o.Formula.Set(&v)
}
// SetFormulaNil sets the value for Formula to be an explicit nil
func (o *InlineObject1619) SetFormulaNil() {
	o.Formula.Set(nil)
}

// UnsetFormula ensures that no value is present for Formula, not even an explicit nil
func (o *InlineObject1619) UnsetFormula() {
	o.Formula.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1619) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1619) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *InlineObject1619) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *InlineObject1619) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *InlineObject1619) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *InlineObject1619) UnsetComment() {
	o.Comment.Unset()
}

func (o InlineObject1619) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Formula.IsSet() {
		toSerialize["formula"] = o.Formula.Get()
	}
	if o.Comment.IsSet() {
		toSerialize["comment"] = o.Comment.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1619 struct {
	value *InlineObject1619
	isSet bool
}

func (v NullableInlineObject1619) Get() *InlineObject1619 {
	return v.value
}

func (v *NullableInlineObject1619) Set(val *InlineObject1619) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1619) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1619) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1619(val *InlineObject1619) *NullableInlineObject1619 {
	return &NullableInlineObject1619{value: val, isSet: true}
}

func (v NullableInlineObject1619) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1619) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


