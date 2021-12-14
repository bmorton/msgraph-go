# MicrosoftGraphPrinterLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AltitudeInMeters** | Pointer to **NullableInt32** | The altitude, in meters, that the printer is located at. | [optional] 
**Building** | Pointer to **NullableString** | The building that the printer is located in. | [optional] 
**City** | Pointer to **NullableString** | The city that the printer is located in. | [optional] 
**CountryOrRegion** | Pointer to **NullableString** | The country or region that the printer is located in. | [optional] 
**Floor** | Pointer to **NullableString** | The floor that the printer is located on. Only numerical values are supported right now. | [optional] 
**FloorDescription** | Pointer to **NullableString** | The description of the floor that the printer is located on. | [optional] 
**Latitude** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | The latitude that the printer is located at. | [optional] 
**Longitude** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | The longitude that the printer is located at. | [optional] 
**Organization** | Pointer to **[]string** | The organizational hierarchy that the printer belongs to. The elements should be in hierarchical order. | [optional] 
**PostalCode** | Pointer to **NullableString** | The postal code that the printer is located in. | [optional] 
**RoomDescription** | Pointer to **NullableString** | The description of the room that the printer is located in. | [optional] 
**RoomName** | Pointer to **NullableString** | The room that the printer is located in. Only numerical values are supported right now. | [optional] 
**Site** | Pointer to **NullableString** | The site that the printer is located in. | [optional] 
**StateOrProvince** | Pointer to **NullableString** | The state or province that the printer is located in. | [optional] 
**StreetAddress** | Pointer to **NullableString** | The street address where the printer is located. | [optional] 
**Subdivision** | Pointer to **[]string** | The subdivision that the printer is located in. The elements should be in hierarchical order. | [optional] 
**Subunit** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMicrosoftGraphPrinterLocation

`func NewMicrosoftGraphPrinterLocation() *MicrosoftGraphPrinterLocation`

NewMicrosoftGraphPrinterLocation instantiates a new MicrosoftGraphPrinterLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPrinterLocationWithDefaults

`func NewMicrosoftGraphPrinterLocationWithDefaults() *MicrosoftGraphPrinterLocation`

NewMicrosoftGraphPrinterLocationWithDefaults instantiates a new MicrosoftGraphPrinterLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAltitudeInMeters

`func (o *MicrosoftGraphPrinterLocation) GetAltitudeInMeters() int32`

GetAltitudeInMeters returns the AltitudeInMeters field if non-nil, zero value otherwise.

### GetAltitudeInMetersOk

`func (o *MicrosoftGraphPrinterLocation) GetAltitudeInMetersOk() (*int32, bool)`

GetAltitudeInMetersOk returns a tuple with the AltitudeInMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitudeInMeters

`func (o *MicrosoftGraphPrinterLocation) SetAltitudeInMeters(v int32)`

SetAltitudeInMeters sets AltitudeInMeters field to given value.

### HasAltitudeInMeters

`func (o *MicrosoftGraphPrinterLocation) HasAltitudeInMeters() bool`

HasAltitudeInMeters returns a boolean if a field has been set.

### SetAltitudeInMetersNil

`func (o *MicrosoftGraphPrinterLocation) SetAltitudeInMetersNil(b bool)`

 SetAltitudeInMetersNil sets the value for AltitudeInMeters to be an explicit nil

### UnsetAltitudeInMeters
`func (o *MicrosoftGraphPrinterLocation) UnsetAltitudeInMeters()`

UnsetAltitudeInMeters ensures that no value is present for AltitudeInMeters, not even an explicit nil
### GetBuilding

`func (o *MicrosoftGraphPrinterLocation) GetBuilding() string`

GetBuilding returns the Building field if non-nil, zero value otherwise.

### GetBuildingOk

`func (o *MicrosoftGraphPrinterLocation) GetBuildingOk() (*string, bool)`

GetBuildingOk returns a tuple with the Building field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilding

`func (o *MicrosoftGraphPrinterLocation) SetBuilding(v string)`

SetBuilding sets Building field to given value.

### HasBuilding

`func (o *MicrosoftGraphPrinterLocation) HasBuilding() bool`

HasBuilding returns a boolean if a field has been set.

### SetBuildingNil

`func (o *MicrosoftGraphPrinterLocation) SetBuildingNil(b bool)`

 SetBuildingNil sets the value for Building to be an explicit nil

### UnsetBuilding
`func (o *MicrosoftGraphPrinterLocation) UnsetBuilding()`

UnsetBuilding ensures that no value is present for Building, not even an explicit nil
### GetCity

