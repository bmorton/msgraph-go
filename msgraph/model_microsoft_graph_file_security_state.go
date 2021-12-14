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

// MicrosoftGraphFileSecurityState struct for MicrosoftGraphFileSecurityState
type MicrosoftGraphFileSecurityState struct {
	// Complex type containing file hashes (cryptographic and location-sensitive).
	FileHash AnyOfmicrosoftGraphFileHash `json:"fileHash,omitempty"`
	// File name (without path).
	Name NullableString `json:"name,omitempty"`
	// Full file path of the file/imageFile.
	Path NullableString `json:"path,omitempty"`
	// Provider generated/calculated risk score of the alert file. Recommended value range of 0-1, which equates to a percentage.
	RiskScore NullableString `json:"riskScore,omitempty"`
}

// NewMicrosoftGraphFileSecurityState instantiates a new MicrosoftGraphFileSecurityState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphFileSecurityState() *MicrosoftGraphFileSecurityState {
	this := MicrosoftGraphFileSecurityState{}
	return &this
}

// NewMicrosoftGraphFileSecurityStateWithDefaults instantiates a new MicrosoftGraphFileSecurityState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphFileSecurityStateWithDefaults() *MicrosoftGraphFileSecurityState {
	this := MicrosoftGraphFileSecurityState{}
	return &this
}

// GetFileHash returns the FileHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphFileSecurityState) GetFileHash() AnyOfmicrosoftGraphFileHash {
	if o == nil  {
		var ret AnyOfmicrosoftGraphFileHash
		return ret
	}
	return o.FileHash
}

// GetFileHashOk returns a tuple with the FileHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphFileSecurityState) GetFileHashOk() (*AnyOfmicrosoftGraphFileHash, bool) {
	if o == nil || o.FileHash == nil {
		return nil, false
	}
	return &o.FileHash, true
}

// HasFileHash returns a boolean if a field has been set.
func (o *MicrosoftGraphFileSecurityState) HasFileHash() bool {
	if o != nil && o.FileHash != nil {
		return true
	}

	return false
}

// SetFileHash gets a reference to the given AnyOfmicrosoftGraphFileHash and assigns it to the FileHash field.
func (o *MicrosoftGraphFileSecurityState) SetFileHash(v AnyOfmicrosoftGraphFileHash) {
	o.FileHash = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphFileSecurityState) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphFileSecurityState) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftGraphFileSecurityState) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MicrosoftGraphFileSecurityState) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MicrosoftGraphFileSecurityState) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MicrosoftGraphFileSecurityState) UnsetName() {
	o.Name.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphFileSecurityState) GetPath() string {
	if o == nil || o.Path.Get() == nil {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphFileSecurityState) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *MicrosoftGraphFileSecurityState) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *MicrosoftGraphFileSecurityState) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *MicrosoftGraphFileSecurityState) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *MicrosoftGraphFileSecurityState) UnsetPath() {
	o.Path.Unset()
}

// GetRiskScore returns the RiskScore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphFileSecurityState) GetRiskScore() string {
	if o == nil || o.RiskScore.Get() == nil {
		var ret string
		return ret
	}
	return *o.RiskScore.Get()
}

// GetRiskScoreOk returns a tuple with the RiskScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphFileSecurityState) GetRiskScoreOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RiskScore.Get(), o.RiskScore.IsSet()
}

// HasRiskScore returns a boolean if a field has been set.
func (o *MicrosoftGraphFileSecurityState) HasRiskScore() bool {
	if o != nil && o.RiskScore.IsSet() {
		return true
	}

	return false
}

// SetRiskScore gets a reference to the given NullableString and assigns it to the RiskScore field.
func (o *MicrosoftGraphFileSecurityState) SetRiskScore(v string) {
	o.RiskScore.Set(&v)
}
// SetRiskScoreNil sets the value for RiskScore to be an explicit nil
func (o *MicrosoftGraphFileSecurityState) SetRiskScoreNil() {
	o.RiskScore.Set(nil)
}

// UnsetRiskScore ensures that no value is present for RiskScore, not even an explicit nil
func (o *MicrosoftGraphFileSecurityState) UnsetRiskScore() {
	o.RiskScore.Unset()
}

func (o MicrosoftGraphFileSecurityState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FileHash != nil {
		toSerialize["fileHash"] = o.FileHash
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.RiskScore.IsSet() {
		toSerialize["riskScore"] = o.RiskScore.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphFileSecurityState struct {
	value *MicrosoftGraphFileSecurityState
	isSet bool
}

func (v NullableMicrosoftGraphFileSecurityState) Get() *MicrosoftGraphFileSecurityState {
	return v.value
}

func (v *NullableMicrosoftGraphFileSecurityState) Set(val *MicrosoftGraphFileSecurityState) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphFileSecurityState) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphFileSecurityState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphFileSecurityState(val *MicrosoftGraphFileSecurityState) *NullableMicrosoftGraphFileSecurityState {
	return &NullableMicrosoftGraphFileSecurityState{value: val, isSet: true}
}

func (v NullableMicrosoftGraphFileSecurityState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphFileSecurityState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


