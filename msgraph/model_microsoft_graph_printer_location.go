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

// MicrosoftGraphPrinterLocation struct for MicrosoftGraphPrinterLocation
type MicrosoftGraphPrinterLocation struct {
	// The altitude, in meters, that the printer is located at.
	AltitudeInMeters NullableInt32 `json:"altitudeInMeters,omitempty"`
	// The building that the printer is located in.
	Building NullableString `json:"building,omitempty"`
	// The city that the printer is located in.
	City NullableString `json:"city,omitempty"`
	// The country or region that the printer is located in.
	CountryOrRegion NullableString `json:"countryOrRegion,omitempty"`
	// The floor that the printer is located on. Only numerical values are supported right now.
	Floor NullableString `json:"floor,omitempty"`
	// The description of the floor that the printer is located on.
	FloorDescription NullableString `json:"floorDescription,omitempty"`
	// The latitude that the printer is located at.
	Latitude AnyOfnumberstringstring `json:"latitude,omitempty"`
	// The longitude that the printer is located at.
	Longitude AnyOfnumberstringstring `json:"longitude,omitempty"`
	// The organizational hierarchy that the printer belongs to. The elements should be in hierarchical order.
	Organization *[]*string `json:"organization,omitempty"`
	// The postal code that the printer is located in.
	PostalCode NullableString `json:"postalCode,omitempty"`
	// The description of the room that the printer is located in.
	RoomDescription NullableString `json:"roomDescription,omitempty"`
	// The room that the printer is located in. Only numerical values are supported right now.
	RoomName NullableString `json:"roomName,omitempty"`
	// The site that the printer is located in.
	Site NullableString `json:"site,omitempty"`
	// The state or province that the printer is located in.
	StateOrProvince NullableString `json:"stateOrProvince,omitempty"`
	// The street address where the printer is located.
	StreetAddress NullableString `json:"streetAddress,omitempty"`
	// The subdivision that the printer is located in. The elements should be in hierarchical order.
	Subdivision *[]*string `json:"subdivision,omitempty"`
	Subunit *[]*string `json:"subunit,omitempty"`
}

// NewMicrosoftGraphPrinterLocation instantiates a new MicrosoftGraphPrinterLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPrinterLocation() *MicrosoftGraphPrinterLocation {
	this := MicrosoftGraphPrinterLocation{}
	return &this
}

// NewMicrosoftGraphPrinterLocationWithDefaults instantiates a new MicrosoftGraphPrinterLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPrinterLocationWithDefaults() *MicrosoftGraphPrinterLocation {
	this := MicrosoftGraphPrinterLocation{}
	return &this
}

// GetAltitudeInMeters returns the AltitudeInMeters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrinterLocation) GetAltitudeInMeters() int32 {
	if o == nil || o.AltitudeInMeters.Get() == nil {
		var ret int32
		return ret
	}
	return *o.AltitudeInMeters.Get()
}

// GetAltitudeInMetersOk returns a tuple with the AltitudeInMeters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrinterLocation) GetAltitudeInMetersOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AltitudeInMeters.Get(), o.AltitudeInMeters.IsSet()
}

// HasAltitudeInMeters returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterLocation) HasAltitudeInMeters() bool {
	if o != nil && o.AltitudeInMeters.IsSet() {
		return true
	}

	return false
}

// SetAltitudeInMeters gets a reference to the given NullableInt32 and assigns it to the AltitudeInMeters field.
func (o *MicrosoftGraphPrinterLocation) SetAltitudeInMeters(v int32) {
	o.AltitudeInMeters.Set(&v)
}
// SetAltitudeInMetersNil sets the value for AltitudeInMeters to be an explicit nil
func (o *MicrosoftGraphPrinterLocation) SetAltitudeInMetersNil() {
	o.AltitudeInMeters.Set(nil)
}

// UnsetAltitudeInMeters ensures that no value is present for AltitudeInMeters, not even an explicit nil
func (o *MicrosoftGraphPrinterLocation) UnsetAltitudeInMeters() {
	o.AltitudeInMeters.Unset()
}

// GetBuilding returns the Building field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrinterLocation) GetBuilding() string {
	if o == nil || o.Building.Get() == nil {
		var ret string
		return ret
	}
	return *o.Building.Get()
}

