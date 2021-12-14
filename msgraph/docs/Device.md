# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewDevice

`func NewDevice() *Device`

NewDevice instantiates a new Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithDefaults

`func NewDeviceWithDefaults() *Device`

NewDeviceWithDefaults instantiates a new Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountEnabled

`func (o *Device) GetAccountEnabled() bool`

GetAccountEnabled returns the AccountEnabled field if non-nil, zero value otherwise.

### GetAccountEnabledOk

`func (o *Device) GetAccountEnabledOk() (*bool, bool)`

GetAccountEnabledOk returns a tuple with the AccountEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountEnabled

`func (o *Device) SetAccountEnabled(v bool)`

SetAccountEnabled sets AccountEnabled field to given value.

### HasAccountEnabled

`func (o *Device) HasAccountEnabled() bool`

HasAccountEnabled returns a boolean if a field has been set.

### SetAccountEnabledNil

`func (o *Device) SetAccountEnabledNil(b bool)`

 SetAccountEnabledNil sets the value for AccountEnabled to be an explicit nil

### UnsetAccountEnabled
`func (o *Device) UnsetAccountEnabled()`

UnsetAccountEnabled ensures that no value is present for AccountEnabled, not even an explicit nil
### GetAlternativeSecurityIds

`func (o *Device) GetAlternativeSecurityIds() []MicrosoftGraphAlternativeSecurityId`

GetAlternativeSecurityIds returns the AlternativeSecurityIds field if non-nil, zero value otherwise.

### GetAlternativeSecurityIdsOk

`func (o *Device) GetAlternativeSecurityIdsOk() (*[]MicrosoftGraphAlternativeSecurityId, bool)`

GetAlternativeSecurityIdsOk returns a tuple with the AlternativeSecurityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeSecurityIds

`func (o *Device) SetAlternativeSecurityIds(v []MicrosoftGraphAlternativeSecurityId)`

SetAlternativeSecurityIds sets AlternativeSecurityIds field to given value.

### HasAlternativeSecurityIds

`func (o *Device) HasAlternativeSecurityIds() bool`

HasAlternativeSecurityIds returns a boolean if a field has been set.

### GetApproximateLastSignInDateTime

`func (o *Device) GetApproximateLastSignInDateTime() time.Time`

GetApproximateLastSignInDateTime returns the ApproximateLastSignInDateTime field if non-nil, zero value otherwise.

### GetApproximateLastSignInDateTimeOk

`func (o *Device) GetApproximateLastSignInDateTimeOk() (*time.Time, bool)`

GetApproximateLastSignInDateTimeOk returns a tuple with the ApproximateLastSignInDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproximateLastSignInDateTime

`func (o *Device) SetApproximateLastSignInDateTime(v time.Time)`

SetApproximateLastSignInDateTime sets ApproximateLastSignInDateTime field to given value.

### HasApproximateLastSignInDateTime

`func (o *Device) HasApproximateLastSignInDateTime() bool`

HasApproximateLastSignInDateTime returns a boolean if a field has been set.

### SetApproximateLastSignInDateTimeNil

`func (o *Device) SetApproximateLastSignInDateTimeNil(b bool)`

 SetApproximateLastSignInDateTimeNil sets the value for ApproximateLastSignInDateTime to be an explicit nil

### UnsetApproximateLastSignInDateTime
`func (o *Device) UnsetApproximateLastSignInDateTime()`

UnsetApproximateLastSignInDateTime ensures that no value is present for ApproximateLastSignInDateTime, not even an explicit nil
### GetComplianceExpirationDateTime

`func (o *Device) GetComplianceExpirationDateTime() time.Time`

GetComplianceExpirationDateTime returns the ComplianceExpirationDateTime field if non-nil, zero value otherwise.

### GetComplianceExpirationDateTimeOk

`func (o *Device) GetComplianceExpirationDateTimeOk() (*time.Time, bool)`

GetComplianceExpirationDateTimeOk returns a tuple with the ComplianceExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplianceExpirationDateTime

`func (o *Device) SetComplianceExpirationDateTime(v time.Time)`

SetComplianceExpirationDateTime sets ComplianceExpirationDateTime field to given value.

### HasComplianceExpirationDateTime

`func (o *Device) HasComplianceExpirationDateTime() bool`

HasComplianceExpirationDateTime returns a boolean if a field has been set.

### SetComplianceExpirationDateTimeNil

`func (o *Device) SetComplianceExpirationDateTimeNil(b bool)`

 SetComplianceExpirationDateTimeNil sets the value for ComplianceExpirationDateTime to be an explicit nil

### UnsetComplianceExpirationDateTime
`func (o *Device) UnsetComplianceExpirationDateTime()`

UnsetComplianceExpirationDateTime ensures that no value is present for ComplianceExpirationDateTime, not even an explicit nil
### GetDeviceId

`func (o *Device) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *Device) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *Device) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *Device) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *Device) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *Device) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetDeviceMetadata

`func (o *Device) GetDeviceMetadata() string`

GetDeviceMetadata returns the DeviceMetadata field if non-nil, zero value otherwise.

### GetDeviceMetadataOk

`func (o *Device) GetDeviceMetadataOk() (*string, bool)`

GetDeviceMetadataOk returns a tuple with the DeviceMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMetadata

`func (o *Device) SetDeviceMetadata(v string)`

SetDeviceMetadata sets DeviceMetadata field to given value.

### HasDeviceMetadata

`func (o *Device) HasDeviceMetadata() bool`

HasDeviceMetadata returns a boolean if a field has been set.

### SetDeviceMetadataNil

`func (o *Device) SetDeviceMetadataNil(b bool)`

 SetDeviceMetadataNil sets the value for DeviceMetadata to be an explicit nil

### UnsetDeviceMetadata
`func (o *Device) UnsetDeviceMetadata()`

UnsetDeviceMetadata ensures that no value is present for DeviceMetadata, not even an explicit nil
### GetDeviceVersion

`func (o *Device) GetDeviceVersion() int32`

GetDeviceVersion returns the DeviceVersion field if non-nil, zero value otherwise.

### GetDeviceVersionOk

`func (o *Device) GetDeviceVersionOk() (*int32, bool)`

GetDeviceVersionOk returns a tuple with the DeviceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceVersion

`func (o *Device) SetDeviceVersion(v int32)`

SetDeviceVersion sets DeviceVersion field to given value.

### HasDeviceVersion

`func (o *Device) HasDeviceVersion() bool`

HasDeviceVersion returns a boolean if a field has been set.

### SetDeviceVersionNil

`func (o *Device) SetDeviceVersionNil(b bool)`

 SetDeviceVersionNil sets the value for DeviceVersion to be an explicit nil

### UnsetDeviceVersion
`func (o *Device) UnsetDeviceVersion()`

UnsetDeviceVersion ensures that no value is present for DeviceVersion, not even an explicit nil
### GetDisplayName

`func (o *Device) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Device) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Device) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Device) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *Device) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *Device) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetIsCompliant

`func (o *Device) GetIsCompliant() bool`

GetIsCompliant returns the IsCompliant field if non-nil, zero value otherwise.

### GetIsCompliantOk

`func (o *Device) GetIsCompliantOk() (*bool, bool)`

GetIsCompliantOk returns a tuple with the IsCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompliant

`func (o *Device) SetIsCompliant(v bool)`

SetIsCompliant sets IsCompliant field to given value.

### HasIsCompliant

`func (o *Device) HasIsCompliant() bool`

HasIsCompliant returns a boolean if a field has been set.

### SetIsCompliantNil

`func (o *Device) SetIsCompliantNil(b bool)`

 SetIsCompliantNil sets the value for IsCompliant to be an explicit nil

### UnsetIsCompliant
`func (o *Device) UnsetIsCompliant()`

UnsetIsCompliant ensures that no value is present for IsCompliant, not even an explicit nil
### GetIsManaged

`func (o *Device) GetIsManaged() bool`

