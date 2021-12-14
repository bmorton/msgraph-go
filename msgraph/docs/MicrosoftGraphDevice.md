# MicrosoftGraphDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DeletedDateTime** | Pointer to **NullableTime** |  | [optional] 
**AccountEnabled** | Pointer to **NullableBool** | true if the account is enabled; otherwise, false. Required. Default is true.  Supports $filter (eq, ne, not, in). Only callers in Global Administrator and Cloud Device Administrator roles can set this property. | [optional] 
**AlternativeSecurityIds** | Pointer to [**[]MicrosoftGraphAlternativeSecurityId**](MicrosoftGraphAlternativeSecurityId.md) | For internal use only. Not nullable. Supports $filter (eq, not, ge, le). | [optional] 
**ApproximateLastSignInDateTime** | Pointer to **NullableTime** | The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. Supports $filter (eq, ne, not, ge, le, and eq on null values) and $orderBy. | [optional] 
**ComplianceExpirationDateTime** | Pointer to **NullableTime** | The timestamp when the device is no longer deemed compliant. The timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. | [optional] 
**DeviceId** | Pointer to **NullableString** | Unique identifier set by Azure Device Registration Service at the time of registration. Supports $filter (eq, ne, not, startsWith). | [optional] 
**DeviceMetadata** | Pointer to **NullableString** | For internal use only. Set to null. | [optional] 
**DeviceVersion** | Pointer to **NullableInt32** | For internal use only. | [optional] 
**DisplayName** | Pointer to **NullableString** | The display name for the device. Required. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy. | [optional] 
**IsCompliant** | Pointer to **NullableBool** | true if the device complies with Mobile Device Management (MDM) policies; otherwise, false. Read-only. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not). | [optional] 
**IsManaged** | Pointer to **NullableBool** | true if the device is managed by a Mobile Device Management (MDM) app; otherwise, false. This can only be updated by Intune for any device OS type or by an approved MDM app for Windows OS devices. Supports $filter (eq, ne, not). | [optional] 
**MdmAppId** | Pointer to **NullableString** | Application identifier used to register device into MDM. Read-only. Supports $filter (eq, ne, not, startsWith). | [optional] 
**OnPremisesLastSyncDateTime** | Pointer to **NullableTime** | The last time at which the object was synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z Read-only. Supports $filter (eq, ne, not, ge, le, in). | [optional] 
**OnPremisesSyncEnabled** | Pointer to **NullableBool** | true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; null if this object has never been synced from an on-premises directory (default). Read-only. Supports $filter (eq, ne, not, in, and eq on null values). | [optional] 
**OperatingSystem** | Pointer to **NullableString** | The type of operating system on the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null values). | [optional] 
**OperatingSystemVersion** | Pointer to **NullableString** | The version of the operating system on the device. Required. Supports $filter (eq, ne, not, ge, le, startsWith, and eq on null values). | [optional] 
**PhysicalIds** | Pointer to **[]string** | For internal use only. Not nullable. Supports $filter (eq, not, ge, le, startsWith). | [optional] 
**ProfileType** | Pointer to **NullableString** | The profile type of the device. Possible values: RegisteredDevice (default), SecureVM, Printer, Shared, IoT. | [optional] 
**SystemLabels** | Pointer to **[]string** | List of labels applied to the device by the system. | [optional] 
**TrustType** | Pointer to **NullableString** | Type of trust for the joined device. Read-only. Possible values:  Workplace (indicates bring your own personal devices), AzureAd (Cloud only joined devices), ServerAd (on-premises domain joined devices joined to Azure AD). For more details, see Introduction to device management in Azure Active Directory | [optional] 
**MemberOf** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Groups that this device is a member of. Read-only. Nullable. Supports $expand. | [optional] 
**RegisteredOwners** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | The user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand. | [optional] 
**RegisteredUsers** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand. | [optional] 
**TransitiveMemberOf** | Pointer to [**[]MicrosoftGraphDirectoryObject**](MicrosoftGraphDirectoryObject.md) | Groups that the device is a member of. This operation is transitive. Supports $expand. | [optional] 
**Extensions** | Pointer to [**[]MicrosoftGraphExtension**](MicrosoftGraphExtension.md) | The collection of open extensions defined for the device. Read-only. Nullable. | [optional] 

