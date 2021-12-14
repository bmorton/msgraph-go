# MicrosoftGraphDeviceManagementPartner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DisplayName** | Pointer to **NullableString** | Partner display name | [optional] 
**IsConfigured** | Pointer to **bool** | Whether device management partner is configured or not | [optional] 
**LastHeartbeatDateTime** | Pointer to **time.Time** | Timestamp of last heartbeat after admin enabled option Connect to Device management Partner | [optional] 
**PartnerAppType** | Pointer to [**AnyOfmicrosoftGraphDeviceManagementPartnerAppType**](anyOf&lt;microsoft.graph.deviceManagementPartnerAppType&gt;.md) | Partner App type. Possible values are: unknown, singleTenantApp, multiTenantApp. | [optional] 
**PartnerState** | Pointer to [**AnyOfmicrosoftGraphDeviceManagementPartnerTenantState**](anyOf&lt;microsoft.graph.deviceManagementPartnerTenantState&gt;.md) | Partner state of this tenant. Possible values are: unknown, unavailable, enabled, terminated, rejected, unresponsive. | [optional] 
**SingleTenantAppId** | Pointer to **NullableString** | Partner Single tenant App id | [optional] 
**WhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime** | Pointer to **NullableTime** | DateTime in UTC when PartnerDevices will be marked as NonCompliant | [optional] 
**WhenPartnerDevicesWillBeRemovedDateTime** | Pointer to **NullableTime** | DateTime in UTC when PartnerDevices will be removed | [optional] 

## Methods

### NewMicrosoftGraphDeviceManagementPartner

`func NewMicrosoftGraphDeviceManagementPartner() *MicrosoftGraphDeviceManagementPartner`

NewMicrosoftGraphDeviceManagementPartner instantiates a new MicrosoftGraphDeviceManagementPartner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDeviceManagementPartnerWithDefaults

`func NewMicrosoftGraphDeviceManagementPartnerWithDefaults() *MicrosoftGraphDeviceManagementPartner`

NewMicrosoftGraphDeviceManagementPartnerWithDefaults instantiates a new MicrosoftGraphDeviceManagementPartner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDeviceManagementPartner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceManagementPartner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDeviceManagementPartner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDeviceManagementPartner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *MicrosoftGraphDeviceManagementPartner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphDeviceManagementPartner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphDeviceManagementPartner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphDeviceManagementPartner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphDeviceManagementPartner) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphDeviceManagementPartner) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetIsConfigured

`func (o *MicrosoftGraphDeviceManagementPartner) GetIsConfigured() bool`

GetIsConfigured returns the IsConfigured field if non-nil, zero value otherwise.

### GetIsConfiguredOk

`func (o *MicrosoftGraphDeviceManagementPartner) GetIsConfiguredOk() (*bool, bool)`

GetIsConfiguredOk returns a tuple with the IsConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfigured

`func (o *MicrosoftGraphDeviceManagementPartner) SetIsConfigured(v bool)`

SetIsConfigured sets IsConfigured field to given value.

### HasIsConfigured

`func (o *MicrosoftGraphDeviceManagementPartner) HasIsConfigured() bool`

HasIsConfigured returns a boolean if a field has been set.

### GetLastHeartbeatDateTime

`func (o *MicrosoftGraphDeviceManagementPartner) GetLastHeartbeatDateTime() time.Time`

GetLastHeartbeatDateTime returns the LastHeartbeatDateTime field if non-nil, zero value otherwise.

### GetLastHeartbeatDateTimeOk

`func (o *MicrosoftGraphDeviceManagementPartner) GetLastHeartbeatDateTimeOk() (*time.Time, bool)`

GetLastHeartbeatDateTimeOk returns a tuple with the LastHeartbeatDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeatDateTime

`func (o *MicrosoftGraphDeviceManagementPartner) SetLastHeartbeatDateTime(v time.Time)`

SetLastHeartbeatDateTime sets LastHeartbeatDateTime field to given value.

### HasLastHeartbeatDateTime

`func (o *MicrosoftGraphDeviceManagementPartner) HasLastHeartbeatDateTime() bool`

HasLastHeartbeatDateTime returns a boolean if a field has been set.

### GetPartnerAppType

`func (o *MicrosoftGraphDeviceManagementPartner) GetPartnerAppType() AnyOfmicrosoftGraphDeviceManagementPartnerAppType`

GetPartnerAppType returns the PartnerAppType field if non-nil, zero value otherwise.

### GetPartnerAppTypeOk

`func (o *MicrosoftGraphDeviceManagementPartner) GetPartnerAppTypeOk() (*AnyOfmicrosoftGraphDeviceManagementPartnerAppType, bool)`

GetPartnerAppTypeOk returns a tuple with the PartnerAppType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerAppType

`func (o *MicrosoftGraphDeviceManagementPartner) SetPartnerAppType(v AnyOfmicrosoftGraphDeviceManagementPartnerAppType)`

SetPartnerAppType sets PartnerAppType field to given value.

### HasPartnerAppType

`func (o *MicrosoftGraphDeviceManagementPartner) HasPartnerAppType() bool`

HasPartnerAppType returns a boolean if a field has been set.

### SetPartnerAppTypeNil

`func (o *MicrosoftGraphDeviceManagementPartner) SetPartnerAppTypeNil(b bool)`

 SetPartnerAppTypeNil sets the value for PartnerAppType to be an explicit nil

### UnsetPartnerAppType
`func (o *MicrosoftGraphDeviceManagementPartner) UnsetPartnerAppType()`

UnsetPartnerAppType ensures that no value is present for PartnerAppType, not even an explicit nil
### GetPartnerState

`func (o *MicrosoftGraphDeviceManagementPartner) GetPartnerState() AnyOfmicrosoftGraphDeviceManagementPartnerTenantState`

