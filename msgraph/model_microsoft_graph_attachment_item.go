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

// MicrosoftGraphAttachmentItem struct for MicrosoftGraphAttachmentItem
type MicrosoftGraphAttachmentItem struct {
	// The type of attachment. Possible values are: file, item, reference. Required.
	AttachmentType AnyOfmicrosoftGraphAttachmentType `json:"attachmentType,omitempty"`
	// The nature of the data in the attachment. Optional.
	ContentType NullableString `json:"contentType,omitempty"`
	// true if the attachment is an inline attachment; otherwise, false. Optional.
	IsInline NullableBool `json:"isInline,omitempty"`
	// The display name of the attachment. This can be a descriptive string and does not have to be the actual file name. Required.
	Name NullableString `json:"name,omitempty"`
	// The length of the attachment in bytes. Required.
	Size NullableInt64 `json:"size,omitempty"`
}

// NewMicrosoftGraphAttachmentItem instantiates a new MicrosoftGraphAttachmentItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphAttachmentItem() *MicrosoftGraphAttachmentItem {
	this := MicrosoftGraphAttachmentItem{}
	return &this
}

// NewMicrosoftGraphAttachmentItemWithDefaults instantiates a new MicrosoftGraphAttachmentItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphAttachmentItemWithDefaults() *MicrosoftGraphAttachmentItem {
	this := MicrosoftGraphAttachmentItem{}
	return &this
}

// GetAttachmentType returns the AttachmentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAttachmentItem) GetAttachmentType() AnyOfmicrosoftGraphAttachmentType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphAttachmentType
		return ret
	}
	return o.AttachmentType
}

// GetAttachmentTypeOk returns a tuple with the AttachmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAttachmentItem) GetAttachmentTypeOk() (*AnyOfmicrosoftGraphAttachmentType, bool) {
	if o == nil || o.AttachmentType == nil {
		return nil, false
	}
	return &o.AttachmentType, true
}

// HasAttachmentType returns a boolean if a field has been set.
func (o *MicrosoftGraphAttachmentItem) HasAttachmentType() bool {
	if o != nil && o.AttachmentType != nil {
		return true
	}

	return false
}

// SetAttachmentType gets a reference to the given AnyOfmicrosoftGraphAttachmentType and assigns it to the AttachmentType field.
func (o *MicrosoftGraphAttachmentItem) SetAttachmentType(v AnyOfmicrosoftGraphAttachmentType) {
	o.AttachmentType = v
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAttachmentItem) GetContentType() string {
	if o == nil || o.ContentType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAttachmentItem) GetContentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *MicrosoftGraphAttachmentItem) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *MicrosoftGraphAttachmentItem) SetContentType(v string) {
	o.ContentType.Set(&v)
}
// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *MicrosoftGraphAttachmentItem) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *MicrosoftGraphAttachmentItem) UnsetContentType() {
	o.ContentType.Unset()
}

// GetIsInline returns the IsInline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAttachmentItem) GetIsInline() bool {
	if o == nil || o.IsInline.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsInline.Get()
}

// GetIsInlineOk returns a tuple with the IsInline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAttachmentItem) GetIsInlineOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsInline.Get(), o.IsInline.IsSet()
}

// HasIsInline returns a boolean if a field has been set.
func (o *MicrosoftGraphAttachmentItem) HasIsInline() bool {
	if o != nil && o.IsInline.IsSet() {
		return true
	}

	return false
}

// SetIsInline gets a reference to the given NullableBool and assigns it to the IsInline field.
func (o *MicrosoftGraphAttachmentItem) SetIsInline(v bool) {
	o.IsInline.Set(&v)
}
// SetIsInlineNil sets the value for IsInline to be an explicit nil
func (o *MicrosoftGraphAttachmentItem) SetIsInlineNil() {
	o.IsInline.Set(nil)
}

// UnsetIsInline ensures that no value is present for IsInline, not even an explicit nil
func (o *MicrosoftGraphAttachmentItem) UnsetIsInline() {
	o.IsInline.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAttachmentItem) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAttachmentItem) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftGraphAttachmentItem) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MicrosoftGraphAttachmentItem) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MicrosoftGraphAttachmentItem) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MicrosoftGraphAttachmentItem) UnsetName() {
	o.Name.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAttachmentItem) GetSize() int64 {
	if o == nil || o.Size.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAttachmentItem) GetSizeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field has been set.
func (o *MicrosoftGraphAttachmentItem) HasSize() bool {
	if o != nil && o.Size.IsSet() {
		return true
	}

	return false
}

// SetSize gets a reference to the given NullableInt64 and assigns it to the Size field.
func (o *MicrosoftGraphAttachmentItem) SetSize(v int64) {
	o.Size.Set(&v)
}
// SetSizeNil sets the value for Size to be an explicit nil
func (o *MicrosoftGraphAttachmentItem) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *MicrosoftGraphAttachmentItem) UnsetSize() {
	o.Size.Unset()
}

func (o MicrosoftGraphAttachmentItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AttachmentType != nil {
		toSerialize["attachmentType"] = o.AttachmentType
	}
	if o.ContentType.IsSet() {
		toSerialize["contentType"] = o.ContentType.Get()
	}
	if o.IsInline.IsSet() {
		toSerialize["isInline"] = o.IsInline.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Size.IsSet() {
		toSerialize["size"] = o.Size.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphAttachmentItem struct {
	value *MicrosoftGraphAttachmentItem
	isSet bool
}

func (v NullableMicrosoftGraphAttachmentItem) Get() *MicrosoftGraphAttachmentItem {
	return v.value
}

func (v *NullableMicrosoftGraphAttachmentItem) Set(val *MicrosoftGraphAttachmentItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAttachmentItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAttachmentItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAttachmentItem(val *MicrosoftGraphAttachmentItem) *NullableMicrosoftGraphAttachmentItem {
	return &NullableMicrosoftGraphAttachmentItem{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAttachmentItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAttachmentItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


