# MicrosoftGraphProvisioningObjectSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ActivityDateTime** | Pointer to **time.Time** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**ChangeId** | Pointer to **NullableString** | Unique ID of this change in this cycle. | [optional] 
**CycleId** | Pointer to **NullableString** | Unique ID per job iteration. | [optional] 
**DurationInMilliseconds** | Pointer to **NullableInt32** | Indicates how long this provisioning action took to finish. Measured in milliseconds. | [optional] 
**InitiatedBy** | Pointer to [**AnyOfmicrosoftGraphInitiator**](anyOf&lt;microsoft.graph.initiator&gt;.md) | Details of who initiated this provisioning. | [optional] 
**JobId** | Pointer to **NullableString** | The unique ID for the whole provisioning job. | [optional] 
**ModifiedProperties** | Pointer to [**[]AnyOfmicrosoftGraphModifiedProperty**](AnyOfmicrosoftGraphModifiedProperty.md) | Details of each property that was modified in this provisioning action on this object. | [optional] 
**ProvisioningAction** | Pointer to [**AnyOfmicrosoftGraphProvisioningAction**](anyOf&lt;microsoft.graph.provisioningAction&gt;.md) | Indicates the activity name or the operation name. Possible values are: create, update, delete, stageddelete, disable, other and unknownFutureValue. For a list of activities logged, refer to Azure AD activity list. | [optional] 
**ProvisioningStatusInfo** | Pointer to [**AnyOfmicrosoftGraphProvisioningStatusInfo**](anyOf&lt;microsoft.graph.provisioningStatusInfo&gt;.md) | Details of provisioning status. | [optional] 
**ProvisioningSteps** | Pointer to [**[]AnyOfmicrosoftGraphProvisioningStep**](AnyOfmicrosoftGraphProvisioningStep.md) | Details of each step in provisioning. | [optional] 
**ServicePrincipal** | Pointer to [**AnyOfmicrosoftGraphProvisioningServicePrincipal**](anyOf&lt;microsoft.graph.provisioningServicePrincipal&gt;.md) | Represents the service principal used for provisioning. | [optional] 
**SourceIdentity** | Pointer to [**AnyOfmicrosoftGraphProvisionedIdentity**](anyOf&lt;microsoft.graph.provisionedIdentity&gt;.md) | Details of source object being provisioned. | [optional] 
**SourceSystem** | Pointer to [**AnyOfmicrosoftGraphProvisioningSystem**](anyOf&lt;microsoft.graph.provisioningSystem&gt;.md) | Details of source system of the object being provisioned. | [optional] 
**TargetIdentity** | Pointer to [**AnyOfmicrosoftGraphProvisionedIdentity**](anyOf&lt;microsoft.graph.provisionedIdentity&gt;.md) | Details of target object being provisioned. | [optional] 
**TargetSystem** | Pointer to [**AnyOfmicrosoftGraphProvisioningSystem**](anyOf&lt;microsoft.graph.provisioningSystem&gt;.md) | Details of target system of the object being provisioned. | [optional] 
**TenantId** | Pointer to **NullableString** | Unique Azure AD tenant ID. | [optional] 

## Methods

### NewMicrosoftGraphProvisioningObjectSummary

`func NewMicrosoftGraphProvisioningObjectSummary() *MicrosoftGraphProvisioningObjectSummary`

NewMicrosoftGraphProvisioningObjectSummary instantiates a new MicrosoftGraphProvisioningObjectSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphProvisioningObjectSummaryWithDefaults

`func NewMicrosoftGraphProvisioningObjectSummaryWithDefaults() *MicrosoftGraphProvisioningObjectSummary`

NewMicrosoftGraphProvisioningObjectSummaryWithDefaults instantiates a new MicrosoftGraphProvisioningObjectSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphProvisioningObjectSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphProvisioningObjectSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphProvisioningObjectSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphProvisioningObjectSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActivityDateTime

`func (o *MicrosoftGraphProvisioningObjectSummary) GetActivityDateTime() time.Time`

GetActivityDateTime returns the ActivityDateTime field if non-nil, zero value otherwise.

### GetActivityDateTimeOk

`func (o *MicrosoftGraphProvisioningObjectSummary) GetActivityDateTimeOk() (*time.Time, bool)`

GetActivityDateTimeOk returns a tuple with the ActivityDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityDateTime

`func (o *MicrosoftGraphProvisioningObjectSummary) SetActivityDateTime(v time.Time)`

SetActivityDateTime sets ActivityDateTime field to given value.

