# DirectoryAudit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewDirectoryAudit

`func NewDirectoryAudit() *DirectoryAudit`

NewDirectoryAudit instantiates a new DirectoryAudit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryAuditWithDefaults

`func NewDirectoryAuditWithDefaults() *DirectoryAudit`

NewDirectoryAuditWithDefaults instantiates a new DirectoryAudit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityDateTime

`func (o *DirectoryAudit) GetActivityDateTime() time.Time`

GetActivityDateTime returns the ActivityDateTime field if non-nil, zero value otherwise.

### GetActivityDateTimeOk

`func (o *DirectoryAudit) GetActivityDateTimeOk() (*time.Time, bool)`

GetActivityDateTimeOk returns a tuple with the ActivityDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityDateTime

`func (o *DirectoryAudit) SetActivityDateTime(v time.Time)`

SetActivityDateTime sets ActivityDateTime field to given value.

### HasActivityDateTime

`func (o *DirectoryAudit) HasActivityDateTime() bool`

HasActivityDateTime returns a boolean if a field has been set.

### GetActivityDisplayName

`func (o *DirectoryAudit) GetActivityDisplayName() string`

GetActivityDisplayName returns the ActivityDisplayName field if non-nil, zero value otherwise.

### GetActivityDisplayNameOk

`func (o *DirectoryAudit) GetActivityDisplayNameOk() (*string, bool)`

GetActivityDisplayNameOk returns a tuple with the ActivityDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityDisplayName

`func (o *DirectoryAudit) SetActivityDisplayName(v string)`

SetActivityDisplayName sets ActivityDisplayName field to given value.

### HasActivityDisplayName

`func (o *DirectoryAudit) HasActivityDisplayName() bool`

HasActivityDisplayName returns a boolean if a field has been set.

### GetAdditionalDetails

`func (o *DirectoryAudit) GetAdditionalDetails() []*AnyOfmicrosoftGraphKeyValue`

GetAdditionalDetails returns the AdditionalDetails field if non-nil, zero value otherwise.

### GetAdditionalDetailsOk

`func (o *DirectoryAudit) GetAdditionalDetailsOk() (*[]*AnyOfmicrosoftGraphKeyValue, bool)`

GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDetails

`func (o *DirectoryAudit) SetAdditionalDetails(v []*AnyOfmicrosoftGraphKeyValue)`

SetAdditionalDetails sets AdditionalDetails field to given value.

### HasAdditionalDetails

`func (o *DirectoryAudit) HasAdditionalDetails() bool`

HasAdditionalDetails returns a boolean if a field has been set.

### GetCategory

`func (o *DirectoryAudit) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *DirectoryAudit) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *DirectoryAudit) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *DirectoryAudit) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCorrelationId

`func (o *DirectoryAudit) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *DirectoryAudit) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *DirectoryAudit) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *DirectoryAudit) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### SetCorrelationIdNil

`func (o *DirectoryAudit) SetCorrelationIdNil(b bool)`

 SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil

### UnsetCorrelationId
`func (o *DirectoryAudit) UnsetCorrelationId()`

UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
### GetInitiatedBy

`func (o *DirectoryAudit) GetInitiatedBy() MicrosoftGraphAuditActivityInitiator`

GetInitiatedBy returns the InitiatedBy field if non-nil, zero value otherwise.

### GetInitiatedByOk

`func (o *DirectoryAudit) GetInitiatedByOk() (*MicrosoftGraphAuditActivityInitiator, bool)`

GetInitiatedByOk returns a tuple with the InitiatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatedBy

`func (o *DirectoryAudit) SetInitiatedBy(v MicrosoftGraphAuditActivityInitiator)`

SetInitiatedBy sets InitiatedBy field to given value.

### HasInitiatedBy

`func (o *DirectoryAudit) HasInitiatedBy() bool`

HasInitiatedBy returns a boolean if a field has been set.

### GetLoggedByService

`func (o *DirectoryAudit) GetLoggedByService() string`

GetLoggedByService returns the LoggedByService field if non-nil, zero value otherwise.

### GetLoggedByServiceOk

`func (o *DirectoryAudit) GetLoggedByServiceOk() (*string, bool)`

GetLoggedByServiceOk returns a tuple with the LoggedByService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoggedByService

`func (o *DirectoryAudit) SetLoggedByService(v string)`

SetLoggedByService sets LoggedByService field to given value.

### HasLoggedByService

`func (o *DirectoryAudit) HasLoggedByService() bool`

HasLoggedByService returns a boolean if a field has been set.

### SetLoggedByServiceNil

`func (o *DirectoryAudit) SetLoggedByServiceNil(b bool)`

 SetLoggedByServiceNil sets the value for LoggedByService to be an explicit nil

### UnsetLoggedByService
`func (o *DirectoryAudit) UnsetLoggedByService()`

UnsetLoggedByService ensures that no value is present for LoggedByService, not even an explicit nil
### GetOperationType

`func (o *DirectoryAudit) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *DirectoryAudit) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *DirectoryAudit) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *DirectoryAudit) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### SetOperationTypeNil

`func (o *DirectoryAudit) SetOperationTypeNil(b bool)`

 SetOperationTypeNil sets the value for OperationType to be an explicit nil

### UnsetOperationType
`func (o *DirectoryAudit) UnsetOperationType()`

UnsetOperationType ensures that no value is present for OperationType, not even an explicit nil
### GetResult

`func (o *DirectoryAudit) GetResult() AnyOfmicrosoftGraphOperationResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *DirectoryAudit) GetResultOk() (*AnyOfmicrosoftGraphOperationResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *DirectoryAudit) SetResult(v AnyOfmicrosoftGraphOperationResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *DirectoryAudit) HasResult() bool`

HasResult returns a boolean if a field has been set.

### SetResultNil

`func (o *DirectoryAudit) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *DirectoryAudit) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetResultReason

`func (o *DirectoryAudit) GetResultReason() string`

GetResultReason returns the ResultReason field if non-nil, zero value otherwise.

### GetResultReasonOk

`func (o *DirectoryAudit) GetResultReasonOk() (*string, bool)`

GetResultReasonOk returns a tuple with the ResultReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultReason

`func (o *DirectoryAudit) SetResultReason(v string)`

SetResultReason sets ResultReason field to given value.

### HasResultReason

`func (o *DirectoryAudit) HasResultReason() bool`

HasResultReason returns a boolean if a field has been set.

### SetResultReasonNil

`func (o *DirectoryAudit) SetResultReasonNil(b bool)`

 SetResultReasonNil sets the value for ResultReason to be an explicit nil

### UnsetResultReason
`func (o *DirectoryAudit) UnsetResultReason()`

UnsetResultReason ensures that no value is present for ResultReason, not even an explicit nil
### GetTargetResources

`func (o *DirectoryAudit) GetTargetResources() []*AnyOfmicrosoftGraphTargetResource`

GetTargetResources returns the TargetResources field if non-nil, zero value otherwise.

### GetTargetResourcesOk

`func (o *DirectoryAudit) GetTargetResourcesOk() (*[]*AnyOfmicrosoftGraphTargetResource, bool)`

GetTargetResourcesOk returns a tuple with the TargetResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResources

`func (o *DirectoryAudit) SetTargetResources(v []*AnyOfmicrosoftGraphTargetResource)`

SetTargetResources sets TargetResources field to given value.

### HasTargetResources

`func (o *DirectoryAudit) HasTargetResources() bool`

HasTargetResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


