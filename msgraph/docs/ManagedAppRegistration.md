# ManagedAppRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppIdentifier** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | The app package Identifier | [optional] 
**ApplicationVersion** | Pointer to **NullableString** | App version | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | Date and time of creation | [optional] 
**DeviceName** | Pointer to **NullableString** | Host device name | [optional] 
**DeviceTag** | Pointer to **NullableString** | App management SDK generated tag, which helps relate apps hosted on the same device. Not guaranteed to relate apps in all conditions. | [optional] 
**DeviceType** | Pointer to **NullableString** | Host device type | [optional] 
**FlaggedReasons** | Pointer to [**[]AnyOfmicrosoftGraphManagedAppFlaggedReason**](AnyOfmicrosoftGraphManagedAppFlaggedReason.md) | Zero or more reasons an app registration is flagged. E.g. app running on rooted device | [optional] 
**LastSyncDateTime** | Pointer to **time.Time** | Date and time of last the app synced with management service. | [optional] 
**ManagementSdkVersion** | Pointer to **NullableString** | App management SDK version | [optional] 
**PlatformVersion** | Pointer to **NullableString** | Operating System version | [optional] 
**UserId** | Pointer to **NullableString** | The user Id to who this app registration belongs. | [optional] 
**Version** | Pointer to **NullableString** | Version of the entity. | [optional] 
**AppliedPolicies** | Pointer to [**[]MicrosoftGraphManagedAppPolicy**](MicrosoftGraphManagedAppPolicy.md) | Zero or more policys already applied on the registered app when it last synchronized with managment service. | [optional] 
**IntendedPolicies** | Pointer to [**[]MicrosoftGraphManagedAppPolicy**](MicrosoftGraphManagedAppPolicy.md) | Zero or more policies admin intended for the app as of now. | [optional] 
**Operations** | Pointer to [**[]MicrosoftGraphManagedAppOperation**](MicrosoftGraphManagedAppOperation.md) | Zero or more long running operations triggered on the app registration. | [optional] 

## Methods

### NewManagedAppRegistration

`func NewManagedAppRegistration() *ManagedAppRegistration`

NewManagedAppRegistration instantiates a new ManagedAppRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedAppRegistrationWithDefaults

`func NewManagedAppRegistrationWithDefaults() *ManagedAppRegistration`

NewManagedAppRegistrationWithDefaults instantiates a new ManagedAppRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppIdentifier

`func (o *ManagedAppRegistration) GetAppIdentifier() AnyOfobject`

GetAppIdentifier returns the AppIdentifier field if non-nil, zero value otherwise.

### GetAppIdentifierOk

`func (o *ManagedAppRegistration) GetAppIdentifierOk() (*AnyOfobject, bool)`

GetAppIdentifierOk returns a tuple with the AppIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIdentifier

`func (o *ManagedAppRegistration) SetAppIdentifier(v AnyOfobject)`

SetAppIdentifier sets AppIdentifier field to given value.

### HasAppIdentifier

`func (o *ManagedAppRegistration) HasAppIdentifier() bool`

HasAppIdentifier returns a boolean if a field has been set.

### SetAppIdentifierNil

`func (o *ManagedAppRegistration) SetAppIdentifierNil(b bool)`

 SetAppIdentifierNil sets the value for AppIdentifier to be an explicit nil

### UnsetAppIdentifier
`func (o *ManagedAppRegistration) UnsetAppIdentifier()`

UnsetAppIdentifier ensures that no value is present for AppIdentifier, not even an explicit nil
### GetApplicationVersion

`func (o *ManagedAppRegistration) GetApplicationVersion() string`

GetApplicationVersion returns the ApplicationVersion field if non-nil, zero value otherwise.

### GetApplicationVersionOk

`func (o *ManagedAppRegistration) GetApplicationVersionOk() (*string, bool)`

GetApplicationVersionOk returns a tuple with the ApplicationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationVersion

`func (o *ManagedAppRegistration) SetApplicationVersion(v string)`

SetApplicationVersion sets ApplicationVersion field to given value.

### HasApplicationVersion

`func (o *ManagedAppRegistration) HasApplicationVersion() bool`

HasApplicationVersion returns a boolean if a field has been set.

### SetApplicationVersionNil

`func (o *ManagedAppRegistration) SetApplicationVersionNil(b bool)`

 SetApplicationVersionNil sets the value for ApplicationVersion to be an explicit nil

