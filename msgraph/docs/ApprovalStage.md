# ApprovalStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedToMe** | Pointer to **NullableBool** | Indicates whether the stage is assigned to the calling user to review. Read-only. | [optional] 
**DisplayName** | Pointer to **NullableString** | The label provided by the policy creator to identify an approval stage. Read-only. | [optional] 
**Justification** | Pointer to **NullableString** | The justification associated with the approval stage decision. | [optional] 
**ReviewedBy** | Pointer to [**AnyOfmicrosoftGraphIdentity**](anyOf&lt;microsoft.graph.identity&gt;.md) | The identifier of the reviewer. Read-only. | [optional] 
**ReviewedDateTime** | Pointer to **NullableTime** | The date and time when a decision was recorded. The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. | [optional] 
**ReviewResult** | Pointer to **NullableString** | The result of this approval record. Possible values include: NotReviewed, Approved, Denied. | [optional] 
**Status** | Pointer to **NullableString** | The stage status. Possible values: InProgress, Initializing, Completed, Expired. Read-only. | [optional] 

## Methods

### NewApprovalStage

`func NewApprovalStage() *ApprovalStage`

NewApprovalStage instantiates a new ApprovalStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalStageWithDefaults

`func NewApprovalStageWithDefaults() *ApprovalStage`

NewApprovalStageWithDefaults instantiates a new ApprovalStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedToMe

`func (o *ApprovalStage) GetAssignedToMe() bool`

GetAssignedToMe returns the AssignedToMe field if non-nil, zero value otherwise.

### GetAssignedToMeOk

`func (o *ApprovalStage) GetAssignedToMeOk() (*bool, bool)`

GetAssignedToMeOk returns a tuple with the AssignedToMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToMe

`func (o *ApprovalStage) SetAssignedToMe(v bool)`

SetAssignedToMe sets AssignedToMe field to given value.

### HasAssignedToMe

`func (o *ApprovalStage) HasAssignedToMe() bool`

HasAssignedToMe returns a boolean if a field has been set.

### SetAssignedToMeNil

`func (o *ApprovalStage) SetAssignedToMeNil(b bool)`

 SetAssignedToMeNil sets the value for AssignedToMe to be an explicit nil

### UnsetAssignedToMe
`func (o *ApprovalStage) UnsetAssignedToMe()`

UnsetAssignedToMe ensures that no value is present for AssignedToMe, not even an explicit nil
### GetDisplayName

`func (o *ApprovalStage) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApprovalStage) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApprovalStage) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ApprovalStage) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *ApprovalStage) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ApprovalStage) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetJustification

`func (o *ApprovalStage) GetJustification() string`

GetJustification returns the Justification field if non-nil, zero value otherwise.

### GetJustificationOk

`func (o *ApprovalStage) GetJustificationOk() (*string, bool)`

GetJustificationOk returns a tuple with the Justification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJustification

`func (o *ApprovalStage) SetJustification(v string)`

SetJustification sets Justification field to given value.

### HasJustification

`func (o *ApprovalStage) HasJustification() bool`

HasJustification returns a boolean if a field has been set.

### SetJustificationNil

`func (o *ApprovalStage) SetJustificationNil(b bool)`

 SetJustificationNil sets the value for Justification to be an explicit nil

### UnsetJustification
`func (o *ApprovalStage) UnsetJustification()`

UnsetJustification ensures that no value is present for Justification, not even an explicit nil
### GetReviewedBy

`func (o *ApprovalStage) GetReviewedBy() AnyOfmicrosoftGraphIdentity`

GetReviewedBy returns the ReviewedBy field if non-nil, zero value otherwise.

### GetReviewedByOk

`func (o *ApprovalStage) GetReviewedByOk() (*AnyOfmicrosoftGraphIdentity, bool)`

GetReviewedByOk returns a tuple with the ReviewedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedBy

`func (o *ApprovalStage) SetReviewedBy(v AnyOfmicrosoftGraphIdentity)`

SetReviewedBy sets ReviewedBy field to given value.

### HasReviewedBy

`func (o *ApprovalStage) HasReviewedBy() bool`

HasReviewedBy returns a boolean if a field has been set.

### SetReviewedByNil

`func (o *ApprovalStage) SetReviewedByNil(b bool)`

 SetReviewedByNil sets the value for ReviewedBy to be an explicit nil

### UnsetReviewedBy
`func (o *ApprovalStage) UnsetReviewedBy()`

UnsetReviewedBy ensures that no value is present for ReviewedBy, not even an explicit nil
### GetReviewedDateTime

`func (o *ApprovalStage) GetReviewedDateTime() time.Time`

GetReviewedDateTime returns the ReviewedDateTime field if non-nil, zero value otherwise.

### GetReviewedDateTimeOk

`func (o *ApprovalStage) GetReviewedDateTimeOk() (*time.Time, bool)`

GetReviewedDateTimeOk returns a tuple with the ReviewedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedDateTime

`func (o *ApprovalStage) SetReviewedDateTime(v time.Time)`

SetReviewedDateTime sets ReviewedDateTime field to given value.

### HasReviewedDateTime

`func (o *ApprovalStage) HasReviewedDateTime() bool`

HasReviewedDateTime returns a boolean if a field has been set.

### SetReviewedDateTimeNil

`func (o *ApprovalStage) SetReviewedDateTimeNil(b bool)`

 SetReviewedDateTimeNil sets the value for ReviewedDateTime to be an explicit nil

### UnsetReviewedDateTime
`func (o *ApprovalStage) UnsetReviewedDateTime()`

UnsetReviewedDateTime ensures that no value is present for ReviewedDateTime, not even an explicit nil
### GetReviewResult

`func (o *ApprovalStage) GetReviewResult() string`

GetReviewResult returns the ReviewResult field if non-nil, zero value otherwise.

### GetReviewResultOk

`func (o *ApprovalStage) GetReviewResultOk() (*string, bool)`

GetReviewResultOk returns a tuple with the ReviewResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewResult

`func (o *ApprovalStage) SetReviewResult(v string)`

SetReviewResult sets ReviewResult field to given value.

### HasReviewResult

`func (o *ApprovalStage) HasReviewResult() bool`

HasReviewResult returns a boolean if a field has been set.

### SetReviewResultNil

`func (o *ApprovalStage) SetReviewResultNil(b bool)`

 SetReviewResultNil sets the value for ReviewResult to be an explicit nil

### UnsetReviewResult
`func (o *ApprovalStage) UnsetReviewResult()`

UnsetReviewResult ensures that no value is present for ReviewResult, not even an explicit nil
### GetStatus

`func (o *ApprovalStage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApprovalStage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApprovalStage) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApprovalStage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ApprovalStage) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ApprovalStage) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


