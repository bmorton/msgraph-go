# MicrosoftGraphLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**AnyOfmicrosoftGraphPhysicalAddress**](anyOf&lt;microsoft.graph.physicalAddress&gt;.md) | The street address of the location. | [optional] 
**Coordinates** | Pointer to [**AnyOfmicrosoftGraphOutlookGeoCoordinates**](anyOf&lt;microsoft.graph.outlookGeoCoordinates&gt;.md) | The geographic coordinates and elevation of the location. | [optional] 
**DisplayName** | Pointer to **NullableString** | The name associated with the location. | [optional] 
**LocationEmailAddress** | Pointer to **NullableString** | Optional email address of the location. | [optional] 
**LocationType** | Pointer to [**AnyOfmicrosoftGraphLocationType**](anyOf&lt;microsoft.graph.locationType&gt;.md) | The type of location. The possible values are: default, conferenceRoom, homeAddress, businessAddress,geoCoordinates, streetAddress, hotel, restaurant, localBusiness, postalAddress. Read-only. | [optional] 
**LocationUri** | Pointer to **NullableString** | Optional URI representing the location. | [optional] 
**UniqueId** | Pointer to **NullableString** | For internal use only. | [optional] 
**UniqueIdType** | Pointer to [**AnyOfmicrosoftGraphLocationUniqueIdType**](anyOf&lt;microsoft.graph.locationUniqueIdType&gt;.md) | For internal use only. | [optional] 

## Methods

### NewMicrosoftGraphLocation

`func NewMicrosoftGraphLocation() *MicrosoftGraphLocation`

NewMicrosoftGraphLocation instantiates a new MicrosoftGraphLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphLocationWithDefaults

`func NewMicrosoftGraphLocationWithDefaults() *MicrosoftGraphLocation`

NewMicrosoftGraphLocationWithDefaults instantiates a new MicrosoftGraphLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *MicrosoftGraphLocation) GetAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MicrosoftGraphLocation) GetAddressOk() (*AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MicrosoftGraphLocation) SetAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MicrosoftGraphLocation) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *MicrosoftGraphLocation) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *MicrosoftGraphLocation) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetCoordinates

`func (o *MicrosoftGraphLocation) GetCoordinates() AnyOfmicrosoftGraphOutlookGeoCoordinates`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *MicrosoftGraphLocation) GetCoordinatesOk() (*AnyOfmicrosoftGraphOutlookGeoCoordinates, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *MicrosoftGraphLocation) SetCoordinates(v AnyOfmicrosoftGraphOutlookGeoCoordinates)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *MicrosoftGraphLocation) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.

### SetCoordinatesNil

`func (o *MicrosoftGraphLocation) SetCoordinatesNil(b bool)`

 SetCoordinatesNil sets the value for Coordinates to be an explicit nil

### UnsetCoordinates
`func (o *MicrosoftGraphLocation) UnsetCoordinates()`

UnsetCoordinates ensures that no value is present for Coordinates, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphLocation) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphLocation) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphLocation) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphLocation) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphLocation) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphLocation) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetLocationEmailAddress

`func (o *MicrosoftGraphLocation) GetLocationEmailAddress() string`

GetLocationEmailAddress returns the LocationEmailAddress field if non-nil, zero value otherwise.

### GetLocationEmailAddressOk

`func (o *MicrosoftGraphLocation) GetLocationEmailAddressOk() (*string, bool)`

GetLocationEmailAddressOk returns a tuple with the LocationEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationEmailAddress

`func (o *MicrosoftGraphLocation) SetLocationEmailAddress(v string)`

SetLocationEmailAddress sets LocationEmailAddress field to given value.

### HasLocationEmailAddress

`func (o *MicrosoftGraphLocation) HasLocationEmailAddress() bool`

HasLocationEmailAddress returns a boolean if a field has been set.

### SetLocationEmailAddressNil

`func (o *MicrosoftGraphLocation) SetLocationEmailAddressNil(b bool)`

 SetLocationEmailAddressNil sets the value for LocationEmailAddress to be an explicit nil

