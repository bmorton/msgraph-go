/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// MicrosoftGraphPhoto struct for MicrosoftGraphPhoto
type MicrosoftGraphPhoto struct {
	// Camera manufacturer. Read-only.
	CameraMake NullableString `json:"cameraMake,omitempty"`
	// Camera model. Read-only.
	CameraModel NullableString `json:"cameraModel,omitempty"`
	// The denominator for the exposure time fraction from the camera. Read-only.
	ExposureDenominator AnyOfnumberstringstring `json:"exposureDenominator,omitempty"`
	// The numerator for the exposure time fraction from the camera. Read-only.
	ExposureNumerator AnyOfnumberstringstring `json:"exposureNumerator,omitempty"`
	// The F-stop value from the camera. Read-only.
	FNumber AnyOfnumberstringstring `json:"fNumber,omitempty"`
	// The focal length from the camera. Read-only.
	FocalLength AnyOfnumberstringstring `json:"focalLength,omitempty"`
	// The ISO value from the camera. Read-only.
	Iso NullableInt32 `json:"iso,omitempty"`
	// The orientation value from the camera. Writable on OneDrive Personal.
	Orientation NullableInt32 `json:"orientation,omitempty"`
	// Represents the date and time the photo was taken. Read-only.
	TakenDateTime NullableTime `json:"takenDateTime,omitempty"`
}

// NewMicrosoftGraphPhoto instantiates a new MicrosoftGraphPhoto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPhoto() *MicrosoftGraphPhoto {
	this := MicrosoftGraphPhoto{}
	return &this
}

// NewMicrosoftGraphPhotoWithDefaults instantiates a new MicrosoftGraphPhoto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPhotoWithDefaults() *MicrosoftGraphPhoto {
	this := MicrosoftGraphPhoto{}
	return &this
}

// GetCameraMake returns the CameraMake field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPhoto) GetCameraMake() string {
	if o == nil || o.CameraMake.Get() == nil {
		var ret string
		return ret
	}
	return *o.CameraMake.Get()
}

// GetCameraMakeOk returns a tuple with the CameraMake field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPhoto) GetCameraMakeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CameraMake.Get(), o.CameraMake.IsSet()
}

// HasCameraMake returns a boolean if a field has been set.
func (o *MicrosoftGraphPhoto) HasCameraMake() bool {
	if o != nil && o.CameraMake.IsSet() {
		return true
	}

	return false
}

// SetCameraMake gets a reference to the given NullableString and assigns it to the CameraMake field.
func (o *MicrosoftGraphPhoto) SetCameraMake(v string) {
	o.CameraMake.Set(&v)
}
// SetCameraMakeNil sets the value for CameraMake to be an explicit nil
func (o *MicrosoftGraphPhoto) SetCameraMakeNil() {
	o.CameraMake.Set(nil)
}

// UnsetCameraMake ensures that no value is present for CameraMake, not even an explicit nil
func (o *MicrosoftGraphPhoto) UnsetCameraMake() {
	o.CameraMake.Unset()
}

// GetCameraModel returns the CameraModel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPhoto) GetCameraModel() string {
	if o == nil || o.CameraModel.Get() == nil {
		var ret string
		return ret
	}
	return *o.CameraModel.Get()
}

// GetCameraModelOk returns a tuple with the CameraModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPhoto) GetCameraModelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CameraModel.Get(), o.CameraModel.IsSet()
}

// HasCameraModel returns a boolean if a field has been set.
func (o *MicrosoftGraphPhoto) HasCameraModel() bool {
	if o != nil && o.CameraModel.IsSet() {
		return true
	}

	return false
}

// SetCameraModel gets a reference to the given NullableString and assigns it to the CameraModel field.
func (o *MicrosoftGraphPhoto) SetCameraModel(v string) {
	o.CameraModel.Set(&v)
}
// SetCameraModelNil sets the value for CameraModel to be an explicit nil
func (o *MicrosoftGraphPhoto) SetCameraModelNil() {
	o.CameraModel.Set(nil)
}

// UnsetCameraModel ensures that no value is present for CameraModel, not even an explicit nil
func (o *MicrosoftGraphPhoto) UnsetCameraModel() {
	o.CameraModel.Unset()
}

