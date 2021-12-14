# MicrosoftGraphWindowsAutopilotDeviceIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**AddressableUserName** | Pointer to **NullableString** | Addressable user name. | [optional] 
**AzureActiveDirectoryDeviceId** | Pointer to **NullableString** | AAD Device ID - to be deprecated | [optional] 
**DisplayName** | Pointer to **NullableString** | Display Name | [optional] 
**EnrollmentState** | Pointer to [**AnyOfmicrosoftGraphEnrollmentState**](anyOf&lt;microsoft.graph.enrollmentState&gt;.md) | Intune enrollment state of the Windows autopilot device. Possible values are: unknown, enrolled, pendingReset, failed, notContacted. | [optional] 
**GroupTag** | Pointer to **NullableString** | Group Tag of the Windows autopilot device. | [optional] 
**LastContactedDateTime** | Pointer to **time.Time** | Intune Last Contacted Date Time of the Windows autopilot device. | [optional] 
**ManagedDeviceId** | Pointer to **NullableString** | Managed Device ID | [optional] 
**Manufacturer** | Pointer to **NullableString** | Oem manufacturer of the Windows autopilot device. | [optional] 
**Model** | Pointer to **NullableString** | Model name of the Windows autopilot device. | [optional] 
**ProductKey** | Pointer to **NullableString** | Product Key of the Windows autopilot device. | [optional] 
**PurchaseOrderIdentifier** | Pointer to **NullableString** | Purchase Order Identifier of the Windows autopilot device. | [optional] 
**ResourceName** | Pointer to **NullableString** | Resource Name. | [optional] 
**SerialNumber** | Pointer to **NullableString** | Serial number of the Windows autopilot device. | [optional] 
**SkuNumber** | Pointer to **NullableString** | SKU Number | [optional] 
**SystemFamily** | Pointer to **NullableString** | System Family | [optional] 
**UserPrincipalName** | Pointer to **NullableString** | User Principal Name. | [optional] 

## Methods

### NewMicrosoftGraphWindowsAutopilotDeviceIdentity

`func NewMicrosoftGraphWindowsAutopilotDeviceIdentity() *MicrosoftGraphWindowsAutopilotDeviceIdentity`

NewMicrosoftGraphWindowsAutopilotDeviceIdentity instantiates a new MicrosoftGraphWindowsAutopilotDeviceIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWindowsAutopilotDeviceIdentityWithDefaults

`func NewMicrosoftGraphWindowsAutopilotDeviceIdentityWithDefaults() *MicrosoftGraphWindowsAutopilotDeviceIdentity`

NewMicrosoftGraphWindowsAutopilotDeviceIdentityWithDefaults instantiates a new MicrosoftGraphWindowsAutopilotDeviceIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAddressableUserName

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetAddressableUserName() string`

GetAddressableUserName returns the AddressableUserName field if non-nil, zero value otherwise.

### GetAddressableUserNameOk

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetAddressableUserNameOk() (*string, bool)`

GetAddressableUserNameOk returns a tuple with the AddressableUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressableUserName

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetAddressableUserName(v string)`

SetAddressableUserName sets AddressableUserName field to given value.

### HasAddressableUserName

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) HasAddressableUserName() bool`

HasAddressableUserName returns a boolean if a field has been set.

### SetAddressableUserNameNil

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetAddressableUserNameNil(b bool)`

 SetAddressableUserNameNil sets the value for AddressableUserName to be an explicit nil

### UnsetAddressableUserName
`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) UnsetAddressableUserName()`

UnsetAddressableUserName ensures that no value is present for AddressableUserName, not even an explicit nil
### GetAzureActiveDirectoryDeviceId

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetAzureActiveDirectoryDeviceId() string`

GetAzureActiveDirectoryDeviceId returns the AzureActiveDirectoryDeviceId field if non-nil, zero value otherwise.

### GetAzureActiveDirectoryDeviceIdOk

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetAzureActiveDirectoryDeviceIdOk() (*string, bool)`

GetAzureActiveDirectoryDeviceIdOk returns a tuple with the AzureActiveDirectoryDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureActiveDirectoryDeviceId

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetAzureActiveDirectoryDeviceId(v string)`