GetPartnerState returns the PartnerState field if non-nil, zero value otherwise.

### GetPartnerStateOk

`func (o *MicrosoftGraphDeviceManagementPartner) GetPartnerStateOk() (*AnyOfmicrosoftGraphDeviceManagementPartnerTenantState, bool)`

GetPartnerStateOk returns a tuple with the PartnerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerState

`func (o *MicrosoftGraphDeviceManagementPartner) SetPartnerState(v AnyOfmicrosoftGraphDeviceManagementPartnerTenantState)`

SetPartnerState sets PartnerState field to given value.

### HasPartnerState

`func (o *MicrosoftGraphDeviceManagementPartner) HasPartnerState() bool`

HasPartnerState returns a boolean if a field has been set.

### SetPartnerStateNil

`func (o *MicrosoftGraphDeviceManagementPartner) SetPartnerStateNil(b bool)`

 SetPartnerStateNil sets the value for PartnerState to be an explicit nil

### UnsetPartnerState
`func (o *MicrosoftGraphDeviceManagementPartner) UnsetPartnerState()`

UnsetPartnerState ensures that no value is present for PartnerState, not even an explicit nil
### GetSingleTenantAppId

`func (o *MicrosoftGraphDeviceManagementPartner) GetSingleTenantAppId() string`

GetSingleTenantAppId returns the SingleTenantAppId field if non-nil, zero value otherwise.

### GetSingleTenantAppIdOk

`func (o *MicrosoftGraphDeviceManagementPartner) GetSingleTenantAppIdOk() (*string, bool)`

GetSingleTenantAppIdOk returns a tuple with the SingleTenantAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleTenantAppId

`func (o *MicrosoftGraphDeviceManagementPartner) SetSingleTenantAppId(v string)`

SetSingleTenantAppId sets SingleTenantAppId field to given value.

### HasSingleTenantAppId

`func (o *MicrosoftGraphDeviceManagementPartner) HasSingleTenantAppId() bool`

HasSingleTenantAppId returns a boolean if a field has been set.

### SetSingleTenantAppIdNil

`func (o *MicrosoftGraphDeviceManagementPartner) SetSingleTenantAppIdNil(b bool)`

 SetSingleTenantAppIdNil sets the value for SingleTenantAppId to be an explicit nil

### UnsetSingleTenantAppId
`func (o *MicrosoftGraphDeviceManagementPartner) UnsetSingleTenantAppId()`

UnsetSingleTenantAppId ensures that no value is present for SingleTenantAppId, not even an explicit nil
### GetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime

`func (o *MicrosoftGraphDeviceManagementPartner) GetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime() time.Time`

GetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime returns the WhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime field if non-nil, zero value otherwise.

### GetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTimeOk

`func (o *MicrosoftGraphDeviceManagementPartner) GetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTimeOk() (*time.Time, bool)`

GetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTimeOk returns a tuple with the WhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime

`func (o *MicrosoftGraphDeviceManagementPartner) SetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime(v time.Time)`

SetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime sets WhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime field to given value.

### HasWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime

`func (o *MicrosoftGraphDeviceManagementPartner) HasWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime() bool`

HasWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime returns a boolean if a field has been set.

### SetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTimeNil

`func (o *MicrosoftGraphDeviceManagementPartner) SetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTimeNil(b bool)`

 SetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTimeNil sets the value for WhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime to be an explicit nil

### UnsetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime
`func (o *MicrosoftGraphDeviceManagementPartner) UnsetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime()`

UnsetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime ensures that no value is present for WhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime, not even an explicit nil
### GetWhenPartnerDevicesWillBeRemovedDateTime

`func (o *MicrosoftGraphDeviceManagementPartner) GetWhenPartnerDevicesWillBeRemovedDateTime() time.Time`

GetWhenPartnerDevicesWillBeRemovedDateTime returns the WhenPartnerDevicesWillBeRemovedDateTime field if non-nil, zero value otherwise.

### GetWhenPartnerDevicesWillBeRemovedDateTimeOk

`func (o *MicrosoftGraphDeviceManagementPartner) GetWhenPartnerDevicesWillBeRemovedDateTimeOk() (*time.Time, bool)`

GetWhenPartnerDevicesWillBeRemovedDateTimeOk returns a tuple with the WhenPartnerDevicesWillBeRemovedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhenPartnerDevicesWillBeRemovedDateTime

`func (o *MicrosoftGraphDeviceManagementPartner) SetWhenPartnerDevicesWillBeRemovedDateTime(v time.Time)`

SetWhenPartnerDevicesWillBeRemovedDateTime sets WhenPartnerDevicesWillBeRemovedDateTime field to given value.

### HasWhenPartnerDevicesWillBeRemovedDateTime

`func (o *MicrosoftGraphDeviceManagementPartner) HasWhenPartnerDevicesWillBeRemovedDateTime() bool`

HasWhenPartnerDevicesWillBeRemovedDateTime returns a boolean if a field has been set.

### SetWhenPartnerDevicesWillBeRemovedDateTimeNil

`func (o *MicrosoftGraphDeviceManagementPartner) SetWhenPartnerDevicesWillBeRemovedDateTimeNil(b bool)`

 SetWhenPartnerDevicesWillBeRemovedDateTimeNil sets the value for WhenPartnerDevicesWillBeRemovedDateTime to be an explicit nil

### UnsetWhenPartnerDevicesWillBeRemovedDateTime
`func (o *MicrosoftGraphDeviceManagementPartner) UnsetWhenPartnerDevicesWillBeRemovedDateTime()`

UnsetWhenPartnerDevicesWillBeRemovedDateTime ensures that no value is present for WhenPartnerDevicesWillBeRemovedDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


