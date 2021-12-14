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

// ExternalItem struct for ExternalItem
type ExternalItem struct {
	// An array of access control entries. Each entry specifies the access granted to a user or group. Required.
	Acl *[]*AnyOfmicrosoftGraphExternalConnectorsAcl `json:"acl,omitempty"`
	// A plain-text  representation of the contents of the item. The text in this property is full-text indexed. Optional.
	Content AnyOfmicrosoftGraphExternalConnectorsExternalItemContent `json:"content,omitempty"`
	// A property bag with the properties of the item. The properties MUST conform to the schema defined for the externalConnection. Required.
	Properties AnyOfobject `json:"properties,omitempty"`
}

// NewExternalItem instantiates a new ExternalItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalItem() *ExternalItem {
	this := ExternalItem{}
	return &this
}

// NewExternalItemWithDefaults instantiates a new ExternalItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalItemWithDefaults() *ExternalItem {
	this := ExternalItem{}
	return &this
}

// GetAcl returns the Acl field value if set, zero value otherwise.
func (o *ExternalItem) GetAcl() []*AnyOfmicrosoftGraphExternalConnectorsAcl {
	if o == nil || o.Acl == nil {
		var ret []*AnyOfmicrosoftGraphExternalConnectorsAcl
		return ret
	}
	return *o.Acl
}

// GetAclOk returns a tuple with the Acl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalItem) GetAclOk() (*[]*AnyOfmicrosoftGraphExternalConnectorsAcl, bool) {
	if o == nil || o.Acl == nil {
		return nil, false
	}
	return o.Acl, true
}

// HasAcl returns a boolean if a field has been set.
func (o *ExternalItem) HasAcl() bool {
	if o != nil && o.Acl != nil {
		return true
	}

	return false
}

// SetAcl gets a reference to the given []*AnyOfmicrosoftGraphExternalConnectorsAcl and assigns it to the Acl field.
func (o *ExternalItem) SetAcl(v []*AnyOfmicrosoftGraphExternalConnectorsAcl) {
	o.Acl = &v
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalItem) GetContent() AnyOfmicrosoftGraphExternalConnectorsExternalItemContent {
	if o == nil  {
		var ret AnyOfmicrosoftGraphExternalConnectorsExternalItemContent
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalItem) GetContentOk() (*AnyOfmicrosoftGraphExternalConnectorsExternalItemContent, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return &o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ExternalItem) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given AnyOfmicrosoftGraphExternalConnectorsExternalItemContent and assigns it to the Content field.
func (o *ExternalItem) SetContent(v AnyOfmicrosoftGraphExternalConnectorsExternalItemContent) {
	o.Content = v
}

// GetProperties returns the Properties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalItem) GetProperties() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalItem) GetPropertiesOk() (*AnyOfobject, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return &o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ExternalItem) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given AnyOfobject and assigns it to the Properties field.
func (o *ExternalItem) SetProperties(v AnyOfobject) {
	o.Properties = v
}

func (o ExternalItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Acl != nil {
		toSerialize["acl"] = o.Acl
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableExternalItem struct {
	value *ExternalItem
	isSet bool
}

func (v NullableExternalItem) Get() *ExternalItem {
	return v.value
}

func (v *NullableExternalItem) Set(val *ExternalItem) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalItem) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalItem(val *ExternalItem) *NullableExternalItem {
	return &NullableExternalItem{value: val, isSet: true}
}

func (v NullableExternalItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


