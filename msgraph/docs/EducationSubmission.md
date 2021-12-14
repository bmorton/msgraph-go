# EducationSubmission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipient** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Who this submission is assigned to. | [optional] 
**ResourcesFolderUrl** | Pointer to **NullableString** | Folder where all file resources for this submission need to be stored. | [optional] 
**ReturnedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | User who moved the status of this submission to returned. | [optional] 
**ReturnedDateTime** | Pointer to **NullableTime** | Moment in time when the submission was returned. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphEducationSubmissionStatus**](anyOf&lt;microsoft.graph.educationSubmissionStatus&gt;.md) | Read-Only. Possible values are: working, submitted, released, returned. | [optional] 
**SubmittedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | User who moved the resource into the submitted state. | [optional] 
**SubmittedDateTime** | Pointer to **NullableTime** | Moment in time when the submission was moved into the submitted state. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**UnsubmittedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | User who moved the resource from submitted into the working state. | [optional] 
**UnsubmittedDateTime** | Pointer to **NullableTime** | Moment in time when the submission was moved from submitted into the working state. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**Outcomes** | Pointer to [**[]MicrosoftGraphEducationOutcome**](MicrosoftGraphEducationOutcome.md) | Read-Write. Nullable. | [optional] 
**Resources** | Pointer to [**[]MicrosoftGraphEducationSubmissionResource**](MicrosoftGraphEducationSubmissionResource.md) | Nullable. | [optional] 
**SubmittedResources** | Pointer to [**[]MicrosoftGraphEducationSubmissionResource**](MicrosoftGraphEducationSubmissionResource.md) | Read-only. Nullable. | [optional] 

## Methods

### NewEducationSubmission

`func NewEducationSubmission() *EducationSubmission`

NewEducationSubmission instantiates a new EducationSubmission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEducationSubmissionWithDefaults

`func NewEducationSubmissionWithDefaults() *EducationSubmission`

NewEducationSubmissionWithDefaults instantiates a new EducationSubmission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipient

`func (o *EducationSubmission) GetRecipient() AnyOfobject`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *EducationSubmission) GetRecipientOk() (*AnyOfobject, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *EducationSubmission) SetRecipient(v AnyOfobject)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *EducationSubmission) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### SetRecipientNil

`func (o *EducationSubmission) SetRecipientNil(b bool)`

 SetRecipientNil sets the value for Recipient to be an explicit nil

### UnsetRecipient
`func (o *EducationSubmission) UnsetRecipient()`

UnsetRecipient ensures that no value is present for Recipient, not even an explicit nil
### GetResourcesFolderUrl

`func (o *EducationSubmission) GetResourcesFolderUrl() string`

GetResourcesFolderUrl returns the ResourcesFolderUrl field if non-nil, zero value otherwise.

### GetResourcesFolderUrlOk

`func (o *EducationSubmission) GetResourcesFolderUrlOk() (*string, bool)`

GetResourcesFolderUrlOk returns a tuple with the ResourcesFolderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesFolderUrl

`func (o *EducationSubmission) SetResourcesFolderUrl(v string)`

SetResourcesFolderUrl sets ResourcesFolderUrl field to given value.

### HasResourcesFolderUrl

`func (o *EducationSubmission) HasResourcesFolderUrl() bool`

HasResourcesFolderUrl returns a boolean if a field has been set.

### SetResourcesFolderUrlNil

`func (o *EducationSubmission) SetResourcesFolderUrlNil(b bool)`

 SetResourcesFolderUrlNil sets the value for ResourcesFolderUrl to be an explicit nil

### UnsetResourcesFolderUrl
`func (o *EducationSubmission) UnsetResourcesFolderUrl()`

UnsetResourcesFolderUrl ensures that no value is present for ResourcesFolderUrl, not even an explicit nil
### GetReturnedBy

`func (o *EducationSubmission) GetReturnedBy() AnyOfmicrosoftGraphIdentitySet`

GetReturnedBy returns the ReturnedBy field if non-nil, zero value otherwise.

### GetReturnedByOk

`func (o *EducationSubmission) GetReturnedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetReturnedByOk returns a tuple with the ReturnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedBy

`func (o *EducationSubmission) SetReturnedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetReturnedBy sets ReturnedBy field to given value.

### HasReturnedBy

