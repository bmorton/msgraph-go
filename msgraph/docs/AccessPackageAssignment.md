# AccessPackageAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiredDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. | [optional] 
**Schedule** | Pointer to [**AnyOfmicrosoftGraphEntitlementManagementSchedule**](anyOf&lt;microsoft.graph.entitlementManagementSchedule&gt;.md) | When the access assignment is to be in place. Read-only. | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphAccessPackageAssignmentState**](anyOf&lt;microsoft.graph.accessPackageAssignmentState&gt;.md) | The state of the access package assignment. The possible values are: delivering, partiallyDelivered, delivered, expired, deliveryFailed, unknownFutureValue. Read-only. | [optional] 
**Status** | Pointer to **NullableString** | More information about the assignment lifecycle.  Possible values include Delivering, Delivered, NearExpiry1DayNotificationTriggered, or ExpiredNotificationTriggered.  Read-only. | [optional] 
**AccessPackage** | Pointer to [**AnyOfmicrosoftGraphAccessPackage**](anyOf&lt;microsoft.graph.accessPackage&gt;.md) | Read-only. Nullable. | [optional] 
**Target** | Pointer to [**AnyOfmicrosoftGraphAccessPackageSubject**](anyOf&lt;microsoft.graph.accessPackageSubject&gt;.md) | The subject of the access package assignment. Read-only. Nullable. | [optional] 

## Methods

### NewAccessPackageAssignment

`func NewAccessPackageAssignment() *AccessPackageAssignment`

NewAccessPackageAssignment instantiates a new AccessPackageAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPackageAssignmentWithDefaults

`func NewAccessPackageAssignmentWithDefaults() *AccessPackageAssignment`

NewAccessPackageAssignmentWithDefaults instantiates a new AccessPackageAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiredDateTime

`func (o *AccessPackageAssignment) GetExpiredDateTime() time.Time`

GetExpiredDateTime returns the ExpiredDateTime field if non-nil, zero value otherwise.

### GetExpiredDateTimeOk

`func (o *AccessPackageAssignment) GetExpiredDateTimeOk() (*time.Time, bool)`

GetExpiredDateTimeOk returns a tuple with the ExpiredDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredDateTime

`func (o *AccessPackageAssignment) SetExpiredDateTime(v time.Time)`

SetExpiredDateTime sets ExpiredDateTime field to given value.

### HasExpiredDateTime

`func (o *AccessPackageAssignment) HasExpiredDateTime() bool`

HasExpiredDateTime returns a boolean if a field has been set.

### SetExpiredDateTimeNil

`func (o *AccessPackageAssignment) SetExpiredDateTimeNil(b bool)`

 SetExpiredDateTimeNil sets the value for ExpiredDateTime to be an explicit nil

### UnsetExpiredDateTime
`func (o *AccessPackageAssignment) UnsetExpiredDateTime()`

UnsetExpiredDateTime ensures that no value is present for ExpiredDateTime, not even an explicit nil
### GetSchedule

`func (o *AccessPackageAssignment) GetSchedule() AnyOfmicrosoftGraphEntitlementManagementSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AccessPackageAssignment) GetScheduleOk() (*AnyOfmicrosoftGraphEntitlementManagementSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AccessPackageAssignment) SetSchedule(v AnyOfmicrosoftGraphEntitlementManagementSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *AccessPackageAssignment) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *AccessPackageAssignment) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *AccessPackageAssignment) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetState

`func (o *AccessPackageAssignment) GetState() AnyOfmicrosoftGraphAccessPackageAssignmentState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AccessPackageAssignment) GetStateOk() (*AnyOfmicrosoftGraphAccessPackageAssignmentState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AccessPackageAssignment) SetState(v AnyOfmicrosoftGraphAccessPackageAssignmentState)`

SetState sets State field to given value.

### HasState

`func (o *AccessPackageAssignment) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *AccessPackageAssignment) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *AccessPackageAssignment) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetStatus

`func (o *AccessPackageAssignment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccessPackageAssignment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccessPackageAssignment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccessPackageAssignment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AccessPackageAssignment) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AccessPackageAssignment) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetAccessPackage

`func (o *AccessPackageAssignment) GetAccessPackage() AnyOfmicrosoftGraphAccessPackage`

GetAccessPackage returns the AccessPackage field if non-nil, zero value otherwise.

### GetAccessPackageOk

`func (o *AccessPackageAssignment) GetAccessPackageOk() (*AnyOfmicrosoftGraphAccessPackage, bool)`

GetAccessPackageOk returns a tuple with the AccessPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPackage

`func (o *AccessPackageAssignment) SetAccessPackage(v AnyOfmicrosoftGraphAccessPackage)`

SetAccessPackage sets AccessPackage field to given value.

### HasAccessPackage

`func (o *AccessPackageAssignment) HasAccessPackage() bool`

HasAccessPackage returns a boolean if a field has been set.

### SetAccessPackageNil

`func (o *AccessPackageAssignment) SetAccessPackageNil(b bool)`

 SetAccessPackageNil sets the value for AccessPackage to be an explicit nil

### UnsetAccessPackage
`func (o *AccessPackageAssignment) UnsetAccessPackage()`

UnsetAccessPackage ensures that no value is present for AccessPackage, not even an explicit nil
### GetTarget

`func (o *AccessPackageAssignment) GetTarget() AnyOfmicrosoftGraphAccessPackageSubject`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AccessPackageAssignment) GetTargetOk() (*AnyOfmicrosoftGraphAccessPackageSubject, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AccessPackageAssignment) SetTarget(v AnyOfmicrosoftGraphAccessPackageSubject)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *AccessPackageAssignment) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *AccessPackageAssignment) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *AccessPackageAssignment) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