SetAzureActiveDirectoryDeviceId sets AzureActiveDirectoryDeviceId field to given value.

### HasAzureActiveDirectoryDeviceId

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) HasAzureActiveDirectoryDeviceId() bool`

HasAzureActiveDirectoryDeviceId returns a boolean if a field has been set.

### SetAzureActiveDirectoryDeviceIdNil

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetAzureActiveDirectoryDeviceIdNil(b bool)`

 SetAzureActiveDirectoryDeviceIdNil sets the value for AzureActiveDirectoryDeviceId to be an explicit nil

### UnsetAzureActiveDirectoryDeviceId
`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) UnsetAzureActiveDirectoryDeviceId()`

UnsetAzureActiveDirectoryDeviceId ensures that no value is present for AzureActiveDirectoryDeviceId, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetEnrollmentState

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetEnrollmentState() AnyOfmicrosoftGraphEnrollmentState`

GetEnrollmentState returns the EnrollmentState field if non-nil, zero value otherwise.

### GetEnrollmentStateOk

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetEnrollmentStateOk() (*AnyOfmicrosoftGraphEnrollmentState, bool)`

GetEnrollmentStateOk returns a tuple with the EnrollmentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentState

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetEnrollmentState(v AnyOfmicrosoftGraphEnrollmentState)`

SetEnrollmentState sets EnrollmentState field to given value.

### HasEnrollmentState

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) HasEnrollmentState() bool`

HasEnrollmentState returns a boolean if a field has been set.

### SetEnrollmentStateNil

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetEnrollmentStateNil(b bool)`

 SetEnrollmentStateNil sets the value for EnrollmentState to be an explicit nil

### UnsetEnrollmentState
`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) UnsetEnrollmentState()`

UnsetEnrollmentState ensures that no value is present for EnrollmentState, not even an explicit nil
### GetGroupTag

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetGroupTag() string`

GetGroupTag returns the GroupTag field if non-nil, zero value otherwise.

### GetGroupTagOk

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetGroupTagOk() (*string, bool)`

GetGroupTagOk returns a tuple with the GroupTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupTag

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetGroupTag(v string)`

SetGroupTag sets GroupTag field to given value.

### HasGroupTag

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) HasGroupTag() bool`

HasGroupTag returns a boolean if a field has been set.

### SetGroupTagNil

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetGroupTagNil(b bool)`

 SetGroupTagNil sets the value for GroupTag to be an explicit nil

### UnsetGroupTag
`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) UnsetGroupTag()`

UnsetGroupTag ensures that no value is present for GroupTag, not even an explicit nil
### GetLastContactedDateTime

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetLastContactedDateTime() time.Time`

GetLastContactedDateTime returns the LastContactedDateTime field if non-nil, zero value otherwise.

### GetLastContactedDateTimeOk

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetLastContactedDateTimeOk() (*time.Time, bool)`

GetLastContactedDateTimeOk returns a tuple with the LastContactedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastContactedDateTime

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetLastContactedDateTime(v time.Time)`

SetLastContactedDateTime sets LastContactedDateTime field to given value.

### HasLastContactedDateTime

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) HasLastContactedDateTime() bool`

HasLastContactedDateTime returns a boolean if a field has been set.

### GetManagedDeviceId

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetManagedDeviceId() string`

GetManagedDeviceId returns the ManagedDeviceId field if non-nil, zero value otherwise.

### GetManagedDeviceIdOk

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetManagedDeviceIdOk() (*string, bool)`

GetManagedDeviceIdOk returns a tuple with the ManagedDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedDeviceId

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetManagedDeviceId(v string)`

SetManagedDeviceId sets ManagedDeviceId field to given value.

### HasManagedDeviceId

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) HasManagedDeviceId() bool`

HasManagedDeviceId returns a boolean if a field has been set.

### SetManagedDeviceIdNil

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetManagedDeviceIdNil(b bool)`

 SetManagedDeviceIdNil sets the value for ManagedDeviceId to be an explicit nil

### UnsetManagedDeviceId
`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) UnsetManagedDeviceId()`

UnsetManagedDeviceId ensures that no value is present for ManagedDeviceId, not even an explicit nil
### GetManufacturer

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetModel

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetProductKey

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetProductKey() string`