## Methods

### NewMicrosoftGraphDevice

`func NewMicrosoftGraphDevice() *MicrosoftGraphDevice`

NewMicrosoftGraphDevice instantiates a new MicrosoftGraphDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDeviceWithDefaults

`func NewMicrosoftGraphDeviceWithDefaults() *MicrosoftGraphDevice`

NewMicrosoftGraphDeviceWithDefaults instantiates a new MicrosoftGraphDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDevice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDevice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDevice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDevice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeletedDateTime

`func (o *MicrosoftGraphDevice) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *MicrosoftGraphDevice) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *MicrosoftGraphDevice) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *MicrosoftGraphDevice) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

### SetDeletedDateTimeNil

`func (o *MicrosoftGraphDevice) SetDeletedDateTimeNil(b bool)`

 SetDeletedDateTimeNil sets the value for DeletedDateTime to be an explicit nil

### UnsetDeletedDateTime
`func (o *MicrosoftGraphDevice) UnsetDeletedDateTime()`

UnsetDeletedDateTime ensures that no value is present for DeletedDateTime, not even an explicit nil
### GetAccountEnabled

`func (o *MicrosoftGraphDevice) GetAccountEnabled() bool`

GetAccountEnabled returns the AccountEnabled field if non-nil, zero value otherwise.

### GetAccountEnabledOk

`func (o *MicrosoftGraphDevice) GetAccountEnabledOk() (*bool, bool)`

GetAccountEnabledOk returns a tuple with the AccountEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEnabled

`func (o *MicrosoftGraphDevice) SetAccountEnabled(v bool)`

SetAccountEnabled sets AccountEnabled field to given value.

### HasAccountEnabled

`func (o *MicrosoftGraphDevice) HasAccountEnabled() bool`

HasAccountEnabled returns a boolean if a field has been set.

### SetAccountEnabledNil

`func (o *MicrosoftGraphDevice) SetAccountEnabledNil(b bool)`

 SetAccountEnabledNil sets the value for AccountEnabled to be an explicit nil

### UnsetAccountEnabled
`func (o *MicrosoftGraphDevice) UnsetAccountEnabled()`

UnsetAccountEnabled ensures that no value is present for AccountEnabled, not even an explicit nil
### GetAlternativeSecurityIds

`func (o *MicrosoftGraphDevice) GetAlternativeSecurityIds() []MicrosoftGraphAlternativeSecurityId`

GetAlternativeSecurityIds returns the AlternativeSecurityIds field if non-nil, zero value otherwise.

### GetAlternativeSecurityIdsOk

`func (o *MicrosoftGraphDevice) GetAlternativeSecurityIdsOk() (*[]MicrosoftGraphAlternativeSecurityId, bool)`

GetAlternativeSecurityIdsOk returns a tuple with the AlternativeSecurityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeSecurityIds

`func (o *MicrosoftGraphDevice) SetAlternativeSecurityIds(v []MicrosoftGraphAlternativeSecurityId)`

SetAlternativeSecurityIds sets AlternativeSecurityIds field to given value.

### HasAlternativeSecurityIds

`func (o *MicrosoftGraphDevice) HasAlternativeSecurityIds() bool`

HasAlternativeSecurityIds returns a boolean if a field has been set.

### GetApproximateLastSignInDateTime

`func (o *MicrosoftGraphDevice) GetApproximateLastSignInDateTime() time.Time`

GetApproximateLastSignInDateTime returns the ApproximateLastSignInDateTime field if non-nil, zero value otherwise.

### GetApproximateLastSignInDateTimeOk

`func (o *MicrosoftGraphDevice) GetApproximateLastSignInDateTimeOk() (*time.Time, bool)`

GetApproximateLastSignInDateTimeOk returns a tuple with the ApproximateLastSignInDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximateLastSignInDateTime

`func (o *MicrosoftGraphDevice) SetApproximateLastSignInDateTime(v time.Time)`

SetApproximateLastSignInDateTime sets ApproximateLastSignInDateTime field to given value.

### HasApproximateLastSignInDateTime

`func (o *MicrosoftGraphDevice) HasApproximateLastSignInDateTime() bool`

HasApproximateLastSignInDateTime returns a boolean if a field has been set.

### SetApproximateLastSignInDateTimeNil

`func (o *MicrosoftGraphDevice) SetApproximateLastSignInDateTimeNil(b bool)`

 SetApproximateLastSignInDateTimeNil sets the value for ApproximateLastSignInDateTime to be an explicit nil

### UnsetApproximateLastSignInDateTime
`func (o *MicrosoftGraphDevice) UnsetApproximateLastSignInDateTime()`

UnsetApproximateLastSignInDateTime ensures that no value is present for ApproximateLastSignInDateTime, not even an explicit nil
### GetComplianceExpirationDateTime

`func (o *MicrosoftGraphDevice) GetComplianceExpirationDateTime() time.Time`

GetComplianceExpirationDateTime returns the ComplianceExpirationDateTime field if non-nil, zero value otherwise.

### GetComplianceExpirationDateTimeOk

`func (o *MicrosoftGraphDevice) GetComplianceExpirationDateTimeOk() (*time.Time, bool)`

GetComplianceExpirationDateTimeOk returns a tuple with the ComplianceExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceExpirationDateTime

`func (o *MicrosoftGraphDevice) SetComplianceExpirationDateTime(v time.Time)`

SetComplianceExpirationDateTime sets ComplianceExpirationDateTime field to given value.

### HasComplianceExpirationDateTime

`func (o *MicrosoftGraphDevice) HasComplianceExpirationDateTime() bool`

HasComplianceExpirationDateTime returns a boolean if a field has been set.

### SetComplianceExpirationDateTimeNil

`func (o *MicrosoftGraphDevice) SetComplianceExpirationDateTimeNil(b bool)`

 SetComplianceExpirationDateTimeNil sets the value for ComplianceExpirationDateTime to be an explicit nil

### UnsetComplianceExpirationDateTime
`func (o *MicrosoftGraphDevice) UnsetComplianceExpirationDateTime()`

UnsetComplianceExpirationDateTime ensures that no value is present for ComplianceExpirationDateTime, not even an explicit nil
### GetDeviceId

`func (o *MicrosoftGraphDevice) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *MicrosoftGraphDevice) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *MicrosoftGraphDevice) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *MicrosoftGraphDevice) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *MicrosoftGraphDevice) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *MicrosoftGraphDevice) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetDeviceMetadata

