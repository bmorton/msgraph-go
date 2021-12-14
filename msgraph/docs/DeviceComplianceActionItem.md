# DeviceComplianceActionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionType** | Pointer to [**AnyOfmicrosoftGraphDeviceComplianceActionType**](anyOf&lt;microsoft.graph.deviceComplianceActionType&gt;.md) | What action to take. Possible values are: noAction, notification, block, retire, wipe, removeResourceAccessProfiles, pushNotification. | [optional] 
**GracePeriodHours** | Pointer to **int32** | Number of hours to wait till the action will be enforced. Valid values 0 to 8760 | [optional] 
**NotificationMessageCCList** | Pointer to **[]string** | A list of group IDs to speicify who to CC this notification message to. | [optional] 
**NotificationTemplateId** | Pointer to **NullableString** | What notification Message template to use | [optional] 

## Methods

### NewDeviceComplianceActionItem

`func NewDeviceComplianceActionItem() *DeviceComplianceActionItem`

NewDeviceComplianceActionItem instantiates a new DeviceComplianceActionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceComplianceActionItemWithDefaults

`func NewDeviceComplianceActionItemWithDefaults() *DeviceComplianceActionItem`

NewDeviceComplianceActionItemWithDefaults instantiates a new DeviceComplianceActionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionType

`func (o *DeviceComplianceActionItem) GetActionType() AnyOfmicrosoftGraphDeviceComplianceActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *DeviceComplianceActionItem) GetActionTypeOk() (*AnyOfmicrosoftGraphDeviceComplianceActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *DeviceComplianceActionItem) SetActionType(v AnyOfmicrosoftGraphDeviceComplianceActionType)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *DeviceComplianceActionItem) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### SetActionTypeNil

`func (o *DeviceComplianceActionItem) SetActionTypeNil(b bool)`

 SetActionTypeNil sets the value for ActionType to be an explicit nil

### UnsetActionType
`func (o *DeviceComplianceActionItem) UnsetActionType()`

UnsetActionType ensures that no value is present for ActionType, not even an explicit nil
### GetGracePeriodHours

`func (o *DeviceComplianceActionItem) GetGracePeriodHours() int32`

GetGracePeriodHours returns the GracePeriodHours field if non-nil, zero value otherwise.

### GetGracePeriodHoursOk

`func (o *DeviceComplianceActionItem) GetGracePeriodHoursOk() (*int32, bool)`

GetGracePeriodHoursOk returns a tuple with the GracePeriodHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodHours

`func (o *DeviceComplianceActionItem) SetGracePeriodHours(v int32)`

SetGracePeriodHours sets GracePeriodHours field to given value.

### HasGracePeriodHours

`func (o *DeviceComplianceActionItem) HasGracePeriodHours() bool`

HasGracePeriodHours returns a boolean if a field has been set.

### GetNotificationMessageCCList

`func (o *DeviceComplianceActionItem) GetNotificationMessageCCList() []*string`

GetNotificationMessageCCList returns the NotificationMessageCCList field if non-nil, zero value otherwise.

### GetNotificationMessageCCListOk

`func (o *DeviceComplianceActionItem) GetNotificationMessageCCListOk() (*[]*string, bool)`

GetNotificationMessageCCListOk returns a tuple with the NotificationMessageCCList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationMessageCCList

`func (o *DeviceComplianceActionItem) SetNotificationMessageCCList(v []*string)`

SetNotificationMessageCCList sets NotificationMessageCCList field to given value.

### HasNotificationMessageCCList

`func (o *DeviceComplianceActionItem) HasNotificationMessageCCList() bool`

HasNotificationMessageCCList returns a boolean if a field has been set.

### GetNotificationTemplateId

`func (o *DeviceComplianceActionItem) GetNotificationTemplateId() string`

GetNotificationTemplateId returns the NotificationTemplateId field if non-nil, zero value otherwise.

### GetNotificationTemplateIdOk

`func (o *DeviceComplianceActionItem) GetNotificationTemplateIdOk() (*string, bool)`

GetNotificationTemplateIdOk returns a tuple with the NotificationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTemplateId

`func (o *DeviceComplianceActionItem) SetNotificationTemplateId(v string)`

SetNotificationTemplateId sets NotificationTemplateId field to given value.

### HasNotificationTemplateId

`func (o *DeviceComplianceActionItem) HasNotificationTemplateId() bool`

HasNotificationTemplateId returns a boolean if a field has been set.

### SetNotificationTemplateIdNil

`func (o *DeviceComplianceActionItem) SetNotificationTemplateIdNil(b bool)`

 SetNotificationTemplateIdNil sets the value for NotificationTemplateId to be an explicit nil

### UnsetNotificationTemplateId
`func (o *DeviceComplianceActionItem) UnsetNotificationTemplateId()`

UnsetNotificationTemplateId ensures that no value is present for NotificationTemplateId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


