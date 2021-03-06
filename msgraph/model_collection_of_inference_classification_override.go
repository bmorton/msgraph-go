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

// CollectionOfInferenceClassificationOverride struct for CollectionOfInferenceClassificationOverride
type CollectionOfInferenceClassificationOverride struct {
	Value *[]MicrosoftGraphInferenceClassificationOverride `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfInferenceClassificationOverride instantiates a new CollectionOfInferenceClassificationOverride object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfInferenceClassificationOverride() *CollectionOfInferenceClassificationOverride {
	this := CollectionOfInferenceClassificationOverride{}
	return &this
}

// NewCollectionOfInferenceClassificationOverrideWithDefaults instantiates a new CollectionOfInferenceClassificationOverride object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfInferenceClassificationOverrideWithDefaults() *CollectionOfInferenceClassificationOverride {
	this := CollectionOfInferenceClassificationOverride{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfInferenceClassificationOverride) GetValue() []MicrosoftGraphInferenceClassificationOverride {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphInferenceClassificationOverride
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfInferenceClassificationOverride) GetValueOk() (*[]MicrosoftGraphInferenceClassificationOverride, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfInferenceClassificationOverride) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphInferenceClassificationOverride and assigns it to the Value field.
func (o *CollectionOfInferenceClassificationOverride) SetValue(v []MicrosoftGraphInferenceClassificationOverride) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfInferenceClassificationOverride) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfInferenceClassificationOverride) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfInferenceClassificationOverride) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfInferenceClassificationOverride) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfInferenceClassificationOverride) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfInferenceClassificationOverride struct {
	value *CollectionOfInferenceClassificationOverride
	isSet bool
}

func (v NullableCollectionOfInferenceClassificationOverride) Get() *CollectionOfInferenceClassificationOverride {
	return v.value
}

func (v *NullableCollectionOfInferenceClassificationOverride) Set(val *CollectionOfInferenceClassificationOverride) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfInferenceClassificationOverride) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfInferenceClassificationOverride) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfInferenceClassificationOverride(val *CollectionOfInferenceClassificationOverride) *NullableCollectionOfInferenceClassificationOverride {
	return &NullableCollectionOfInferenceClassificationOverride{value: val, isSet: true}
}

func (v NullableCollectionOfInferenceClassificationOverride) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfInferenceClassificationOverride) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