GetIsManaged returns the IsManaged field if non-nil, zero value otherwise.

### GetIsManagedOk

`func (o *Device) GetIsManagedOk() (*bool, bool)`

GetIsManagedOk returns a tuple with the IsManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManaged

`func (o *Device) SetIsManaged(v bool)`

SetIsManaged sets IsManaged field to given value.

### HasIsManaged

`func (o *Device) HasIsManaged() bool`

HasIsManaged returns a boolean if a field has been set.

### SetIsManagedNil

`func (o *Device) SetIsManagedNil(b bool)`

 SetIsManagedNil sets the value for IsManaged to be an explicit nil

### UnsetIsManaged
`func (o *Device) UnsetIsManaged()`

UnsetIsManaged ensures that no value is present for IsManaged, not even an explicit nil
### GetMdmAppId

`func (o *Device) GetMdmAppId() string`

GetMdmAppId returns the MdmAppId field if non-nil, zero value otherwise.

### GetMdmAppIdOk

`func (o *Device) GetMdmAppIdOk() (*string, bool)`

GetMdmAppIdOk returns a tuple with the MdmAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdmAppId

`func (o *Device) SetMdmAppId(v string)`

SetMdmAppId sets MdmAppId field to given value.

### HasMdmAppId

`func (o *Device) HasMdmAppId() bool`

HasMdmAppId returns a boolean if a field has been set.

### SetMdmAppIdNil

`func (o *Device) SetMdmAppIdNil(b bool)`

 SetMdmAppIdNil sets the value for MdmAppId to be an explicit nil

### UnsetMdmAppId
`func (o *Device) UnsetMdmAppId()`

UnsetMdmAppId ensures that no value is present for MdmAppId, not even an explicit nil
### GetOnPremisesLastSyncDateTime

`func (o *Device) GetOnPremisesLastSyncDateTime() time.Time`

GetOnPremisesLastSyncDateTime returns the OnPremisesLastSyncDateTime field if non-nil, zero value otherwise.

### GetOnPremisesLastSyncDateTimeOk

`func (o *Device) GetOnPremisesLastSyncDateTimeOk() (*time.Time, bool)`

GetOnPremisesLastSyncDateTimeOk returns a tuple with the OnPremisesLastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesLastSyncDateTime

`func (o *Device) SetOnPremisesLastSyncDateTime(v time.Time)`

SetOnPremisesLastSyncDateTime sets OnPremisesLastSyncDateTime field to given value.

### HasOnPremisesLastSyncDateTime

`func (o *Device) HasOnPremisesLastSyncDateTime() bool`

HasOnPremisesLastSyncDateTime returns a boolean if a field has been set.

### SetOnPremisesLastSyncDateTimeNil

`func (o *Device) SetOnPremisesLastSyncDateTimeNil(b bool)`

 SetOnPremisesLastSyncDateTimeNil sets the value for OnPremisesLastSyncDateTime to be an explicit nil

### UnsetOnPremisesLastSyncDateTime
`func (o *Device) UnsetOnPremisesLastSyncDateTime()`

UnsetOnPremisesLastSyncDateTime ensures that no value is present for OnPremisesLastSyncDateTime, not even an explicit nil
### GetOnPremisesSyncEnabled

`func (o *Device) GetOnPremisesSyncEnabled() bool`

GetOnPremisesSyncEnabled returns the OnPremisesSyncEnabled field if non-nil, zero value otherwise.

### GetOnPremisesSyncEnabledOk

`func (o *Device) GetOnPremisesSyncEnabledOk() (*bool, bool)`

GetOnPremisesSyncEnabledOk returns a tuple with the OnPremisesSyncEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPremisesSyncEnabled

`func (o *Device) SetOnPremisesSyncEnabled(v bool)`

SetOnPremisesSyncEnabled sets OnPremisesSyncEnabled field to given value.

### HasOnPremisesSyncEnabled

`func (o *Device) HasOnPremisesSyncEnabled() bool`

HasOnPremisesSyncEnabled returns a boolean if a field has been set.

