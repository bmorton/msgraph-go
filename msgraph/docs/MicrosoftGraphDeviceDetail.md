# MicrosoftGraphDeviceDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Browser** | Pointer to **NullableString** | Indicates the browser information of the used for signing in. | [optional] 
**DeviceId** | Pointer to **NullableString** | Refers to the UniqueID of the device used for signing in. | [optional] 
**DisplayName** | Pointer to **NullableString** | Refers to the name of the device used for signing in. | [optional] 
**IsCompliant** | Pointer to **NullableBool** | Indicates whether the device is compliant. | [optional] 
**IsManaged** | Pointer to **NullableBool** | Indicates whether the device is managed. | [optional] 
**OperatingSystem** | Pointer to **NullableString** | Indicates the operating system name and version used for signing in. | [optional] 
**TrustType** | Pointer to **NullableString** | Provides information about whether the signed-in device is Workplace Joined, AzureAD Joined, Domain Joined. | [optional] 

## Methods

### NewMicrosoftGraphDeviceDetail

`func NewMicrosoftGraphDeviceDetail() *MicrosoftGraphDeviceDetail`

NewMicrosoftGraphDeviceDetail instantiates a new MicrosoftGraphDeviceDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDeviceDetailWithDefaults

`func NewMicrosoftGraphDeviceDetailWithDefaults() *MicrosoftGraphDeviceDetail`

NewMicrosoftGraphDeviceDetailWithDefaults instantiates a new MicrosoftGraphDeviceDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrowser

`func (o *MicrosoftGraphDeviceDetail) GetBrowser() string`

GetBrowser returns the Browser field if non-nil, zero value otherwise.

### GetBrowserOk

`func (o *MicrosoftGraphDeviceDetail) GetBrowserOk() (*string, bool)`

GetBrowserOk returns a tuple with the Browser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowser

`func (o *MicrosoftGraphDeviceDetail) SetBrowser(v string)`

SetBrowser sets Browser field to given value.

### HasBrowser

`func (o *MicrosoftGraphDeviceDetail) HasBrowser() bool`

HasBrowser returns a boolean if a field has been set.

### SetBrowserNil

`func (o *MicrosoftGraphDeviceDetail) SetBrowserNil(b bool)`

 SetBrowserNil sets the value for Browser to be an explicit nil

### UnsetBrowser
`func (o *MicrosoftGraphDeviceDetail) UnsetBrowser()`

UnsetBrowser ensures that no value is present for Browser, not even an explicit nil
### GetDeviceId

`func (o *MicrosoftGraphDeviceDetail) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *MicrosoftGraphDeviceDetail) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *MicrosoftGraphDeviceDetail) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *MicrosoftGraphDeviceDetail) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *MicrosoftGraphDeviceDetail) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *MicrosoftGraphDeviceDetail) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphDeviceDetail) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphDeviceDetail) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphDeviceDetail) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphDeviceDetail) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphDeviceDetail) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphDeviceDetail) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetIsCompliant

`func (o *MicrosoftGraphDeviceDetail) GetIsCompliant() bool`

GetIsCompliant returns the IsCompliant field if non-nil, zero value otherwise.

### GetIsCompliantOk

`func (o *MicrosoftGraphDeviceDetail) GetIsCompliantOk() (*bool, bool)`

GetIsCompliantOk returns a tuple with the IsCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompliant

`func (o *MicrosoftGraphDeviceDetail) SetIsCompliant(v bool)`

SetIsCompliant sets IsCompliant field to given value.

### HasIsCompliant

`func (o *MicrosoftGraphDeviceDetail) HasIsCompliant() bool`

HasIsCompliant returns a boolean if a field has been set.

### SetIsCompliantNil

`func (o *MicrosoftGraphDeviceDetail) SetIsCompliantNil(b bool)`

 SetIsCompliantNil sets the value for IsCompliant to be an explicit nil

### UnsetIsCompliant
`func (o *MicrosoftGraphDeviceDetail) UnsetIsCompliant()`

UnsetIsCompliant ensures that no value is present for IsCompliant, not even an explicit nil
### GetIsManaged

`func (o *MicrosoftGraphDeviceDetail) GetIsManaged() bool`

GetIsManaged returns the IsManaged field if non-nil, zero value otherwise.

### GetIsManagedOk

`func (o *MicrosoftGraphDeviceDetail) GetIsManagedOk() (*bool, bool)`

GetIsManagedOk returns a tuple with the IsManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManaged

`func (o *MicrosoftGraphDeviceDetail) SetIsManaged(v bool)`

SetIsManaged sets IsManaged field to given value.

### HasIsManaged

`func (o *MicrosoftGraphDeviceDetail) HasIsManaged() bool`

HasIsManaged returns a boolean if a field has been set.

### SetIsManagedNil

`func (o *MicrosoftGraphDeviceDetail) SetIsManagedNil(b bool)`

 SetIsManagedNil sets the value for IsManaged to be an explicit nil

### UnsetIsManaged
`func (o *MicrosoftGraphDeviceDetail) UnsetIsManaged()`

UnsetIsManaged ensures that no value is present for IsManaged, not even an explicit nil
### GetOperatingSystem

`func (o *MicrosoftGraphDeviceDetail) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *MicrosoftGraphDeviceDetail) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *MicrosoftGraphDeviceDetail) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *MicrosoftGraphDeviceDetail) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### SetOperatingSystemNil

`func (o *MicrosoftGraphDeviceDetail) SetOperatingSystemNil(b bool)`

 SetOperatingSystemNil sets the value for OperatingSystem to be an explicit nil

### UnsetOperatingSystem
`func (o *MicrosoftGraphDeviceDetail) UnsetOperatingSystem()`

UnsetOperatingSystem ensures that no value is present for OperatingSystem, not even an explicit nil
### GetTrustType

`func (o *MicrosoftGraphDeviceDetail) GetTrustType() string`

GetTrustType returns the TrustType field if non-nil, zero value otherwise.

### GetTrustTypeOk

`func (o *MicrosoftGraphDeviceDetail) GetTrustTypeOk() (*string, bool)`

GetTrustTypeOk returns a tuple with the TrustType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustType

`func (o *MicrosoftGraphDeviceDetail) SetTrustType(v string)`

SetTrustType sets TrustType field to given value.

### HasTrustType

`func (o *MicrosoftGraphDeviceDetail) HasTrustType() bool`

HasTrustType returns a boolean if a field has been set.

### SetTrustTypeNil

`func (o *MicrosoftGraphDeviceDetail) SetTrustTypeNil(b bool)`

 SetTrustTypeNil sets the value for TrustType to be an explicit nil

### UnsetTrustType
`func (o *MicrosoftGraphDeviceDetail) UnsetTrustType()`

UnsetTrustType ensures that no value is present for TrustType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


