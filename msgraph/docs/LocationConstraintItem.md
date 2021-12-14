# LocationConstraintItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResolveAvailability** | Pointer to **NullableBool** | If set to true and the specified resource is busy, findMeetingTimes looks for another resource that is free. If set to false and the specified resource is busy, findMeetingTimes returns the resource best ranked in the user&#39;s cache without checking if it&#39;s free. Default is true. | [optional] 

## Methods

### NewLocationConstraintItem

`func NewLocationConstraintItem() *LocationConstraintItem`

NewLocationConstraintItem instantiates a new LocationConstraintItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationConstraintItemWithDefaults

`func NewLocationConstraintItemWithDefaults() *LocationConstraintItem`

NewLocationConstraintItemWithDefaults instantiates a new LocationConstraintItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResolveAvailability

`func (o *LocationConstraintItem) GetResolveAvailability() bool`

GetResolveAvailability returns the ResolveAvailability field if non-nil, zero value otherwise.

### GetResolveAvailabilityOk

`func (o *LocationConstraintItem) GetResolveAvailabilityOk() (*bool, bool)`

GetResolveAvailabilityOk returns a tuple with the ResolveAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveAvailability

`func (o *LocationConstraintItem) SetResolveAvailability(v bool)`

SetResolveAvailability sets ResolveAvailability field to given value.

### HasResolveAvailability

`func (o *LocationConstraintItem) HasResolveAvailability() bool`

HasResolveAvailability returns a boolean if a field has been set.

### SetResolveAvailabilityNil

`func (o *LocationConstraintItem) SetResolveAvailabilityNil(b bool)`

 SetResolveAvailabilityNil sets the value for ResolveAvailability to be an explicit nil

### UnsetResolveAvailability
`func (o *LocationConstraintItem) UnsetResolveAvailability()`

UnsetResolveAvailability ensures that no value is present for ResolveAvailability, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