`func (o *EducationSubmission) HasReturnedBy() bool`

HasReturnedBy returns a boolean if a field has been set.

### SetReturnedByNil

`func (o *EducationSubmission) SetReturnedByNil(b bool)`

 SetReturnedByNil sets the value for ReturnedBy to be an explicit nil

### UnsetReturnedBy
`func (o *EducationSubmission) UnsetReturnedBy()`

UnsetReturnedBy ensures that no value is present for ReturnedBy, not even an explicit nil
### GetReturnedDateTime

`func (o *EducationSubmission) GetReturnedDateTime() time.Time`

GetReturnedDateTime returns the ReturnedDateTime field if non-nil, zero value otherwise.

### GetReturnedDateTimeOk

`func (o *EducationSubmission) GetReturnedDateTimeOk() (*time.Time, bool)`

GetReturnedDateTimeOk returns a tuple with the ReturnedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedDateTime

`func (o *EducationSubmission) SetReturnedDateTime(v time.Time)`

SetReturnedDateTime sets ReturnedDateTime field to given value.

### HasReturnedDateTime

`func (o *EducationSubmission) HasReturnedDateTime() bool`

HasReturnedDateTime returns a boolean if a field has been set.

### SetReturnedDateTimeNil

`func (o *EducationSubmission) SetReturnedDateTimeNil(b bool)`

 SetReturnedDateTimeNil sets the value for ReturnedDateTime to be an explicit nil

### UnsetReturnedDateTime
`func (o *EducationSubmission) UnsetReturnedDateTime()`

UnsetReturnedDateTime ensures that no value is present for ReturnedDateTime, not even an explicit nil
### GetStatus

`func (o *EducationSubmission) GetStatus() AnyOfmicrosoftGraphEducationSubmissionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EducationSubmission) GetStatusOk() (*AnyOfmicrosoftGraphEducationSubmissionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EducationSubmission) SetStatus(v AnyOfmicrosoftGraphEducationSubmissionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EducationSubmission) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *EducationSubmission) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *EducationSubmission) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetSubmittedBy

`func (o *EducationSubmission) GetSubmittedBy() AnyOfmicrosoftGraphIdentitySet`

GetSubmittedBy returns the SubmittedBy field if non-nil, zero value otherwise.

### GetSubmittedByOk

`func (o *EducationSubmission) GetSubmittedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetSubmittedByOk returns a tuple with the SubmittedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedBy

`func (o *EducationSubmission) SetSubmittedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetSubmittedBy sets SubmittedBy field to given value.

### HasSubmittedBy

`func (o *EducationSubmission) HasSubmittedBy() bool`

HasSubmittedBy returns a boolean if a field has been set.

### SetSubmittedByNil

`func (o *EducationSubmission) SetSubmittedByNil(b bool)`

 SetSubmittedByNil sets the value for SubmittedBy to be an explicit nil

### UnsetSubmittedBy
`func (o *EducationSubmission) UnsetSubmittedBy()`

UnsetSubmittedBy ensures that no value is present for SubmittedBy, not even an explicit nil
### GetSubmittedDateTime

`func (o *EducationSubmission) GetSubmittedDateTime() time.Time`

GetSubmittedDateTime returns the SubmittedDateTime field if non-nil, zero value otherwise.

### GetSubmittedDateTimeOk

`func (o *EducationSubmission) GetSubmittedDateTimeOk() (*time.Time, bool)`

GetSubmittedDateTimeOk returns a tuple with the SubmittedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedDateTime

`func (o *EducationSubmission) SetSubmittedDateTime(v time.Time)`

SetSubmittedDateTime sets SubmittedDateTime field to given value.

### HasSubmittedDateTime

`func (o *EducationSubmission) HasSubmittedDateTime() bool`

HasSubmittedDateTime returns a boolean if a field has been set.

### SetSubmittedDateTimeNil

`func (o *EducationSubmission) SetSubmittedDateTimeNil(b bool)`

 SetSubmittedDateTimeNil sets the value for SubmittedDateTime to be an explicit nil

### UnsetSubmittedDateTime
`func (o *EducationSubmission) UnsetSubmittedDateTime()`

UnsetSubmittedDateTime ensures that no value is present for SubmittedDateTime, not even an explicit nil
### GetUnsubmittedBy

`func (o *EducationSubmission) GetUnsubmittedBy() AnyOfmicrosoftGraphIdentitySet`

GetUnsubmittedBy returns the UnsubmittedBy field if non-nil, zero value otherwise.

### GetUnsubmittedByOk

`func (o *EducationSubmission) GetUnsubmittedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetUnsubmittedByOk returns a tuple with the UnsubmittedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubmittedBy

