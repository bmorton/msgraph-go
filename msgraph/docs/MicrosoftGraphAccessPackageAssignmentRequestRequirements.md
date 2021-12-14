# MicrosoftGraphAccessPackageAssignmentRequestRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowCustomAssignmentSchedule** | Pointer to **NullableBool** | Indicates whether the requestor is allowed to set a custom schedule. | [optional] 
**IsApprovalRequiredForAdd** | Pointer to **NullableBool** | Indicates whether a request to add must be approved by an approver. | [optional] 
**IsApprovalRequiredForUpdate** | Pointer to **NullableBool** | Indicates whether a request to update must be approved by an approver. | [optional] 
**PolicyDescription** | Pointer to **NullableString** | The description of the policy that the user is trying to request access using. | [optional] 
**PolicyDisplayName** | Pointer to **NullableString** | The display name of the policy that the user is trying to request access using. | [optional] 
**PolicyId** | Pointer to **NullableString** | The identifier of the policy that these requirements are associated with. This identifier can be used when creating a new assignment request. | [optional] 
**Schedule** | Pointer to [**AnyOfmicrosoftGraphEntitlementManagementSchedule**](anyOf&lt;microsoft.graph.entitlementManagementSchedule&gt;.md) | Schedule restrictions enforced, if any. | [optional] 

## Methods

### NewMicrosoftGraphAccessPackageAssignmentRequestRequirements

`func NewMicrosoftGraphAccessPackageAssignmentRequestRequirements() *MicrosoftGraphAccessPackageAssignmentRequestRequirements`

NewMicrosoftGraphAccessPackageAssignmentRequestRequirements instantiates a new MicrosoftGraphAccessPackageAssignmentRequestRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAccessPackageAssignmentRequestRequirementsWithDefaults

`func NewMicrosoftGraphAccessPackageAssignmentRequestRequirementsWithDefaults() *MicrosoftGraphAccessPackageAssignmentRequestRequirements`

NewMicrosoftGraphAccessPackageAssignmentRequestRequirementsWithDefaults instantiates a new MicrosoftGraphAccessPackageAssignmentRequestRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowCustomAssignmentSchedule

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetAllowCustomAssignmentSchedule() bool`

GetAllowCustomAssignmentSchedule returns the AllowCustomAssignmentSchedule field if non-nil, zero value otherwise.

### GetAllowCustomAssignmentScheduleOk

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetAllowCustomAssignmentScheduleOk() (*bool, bool)`

GetAllowCustomAssignmentScheduleOk returns a tuple with the AllowCustomAssignmentSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomAssignmentSchedule

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetAllowCustomAssignmentSchedule(v bool)`

SetAllowCustomAssignmentSchedule sets AllowCustomAssignmentSchedule field to given value.

### HasAllowCustomAssignmentSchedule

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) HasAllowCustomAssignmentSchedule() bool`

HasAllowCustomAssignmentSchedule returns a boolean if a field has been set.

### SetAllowCustomAssignmentScheduleNil

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetAllowCustomAssignmentScheduleNil(b bool)`

 SetAllowCustomAssignmentScheduleNil sets the value for AllowCustomAssignmentSchedule to be an explicit nil

### UnsetAllowCustomAssignmentSchedule
`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) UnsetAllowCustomAssignmentSchedule()`

UnsetAllowCustomAssignmentSchedule ensures that no value is present for AllowCustomAssignmentSchedule, not even an explicit nil
### GetIsApprovalRequiredForAdd

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetIsApprovalRequiredForAdd() bool`

GetIsApprovalRequiredForAdd returns the IsApprovalRequiredForAdd field if non-nil, zero value otherwise.

### GetIsApprovalRequiredForAddOk

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetIsApprovalRequiredForAddOk() (*bool, bool)`

GetIsApprovalRequiredForAddOk returns a tuple with the IsApprovalRequiredForAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApprovalRequiredForAdd

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetIsApprovalRequiredForAdd(v bool)`

SetIsApprovalRequiredForAdd sets IsApprovalRequiredForAdd field to given value.

