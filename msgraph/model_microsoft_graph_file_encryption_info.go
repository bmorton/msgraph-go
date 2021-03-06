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

// MicrosoftGraphFileEncryptionInfo Contains properties for file encryption information for the content version of a line of business app.
type MicrosoftGraphFileEncryptionInfo struct {
	// The key used to encrypt the file content.
	EncryptionKey NullableString `json:"encryptionKey,omitempty"`
	// The file digest prior to encryption.
	FileDigest NullableString `json:"fileDigest,omitempty"`
	// The file digest algorithm.
	FileDigestAlgorithm NullableString `json:"fileDigestAlgorithm,omitempty"`
	// The initialization vector used for the encryption algorithm.
	InitializationVector NullableString `json:"initializationVector,omitempty"`
	// The hash of the encrypted file content + IV (content hash).
	Mac NullableString `json:"mac,omitempty"`
	// The key used to get mac.
	MacKey NullableString `json:"macKey,omitempty"`
	// The the profile identifier.
	ProfileIdentifier NullableString `json:"profileIdentifier,omitempty"`
}

// NewMicrosoftGraphFileEncryptionInfo instantiates a new MicrosoftGraphFileEncryptionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphFileEncryptionInfo() *MicrosoftGraphFileEncryptionInfo {
	this := MicrosoftGraphFileEncryptionInfo{}
	return &this
}

// NewMicrosoftGraphFileEncryptionInfoWithDefaults instantiates a new MicrosoftGraphFileEncryptionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphFileEncryptionInfoWithDefaults() *MicrosoftGraphFileEncryptionInfo {
	this := MicrosoftGraphFileEncryptionInfo{}
	return &this
}

// GetEncryptionKey returns the EncryptionKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphFileEncryptionInfo) GetEncryptionKey() string {
	if o == nil || o.EncryptionKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.EncryptionKey.Get()
}

// GetEncryptionKeyOk returns a tuple with the EncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphFileEncryptionInfo) GetEncryptionKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EncryptionKey.Get(), o.EncryptionKey.IsSet()
}

// HasEncryptionKey returns a boolean if a field has been set.
func (o *MicrosoftGraphFileEncryptionInfo) HasEncryptionKey() bool {
	if o != nil && o.EncryptionKey.IsSet() {
		return true
	}

	return false
}

// SetEncryptionKey gets a reference to the given NullableString and assigns it to the EncryptionKey field.
func (o *MicrosoftGraphFileEncryptionInfo) SetEncryptionKey(v string) {
	o.EncryptionKey.Set(&v)
}
// SetEncryptionKeyNil sets the value for EncryptionKey to be an explicit nil
func (o *MicrosoftGraphFileEncryptionInfo) SetEncryptionKeyNil() {
	o.EncryptionKey.Set(nil)
}

// UnsetEncryptionKey ensures that no value is present for EncryptionKey, not even an explicit nil
func (o *MicrosoftGraphFileEncryptionInfo) UnsetEncryptionKey() {
	o.EncryptionKey.Unset()
}

// GetFileDigest returns the FileDigest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphFileEncryptionInfo) GetFileDigest() string {
	if o == nil || o.FileDigest.Get() == nil {
		var ret string
		return ret
	}
	return *o.FileDigest.Get()
}

// GetFileDigestOk returns a tuple with the FileDigest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphFileEncryptionInfo) GetFileDigestOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FileDigest.Get(), o.FileDigest.IsSet()
}

// HasFileDigest returns a boolean if a field has been set.
func (o *MicrosoftGraphFileEncryptionInfo) HasFileDigest() bool {
	if o != nil && o.FileDigest.IsSet() {
		return true
	}

	return false
}

// SetFileDigest gets a reference to the given NullableString and assigns it to the FileDigest field.
func (o *MicrosoftGraphFileEncryptionInfo) SetFileDigest(v string) {
	o.FileDigest.Set(&v)
}
// SetFileDigestNil sets the value for FileDigest to be an explicit nil
func (o *MicrosoftGraphFileEncryptionInfo) SetFileDigestNil() {
	o.FileDigest.Set(nil)
}

// UnsetFileDigest ensures that no value is present for FileDigest, not even an explicit nil
func (o *MicrosoftGraphFileEncryptionInfo) UnsetFileDigest() {
	o.FileDigest.Unset()
}

