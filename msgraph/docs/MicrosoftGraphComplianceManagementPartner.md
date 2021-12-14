# MicrosoftGraphComplianceManagementPartner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**AndroidEnrollmentAssignments** | Pointer to [**[]AnyOfmicrosoftGraphComplianceManagementPartnerAssignment**](AnyOfmicrosoftGraphComplianceManagementPartnerAssignment.md) | User groups which enroll Android devices through partner. | [optional] 
**AndroidOnboarded** | Pointer to **bool** | Partner onboarded for Android devices. | [optional] 
**DisplayName** | Pointer to **NullableString** | Partner display name | [optional] 
**IosEnrollmentAssignments** | Pointer to [**[]AnyOfmicrosoftGraphComplianceManagementPartnerAssignment**](AnyOfmicrosoftGraphComplianceManagementPartnerAssignment.md) | User groups which enroll ios devices through partner. | [optional] 
**IosOnboarded** | Pointer to **bool** | Partner onboarded for ios devices. | [optional] 
**LastHeartbeatDateTime** | Pointer to **time.Time** | Timestamp of last heartbeat after admin onboarded to the compliance management partner | [optional] 
**MacOsEnrollmentAssignments** | Pointer to [**[]AnyOfmicrosoftGraphComplianceManagementPartnerAssignment**](AnyOfmicrosoftGraphComplianceManagementPartnerAssignment.md) | User groups which enroll Mac devices through partner. | [optional] 
**MacOsOnboarded** | Pointer to **bool** | Partner onboarded for Mac devices. | [optional] 
**PartnerState** | Pointer to [**AnyOfmicrosoftGraphDeviceManagementPartnerTenantState**](anyOf&lt;microsoft.graph.deviceManagementPartnerTenantState&gt;.md) | Partner state of this tenant. Possible values are: unknown, unavailable, enabled, terminated, rejected, unresponsive. | [optional] 

## Methods

### NewMicrosoftGraphComplianceManagementPartner

`func NewMicrosoftGraphComplianceManagementPartner() *MicrosoftGraphComplianceManagementPartner`

NewMicrosoftGraphComplianceManagementPartner instantiates a new MicrosoftGraphComplianceManagementPartner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphComplianceManagementPartnerWithDefaults

`func NewMicrosoftGraphComplianceManagementPartnerWithDefaults() *MicrosoftGraphComplianceManagementPartner`

NewMicrosoftGraphComplianceManagementPartnerWithDefaults instantiates a new MicrosoftGraphComplianceManagementPartner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphComplianceManagementPartner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphComplianceManagementPartner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphComplianceManagementPartner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphComplianceManagementPartner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAndroidEnrollmentAssignments

`func (o *MicrosoftGraphComplianceManagementPartner) GetAndroidEnrollmentAssignments() []*AnyOfmicrosoftGraphComplianceManagementPartnerAssignment`

GetAndroidEnrollmentAssignments returns the AndroidEnrollmentAssignments field if non-nil, zero value otherwise.

### GetAndroidEnrollmentAssignmentsOk

`func (o *MicrosoftGraphComplianceManagementPartner) GetAndroidEnrollmentAssignmentsOk() (*[]*AnyOfmicrosoftGraphComplianceManagementPartnerAssignment, bool)`

GetAndroidEnrollmentAssignmentsOk returns a tuple with the AndroidEnrollmentAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidEnrollmentAssignments

`func (o *MicrosoftGraphComplianceManagementPartner) SetAndroidEnrollmentAssignments(v []*AnyOfmicrosoftGraphComplianceManagementPartnerAssignment)`

SetAndroidEnrollmentAssignments sets AndroidEnrollmentAssignments field to given value.

### HasAndroidEnrollmentAssignments

`func (o *MicrosoftGraphComplianceManagementPartner) HasAndroidEnrollmentAssignments() bool`

HasAndroidEnrollmentAssignments returns a boolean if a field has been set.

### GetAndroidOnboarded

`func (o *MicrosoftGraphComplianceManagementPartner) GetAndroidOnboarded() bool`

GetAndroidOnboarded returns the AndroidOnboarded field if non-nil, zero value otherwise.

### GetAndroidOnboardedOk

`func (o *MicrosoftGraphComplianceManagementPartner) GetAndroidOnboardedOk() (*bool, bool)`

GetAndroidOnboardedOk returns a tuple with the AndroidOnboarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidOnboarded

`func (o *MicrosoftGraphComplianceManagementPartner) SetAndroidOnboarded(v bool)`

SetAndroidOnboarded sets AndroidOnboarded field to given value.

### HasAndroidOnboarded

`func (o *MicrosoftGraphComplianceManagementPartner) HasAndroidOnboarded() bool`

HasAndroidOnboarded returns a boolean if a field has been set.

### GetDisplayName