### HasIsApprovalRequiredForAdd

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) HasIsApprovalRequiredForAdd() bool`

HasIsApprovalRequiredForAdd returns a boolean if a field has been set.

### SetIsApprovalRequiredForAddNil

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetIsApprovalRequiredForAddNil(b bool)`

 SetIsApprovalRequiredForAddNil sets the value for IsApprovalRequiredForAdd to be an explicit nil

### UnsetIsApprovalRequiredForAdd
`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) UnsetIsApprovalRequiredForAdd()`

UnsetIsApprovalRequiredForAdd ensures that no value is present for IsApprovalRequiredForAdd, not even an explicit nil
### GetIsApprovalRequiredForUpdate

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetIsApprovalRequiredForUpdate() bool`

GetIsApprovalRequiredForUpdate returns the IsApprovalRequiredForUpdate field if non-nil, zero value otherwise.

### GetIsApprovalRequiredForUpdateOk

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetIsApprovalRequiredForUpdateOk() (*bool, bool)`

GetIsApprovalRequiredForUpdateOk returns a tuple with the IsApprovalRequiredForUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsApprovalRequiredForUpdate

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetIsApprovalRequiredForUpdate(v bool)`

SetIsApprovalRequiredForUpdate sets IsApprovalRequiredForUpdate field to given value.

### HasIsApprovalRequiredForUpdate

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) HasIsApprovalRequiredForUpdate() bool`

HasIsApprovalRequiredForUpdate returns a boolean if a field has been set.

### SetIsApprovalRequiredForUpdateNil

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetIsApprovalRequiredForUpdateNil(b bool)`

 SetIsApprovalRequiredForUpdateNil sets the value for IsApprovalRequiredForUpdate to be an explicit nil

### UnsetIsApprovalRequiredForUpdate
`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) UnsetIsApprovalRequiredForUpdate()`

UnsetIsApprovalRequiredForUpdate ensures that no value is present for IsApprovalRequiredForUpdate, not even an explicit nil
### GetPolicyDescription

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetPolicyDescription() string`

GetPolicyDescription returns the PolicyDescription field if non-nil, zero value otherwise.

### GetPolicyDescriptionOk

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetPolicyDescriptionOk() (*string, bool)`

GetPolicyDescriptionOk returns a tuple with the PolicyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDescription

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetPolicyDescription(v string)`

SetPolicyDescription sets PolicyDescription field to given value.

### HasPolicyDescription

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) HasPolicyDescription() bool`

HasPolicyDescription returns a boolean if a field has been set.

### SetPolicyDescriptionNil

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetPolicyDescriptionNil(b bool)`

 SetPolicyDescriptionNil sets the value for PolicyDescription to be an explicit nil

### UnsetPolicyDescription
`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) UnsetPolicyDescription()`

UnsetPolicyDescription ensures that no value is present for PolicyDescription, not even an explicit nil
### GetPolicyDisplayName

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetPolicyDisplayName() string`

GetPolicyDisplayName returns the PolicyDisplayName field if non-nil, zero value otherwise.

### GetPolicyDisplayNameOk

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetPolicyDisplayNameOk() (*string, bool)`

GetPolicyDisplayNameOk returns a tuple with the PolicyDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDisplayName

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetPolicyDisplayName(v string)`

SetPolicyDisplayName sets PolicyDisplayName field to given value.

### HasPolicyDisplayName

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) HasPolicyDisplayName() bool`

HasPolicyDisplayName returns a boolean if a field has been set.

### SetPolicyDisplayNameNil

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetPolicyDisplayNameNil(b bool)`

 SetPolicyDisplayNameNil sets the value for PolicyDisplayName to be an explicit nil

### UnsetPolicyDisplayName
`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) UnsetPolicyDisplayName()`

UnsetPolicyDisplayName ensures that no value is present for PolicyDisplayName, not even an explicit nil
### GetPolicyId

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetSchedule

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetSchedule() AnyOfmicrosoftGraphEntitlementManagementSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) GetScheduleOk() (*AnyOfmicrosoftGraphEntitlementManagementSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetSchedule(v AnyOfmicrosoftGraphEntitlementManagementSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *MicrosoftGraphAccessPackageAssignmentRequestRequirements) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


