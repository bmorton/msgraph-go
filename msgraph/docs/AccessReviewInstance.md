# AccessReviewInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDateTime** | Pointer to **NullableTime** | DateTime when review instance is scheduled to end.The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $select. Read-only. | [optional] 
**FallbackReviewers** | Pointer to [**[]AnyOfmicrosoftGraphAccessReviewReviewerScope**](AnyOfmicrosoftGraphAccessReviewReviewerScope.md) | This collection of reviewer scopes is used to define the list of fallback reviewers. These fallback reviewers will be notified to take action if no users are found from the list of reviewers specified. This could occur when either the group owner is specified as the reviewer but the group owner does not exist, or manager is specified as reviewer but a user&#39;s manager does not exist. Supports $select. | [optional] 
**Reviewers** | Pointer to [**[]AnyOfmicrosoftGraphAccessReviewReviewerScope**](AnyOfmicrosoftGraphAccessReviewReviewerScope.md) | This collection of access review scopes is used to define who the reviewers are. Supports $select. For examples of options for assigning reviewers, see Assign reviewers to your access review definition using the Microsoft Graph API. | [optional] 
**Scope** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Created based on scope and instanceEnumerationScope at the accessReviewScheduleDefinition level. Defines the scope of users reviewed in a group. Supports $select and $filter (contains only). Read-only. | [optional] 
**StartDateTime** | Pointer to **NullableTime** | DateTime when review instance is scheduled to start. May be in the future. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $select. Read-only. | [optional] 
**Status** | Pointer to **NullableString** | Specifies the status of an accessReview. Possible values: Initializing, NotStarted, Starting, InProgress, Completing, Completed, AutoReviewing, and AutoReviewed. Supports $select, $orderby, and $filter (eq only). Read-only. | [optional] 
**Decisions** | Pointer to [**[]MicrosoftGraphAccessReviewInstanceDecisionItem**](MicrosoftGraphAccessReviewInstanceDecisionItem.md) | Each principal reviewed in an accessReviewInstance has a decision item representing if they were approved, denied, or not yet reviewed. | [optional] 

## Methods

### NewAccessReviewInstance

`func NewAccessReviewInstance() *AccessReviewInstance`

NewAccessReviewInstance instantiates a new AccessReviewInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessReviewInstanceWithDefaults

`func NewAccessReviewInstanceWithDefaults() *AccessReviewInstance`

NewAccessReviewInstanceWithDefaults instantiates a new AccessReviewInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDateTime

`func (o *AccessReviewInstance) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *AccessReviewInstance) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *AccessReviewInstance) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *AccessReviewInstance) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### SetEndDateTimeNil

`func (o *AccessReviewInstance) SetEndDateTimeNil(b bool)`

 SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil

### UnsetEndDateTime
`func (o *AccessReviewInstance) UnsetEndDateTime()`

UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
### GetFallbackReviewers

`func (o *AccessReviewInstance) GetFallbackReviewers() []*AnyOfmicrosoftGraphAccessReviewReviewerScope`

GetFallbackReviewers returns the FallbackReviewers field if non-nil, zero value otherwise.

### GetFallbackReviewersOk

`func (o *AccessReviewInstance) GetFallbackReviewersOk() (*[]*AnyOfmicrosoftGraphAccessReviewReviewerScope, bool)`

GetFallbackReviewersOk returns a tuple with the FallbackReviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackReviewers

`func (o *AccessReviewInstance) SetFallbackReviewers(v []*AnyOfmicrosoftGraphAccessReviewReviewerScope)`

SetFallbackReviewers sets FallbackReviewers field to given value.

### HasFallbackReviewers

`func (o *AccessReviewInstance) HasFallbackReviewers() bool`

HasFallbackReviewers returns a boolean if a field has been set.

### GetReviewers

`func (o *AccessReviewInstance) GetReviewers() []*AnyOfmicrosoftGraphAccessReviewReviewerScope`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *AccessReviewInstance) GetReviewersOk() (*[]*AnyOfmicrosoftGraphAccessReviewReviewerScope, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *AccessReviewInstance) SetReviewers(v []*AnyOfmicrosoftGraphAccessReviewReviewerScope)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *AccessReviewInstance) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.

### GetScope

`func (o *AccessReviewInstance) GetScope() AnyOfobject`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AccessReviewInstance) GetScopeOk() (*AnyOfobject, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AccessReviewInstance) SetScope(v AnyOfobject)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AccessReviewInstance) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *AccessReviewInstance) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *AccessReviewInstance) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetStartDateTime

`func (o *AccessReviewInstance) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *AccessReviewInstance) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *AccessReviewInstance) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *AccessReviewInstance) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTimeNil

`func (o *AccessReviewInstance) SetStartDateTimeNil(b bool)`

 SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil

### UnsetStartDateTime
`func (o *AccessReviewInstance) UnsetStartDateTime()`

UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
### GetStatus

`func (o *AccessReviewInstance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccessReviewInstance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccessReviewInstance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccessReviewInstance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AccessReviewInstance) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AccessReviewInstance) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDecisions

`func (o *AccessReviewInstance) GetDecisions() []MicrosoftGraphAccessReviewInstanceDecisionItem`

GetDecisions returns the Decisions field if non-nil, zero value otherwise.

### GetDecisionsOk

`func (o *AccessReviewInstance) GetDecisionsOk() (*[]MicrosoftGraphAccessReviewInstanceDecisionItem, bool)`

GetDecisionsOk returns a tuple with the Decisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisions

`func (o *AccessReviewInstance) SetDecisions(v []MicrosoftGraphAccessReviewInstanceDecisionItem)`

SetDecisions sets Decisions field to given value.

### HasDecisions

`func (o *AccessReviewInstance) HasDecisions() bool`

HasDecisions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


