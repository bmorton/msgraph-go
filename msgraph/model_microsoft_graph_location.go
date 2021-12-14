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

// MicrosoftGraphLocation struct for MicrosoftGraphLocation
type MicrosoftGraphLocation struct {
	// The street address of the location.
	Address AnyOfmicrosoftGraphPhysicalAddress `json:"address,omitempty"`
	// The geographic coordinates and elevation of the location.
	Coordinates AnyOfmicrosoftGraphOutlookGeoCoordinates `json:"coordinates,omitempty"`
	// The name associated with the location.
	DisplayName NullableString `json:"displayName,omitempty"`
	// Optional email address of the location.
	LocationEmailAddress NullableString `json:"locationEmailAddress,omitempty"`
	// The type of location. The possible values are: default, conferenceRoom, homeAddress, businessAddress,geoCoordinates, streetAddress, hotel, restaurant, localBusiness, postalAddress. Read-only.
	LocationType AnyOfmicrosoftGraphLocationType `json:"locationType,omitempty"`
	// Optional URI representing the location.
	LocationUri NullableString `json:"locationUri,omitempty"`
	// For internal use only.
	UniqueId NullableString `json:"uniqueId,omitempty"`
	// For internal use only.
	UniqueIdType AnyOfmicrosoftGraphLocationUniqueIdType `json:"uniqueIdType,omitempty"`
}

// NewMicrosoftGraphLocation instantiates a new MicrosoftGraphLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphLocation() *MicrosoftGraphLocation {
	this := MicrosoftGraphLocation{}
	return &this
}

// NewMicrosoftGraphLocationWithDefaults instantiates a new MicrosoftGraphLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphLocationWithDefaults() *MicrosoftGraphLocation {
	this := MicrosoftGraphLocation{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphLocation) GetAddress() AnyOfmicrosoftGraphPhysicalAddress {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPhysicalAddress
		return ret
	}
	return o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphLocation) GetAddressOk() (*AnyOfmicrosoftGraphPhysicalAddress, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return &o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *MicrosoftGraphLocation) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given AnyOfmicrosoftGraphPhysicalAddress and assigns it to the Address field.
func (o *MicrosoftGraphLocation) SetAddress(v AnyOfmicrosoftGraphPhysicalAddress) {
	o.Address = v
}

// GetCoordinates returns the Coordinates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphLocation) GetCoordinates() AnyOfmicrosoftGraphOutlookGeoCoordinates {
	if o == nil  {
		var ret AnyOfmicrosoftGraphOutlookGeoCoordinates
		return ret
	}
	return o.Coordinates
}

// GetCoordinatesOk returns a tuple with the Coordinates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphLocation) GetCoordinatesOk() (*AnyOfmicrosoftGraphOutlookGeoCoordinates, bool) {
	if o == nil || o.Coordinates == nil {
		return nil, false
	}
	return &o.Coordinates, true
}

// HasCoordinates returns a boolean if a field has been set.
func (o *MicrosoftGraphLocation) HasCoordinates() bool {
	if o != nil && o.Coordinates != nil {
		return true
	}

	return false
}

// SetCoordinates gets a reference to the given AnyOfmicrosoftGraphOutlookGeoCoordinates and assigns it to the Coordinates field.
func (o *MicrosoftGraphLocation) SetCoordinates(v AnyOfmicrosoftGraphOutlookGeoCoordinates) {
	o.Coordinates = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphLocation) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphLocation) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphLocation) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphLocation) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphLocation) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphLocation) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetLocationEmailAddress returns the LocationEmailAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphLocation) GetLocationEmailAddress() string {
	if o == nil || o.LocationEmailAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.LocationEmailAddress.Get()
}

// GetLocationEmailAddressOk returns a tuple with the LocationEmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphLocation) GetLocationEmailAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LocationEmailAddress.Get(), o.LocationEmailAddress.IsSet()
}

// HasLocationEmailAddress returns a boolean if a field has been set.
func (o *MicrosoftGraphLocation) HasLocationEmailAddress() bool {
	if o != nil && o.LocationEmailAddress.IsSet() {
		return true
	}

	return false
}

// SetLocationEmailAddress gets a reference to the given NullableString and assigns it to the LocationEmailAddress field.
func (o *MicrosoftGraphLocation) SetLocationEmailAddress(v string) {
	o.LocationEmailAddress.Set(&v)
}
// SetLocationEmailAddressNil sets the value for LocationEmailAddress to be an explicit nil
func (o *MicrosoftGraphLocation) SetLocationEmailAddressNil() {
	o.LocationEmailAddress.Set(nil)
}

// UnsetLocationEmailAddress ensures that no value is present for LocationEmailAddress, not even an explicit nil
func (o *MicrosoftGraphLocation) UnsetLocationEmailAddress() {
	o.LocationEmailAddress.Unset()
}

// GetLocationType returns the LocationType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphLocation) GetLocationType() AnyOfmicrosoftGraphLocationType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphLocationType
		return ret
	}
	return o.LocationType
}

// GetLocationTypeOk returns a tuple with the LocationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphLocation) GetLocationTypeOk() (*AnyOfmicrosoftGraphLocationType, bool) {
	if o == nil || o.LocationType == nil {
		return nil, false
	}
	return &o.LocationType, true
}