`func (o *MicrosoftGraphDevice) GetDeviceMetadata() string`

GetDeviceMetadata returns the DeviceMetadata field if non-nil, zero value otherwise.

### GetDeviceMetadataOk

`func (o *MicrosoftGraphDevice) GetDeviceMetadataOk() (*string, bool)`

GetDeviceMetadataOk returns a tuple with the DeviceMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMetadata

`func (o *MicrosoftGraphDevice) SetDeviceMetadata(v string)`

SetDeviceMetadata sets DeviceMetadata field to given value.

### HasDeviceMetadata

`func (o *MicrosoftGraphDevice) HasDeviceMetadata() bool`

HasDeviceMetadata returns a boolean if a field has been set.

### SetDeviceMetadataNil

`func (o *MicrosoftGraphDevice) SetDeviceMetadataNil(b bool)`

 SetDeviceMetadataNil sets the value for DeviceMetadata to be an explicit nil

### UnsetDeviceMetadata
`func (o *MicrosoftGraphDevice) UnsetDeviceMetadata()`

UnsetDeviceMetadata ensures that no value is present for DeviceMetadata, not even an explicit nil
### GetDeviceVersion

`func (o *MicrosoftGraphDevice) GetDeviceVersion() int32`

GetDeviceVersion returns the DeviceVersion field if non-nil, zero value otherwise.