### UnsetApplicationVersion
`func (o *ManagedAppRegistration) UnsetApplicationVersion()`

UnsetApplicationVersion ensures that no value is present for ApplicationVersion, not even an explicit nil
### GetCreatedDateTime

`func (o *ManagedAppRegistration) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *ManagedAppRegistration) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *ManagedAppRegistration) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *ManagedAppRegistration) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetDeviceName

`func (o *ManagedAppRegistration) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *ManagedAppRegistration) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *ManagedAppRegistration) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *ManagedAppRegistration) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### SetDeviceNameNil

`func (o *ManagedAppRegistration) SetDeviceNameNil(b bool)`

 SetDeviceNameNil sets the value for DeviceName to be an explicit nil

### UnsetDeviceName
`func (o *ManagedAppRegistration) UnsetDeviceName()`

UnsetDeviceName ensures that no value is present for DeviceName, not even an explicit nil
### GetDeviceTag

`func (o *ManagedAppRegistration) GetDeviceTag() string`

GetDeviceTag returns the DeviceTag field if non-nil, zero value otherwise.

### GetDeviceTagOk

`func (o *ManagedAppRegistration) GetDeviceTagOk() (*string, bool)`

GetDeviceTagOk returns a tuple with the DeviceTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTag

`func (o *ManagedAppRegistration) SetDeviceTag(v string)`

SetDeviceTag sets DeviceTag field to given value.

### HasDeviceTag

`func (o *ManagedAppRegistration) HasDeviceTag() bool`

HasDeviceTag returns a boolean if a field has been set.

### SetDeviceTagNil

`func (o *ManagedAppRegistration) SetDeviceTagNil(b bool)`

 SetDeviceTagNil sets the value for DeviceTag to be an explicit nil

### UnsetDeviceTag
`func (o *ManagedAppRegistration) UnsetDeviceTag()`

UnsetDeviceTag ensures that no value is present for DeviceTag, not even an explicit nil
### GetDeviceType

`func (o *ManagedAppRegistration) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ManagedAppRegistration) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ManagedAppRegistration) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *ManagedAppRegistration) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *ManagedAppRegistration) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *ManagedAppRegistration) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetFlaggedReasons

`func (o *ManagedAppRegistration) GetFlaggedReasons() []AnyOfmicrosoftGraphManagedAppFlaggedReason`

GetFlaggedReasons returns the FlaggedReasons field if non-nil, zero value otherwise.

### GetFlaggedReasonsOk

`func (o *ManagedAppRegistration) GetFlaggedReasonsOk() (*[]AnyOfmicrosoftGraphManagedAppFlaggedReason, bool)`

GetFlaggedReasonsOk returns a tuple with the FlaggedReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlaggedReasons

`func (o *ManagedAppRegistration) SetFlaggedReasons(v []AnyOfmicrosoftGraphManagedAppFlaggedReason)`

SetFlaggedReasons sets FlaggedReasons field to given value.

### HasFlaggedReasons

`func (o *ManagedAppRegistration) HasFlaggedReasons() bool`

HasFlaggedReasons returns a boolean if a field has been set.

### GetLastSyncDateTime

`func (o *ManagedAppRegistration) GetLastSyncDateTime() time.Time`

GetLastSyncDateTime returns the LastSyncDateTime field if non-nil, zero value otherwise.

### GetLastSyncDateTimeOk

`func (o *ManagedAppRegistration) GetLastSyncDateTimeOk() (*time.Time, bool)`

GetLastSyncDateTimeOk returns a tuple with the LastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDateTime

`func (o *ManagedAppRegistration) SetLastSyncDateTime(v time.Time)`

SetLastSyncDateTime sets LastSyncDateTime field to given value.

### HasLastSyncDateTime

`func (o *ManagedAppRegistration) HasLastSyncDateTime() bool`

HasLastSyncDateTime returns a boolean if a field has been set.

### GetManagementSdkVersion

`func (o *ManagedAppRegistration) GetManagementSdkVersion() string`

GetManagementSdkVersion returns the ManagementSdkVersion field if non-nil, zero value otherwise.

### GetManagementSdkVersionOk

`func (o *ManagedAppRegistration) GetManagementSdkVersionOk() (*string, bool)`

GetManagementSdkVersionOk returns a tuple with the ManagementSdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementSdkVersion

`func (o *ManagedAppRegistration) SetManagementSdkVersion(v string)`

SetManagementSdkVersion sets ManagementSdkVersion field to given value.

### HasManagementSdkVersion

`func (o *ManagedAppRegistration) HasManagementSdkVersion() bool`

HasManagementSdkVersion returns a boolean if a field has been set.

### SetManagementSdkVersionNil

`func (o *ManagedAppRegistration) SetManagementSdkVersionNil(b bool)`

 SetManagementSdkVersionNil sets the value for ManagementSdkVersion to be an explicit nil

### UnsetManagementSdkVersion
`func (o *ManagedAppRegistration) UnsetManagementSdkVersion()`

UnsetManagementSdkVersion ensures that no value is present for ManagementSdkVersion, not even an explicit nil
### GetPlatformVersion

`func (o *ManagedAppRegistration) GetPlatformVersion() string`

GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.

### GetPlatformVersionOk

`func (o *ManagedAppRegistration) GetPlatformVersionOk() (*string, bool)`

GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformVersion

`func (o *ManagedAppRegistration) SetPlatformVersion(v string)`

SetPlatformVersion sets PlatformVersion field to given value.

### HasPlatformVersion

`func (o *ManagedAppRegistration) HasPlatformVersion() bool`

HasPlatformVersion returns a boolean if a field has been set.

### SetPlatformVersionNil

`func (o *ManagedAppRegistration) SetPlatformVersionNil(b bool)`

 SetPlatformVersionNil sets the value for PlatformVersion to be an explicit nil

### UnsetPlatformVersion
`func (o *ManagedAppRegistration) UnsetPlatformVersion()`

UnsetPlatformVersion ensures that no value is present for PlatformVersion, not even an explicit nil
### GetUserId

`func (o *ManagedAppRegistration) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ManagedAppRegistration) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ManagedAppRegistration) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ManagedAppRegistration) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ManagedAppRegistration) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ManagedAppRegistration) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetVersion

