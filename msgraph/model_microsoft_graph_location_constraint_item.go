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

// MicrosoftGraphLocationConstraintItem struct for MicrosoftGraphLocationConstraintItem
type MicrosoftGraphLocationConstraintItem struct {
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
	// If set to true and the specified resource is busy, findMeetingTimes looks for another resource that is free. If set to false and the specified resource is busy, findMeetingTimes returns the resource best ranked in the user's cache without checking if it's free. Default is true.
	ResolveAvailability NullableBool `json:"resolveAvailability,omitempty"`
}

// NewMicrosoftGraphLocationConstraintItem instantiates a new MicrosoftGraphLocationConstraintItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphLocationConstraintItem() *MicrosoftGraphLocationConstraintItem {
	this := MicrosoftGraphLocationConstraintItem{}
	return &this
}

// NewMicrosoftGraphLocationConstraintItemWithDefaults instantiates a new MicrosoftGraphLocationConstraintItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphLocationConstraintItemWithDefaults() *MicrosoftGraphLocationConstraintItem {
	this := MicrosoftGraphLocationConstraintItem{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphLocationConstraintItem) GetAddress() AnyOfmicrosoftGraphPhysicalAddress {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPhysicalAddress
		return ret
	}
	return o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphLocationConstraintItem) GetAddressOk() (*AnyOfmicrosoftGraphPhysicalAddress, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return &o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *MicrosoftGraphLocationConstraintItem) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given AnyOfmicrosoftGraphPhysicalAddress and assigns it to the Address field.
func (o *MicrosoftGraphLocationConstraintItem) SetAddress(v AnyOfmicrosoftGraphPhysicalAddress) {
	o.Address = v
}

// GetCoordinates returns the Coordinates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphLocationConstraintItem) GetCoordinates() AnyOfmicrosoftGraphOutlookGeoCoordinates {
	if o == nil  {
		var ret AnyOfmicrosoftGraphOutlookGeoCoordinates
		return ret
	}
	return o.Coordinates
}

// GetCoordinatesOk returns a tuple with the Coordinates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphLocationConstraintItem) GetCoordinatesOk() (*AnyOfmicrosoftGraphOutlookGeoCoordinates, bool) {
	if o == nil || o.Coordinates == nil {
		return nil, false
	}
	return &o.Coordinates, true
}

// HasCoordinates returns a boolean if a field has been set.
func (o *MicrosoftGraphLocationConstraintItem) HasCoordinates() bool {
	if o != nil && o.Coordinates != nil {
		return true
	}

	return false
}

// SetCoordinates gets a reference to the given AnyOfmicrosoftGraphOutlookGeoCoordinates and assigns it to the Coordinates field.
func (o *MicrosoftGraphLocationConstraintItem) SetCoordinates(v AnyOfmicrosoftGraphOutlookGeoCoordinates) {
	o.Coordinates = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphLocationConstraintItem) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphLocationConstraintItem) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphLocationConstraintItem) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphLocationConstraintItem) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphLocationConstraintItem) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphLocationConstraintItem) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetLocationEmailAddress returns the LocationEmailAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphLocationConstraintItem) GetLocationEmailAddress() string {
	if o == nil || o.LocationEmailAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.LocationEmailAddress.Get()
}

// GetLocationEmailAddressOk returns a tuple with the LocationEmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphLocationConstraintItem) GetLocationEmailAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LocationEmailAddress.Get(), o.LocationEmailAddress.IsSet()
}

// HasLocationEmailAddress returns a boolean if a field has been set.
func (o *MicrosoftGraphLocationConstraintItem) HasLocationEmailAddress() bool {
	if o != nil && o.LocationEmailAddress.IsSet() {
		return true
	}

	return false
}

// SetLocationEmailAddress gets a reference to the given NullableString and assigns it to the LocationEmailAddress field.
func (o *MicrosoftGraphLocationConstraintItem) SetLocationEmailAddress(v string) {
	o.LocationEmailAddress.Set(&v)
}
// SetLocationEmailAddressNil sets the value for LocationEmailAddress to be an explicit nil
func (o *MicrosoftGraphLocationConstraintItem) SetLocationEmailAddressNil() {
	o.LocationEmailAddress.Set(nil)
}

// UnsetLocationEmailAddress ensures that no value is present for LocationEmailAddress, not even an explicit nil
func (o *MicrosoftGraphLocationConstraintItem) UnsetLocationEmailAddress() {
	o.LocationEmailAddress.Unset()
}

// GetLocationType returns the LocationType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphLocationConstraintItem) GetLocationType() AnyOfmicrosoftGraphLocationType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphLocationType
		return ret
	}
	return o.LocationType
}

// GetLocationTypeOk returns a tuple with the LocationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphLocationConstraintItem) GetLocationTypeOk() (*AnyOfmicrosoftGraphLocationType, bool) {
	if o == nil || o.LocationType == nil {
		return nil, false
	}
	return &o.LocationType, true
}

// HasLocationType returns a boolean if a field has been set.
func (o *MicrosoftGraphLocationConstraintItem) HasLocationType() bool {
	if o != nil && o.LocationType != nil {
		return true
	}

	return false
}

// SetLocationType gets a reference to the given AnyOfmicrosoftGraphLocationType and assigns it to the LocationType field.
func (o *MicrosoftGraphLocationConstraintItem) SetLocationType(v AnyOfmicrosoftGraphLocationType) {
	o.LocationType = v
}

// GetLocationUri returns the LocationUri field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphLocationConstraintItem) GetLocationUri() string {
	if o == nil || o.LocationUri.Get() == nil {
		var ret string
		return ret
	}
	return *o.LocationUri.Get()
}

// GetLocationUriOk returns a tuple with the LocationUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphLocationConstraintItem) GetLocationUriOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LocationUri.Get(), o.LocationUri.IsSet()
}

