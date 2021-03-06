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

// InlineObject1441 struct for InlineObject1441
type InlineObject1441 struct {
	Probability AnyOfobject `json:"probability,omitempty"`
	Mean AnyOfobject `json:"mean,omitempty"`
	StandardDev AnyOfobject `json:"standardDev,omitempty"`
}

// NewInlineObject1441 instantiates a new InlineObject1441 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1441() *InlineObject1441 {
	this := InlineObject1441{}
	return &this
}

// NewInlineObject1441WithDefaults instantiates a new InlineObject1441 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1441WithDefaults() *InlineObject1441 {
	this := InlineObject1441{}
	return &this
}

// GetProbability returns the Probability field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1441) GetProbability() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Probability
}

// GetProbabilityOk returns a tuple with the Probability field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1441) GetProbabilityOk() (*AnyOfobject, bool) {
	if o == nil || o.Probability == nil {
		return nil, false
	}
	return &o.Probability, true
}

// HasProbability returns a boolean if a field has been set.
func (o *InlineObject1441) HasProbability() bool {
	if o != nil && o.Probability != nil {
		return true
	}

	return false
}

// SetProbability gets a reference to the given AnyOfobject and assigns it to the Probability field.
func (o *InlineObject1441) SetProbability(v AnyOfobject) {
	o.Probability = v
}

// GetMean returns the Mean field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1441) GetMean() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Mean
}

// GetMeanOk returns a tuple with the Mean field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1441) GetMeanOk() (*AnyOfobject, bool) {
	if o == nil || o.Mean == nil {
		return nil, false
	}
	return &o.Mean, true
}

// HasMean returns a boolean if a field has been set.
func (o *InlineObject1441) HasMean() bool {
	if o != nil && o.Mean != nil {
		return true
	}

	return false
}

// SetMean gets a reference to the given AnyOfobject and assigns it to the Mean field.
func (o *InlineObject1441) SetMean(v AnyOfobject) {
	o.Mean = v
}

// GetStandardDev returns the StandardDev field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1441) GetStandardDev() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.StandardDev
}

// GetStandardDevOk returns a tuple with the StandardDev field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1441) GetStandardDevOk() (*AnyOfobject, bool) {
	if o == nil || o.StandardDev == nil {
		return nil, false
	}
	return &o.StandardDev, true
}

// HasStandardDev returns a boolean if a field has been set.
func (o *InlineObject1441) HasStandardDev() bool {
	if o != nil && o.StandardDev != nil {
		return true
	}

	return false
}

// SetStandardDev gets a reference to the given AnyOfobject and assigns it to the StandardDev field.
func (o *InlineObject1441) SetStandardDev(v AnyOfobject) {
	o.StandardDev = v
}

func (o InlineObject1441) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Probability != nil {
		toSerialize["probability"] = o.Probability
	}
	if o.Mean != nil {
		toSerialize["mean"] = o.Mean
	}
	if o.StandardDev != nil {
		toSerialize["standardDev"] = o.StandardDev
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1441 struct {
	value *InlineObject1441
	isSet bool
}

func (v NullableInlineObject1441) Get() *InlineObject1441 {
	return v.value
}

func (v *NullableInlineObject1441) Set(val *InlineObject1441) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1441) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1441) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1441(val *InlineObject1441) *NullableInlineObject1441 {
	return &NullableInlineObject1441{value: val, isSet: true}
}

func (v NullableInlineObject1441) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1441) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


