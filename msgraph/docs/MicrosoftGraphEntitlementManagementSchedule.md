# MicrosoftGraphEntitlementManagementSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiration** | Pointer to [**AnyOfmicrosoftGraphExpirationPattern**](anyOf&lt;microsoft.graph.expirationPattern&gt;.md) | When the access should expire. | [optional] 
**Recurrence** | Pointer to [**AnyOfmicrosoftGraphPatternedRecurrence**](anyOf&lt;microsoft.graph.patternedRecurrence&gt;.md) | For recurring access. Not used at present. | [optional] 
**StartDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. | [optional] 

## Methods

### NewMicrosoftGraphEntitlementManagementSchedule

`func NewMicrosoftGraphEntitlementManagementSchedule() *MicrosoftGraphEntitlementManagementSchedule`

NewMicrosoftGraphEntitlementManagementSchedule instantiates a new MicrosoftGraphEntitlementManagementSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphEntitlementManagementScheduleWithDefaults

`func NewMicrosoftGraphEntitlementManagementScheduleWithDefaults() *MicrosoftGraphEntitlementManagementSchedule`

NewMicrosoftGraphEntitlementManagementScheduleWithDefaults instantiates a new MicrosoftGraphEntitlementManagementSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiration

`func (o *MicrosoftGraphEntitlementManagementSchedule) GetExpiration() AnyOfmicrosoftGraphExpirationPattern`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *MicrosoftGraphEntitlementManagementSchedule) GetExpirationOk() (*AnyOfmicrosoftGraphExpirationPattern, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *MicrosoftGraphEntitlementManagementSchedule) SetExpiration(v AnyOfmicrosoftGraphExpirationPattern)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *MicrosoftGraphEntitlementManagementSchedule) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### SetExpirationNil

`func (o *MicrosoftGraphEntitlementManagementSchedule) SetExpirationNil(b bool)`

 SetExpirationNil sets the value for Expiration to be an explicit nil

### UnsetExpiration
`func (o *MicrosoftGraphEntitlementManagementSchedule) UnsetExpiration()`

UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
### GetRecurrence

`func (o *MicrosoftGraphEntitlementManagementSchedule) GetRecurrence() AnyOfmicrosoftGraphPatternedRecurrence`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *MicrosoftGraphEntitlementManagementSchedule) GetRecurrenceOk() (*AnyOfmicrosoftGraphPatternedRecurrence, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrence

`func (o *MicrosoftGraphEntitlementManagementSchedule) SetRecurrence(v AnyOfmicrosoftGraphPatternedRecurrence)`

SetRecurrence sets Recurrence field to given value.

### HasRecurrence

`func (o *MicrosoftGraphEntitlementManagementSchedule) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.

### SetRecurrenceNil

`func (o *MicrosoftGraphEntitlementManagementSchedule) SetRecurrenceNil(b bool)`

 SetRecurrenceNil sets the value for Recurrence to be an explicit nil

### UnsetRecurrence
`func (o *MicrosoftGraphEntitlementManagementSchedule) UnsetRecurrence()`

UnsetRecurrence ensures that no value is present for Recurrence, not even an explicit nil
### GetStartDateTime

`func (o *MicrosoftGraphEntitlementManagementSchedule) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MicrosoftGraphEntitlementManagementSchedule) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MicrosoftGraphEntitlementManagementSchedule) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MicrosoftGraphEntitlementManagementSchedule) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTimeNil

`func (o *MicrosoftGraphEntitlementManagementSchedule) SetStartDateTimeNil(b bool)`

 SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil

### UnsetStartDateTime
`func (o *MicrosoftGraphEntitlementManagementSchedule) UnsetStartDateTime()`

UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