`func (o *MicrosoftGraphPrinterLocation) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *MicrosoftGraphPrinterLocation) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *MicrosoftGraphPrinterLocation) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *MicrosoftGraphPrinterLocation) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *MicrosoftGraphPrinterLocation) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *MicrosoftGraphPrinterLocation) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCountryOrRegion

`func (o *MicrosoftGraphPrinterLocation) GetCountryOrRegion() string`

GetCountryOrRegion returns the CountryOrRegion field if non-nil, zero value otherwise.

### GetCountryOrRegionOk

`func (o *MicrosoftGraphPrinterLocation) GetCountryOrRegionOk() (*string, bool)`

GetCountryOrRegionOk returns a tuple with the CountryOrRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOrRegion

`func (o *MicrosoftGraphPrinterLocation) SetCountryOrRegion(v string)`

SetCountryOrRegion sets CountryOrRegion field to given value.

### HasCountryOrRegion

`func (o *MicrosoftGraphPrinterLocation) HasCountryOrRegion() bool`

HasCountryOrRegion returns a boolean if a field has been set.

### SetCountryOrRegionNil

`func (o *MicrosoftGraphPrinterLocation) SetCountryOrRegionNil(b bool)`

 SetCountryOrRegionNil sets the value for CountryOrRegion to be an explicit nil

### UnsetCountryOrRegion
`func (o *MicrosoftGraphPrinterLocation) UnsetCountryOrRegion()`

UnsetCountryOrRegion ensures that no value is present for CountryOrRegion, not even an explicit nil
### GetFloor

`func (o *MicrosoftGraphPrinterLocation) GetFloor() string`

GetFloor returns the Floor field if non-nil, zero value otherwise.

### GetFloorOk

`func (o *MicrosoftGraphPrinterLocation) GetFloorOk() (*string, bool)`

GetFloorOk returns a tuple with the Floor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloor

`func (o *MicrosoftGraphPrinterLocation) SetFloor(v string)`

SetFloor sets Floor field to given value.

### HasFloor

`func (o *MicrosoftGraphPrinterLocation) HasFloor() bool`

HasFloor returns a boolean if a field has been set.

### SetFloorNil

`func (o *MicrosoftGraphPrinterLocation) SetFloorNil(b bool)`

 SetFloorNil sets the value for Floor to be an explicit nil

### UnsetFloor
`func (o *MicrosoftGraphPrinterLocation) UnsetFloor()`

UnsetFloor ensures that no value is present for Floor, not even an explicit nil
### GetFloorDescription

`func (o *MicrosoftGraphPrinterLocation) GetFloorDescription() string`

GetFloorDescription returns the FloorDescription field if non-nil, zero value otherwise.

### GetFloorDescriptionOk

`func (o *MicrosoftGraphPrinterLocation) GetFloorDescriptionOk() (*string, bool)`

GetFloorDescriptionOk returns a tuple with the FloorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorDescription

`func (o *MicrosoftGraphPrinterLocation) SetFloorDescription(v string)`

SetFloorDescription sets FloorDescription field to given value.

### HasFloorDescription

`func (o *MicrosoftGraphPrinterLocation) HasFloorDescription() bool`

HasFloorDescription returns a boolean if a field has been set.

### SetFloorDescriptionNil

`func (o *MicrosoftGraphPrinterLocation) SetFloorDescriptionNil(b bool)`

 SetFloorDescriptionNil sets the value for FloorDescription to be an explicit nil

### UnsetFloorDescription
`func (o *MicrosoftGraphPrinterLocation) UnsetFloorDescription()`

UnsetFloorDescription ensures that no value is present for FloorDescription, not even an explicit nil
### GetLatitude

`func (o *MicrosoftGraphPrinterLocation) GetLatitude() AnyOfnumberstringstring`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *MicrosoftGraphPrinterLocation) GetLatitudeOk() (*AnyOfnumberstringstring, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *MicrosoftGraphPrinterLocation) SetLatitude(v AnyOfnumberstringstring)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *MicrosoftGraphPrinterLocation) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *MicrosoftGraphPrinterLocation) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *MicrosoftGraphPrinterLocation) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *MicrosoftGraphPrinterLocation) GetLongitude() AnyOfnumberstringstring`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *MicrosoftGraphPrinterLocation) GetLongitudeOk() (*AnyOfnumberstringstring, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *MicrosoftGraphPrinterLocation) SetLongitude(v AnyOfnumberstringstring)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *MicrosoftGraphPrinterLocation) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *MicrosoftGraphPrinterLocation) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *MicrosoftGraphPrinterLocation) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetOrganization

`func (o *MicrosoftGraphPrinterLocation) GetOrganization() []*string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *MicrosoftGraphPrinterLocation) GetOrganizationOk() (*[]*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *MicrosoftGraphPrinterLocation) SetOrganization(v []*string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *MicrosoftGraphPrinterLocation) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPostalCode

`func (o *MicrosoftGraphPrinterLocation) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *MicrosoftGraphPrinterLocation) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *MicrosoftGraphPrinterLocation) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *MicrosoftGraphPrinterLocation) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *MicrosoftGraphPrinterLocation) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *MicrosoftGraphPrinterLocation) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetRoomDescription

`func (o *MicrosoftGraphPrinterLocation) GetRoomDescription() string`

GetRoomDescription returns the RoomDescription field if non-nil, zero value otherwise.

### GetRoomDescriptionOk

`func (o *MicrosoftGraphPrinterLocation) GetRoomDescriptionOk() (*string, bool)`

GetRoomDescriptionOk returns a tuple with the RoomDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomDescription

`func (o *MicrosoftGraphPrinterLocation) SetRoomDescription(v string)`