`func (o *EducationSubmission) SetUnsubmittedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetUnsubmittedBy sets UnsubmittedBy field to given value.

### HasUnsubmittedBy

`func (o *EducationSubmission) HasUnsubmittedBy() bool`

HasUnsubmittedBy returns a boolean if a field has been set.

### SetUnsubmittedByNil

`func (o *EducationSubmission) SetUnsubmittedByNil(b bool)`

 SetUnsubmittedByNil sets the value for UnsubmittedBy to be an explicit nil

### UnsetUnsubmittedBy
`func (o *EducationSubmission) UnsetUnsubmittedBy()`

UnsetUnsubmittedBy ensures that no value is present for UnsubmittedBy, not even an explicit nil
### GetUnsubmittedDateTime

`func (o *EducationSubmission) GetUnsubmittedDateTime() time.Time`

GetUnsubmittedDateTime returns the UnsubmittedDateTime field if non-nil, zero value otherwise.

### GetUnsubmittedDateTimeOk

`func (o *EducationSubmission) GetUnsubmittedDateTimeOk() (*time.Time, bool)`

GetUnsubmittedDateTimeOk returns a tuple with the UnsubmittedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubmittedDateTime

`func (o *EducationSubmission) SetUnsubmittedDateTime(v time.Time)`

SetUnsubmittedDateTime sets UnsubmittedDateTime field to given value.

### HasUnsubmittedDateTime

`func (o *EducationSubmission) HasUnsubmittedDateTime() bool`

HasUnsubmittedDateTime returns a boolean if a field has been set.

### SetUnsubmittedDateTimeNil

`func (o *EducationSubmission) SetUnsubmittedDateTimeNil(b bool)`

 SetUnsubmittedDateTimeNil sets the value for UnsubmittedDateTime to be an explicit nil

### UnsetUnsubmittedDateTime
`func (o *EducationSubmission) UnsetUnsubmittedDateTime()`

UnsetUnsubmittedDateTime ensures that no value is present for UnsubmittedDateTime, not even an explicit nil
### GetOutcomes

`func (o *EducationSubmission) GetOutcomes() []MicrosoftGraphEducationOutcome`

GetOutcomes returns the Outcomes field if non-nil, zero value otherwise.

### GetOutcomesOk

`func (o *EducationSubmission) GetOutcomesOk() (*[]MicrosoftGraphEducationOutcome, bool)`

GetOutcomesOk returns a tuple with the Outcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcomes

`func (o *EducationSubmission) SetOutcomes(v []MicrosoftGraphEducationOutcome)`

SetOutcomes sets Outcomes field to given value.

### HasOutcomes

`func (o *EducationSubmission) HasOutcomes() bool`

HasOutcomes returns a boolean if a field has been set.

### GetResources

`func (o *EducationSubmission) GetResources() []MicrosoftGraphEducationSubmissionResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *EducationSubmission) GetResourcesOk() (*[]MicrosoftGraphEducationSubmissionResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *EducationSubmission) SetResources(v []MicrosoftGraphEducationSubmissionResource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *EducationSubmission) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetSubmittedResources

`func (o *EducationSubmission) GetSubmittedResources() []MicrosoftGraphEducationSubmissionResource`

GetSubmittedResources returns the SubmittedResources field if non-nil, zero value otherwise.

### GetSubmittedResourcesOk

`func (o *EducationSubmission) GetSubmittedResourcesOk() (*[]MicrosoftGraphEducationSubmissionResource, bool)`

GetSubmittedResourcesOk returns a tuple with the SubmittedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedResources

`func (o *EducationSubmission) SetSubmittedResources(v []MicrosoftGraphEducationSubmissionResource)`

SetSubmittedResources sets SubmittedResources field to given value.

### HasSubmittedResources

`func (o *EducationSubmission) HasSubmittedResources() bool`

HasSubmittedResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


