# Presence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activity** | Pointer to **NullableString** | The supplemental information to a user&#39;s availability. Possible values are Available, Away, BeRightBack, Busy, DoNotDisturb, InACall, InAConferenceCall, Inactive, InAMeeting, Offline, OffWork, OutOfOffice, PresenceUnknown, Presenting, UrgentInterruptionsOnly. | [optional] 
**Availability** | Pointer to **NullableString** | The base presence information for a user. Possible values are Available, AvailableIdle,  Away, BeRightBack, Busy, BusyIdle, DoNotDisturb, Offline, PresenceUnknown | [optional] 

## Methods

### NewPresence

`func NewPresence() *Presence`

NewPresence instantiates a new Presence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresenceWithDefaults

`func NewPresenceWithDefaults() *Presence`

NewPresenceWithDefaults instantiates a new Presence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivity

`func (o *Presence) GetActivity() string`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *Presence) GetActivityOk() (*string, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *Presence) SetActivity(v string)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *Presence) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### SetActivityNil

`func (o *Presence) SetActivityNil(b bool)`

 SetActivityNil sets the value for Activity to be an explicit nil

### UnsetActivity
`func (o *Presence) UnsetActivity()`

UnsetActivity ensures that no value is present for Activity, not even an explicit nil
### GetAvailability

`func (o *Presence) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *Presence) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *Presence) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *Presence) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### SetAvailabilityNil

`func (o *Presence) SetAvailabilityNil(b bool)`

 SetAvailabilityNil sets the value for Availability to be an explicit nil

### UnsetAvailability
`func (o *Presence) UnsetAvailability()`

UnsetAvailability ensures that no value is present for Availability, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