// HasLocationType returns a boolean if a field has been set.
func (o *MicrosoftGraphLocation) HasLocationType() bool {
	if o != nil && o.LocationType != nil {
		return true
	}

	return false
}

// SetLocationType gets a reference to the given AnyOfmicrosoftGraphLocationType and assigns it to the LocationType field.
func (o *MicrosoftGraphLocation) SetLocationType(v AnyOfmicrosoftGraphLocationType) {
	o.LocationType = v
}

// GetLocationUri returns the LocationUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphLocation) GetLocationUri() string {
	if o == nil || o.LocationUri.Get() == nil {
		var ret string
		return ret
	}
	return *o.LocationUri.Get()
}

// GetLocationUriOk returns a tuple with the LocationUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphLocation) GetLocationUriOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LocationUri.Get(), o.LocationUri.IsSet()
}

// HasLocationUri returns a boolean if a field has been set.
func (o *MicrosoftGraphLocation) HasLocationUri() bool {
	if o != nil && o.LocationUri.IsSet() {
		return true
	}

	return false
}

// SetLocationUri gets a reference to the given NullableString and assigns it to the LocationUri field.
func (o *MicrosoftGraphLocation) SetLocationUri(v string) {
	o.LocationUri.Set(&v)
}
// SetLocationUriNil sets the value for LocationUri to be an explicit nil
func (o *MicrosoftGraphLocation) SetLocationUriNil() {
	o.LocationUri.Set(nil)
}

// UnsetLocationUri ensures that no value is present for LocationUri, not even an explicit nil
func (o *MicrosoftGraphLocation) UnsetLocationUri() {
	o.LocationUri.Unset()
}

// GetUniqueId returns the UniqueId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphLocation) GetUniqueId() string {
	if o == nil || o.UniqueId.Get() == nil {
		var ret string
		return ret
	}
	return *o.UniqueId.Get()
}

// GetUniqueIdOk returns a tuple with the UniqueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphLocation) GetUniqueIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UniqueId.Get(), o.UniqueId.IsSet()
}

// HasUniqueId returns a boolean if a field has been set.
func (o *MicrosoftGraphLocation) HasUniqueId() bool {
	if o != nil && o.UniqueId.IsSet() {
		return true
	}

	return false
}

// SetUniqueId gets a reference to the given NullableString and assigns it to the UniqueId field.
func (o *MicrosoftGraphLocation) SetUniqueId(v string) {
	o.UniqueId.Set(&v)
}
// SetUniqueIdNil sets the value for UniqueId to be an explicit nil
func (o *MicrosoftGraphLocation) SetUniqueIdNil() {
	o.UniqueId.Set(nil)
}

// UnsetUniqueId ensures that no value is present for UniqueId, not even an explicit nil
func (o *MicrosoftGraphLocation) UnsetUniqueId() {
	o.UniqueId.Unset()
}

// GetUniqueIdType returns the UniqueIdType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphLocation) GetUniqueIdType() AnyOfmicrosoftGraphLocationUniqueIdType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphLocationUniqueIdType
		return ret
	}
	return o.UniqueIdType
}

// GetUniqueIdTypeOk returns a tuple with the UniqueIdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphLocation) GetUniqueIdTypeOk() (*AnyOfmicrosoftGraphLocationUniqueIdType, bool) {
	if o == nil || o.UniqueIdType == nil {
		return nil, false
	}
	return &o.UniqueIdType, true
}

// HasUniqueIdType returns a boolean if a field has been set.
func (o *MicrosoftGraphLocation) HasUniqueIdType() bool {
	if o != nil && o.UniqueIdType != nil {
		return true
	}

	return false
}

// SetUniqueIdType gets a reference to the given AnyOfmicrosoftGraphLocationUniqueIdType and assigns it to the UniqueIdType field.
func (o *MicrosoftGraphLocation) SetUniqueIdType(v AnyOfmicrosoftGraphLocationUniqueIdType) {
	o.UniqueIdType = v
}

func (o MicrosoftGraphLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Coordinates != nil {
		toSerialize["coordinates"] = o.Coordinates
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.LocationEmailAddress.IsSet() {
		toSerialize["locationEmailAddress"] = o.LocationEmailAddress.Get()
	}
	if o.LocationType != nil {
		toSerialize["locationType"] = o.LocationType
	}
	if o.LocationUri.IsSet() {
		toSerialize["locationUri"] = o.LocationUri.Get()
	}
	if o.UniqueId.IsSet() {
		toSerialize["uniqueId"] = o.UniqueId.Get()
	}
	if o.UniqueIdType != nil {
		toSerialize["uniqueIdType"] = o.UniqueIdType
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphLocation struct {
	value *MicrosoftGraphLocation
	isSet bool
}

func (v NullableMicrosoftGraphLocation) Get() *MicrosoftGraphLocation {
	return v.value
}

func (v *NullableMicrosoftGraphLocation) Set(val *MicrosoftGraphLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphLocation(val *MicrosoftGraphLocation) *NullableMicrosoftGraphLocation {
	return &NullableMicrosoftGraphLocation{value: val, isSet: true}
}

func (v NullableMicrosoftGraphLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


