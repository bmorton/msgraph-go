# MicrosoftGraphDeviceComplianceActionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ActionType** | Pointer to [**AnyOfmicrosoftGraphDeviceComplianceActionType**](anyOf&lt;microsoft.graph.deviceComplianceActionType&gt;.md) | What action to take. Possible values are: noAction, notification, block, retire, wipe, removeResourceAccessProfiles, pushNotification. | [optional] 
**GracePeriodHours** | Pointer to **int32** | Number of hours to wait till the action will be enforced. Valid values 0 to 8760 | [optional] 
**NotificationMessageCCList** | Pointer to **[]string** | A list of group IDs to speicify who to CC this notification message to. | [optional] 
**NotificationTemplateId** | Pointer to **NullableString** | What notification Message template to use | [optional] 

## Methods

### NewMicrosoftGraphDeviceComplianceActionItem

`func NewMicrosoftGraphDeviceComplianceActionItem() *MicrosoftGraphDeviceComplianceActionItem`

NewMicrosoftGraphDeviceComplianceActionItem instantiates a new MicrosoftGraphDeviceComplianceActionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDeviceComplianceActionItemWithDefaults

`func NewMicrosoftGraphDeviceComplianceActionItemWithDefaults() *MicrosoftGraphDeviceComplianceActionItem`

NewMicrosoftGraphDeviceComplianceActionItemWithDefaults instantiates a new MicrosoftGraphDeviceComplianceActionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDeviceComplianceActionItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceComplianceActionItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDeviceComplianceActionItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDeviceComplianceActionItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActionType

`func (o *MicrosoftGraphDeviceComplianceActionItem) GetActionType() AnyOfmicrosoftGraphDeviceComplianceActionType`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *MicrosoftGraphDeviceComplianceActionItem) GetActionTypeOk() (*AnyOfmicrosoftGraphDeviceComplianceActionType, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *MicrosoftGraphDeviceComplianceActionItem) SetActionType(v AnyOfmicrosoftGraphDeviceComplianceActionType)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *MicrosoftGraphDeviceComplianceActionItem) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### SetActionTypeNil

`func (o *MicrosoftGraphDeviceComplianceActionItem) SetActionTypeNil(b bool)`

 SetActionTypeNil sets the value for ActionType to be an explicit nil

### UnsetActionType
`func (o *MicrosoftGraphDeviceComplianceActionItem) UnsetActionType()`

UnsetActionType ensures that no value is present for ActionType, not even an explicit nil
### GetGracePeriodHours

`func (o *MicrosoftGraphDeviceComplianceActionItem) GetGracePeriodHours() int32`

GetGracePeriodHours returns the GracePeriodHours field if non-nil, zero value otherwise.

### GetGracePeriodHoursOk

`func (o *MicrosoftGraphDeviceComplianceActionItem) GetGracePeriodHoursOk() (*int32, bool)`

GetGracePeriodHoursOk returns a tuple with the GracePeriodHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriodHours

`func (o *MicrosoftGraphDeviceComplianceActionItem) SetGracePeriodHours(v int32)`

SetGracePeriodHours sets GracePeriodHours field to given value.

### HasGracePeriodHours

`func (o *MicrosoftGraphDeviceComplianceActionItem) HasGracePeriodHours() bool`

HasGracePeriodHours returns a boolean if a field has been set.

### GetNotificationMessageCCList

`func (o *MicrosoftGraphDeviceComplianceActionItem) GetNotificationMessageCCList() []*string`

GetNotificationMessageCCList returns the NotificationMessageCCList field if non-nil, zero value otherwise.

### GetNotificationMessageCCListOk

`func (o *MicrosoftGraphDeviceComplianceActionItem) GetNotificationMessageCCListOk() (*[]*string, bool)`

GetNotificationMessageCCListOk returns a tuple with the NotificationMessageCCList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationMessageCCList

`func (o *MicrosoftGraphDeviceComplianceActionItem) SetNotificationMessageCCList(v []*string)`

SetNotificationMessageCCList sets NotificationMessageCCList field to given value.

### HasNotificationMessageCCList

`func (o *MicrosoftGraphDeviceComplianceActionItem) HasNotificationMessageCCList() bool`

HasNotificationMessageCCList returns a boolean if a field has been set.

### GetNotificationTemplateId

`func (o *MicrosoftGraphDeviceComplianceActionItem) GetNotificationTemplateId() string`

GetNotificationTemplateId returns the NotificationTemplateId field if non-nil, zero value otherwise.

### GetNotificationTemplateIdOk

`func (o *MicrosoftGraphDeviceComplianceActionItem) GetNotificationTemplateIdOk() (*string, bool)`

GetNotificationTemplateIdOk returns a tuple with the NotificationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTemplateId

`func (o *MicrosoftGraphDeviceComplianceActionItem) SetNotificationTemplateId(v string)`

SetNotificationTemplateId sets NotificationTemplateId field to given value.

### HasNotificationTemplateId

`func (o *MicrosoftGraphDeviceComplianceActionItem) HasNotificationTemplateId() bool`

HasNotificationTemplateId returns a boolean if a field has been set.

### SetNotificationTemplateIdNil

`func (o *MicrosoftGraphDeviceComplianceActionItem) SetNotificationTemplateIdNil(b bool)`

 SetNotificationTemplateIdNil sets the value for NotificationTemplateId to be an explicit nil

### UnsetNotificationTemplateId
`func (o *MicrosoftGraphDeviceComplianceActionItem) UnsetNotificationTemplateId()`

UnsetNotificationTemplateId ensures that no value is present for NotificationTemplateId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


