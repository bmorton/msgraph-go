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

// MicrosoftGraphSignInStatus struct for MicrosoftGraphSignInStatus
type MicrosoftGraphSignInStatus struct {
	// Provides additional details on the sign-in activity
	AdditionalDetails NullableString `json:"additionalDetails,omitempty"`
	// Provides the 5-6 digit error code that's generated during a sign-in failure. Check out the list of error codes and messages.
	ErrorCode NullableInt32 `json:"errorCode,omitempty"`
	// Provides the error message or the reason for failure for the corresponding sign-in activity. Check out the list of error codes and messages.
	FailureReason NullableString `json:"failureReason,omitempty"`
}

// NewMicrosoftGraphSignInStatus instantiates a new MicrosoftGraphSignInStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphSignInStatus() *MicrosoftGraphSignInStatus {
	this := MicrosoftGraphSignInStatus{}
	return &this
}

// NewMicrosoftGraphSignInStatusWithDefaults instantiates a new MicrosoftGraphSignInStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphSignInStatusWithDefaults() *MicrosoftGraphSignInStatus {
	this := MicrosoftGraphSignInStatus{}
	return &this
}

// GetAdditionalDetails returns the AdditionalDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSignInStatus) GetAdditionalDetails() string {
	if o == nil || o.AdditionalDetails.Get() == nil {
		var ret string
		return ret
	}
	return *o.AdditionalDetails.Get()
}

// GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSignInStatus) GetAdditionalDetailsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AdditionalDetails.Get(), o.AdditionalDetails.IsSet()
}

// HasAdditionalDetails returns a boolean if a field has been set.
func (o *MicrosoftGraphSignInStatus) HasAdditionalDetails() bool {
	if o != nil && o.AdditionalDetails.IsSet() {
		return true
	}

	return false
}

// SetAdditionalDetails gets a reference to the given NullableString and assigns it to the AdditionalDetails field.
func (o *MicrosoftGraphSignInStatus) SetAdditionalDetails(v string) {
	o.AdditionalDetails.Set(&v)
}
// SetAdditionalDetailsNil sets the value for AdditionalDetails to be an explicit nil
func (o *MicrosoftGraphSignInStatus) SetAdditionalDetailsNil() {
	o.AdditionalDetails.Set(nil)
}

// UnsetAdditionalDetails ensures that no value is present for AdditionalDetails, not even an explicit nil
func (o *MicrosoftGraphSignInStatus) UnsetAdditionalDetails() {
	o.AdditionalDetails.Unset()
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSignInStatus) GetErrorCode() int32 {
	if o == nil || o.ErrorCode.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCode.Get()
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSignInStatus) GetErrorCodeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ErrorCode.Get(), o.ErrorCode.IsSet()
}

// HasErrorCode returns a boolean if a field has been set.
func (o *MicrosoftGraphSignInStatus) HasErrorCode() bool {
	if o != nil && o.ErrorCode.IsSet() {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given NullableInt32 and assigns it to the ErrorCode field.
func (o *MicrosoftGraphSignInStatus) SetErrorCode(v int32) {
	o.ErrorCode.Set(&v)
}
// SetErrorCodeNil sets the value for ErrorCode to be an explicit nil
func (o *MicrosoftGraphSignInStatus) SetErrorCodeNil() {
	o.ErrorCode.Set(nil)
}

// UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
func (o *MicrosoftGraphSignInStatus) UnsetErrorCode() {
	o.ErrorCode.Unset()
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphSignInStatus) GetFailureReason() string {
	if o == nil || o.FailureReason.Get() == nil {
		var ret string
		return ret
	}
	return *o.FailureReason.Get()
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphSignInStatus) GetFailureReasonOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FailureReason.Get(), o.FailureReason.IsSet()
}

// HasFailureReason returns a boolean if a field has been set.
func (o *MicrosoftGraphSignInStatus) HasFailureReason() bool {
	if o != nil && o.FailureReason.IsSet() {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given NullableString and assigns it to the FailureReason field.
func (o *MicrosoftGraphSignInStatus) SetFailureReason(v string) {
	o.FailureReason.Set(&v)
}
// SetFailureReasonNil sets the value for FailureReason to be an explicit nil
func (o *MicrosoftGraphSignInStatus) SetFailureReasonNil() {
	o.FailureReason.Set(nil)
}

// UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
func (o *MicrosoftGraphSignInStatus) UnsetFailureReason() {
	o.FailureReason.Unset()
}

func (o MicrosoftGraphSignInStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalDetails.IsSet() {
		toSerialize["additionalDetails"] = o.AdditionalDetails.Get()
	}
	if o.ErrorCode.IsSet() {
		toSerialize["errorCode"] = o.ErrorCode.Get()
	}
	if o.FailureReason.IsSet() {
		toSerialize["failureReason"] = o.FailureReason.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphSignInStatus struct {
	value *MicrosoftGraphSignInStatus
	isSet bool
}

func (v NullableMicrosoftGraphSignInStatus) Get() *MicrosoftGraphSignInStatus {
	return v.value
}

func (v *NullableMicrosoftGraphSignInStatus) Set(val *MicrosoftGraphSignInStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphSignInStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphSignInStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphSignInStatus(val *MicrosoftGraphSignInStatus) *NullableMicrosoftGraphSignInStatus {
	return &NullableMicrosoftGraphSignInStatus{value: val, isSet: true}
}

func (v NullableMicrosoftGraphSignInStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphSignInStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


