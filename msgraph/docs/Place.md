# Place

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**AnyOfmicrosoftGraphPhysicalAddress**](anyOf&lt;microsoft.graph.physicalAddress&gt;.md) | The street address of the place. | [optional] 
**DisplayName** | Pointer to **string** | The name associated with the place. | [optional] 
**GeoCoordinates** | Pointer to [**AnyOfmicrosoftGraphOutlookGeoCoordinates**](anyOf&lt;microsoft.graph.outlookGeoCoordinates&gt;.md) | Specifies the place location in latitude, longitude and (optionally) altitude coordinates. | [optional] 
**Phone** | Pointer to **NullableString** | The phone number of the place. | [optional] 

## Methods

### NewPlace

`func NewPlace() *Place`

NewPlace instantiates a new Place object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaceWithDefaults

`func NewPlaceWithDefaults() *Place`

NewPlaceWithDefaults instantiates a new Place object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Place) GetAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Place) GetAddressOk() (*AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Place) SetAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Place) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *Place) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *Place) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetDisplayName

`func (o *Place) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Place) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Place) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Place) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetGeoCoordinates

`func (o *Place) GetGeoCoordinates() AnyOfmicrosoftGraphOutlookGeoCoordinates`

GetGeoCoordinates returns the GeoCoordinates field if non-nil, zero value otherwise.

### GetGeoCoordinatesOk

`func (o *Place) GetGeoCoordinatesOk() (*AnyOfmicrosoftGraphOutlookGeoCoordinates, bool)`

GetGeoCoordinatesOk returns a tuple with the GeoCoordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoCoordinates

`func (o *Place) SetGeoCoordinates(v AnyOfmicrosoftGraphOutlookGeoCoordinates)`

SetGeoCoordinates sets GeoCoordinates field to given value.

### HasGeoCoordinates

`func (o *Place) HasGeoCoordinates() bool`

HasGeoCoordinates returns a boolean if a field has been set.

### SetGeoCoordinatesNil

`func (o *Place) SetGeoCoordinatesNil(b bool)`

 SetGeoCoordinatesNil sets the value for GeoCoordinates to be an explicit nil

### UnsetGeoCoordinates
`func (o *Place) UnsetGeoCoordinates()`

UnsetGeoCoordinates ensures that no value is present for GeoCoordinates, not even an explicit nil
### GetPhone

`func (o *Place) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Place) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Place) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Place) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *Place) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *Place) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


