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

// InlineObject739 struct for InlineObject739
type InlineObject739 struct {
	HubSiteUrls *[]string `json:"hubSiteUrls,omitempty"`
	PropagateToExistingLists NullableBool `json:"propagateToExistingLists,omitempty"`
}

// NewInlineObject739 instantiates a new InlineObject739 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject739() *InlineObject739 {
	this := InlineObject739{}
	var propagateToExistingLists bool = false
	this.PropagateToExistingLists = *NewNullableBool(&propagateToExistingLists)
	return &this
}

// NewInlineObject739WithDefaults instantiates a new InlineObject739 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject739WithDefaults() *InlineObject739 {
	this := InlineObject739{}
	var propagateToExistingLists bool = false
	this.PropagateToExistingLists = *NewNullableBool(&propagateToExistingLists)
	return &this
}

// GetHubSiteUrls returns the HubSiteUrls field value if set, zero value otherwise.
func (o *InlineObject739) GetHubSiteUrls() []string {
	if o == nil || o.HubSiteUrls == nil {
		var ret []string
		return ret
	}
	return *o.HubSiteUrls
}

// GetHubSiteUrlsOk returns a tuple with the HubSiteUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject739) GetHubSiteUrlsOk() (*[]string, bool) {
	if o == nil || o.HubSiteUrls == nil {
		return nil, false
	}
	return o.HubSiteUrls, true
}

// HasHubSiteUrls returns a boolean if a field has been set.
func (o *InlineObject739) HasHubSiteUrls() bool {
	if o != nil && o.HubSiteUrls != nil {
		return true
	}

	return false
}

// SetHubSiteUrls gets a reference to the given []string and assigns it to the HubSiteUrls field.
func (o *InlineObject739) SetHubSiteUrls(v []string) {
	o.HubSiteUrls = &v
}

// GetPropagateToExistingLists returns the PropagateToExistingLists field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject739) GetPropagateToExistingLists() bool {
	if o == nil || o.PropagateToExistingLists.Get() == nil {
		var ret bool
		return ret
	}
	return *o.PropagateToExistingLists.Get()
}

// GetPropagateToExistingListsOk returns a tuple with the PropagateToExistingLists field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject739) GetPropagateToExistingListsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PropagateToExistingLists.Get(), o.PropagateToExistingLists.IsSet()
}

// HasPropagateToExistingLists returns a boolean if a field has been set.
func (o *InlineObject739) HasPropagateToExistingLists() bool {
	if o != nil && o.PropagateToExistingLists.IsSet() {
		return true
	}

	return false
}

// SetPropagateToExistingLists gets a reference to the given NullableBool and assigns it to the PropagateToExistingLists field.
func (o *InlineObject739) SetPropagateToExistingLists(v bool) {
	o.PropagateToExistingLists.Set(&v)
}
// SetPropagateToExistingListsNil sets the value for PropagateToExistingLists to be an explicit nil
func (o *InlineObject739) SetPropagateToExistingListsNil() {
	o.PropagateToExistingLists.Set(nil)
}

// UnsetPropagateToExistingLists ensures that no value is present for PropagateToExistingLists, not even an explicit nil
func (o *InlineObject739) UnsetPropagateToExistingLists() {
	o.PropagateToExistingLists.Unset()
}

func (o InlineObject739) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HubSiteUrls != nil {
		toSerialize["hubSiteUrls"] = o.HubSiteUrls
	}
	if o.PropagateToExistingLists.IsSet() {
		toSerialize["propagateToExistingLists"] = o.PropagateToExistingLists.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject739 struct {
	value *InlineObject739
	isSet bool
}

func (v NullableInlineObject739) Get() *InlineObject739 {
	return v.value
}

func (v *NullableInlineObject739) Set(val *InlineObject739) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject739) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject739) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject739(val *InlineObject739) *NullableInlineObject739 {
	return &NullableInlineObject739{value: val, isSet: true}
}

func (v NullableInlineObject739) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject739) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