`func (o *MicrosoftGraphComplianceManagementPartner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphComplianceManagementPartner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphComplianceManagementPartner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphComplianceManagementPartner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphComplianceManagementPartner) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphComplianceManagementPartner) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetIosEnrollmentAssignments

`func (o *MicrosoftGraphComplianceManagementPartner) GetIosEnrollmentAssignments() []*AnyOfmicrosoftGraphComplianceManagementPartnerAssignment`

GetIosEnrollmentAssignments returns the IosEnrollmentAssignments field if non-nil, zero value otherwise.

### GetIosEnrollmentAssignmentsOk

`func (o *MicrosoftGraphComplianceManagementPartner) GetIosEnrollmentAssignmentsOk() (*[]*AnyOfmicrosoftGraphComplianceManagementPartnerAssignment, bool)`

GetIosEnrollmentAssignmentsOk returns a tuple with the IosEnrollmentAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosEnrollmentAssignments

`func (o *MicrosoftGraphComplianceManagementPartner) SetIosEnrollmentAssignments(v []*AnyOfmicrosoftGraphComplianceManagementPartnerAssignment)`

SetIosEnrollmentAssignments sets IosEnrollmentAssignments field to given value.

### HasIosEnrollmentAssignments

`func (o *MicrosoftGraphComplianceManagementPartner) HasIosEnrollmentAssignments() bool`

HasIosEnrollmentAssignments returns a boolean if a field has been set.

### GetIosOnboarded

`func (o *MicrosoftGraphComplianceManagementPartner) GetIosOnboarded() bool`

GetIosOnboarded returns the IosOnboarded field if non-nil, zero value otherwise.

### GetIosOnboardedOk

`func (o *MicrosoftGraphComplianceManagementPartner) GetIosOnboardedOk() (*bool, bool)`

GetIosOnboardedOk returns a tuple with the IosOnboarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIosOnboarded

`func (o *MicrosoftGraphComplianceManagementPartner) SetIosOnboarded(v bool)`

SetIosOnboarded sets IosOnboarded field to given value.

### HasIosOnboarded

`func (o *MicrosoftGraphComplianceManagementPartner) HasIosOnboarded() bool`

HasIosOnboarded returns a boolean if a field has been set.

### GetLastHeartbeatDateTime

`func (o *MicrosoftGraphComplianceManagementPartner) GetLastHeartbeatDateTime() time.Time`

GetLastHeartbeatDateTime returns the LastHeartbeatDateTime field if non-nil, zero value otherwise.

### GetLastHeartbeatDateTimeOk

`func (o *MicrosoftGraphComplianceManagementPartner) GetLastHeartbeatDateTimeOk() (*time.Time, bool)`

GetLastHeartbeatDateTimeOk returns a tuple with the LastHeartbeatDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeatDateTime

`func (o *MicrosoftGraphComplianceManagementPartner) SetLastHeartbeatDateTime(v time.Time)`

SetLastHeartbeatDateTime sets LastHeartbeatDateTime field to given value.

### HasLastHeartbeatDateTime

`func (o *MicrosoftGraphComplianceManagementPartner) HasLastHeartbeatDateTime() bool`

HasLastHeartbeatDateTime returns a boolean if a field has been set.

### GetMacOsEnrollmentAssignments

`func (o *MicrosoftGraphComplianceManagementPartner) GetMacOsEnrollmentAssignments() []*AnyOfmicrosoftGraphComplianceManagementPartnerAssignment`

GetMacOsEnrollmentAssignments returns the MacOsEnrollmentAssignments field if non-nil, zero value otherwise.

### GetMacOsEnrollmentAssignmentsOk

`func (o *MicrosoftGraphComplianceManagementPartner) GetMacOsEnrollmentAssignmentsOk() (*[]*AnyOfmicrosoftGraphComplianceManagementPartnerAssignment, bool)`

GetMacOsEnrollmentAssignmentsOk returns a tuple with the MacOsEnrollmentAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacOsEnrollmentAssignments

`func (o *MicrosoftGraphComplianceManagementPartner) SetMacOsEnrollmentAssignments(v []*AnyOfmicrosoftGraphComplianceManagementPartnerAssignment)`

SetMacOsEnrollmentAssignments sets MacOsEnrollmentAssignments field to given value.

### HasMacOsEnrollmentAssignments

`func (o *MicrosoftGraphComplianceManagementPartner) HasMacOsEnrollmentAssignments() bool`

HasMacOsEnrollmentAssignments returns a boolean if a field has been set.

### GetMacOsOnboarded

`func (o *MicrosoftGraphComplianceManagementPartner) GetMacOsOnboarded() bool`

GetMacOsOnboarded returns the MacOsOnboarded field if non-nil, zero value otherwise.

### GetMacOsOnboardedOk

`func (o *MicrosoftGraphComplianceManagementPartner) GetMacOsOnboardedOk() (*bool, bool)`

GetMacOsOnboardedOk returns a tuple with the MacOsOnboarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacOsOnboarded

`func (o *MicrosoftGraphComplianceManagementPartner) SetMacOsOnboarded(v bool)`

SetMacOsOnboarded sets MacOsOnboarded field to given value.

### HasMacOsOnboarded

`func (o *MicrosoftGraphComplianceManagementPartner) HasMacOsOnboarded() bool`

HasMacOsOnboarded returns a boolean if a field has been set.

### GetPartnerState

`func (o *MicrosoftGraphComplianceManagementPartner) GetPartnerState() AnyOfmicrosoftGraphDeviceManagementPartnerTenantState`

GetPartnerState returns the PartnerState field if non-nil, zero value otherwise.

### GetPartnerStateOk

`func (o *MicrosoftGraphComplianceManagementPartner) GetPartnerStateOk() (*AnyOfmicrosoftGraphDeviceManagementPartnerTenantState, bool)`

GetPartnerStateOk returns a tuple with the PartnerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerState

`func (o *MicrosoftGraphComplianceManagementPartner) SetPartnerState(v AnyOfmicrosoftGraphDeviceManagementPartnerTenantState)`

SetPartnerState sets PartnerState field to given value.

### HasPartnerState

`func (o *MicrosoftGraphComplianceManagementPartner) HasPartnerState() bool`

HasPartnerState returns a boolean if a field has been set.

### SetPartnerStateNil

`func (o *MicrosoftGraphComplianceManagementPartner) SetPartnerStateNil(b bool)`

 SetPartnerStateNil sets the value for PartnerState to be an explicit nil

### UnsetPartnerState
`func (o *MicrosoftGraphComplianceManagementPartner) UnsetPartnerState()`

UnsetPartnerState ensures that no value is present for PartnerState, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


