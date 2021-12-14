# MicrosoftGraphSignInLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **NullableString** | Provides the city where the sign-in originated. This is calculated using latitude/longitude information from the sign-in activity. | [optional] 
**CountryOrRegion** | Pointer to **NullableString** | Provides the country code info (2 letter code) where the sign-in originated.  This is calculated using latitude/longitude information from the sign-in activity. | [optional] 
**GeoCoordinates** | Pointer to [**AnyOfmicrosoftGraphGeoCoordinates**](anyOf&lt;microsoft.graph.geoCoordinates&gt;.md) | Provides the latitude, longitude and altitude where the sign-in originated. | [optional] 
**State** | Pointer to **NullableString** | Provides the State where the sign-in originated. This is calculated using latitude/longitude information from the sign-in activity. | [optional] 

## Methods

### NewMicrosoftGraphSignInLocation

`func NewMicrosoftGraphSignInLocation() *MicrosoftGraphSignInLocation`

NewMicrosoftGraphSignInLocation instantiates a new MicrosoftGraphSignInLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSignInLocationWithDefaults

`func NewMicrosoftGraphSignInLocationWithDefaults() *MicrosoftGraphSignInLocation`

NewMicrosoftGraphSignInLocationWithDefaults instantiates a new MicrosoftGraphSignInLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *MicrosoftGraphSignInLocation) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *MicrosoftGraphSignInLocation) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *MicrosoftGraphSignInLocation) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *MicrosoftGraphSignInLocation) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *MicrosoftGraphSignInLocation) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *MicrosoftGraphSignInLocation) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCountryOrRegion

`func (o *MicrosoftGraphSignInLocation) GetCountryOrRegion() string`

GetCountryOrRegion returns the CountryOrRegion field if non-nil, zero value otherwise.

### GetCountryOrRegionOk

`func (o *MicrosoftGraphSignInLocation) GetCountryOrRegionOk() (*string, bool)`

GetCountryOrRegionOk returns a tuple with the CountryOrRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOrRegion

`func (o *MicrosoftGraphSignInLocation) SetCountryOrRegion(v string)`

SetCountryOrRegion sets CountryOrRegion field to given value.

### HasCountryOrRegion

`func (o *MicrosoftGraphSignInLocation) HasCountryOrRegion() bool`

HasCountryOrRegion returns a boolean if a field has been set.

### SetCountryOrRegionNil

`func (o *MicrosoftGraphSignInLocation) SetCountryOrRegionNil(b bool)`

 SetCountryOrRegionNil sets the value for CountryOrRegion to be an explicit nil

### UnsetCountryOrRegion
`func (o *MicrosoftGraphSignInLocation) UnsetCountryOrRegion()`

UnsetCountryOrRegion ensures that no value is present for CountryOrRegion, not even an explicit nil
### GetGeoCoordinates

`func (o *MicrosoftGraphSignInLocation) GetGeoCoordinates() AnyOfmicrosoftGraphGeoCoordinates`

GetGeoCoordinates returns the GeoCoordinates field if non-nil, zero value otherwise.

### GetGeoCoordinatesOk

`func (o *MicrosoftGraphSignInLocation) GetGeoCoordinatesOk() (*AnyOfmicrosoftGraphGeoCoordinates, bool)`

GetGeoCoordinatesOk returns a tuple with the GeoCoordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoCoordinates

`func (o *MicrosoftGraphSignInLocation) SetGeoCoordinates(v AnyOfmicrosoftGraphGeoCoordinates)`

SetGeoCoordinates sets GeoCoordinates field to given value.

### HasGeoCoordinates

`func (o *MicrosoftGraphSignInLocation) HasGeoCoordinates() bool`

HasGeoCoordinates returns a boolean if a field has been set.

### SetGeoCoordinatesNil

`func (o *MicrosoftGraphSignInLocation) SetGeoCoordinatesNil(b bool)`

 SetGeoCoordinatesNil sets the value for GeoCoordinates to be an explicit nil

### UnsetGeoCoordinates
`func (o *MicrosoftGraphSignInLocation) UnsetGeoCoordinates()`

UnsetGeoCoordinates ensures that no value is present for GeoCoordinates, not even an explicit nil
### GetState

`func (o *MicrosoftGraphSignInLocation) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphSignInLocation) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphSignInLocation) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphSignInLocation) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphSignInLocation) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphSignInLocation) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


