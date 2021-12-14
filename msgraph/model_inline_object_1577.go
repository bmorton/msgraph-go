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

// InlineObject1577 struct for InlineObject1577
type InlineObject1577 struct {
	Options AnyOfmicrosoftGraphWorkbookWorksheetProtectionOptions `json:"options,omitempty"`
}

// NewInlineObject1577 instantiates a new InlineObject1577 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1577() *InlineObject1577 {
	this := InlineObject1577{}
	return &this
}

// NewInlineObject1577WithDefaults instantiates a new InlineObject1577 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1577WithDefaults() *InlineObject1577 {
	this := InlineObject1577{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1577) GetOptions() AnyOfmicrosoftGraphWorkbookWorksheetProtectionOptions {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookWorksheetProtectionOptions
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1577) GetOptionsOk() (*AnyOfmicrosoftGraphWorkbookWorksheetProtectionOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return &o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *InlineObject1577) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given AnyOfmicrosoftGraphWorkbookWorksheetProtectionOptions and assigns it to the Options field.
func (o *InlineObject1577) SetOptions(v AnyOfmicrosoftGraphWorkbookWorksheetProtectionOptions) {
	o.Options = v
}

func (o InlineObject1577) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1577 struct {
	value *InlineObject1577
	isSet bool
}

func (v NullableInlineObject1577) Get() *InlineObject1577 {
	return v.value
}

func (v *NullableInlineObject1577) Set(val *InlineObject1577) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1577) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1577) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1577(val *InlineObject1577) *NullableInlineObject1577 {
	return &NullableInlineObject1577{value: val, isSet: true}
}

func (v NullableInlineObject1577) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1577) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