SetRoomDescription sets RoomDescription field to given value.

### HasRoomDescription

`func (o *MicrosoftGraphPrinterLocation) HasRoomDescription() bool`

HasRoomDescription returns a boolean if a field has been set.

### SetRoomDescriptionNil

`func (o *MicrosoftGraphPrinterLocation) SetRoomDescriptionNil(b bool)`

 SetRoomDescriptionNil sets the value for RoomDescription to be an explicit nil

### UnsetRoomDescription
`func (o *MicrosoftGraphPrinterLocation) UnsetRoomDescription()`

UnsetRoomDescription ensures that no value is present for RoomDescription, not even an explicit nil
### GetRoomName

`func (o *MicrosoftGraphPrinterLocation) GetRoomName() string`

GetRoomName returns the RoomName field if non-nil, zero value otherwise.

### GetRoomNameOk

`func (o *MicrosoftGraphPrinterLocation) GetRoomNameOk() (*string, bool)`

GetRoomNameOk returns a tuple with the RoomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoomName

`func (o *MicrosoftGraphPrinterLocation) SetRoomName(v string)`

SetRoomName sets RoomName field to given value.

### HasRoomName

`func (o *MicrosoftGraphPrinterLocation) HasRoomName() bool`

HasRoomName returns a boolean if a field has been set.

### SetRoomNameNil

`func (o *MicrosoftGraphPrinterLocation) SetRoomNameNil(b bool)`

 SetRoomNameNil sets the value for RoomName to be an explicit nil

### UnsetRoomName
`func (o *MicrosoftGraphPrinterLocation) UnsetRoomName()`

UnsetRoomName ensures that no value is present for RoomName, not even an explicit nil
### GetSite

`func (o *MicrosoftGraphPrinterLocation) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *MicrosoftGraphPrinterLocation) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *MicrosoftGraphPrinterLocation) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *MicrosoftGraphPrinterLocation) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *MicrosoftGraphPrinterLocation) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *MicrosoftGraphPrinterLocation) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetStateOrProvince

`func (o *MicrosoftGraphPrinterLocation) GetStateOrProvince() string`

GetStateOrProvince returns the StateOrProvince field if non-nil, zero value otherwise.

### GetStateOrProvinceOk

`func (o *MicrosoftGraphPrinterLocation) GetStateOrProvinceOk() (*string, bool)`

GetStateOrProvinceOk returns a tuple with the StateOrProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOrProvince

`func (o *MicrosoftGraphPrinterLocation) SetStateOrProvince(v string)`

SetStateOrProvince sets StateOrProvince field to given value.

### HasStateOrProvince

`func (o *MicrosoftGraphPrinterLocation) HasStateOrProvince() bool`

HasStateOrProvince returns a boolean if a field has been set.

### SetStateOrProvinceNil

`func (o *MicrosoftGraphPrinterLocation) SetStateOrProvinceNil(b bool)`

 SetStateOrProvinceNil sets the value for StateOrProvince to be an explicit nil

### UnsetStateOrProvince
`func (o *MicrosoftGraphPrinterLocation) UnsetStateOrProvince()`

UnsetStateOrProvince ensures that no value is present for StateOrProvince, not even an explicit nil
### GetStreetAddress

`func (o *MicrosoftGraphPrinterLocation) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *MicrosoftGraphPrinterLocation) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *MicrosoftGraphPrinterLocation) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *MicrosoftGraphPrinterLocation) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### SetStreetAddressNil

`func (o *MicrosoftGraphPrinterLocation) SetStreetAddressNil(b bool)`

 SetStreetAddressNil sets the value for StreetAddress to be an explicit nil

### UnsetStreetAddress
`func (o *MicrosoftGraphPrinterLocation) UnsetStreetAddress()`

UnsetStreetAddress ensures that no value is present for StreetAddress, not even an explicit nil
### GetSubdivision

`func (o *MicrosoftGraphPrinterLocation) GetSubdivision() []*string`

GetSubdivision returns the Subdivision field if non-nil, zero value otherwise.

### GetSubdivisionOk

`func (o *MicrosoftGraphPrinterLocation) GetSubdivisionOk() (*[]*string, bool)`

GetSubdivisionOk returns a tuple with the Subdivision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdivision

`func (o *MicrosoftGraphPrinterLocation) SetSubdivision(v []*string)`

SetSubdivision sets Subdivision field to given value.

### HasSubdivision

`func (o *MicrosoftGraphPrinterLocation) HasSubdivision() bool`

HasSubdivision returns a boolean if a field has been set.

### GetSubunit

`func (o *MicrosoftGraphPrinterLocation) GetSubunit() []*string`

GetSubunit returns the Subunit field if non-nil, zero value otherwise.

### GetSubunitOk

`func (o *MicrosoftGraphPrinterLocation) GetSubunitOk() (*[]*string, bool)`

GetSubunitOk returns a tuple with the Subunit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubunit

`func (o *MicrosoftGraphPrinterLocation) SetSubunit(v []*string)`

SetSubunit sets Subunit field to given value.

### HasSubunit

`func (o *MicrosoftGraphPrinterLocation) HasSubunit() bool`

HasSubunit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


