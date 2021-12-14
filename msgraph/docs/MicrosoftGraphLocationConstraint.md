# MicrosoftGraphLocationConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsRequired** | Pointer to **NullableBool** | The client requests the service to include in the response a meeting location for the meeting. If this is true and all the resources are busy, findMeetingTimes will not return any meeting time suggestions. If this is false and all the resources are busy, findMeetingTimes would still look for meeting times without locations. | [optional] 
**Locations** | Pointer to [**[]AnyOfmicrosoftGraphLocationConstraintItem**](AnyOfmicrosoftGraphLocationConstraintItem.md) | Constraint information for one or more locations that the client requests for the meeting. | [optional] 
**SuggestLocation** | Pointer to **NullableBool** | The client requests the service to suggest one or more meeting locations. | [optional] 

## Methods

### NewMicrosoftGraphLocationConstraint

`func NewMicrosoftGraphLocationConstraint() *MicrosoftGraphLocationConstraint`

NewMicrosoftGraphLocationConstraint instantiates a new MicrosoftGraphLocationConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphLocationConstraintWithDefaults

`func NewMicrosoftGraphLocationConstraintWithDefaults() *MicrosoftGraphLocationConstraint`

NewMicrosoftGraphLocationConstraintWithDefaults instantiates a new MicrosoftGraphLocationConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsRequired

`func (o *MicrosoftGraphLocationConstraint) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *MicrosoftGraphLocationConstraint) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *MicrosoftGraphLocationConstraint) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *MicrosoftGraphLocationConstraint) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### SetIsRequiredNil

`func (o *MicrosoftGraphLocationConstraint) SetIsRequiredNil(b bool)`

 SetIsRequiredNil sets the value for IsRequired to be an explicit nil

### UnsetIsRequired
`func (o *MicrosoftGraphLocationConstraint) UnsetIsRequired()`

UnsetIsRequired ensures that no value is present for IsRequired, not even an explicit nil
### GetLocations

`func (o *MicrosoftGraphLocationConstraint) GetLocations() []*AnyOfmicrosoftGraphLocationConstraintItem`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *MicrosoftGraphLocationConstraint) GetLocationsOk() (*[]*AnyOfmicrosoftGraphLocationConstraintItem, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *MicrosoftGraphLocationConstraint) SetLocations(v []*AnyOfmicrosoftGraphLocationConstraintItem)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *MicrosoftGraphLocationConstraint) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetSuggestLocation

`func (o *MicrosoftGraphLocationConstraint) GetSuggestLocation() bool`

GetSuggestLocation returns the SuggestLocation field if non-nil, zero value otherwise.

### GetSuggestLocationOk

`func (o *MicrosoftGraphLocationConstraint) GetSuggestLocationOk() (*bool, bool)`

GetSuggestLocationOk returns a tuple with the SuggestLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestLocation

`func (o *MicrosoftGraphLocationConstraint) SetSuggestLocation(v bool)`

SetSuggestLocation sets SuggestLocation field to given value.

### HasSuggestLocation

`func (o *MicrosoftGraphLocationConstraint) HasSuggestLocation() bool`

HasSuggestLocation returns a boolean if a field has been set.

### SetSuggestLocationNil

`func (o *MicrosoftGraphLocationConstraint) SetSuggestLocationNil(b bool)`

 SetSuggestLocationNil sets the value for SuggestLocation to be an explicit nil

### UnsetSuggestLocation
`func (o *MicrosoftGraphLocationConstraint) UnsetSuggestLocation()`

UnsetSuggestLocation ensures that no value is present for SuggestLocation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


