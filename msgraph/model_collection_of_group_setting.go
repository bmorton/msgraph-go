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

// CollectionOfGroupSetting struct for CollectionOfGroupSetting
type CollectionOfGroupSetting struct {
	Value *[]MicrosoftGraphGroupSetting `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfGroupSetting instantiates a new CollectionOfGroupSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfGroupSetting() *CollectionOfGroupSetting {
	this := CollectionOfGroupSetting{}
	return &this
}

// NewCollectionOfGroupSettingWithDefaults instantiates a new CollectionOfGroupSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfGroupSettingWithDefaults() *CollectionOfGroupSetting {
	this := CollectionOfGroupSetting{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfGroupSetting) GetValue() []MicrosoftGraphGroupSetting {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphGroupSetting
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfGroupSetting) GetValueOk() (*[]MicrosoftGraphGroupSetting, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfGroupSetting) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphGroupSetting and assigns it to the Value field.
func (o *CollectionOfGroupSetting) SetValue(v []MicrosoftGraphGroupSetting) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfGroupSetting) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfGroupSetting) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfGroupSetting) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfGroupSetting) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfGroupSetting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfGroupSetting struct {
	value *CollectionOfGroupSetting
	isSet bool
}

func (v NullableCollectionOfGroupSetting) Get() *CollectionOfGroupSetting {
	return v.value
}

func (v *NullableCollectionOfGroupSetting) Set(val *CollectionOfGroupSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfGroupSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfGroupSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfGroupSetting(val *CollectionOfGroupSetting) *NullableCollectionOfGroupSetting {
	return &NullableCollectionOfGroupSetting{value: val, isSet: true}
}

func (v NullableCollectionOfGroupSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfGroupSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