### SetOnPremisesSyncEnabledNil

`func (o *Device) SetOnPremisesSyncEnabledNil(b bool)`

 SetOnPremisesSyncEnabledNil sets the value for OnPremisesSyncEnabled to be an explicit nil

### UnsetOnPremisesSyncEnabled
`func (o *Device) UnsetOnPremisesSyncEnabled()`

UnsetOnPremisesSyncEnabled ensures that no value is present for OnPremisesSyncEnabled, not even an explicit nil
### GetOperatingSystem

`func (o *Device) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *Device) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *Device) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *Device) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### SetOperatingSystemNil

`func (o *Device) SetOperatingSystemNil(b bool)`

 SetOperatingSystemNil sets the value for OperatingSystem to be an explicit nil

### UnsetOperatingSystem
`func (o *Device) UnsetOperatingSystem()`

UnsetOperatingSystem ensures that no value is present for OperatingSystem, not even an explicit nil
### GetOperatingSystemVersion

`func (o *Device) GetOperatingSystemVersion() string`

GetOperatingSystemVersion returns the OperatingSystemVersion field if non-nil, zero value otherwise.

### GetOperatingSystemVersionOk

`func (o *Device) GetOperatingSystemVersionOk() (*string, bool)`

GetOperatingSystemVersionOk returns a tuple with the OperatingSystemVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemVersion

`func (o *Device) SetOperatingSystemVersion(v string)`

SetOperatingSystemVersion sets OperatingSystemVersion field to given value.

### HasOperatingSystemVersion

`func (o *Device) HasOperatingSystemVersion() bool`

HasOperatingSystemVersion returns a boolean if a field has been set.

### SetOperatingSystemVersionNil

`func (o *Device) SetOperatingSystemVersionNil(b bool)`

 SetOperatingSystemVersionNil sets the value for OperatingSystemVersion to be an explicit nil

### UnsetOperatingSystemVersion
`func (o *Device) UnsetOperatingSystemVersion()`

UnsetOperatingSystemVersion ensures that no value is present for OperatingSystemVersion, not even an explicit nil
### GetPhysicalIds

`func (o *Device) GetPhysicalIds() []string`

GetPhysicalIds returns the PhysicalIds field if non-nil, zero value otherwise.

### GetPhysicalIdsOk

`func (o *Device) GetPhysicalIdsOk() (*[]string, bool)`

GetPhysicalIdsOk returns a tuple with the PhysicalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalIds

`func (o *Device) SetPhysicalIds(v []string)`

SetPhysicalIds sets PhysicalIds field to given value.

### HasPhysicalIds

`func (o *Device) HasPhysicalIds() bool`

HasPhysicalIds returns a boolean if a field has been set.

### GetProfileType

`func (o *Device) GetProfileType() string`

GetProfileType returns the ProfileType field if non-nil, zero value otherwise.

### GetProfileTypeOk

`func (o *Device) GetProfileTypeOk() (*string, bool)`

GetProfileTypeOk returns a tuple with the ProfileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileType

`func (o *Device) SetProfileType(v string)`

SetProfileType sets ProfileType field to given value.

### HasProfileType

`func (o *Device) HasProfileType() bool`

HasProfileType returns a boolean if a field has been set.

### SetProfileTypeNil

`func (o *Device) SetProfileTypeNil(b bool)`

 SetProfileTypeNil sets the value for ProfileType to be an explicit nil

### UnsetProfileType
`func (o *Device) UnsetProfileType()`

UnsetProfileType ensures that no value is present for ProfileType, not even an explicit nil
### GetSystemLabels

`func (o *Device) GetSystemLabels() []string`

GetSystemLabels returns the SystemLabels field if non-nil, zero value otherwise.

### GetSystemLabelsOk

`func (o *Device) GetSystemLabelsOk() (*[]string, bool)`

GetSystemLabelsOk returns a tuple with the SystemLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemLabels

`func (o *Device) SetSystemLabels(v []string)`

SetSystemLabels sets SystemLabels field to given value.

### HasSystemLabels

`func (o *Device) HasSystemLabels() bool`

HasSystemLabels returns a boolean if a field has been set.

### GetTrustType

`func (o *Device) GetTrustType() string`

GetTrustType returns the TrustType field if non-nil, zero value otherwise.

### GetTrustTypeOk

`func (o *Device) GetTrustTypeOk() (*string, bool)`

GetTrustTypeOk returns a tuple with the TrustType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustType

`func (o *Device) SetTrustType(v string)`

SetTrustType sets TrustType field to given value.

### HasTrustType

`func (o *Device) HasTrustType() bool`

HasTrustType returns a boolean if a field has been set.

### SetTrustTypeNil

`func (o *Device) SetTrustTypeNil(b bool)`

 SetTrustTypeNil sets the value for TrustType to be an explicit nil

### UnsetTrustType
`func (o *Device) UnsetTrustType()`

UnsetTrustType ensures that no value is present for TrustType, not even an explicit nil
### GetMemberOf

`func (o *Device) GetMemberOf() []MicrosoftGraphDirectoryObject`

GetMemberOf returns the MemberOf field if non-nil, zero value otherwise.

### GetMemberOfOk

`func (o *Device) GetMemberOfOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetMemberOfOk returns a tuple with the MemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOf

