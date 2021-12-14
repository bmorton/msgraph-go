# MicrosoftGraphDirectoryAudit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ActivityDateTime** | Pointer to **time.Time** | Indicates the date and time the activity was performed. The Timestamp type is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. | [optional] 
**ActivityDisplayName** | Pointer to **string** | Indicates the activity name or the operation name (examples: &#39;Create User&#39; and &#39;Add member to group&#39;). For full list, see Azure AD activity list. | [optional] 
**AdditionalDetails** | Pointer to [**[]AnyOfmicrosoftGraphKeyValue**](AnyOfmicrosoftGraphKeyValue.md) | Indicates additional details on the activity. | [optional] 
**Category** | Pointer to **string** | Indicates which resource category that&#39;s targeted by the activity. (For example: User Management, Group Management etc..) | [optional] 
**CorrelationId** | Pointer to **NullableString** | Indicates a unique ID that helps correlate activities that span across various services. Can be used to trace logs across services. | [optional] 
**InitiatedBy** | Pointer to [**MicrosoftGraphAuditActivityInitiator**](MicrosoftGraphAuditActivityInitiator.md) |  | [optional] 
**LoggedByService** | Pointer to **NullableString** | Indicates information on which service initiated the activity (For example: Self-service Password Management, Core Directory, B2C, Invited Users, Microsoft Identity Manager, Privileged Identity Management. | [optional] 
**OperationType** | Pointer to **NullableString** |  | [optional] 
**Result** | Pointer to [**AnyOfmicrosoftGraphOperationResult**](anyOf&lt;microsoft.graph.operationResult&gt;.md) | Indicates the result of the activity. Possible values are: success, failure, timeout, unknownFutureValue. | [optional] 
**ResultReason** | Pointer to **NullableString** | Indicates the reason for failure if the result is failure or timeout. | [optional] 
**TargetResources** | Pointer to [**[]AnyOfmicrosoftGraphTargetResource**](AnyOfmicrosoftGraphTargetResource.md) | Indicates information on which resource was changed due to the activity. Target Resource Type can be User, Device, Directory, App, Role, Group, Policy or Other. | [optional] 

## Methods

### NewMicrosoftGraphDirectoryAudit

`func NewMicrosoftGraphDirectoryAudit() *MicrosoftGraphDirectoryAudit`

NewMicrosoftGraphDirectoryAudit instantiates a new MicrosoftGraphDirectoryAudit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDirectoryAuditWithDefaults

`func NewMicrosoftGraphDirectoryAuditWithDefaults() *MicrosoftGraphDirectoryAudit`

NewMicrosoftGraphDirectoryAuditWithDefaults instantiates a new MicrosoftGraphDirectoryAudit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDirectoryAudit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDirectoryAudit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDirectoryAudit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDirectoryAudit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActivityDateTime

`func (o *MicrosoftGraphDirectoryAudit) GetActivityDateTime() time.Time`

GetActivityDateTime returns the ActivityDateTime field if non-nil, zero value otherwise.

### GetActivityDateTimeOk

`func (o *MicrosoftGraphDirectoryAudit) GetActivityDateTimeOk() (*time.Time, bool)`

GetActivityDateTimeOk returns a tuple with the ActivityDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityDateTime

`func (o *MicrosoftGraphDirectoryAudit) SetActivityDateTime(v time.Time)`

SetActivityDateTime sets ActivityDateTime field to given value.

### HasActivityDateTime

`func (o *MicrosoftGraphDirectoryAudit) HasActivityDateTime() bool`

HasActivityDateTime returns a boolean if a field has been set.

### GetActivityDisplayName

`func (o *MicrosoftGraphDirectoryAudit) GetActivityDisplayName() string`

GetActivityDisplayName returns the ActivityDisplayName field if non-nil, zero value otherwise.

### GetActivityDisplayNameOk

`func (o *MicrosoftGraphDirectoryAudit) GetActivityDisplayNameOk() (*string, bool)`

GetActivityDisplayNameOk returns a tuple with the ActivityDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityDisplayName

`func (o *MicrosoftGraphDirectoryAudit) SetActivityDisplayName(v string)`

SetActivityDisplayName sets ActivityDisplayName field to given value.

### HasActivityDisplayName

`func (o *MicrosoftGraphDirectoryAudit) HasActivityDisplayName() bool`

HasActivityDisplayName returns a boolean if a field has been set.

### GetAdditionalDetails

`func (o *MicrosoftGraphDirectoryAudit) GetAdditionalDetails() []*AnyOfmicrosoftGraphKeyValue`

GetAdditionalDetails returns the AdditionalDetails field if non-nil, zero value otherwise.

### GetAdditionalDetailsOk

`func (o *MicrosoftGraphDirectoryAudit) GetAdditionalDetailsOk() (*[]*AnyOfmicrosoftGraphKeyValue, bool)`

GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDetails

`func (o *MicrosoftGraphDirectoryAudit) SetAdditionalDetails(v []*AnyOfmicrosoftGraphKeyValue)`

SetAdditionalDetails sets AdditionalDetails field to given value.

### HasAdditionalDetails

`func (o *MicrosoftGraphDirectoryAudit) HasAdditionalDetails() bool`

HasAdditionalDetails returns a boolean if a field has been set.

### GetCategory

`func (o *MicrosoftGraphDirectoryAudit) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MicrosoftGraphDirectoryAudit) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *MicrosoftGraphDirectoryAudit) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *MicrosoftGraphDirectoryAudit) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCorrelationId

`func (o *MicrosoftGraphDirectoryAudit) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *MicrosoftGraphDirectoryAudit) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *MicrosoftGraphDirectoryAudit) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *MicrosoftGraphDirectoryAudit) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *MicrosoftGraphDirectoryAudit) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *MicrosoftGraphDirectoryAudit) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetInitiatedBy

