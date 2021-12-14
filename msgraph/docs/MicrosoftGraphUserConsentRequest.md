# MicrosoftGraphUserConsentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ApprovalId** | Pointer to **NullableString** | The identifier of the approval of the request. | [optional] 
**CompletedDateTime** | Pointer to **NullableTime** | The request completion date time. | [optional] 
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | The user who created this request. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The request creation date time. | [optional] 
**CustomData** | Pointer to **NullableString** | Free text field to define any custom data for the request. Not used. | [optional] 
**Status** | Pointer to **string** | The status of the request. Not nullable. The possible values are: Canceled, Denied, Failed, Granted, PendingAdminDecision, PendingApproval, PendingProvisioning, PendingScheduleCreation, Provisioned, Revoked, and ScheduleCreated. Not nullable. | [optional] 
**Reason** | Pointer to **NullableString** | The user&#39;s justification for requiring access to the app. Supports $filter (eq only) and $orderby. | [optional] 
**Approval** | Pointer to [**AnyOfmicrosoftGraphApproval**](anyOf&lt;microsoft.graph.approval&gt;.md) | Approval decisions associated with a request. | [optional] 

## Methods

### NewMicrosoftGraphUserConsentRequest

`func NewMicrosoftGraphUserConsentRequest() *MicrosoftGraphUserConsentRequest`

NewMicrosoftGraphUserConsentRequest instantiates a new MicrosoftGraphUserConsentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphUserConsentRequestWithDefaults

`func NewMicrosoftGraphUserConsentRequestWithDefaults() *MicrosoftGraphUserConsentRequest`

NewMicrosoftGraphUserConsentRequestWithDefaults instantiates a new MicrosoftGraphUserConsentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphUserConsentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphUserConsentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphUserConsentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphUserConsentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetApprovalId

`func (o *MicrosoftGraphUserConsentRequest) GetApprovalId() string`

GetApprovalId returns the ApprovalId field if non-nil, zero value otherwise.

### GetApprovalIdOk

`func (o *MicrosoftGraphUserConsentRequest) GetApprovalIdOk() (*string, bool)`

GetApprovalIdOk returns a tuple with the ApprovalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalId

`func (o *MicrosoftGraphUserConsentRequest) SetApprovalId(v string)`

SetApprovalId sets ApprovalId field to given value.

### HasApprovalId

`func (o *MicrosoftGraphUserConsentRequest) HasApprovalId() bool`

HasApprovalId returns a boolean if a field has been set.

### SetApprovalIdNil

`func (o *MicrosoftGraphUserConsentRequest) SetApprovalIdNil(b bool)`

 SetApprovalIdNil sets the value for ApprovalId to be an explicit nil

### UnsetApprovalId
`func (o *MicrosoftGraphUserConsentRequest) UnsetApprovalId()`

UnsetApprovalId ensures that no value is present for ApprovalId, not even an explicit nil
### GetCompletedDateTime

`func (o *MicrosoftGraphUserConsentRequest) GetCompletedDateTime() time.Time`

GetCompletedDateTime returns the CompletedDateTime field if non-nil, zero value otherwise.

### GetCompletedDateTimeOk

`func (o *MicrosoftGraphUserConsentRequest) GetCompletedDateTimeOk() (*time.Time, bool)`

GetCompletedDateTimeOk returns a tuple with the CompletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDateTime

`func (o *MicrosoftGraphUserConsentRequest) SetCompletedDateTime(v time.Time)`

SetCompletedDateTime sets CompletedDateTime field to given value.

### HasCompletedDateTime

`func (o *MicrosoftGraphUserConsentRequest) HasCompletedDateTime() bool`

HasCompletedDateTime returns a boolean if a field has been set.

### SetCompletedDateTimeNil

`func (o *MicrosoftGraphUserConsentRequest) SetCompletedDateTimeNil(b bool)`

 SetCompletedDateTimeNil sets the value for CompletedDateTime to be an explicit nil

### UnsetCompletedDateTime
`func (o *MicrosoftGraphUserConsentRequest) UnsetCompletedDateTime()`

UnsetCompletedDateTime ensures that no value is present for CompletedDateTime, not even an explicit nil
### GetCreatedBy

`func (o *MicrosoftGraphUserConsentRequest) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MicrosoftGraphUserConsentRequest) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MicrosoftGraphUserConsentRequest) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MicrosoftGraphUserConsentRequest) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *MicrosoftGraphUserConsentRequest) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *MicrosoftGraphUserConsentRequest) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphUserConsentRequest) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphUserConsentRequest) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphUserConsentRequest) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphUserConsentRequest) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphUserConsentRequest) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphUserConsentRequest) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetCustomData

`func (o *MicrosoftGraphUserConsentRequest) GetCustomData() string`

GetCustomData returns the CustomData field if non-nil, zero value otherwise.

### GetCustomDataOk

`func (o *MicrosoftGraphUserConsentRequest) GetCustomDataOk() (*string, bool)`

GetCustomDataOk returns a tuple with the CustomData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomData

`func (o *MicrosoftGraphUserConsentRequest) SetCustomData(v string)`

SetCustomData sets CustomData field to given value.

### HasCustomData

`func (o *MicrosoftGraphUserConsentRequest) HasCustomData() bool`

HasCustomData returns a boolean if a field has been set.

### SetCustomDataNil

`func (o *MicrosoftGraphUserConsentRequest) SetCustomDataNil(b bool)`

 SetCustomDataNil sets the value for CustomData to be an explicit nil

### UnsetCustomData
`func (o *MicrosoftGraphUserConsentRequest) UnsetCustomData()`

UnsetCustomData ensures that no value is present for CustomData, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphUserConsentRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphUserConsentRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphUserConsentRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphUserConsentRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *MicrosoftGraphUserConsentRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MicrosoftGraphUserConsentRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MicrosoftGraphUserConsentRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MicrosoftGraphUserConsentRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *MicrosoftGraphUserConsentRequest) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *MicrosoftGraphUserConsentRequest) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetApproval

`func (o *MicrosoftGraphUserConsentRequest) GetApproval() AnyOfmicrosoftGraphApproval`

GetApproval returns the Approval field if non-nil, zero value otherwise.

### GetApprovalOk

`func (o *MicrosoftGraphUserConsentRequest) GetApprovalOk() (*AnyOfmicrosoftGraphApproval, bool)`

GetApprovalOk returns a tuple with the Approval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproval

`func (o *MicrosoftGraphUserConsentRequest) SetApproval(v AnyOfmicrosoftGraphApproval)`

SetApproval sets Approval field to given value.

### HasApproval

`func (o *MicrosoftGraphUserConsentRequest) HasApproval() bool`

HasApproval returns a boolean if a field has been set.

### SetApprovalNil

`func (o *MicrosoftGraphUserConsentRequest) SetApprovalNil(b bool)`

 SetApprovalNil sets the value for Approval to be an explicit nil

### UnsetApproval
`func (o *MicrosoftGraphUserConsentRequest) UnsetApproval()`

UnsetApproval ensures that no value is present for Approval, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