### GetDeviceVersionOk

`func (o *MicrosoftGraphDevice) GetDeviceVersionOk() (*int32, bool)`

GetDeviceVersionOk returns a tuple with the DeviceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceVersion

`func (o *MicrosoftGraphDevice) SetDeviceVersion(v int32)`

SetDeviceVersion sets DeviceVersion field to given value.

### HasDeviceVersion

`func (o *MicrosoftGraphDevice) HasDeviceVersion() bool`

HasDeviceVersion returns a boolean if a field has been set.

### SetDeviceVersionNil

`func (o *MicrosoftGraphDevice) SetDeviceVersionNil(b bool)`

 SetDeviceVersionNil sets the value for DeviceVersion to be an explicit nil

### UnsetDeviceVersion
`func (o *MicrosoftGraphDevice) UnsetDeviceVersion()`

UnsetDeviceVersion ensures that no value is present for DeviceVersion, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphDevice) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphDevice) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphDevice) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphDevice) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphDevice) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphDevice) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetIsCompliant

`func (o *MicrosoftGraphDevice) GetIsCompliant() bool`

GetIsCompliant returns the IsCompliant field if non-nil, zero value otherwise.

### GetIsCompliantOk

`func (o *MicrosoftGraphDevice) GetIsCompliantOk() (*bool, bool)`

GetIsCompliantOk returns a tuple with the IsCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompliant

`func (o *MicrosoftGraphDevice) SetIsCompliant(v bool)`

SetIsCompliant sets IsCompliant field to given value.

### HasIsCompliant

`func (o *MicrosoftGraphDevice) HasIsCompliant() bool`

HasIsCompliant returns a boolean if a field has been set.

### SetIsCompliantNil

`func (o *MicrosoftGraphDevice) SetIsCompliantNil(b bool)`

 SetIsCompliantNil sets the value for IsCompliant to be an explicit nil

### UnsetIsCompliant
`func (o *MicrosoftGraphDevice) UnsetIsCompliant()`

UnsetIsCompliant ensures that no value is present for IsCompliant, not even an explicit nil
### GetIsManaged

`func (o *MicrosoftGraphDevice) GetIsManaged() bool`

GetIsManaged returns the IsManaged field if non-nil, zero value otherwise.

### GetIsManagedOk

`func (o *MicrosoftGraphDevice) GetIsManagedOk() (*bool, bool)`

GetIsManagedOk returns a tuple with the IsManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManaged

`func (o *MicrosoftGraphDevice) SetIsManaged(v bool)`

SetIsManaged sets IsManaged field to given value.

### HasIsManaged

`func (o *MicrosoftGraphDevice) HasIsManaged() bool`

HasIsManaged returns a boolean if a field has been set.

### SetIsManagedNil

`func (o *MicrosoftGraphDevice) SetIsManagedNil(b bool)`

 SetIsManagedNil sets the value for IsManaged to be an explicit nil

### UnsetIsManaged
`func (o *MicrosoftGraphDevice) UnsetIsManaged()`

UnsetIsManaged ensures that no value is present for IsManaged, not even an explicit nil
### GetMdmAppId

`func (o *MicrosoftGraphDevice) GetMdmAppId() string`

GetMdmAppId returns the MdmAppId field if non-nil, zero value otherwise.

### GetMdmAppIdOk

`func (o *MicrosoftGraphDevice) GetMdmAppIdOk() (*string, bool)`

GetMdmAppIdOk returns a tuple with the MdmAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmAppId

`func (o *MicrosoftGraphDevice) SetMdmAppId(v string)`

SetMdmAppId sets MdmAppId field to given value.

### HasMdmAppId

`func (o *MicrosoftGraphDevice) HasMdmAppId() bool`

HasMdmAppId returns a boolean if a field has been set.

### SetMdmAppIdNil

`func (o *MicrosoftGraphDevice) SetMdmAppIdNil(b bool)`

 SetMdmAppIdNil sets the value for MdmAppId to be an explicit nil

