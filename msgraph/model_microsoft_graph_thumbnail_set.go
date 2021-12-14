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

// MicrosoftGraphThumbnailSet struct for MicrosoftGraphThumbnailSet
type MicrosoftGraphThumbnailSet struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// A 1920x1920 scaled thumbnail.
	Large AnyOfmicrosoftGraphThumbnail `json:"large,omitempty"`
	// A 176x176 scaled thumbnail.
	Medium AnyOfmicrosoftGraphThumbnail `json:"medium,omitempty"`
	// A 48x48 cropped thumbnail.
	Small AnyOfmicrosoftGraphThumbnail `json:"small,omitempty"`
	// A custom thumbnail image or the original image used to generate other thumbnails.
	Source AnyOfmicrosoftGraphThumbnail `json:"source,omitempty"`
}

// NewMicrosoftGraphThumbnailSet instantiates a new MicrosoftGraphThumbnailSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphThumbnailSet() *MicrosoftGraphThumbnailSet {
	this := MicrosoftGraphThumbnailSet{}
	return &this
}

// NewMicrosoftGraphThumbnailSetWithDefaults instantiates a new MicrosoftGraphThumbnailSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphThumbnailSetWithDefaults() *MicrosoftGraphThumbnailSet {
	this := MicrosoftGraphThumbnailSet{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphThumbnailSet) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphThumbnailSet) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphThumbnailSet) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphThumbnailSet) SetId(v string) {
	o.Id = &v
}

// GetLarge returns the Large field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphThumbnailSet) GetLarge() AnyOfmicrosoftGraphThumbnail {
	if o == nil  {
		var ret AnyOfmicrosoftGraphThumbnail
		return ret
	}
	return o.Large
}

// GetLargeOk returns a tuple with the Large field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphThumbnailSet) GetLargeOk() (*AnyOfmicrosoftGraphThumbnail, bool) {
	if o == nil || o.Large == nil {
		return nil, false
	}
	return &o.Large, true
}

// HasLarge returns a boolean if a field has been set.
func (o *MicrosoftGraphThumbnailSet) HasLarge() bool {
	if o != nil && o.Large != nil {
		return true
	}

	return false
}

// SetLarge gets a reference to the given AnyOfmicrosoftGraphThumbnail and assigns it to the Large field.
func (o *MicrosoftGraphThumbnailSet) SetLarge(v AnyOfmicrosoftGraphThumbnail) {
	o.Large = v
}

// GetMedium returns the Medium field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphThumbnailSet) GetMedium() AnyOfmicrosoftGraphThumbnail {
	if o == nil  {
		var ret AnyOfmicrosoftGraphThumbnail
		return ret
	}
	return o.Medium
}

// GetMediumOk returns a tuple with the Medium field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphThumbnailSet) GetMediumOk() (*AnyOfmicrosoftGraphThumbnail, bool) {
	if o == nil || o.Medium == nil {
		return nil, false
	}
	return &o.Medium, true
}

// HasMedium returns a boolean if a field has been set.
func (o *MicrosoftGraphThumbnailSet) HasMedium() bool {
	if o != nil && o.Medium != nil {
		return true
	}

	return false
}

// SetMedium gets a reference to the given AnyOfmicrosoftGraphThumbnail and assigns it to the Medium field.
func (o *MicrosoftGraphThumbnailSet) SetMedium(v AnyOfmicrosoftGraphThumbnail) {
	o.Medium = v
}

// GetSmall returns the Small field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphThumbnailSet) GetSmall() AnyOfmicrosoftGraphThumbnail {
	if o == nil  {
		var ret AnyOfmicrosoftGraphThumbnail
		return ret
	}
	return o.Small
}

// GetSmallOk returns a tuple with the Small field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphThumbnailSet) GetSmallOk() (*AnyOfmicrosoftGraphThumbnail, bool) {
	if o == nil || o.Small == nil {
		return nil, false
	}
	return &o.Small, true
}

// HasSmall returns a boolean if a field has been set.
func (o *MicrosoftGraphThumbnailSet) HasSmall() bool {
	if o != nil && o.Small != nil {
		return true
	}

	return false
}

// SetSmall gets a reference to the given AnyOfmicrosoftGraphThumbnail and assigns it to the Small field.
func (o *MicrosoftGraphThumbnailSet) SetSmall(v AnyOfmicrosoftGraphThumbnail) {
	o.Small = v
}

// GetSource returns the Source field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphThumbnailSet) GetSource() AnyOfmicrosoftGraphThumbnail {
	if o == nil  {
		var ret AnyOfmicrosoftGraphThumbnail
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphThumbnailSet) GetSourceOk() (*AnyOfmicrosoftGraphThumbnail, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return &o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *MicrosoftGraphThumbnailSet) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given AnyOfmicrosoftGraphThumbnail and assigns it to the Source field.
func (o *MicrosoftGraphThumbnailSet) SetSource(v AnyOfmicrosoftGraphThumbnail) {
	o.Source = v
}

func (o MicrosoftGraphThumbnailSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Large != nil {
		toSerialize["large"] = o.Large
	}
	if o.Medium != nil {
		toSerialize["medium"] = o.Medium
	}
	if o.Small != nil {
		toSerialize["small"] = o.Small
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphThumbnailSet struct {
	value *MicrosoftGraphThumbnailSet
	isSet bool
}

func (v NullableMicrosoftGraphThumbnailSet) Get() *MicrosoftGraphThumbnailSet {
	return v.value
}

func (v *NullableMicrosoftGraphThumbnailSet) Set(val *MicrosoftGraphThumbnailSet) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphThumbnailSet) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphThumbnailSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphThumbnailSet(val *MicrosoftGraphThumbnailSet) *NullableMicrosoftGraphThumbnailSet {
	return &NullableMicrosoftGraphThumbnailSet{value: val, isSet: true}
}

func (v NullableMicrosoftGraphThumbnailSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphThumbnailSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