// GetExposureDenominator returns the ExposureDenominator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPhoto) GetExposureDenominator() AnyOfnumberstringstring {
	if o == nil  {
		var ret AnyOfnumberstringstring
		return ret
	}
	return o.ExposureDenominator
}

// GetExposureDenominatorOk returns a tuple with the ExposureDenominator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPhoto) GetExposureDenominatorOk() (*AnyOfnumberstringstring, bool) {
	if o == nil || o.ExposureDenominator == nil {
		return nil, false
	}
	return &o.ExposureDenominator, true
}

// HasExposureDenominator returns a boolean if a field has been set.
func (o *MicrosoftGraphPhoto) HasExposureDenominator() bool {
	if o != nil && o.ExposureDenominator != nil {
		return true
	}

	return false
}

// SetExposureDenominator gets a reference to the given AnyOfnumberstringstring and assigns it to the ExposureDenominator field.
func (o *MicrosoftGraphPhoto) SetExposureDenominator(v AnyOfnumberstringstring) {
	o.ExposureDenominator = v
}

// GetExposureNumerator returns the ExposureNumerator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPhoto) GetExposureNumerator() AnyOfnumberstringstring {
	if o == nil  {
		var ret AnyOfnumberstringstring
		return ret
	}
	return o.ExposureNumerator
}

// GetExposureNumeratorOk returns a tuple with the ExposureNumerator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPhoto) GetExposureNumeratorOk() (*AnyOfnumberstringstring, bool) {
	if o == nil || o.ExposureNumerator == nil {
		return nil, false
	}
	return &o.ExposureNumerator, true
}

// HasExposureNumerator returns a boolean if a field has been set.
func (o *MicrosoftGraphPhoto) HasExposureNumerator() bool {
	if o != nil && o.ExposureNumerator != nil {
		return true
	}

	return false
}

// SetExposureNumerator gets a reference to the given AnyOfnumberstringstring and assigns it to the ExposureNumerator field.
func (o *MicrosoftGraphPhoto) SetExposureNumerator(v AnyOfnumberstringstring) {
	o.ExposureNumerator = v
}

// GetFNumber returns the FNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPhoto) GetFNumber() AnyOfnumberstringstring {
	if o == nil  {
		var ret AnyOfnumberstringstring
		return ret
	}
	return o.FNumber
}

// GetFNumberOk returns a tuple with the FNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPhoto) GetFNumberOk() (*AnyOfnumberstringstring, bool) {
	if o == nil || o.FNumber == nil {
		return nil, false
	}
	return &o.FNumber, true
}

// HasFNumber returns a boolean if a field has been set.
func (o *MicrosoftGraphPhoto) HasFNumber() bool {
	if o != nil && o.FNumber != nil {
		return true
	}

	return false
}

// SetFNumber gets a reference to the given AnyOfnumberstringstring and assigns it to the FNumber field.
func (o *MicrosoftGraphPhoto) SetFNumber(v AnyOfnumberstringstring) {
	o.FNumber = v
}

// GetFocalLength returns the FocalLength field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPhoto) GetFocalLength() AnyOfnumberstringstring {
	if o == nil  {
		var ret AnyOfnumberstringstring
		return ret
	}
	return o.FocalLength
}

// GetFocalLengthOk returns a tuple with the FocalLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPhoto) GetFocalLengthOk() (*AnyOfnumberstringstring, bool) {
	if o == nil || o.FocalLength == nil {
		return nil, false
	}
	return &o.FocalLength, true
}

// HasFocalLength returns a boolean if a field has been set.
func (o *MicrosoftGraphPhoto) HasFocalLength() bool {
	if o != nil && o.FocalLength != nil {
		return true
	}

	return false
}

// SetFocalLength gets a reference to the given AnyOfnumberstringstring and assigns it to the FocalLength field.
func (o *MicrosoftGraphPhoto) SetFocalLength(v AnyOfnumberstringstring) {
	o.FocalLength = v
}

// GetIso returns the Iso field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPhoto) GetIso() int32 {
	if o == nil || o.Iso.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Iso.Get()
}

// GetIsoOk returns a tuple with the Iso field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPhoto) GetIsoOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Iso.Get(), o.Iso.IsSet()
}

// HasIso returns a boolean if a field has been set.
func (o *MicrosoftGraphPhoto) HasIso() bool {
	if o != nil && o.Iso.IsSet() {
		return true
	}

	return false
}