### UnsetMdmAppId
`func (o *MicrosoftGraphDevice) UnsetMdmAppId()`

UnsetMdmAppId ensures that no value is present for MdmAppId, not even an explicit nil
### GetOnPremisesLastSyncDateTime

`func (o *MicrosoftGraphDevice) GetOnPremisesLastSyncDateTime() time.Time`

GetOnPremisesLastSyncDateTime returns the OnPremisesLastSyncDateTime field if non-nil, zero value otherwise.

### GetOnPremisesLastSyncDateTimeOk

`func (o *MicrosoftGraphDevice) GetOnPremisesLastSyncDateTimeOk() (*time.Time, bool)`

GetOnPremisesLastSyncDateTimeOk returns a tuple with the OnPremisesLastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesLastSyncDateTime

`func (o *MicrosoftGraphDevice) SetOnPremisesLastSyncDateTime(v time.Time)`

SetOnPremisesLastSyncDateTime sets OnPremisesLastSyncDateTime field to given value.

### HasOnPremisesLastSyncDateTime

`func (o *MicrosoftGraphDevice) HasOnPremisesLastSyncDateTime() bool`

HasOnPremisesLastSyncDateTime returns a boolean if a field has been set.

### SetOnPremisesLastSyncDateTimeNil

`func (o *MicrosoftGraphDevice) SetOnPremisesLastSyncDateTimeNil(b bool)`

 SetOnPremisesLastSyncDateTimeNil sets the value for OnPremisesLastSyncDateTime to be an explicit nil

### UnsetOnPremisesLastSyncDateTime
`func (o *MicrosoftGraphDevice) UnsetOnPremisesLastSyncDateTime()`

UnsetOnPremisesLastSyncDateTime ensures that no value is present for OnPremisesLastSyncDateTime, not even an explicit nil
### GetOnPremisesSyncEnabled

`func (o *MicrosoftGraphDevice) GetOnPremisesSyncEnabled() bool`

GetOnPremisesSyncEnabled returns the OnPremisesSyncEnabled field if non-nil, zero value otherwise.

### GetOnPremisesSyncEnabledOk

`func (o *MicrosoftGraphDevice) GetOnPremisesSyncEnabledOk() (*bool, bool)`

GetOnPremisesSyncEnabledOk returns a tuple with the OnPremisesSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesSyncEnabled

`func (o *MicrosoftGraphDevice) SetOnPremisesSyncEnabled(v bool)`

SetOnPremisesSyncEnabled sets OnPremisesSyncEnabled field to given value.

### HasOnPremisesSyncEnabled

`func (o *MicrosoftGraphDevice) HasOnPremisesSyncEnabled() bool`

HasOnPremisesSyncEnabled returns a boolean if a field has been set.

### SetOnPremisesSyncEnabledNil

`func (o *MicrosoftGraphDevice) SetOnPremisesSyncEnabledNil(b bool)`

 SetOnPremisesSyncEnabledNil sets the value for OnPremisesSyncEnabled to be an explicit nil

### UnsetOnPremisesSyncEnabled
`func (o *MicrosoftGraphDevice) UnsetOnPremisesSyncEnabled()`

UnsetOnPremisesSyncEnabled ensures that no value is present for OnPremisesSyncEnabled, not even an explicit nil
### GetOperatingSystem

