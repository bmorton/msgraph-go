# MicrosoftGraphLicenseAssignmentState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedByGroup** | Pointer to **NullableString** | The id of the group that assigns this license. If the assignment is a direct-assigned license, this field will be Null. Read-Only. | [optional] 
**DisabledPlans** | Pointer to **[]string** | The service plans that are disabled in this assignment. Read-Only. | [optional] 
**Error** | Pointer to **NullableString** | License assignment failure error. If the license is assigned successfully, this field will be Null. Read-Only. Possible values: CountViolation, MutuallyExclusiveViolation, DependencyViolation, ProhibitedInUsageLocationViolation, UniquenessViolation, and Others. For more information on how to identify and resolve license assignment errors see here. | [optional] 
**SkuId** | Pointer to **NullableString** | The unique identifier for the SKU. Read-Only. | [optional] 
**State** | Pointer to **NullableString** | Indicate the current state of this assignment. Read-Only. Possible values: Active, ActiveWithError, Disabled and Error. | [optional] 

## Methods

### NewMicrosoftGraphLicenseAssignmentState

`func NewMicrosoftGraphLicenseAssignmentState() *MicrosoftGraphLicenseAssignmentState`

NewMicrosoftGraphLicenseAssignmentState instantiates a new MicrosoftGraphLicenseAssignmentState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphLicenseAssignmentStateWithDefaults

`func NewMicrosoftGraphLicenseAssignmentStateWithDefaults() *MicrosoftGraphLicenseAssignmentState`

NewMicrosoftGraphLicenseAssignmentStateWithDefaults instantiates a new MicrosoftGraphLicenseAssignmentState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedByGroup

`func (o *MicrosoftGraphLicenseAssignmentState) GetAssignedByGroup() string`

GetAssignedByGroup returns the AssignedByGroup field if non-nil, zero value otherwise.

### GetAssignedByGroupOk

`func (o *MicrosoftGraphLicenseAssignmentState) GetAssignedByGroupOk() (*string, bool)`

GetAssignedByGroupOk returns a tuple with the AssignedByGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedByGroup

`func (o *MicrosoftGraphLicenseAssignmentState) SetAssignedByGroup(v string)`

SetAssignedByGroup sets AssignedByGroup field to given value.

### HasAssignedByGroup

`func (o *MicrosoftGraphLicenseAssignmentState) HasAssignedByGroup() bool`

HasAssignedByGroup returns a boolean if a field has been set.

### SetAssignedByGroupNil

`func (o *MicrosoftGraphLicenseAssignmentState) SetAssignedByGroupNil(b bool)`

 SetAssignedByGroupNil sets the value for AssignedByGroup to be an explicit nil

### UnsetAssignedByGroup
`func (o *MicrosoftGraphLicenseAssignmentState) UnsetAssignedByGroup()`

UnsetAssignedByGroup ensures that no value is present for AssignedByGroup, not even an explicit nil
### GetDisabledPlans

`func (o *MicrosoftGraphLicenseAssignmentState) GetDisabledPlans() []*string`

GetDisabledPlans returns the DisabledPlans field if non-nil, zero value otherwise.

### GetDisabledPlansOk

`func (o *MicrosoftGraphLicenseAssignmentState) GetDisabledPlansOk() (*[]*string, bool)`

GetDisabledPlansOk returns a tuple with the DisabledPlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledPlans

`func (o *MicrosoftGraphLicenseAssignmentState) SetDisabledPlans(v []*string)`

SetDisabledPlans sets DisabledPlans field to given value.

### HasDisabledPlans

`func (o *MicrosoftGraphLicenseAssignmentState) HasDisabledPlans() bool`

HasDisabledPlans returns a boolean if a field has been set.

### GetError

`func (o *MicrosoftGraphLicenseAssignmentState) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MicrosoftGraphLicenseAssignmentState) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MicrosoftGraphLicenseAssignmentState) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *MicrosoftGraphLicenseAssignmentState) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *MicrosoftGraphLicenseAssignmentState) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *MicrosoftGraphLicenseAssignmentState) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetSkuId

`func (o *MicrosoftGraphLicenseAssignmentState) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *MicrosoftGraphLicenseAssignmentState) GetSkuIdOk() (*string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *MicrosoftGraphLicenseAssignmentState) SetSkuId(v string)`

SetSkuId sets SkuId field to given value.

### HasSkuId

`func (o *MicrosoftGraphLicenseAssignmentState) HasSkuId() bool`

HasSkuId returns a boolean if a field has been set.

### SetSkuIdNil

`func (o *MicrosoftGraphLicenseAssignmentState) SetSkuIdNil(b bool)`

 SetSkuIdNil sets the value for SkuId to be an explicit nil

### UnsetSkuId
`func (o *MicrosoftGraphLicenseAssignmentState) UnsetSkuId()`

UnsetSkuId ensures that no value is present for SkuId, not even an explicit nil
### GetState

`func (o *MicrosoftGraphLicenseAssignmentState) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MicrosoftGraphLicenseAssignmentState) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MicrosoftGraphLicenseAssignmentState) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MicrosoftGraphLicenseAssignmentState) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *MicrosoftGraphLicenseAssignmentState) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *MicrosoftGraphLicenseAssignmentState) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


