# MicrosoftGraphAccessReviewInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**EndDateTime** | Pointer to **NullableTime** | DateTime when review instance is scheduled to end.The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $select. Read-only. | [optional] 
**FallbackReviewers** | Pointer to [**[]AnyOfmicrosoftGraphAccessReviewReviewerScope**](AnyOfmicrosoftGraphAccessReviewReviewerScope.md) | This collection of reviewer scopes is used to define the list of fallback reviewers. These fallback reviewers will be notified to take action if no users are found from the list of reviewers specified. This could occur when either the group owner is specified as the reviewer but the group owner does not exist, or manager is specified as reviewer but a user&#39;s manager does not exist. Supports $select. | [optional] 
**Reviewers** | Pointer to [**[]AnyOfmicrosoftGraphAccessReviewReviewerScope**](AnyOfmicrosoftGraphAccessReviewReviewerScope.md) | This collection of access review scopes is used to define who the reviewers are. Supports $select. For examples of options for assigning reviewers, see Assign reviewers to your access review definition using the Microsoft Graph API. | [optional] 
**Scope** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Created based on scope and instanceEnumerationScope at the accessReviewScheduleDefinition level. Defines the scope of users reviewed in a group. Supports $select and $filter (contains only). Read-only. | [optional] 
**StartDateTime** | Pointer to **NullableTime** | DateTime when review instance is scheduled to start. May be in the future. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $select. Read-only. | [optional] 
**Status** | Pointer to **NullableString** | Specifies the status of an accessReview. Possible values: Initializing, NotStarted, Starting, InProgress, Completing, Completed, AutoReviewing, and AutoReviewed. Supports $select, $orderby, and $filter (eq only). Read-only. | [optional] 
**Decisions** | Pointer to [**[]MicrosoftGraphAccessReviewInstanceDecisionItem**](MicrosoftGraphAccessReviewInstanceDecisionItem.md) | Each principal reviewed in an accessReviewInstance has a decision item representing if they were approved, denied, or not yet reviewed. | [optional] 

## Methods

### NewMicrosoftGraphAccessReviewInstance

`func NewMicrosoftGraphAccessReviewInstance() *MicrosoftGraphAccessReviewInstance`

NewMicrosoftGraphAccessReviewInstance instantiates a new MicrosoftGraphAccessReviewInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAccessReviewInstanceWithDefaults

`func NewMicrosoftGraphAccessReviewInstanceWithDefaults() *MicrosoftGraphAccessReviewInstance`

NewMicrosoftGraphAccessReviewInstanceWithDefaults instantiates a new MicrosoftGraphAccessReviewInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphAccessReviewInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAccessReviewInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphAccessReviewInstance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphAccessReviewInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEndDateTime

`func (o *MicrosoftGraphAccessReviewInstance) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *MicrosoftGraphAccessReviewInstance) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *MicrosoftGraphAccessReviewInstance) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *MicrosoftGraphAccessReviewInstance) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### SetEndDateTimeNil

`func (o *MicrosoftGraphAccessReviewInstance) SetEndDateTimeNil(b bool)`

 SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil

### UnsetEndDateTime
`func (o *MicrosoftGraphAccessReviewInstance) UnsetEndDateTime()`

UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
### GetFallbackReviewers

`func (o *MicrosoftGraphAccessReviewInstance) GetFallbackReviewers() []*AnyOfmicrosoftGraphAccessReviewReviewerScope`

GetFallbackReviewers returns the FallbackReviewers field if non-nil, zero value otherwise.

### GetFallbackReviewersOk

`func (o *MicrosoftGraphAccessReviewInstance) GetFallbackReviewersOk() (*[]*AnyOfmicrosoftGraphAccessReviewReviewerScope, bool)`

GetFallbackReviewersOk returns a tuple with the FallbackReviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackReviewers

`func (o *MicrosoftGraphAccessReviewInstance) SetFallbackReviewers(v []*AnyOfmicrosoftGraphAccessReviewReviewerScope)`

SetFallbackReviewers sets FallbackReviewers field to given value.

### HasFallbackReviewers

`func (o *MicrosoftGraphAccessReviewInstance) HasFallbackReviewers() bool`

HasFallbackReviewers returns a boolean if a field has been set.

### GetReviewers

`func (o *MicrosoftGraphAccessReviewInstance) GetReviewers() []*AnyOfmicrosoftGraphAccessReviewReviewerScope`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *MicrosoftGraphAccessReviewInstance) GetReviewersOk() (*[]*AnyOfmicrosoftGraphAccessReviewReviewerScope, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *MicrosoftGraphAccessReviewInstance) SetReviewers(v []*AnyOfmicrosoftGraphAccessReviewReviewerScope)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *MicrosoftGraphAccessReviewInstance) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.

### GetScope

`func (o *MicrosoftGraphAccessReviewInstance) GetScope() AnyOfobject`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *MicrosoftGraphAccessReviewInstance) GetScopeOk() (*AnyOfobject, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *MicrosoftGraphAccessReviewInstance) SetScope(v AnyOfobject)`

SetScope sets Scope field to given value.

### HasScope

`func (o *MicrosoftGraphAccessReviewInstance) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *MicrosoftGraphAccessReviewInstance) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *MicrosoftGraphAccessReviewInstance) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetStartDateTime

`func (o *MicrosoftGraphAccessReviewInstance) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MicrosoftGraphAccessReviewInstance) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MicrosoftGraphAccessReviewInstance) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MicrosoftGraphAccessReviewInstance) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTimeNil

`func (o *MicrosoftGraphAccessReviewInstance) SetStartDateTimeNil(b bool)`

 SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil

### UnsetStartDateTime
`func (o *MicrosoftGraphAccessReviewInstance) UnsetStartDateTime()`

UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphAccessReviewInstance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphAccessReviewInstance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphAccessReviewInstance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphAccessReviewInstance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphAccessReviewInstance) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphAccessReviewInstance) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDecisions

`func (o *MicrosoftGraphAccessReviewInstance) GetDecisions() []MicrosoftGraphAccessReviewInstanceDecisionItem`

GetDecisions returns the Decisions field if non-nil, zero value otherwise.

### GetDecisionsOk

`func (o *MicrosoftGraphAccessReviewInstance) GetDecisionsOk() (*[]MicrosoftGraphAccessReviewInstanceDecisionItem, bool)`

GetDecisionsOk returns a tuple with the Decisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisions

`func (o *MicrosoftGraphAccessReviewInstance) SetDecisions(v []MicrosoftGraphAccessReviewInstanceDecisionItem)`

SetDecisions sets Decisions field to given value.

### HasDecisions

`func (o *MicrosoftGraphAccessReviewInstance) HasDecisions() bool`

HasDecisions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


