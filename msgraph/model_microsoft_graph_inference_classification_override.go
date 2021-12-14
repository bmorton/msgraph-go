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

// MicrosoftGraphInferenceClassificationOverride struct for MicrosoftGraphInferenceClassificationOverride
type MicrosoftGraphInferenceClassificationOverride struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Specifies how incoming messages from a specific sender should always be classified as. The possible values are: focused, other.
	ClassifyAs AnyOfmicrosoftGraphInferenceClassificationType `json:"classifyAs,omitempty"`
	// The email address information of the sender for whom the override is created.
	SenderEmailAddress AnyOfmicrosoftGraphEmailAddress `json:"senderEmailAddress,omitempty"`
}

// NewMicrosoftGraphInferenceClassificationOverride instantiates a new MicrosoftGraphInferenceClassificationOverride object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphInferenceClassificationOverride() *MicrosoftGraphInferenceClassificationOverride {
	this := MicrosoftGraphInferenceClassificationOverride{}
	return &this
}

// NewMicrosoftGraphInferenceClassificationOverrideWithDefaults instantiates a new MicrosoftGraphInferenceClassificationOverride object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphInferenceClassificationOverrideWithDefaults() *MicrosoftGraphInferenceClassificationOverride {
	this := MicrosoftGraphInferenceClassificationOverride{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphInferenceClassificationOverride) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphInferenceClassificationOverride) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphInferenceClassificationOverride) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphInferenceClassificationOverride) SetId(v string) {
	o.Id = &v
}

// GetClassifyAs returns the ClassifyAs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphInferenceClassificationOverride) GetClassifyAs() AnyOfmicrosoftGraphInferenceClassificationType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphInferenceClassificationType
		return ret
	}
	return o.ClassifyAs
}

// GetClassifyAsOk returns a tuple with the ClassifyAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphInferenceClassificationOverride) GetClassifyAsOk() (*AnyOfmicrosoftGraphInferenceClassificationType, bool) {
	if o == nil || o.ClassifyAs == nil {
		return nil, false
	}
	return &o.ClassifyAs, true
}

// HasClassifyAs returns a boolean if a field has been set.
func (o *MicrosoftGraphInferenceClassificationOverride) HasClassifyAs() bool {
	if o != nil && o.ClassifyAs != nil {
		return true
	}

	return false
}

// SetClassifyAs gets a reference to the given AnyOfmicrosoftGraphInferenceClassificationType and assigns it to the ClassifyAs field.
func (o *MicrosoftGraphInferenceClassificationOverride) SetClassifyAs(v AnyOfmicrosoftGraphInferenceClassificationType) {
	o.ClassifyAs = v
}

// GetSenderEmailAddress returns the SenderEmailAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphInferenceClassificationOverride) GetSenderEmailAddress() AnyOfmicrosoftGraphEmailAddress {
	if o == nil  {
		var ret AnyOfmicrosoftGraphEmailAddress
		return ret
	}
	return o.SenderEmailAddress
}

// GetSenderEmailAddressOk returns a tuple with the SenderEmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphInferenceClassificationOverride) GetSenderEmailAddressOk() (*AnyOfmicrosoftGraphEmailAddress, bool) {
	if o == nil || o.SenderEmailAddress == nil {
		return nil, false
	}
	return &o.SenderEmailAddress, true
}

// HasSenderEmailAddress returns a boolean if a field has been set.
func (o *MicrosoftGraphInferenceClassificationOverride) HasSenderEmailAddress() bool {
	if o != nil && o.SenderEmailAddress != nil {
		return true
	}

	return false
}

// SetSenderEmailAddress gets a reference to the given AnyOfmicrosoftGraphEmailAddress and assigns it to the SenderEmailAddress field.
func (o *MicrosoftGraphInferenceClassificationOverride) SetSenderEmailAddress(v AnyOfmicrosoftGraphEmailAddress) {
	o.SenderEmailAddress = v
}

func (o MicrosoftGraphInferenceClassificationOverride) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ClassifyAs != nil {
		toSerialize["classifyAs"] = o.ClassifyAs
	}
	if o.SenderEmailAddress != nil {
		toSerialize["senderEmailAddress"] = o.SenderEmailAddress
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphInferenceClassificationOverride struct {
	value *MicrosoftGraphInferenceClassificationOverride
	isSet bool
}

func (v NullableMicrosoftGraphInferenceClassificationOverride) Get() *MicrosoftGraphInferenceClassificationOverride {
	return v.value
}

func (v *NullableMicrosoftGraphInferenceClassificationOverride) Set(val *MicrosoftGraphInferenceClassificationOverride) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphInferenceClassificationOverride) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphInferenceClassificationOverride) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphInferenceClassificationOverride(val *MicrosoftGraphInferenceClassificationOverride) *NullableMicrosoftGraphInferenceClassificationOverride {
	return &NullableMicrosoftGraphInferenceClassificationOverride{value: val, isSet: true}
}

func (v NullableMicrosoftGraphInferenceClassificationOverride) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphInferenceClassificationOverride) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