### UnsetLocationEmailAddress
`func (o *MicrosoftGraphLocation) UnsetLocationEmailAddress()`

UnsetLocationEmailAddress ensures that no value is present for LocationEmailAddress, not even an explicit nil
### GetLocationType

`func (o *MicrosoftGraphLocation) GetLocationType() AnyOfmicrosoftGraphLocationType`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *MicrosoftGraphLocation) GetLocationTypeOk() (*AnyOfmicrosoftGraphLocationType, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *MicrosoftGraphLocation) SetLocationType(v AnyOfmicrosoftGraphLocationType)`

SetLocationType sets LocationType field to given value.

### HasLocationType

`func (o *MicrosoftGraphLocation) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### SetLocationTypeNil

`func (o *MicrosoftGraphLocation) SetLocationTypeNil(b bool)`

 SetLocationTypeNil sets the value for LocationType to be an explicit nil

### UnsetLocationType
`func (o *MicrosoftGraphLocation) UnsetLocationType()`

UnsetLocationType ensures that no value is present for LocationType, not even an explicit nil
### GetLocationUri

`func (o *MicrosoftGraphLocation) GetLocationUri() string`

GetLocationUri returns the LocationUri field if non-nil, zero value otherwise.

### GetLocationUriOk

`func (o *MicrosoftGraphLocation) GetLocationUriOk() (*string, bool)`

GetLocationUriOk returns a tuple with the LocationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationUri

`func (o *MicrosoftGraphLocation) SetLocationUri(v string)`

SetLocationUri sets LocationUri field to given value.

### HasLocationUri

`func (o *MicrosoftGraphLocation) HasLocationUri() bool`

HasLocationUri returns a boolean if a field has been set.

### SetLocationUriNil

`func (o *MicrosoftGraphLocation) SetLocationUriNil(b bool)`

 SetLocationUriNil sets the value for LocationUri to be an explicit nil

### UnsetLocationUri
`func (o *MicrosoftGraphLocation) UnsetLocationUri()`

UnsetLocationUri ensures that no value is present for LocationUri, not even an explicit nil
### GetUniqueId

`func (o *MicrosoftGraphLocation) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *MicrosoftGraphLocation) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *MicrosoftGraphLocation) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *MicrosoftGraphLocation) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### SetUniqueIdNil

`func (o *MicrosoftGraphLocation) SetUniqueIdNil(b bool)`

 SetUniqueIdNil sets the value for UniqueId to be an explicit nil

### UnsetUniqueId
`func (o *MicrosoftGraphLocation) UnsetUniqueId()`

UnsetUniqueId ensures that no value is present for UniqueId, not even an explicit nil
### GetUniqueIdType

`func (o *MicrosoftGraphLocation) GetUniqueIdType() AnyOfmicrosoftGraphLocationUniqueIdType`

GetUniqueIdType returns the UniqueIdType field if non-nil, zero value otherwise.

### GetUniqueIdTypeOk

`func (o *MicrosoftGraphLocation) GetUniqueIdTypeOk() (*AnyOfmicrosoftGraphLocationUniqueIdType, bool)`

GetUniqueIdTypeOk returns a tuple with the UniqueIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdType

`func (o *MicrosoftGraphLocation) SetUniqueIdType(v AnyOfmicrosoftGraphLocationUniqueIdType)`

SetUniqueIdType sets UniqueIdType field to given value.

### HasUniqueIdType

`func (o *MicrosoftGraphLocation) HasUniqueIdType() bool`

HasUniqueIdType returns a boolean if a field has been set.

### SetUniqueIdTypeNil

`func (o *MicrosoftGraphLocation) SetUniqueIdTypeNil(b bool)`

 SetUniqueIdTypeNil sets the value for UniqueIdType to be an explicit nil

### UnsetUniqueIdType
`func (o *MicrosoftGraphLocation) UnsetUniqueIdType()`

UnsetUniqueIdType ensures that no value is present for UniqueIdType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