GetProductKey returns the ProductKey field if non-nil, zero value otherwise.

### GetProductKeyOk

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetProductKeyOk() (*string, bool)`

GetProductKeyOk returns a tuple with the ProductKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductKey

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetProductKey(v string)`

SetProductKey sets ProductKey field to given value.

### HasProductKey

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) HasProductKey() bool`

HasProductKey returns a boolean if a field has been set.

### SetProductKeyNil

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetProductKeyNil(b bool)`

 SetProductKeyNil sets the value for ProductKey to be an explicit nil

### UnsetProductKey
`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) UnsetProductKey()`

UnsetProductKey ensures that no value is present for ProductKey, not even an explicit nil
### GetPurchaseOrderIdentifier

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetPurchaseOrderIdentifier() string`

GetPurchaseOrderIdentifier returns the PurchaseOrderIdentifier field if non-nil, zero value otherwise.

### GetPurchaseOrderIdentifierOk

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetPurchaseOrderIdentifierOk() (*string, bool)`

GetPurchaseOrderIdentifierOk returns a tuple with the PurchaseOrderIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderIdentifier

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetPurchaseOrderIdentifier(v string)`

SetPurchaseOrderIdentifier sets PurchaseOrderIdentifier field to given value.

### HasPurchaseOrderIdentifier

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) HasPurchaseOrderIdentifier() bool`

HasPurchaseOrderIdentifier returns a boolean if a field has been set.

### SetPurchaseOrderIdentifierNil

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetPurchaseOrderIdentifierNil(b bool)`

 SetPurchaseOrderIdentifierNil sets the value for PurchaseOrderIdentifier to be an explicit nil

### UnsetPurchaseOrderIdentifier
`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) UnsetPurchaseOrderIdentifier()`

UnsetPurchaseOrderIdentifier ensures that no value is present for PurchaseOrderIdentifier, not even an explicit nil
### GetResourceName

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### SetResourceNameNil

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
### GetSerialNumber

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetSkuNumber

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetSkuNumber() string`

GetSkuNumber returns the SkuNumber field if non-nil, zero value otherwise.

### GetSkuNumberOk

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetSkuNumberOk() (*string, bool)`

GetSkuNumberOk returns a tuple with the SkuNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuNumber

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetSkuNumber(v string)`

SetSkuNumber sets SkuNumber field to given value.

### HasSkuNumber

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) HasSkuNumber() bool`

HasSkuNumber returns a boolean if a field has been set.

### SetSkuNumberNil

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetSkuNumberNil(b bool)`

 SetSkuNumberNil sets the value for SkuNumber to be an explicit nil

### UnsetSkuNumber
`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) UnsetSkuNumber()`

UnsetSkuNumber ensures that no value is present for SkuNumber, not even an explicit nil
### GetSystemFamily

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetSystemFamily() string`

GetSystemFamily returns the SystemFamily field if non-nil, zero value otherwise.

### GetSystemFamilyOk

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetSystemFamilyOk() (*string, bool)`

GetSystemFamilyOk returns a tuple with the SystemFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemFamily

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetSystemFamily(v string)`

SetSystemFamily sets SystemFamily field to given value.

### HasSystemFamily

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) HasSystemFamily() bool`

HasSystemFamily returns a boolean if a field has been set.

### SetSystemFamilyNil

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetSystemFamilyNil(b bool)`

 SetSystemFamilyNil sets the value for SystemFamily to be an explicit nil

### UnsetSystemFamily
`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) UnsetSystemFamily()`

UnsetSystemFamily ensures that no value is present for SystemFamily, not even an explicit nil
### GetUserPrincipalName

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### SetUserPrincipalNameNil

`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) SetUserPrincipalNameNil(b bool)`

 SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil

### UnsetUserPrincipalName
`func (o *MicrosoftGraphWindowsAutopilotDeviceIdentity) UnsetUserPrincipalName()`

UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