// GetFileDigestAlgorithm returns the FileDigestAlgorithm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphFileEncryptionInfo) GetFileDigestAlgorithm() string {
	if o == nil || o.FileDigestAlgorithm.Get() == nil {
		var ret string
		return ret
	}
	return *o.FileDigestAlgorithm.Get()
}

// GetFileDigestAlgorithmOk returns a tuple with the FileDigestAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphFileEncryptionInfo) GetFileDigestAlgorithmOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FileDigestAlgorithm.Get(), o.FileDigestAlgorithm.IsSet()
}

// HasFileDigestAlgorithm returns a boolean if a field has been set.
func (o *MicrosoftGraphFileEncryptionInfo) HasFileDigestAlgorithm() bool {
	if o != nil && o.FileDigestAlgorithm.IsSet() {
		return true
	}

	return false
}

// SetFileDigestAlgorithm gets a reference to the given NullableString and assigns it to the FileDigestAlgorithm field.
func (o *MicrosoftGraphFileEncryptionInfo) SetFileDigestAlgorithm(v string) {
	o.FileDigestAlgorithm.Set(&v)
}
// SetFileDigestAlgorithmNil sets the value for FileDigestAlgorithm to be an explicit nil
func (o *MicrosoftGraphFileEncryptionInfo) SetFileDigestAlgorithmNil() {
	o.FileDigestAlgorithm.Set(nil)
}

// UnsetFileDigestAlgorithm ensures that no value is present for FileDigestAlgorithm, not even an explicit nil
func (o *MicrosoftGraphFileEncryptionInfo) UnsetFileDigestAlgorithm() {
	o.FileDigestAlgorithm.Unset()
}

// GetInitializationVector returns the InitializationVector field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphFileEncryptionInfo) GetInitializationVector() string {
	if o == nil || o.InitializationVector.Get() == nil {
		var ret string
		return ret
	}
	return *o.InitializationVector.Get()
}

// GetInitializationVectorOk returns a tuple with the InitializationVector field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphFileEncryptionInfo) GetInitializationVectorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InitializationVector.Get(), o.InitializationVector.IsSet()
}

// HasInitializationVector returns a boolean if a field has been set.
func (o *MicrosoftGraphFileEncryptionInfo) HasInitializationVector() bool {
	if o != nil && o.InitializationVector.IsSet() {
		return true
	}

	return false
}

// SetInitializationVector gets a reference to the given NullableString and assigns it to the InitializationVector field.
func (o *MicrosoftGraphFileEncryptionInfo) SetInitializationVector(v string) {
	o.InitializationVector.Set(&v)
}
// SetInitializationVectorNil sets the value for InitializationVector to be an explicit nil
func (o *MicrosoftGraphFileEncryptionInfo) SetInitializationVectorNil() {
	o.InitializationVector.Set(nil)
}

// UnsetInitializationVector ensures that no value is present for InitializationVector, not even an explicit nil
func (o *MicrosoftGraphFileEncryptionInfo) UnsetInitializationVector() {
	o.InitializationVector.Unset()
}

// GetMac returns the Mac field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphFileEncryptionInfo) GetMac() string {
	if o == nil || o.Mac.Get() == nil {
		var ret string
		return ret
	}
	return *o.Mac.Get()
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphFileEncryptionInfo) GetMacOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Mac.Get(), o.Mac.IsSet()
}

// HasMac returns a boolean if a field has been set.
func (o *MicrosoftGraphFileEncryptionInfo) HasMac() bool {
	if o != nil && o.Mac.IsSet() {
		return true
	}

	return false
}

// SetMac gets a reference to the given NullableString and assigns it to the Mac field.
func (o *MicrosoftGraphFileEncryptionInfo) SetMac(v string) {
	o.Mac.Set(&v)
}
// SetMacNil sets the value for Mac to be an explicit nil
func (o *MicrosoftGraphFileEncryptionInfo) SetMacNil() {
	o.Mac.Set(nil)
}

// UnsetMac ensures that no value is present for Mac, not even an explicit nil
func (o *MicrosoftGraphFileEncryptionInfo) UnsetMac() {
	o.Mac.Unset()
}