// GetBuildingOk returns a tuple with the Building field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrinterLocation) GetBuildingOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Building.Get(), o.Building.IsSet()
}

// HasBuilding returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterLocation) HasBuilding() bool {
	if o != nil && o.Building.IsSet() {
		return true
	}

	return false
}

// SetBuilding gets a reference to the given NullableString and assigns it to the Building field.
func (o *MicrosoftGraphPrinterLocation) SetBuilding(v string) {
	o.Building.Set(&v)
}
// SetBuildingNil sets the value for Building to be an explicit nil
func (o *MicrosoftGraphPrinterLocation) SetBuildingNil() {
	o.Building.Set(nil)
}

// UnsetBuilding ensures that no value is present for Building, not even an explicit nil
func (o *MicrosoftGraphPrinterLocation) UnsetBuilding() {
	o.Building.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrinterLocation) GetCity() string {
	if o == nil || o.City.Get() == nil {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrinterLocation) GetCityOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterLocation) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *MicrosoftGraphPrinterLocation) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *MicrosoftGraphPrinterLocation) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *MicrosoftGraphPrinterLocation) UnsetCity() {
	o.City.Unset()
}

// GetCountryOrRegion returns the CountryOrRegion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrinterLocation) GetCountryOrRegion() string {
	if o == nil || o.CountryOrRegion.Get() == nil {
		var ret string
		return ret
	}
	return *o.CountryOrRegion.Get()
}

// GetCountryOrRegionOk returns a tuple with the CountryOrRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrinterLocation) GetCountryOrRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CountryOrRegion.Get(), o.CountryOrRegion.IsSet()
}

// HasCountryOrRegion returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterLocation) HasCountryOrRegion() bool {
	if o != nil && o.CountryOrRegion.IsSet() {
		return true
	}

	return false
}

// SetCountryOrRegion gets a reference to the given NullableString and assigns it to the CountryOrRegion field.
func (o *MicrosoftGraphPrinterLocation) SetCountryOrRegion(v string) {
	o.CountryOrRegion.Set(&v)
}
// SetCountryOrRegionNil sets the value for CountryOrRegion to be an explicit nil
func (o *MicrosoftGraphPrinterLocation) SetCountryOrRegionNil() {
	o.CountryOrRegion.Set(nil)
}

// UnsetCountryOrRegion ensures that no value is present for CountryOrRegion, not even an explicit nil
func (o *MicrosoftGraphPrinterLocation) UnsetCountryOrRegion() {
	o.CountryOrRegion.Unset()
}

// GetFloor returns the Floor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrinterLocation) GetFloor() string {
	if o == nil || o.Floor.Get() == nil {
		var ret string
		return ret
	}
	return *o.Floor.Get()
}

// GetFloorOk returns a tuple with the Floor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrinterLocation) GetFloorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Floor.Get(), o.Floor.IsSet()
}

// HasFloor returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterLocation) HasFloor() bool {
	if o != nil && o.Floor.IsSet() {
		return true
	}

	return false
}

// SetFloor gets a reference to the given NullableString and assigns it to the Floor field.
func (o *MicrosoftGraphPrinterLocation) SetFloor(v string) {
	o.Floor.Set(&v)
}
// SetFloorNil sets the value for Floor to be an explicit nil
func (o *MicrosoftGraphPrinterLocation) SetFloorNil() {
	o.Floor.Set(nil)
}

// UnsetFloor ensures that no value is present for Floor, not even an explicit nil
func (o *MicrosoftGraphPrinterLocation) UnsetFloor() {
	o.Floor.Unset()
}

// GetFloorDescription returns the FloorDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrinterLocation) GetFloorDescription() string {
	if o == nil || o.FloorDescription.Get() == nil {
		var ret string
		return ret
	}
	return *o.FloorDescription.Get()
}

// GetFloorDescriptionOk returns a tuple with the FloorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrinterLocation) GetFloorDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FloorDescription.Get(), o.FloorDescription.IsSet()
}

// HasFloorDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterLocation) HasFloorDescription() bool {
	if o != nil && o.FloorDescription.IsSet() {
		return true
	}

	return false
}

