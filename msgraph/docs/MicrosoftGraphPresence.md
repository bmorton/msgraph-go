# MicrosoftGraphPresence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Activity** | Pointer to **NullableString** | The supplemental information to a user&#39;s availability. Possible values are Available, Away, BeRightBack, Busy, DoNotDisturb, InACall, InAConferenceCall, Inactive, InAMeeting, Offline, OffWork, OutOfOffice, PresenceUnknown, Presenting, UrgentInterruptionsOnly. | [optional] 
**Availability** | Pointer to **NullableString** | The base presence information for a user. Possible values are Available, AvailableIdle,  Away, BeRightBack, Busy, BusyIdle, DoNotDisturb, Offline, PresenceUnknown | [optional] 

## Methods

### NewMicrosoftGraphPresence

`func NewMicrosoftGraphPresence() *MicrosoftGraphPresence`

NewMicrosoftGraphPresence instantiates a new MicrosoftGraphPresence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPresenceWithDefaults

`func NewMicrosoftGraphPresenceWithDefaults() *MicrosoftGraphPresence`

NewMicrosoftGraphPresenceWithDefaults instantiates a new MicrosoftGraphPresence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphPresence) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPresence) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphPresence) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphPresence) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActivity

`func (o *MicrosoftGraphPresence) GetActivity() string`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *MicrosoftGraphPresence) GetActivityOk() (*string, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *MicrosoftGraphPresence) SetActivity(v string)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *MicrosoftGraphPresence) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### SetActivityNil

`func (o *MicrosoftGraphPresence) SetActivityNil(b bool)`

 SetActivityNil sets the value for Activity to be an explicit nil

### UnsetActivity
`func (o *MicrosoftGraphPresence) UnsetActivity()`

UnsetActivity ensures that no value is present for Activity, not even an explicit nil
### GetAvailability

`func (o *MicrosoftGraphPresence) GetAvailability() string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *MicrosoftGraphPresence) GetAvailabilityOk() (*string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *MicrosoftGraphPresence) SetAvailability(v string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *MicrosoftGraphPresence) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### SetAvailabilityNil

`func (o *MicrosoftGraphPresence) SetAvailabilityNil(b bool)`

 SetAvailabilityNil sets the value for Availability to be an explicit nil

### UnsetAvailability
`func (o *MicrosoftGraphPresence) UnsetAvailability()`

UnsetAvailability ensures that no value is present for Availability, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