`func (o *ManagedAppRegistration) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ManagedAppRegistration) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ManagedAppRegistration) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ManagedAppRegistration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *ManagedAppRegistration) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *ManagedAppRegistration) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetAppliedPolicies

`func (o *ManagedAppRegistration) GetAppliedPolicies() []MicrosoftGraphManagedAppPolicy`

GetAppliedPolicies returns the AppliedPolicies field if non-nil, zero value otherwise.

### GetAppliedPoliciesOk

`func (o *ManagedAppRegistration) GetAppliedPoliciesOk() (*[]MicrosoftGraphManagedAppPolicy, bool)`

GetAppliedPoliciesOk returns a tuple with the AppliedPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedPolicies

`func (o *ManagedAppRegistration) SetAppliedPolicies(v []MicrosoftGraphManagedAppPolicy)`

SetAppliedPolicies sets AppliedPolicies field to given value.

### HasAppliedPolicies

`func (o *ManagedAppRegistration) HasAppliedPolicies() bool`

HasAppliedPolicies returns a boolean if a field has been set.

### GetIntendedPolicies

`func (o *ManagedAppRegistration) GetIntendedPolicies() []MicrosoftGraphManagedAppPolicy`

GetIntendedPolicies returns the IntendedPolicies field if non-nil, zero value otherwise.

### GetIntendedPoliciesOk

`func (o *ManagedAppRegistration) GetIntendedPoliciesOk() (*[]MicrosoftGraphManagedAppPolicy, bool)`

GetIntendedPoliciesOk returns a tuple with the IntendedPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntendedPolicies

`func (o *ManagedAppRegistration) SetIntendedPolicies(v []MicrosoftGraphManagedAppPolicy)`

SetIntendedPolicies sets IntendedPolicies field to given value.

### HasIntendedPolicies

`func (o *ManagedAppRegistration) HasIntendedPolicies() bool`

HasIntendedPolicies returns a boolean if a field has been set.

### GetOperations

`func (o *ManagedAppRegistration) GetOperations() []MicrosoftGraphManagedAppOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *ManagedAppRegistration) GetOperationsOk() (*[]MicrosoftGraphManagedAppOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *ManagedAppRegistration) SetOperations(v []MicrosoftGraphManagedAppOperation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *ManagedAppRegistration) HasOperations() bool`

HasOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


