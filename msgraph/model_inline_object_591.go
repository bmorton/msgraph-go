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

// InlineObject591 struct for InlineObject591
type InlineObject591 struct {
	EmailAddresses *[]string `json:"EmailAddresses,omitempty"`
	MailTipsOptions AnyOfmicrosoftGraphMailTipsType `json:"MailTipsOptions,omitempty"`
}

// NewInlineObject591 instantiates a new InlineObject591 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject591() *InlineObject591 {
	this := InlineObject591{}
	return &this
}

// NewInlineObject591WithDefaults instantiates a new InlineObject591 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject591WithDefaults() *InlineObject591 {
	this := InlineObject591{}
	return &this
}

// GetEmailAddresses returns the EmailAddresses field value if set, zero value otherwise.
func (o *InlineObject591) GetEmailAddresses() []string {
	if o == nil || o.EmailAddresses == nil {
		var ret []string
		return ret
	}
	return *o.EmailAddresses
}

// GetEmailAddressesOk returns a tuple with the EmailAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject591) GetEmailAddressesOk() (*[]string, bool) {
	if o == nil || o.EmailAddresses == nil {
		return nil, false
	}
	return o.EmailAddresses, true
}

// HasEmailAddresses returns a boolean if a field has been set.
func (o *InlineObject591) HasEmailAddresses() bool {
	if o != nil && o.EmailAddresses != nil {
		return true
	}

	return false
}

// SetEmailAddresses gets a reference to the given []string and assigns it to the EmailAddresses field.
func (o *InlineObject591) SetEmailAddresses(v []string) {
	o.EmailAddresses = &v
}

// GetMailTipsOptions returns the MailTipsOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject591) GetMailTipsOptions() AnyOfmicrosoftGraphMailTipsType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphMailTipsType
		return ret
	}
	return o.MailTipsOptions
}

// GetMailTipsOptionsOk returns a tuple with the MailTipsOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject591) GetMailTipsOptionsOk() (*AnyOfmicrosoftGraphMailTipsType, bool) {
	if o == nil || o.MailTipsOptions == nil {
		return nil, false
	}
	return &o.MailTipsOptions, true
}

// HasMailTipsOptions returns a boolean if a field has been set.
func (o *InlineObject591) HasMailTipsOptions() bool {
	if o != nil && o.MailTipsOptions != nil {
		return true
	}

	return false
}

// SetMailTipsOptions gets a reference to the given AnyOfmicrosoftGraphMailTipsType and assigns it to the MailTipsOptions field.
func (o *InlineObject591) SetMailTipsOptions(v AnyOfmicrosoftGraphMailTipsType) {
	o.MailTipsOptions = v
}

func (o InlineObject591) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EmailAddresses != nil {
		toSerialize["EmailAddresses"] = o.EmailAddresses
	}
	if o.MailTipsOptions != nil {
		toSerialize["MailTipsOptions"] = o.MailTipsOptions
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject591 struct {
	value *InlineObject591
	isSet bool
}

func (v NullableInlineObject591) Get() *InlineObject591 {
	return v.value
}

func (v *NullableInlineObject591) Set(val *InlineObject591) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject591) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject591) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject591(val *InlineObject591) *NullableInlineObject591 {
	return &NullableInlineObject591{value: val, isSet: true}
}

func (v NullableInlineObject591) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject591) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


