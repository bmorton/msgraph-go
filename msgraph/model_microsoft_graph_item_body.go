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

// MicrosoftGraphItemBody struct for MicrosoftGraphItemBody
type MicrosoftGraphItemBody struct {
	// The content of the item.
	Content NullableString `json:"content,omitempty"`
	// The type of the content. Possible values are text and html.
	ContentType AnyOfmicrosoftGraphBodyType `json:"contentType,omitempty"`
}

// NewMicrosoftGraphItemBody instantiates a new MicrosoftGraphItemBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphItemBody() *MicrosoftGraphItemBody {
	this := MicrosoftGraphItemBody{}
	return &this
}

// NewMicrosoftGraphItemBodyWithDefaults instantiates a new MicrosoftGraphItemBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphItemBodyWithDefaults() *MicrosoftGraphItemBody {
	this := MicrosoftGraphItemBody{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphItemBody) GetContent() string {
	if o == nil || o.Content.Get() == nil {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphItemBody) GetContentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *MicrosoftGraphItemBody) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *MicrosoftGraphItemBody) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *MicrosoftGraphItemBody) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *MicrosoftGraphItemBody) UnsetContent() {
	o.Content.Unset()
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphItemBody) GetContentType() AnyOfmicrosoftGraphBodyType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphBodyType
		return ret
	}
	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphItemBody) GetContentTypeOk() (*AnyOfmicrosoftGraphBodyType, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return &o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *MicrosoftGraphItemBody) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given AnyOfmicrosoftGraphBodyType and assigns it to the ContentType field.
func (o *MicrosoftGraphItemBody) SetContentType(v AnyOfmicrosoftGraphBodyType) {
	o.ContentType = v
}

func (o MicrosoftGraphItemBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Content.IsSet() {
		toSerialize["content"] = o.Content.Get()
	}
	if o.ContentType != nil {
		toSerialize["contentType"] = o.ContentType
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphItemBody struct {
	value *MicrosoftGraphItemBody
	isSet bool
}

func (v NullableMicrosoftGraphItemBody) Get() *MicrosoftGraphItemBody {
	return v.value
}

func (v *NullableMicrosoftGraphItemBody) Set(val *MicrosoftGraphItemBody) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphItemBody) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphItemBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphItemBody(val *MicrosoftGraphItemBody) *NullableMicrosoftGraphItemBody {
	return &NullableMicrosoftGraphItemBody{value: val, isSet: true}
}

func (v NullableMicrosoftGraphItemBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphItemBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


