# ItemActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | An item was accessed. | [optional] 
**ActivityDateTime** | Pointer to **NullableTime** | Details about when the activity took place. Read-only. | [optional] 
**Actor** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | Identity of who performed the action. Read-only. | [optional] 
**DriveItem** | Pointer to [**AnyOfmicrosoftGraphDriveItem**](anyOf&lt;microsoft.graph.driveItem&gt;.md) | Exposes the driveItem that was the target of this activity. | [optional] 

## Methods

### NewItemActivity

`func NewItemActivity() *ItemActivity`

NewItemActivity instantiates a new ItemActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemActivityWithDefaults

`func NewItemActivityWithDefaults() *ItemActivity`

NewItemActivityWithDefaults instantiates a new ItemActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *ItemActivity) GetAccess() AnyOfobject`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ItemActivity) GetAccessOk() (*AnyOfobject, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ItemActivity) SetAccess(v AnyOfobject)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ItemActivity) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### SetAccessNil

`func (o *ItemActivity) SetAccessNil(b bool)`

 SetAccessNil sets the value for Access to be an explicit nil

### UnsetAccess
`func (o *ItemActivity) UnsetAccess()`

UnsetAccess ensures that no value is present for Access, not even an explicit nil
### GetActivityDateTime

`func (o *ItemActivity) GetActivityDateTime() time.Time`

GetActivityDateTime returns the ActivityDateTime field if non-nil, zero value otherwise.

### GetActivityDateTimeOk

`func (o *ItemActivity) GetActivityDateTimeOk() (*time.Time, bool)`

GetActivityDateTimeOk returns a tuple with the ActivityDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityDateTime

`func (o *ItemActivity) SetActivityDateTime(v time.Time)`

SetActivityDateTime sets ActivityDateTime field to given value.

### HasActivityDateTime

`func (o *ItemActivity) HasActivityDateTime() bool`

HasActivityDateTime returns a boolean if a field has been set.

### SetActivityDateTimeNil

`func (o *ItemActivity) SetActivityDateTimeNil(b bool)`

 SetActivityDateTimeNil sets the value for ActivityDateTime to be an explicit nil

### UnsetActivityDateTime
`func (o *ItemActivity) UnsetActivityDateTime()`

UnsetActivityDateTime ensures that no value is present for ActivityDateTime, not even an explicit nil
### GetActor

`func (o *ItemActivity) GetActor() AnyOfmicrosoftGraphIdentitySet`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *ItemActivity) GetActorOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *ItemActivity) SetActor(v AnyOfmicrosoftGraphIdentitySet)`

SetActor sets Actor field to given value.

### HasActor

`func (o *ItemActivity) HasActor() bool`

HasActor returns a boolean if a field has been set.

### SetActorNil

`func (o *ItemActivity) SetActorNil(b bool)`

 SetActorNil sets the value for Actor to be an explicit nil

### UnsetActor
`func (o *ItemActivity) UnsetActor()`

UnsetActor ensures that no value is present for Actor, not even an explicit nil
### GetDriveItem

`func (o *ItemActivity) GetDriveItem() AnyOfmicrosoftGraphDriveItem`

GetDriveItem returns the DriveItem field if non-nil, zero value otherwise.

### GetDriveItemOk

`func (o *ItemActivity) GetDriveItemOk() (*AnyOfmicrosoftGraphDriveItem, bool)`

GetDriveItemOk returns a tuple with the DriveItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveItem

`func (o *ItemActivity) SetDriveItem(v AnyOfmicrosoftGraphDriveItem)`

SetDriveItem sets DriveItem field to given value.

### HasDriveItem

`func (o *ItemActivity) HasDriveItem() bool`

HasDriveItem returns a boolean if a field has been set.

### SetDriveItemNil

`func (o *ItemActivity) SetDriveItemNil(b bool)`

 SetDriveItemNil sets the value for DriveItem to be an explicit nil

### UnsetDriveItem
`func (o *ItemActivity) UnsetDriveItem()`

UnsetDriveItem ensures that no value is present for DriveItem, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


