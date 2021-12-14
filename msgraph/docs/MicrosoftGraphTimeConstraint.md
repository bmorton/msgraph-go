# MicrosoftGraphTimeConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityDomain** | Pointer to [**AnyOfmicrosoftGraphActivityDomain**](anyOf&lt;microsoft.graph.activityDomain&gt;.md) | The nature of the activity, optional. The possible values are: work, personal, unrestricted, or unknown. | [optional] 
**TimeSlots** | Pointer to [**[]AnyOfmicrosoftGraphTimeSlot**](AnyOfmicrosoftGraphTimeSlot.md) |  | [optional] 

## Methods

### NewMicrosoftGraphTimeConstraint

`func NewMicrosoftGraphTimeConstraint() *MicrosoftGraphTimeConstraint`

NewMicrosoftGraphTimeConstraint instantiates a new MicrosoftGraphTimeConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphTimeConstraintWithDefaults

`func NewMicrosoftGraphTimeConstraintWithDefaults() *MicrosoftGraphTimeConstraint`

NewMicrosoftGraphTimeConstraintWithDefaults instantiates a new MicrosoftGraphTimeConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityDomain

`func (o *MicrosoftGraphTimeConstraint) GetActivityDomain() AnyOfmicrosoftGraphActivityDomain`

GetActivityDomain returns the ActivityDomain field if non-nil, zero value otherwise.

### GetActivityDomainOk

`func (o *MicrosoftGraphTimeConstraint) GetActivityDomainOk() (*AnyOfmicrosoftGraphActivityDomain, bool)`

GetActivityDomainOk returns a tuple with the ActivityDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityDomain

`func (o *MicrosoftGraphTimeConstraint) SetActivityDomain(v AnyOfmicrosoftGraphActivityDomain)`

SetActivityDomain sets ActivityDomain field to given value.

### HasActivityDomain

`func (o *MicrosoftGraphTimeConstraint) HasActivityDomain() bool`

HasActivityDomain returns a boolean if a field has been set.

### SetActivityDomainNil

`func (o *MicrosoftGraphTimeConstraint) SetActivityDomainNil(b bool)`

 SetActivityDomainNil sets the value for ActivityDomain to be an explicit nil

### UnsetActivityDomain
`func (o *MicrosoftGraphTimeConstraint) UnsetActivityDomain()`

UnsetActivityDomain ensures that no value is present for ActivityDomain, not even an explicit nil
### GetTimeSlots

`func (o *MicrosoftGraphTimeConstraint) GetTimeSlots() []*AnyOfmicrosoftGraphTimeSlot`

GetTimeSlots returns the TimeSlots field if non-nil, zero value otherwise.

### GetTimeSlotsOk

`func (o *MicrosoftGraphTimeConstraint) GetTimeSlotsOk() (*[]*AnyOfmicrosoftGraphTimeSlot, bool)`

GetTimeSlotsOk returns a tuple with the TimeSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSlots

`func (o *MicrosoftGraphTimeConstraint) SetTimeSlots(v []*AnyOfmicrosoftGraphTimeSlot)`

SetTimeSlots sets TimeSlots field to given value.

### HasTimeSlots

`func (o *MicrosoftGraphTimeConstraint) HasTimeSlots() bool`

HasTimeSlots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