// SetFloorDescription gets a reference to the given NullableString and assigns it to the FloorDescription field.
func (o *MicrosoftGraphPrinterLocation) SetFloorDescription(v string) {
	o.FloorDescription.Set(&v)
}
// SetFloorDescriptionNil sets the value for FloorDescription to be an explicit nil
func (o *MicrosoftGraphPrinterLocation) SetFloorDescriptionNil() {
	o.FloorDescription.Set(nil)
}

// UnsetFloorDescription ensures that no value is present for FloorDescription, not even an explicit nil
func (o *MicrosoftGraphPrinterLocation) UnsetFloorDescription() {
	o.FloorDescription.Unset()
}

// GetLatitude returns the Latitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrinterLocation) GetLatitude() AnyOfnumberstringstring {
	if o == nil  {
		var ret AnyOfnumberstringstring
		return ret
	}
	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrinterLocation) GetLatitudeOk() (*AnyOfnumberstringstring, bool) {
	if o == nil || o.Latitude == nil {
		return nil, false
	}
	return &o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterLocation) HasLatitude() bool {
	if o != nil && o.Latitude != nil {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given AnyOfnumberstringstring and assigns it to the Latitude field.
func (o *MicrosoftGraphPrinterLocation) SetLatitude(v AnyOfnumberstringstring) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrinterLocation) GetLongitude() AnyOfnumberstringstring {
	if o == nil  {
		var ret AnyOfnumberstringstring
		return ret
	}
	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrinterLocation) GetLongitudeOk() (*AnyOfnumberstringstring, bool) {
	if o == nil || o.Longitude == nil {
		return nil, false
	}
	return &o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterLocation) HasLongitude() bool {
	if o != nil && o.Longitude != nil {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given AnyOfnumberstringstring and assigns it to the Longitude field.
func (o *MicrosoftGraphPrinterLocation) SetLongitude(v AnyOfnumberstringstring) {
	o.Longitude = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *MicrosoftGraphPrinterLocation) GetOrganization() []*string {
	if o == nil || o.Organization == nil {
		var ret []*string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrinterLocation) GetOrganizationOk() (*[]*string, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterLocation) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given []*string and assigns it to the Organization field.
func (o *MicrosoftGraphPrinterLocation) SetOrganization(v []*string) {
	o.Organization = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrinterLocation) GetPostalCode() string {
	if o == nil || o.PostalCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrinterLocation) GetPostalCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// HasPostalCode returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterLocation) HasPostalCode() bool {
	if o != nil && o.PostalCode.IsSet() {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given NullableString and assigns it to the PostalCode field.
func (o *MicrosoftGraphPrinterLocation) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}
// SetPostalCodeNil sets the value for PostalCode to be an explicit nil
func (o *MicrosoftGraphPrinterLocation) SetPostalCodeNil() {
	o.PostalCode.Set(nil)
}

// UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
func (o *MicrosoftGraphPrinterLocation) UnsetPostalCode() {
	o.PostalCode.Unset()
}

// GetRoomDescription returns the RoomDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrinterLocation) GetRoomDescription() string {
	if o == nil || o.RoomDescription.Get() == nil {
		var ret string
		return ret
	}
	return *o.RoomDescription.Get()
}

// GetRoomDescriptionOk returns a tuple with the RoomDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrinterLocation) GetRoomDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RoomDescription.Get(), o.RoomDescription.IsSet()
}

// HasRoomDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterLocation) HasRoomDescription() bool {
	if o != nil && o.RoomDescription.IsSet() {
		return true
	}

	return false
}

// SetRoomDescription gets a reference to the given NullableString and assigns it to the RoomDescription field.
func (o *MicrosoftGraphPrinterLocation) SetRoomDescription(v string) {
	o.RoomDescription.Set(&v)
}
// SetRoomDescriptionNil sets the value for RoomDescription to be an explicit nil
func (o *MicrosoftGraphPrinterLocation) SetRoomDescriptionNil() {
	o.RoomDescription.Set(nil)
}

// UnsetRoomDescription ensures that no value is present for RoomDescription, not even an explicit nil
func (o *MicrosoftGraphPrinterLocation) UnsetRoomDescription() {
	o.RoomDescription.Unset()
}

// GetRoomName returns the RoomName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrinterLocation) GetRoomName() string {
	if o == nil || o.RoomName.Get() == nil {
		var ret string
		return ret
	}
	return *o.RoomName.Get()
}