### HasActivityDateTime

`func (o *MicrosoftGraphProvisioningObjectSummary) HasActivityDateTime() bool`

HasActivityDateTime returns a boolean if a field has been set.

### GetChangeId

`func (o *MicrosoftGraphProvisioningObjectSummary) GetChangeId() string`

GetChangeId returns the ChangeId field if non-nil, zero value otherwise.

### GetChangeIdOk

`func (o *MicrosoftGraphProvisioningObjectSummary) GetChangeIdOk() (*string, bool)`

GetChangeIdOk returns a tuple with the ChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeId

`func (o *MicrosoftGraphProvisioningObjectSummary) SetChangeId(v string)`

SetChangeId sets ChangeId field to given value.

### HasChangeId

`func (o *MicrosoftGraphProvisioningObjectSummary) HasChangeId() bool`

HasChangeId returns a boolean if a field has been set.

### SetChangeIdNil

`func (o *MicrosoftGraphProvisioningObjectSummary) SetChangeIdNil(b bool)`

 SetChangeIdNil sets the value for ChangeId to be an explicit nil

### UnsetChangeId
`func (o *MicrosoftGraphProvisioningObjectSummary) UnsetChangeId()`

UnsetChangeId ensures that no value is present for ChangeId, not even an explicit nil
### GetCycleId

`func (o *MicrosoftGraphProvisioningObjectSummary) GetCycleId() string`

GetCycleId returns the CycleId field if non-nil, zero value otherwise.

### GetCycleIdOk

`func (o *MicrosoftGraphProvisioningObjectSummary) GetCycleIdOk() (*string, bool)`

GetCycleIdOk returns a tuple with the CycleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleId

`func (o *MicrosoftGraphProvisioningObjectSummary) SetCycleId(v string)`

SetCycleId sets CycleId field to given value.

### HasCycleId

`func (o *MicrosoftGraphProvisioningObjectSummary) HasCycleId() bool`

HasCycleId returns a boolean if a field has been set.

### SetCycleIdNil

`func (o *MicrosoftGraphProvisioningObjectSummary) SetCycleIdNil(b bool)`

 SetCycleIdNil sets the value for CycleId to be an explicit nil

### UnsetCycleId
`func (o *MicrosoftGraphProvisioningObjectSummary) UnsetCycleId()`

UnsetCycleId ensures that no value is present for CycleId, not even an explicit nil
### GetDurationInMilliseconds

`func (o *MicrosoftGraphProvisioningObjectSummary) GetDurationInMilliseconds() int32`

GetDurationInMilliseconds returns the DurationInMilliseconds field if non-nil, zero value otherwise.

### GetDurationInMillisecondsOk

`func (o *MicrosoftGraphProvisioningObjectSummary) GetDurationInMillisecondsOk() (*int32, bool)`

GetDurationInMillisecondsOk returns a tuple with the DurationInMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationInMilliseconds

`func (o *MicrosoftGraphProvisioningObjectSummary) SetDurationInMilliseconds(v int32)`

SetDurationInMilliseconds sets DurationInMilliseconds field to given value.

### HasDurationInMilliseconds

`func (o *MicrosoftGraphProvisioningObjectSummary) HasDurationInMilliseconds() bool`

HasDurationInMilliseconds returns a boolean if a field has been set.

### SetDurationInMillisecondsNil

`func (o *MicrosoftGraphProvisioningObjectSummary) SetDurationInMillisecondsNil(b bool)`

 SetDurationInMillisecondsNil sets the value for DurationInMilliseconds to be an explicit nil

### UnsetDurationInMilliseconds
`func (o *MicrosoftGraphProvisioningObjectSummary) UnsetDurationInMilliseconds()`

UnsetDurationInMilliseconds ensures that no value is present for DurationInMilliseconds, not even an explicit nil
### GetInitiatedBy

`func (o *MicrosoftGraphProvisioningObjectSummary) GetInitiatedBy() AnyOfmicrosoftGraphInitiator`

GetInitiatedBy returns the InitiatedBy field if non-nil, zero value otherwise.

### GetInitiatedByOk

`func (o *MicrosoftGraphProvisioningObjectSummary) GetInitiatedByOk() (*AnyOfmicrosoftGraphInitiator, bool)`

GetInitiatedByOk returns a tuple with the InitiatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatedBy

`func (o *MicrosoftGraphProvisioningObjectSummary) SetInitiatedBy(v AnyOfmicrosoftGraphInitiator)`

