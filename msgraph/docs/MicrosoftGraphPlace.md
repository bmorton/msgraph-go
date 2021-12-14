# MicrosoftGraphPlace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Address** | Pointer to [**AnyOfmicrosoftGraphPhysicalAddress**](anyOf&lt;microsoft.graph.physicalAddress&gt;.md) | The street address of the place. | [optional] 
**DisplayName** | Pointer to **string** | The name associated with the place. | [optional] 
**GeoCoordinates** | Pointer to [**AnyOfmicrosoftGraphOutlookGeoCoordinates**](anyOf&lt;microsoft.graph.outlookGeoCoordinates&gt;.md) | Specifies the place location in latitude, longitude and (optionally) altitude coordinates. | [optional] 
**Phone** | Pointer to **NullableString** | The phone number of the place. | [optional] 

## Methods

### NewMicrosoftGraphPlace

`func NewMicrosoftGraphPlace() *MicrosoftGraphPlace`

NewMicrosoftGraphPlace instantiates a new MicrosoftGraphPlace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPlaceWithDefaults

`func NewMicrosoftGraphPlaceWithDefaults() *MicrosoftGraphPlace`

NewMicrosoftGraphPlaceWithDefaults instantiates a new MicrosoftGraphPlace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphPlace) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPlace) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphPlace) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphPlace) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAddress

`func (o *MicrosoftGraphPlace) GetAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MicrosoftGraphPlace) GetAddressOk() (*AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MicrosoftGraphPlace) SetAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MicrosoftGraphPlace) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *MicrosoftGraphPlace) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *MicrosoftGraphPlace) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphPlace) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphPlace) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphPlace) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphPlace) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetGeoCoordinates

`func (o *MicrosoftGraphPlace) GetGeoCoordinates() AnyOfmicrosoftGraphOutlookGeoCoordinates`

GetGeoCoordinates returns the GeoCoordinates field if non-nil, zero value otherwise.

### GetGeoCoordinatesOk

`func (o *MicrosoftGraphPlace) GetGeoCoordinatesOk() (*AnyOfmicrosoftGraphOutlookGeoCoordinates, bool)`

GetGeoCoordinatesOk returns a tuple with the GeoCoordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoCoordinates

`func (o *MicrosoftGraphPlace) SetGeoCoordinates(v AnyOfmicrosoftGraphOutlookGeoCoordinates)`

SetGeoCoordinates sets GeoCoordinates field to given value.

### HasGeoCoordinates

`func (o *MicrosoftGraphPlace) HasGeoCoordinates() bool`

HasGeoCoordinates returns a boolean if a field has been set.

### SetGeoCoordinatesNil

`func (o *MicrosoftGraphPlace) SetGeoCoordinatesNil(b bool)`

 SetGeoCoordinatesNil sets the value for GeoCoordinates to be an explicit nil

### UnsetGeoCoordinates
`func (o *MicrosoftGraphPlace) UnsetGeoCoordinates()`

UnsetGeoCoordinates ensures that no value is present for GeoCoordinates, not even an explicit nil
### GetPhone

`func (o *MicrosoftGraphPlace) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *MicrosoftGraphPlace) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *MicrosoftGraphPlace) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *MicrosoftGraphPlace) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *MicrosoftGraphPlace) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *MicrosoftGraphPlace) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