// GetRoomNameOk returns a tuple with the RoomName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrinterLocation) GetRoomNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RoomName.Get(), o.RoomName.IsSet()
}

// HasRoomName returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterLocation) HasRoomName() bool {
	if o != nil && o.RoomName.IsSet() {
		return true
	}

	return false
}

// SetRoomName gets a reference to the given NullableString and assigns it to the RoomName field.
func (o *MicrosoftGraphPrinterLocation) SetRoomName(v string) {
	o.RoomName.Set(&v)
}
// SetRoomNameNil sets the value for RoomName to be an explicit nil
func (o *MicrosoftGraphPrinterLocation) SetRoomNameNil() {
	o.RoomName.Set(nil)
}

// UnsetRoomName ensures that no value is present for RoomName, not even an explicit nil
func (o *MicrosoftGraphPrinterLocation) UnsetRoomName() {
	o.RoomName.Unset()
}

// GetSite returns the Site field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrinterLocation) GetSite() string {
	if o == nil || o.Site.Get() == nil {
		var ret string
		return ret
	}
	return *o.Site.Get()
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrinterLocation) GetSiteOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Site.Get(), o.Site.IsSet()
}

// HasSite returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterLocation) HasSite() bool {
	if o != nil && o.Site.IsSet() {
		return true
	}

	return false
}

// SetSite gets a reference to the given NullableString and assigns it to the Site field.
func (o *MicrosoftGraphPrinterLocation) SetSite(v string) {
	o.Site.Set(&v)
}
// SetSiteNil sets the value for Site to be an explicit nil
func (o *MicrosoftGraphPrinterLocation) SetSiteNil() {
	o.Site.Set(nil)
}

// UnsetSite ensures that no value is present for Site, not even an explicit nil
func (o *MicrosoftGraphPrinterLocation) UnsetSite() {
	o.Site.Unset()
}

// GetStateOrProvince returns the StateOrProvince field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrinterLocation) GetStateOrProvince() string {
	if o == nil || o.StateOrProvince.Get() == nil {
		var ret string
		return ret
	}
	return *o.StateOrProvince.Get()
}

// GetStateOrProvinceOk returns a tuple with the StateOrProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrinterLocation) GetStateOrProvinceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StateOrProvince.Get(), o.StateOrProvince.IsSet()
}

// HasStateOrProvince returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterLocation) HasStateOrProvince() bool {
	if o != nil && o.StateOrProvince.IsSet() {
		return true
	}

	return false
}

// SetStateOrProvince gets a reference to the given NullableString and assigns it to the StateOrProvince field.
func (o *MicrosoftGraphPrinterLocation) SetStateOrProvince(v string) {
	o.StateOrProvince.Set(&v)
}
// SetStateOrProvinceNil sets the value for StateOrProvince to be an explicit nil
func (o *MicrosoftGraphPrinterLocation) SetStateOrProvinceNil() {
	o.StateOrProvince.Set(nil)
}

// UnsetStateOrProvince ensures that no value is present for StateOrProvince, not even an explicit nil
func (o *MicrosoftGraphPrinterLocation) UnsetStateOrProvince() {
	o.StateOrProvince.Unset()
}

// GetStreetAddress returns the StreetAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrinterLocation) GetStreetAddress() string {
	if o == nil || o.StreetAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.StreetAddress.Get()
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrinterLocation) GetStreetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StreetAddress.Get(), o.StreetAddress.IsSet()
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterLocation) HasStreetAddress() bool {
	if o != nil && o.StreetAddress.IsSet() {
		return true
	}

	return false
}

// SetStreetAddress gets a reference to the given NullableString and assigns it to the StreetAddress field.
func (o *MicrosoftGraphPrinterLocation) SetStreetAddress(v string) {
	o.StreetAddress.Set(&v)
}
// SetStreetAddressNil sets the value for StreetAddress to be an explicit nil
func (o *MicrosoftGraphPrinterLocation) SetStreetAddressNil() {
	o.StreetAddress.Set(nil)
}