SetInitiatedBy sets InitiatedBy field to given value.

### HasInitiatedBy

`func (o *MicrosoftGraphProvisioningObjectSummary) HasInitiatedBy() bool`

HasInitiatedBy returns a boolean if a field has been set.

### SetInitiatedByNil

`func (o *MicrosoftGraphProvisioningObjectSummary) SetInitiatedByNil(b bool)`

 SetInitiatedByNil sets the value for InitiatedBy to be an explicit nil

### UnsetInitiatedBy
`func (o *MicrosoftGraphProvisioningObjectSummary) UnsetInitiatedBy()`

UnsetInitiatedBy ensures that no value is present for InitiatedBy, not even an explicit nil
### GetJobId

`func (o *MicrosoftGraphProvisioningObjectSummary) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *MicrosoftGraphProvisioningObjectSummary) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *MicrosoftGraphProvisioningObjectSummary) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *MicrosoftGraphProvisioningObjectSummary) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *MicrosoftGraphProvisioningObjectSummary) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *MicrosoftGraphProvisioningObjectSummary) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil
### GetModifiedProperties

`func (o *MicrosoftGraphProvisioningObjectSummary) GetModifiedProperties() []*AnyOfmicrosoftGraphModifiedProperty`

GetModifiedProperties returns the ModifiedProperties field if non-nil, zero value otherwise.

### GetModifiedPropertiesOk

`func (o *MicrosoftGraphProvisioningObjectSummary) GetModifiedPropertiesOk() (*[]*AnyOfmicrosoftGraphModifiedProperty, bool)`

GetModifiedPropertiesOk returns a tuple with the ModifiedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedProperties

`func (o *MicrosoftGraphProvisioningObjectSummary) SetModifiedProperties(v []*AnyOfmicrosoftGraphModifiedProperty)`

SetModifiedProperties sets ModifiedProperties field to given value.

### HasModifiedProperties

`func (o *MicrosoftGraphProvisioningObjectSummary) HasModifiedProperties() bool`

HasModifiedProperties returns a boolean if a field has been set.

### GetProvisioningAction

`func (o *MicrosoftGraphProvisioningObjectSummary) GetProvisioningAction() AnyOfmicrosoftGraphProvisioningAction`

GetProvisioningAction returns the ProvisioningAction field if non-nil, zero value otherwise.

### GetProvisioningActionOk

`func (o *MicrosoftGraphProvisioningObjectSummary) GetProvisioningActionOk() (*AnyOfmicrosoftGraphProvisioningAction, bool)`

GetProvisioningActionOk returns a tuple with the ProvisioningAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningAction

`func (o *MicrosoftGraphProvisioningObjectSummary) SetProvisioningAction(v AnyOfmicrosoftGraphProvisioningAction)`

SetProvisioningAction sets ProvisioningAction field to given value.

### HasProvisioningAction

`func (o *MicrosoftGraphProvisioningObjectSummary) HasProvisioningAction() bool`

HasProvisioningAction returns a boolean if a field has been set.

### SetProvisioningActionNil

`func (o *MicrosoftGraphProvisioningObjectSummary) SetProvisioningActionNil(b bool)`

 SetProvisioningActionNil sets the value for ProvisioningAction to be an explicit nil

### UnsetProvisioningAction
`func (o *MicrosoftGraphProvisioningObjectSummary) UnsetProvisioningAction()`

UnsetProvisioningAction ensures that no value is present for ProvisioningAction, not even an explicit nil
### GetProvisioningStatusInfo

`func (o *MicrosoftGraphProvisioningObjectSummary) GetProvisioningStatusInfo() AnyOfmicrosoftGraphProvisioningStatusInfo`

GetProvisioningStatusInfo returns the ProvisioningStatusInfo field if non-nil, zero value otherwise.

### GetProvisioningStatusInfoOk

`func (o *MicrosoftGraphProvisioningObjectSummary) GetProvisioningStatusInfoOk() (*AnyOfmicrosoftGraphProvisioningStatusInfo, bool)`

GetProvisioningStatusInfoOk returns a tuple with the ProvisioningStatusInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatusInfo

`func (o *MicrosoftGraphProvisioningObjectSummary) SetProvisioningStatusInfo(v AnyOfmicrosoftGraphProvisioningStatusInfo)`

SetProvisioningStatusInfo sets ProvisioningStatusInfo field to given value.

### HasProvisioningStatusInfo

`func (o *MicrosoftGraphProvisioningObjectSummary) HasProvisioningStatusInfo() bool`

HasProvisioningStatusInfo returns a boolean if a field has been set.

### SetProvisioningStatusInfoNil

`func (o *MicrosoftGraphProvisioningObjectSummary) SetProvisioningStatusInfoNil(b bool)`

 SetProvisioningStatusInfoNil sets the value for ProvisioningStatusInfo to be an explicit nil

### UnsetProvisioningStatusInfo
`func (o *MicrosoftGraphProvisioningObjectSummary) UnsetProvisioningStatusInfo()`

UnsetProvisioningStatusInfo ensures that no value is present for ProvisioningStatusInfo, not even an explicit nil
### GetProvisioningSteps

`func (o *MicrosoftGraphProvisioningObjectSummary) GetProvisioningSteps() []*AnyOfmicrosoftGraphProvisioningStep`

GetProvisioningSteps returns the ProvisioningSteps field if non-nil, zero value otherwise.

### GetProvisioningStepsOk

`func (o *MicrosoftGraphProvisioningObjectSummary) GetProvisioningStepsOk() (*[]*AnyOfmicrosoftGraphProvisioningStep, bool)`

GetProvisioningStepsOk returns a tuple with the ProvisioningSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSteps

`func (o *MicrosoftGraphProvisioningObjectSummary) SetProvisioningSteps(v []*AnyOfmicrosoftGraphProvisioningStep)`

SetProvisioningSteps sets ProvisioningSteps field to given value.

### HasProvisioningSteps

`func (o *MicrosoftGraphProvisioningObjectSummary) HasProvisioningSteps() bool`

HasProvisioningSteps returns a boolean if a field has been set.

### GetServicePrincipal

`func (o *MicrosoftGraphProvisioningObjectSummary) GetServicePrincipal() AnyOfmicrosoftGraphProvisioningServicePrincipal`

GetServicePrincipal returns the ServicePrincipal field if non-nil, zero value otherwise.

### GetServicePrincipalOk

`func (o *MicrosoftGraphProvisioningObjectSummary) GetServicePrincipalOk() (*AnyOfmicrosoftGraphProvisioningServicePrincipal, bool)`

GetServicePrincipalOk returns a tuple with the ServicePrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePrincipal

`func (o *MicrosoftGraphProvisioningObjectSummary) SetServicePrincipal(v AnyOfmicrosoftGraphProvisioningServicePrincipal)`

SetServicePrincipal sets ServicePrincipal field to given value.

### HasServicePrincipal

`func (o *MicrosoftGraphProvisioningObjectSummary) HasServicePrincipal() bool`

HasServicePrincipal returns a boolean if a field has been set.

### SetServicePrincipalNil

`func (o *MicrosoftGraphProvisioningObjectSummary) SetServicePrincipalNil(b bool)`

 SetServicePrincipalNil sets the value for ServicePrincipal to be an explicit nil

### UnsetServicePrincipal
`func (o *MicrosoftGraphProvisioningObjectSummary) UnsetServicePrincipal()`

UnsetServicePrincipal ensures that no value is present for ServicePrincipal, not even an explicit nil
### GetSourceIdentity

`func (o *MicrosoftGraphProvisioningObjectSummary) GetSourceIdentity() AnyOfmicrosoftGraphProvisionedIdentity`

GetSourceIdentity returns the SourceIdentity field if non-nil, zero value otherwise.

### GetSourceIdentityOk

`func (o *MicrosoftGraphProvisioningObjectSummary) GetSourceIdentityOk() (*AnyOfmicrosoftGraphProvisionedIdentity, bool)`

GetSourceIdentityOk returns a tuple with the SourceIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIdentity

`func (o *MicrosoftGraphProvisioningObjectSummary) SetSourceIdentity(v AnyOfmicrosoftGraphProvisionedIdentity)`

SetSourceIdentity sets SourceIdentity field to given value.

### HasSourceIdentity

`func (o *MicrosoftGraphProvisioningObjectSummary) HasSourceIdentity() bool`

HasSourceIdentity returns a boolean if a field has been set.

### SetSourceIdentityNil

`func (o *MicrosoftGraphProvisioningObjectSummary) SetSourceIdentityNil(b bool)`

 SetSourceIdentityNil sets the value for SourceIdentity to be an explicit nil

### UnsetSourceIdentity
`func (o *MicrosoftGraphProvisioningObjectSummary) UnsetSourceIdentity()`

UnsetSourceIdentity ensures that no value is present for SourceIdentity, not even an explicit nil
### GetSourceSystem

`func (o *MicrosoftGraphProvisioningObjectSummary) GetSourceSystem() AnyOfmicrosoftGraphProvisioningSystem`

GetSourceSystem returns the SourceSystem field if non-nil, zero value otherwise.

### GetSourceSystemOk

`func (o *MicrosoftGraphProvisioningObjectSummary) GetSourceSystemOk() (*AnyOfmicrosoftGraphProvisioningSystem, bool)`

GetSourceSystemOk returns a tuple with the SourceSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSystem

`func (o *MicrosoftGraphProvisioningObjectSummary) SetSourceSystem(v AnyOfmicrosoftGraphProvisioningSystem)`

SetSourceSystem sets SourceSystem field to given value.

### HasSourceSystem

`func (o *MicrosoftGraphProvisioningObjectSummary) HasSourceSystem() bool`

HasSourceSystem returns a boolean if a field has been set.

### SetSourceSystemNil

`func (o *MicrosoftGraphProvisioningObjectSummary) SetSourceSystemNil(b bool)`

 SetSourceSystemNil sets the value for SourceSystem to be an explicit nil

### UnsetSourceSystem
`func (o *MicrosoftGraphProvisioningObjectSummary) UnsetSourceSystem()`

UnsetSourceSystem ensures that no value is present for SourceSystem, not even an explicit nil
### GetTargetIdentity

`func (o *MicrosoftGraphProvisioningObjectSummary) GetTargetIdentity() AnyOfmicrosoftGraphProvisionedIdentity`

GetTargetIdentity returns the TargetIdentity field if non-nil, zero value otherwise.

### GetTargetIdentityOk

`func (o *MicrosoftGraphProvisioningObjectSummary) GetTargetIdentityOk() (*AnyOfmicrosoftGraphProvisionedIdentity, bool)`

GetTargetIdentityOk returns a tuple with the TargetIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetIdentity

`func (o *MicrosoftGraphProvisioningObjectSummary) SetTargetIdentity(v AnyOfmicrosoftGraphProvisionedIdentity)`

SetTargetIdentity sets TargetIdentity field to given value.

### HasTargetIdentity

`func (o *MicrosoftGraphProvisioningObjectSummary) HasTargetIdentity() bool`

HasTargetIdentity returns a boolean if a field has been set.

### SetTargetIdentityNil

`func (o *MicrosoftGraphProvisioningObjectSummary) SetTargetIdentityNil(b bool)`

 SetTargetIdentityNil sets the value for TargetIdentity to be an explicit nil

### UnsetTargetIdentity
`func (o *MicrosoftGraphProvisioningObjectSummary) UnsetTargetIdentity()`

UnsetTargetIdentity ensures that no value is present for TargetIdentity, not even an explicit nil
### GetTargetSystem

`func (o *MicrosoftGraphProvisioningObjectSummary) GetTargetSystem() AnyOfmicrosoftGraphProvisioningSystem`

GetTargetSystem returns the TargetSystem field if non-nil, zero value otherwise.

### GetTargetSystemOk

`func (o *MicrosoftGraphProvisioningObjectSummary) GetTargetSystemOk() (*AnyOfmicrosoftGraphProvisioningSystem, bool)`

GetTargetSystemOk returns a tuple with the TargetSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSystem

`func (o *MicrosoftGraphProvisioningObjectSummary) SetTargetSystem(v AnyOfmicrosoftGraphProvisioningSystem)`

SetTargetSystem sets TargetSystem field to given value.

### HasTargetSystem

`func (o *MicrosoftGraphProvisioningObjectSummary) HasTargetSystem() bool`

HasTargetSystem returns a boolean if a field has been set.

### SetTargetSystemNil

`func (o *MicrosoftGraphProvisioningObjectSummary) SetTargetSystemNil(b bool)`

 SetTargetSystemNil sets the value for TargetSystem to be an explicit nil

### UnsetTargetSystem
`func (o *MicrosoftGraphProvisioningObjectSummary) UnsetTargetSystem()`

UnsetTargetSystem ensures that no value is present for TargetSystem, not even an explicit nil
### GetTenantId

`func (o *MicrosoftGraphProvisioningObjectSummary) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *MicrosoftGraphProvisioningObjectSummary) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *MicrosoftGraphProvisioningObjectSummary) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *MicrosoftGraphProvisioningObjectSummary) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *MicrosoftGraphProvisioningObjectSummary) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *MicrosoftGraphProvisioningObjectSummary) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