`func (o *MicrosoftGraphDevice) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *MicrosoftGraphDevice) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *MicrosoftGraphDevice) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *MicrosoftGraphDevice) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### SetOperatingSystemNil

`func (o *MicrosoftGraphDevice) SetOperatingSystemNil(b bool)`

 SetOperatingSystemNil sets the value for OperatingSystem to be an explicit nil

### UnsetOperatingSystem
`func (o *MicrosoftGraphDevice) UnsetOperatingSystem()`

UnsetOperatingSystem ensures that no value is present for OperatingSystem, not even an explicit nil
### GetOperatingSystemVersion

`func (o *MicrosoftGraphDevice) GetOperatingSystemVersion() string`

GetOperatingSystemVersion returns the OperatingSystemVersion field if non-nil, zero value otherwise.

### GetOperatingSystemVersionOk

`func (o *MicrosoftGraphDevice) GetOperatingSystemVersionOk() (*string, bool)`

GetOperatingSystemVersionOk returns a tuple with the OperatingSystemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemVersion

`func (o *MicrosoftGraphDevice) SetOperatingSystemVersion(v string)`

SetOperatingSystemVersion sets OperatingSystemVersion field to given value.

### HasOperatingSystemVersion

`func (o *MicrosoftGraphDevice) HasOperatingSystemVersion() bool`

HasOperatingSystemVersion returns a boolean if a field has been set.

### SetOperatingSystemVersionNil

`func (o *MicrosoftGraphDevice) SetOperatingSystemVersionNil(b bool)`

 SetOperatingSystemVersionNil sets the value for OperatingSystemVersion to be an explicit nil

### UnsetOperatingSystemVersion
`func (o *MicrosoftGraphDevice) UnsetOperatingSystemVersion()`

UnsetOperatingSystemVersion ensures that no value is present for OperatingSystemVersion, not even an explicit nil
### GetPhysicalIds

`func (o *MicrosoftGraphDevice) GetPhysicalIds() []string`

GetPhysicalIds returns the PhysicalIds field if non-nil, zero value otherwise.

### GetPhysicalIdsOk

`func (o *MicrosoftGraphDevice) GetPhysicalIdsOk() (*[]string, bool)`

GetPhysicalIdsOk returns a tuple with the PhysicalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalIds

`func (o *MicrosoftGraphDevice) SetPhysicalIds(v []string)`

SetPhysicalIds sets PhysicalIds field to given value.

### HasPhysicalIds

`func (o *MicrosoftGraphDevice) HasPhysicalIds() bool`

HasPhysicalIds returns a boolean if a field has been set.

### GetProfileType

`func (o *MicrosoftGraphDevice) GetProfileType() string`

GetProfileType returns the ProfileType field if non-nil, zero value otherwise.

### GetProfileTypeOk

`func (o *MicrosoftGraphDevice) GetProfileTypeOk() (*string, bool)`

GetProfileTypeOk returns a tuple with the ProfileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileType

`func (o *MicrosoftGraphDevice) SetProfileType(v string)`

SetProfileType sets ProfileType field to given value.

### HasProfileType

`func (o *MicrosoftGraphDevice) HasProfileType() bool`

HasProfileType returns a boolean if a field has been set.

### SetProfileTypeNil

`func (o *MicrosoftGraphDevice) SetProfileTypeNil(b bool)`

 SetProfileTypeNil sets the value for ProfileType to be an explicit nil

### UnsetProfileType
`func (o *MicrosoftGraphDevice) UnsetProfileType()`

UnsetProfileType ensures that no value is present for ProfileType, not even an explicit nil
### GetSystemLabels

`func (o *MicrosoftGraphDevice) GetSystemLabels() []string`

GetSystemLabels returns the SystemLabels field if non-nil, zero value otherwise.

### GetSystemLabelsOk

`func (o *MicrosoftGraphDevice) GetSystemLabelsOk() (*[]string, bool)`

GetSystemLabelsOk returns a tuple with the SystemLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemLabels

`func (o *MicrosoftGraphDevice) SetSystemLabels(v []string)`

SetSystemLabels sets SystemLabels field to given value.

### HasSystemLabels

`func (o *MicrosoftGraphDevice) HasSystemLabels() bool`

HasSystemLabels returns a boolean if a field has been set.

### GetTrustType

`func (o *MicrosoftGraphDevice) GetTrustType() string`

GetTrustType returns the TrustType field if non-nil, zero value otherwise.

### GetTrustTypeOk

`func (o *MicrosoftGraphDevice) GetTrustTypeOk() (*string, bool)`

GetTrustTypeOk returns a tuple with the TrustType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustType

`func (o *MicrosoftGraphDevice) SetTrustType(v string)`

SetTrustType sets TrustType field to given value.

### HasTrustType

`func (o *MicrosoftGraphDevice) HasTrustType() bool`

HasTrustType returns a boolean if a field has been set.

### SetTrustTypeNil

`func (o *MicrosoftGraphDevice) SetTrustTypeNil(b bool)`

 SetTrustTypeNil sets the value for TrustType to be an explicit nil

### UnsetTrustType
`func (o *MicrosoftGraphDevice) UnsetTrustType()`

UnsetTrustType ensures that no value is present for TrustType, not even an explicit nil
### GetMemberOf

`func (o *MicrosoftGraphDevice) GetMemberOf() []MicrosoftGraphDirectoryObject`

GetMemberOf returns the MemberOf field if non-nil, zero value otherwise.

### GetMemberOfOk

`func (o *MicrosoftGraphDevice) GetMemberOfOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetMemberOfOk returns a tuple with the MemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOf

`func (o *MicrosoftGraphDevice) SetMemberOf(v []MicrosoftGraphDirectoryObject)`

SetMemberOf sets MemberOf field to given value.

### HasMemberOf

`func (o *MicrosoftGraphDevice) HasMemberOf() bool`

HasMemberOf returns a boolean if a field has been set.

### GetRegisteredOwners

`func (o *MicrosoftGraphDevice) GetRegisteredOwners() []MicrosoftGraphDirectoryObject`

GetRegisteredOwners returns the RegisteredOwners field if non-nil, zero value otherwise.

### GetRegisteredOwnersOk

`func (o *MicrosoftGraphDevice) GetRegisteredOwnersOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetRegisteredOwnersOk returns a tuple with the RegisteredOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredOwners

`func (o *MicrosoftGraphDevice) SetRegisteredOwners(v []MicrosoftGraphDirectoryObject)`

SetRegisteredOwners sets RegisteredOwners field to given value.

### HasRegisteredOwners

`func (o *MicrosoftGraphDevice) HasRegisteredOwners() bool`

HasRegisteredOwners returns a boolean if a field has been set.

### GetRegisteredUsers

`func (o *MicrosoftGraphDevice) GetRegisteredUsers() []MicrosoftGraphDirectoryObject`

GetRegisteredUsers returns the RegisteredUsers field if non-nil, zero value otherwise.

### GetRegisteredUsersOk

`func (o *MicrosoftGraphDevice) GetRegisteredUsersOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetRegisteredUsersOk returns a tuple with the RegisteredUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredUsers

`func (o *MicrosoftGraphDevice) SetRegisteredUsers(v []MicrosoftGraphDirectoryObject)`

SetRegisteredUsers sets RegisteredUsers field to given value.

### HasRegisteredUsers

`func (o *MicrosoftGraphDevice) HasRegisteredUsers() bool`

HasRegisteredUsers returns a boolean if a field has been set.

### GetTransitiveMemberOf

`func (o *MicrosoftGraphDevice) GetTransitiveMemberOf() []MicrosoftGraphDirectoryObject`

GetTransitiveMemberOf returns the TransitiveMemberOf field if non-nil, zero value otherwise.

### GetTransitiveMemberOfOk

`func (o *MicrosoftGraphDevice) GetTransitiveMemberOfOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetTransitiveMemberOfOk returns a tuple with the TransitiveMemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitiveMemberOf

`func (o *MicrosoftGraphDevice) SetTransitiveMemberOf(v []MicrosoftGraphDirectoryObject)`

SetTransitiveMemberOf sets TransitiveMemberOf field to given value.

### HasTransitiveMemberOf

`func (o *MicrosoftGraphDevice) HasTransitiveMemberOf() bool`

HasTransitiveMemberOf returns a boolean if a field has been set.

### GetExtensions

`func (o *MicrosoftGraphDevice) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *MicrosoftGraphDevice) GetExtensionsOk() (*[]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *MicrosoftGraphDevice) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *MicrosoftGraphDevice) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