// UnsetStreetAddress ensures that no value is present for StreetAddress, not even an explicit nil
func (o *MicrosoftGraphPrinterLocation) UnsetStreetAddress() {
	o.StreetAddress.Unset()
}

// GetSubdivision returns the Subdivision field value if set, zero value otherwise.
func (o *MicrosoftGraphPrinterLocation) GetSubdivision() []*string {
	if o == nil || o.Subdivision == nil {
		var ret []*string
		return ret
	}
	return *o.Subdivision
}

// GetSubdivisionOk returns a tuple with the Subdivision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrinterLocation) GetSubdivisionOk() (*[]*string, bool) {
	if o == nil || o.Subdivision == nil {
		return nil, false
	}
	return o.Subdivision, true
}

// HasSubdivision returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterLocation) HasSubdivision() bool {
	if o != nil && o.Subdivision != nil {
		return true
	}

	return false
}

// SetSubdivision gets a reference to the given []*string and assigns it to the Subdivision field.
func (o *MicrosoftGraphPrinterLocation) SetSubdivision(v []*string) {
	o.Subdivision = &v
}

// GetSubunit returns the Subunit field value if set, zero value otherwise.
func (o *MicrosoftGraphPrinterLocation) GetSubunit() []*string {
	if o == nil || o.Subunit == nil {
		var ret []*string
		return ret
	}
	return *o.Subunit
}

// GetSubunitOk returns a tuple with the Subunit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrinterLocation) GetSubunitOk() (*[]*string, bool) {
	if o == nil || o.Subunit == nil {
		return nil, false
	}
	return o.Subunit, true
}

// HasSubunit returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterLocation) HasSubunit() bool {
	if o != nil && o.Subunit != nil {
		return true
	}

	return false
}

// SetSubunit gets a reference to the given []*string and assigns it to the Subunit field.
func (o *MicrosoftGraphPrinterLocation) SetSubunit(v []*string) {
	o.Subunit = &v
}

func (o MicrosoftGraphPrinterLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AltitudeInMeters.IsSet() {
		toSerialize["altitudeInMeters"] = o.AltitudeInMeters.Get()
	}
	if o.Building.IsSet() {
		toSerialize["building"] = o.Building.Get()
	}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if o.CountryOrRegion.IsSet() {
		toSerialize["countryOrRegion"] = o.CountryOrRegion.Get()
	}
	if o.Floor.IsSet() {
		toSerialize["floor"] = o.Floor.Get()
	}
	if o.FloorDescription.IsSet() {
		toSerialize["floorDescription"] = o.FloorDescription.Get()
	}
	if o.Latitude != nil {
		toSerialize["latitude"] = o.Latitude
	}
	if o.Longitude != nil {
		toSerialize["longitude"] = o.Longitude
	}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	if o.PostalCode.IsSet() {
		toSerialize["postalCode"] = o.PostalCode.Get()
	}
	if o.RoomDescription.IsSet() {
		toSerialize["roomDescription"] = o.RoomDescription.Get()
	}
	if o.RoomName.IsSet() {
		toSerialize["roomName"] = o.RoomName.Get()
	}
	if o.Site.IsSet() {
		toSerialize["site"] = o.Site.Get()
	}
	if o.StateOrProvince.IsSet() {
		toSerialize["stateOrProvince"] = o.StateOrProvince.Get()
	}
	if o.StreetAddress.IsSet() {
		toSerialize["streetAddress"] = o.StreetAddress.Get()
	}
	if o.Subdivision != nil {
		toSerialize["subdivision"] = o.Subdivision
	}
	if o.Subunit != nil {
		toSerialize["subunit"] = o.Subunit
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPrinterLocation struct {
	value *MicrosoftGraphPrinterLocation
	isSet bool
}

func (v NullableMicrosoftGraphPrinterLocation) Get() *MicrosoftGraphPrinterLocation {
	return v.value
}

func (v *NullableMicrosoftGraphPrinterLocation) Set(val *MicrosoftGraphPrinterLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPrinterLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPrinterLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPrinterLocation(val *MicrosoftGraphPrinterLocation) *NullableMicrosoftGraphPrinterLocation {
	return &NullableMicrosoftGraphPrinterLocation{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPrinterLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPrinterLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

