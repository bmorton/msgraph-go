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

// MicrosoftGraphGeoCoordinates struct for MicrosoftGraphGeoCoordinates
type MicrosoftGraphGeoCoordinates struct {
	// Optional. The altitude (height), in feet,  above sea level for the item. Read-only.
	Altitude AnyOfnumberstringstring `json:"altitude,omitempty"`
	// Optional. The latitude, in decimal, for the item. Read-only.
	Latitude AnyOfnumberstringstring `json:"latitude,omitempty"`
	// Optional. The longitude, in decimal, for the item. Read-only.
	Longitude AnyOfnumberstringstring `json:"longitude,omitempty"`
}

// NewMicrosoftGraphGeoCoordinates instantiates a new MicrosoftGraphGeoCoordinates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphGeoCoordinates() *MicrosoftGraphGeoCoordinates {
	this := MicrosoftGraphGeoCoordinates{}
	return &this
}

// NewMicrosoftGraphGeoCoordinatesWithDefaults instantiates a new MicrosoftGraphGeoCoordinates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphGeoCoordinatesWithDefaults() *MicrosoftGraphGeoCoordinates {
	this := MicrosoftGraphGeoCoordinates{}
	return &this
}

// GetAltitude returns the Altitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphGeoCoordinates) GetAltitude() AnyOfnumberstringstring {
	if o == nil  {
		var ret AnyOfnumberstringstring
		return ret
	}
	return o.Altitude
}

// GetAltitudeOk returns a tuple with the Altitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphGeoCoordinates) GetAltitudeOk() (*AnyOfnumberstringstring, bool) {
	if o == nil || o.Altitude == nil {
		return nil, false
	}
	return &o.Altitude, true
}

// HasAltitude returns a boolean if a field has been set.
func (o *MicrosoftGraphGeoCoordinates) HasAltitude() bool {
	if o != nil && o.Altitude != nil {
		return true
	}

	return false
}

// SetAltitude gets a reference to the given AnyOfnumberstringstring and assigns it to the Altitude field.
func (o *MicrosoftGraphGeoCoordinates) SetAltitude(v AnyOfnumberstringstring) {
	o.Altitude = v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphGeoCoordinates) GetLatitude() AnyOfnumberstringstring {
	if o == nil  {
		var ret AnyOfnumberstringstring
		return ret
	}
	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphGeoCoordinates) GetLatitudeOk() (*AnyOfnumberstringstring, bool) {
	if o == nil || o.Latitude == nil {
		return nil, false
	}
	return &o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *MicrosoftGraphGeoCoordinates) HasLatitude() bool {
	if o != nil && o.Latitude != nil {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given AnyOfnumberstringstring and assigns it to the Latitude field.
func (o *MicrosoftGraphGeoCoordinates) SetLatitude(v AnyOfnumberstringstring) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphGeoCoordinates) GetLongitude() AnyOfnumberstringstring {
	if o == nil  {
		var ret AnyOfnumberstringstring
		return ret
	}
	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphGeoCoordinates) GetLongitudeOk() (*AnyOfnumberstringstring, bool) {
	if o == nil || o.Longitude == nil {
		return nil, false
	}
	return &o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *MicrosoftGraphGeoCoordinates) HasLongitude() bool {
	if o != nil && o.Longitude != nil {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given AnyOfnumberstringstring and assigns it to the Longitude field.
func (o *MicrosoftGraphGeoCoordinates) SetLongitude(v AnyOfnumberstringstring) {
	o.Longitude = v
}

func (o MicrosoftGraphGeoCoordinates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Altitude != nil {
		toSerialize["altitude"] = o.Altitude
	}
	if o.Latitude != nil {
		toSerialize["latitude"] = o.Latitude
	}
	if o.Longitude != nil {
		toSerialize["longitude"] = o.Longitude
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphGeoCoordinates struct {
	value *MicrosoftGraphGeoCoordinates
	isSet bool
}

func (v NullableMicrosoftGraphGeoCoordinates) Get() *MicrosoftGraphGeoCoordinates {
	return v.value
}

func (v *NullableMicrosoftGraphGeoCoordinates) Set(val *MicrosoftGraphGeoCoordinates) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphGeoCoordinates) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphGeoCoordinates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphGeoCoordinates(val *MicrosoftGraphGeoCoordinates) *NullableMicrosoftGraphGeoCoordinates {
	return &NullableMicrosoftGraphGeoCoordinates{value: val, isSet: true}
}

func (v NullableMicrosoftGraphGeoCoordinates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphGeoCoordinates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


