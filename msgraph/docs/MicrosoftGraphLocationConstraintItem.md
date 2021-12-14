# MicrosoftGraphLocationConstraintItem

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
**ResolveAvailability** | Pointer to **NullableBool** | If set to true and the specified resource is busy, findMeetingTimes looks for another resource that is free. If set to false and the specified resource is busy, findMeetingTimes returns the resource best ranked in the user&#39;s cache without checking if it&#39;s free. Default is true. | [optional] 

## Methods

### NewMicrosoftGraphLocationConstraintItem

`func NewMicrosoftGraphLocationConstraintItem() *MicrosoftGraphLocationConstraintItem`

NewMicrosoftGraphLocationConstraintItem instantiates a new MicrosoftGraphLocationConstraintItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphLocationConstraintItemWithDefaults

`func NewMicrosoftGraphLocationConstraintItemWithDefaults() *MicrosoftGraphLocationConstraintItem`

NewMicrosoftGraphLocationConstraintItemWithDefaults instantiates a new MicrosoftGraphLocationConstraintItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *MicrosoftGraphLocationConstraintItem) GetAddress() AnyOfmicrosoftGraphPhysicalAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MicrosoftGraphLocationConstraintItem) GetAddressOk() (*AnyOfmicrosoftGraphPhysicalAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MicrosoftGraphLocationConstraintItem) SetAddress(v AnyOfmicrosoftGraphPhysicalAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MicrosoftGraphLocationConstraintItem) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *MicrosoftGraphLocationConstraintItem) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *MicrosoftGraphLocationConstraintItem) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetCoordinates

`func (o *MicrosoftGraphLocationConstraintItem) GetCoordinates() AnyOfmicrosoftGraphOutlookGeoCoordinates`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *MicrosoftGraphLocationConstraintItem) GetCoordinatesOk() (*AnyOfmicrosoftGraphOutlookGeoCoordinates, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *MicrosoftGraphLocationConstraintItem) SetCoordinates(v AnyOfmicrosoftGraphOutlookGeoCoordinates)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *MicrosoftGraphLocationConstraintItem) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.

### SetCoordinatesNil

`func (o *MicrosoftGraphLocationConstraintItem) SetCoordinatesNil(b bool)`

 SetCoordinatesNil sets the value for Coordinates to be an explicit nil

### UnsetCoordinates
`func (o *MicrosoftGraphLocationConstraintItem) UnsetCoordinates()`

UnsetCoordinates ensures that no value is present for Coordinates, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphLocationConstraintItem) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphLocationConstraintItem) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphLocationConstraintItem) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphLocationConstraintItem) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphLocationConstraintItem) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphLocationConstraintItem) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetLocationEmailAddress

`func (o *MicrosoftGraphLocationConstraintItem) GetLocationEmailAddress() string`

GetLocationEmailAddress returns the LocationEmailAddress field if non-nil, zero value otherwise.

### GetLocationEmailAddressOk

`func (o *MicrosoftGraphLocationConstraintItem) GetLocationEmailAddressOk() (*string, bool)`

GetLocationEmailAddressOk returns a tuple with the LocationEmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationEmailAddress

`func (o *MicrosoftGraphLocationConstraintItem) SetLocationEmailAddress(v string)`

SetLocationEmailAddress sets LocationEmailAddress field to given value.

### HasLocationEmailAddress

`func (o *MicrosoftGraphLocationConstraintItem) HasLocationEmailAddress() bool`

HasLocationEmailAddress returns a boolean if a field has been set.

### SetLocationEmailAddressNil

`func (o *MicrosoftGraphLocationConstraintItem) SetLocationEmailAddressNil(b bool)`

 SetLocationEmailAddressNil sets the value for LocationEmailAddress to be an explicit nil

### UnsetLocationEmailAddress
`func (o *MicrosoftGraphLocationConstraintItem) UnsetLocationEmailAddress()`

UnsetLocationEmailAddress ensures that no value is present for LocationEmailAddress, not even an explicit nil
### GetLocationType

`func (o *MicrosoftGraphLocationConstraintItem) GetLocationType() AnyOfmicrosoftGraphLocationType`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *MicrosoftGraphLocationConstraintItem) GetLocationTypeOk() (*AnyOfmicrosoftGraphLocationType, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *MicrosoftGraphLocationConstraintItem) SetLocationType(v AnyOfmicrosoftGraphLocationType)`

SetLocationType sets LocationType field to given value.

### HasLocationType

`func (o *MicrosoftGraphLocationConstraintItem) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### SetLocationTypeNil

`func (o *MicrosoftGraphLocationConstraintItem) SetLocationTypeNil(b bool)`

 SetLocationTypeNil sets the value for LocationType to be an explicit nil