`func (o *MicrosoftGraphDirectoryAudit) GetInitiatedBy() MicrosoftGraphAuditActivityInitiator`

GetInitiatedBy returns the InitiatedBy field if non-nil, zero value otherwise.

### GetInitiatedByOk

`func (o *MicrosoftGraphDirectoryAudit) GetInitiatedByOk() (*MicrosoftGraphAuditActivityInitiator, bool)`

GetInitiatedByOk returns a tuple with the InitiatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatedBy

`func (o *MicrosoftGraphDirectoryAudit) SetInitiatedBy(v MicrosoftGraphAuditActivityInitiator)`

SetInitiatedBy sets InitiatedBy field to given value.

### HasInitiatedBy

`func (o *MicrosoftGraphDirectoryAudit) HasInitiatedBy() bool`

HasInitiatedBy returns a boolean if a field has been set.

### GetLoggedByService

`func (o *MicrosoftGraphDirectoryAudit) GetLoggedByService() string`

GetLoggedByService returns the LoggedByService field if non-nil, zero value otherwise.

### GetLoggedByServiceOk

`func (o *MicrosoftGraphDirectoryAudit) GetLoggedByServiceOk() (*string, bool)`

GetLoggedByServiceOk returns a tuple with the LoggedByService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedByService

`func (o *MicrosoftGraphDirectoryAudit) SetLoggedByService(v string)`

SetLoggedByService sets LoggedByService field to given value.

### HasLoggedByService

`func (o *MicrosoftGraphDirectoryAudit) HasLoggedByService() bool`

HasLoggedByService returns a boolean if a field has been set.

### SetLoggedByServiceNil

`func (o *MicrosoftGraphDirectoryAudit) SetLoggedByServiceNil(b bool)`

 SetLoggedByServiceNil sets the value for LoggedByService to be an explicit nil

### UnsetLoggedByService
`func (o *MicrosoftGraphDirectoryAudit) UnsetLoggedByService()`

UnsetLoggedByService ensures that no value is present for LoggedByService, not even an explicit nil
### GetOperationType

`func (o *MicrosoftGraphDirectoryAudit) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *MicrosoftGraphDirectoryAudit) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *MicrosoftGraphDirectoryAudit) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *MicrosoftGraphDirectoryAudit) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### SetOperationTypeNil

`func (o *MicrosoftGraphDirectoryAudit) SetOperationTypeNil(b bool)`

 SetOperationTypeNil sets the value for OperationType to be an explicit nil

### UnsetOperationType
`func (o *MicrosoftGraphDirectoryAudit) UnsetOperationType()`

UnsetOperationType ensures that no value is present for OperationType, not even an explicit nil
### GetResult

`func (o *MicrosoftGraphDirectoryAudit) GetResult() AnyOfmicrosoftGraphOperationResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *MicrosoftGraphDirectoryAudit) GetResultOk() (*AnyOfmicrosoftGraphOperationResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *MicrosoftGraphDirectoryAudit) SetResult(v AnyOfmicrosoftGraphOperationResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *MicrosoftGraphDirectoryAudit) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *MicrosoftGraphDirectoryAudit) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *MicrosoftGraphDirectoryAudit) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetResultReason

`func (o *MicrosoftGraphDirectoryAudit) GetResultReason() string`

GetResultReason returns the ResultReason field if non-nil, zero value otherwise.

### GetResultReasonOk

`func (o *MicrosoftGraphDirectoryAudit) GetResultReasonOk() (*string, bool)`

GetResultReasonOk returns a tuple with the ResultReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultReason

`func (o *MicrosoftGraphDirectoryAudit) SetResultReason(v string)`

SetResultReason sets ResultReason field to given value.

### HasResultReason

`func (o *MicrosoftGraphDirectoryAudit) HasResultReason() bool`

HasResultReason returns a boolean if a field has been set.

### SetResultReasonNil

`func (o *MicrosoftGraphDirectoryAudit) SetResultReasonNil(b bool)`

 SetResultReasonNil sets the value for ResultReason to be an explicit nil

### UnsetResultReason
`func (o *MicrosoftGraphDirectoryAudit) UnsetResultReason()`

UnsetResultReason ensures that no value is present for ResultReason, not even an explicit nil
### GetTargetResources

`func (o *MicrosoftGraphDirectoryAudit) GetTargetResources() []*AnyOfmicrosoftGraphTargetResource`

GetTargetResources returns the TargetResources field if non-nil, zero value otherwise.

### GetTargetResourcesOk

`func (o *MicrosoftGraphDirectoryAudit) GetTargetResourcesOk() (*[]*AnyOfmicrosoftGraphTargetResource, bool)`

GetTargetResourcesOk returns a tuple with the TargetResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResources

`func (o *MicrosoftGraphDirectoryAudit) SetTargetResources(v []*AnyOfmicrosoftGraphTargetResource)`

SetTargetResources sets TargetResources field to given value.

### HasTargetResources

`func (o *MicrosoftGraphDirectoryAudit) HasTargetResources() bool`

HasTargetResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