// HasLocationUri returns a boolean if a field has been set.
func (o *MicrosoftGraphLocationConstraintItem) HasLocationUri() bool {
	if o != nil && o.LocationUri.IsSet() {
		return true
	}

	return false
}

// SetLocationUri gets a reference to the given NullableString and assigns it to the LocationUri field.
func (o *MicrosoftGraphLocationConstraintItem) SetLocationUri(v string) {
	o.LocationUri.Set(&v)
}
// SetLocationUriNil sets the value for LocationUri to be an explicit nil
func (o *MicrosoftGraphLocationConstraintItem) SetLocationUriNil() {
	o.LocationUri.Set(nil)
}

// UnsetLocationUri ensures that no value is present for LocationUri, not even an explicit nil
func (o *MicrosoftGraphLocationConstraintItem) UnsetLocationUri() {
	o.LocationUri.Unset()
}

// GetUniqueId returns the UniqueId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphLocationConstraintItem) GetUniqueId() string {
	if o == nil || o.UniqueId.Get() == nil {
		var ret string
		return ret
	}
	return *o.UniqueId.Get()
}

// GetUniqueIdOk returns a tuple with the UniqueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphLocationConstraintItem) GetUniqueIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UniqueId.Get(), o.UniqueId.IsSet()
}

// HasUniqueId returns a boolean if a field has been set.
func (o *MicrosoftGraphLocationConstraintItem) HasUniqueId() bool {
	if o != nil && o.UniqueId.IsSet() {
		return true
	}

	return false
}

// SetUniqueId gets a reference to the given NullableString and assigns it to the UniqueId field.
func (o *MicrosoftGraphLocationConstraintItem) SetUniqueId(v string) {
	o.UniqueId.Set(&v)
}
// SetUniqueIdNil sets the value for UniqueId to be an explicit nil
func (o *MicrosoftGraphLocationConstraintItem) SetUniqueIdNil() {
	o.UniqueId.Set(nil)
}

// UnsetUniqueId ensures that no value is present for UniqueId, not even an explicit nil
func (o *MicrosoftGraphLocationConstraintItem) UnsetUniqueId() {
	o.UniqueId.Unset()
}

// GetUniqueIdType returns the UniqueIdType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphLocationConstraintItem) GetUniqueIdType() AnyOfmicrosoftGraphLocationUniqueIdType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphLocationUniqueIdType
		return ret
	}
	return o.UniqueIdType
}

// GetUniqueIdTypeOk returns a tuple with the UniqueIdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphLocationConstraintItem) GetUniqueIdTypeOk() (*AnyOfmicrosoftGraphLocationUniqueIdType, bool) {
	if o == nil || o.UniqueIdType == nil {
		return nil, false
	}
	return &o.UniqueIdType, true
}

// HasUniqueIdType returns a boolean if a field has been set.
func (o *MicrosoftGraphLocationConstraintItem) HasUniqueIdType() bool {
	if o != nil && o.UniqueIdType != nil {
		return true
	}

	return false
}

// SetUniqueIdType gets a reference to the given AnyOfmicrosoftGraphLocationUniqueIdType and assigns it to the UniqueIdType field.
func (o *MicrosoftGraphLocationConstraintItem) SetUniqueIdType(v AnyOfmicrosoftGraphLocationUniqueIdType) {
	o.UniqueIdType = v
}

// GetResolveAvailability returns the ResolveAvailability field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphLocationConstraintItem) GetResolveAvailability() bool {
	if o == nil || o.ResolveAvailability.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ResolveAvailability.Get()
}

// GetResolveAvailabilityOk returns a tuple with the ResolveAvailability field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphLocationConstraintItem) GetResolveAvailabilityOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResolveAvailability.Get(), o.ResolveAvailability.IsSet()
}

// HasResolveAvailability returns a boolean if a field has been set.
func (o *MicrosoftGraphLocationConstraintItem) HasResolveAvailability() bool {
	if o != nil && o.ResolveAvailability.IsSet() {
		return true
	}

	return false
}

// SetResolveAvailability gets a reference to the given NullableBool and assigns it to the ResolveAvailability field.
func (o *MicrosoftGraphLocationConstraintItem) SetResolveAvailability(v bool) {
	o.ResolveAvailability.Set(&v)
}
// SetResolveAvailabilityNil sets the value for ResolveAvailability to be an explicit nil
func (o *MicrosoftGraphLocationConstraintItem) SetResolveAvailabilityNil() {
	o.ResolveAvailability.Set(nil)
}

// UnsetResolveAvailability ensures that no value is present for ResolveAvailability, not even an explicit nil
func (o *MicrosoftGraphLocationConstraintItem) UnsetResolveAvailability() {
	o.ResolveAvailability.Unset()
}

func (o MicrosoftGraphLocationConstraintItem) MarshalJSON() ([]byte, error) {
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
	if o.ResolveAvailability.IsSet() {
		toSerialize["resolveAvailability"] = o.ResolveAvailability.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphLocationConstraintItem struct {
	value *MicrosoftGraphLocationConstraintItem
	isSet bool
}

func (v NullableMicrosoftGraphLocationConstraintItem) Get() *MicrosoftGraphLocationConstraintItem {
	return v.value
}

func (v *NullableMicrosoftGraphLocationConstraintItem) Set(val *MicrosoftGraphLocationConstraintItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphLocationConstraintItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphLocationConstraintItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphLocationConstraintItem(val *MicrosoftGraphLocationConstraintItem) *NullableMicrosoftGraphLocationConstraintItem {
	return &NullableMicrosoftGraphLocationConstraintItem{value: val, isSet: true}
}

func (v NullableMicrosoftGraphLocationConstraintItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphLocationConstraintItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


