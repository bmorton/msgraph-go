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

// InlineObject1617 struct for InlineObject1617
type InlineObject1617 struct {
	Type *string `json:"type,omitempty"`
	SourceData AnyOfobject `json:"sourceData,omitempty"`
	SeriesBy *string `json:"seriesBy,omitempty"`
}

// NewInlineObject1617 instantiates a new InlineObject1617 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1617() *InlineObject1617 {
	this := InlineObject1617{}
	return &this
}

// NewInlineObject1617WithDefaults instantiates a new InlineObject1617 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1617WithDefaults() *InlineObject1617 {
	this := InlineObject1617{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineObject1617) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1617) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineObject1617) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineObject1617) SetType(v string) {
	o.Type = &v
}

// GetSourceData returns the SourceData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1617) GetSourceData() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.SourceData
}

// GetSourceDataOk returns a tuple with the SourceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1617) GetSourceDataOk() (*AnyOfobject, bool) {
	if o == nil || o.SourceData == nil {
		return nil, false
	}
	return &o.SourceData, true
}

// HasSourceData returns a boolean if a field has been set.
func (o *InlineObject1617) HasSourceData() bool {
	if o != nil && o.SourceData != nil {
		return true
	}

	return false
}

// SetSourceData gets a reference to the given AnyOfobject and assigns it to the SourceData field.
func (o *InlineObject1617) SetSourceData(v AnyOfobject) {
	o.SourceData = v
}

// GetSeriesBy returns the SeriesBy field value if set, zero value otherwise.
func (o *InlineObject1617) GetSeriesBy() string {
	if o == nil || o.SeriesBy == nil {
		var ret string
		return ret
	}
	return *o.SeriesBy
}

// GetSeriesByOk returns a tuple with the SeriesBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1617) GetSeriesByOk() (*string, bool) {
	if o == nil || o.SeriesBy == nil {
		return nil, false
	}
	return o.SeriesBy, true
}

// HasSeriesBy returns a boolean if a field has been set.
func (o *InlineObject1617) HasSeriesBy() bool {
	if o != nil && o.SeriesBy != nil {
		return true
	}

	return false
}

// SetSeriesBy gets a reference to the given string and assigns it to the SeriesBy field.
func (o *InlineObject1617) SetSeriesBy(v string) {
	o.SeriesBy = &v
}

func (o InlineObject1617) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.SourceData != nil {
		toSerialize["sourceData"] = o.SourceData
	}
	if o.SeriesBy != nil {
		toSerialize["seriesBy"] = o.SeriesBy
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1617 struct {
	value *InlineObject1617
	isSet bool
}

func (v NullableInlineObject1617) Get() *InlineObject1617 {
	return v.value
}

func (v *NullableInlineObject1617) Set(val *InlineObject1617) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1617) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1617) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1617(val *InlineObject1617) *NullableInlineObject1617 {
	return &NullableInlineObject1617{value: val, isSet: true}
}

func (v NullableInlineObject1617) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1617) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