`func (o *Device) SetMemberOf(v []MicrosoftGraphDirectoryObject)`

SetMemberOf sets MemberOf field to given value.

### HasMemberOf

`func (o *Device) HasMemberOf() bool`

HasMemberOf returns a boolean if a field has been set.

### GetRegisteredOwners

`func (o *Device) GetRegisteredOwners() []MicrosoftGraphDirectoryObject`

GetRegisteredOwners returns the RegisteredOwners field if non-nil, zero value otherwise.

### GetRegisteredOwnersOk

`func (o *Device) GetRegisteredOwnersOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetRegisteredOwnersOk returns a tuple with the RegisteredOwners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredOwners

`func (o *Device) SetRegisteredOwners(v []MicrosoftGraphDirectoryObject)`

SetRegisteredOwners sets RegisteredOwners field to given value.

### HasRegisteredOwners

`func (o *Device) HasRegisteredOwners() bool`

HasRegisteredOwners returns a boolean if a field has been set.

### GetRegisteredUsers

`func (o *Device) GetRegisteredUsers() []MicrosoftGraphDirectoryObject`

GetRegisteredUsers returns the RegisteredUsers field if non-nil, zero value otherwise.

### GetRegisteredUsersOk

`func (o *Device) GetRegisteredUsersOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetRegisteredUsersOk returns a tuple with the RegisteredUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredUsers

`func (o *Device) SetRegisteredUsers(v []MicrosoftGraphDirectoryObject)`

SetRegisteredUsers sets RegisteredUsers field to given value.

### HasRegisteredUsers

`func (o *Device) HasRegisteredUsers() bool`

HasRegisteredUsers returns a boolean if a field has been set.

### GetTransitiveMemberOf

`func (o *Device) GetTransitiveMemberOf() []MicrosoftGraphDirectoryObject`

GetTransitiveMemberOf returns the TransitiveMemberOf field if non-nil, zero value otherwise.

### GetTransitiveMemberOfOk

`func (o *Device) GetTransitiveMemberOfOk() (*[]MicrosoftGraphDirectoryObject, bool)`

GetTransitiveMemberOfOk returns a tuple with the TransitiveMemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitiveMemberOf

`func (o *Device) SetTransitiveMemberOf(v []MicrosoftGraphDirectoryObject)`

SetTransitiveMemberOf sets TransitiveMemberOf field to given value.

### HasTransitiveMemberOf

`func (o *Device) HasTransitiveMemberOf() bool`

HasTransitiveMemberOf returns a boolean if a field has been set.

### GetExtensions

`func (o *Device) GetExtensions() []MicrosoftGraphExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Device) GetExtensionsOk() (*[]MicrosoftGraphExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Device) SetExtensions(v []MicrosoftGraphExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Device) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


