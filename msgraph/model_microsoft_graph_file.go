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

// MicrosoftGraphFile struct for MicrosoftGraphFile
type MicrosoftGraphFile struct {
	// Hashes of the file's binary content, if available. Read-only.
	Hashes AnyOfmicrosoftGraphHashes `json:"hashes,omitempty"`
	// The MIME type for the file. This is determined by logic on the server and might not be the value provided when the file was uploaded. Read-only.
	MimeType NullableString `json:"mimeType,omitempty"`
	ProcessingMetadata NullableBool `json:"processingMetadata,omitempty"`
}

// NewMicrosoftGraphFile instantiates a new MicrosoftGraphFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphFile() *MicrosoftGraphFile {
	this := MicrosoftGraphFile{}
	return &this
}

// NewMicrosoftGraphFileWithDefaults instantiates a new MicrosoftGraphFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphFileWithDefaults() *MicrosoftGraphFile {
	this := MicrosoftGraphFile{}
	return &this
}

// GetHashes returns the Hashes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphFile) GetHashes() AnyOfmicrosoftGraphHashes {
	if o == nil  {
		var ret AnyOfmicrosoftGraphHashes
		return ret
	}
	return o.Hashes
}

// GetHashesOk returns a tuple with the Hashes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphFile) GetHashesOk() (*AnyOfmicrosoftGraphHashes, bool) {
	if o == nil || o.Hashes == nil {
		return nil, false
	}
	return &o.Hashes, true
}

// HasHashes returns a boolean if a field has been set.
func (o *MicrosoftGraphFile) HasHashes() bool {
	if o != nil && o.Hashes != nil {
		return true
	}

	return false
}

// SetHashes gets a reference to the given AnyOfmicrosoftGraphHashes and assigns it to the Hashes field.
func (o *MicrosoftGraphFile) SetHashes(v AnyOfmicrosoftGraphHashes) {
	o.Hashes = v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphFile) GetMimeType() string {
	if o == nil || o.MimeType.Get() == nil {
		var ret string
		return ret
	}
	return *o.MimeType.Get()
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphFile) GetMimeTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MimeType.Get(), o.MimeType.IsSet()
}

// HasMimeType returns a boolean if a field has been set.
func (o *MicrosoftGraphFile) HasMimeType() bool {
	if o != nil && o.MimeType.IsSet() {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given NullableString and assigns it to the MimeType field.
func (o *MicrosoftGraphFile) SetMimeType(v string) {
	o.MimeType.Set(&v)
}
// SetMimeTypeNil sets the value for MimeType to be an explicit nil
func (o *MicrosoftGraphFile) SetMimeTypeNil() {
	o.MimeType.Set(nil)
}

// UnsetMimeType ensures that no value is present for MimeType, not even an explicit nil
func (o *MicrosoftGraphFile) UnsetMimeType() {
	o.MimeType.Unset()
}

// GetProcessingMetadata returns the ProcessingMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphFile) GetProcessingMetadata() bool {
	if o == nil || o.ProcessingMetadata.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ProcessingMetadata.Get()
}

// GetProcessingMetadataOk returns a tuple with the ProcessingMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphFile) GetProcessingMetadataOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProcessingMetadata.Get(), o.ProcessingMetadata.IsSet()
}

// HasProcessingMetadata returns a boolean if a field has been set.
func (o *MicrosoftGraphFile) HasProcessingMetadata() bool {
	if o != nil && o.ProcessingMetadata.IsSet() {
		return true
	}

	return false
}

// SetProcessingMetadata gets a reference to the given NullableBool and assigns it to the ProcessingMetadata field.
func (o *MicrosoftGraphFile) SetProcessingMetadata(v bool) {
	o.ProcessingMetadata.Set(&v)
}
// SetProcessingMetadataNil sets the value for ProcessingMetadata to be an explicit nil
func (o *MicrosoftGraphFile) SetProcessingMetadataNil() {
	o.ProcessingMetadata.Set(nil)
}

// UnsetProcessingMetadata ensures that no value is present for ProcessingMetadata, not even an explicit nil
func (o *MicrosoftGraphFile) UnsetProcessingMetadata() {
	o.ProcessingMetadata.Unset()
}

func (o MicrosoftGraphFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hashes != nil {
		toSerialize["hashes"] = o.Hashes
	}
	if o.MimeType.IsSet() {
		toSerialize["mimeType"] = o.MimeType.Get()
	}
	if o.ProcessingMetadata.IsSet() {
		toSerialize["processingMetadata"] = o.ProcessingMetadata.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphFile struct {
	value *MicrosoftGraphFile
	isSet bool
}

func (v NullableMicrosoftGraphFile) Get() *MicrosoftGraphFile {
	return v.value
}

func (v *NullableMicrosoftGraphFile) Set(val *MicrosoftGraphFile) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphFile) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphFile(val *MicrosoftGraphFile) *NullableMicrosoftGraphFile {
	return &NullableMicrosoftGraphFile{value: val, isSet: true}
}

func (v NullableMicrosoftGraphFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