### UnsetLocationType
`func (o *MicrosoftGraphLocationConstraintItem) UnsetLocationType()`

UnsetLocationType ensures that no value is present for LocationType, not even an explicit nil
### GetLocationUri

`func (o *MicrosoftGraphLocationConstraintItem) GetLocationUri() string`

GetLocationUri returns the LocationUri field if non-nil, zero value otherwise.

### GetLocationUriOk

`func (o *MicrosoftGraphLocationConstraintItem) GetLocationUriOk() (*string, bool)`

GetLocationUriOk returns a tuple with the LocationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationUri

`func (o *MicrosoftGraphLocationConstraintItem) SetLocationUri(v string)`

SetLocationUri sets LocationUri field to given value.

### HasLocationUri

`func (o *MicrosoftGraphLocationConstraintItem) HasLocationUri() bool`

HasLocationUri returns a boolean if a field has been set.

### SetLocationUriNil

`func (o *MicrosoftGraphLocationConstraintItem) SetLocationUriNil(b bool)`

 SetLocationUriNil sets the value for LocationUri to be an explicit nil

### UnsetLocationUri
`func (o *MicrosoftGraphLocationConstraintItem) UnsetLocationUri()`

UnsetLocationUri ensures that no value is present for LocationUri, not even an explicit nil
### GetUniqueId

`func (o *MicrosoftGraphLocationConstraintItem) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *MicrosoftGraphLocationConstraintItem) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *MicrosoftGraphLocationConstraintItem) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *MicrosoftGraphLocationConstraintItem) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### SetUniqueIdNil

`func (o *MicrosoftGraphLocationConstraintItem) SetUniqueIdNil(b bool)`

 SetUniqueIdNil sets the value for UniqueId to be an explicit nil

### UnsetUniqueId
`func (o *MicrosoftGraphLocationConstraintItem) UnsetUniqueId()`

UnsetUniqueId ensures that no value is present for UniqueId, not even an explicit nil
### GetUniqueIdType

`func (o *MicrosoftGraphLocationConstraintItem) GetUniqueIdType() AnyOfmicrosoftGraphLocationUniqueIdType`

GetUniqueIdType returns the UniqueIdType field if non-nil, zero value otherwise.

### GetUniqueIdTypeOk

`func (o *MicrosoftGraphLocationConstraintItem) GetUniqueIdTypeOk() (*AnyOfmicrosoftGraphLocationUniqueIdType, bool)`

GetUniqueIdTypeOk returns a tuple with the UniqueIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdType

`func (o *MicrosoftGraphLocationConstraintItem) SetUniqueIdType(v AnyOfmicrosoftGraphLocationUniqueIdType)`

SetUniqueIdType sets UniqueIdType field to given value.

### HasUniqueIdType

`func (o *MicrosoftGraphLocationConstraintItem) HasUniqueIdType() bool`

HasUniqueIdType returns a boolean if a field has been set.

### SetUniqueIdTypeNil

`func (o *MicrosoftGraphLocationConstraintItem) SetUniqueIdTypeNil(b bool)`

 SetUniqueIdTypeNil sets the value for UniqueIdType to be an explicit nil

### UnsetUniqueIdType
`func (o *MicrosoftGraphLocationConstraintItem) UnsetUniqueIdType()`

UnsetUniqueIdType ensures that no value is present for UniqueIdType, not even an explicit nil
### GetResolveAvailability

`func (o *MicrosoftGraphLocationConstraintItem) GetResolveAvailability() bool`

GetResolveAvailability returns the ResolveAvailability field if non-nil, zero value otherwise.

### GetResolveAvailabilityOk

`func (o *MicrosoftGraphLocationConstraintItem) GetResolveAvailabilityOk() (*bool, bool)`

GetResolveAvailabilityOk returns a tuple with the ResolveAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveAvailability

`func (o *MicrosoftGraphLocationConstraintItem) SetResolveAvailability(v bool)`

SetResolveAvailability sets ResolveAvailability field to given value.

### HasResolveAvailability

`func (o *MicrosoftGraphLocationConstraintItem) HasResolveAvailability() bool`

HasResolveAvailability returns a boolean if a field has been set.

### SetResolveAvailabilityNil

`func (o *MicrosoftGraphLocationConstraintItem) SetResolveAvailabilityNil(b bool)`

 SetResolveAvailabilityNil sets the value for ResolveAvailability to be an explicit nil

### UnsetResolveAvailability
`func (o *MicrosoftGraphLocationConstraintItem) UnsetResolveAvailability()`

UnsetResolveAvailability ensures that no value is present for ResolveAvailability, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