// SetIso gets a reference to the given NullableInt32 and assigns it to the Iso field.
func (o *MicrosoftGraphPhoto) SetIso(v int32) {
	o.Iso.Set(&v)
}
// SetIsoNil sets the value for Iso to be an explicit nil
func (o *MicrosoftGraphPhoto) SetIsoNil() {
	o.Iso.Set(nil)
}

// UnsetIso ensures that no value is present for Iso, not even an explicit nil
func (o *MicrosoftGraphPhoto) UnsetIso() {
	o.Iso.Unset()
}

// GetOrientation returns the Orientation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPhoto) GetOrientation() int32 {
	if o == nil || o.Orientation.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Orientation.Get()
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPhoto) GetOrientationOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Orientation.Get(), o.Orientation.IsSet()
}

// HasOrientation returns a boolean if a field has been set.
func (o *MicrosoftGraphPhoto) HasOrientation() bool {
	if o != nil && o.Orientation.IsSet() {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given NullableInt32 and assigns it to the Orientation field.
func (o *MicrosoftGraphPhoto) SetOrientation(v int32) {
	o.Orientation.Set(&v)
}
// SetOrientationNil sets the value for Orientation to be an explicit nil
func (o *MicrosoftGraphPhoto) SetOrientationNil() {
	o.Orientation.Set(nil)
}

// UnsetOrientation ensures that no value is present for Orientation, not even an explicit nil
func (o *MicrosoftGraphPhoto) UnsetOrientation() {
	o.Orientation.Unset()
}

// GetTakenDateTime returns the TakenDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPhoto) GetTakenDateTime() time.Time {
	if o == nil || o.TakenDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.TakenDateTime.Get()
}

// GetTakenDateTimeOk returns a tuple with the TakenDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPhoto) GetTakenDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TakenDateTime.Get(), o.TakenDateTime.IsSet()
}

// HasTakenDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphPhoto) HasTakenDateTime() bool {
	if o != nil && o.TakenDateTime.IsSet() {
		return true
	}

	return false
}

// SetTakenDateTime gets a reference to the given NullableTime and assigns it to the TakenDateTime field.
func (o *MicrosoftGraphPhoto) SetTakenDateTime(v time.Time) {
	o.TakenDateTime.Set(&v)
}
// SetTakenDateTimeNil sets the value for TakenDateTime to be an explicit nil
func (o *MicrosoftGraphPhoto) SetTakenDateTimeNil() {
	o.TakenDateTime.Set(nil)
}

// UnsetTakenDateTime ensures that no value is present for TakenDateTime, not even an explicit nil
func (o *MicrosoftGraphPhoto) UnsetTakenDateTime() {
	o.TakenDateTime.Unset()
}

func (o MicrosoftGraphPhoto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CameraMake.IsSet() {
		toSerialize["cameraMake"] = o.CameraMake.Get()
	}
	if o.CameraModel.IsSet() {
		toSerialize["cameraModel"] = o.CameraModel.Get()
	}
	if o.ExposureDenominator != nil {
		toSerialize["exposureDenominator"] = o.ExposureDenominator
	}
	if o.ExposureNumerator != nil {
		toSerialize["exposureNumerator"] = o.ExposureNumerator
	}
	if o.FNumber != nil {
		toSerialize["fNumber"] = o.FNumber
	}
	if o.FocalLength != nil {
		toSerialize["focalLength"] = o.FocalLength
	}
	if o.Iso.IsSet() {
		toSerialize["iso"] = o.Iso.Get()
	}
	if o.Orientation.IsSet() {
		toSerialize["orientation"] = o.Orientation.Get()
	}
	if o.TakenDateTime.IsSet() {
		toSerialize["takenDateTime"] = o.TakenDateTime.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPhoto struct {
	value *MicrosoftGraphPhoto
	isSet bool
}

func (v NullableMicrosoftGraphPhoto) Get() *MicrosoftGraphPhoto {
	return v.value
}

func (v *NullableMicrosoftGraphPhoto) Set(val *MicrosoftGraphPhoto) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPhoto) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPhoto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPhoto(val *MicrosoftGraphPhoto) *NullableMicrosoftGraphPhoto {
	return &NullableMicrosoftGraphPhoto{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPhoto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPhoto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

