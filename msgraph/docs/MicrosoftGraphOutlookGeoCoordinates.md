# MicrosoftGraphOutlookGeoCoordinates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accuracy** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | The accuracy of the latitude and longitude. As an example, the accuracy can be measured in meters, such as the latitude and longitude are accurate to within 50 meters. | [optional] 
**Altitude** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | The altitude of the location. | [optional] 
**AltitudeAccuracy** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | The accuracy of the altitude. | [optional] 
**Latitude** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | The latitude of the location. | [optional] 
**Longitude** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | The longitude of the location. | [optional] 

## Methods

### NewMicrosoftGraphOutlookGeoCoordinates

`func NewMicrosoftGraphOutlookGeoCoordinates() *MicrosoftGraphOutlookGeoCoordinates`

NewMicrosoftGraphOutlookGeoCoordinates instantiates a new MicrosoftGraphOutlookGeoCoordinates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphOutlookGeoCoordinatesWithDefaults

`func NewMicrosoftGraphOutlookGeoCoordinatesWithDefaults() *MicrosoftGraphOutlookGeoCoordinates`

NewMicrosoftGraphOutlookGeoCoordinatesWithDefaults instantiates a new MicrosoftGraphOutlookGeoCoordinates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccuracy

`func (o *MicrosoftGraphOutlookGeoCoordinates) GetAccuracy() AnyOfnumberstringstring`

GetAccuracy returns the Accuracy field if non-nil, zero value otherwise.

### GetAccuracyOk

`func (o *MicrosoftGraphOutlookGeoCoordinates) GetAccuracyOk() (*AnyOfnumberstringstring, bool)`

GetAccuracyOk returns a tuple with the Accuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracy

`func (o *MicrosoftGraphOutlookGeoCoordinates) SetAccuracy(v AnyOfnumberstringstring)`

SetAccuracy sets Accuracy field to given value.

### HasAccuracy

`func (o *MicrosoftGraphOutlookGeoCoordinates) HasAccuracy() bool`

HasAccuracy returns a boolean if a field has been set.

### SetAccuracyNil

`func (o *MicrosoftGraphOutlookGeoCoordinates) SetAccuracyNil(b bool)`

 SetAccuracyNil sets the value for Accuracy to be an explicit nil

### UnsetAccuracy
`func (o *MicrosoftGraphOutlookGeoCoordinates) UnsetAccuracy()`

UnsetAccuracy ensures that no value is present for Accuracy, not even an explicit nil
### GetAltitude

`func (o *MicrosoftGraphOutlookGeoCoordinates) GetAltitude() AnyOfnumberstringstring`

GetAltitude returns the Altitude field if non-nil, zero value otherwise.

### GetAltitudeOk

`func (o *MicrosoftGraphOutlookGeoCoordinates) GetAltitudeOk() (*AnyOfnumberstringstring, bool)`

GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitude

`func (o *MicrosoftGraphOutlookGeoCoordinates) SetAltitude(v AnyOfnumberstringstring)`

SetAltitude sets Altitude field to given value.

### HasAltitude

`func (o *MicrosoftGraphOutlookGeoCoordinates) HasAltitude() bool`

HasAltitude returns a boolean if a field has been set.

### SetAltitudeNil

`func (o *MicrosoftGraphOutlookGeoCoordinates) SetAltitudeNil(b bool)`

 SetAltitudeNil sets the value for Altitude to be an explicit nil

### UnsetAltitude
`func (o *MicrosoftGraphOutlookGeoCoordinates) UnsetAltitude()`

UnsetAltitude ensures that no value is present for Altitude, not even an explicit nil
### GetAltitudeAccuracy

`func (o *MicrosoftGraphOutlookGeoCoordinates) GetAltitudeAccuracy() AnyOfnumberstringstring`

GetAltitudeAccuracy returns the AltitudeAccuracy field if non-nil, zero value otherwise.

### GetAltitudeAccuracyOk

`func (o *MicrosoftGraphOutlookGeoCoordinates) GetAltitudeAccuracyOk() (*AnyOfnumberstringstring, bool)`

GetAltitudeAccuracyOk returns a tuple with the AltitudeAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitudeAccuracy

`func (o *MicrosoftGraphOutlookGeoCoordinates) SetAltitudeAccuracy(v AnyOfnumberstringstring)`

SetAltitudeAccuracy sets AltitudeAccuracy field to given value.

### HasAltitudeAccuracy

`func (o *MicrosoftGraphOutlookGeoCoordinates) HasAltitudeAccuracy() bool`

HasAltitudeAccuracy returns a boolean if a field has been set.

### SetAltitudeAccuracyNil

`func (o *MicrosoftGraphOutlookGeoCoordinates) SetAltitudeAccuracyNil(b bool)`

 SetAltitudeAccuracyNil sets the value for AltitudeAccuracy to be an explicit nil

### UnsetAltitudeAccuracy
`func (o *MicrosoftGraphOutlookGeoCoordinates) UnsetAltitudeAccuracy()`

UnsetAltitudeAccuracy ensures that no value is present for AltitudeAccuracy, not even an explicit nil
### GetLatitude

`func (o *MicrosoftGraphOutlookGeoCoordinates) GetLatitude() AnyOfnumberstringstring`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *MicrosoftGraphOutlookGeoCoordinates) GetLatitudeOk() (*AnyOfnumberstringstring, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *MicrosoftGraphOutlookGeoCoordinates) SetLatitude(v AnyOfnumberstringstring)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *MicrosoftGraphOutlookGeoCoordinates) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *MicrosoftGraphOutlookGeoCoordinates) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *MicrosoftGraphOutlookGeoCoordinates) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *MicrosoftGraphOutlookGeoCoordinates) GetLongitude() AnyOfnumberstringstring`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *MicrosoftGraphOutlookGeoCoordinates) GetLongitudeOk() (*AnyOfnumberstringstring, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *MicrosoftGraphOutlookGeoCoordinates) SetLongitude(v AnyOfnumberstringstring)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *MicrosoftGraphOutlookGeoCoordinates) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *MicrosoftGraphOutlookGeoCoordinates) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *MicrosoftGraphOutlookGeoCoordinates) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