// GetMacKey returns the MacKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphFileEncryptionInfo) GetMacKey() string {
	if o == nil || o.MacKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.MacKey.Get()
}

// GetMacKeyOk returns a tuple with the MacKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphFileEncryptionInfo) GetMacKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MacKey.Get(), o.MacKey.IsSet()
}

// HasMacKey returns a boolean if a field has been set.
func (o *MicrosoftGraphFileEncryptionInfo) HasMacKey() bool {
	if o != nil && o.MacKey.IsSet() {
		return true
	}

	return false
}

// SetMacKey gets a reference to the given NullableString and assigns it to the MacKey field.
func (o *MicrosoftGraphFileEncryptionInfo) SetMacKey(v string) {
	o.MacKey.Set(&v)
}
// SetMacKeyNil sets the value for MacKey to be an explicit nil
func (o *MicrosoftGraphFileEncryptionInfo) SetMacKeyNil() {
	o.MacKey.Set(nil)
}

// UnsetMacKey ensures that no value is present for MacKey, not even an explicit nil
func (o *MicrosoftGraphFileEncryptionInfo) UnsetMacKey() {
	o.MacKey.Unset()
}

// GetProfileIdentifier returns the ProfileIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphFileEncryptionInfo) GetProfileIdentifier() string {
	if o == nil || o.ProfileIdentifier.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProfileIdentifier.Get()
}

// GetProfileIdentifierOk returns a tuple with the ProfileIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphFileEncryptionInfo) GetProfileIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ProfileIdentifier.Get(), o.ProfileIdentifier.IsSet()
}

// HasProfileIdentifier returns a boolean if a field has been set.
func (o *MicrosoftGraphFileEncryptionInfo) HasProfileIdentifier() bool {
	if o != nil && o.ProfileIdentifier.IsSet() {
		return true
	}

	return false
}

// SetProfileIdentifier gets a reference to the given NullableString and assigns it to the ProfileIdentifier field.
func (o *MicrosoftGraphFileEncryptionInfo) SetProfileIdentifier(v string) {
	o.ProfileIdentifier.Set(&v)
}
// SetProfileIdentifierNil sets the value for ProfileIdentifier to be an explicit nil
func (o *MicrosoftGraphFileEncryptionInfo) SetProfileIdentifierNil() {
	o.ProfileIdentifier.Set(nil)
}

// UnsetProfileIdentifier ensures that no value is present for ProfileIdentifier, not even an explicit nil
func (o *MicrosoftGraphFileEncryptionInfo) UnsetProfileIdentifier() {
	o.ProfileIdentifier.Unset()
}

func (o MicrosoftGraphFileEncryptionInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EncryptionKey.IsSet() {
		toSerialize["encryptionKey"] = o.EncryptionKey.Get()
	}
	if o.FileDigest.IsSet() {
		toSerialize["fileDigest"] = o.FileDigest.Get()
	}
	if o.FileDigestAlgorithm.IsSet() {
		toSerialize["fileDigestAlgorithm"] = o.FileDigestAlgorithm.Get()
	}
	if o.InitializationVector.IsSet() {
		toSerialize["initializationVector"] = o.InitializationVector.Get()
	}
	if o.Mac.IsSet() {
		toSerialize["mac"] = o.Mac.Get()
	}
	if o.MacKey.IsSet() {
		toSerialize["macKey"] = o.MacKey.Get()
	}
	if o.ProfileIdentifier.IsSet() {
		toSerialize["profileIdentifier"] = o.ProfileIdentifier.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphFileEncryptionInfo struct {
	value *MicrosoftGraphFileEncryptionInfo
	isSet bool
}

func (v NullableMicrosoftGraphFileEncryptionInfo) Get() *MicrosoftGraphFileEncryptionInfo {
	return v.value
}

func (v *NullableMicrosoftGraphFileEncryptionInfo) Set(val *MicrosoftGraphFileEncryptionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphFileEncryptionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphFileEncryptionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphFileEncryptionInfo(val *MicrosoftGraphFileEncryptionInfo) *NullableMicrosoftGraphFileEncryptionInfo {
	return &NullableMicrosoftGraphFileEncryptionInfo{value: val, isSet: true}
}

func (v NullableMicrosoftGraphFileEncryptionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphFileEncryptionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


