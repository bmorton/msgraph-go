# MicrosoftGraphItemActivityStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Access** | Pointer to [**AnyOfmicrosoftGraphItemActionStat**](anyOf&lt;microsoft.graph.itemActionStat&gt;.md) | Statistics about the access actions in this interval. Read-only. | [optional] 
**Create** | Pointer to [**AnyOfmicrosoftGraphItemActionStat**](anyOf&lt;microsoft.graph.itemActionStat&gt;.md) | Statistics about the create actions in this interval. Read-only. | [optional] 
**Delete** | Pointer to [**AnyOfmicrosoftGraphItemActionStat**](anyOf&lt;microsoft.graph.itemActionStat&gt;.md) | Statistics about the delete actions in this interval. Read-only. | [optional] 
**Edit** | Pointer to [**AnyOfmicrosoftGraphItemActionStat**](anyOf&lt;microsoft.graph.itemActionStat&gt;.md) | Statistics about the edit actions in this interval. Read-only. | [optional] 
**EndDateTime** | Pointer to **NullableTime** | When the interval ends. Read-only. | [optional] 
**IncompleteData** | Pointer to [**AnyOfmicrosoftGraphIncompleteData**](anyOf&lt;microsoft.graph.incompleteData&gt;.md) | Indicates that the statistics in this interval are based on incomplete data. Read-only. | [optional] 
**IsTrending** | Pointer to **NullableBool** | Indicates whether the item is &#39;trending.&#39; Read-only. | [optional] 
**Move** | Pointer to [**AnyOfmicrosoftGraphItemActionStat**](anyOf&lt;microsoft.graph.itemActionStat&gt;.md) | Statistics about the move actions in this interval. Read-only. | [optional] 
**StartDateTime** | Pointer to **NullableTime** | When the interval starts. Read-only. | [optional] 
**Activities** | Pointer to [**[]MicrosoftGraphItemActivity**](MicrosoftGraphItemActivity.md) | Exposes the itemActivities represented in this itemActivityStat resource. | [optional] 

## Methods

### NewMicrosoftGraphItemActivityStat

`func NewMicrosoftGraphItemActivityStat() *MicrosoftGraphItemActivityStat`

NewMicrosoftGraphItemActivityStat instantiates a new MicrosoftGraphItemActivityStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphItemActivityStatWithDefaults

`func NewMicrosoftGraphItemActivityStatWithDefaults() *MicrosoftGraphItemActivityStat`

NewMicrosoftGraphItemActivityStatWithDefaults instantiates a new MicrosoftGraphItemActivityStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphItemActivityStat) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphItemActivityStat) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphItemActivityStat) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphItemActivityStat) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccess

`func (o *MicrosoftGraphItemActivityStat) GetAccess() AnyOfmicrosoftGraphItemActionStat`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *MicrosoftGraphItemActivityStat) GetAccessOk() (*AnyOfmicrosoftGraphItemActionStat, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *MicrosoftGraphItemActivityStat) SetAccess(v AnyOfmicrosoftGraphItemActionStat)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *MicrosoftGraphItemActivityStat) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### SetAccessNil

`func (o *MicrosoftGraphItemActivityStat) SetAccessNil(b bool)`

 SetAccessNil sets the value for Access to be an explicit nil

### UnsetAccess
`func (o *MicrosoftGraphItemActivityStat) UnsetAccess()`

UnsetAccess ensures that no value is present for Access, not even an explicit nil
### GetCreate

`func (o *MicrosoftGraphItemActivityStat) GetCreate() AnyOfmicrosoftGraphItemActionStat`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *MicrosoftGraphItemActivityStat) GetCreateOk() (*AnyOfmicrosoftGraphItemActionStat, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *MicrosoftGraphItemActivityStat) SetCreate(v AnyOfmicrosoftGraphItemActionStat)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *MicrosoftGraphItemActivityStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### SetCreateNil

`func (o *MicrosoftGraphItemActivityStat) SetCreateNil(b bool)`

 SetCreateNil sets the value for Create to be an explicit nil

### UnsetCreate
`func (o *MicrosoftGraphItemActivityStat) UnsetCreate()`

UnsetCreate ensures that no value is present for Create, not even an explicit nil
### GetDelete

`func (o *MicrosoftGraphItemActivityStat) GetDelete() AnyOfmicrosoftGraphItemActionStat`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *MicrosoftGraphItemActivityStat) GetDeleteOk() (*AnyOfmicrosoftGraphItemActionStat, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *MicrosoftGraphItemActivityStat) SetDelete(v AnyOfmicrosoftGraphItemActionStat)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *MicrosoftGraphItemActivityStat) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### SetDeleteNil

`func (o *MicrosoftGraphItemActivityStat) SetDeleteNil(b bool)`

 SetDeleteNil sets the value for Delete to be an explicit nil

### UnsetDelete
`func (o *MicrosoftGraphItemActivityStat) UnsetDelete()`

UnsetDelete ensures that no value is present for Delete, not even an explicit nil
### GetEdit

`func (o *MicrosoftGraphItemActivityStat) GetEdit() AnyOfmicrosoftGraphItemActionStat`

GetEdit returns the Edit field if non-nil, zero value otherwise.

### GetEditOk

`func (o *MicrosoftGraphItemActivityStat) GetEditOk() (*AnyOfmicrosoftGraphItemActionStat, bool)`

GetEditOk returns a tuple with the Edit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdit

`func (o *MicrosoftGraphItemActivityStat) SetEdit(v AnyOfmicrosoftGraphItemActionStat)`

SetEdit sets Edit field to given value.

### HasEdit

`func (o *MicrosoftGraphItemActivityStat) HasEdit() bool`

HasEdit returns a boolean if a field has been set.

### SetEditNil

`func (o *MicrosoftGraphItemActivityStat) SetEditNil(b bool)`

 SetEditNil sets the value for Edit to be an explicit nil

### UnsetEdit
`func (o *MicrosoftGraphItemActivityStat) UnsetEdit()`

UnsetEdit ensures that no value is present for Edit, not even an explicit nil
### GetEndDateTime

`func (o *MicrosoftGraphItemActivityStat) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *MicrosoftGraphItemActivityStat) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *MicrosoftGraphItemActivityStat) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *MicrosoftGraphItemActivityStat) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### SetEndDateTimeNil

`func (o *MicrosoftGraphItemActivityStat) SetEndDateTimeNil(b bool)`

 SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil

### UnsetEndDateTime
`func (o *MicrosoftGraphItemActivityStat) UnsetEndDateTime()`

UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
### GetIncompleteData

`func (o *MicrosoftGraphItemActivityStat) GetIncompleteData() AnyOfmicrosoftGraphIncompleteData`

GetIncompleteData returns the IncompleteData field if non-nil, zero value otherwise.

### GetIncompleteDataOk

`func (o *MicrosoftGraphItemActivityStat) GetIncompleteDataOk() (*AnyOfmicrosoftGraphIncompleteData, bool)`

GetIncompleteDataOk returns a tuple with the IncompleteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompleteData

`func (o *MicrosoftGraphItemActivityStat) SetIncompleteData(v AnyOfmicrosoftGraphIncompleteData)`

SetIncompleteData sets IncompleteData field to given value.

### HasIncompleteData

`func (o *MicrosoftGraphItemActivityStat) HasIncompleteData() bool`

HasIncompleteData returns a boolean if a field has been set.

### SetIncompleteDataNil

`func (o *MicrosoftGraphItemActivityStat) SetIncompleteDataNil(b bool)`

 SetIncompleteDataNil sets the value for IncompleteData to be an explicit nil

### UnsetIncompleteData
`func (o *MicrosoftGraphItemActivityStat) UnsetIncompleteData()`

UnsetIncompleteData ensures that no value is present for IncompleteData, not even an explicit nil
### GetIsTrending

`func (o *MicrosoftGraphItemActivityStat) GetIsTrending() bool`

GetIsTrending returns the IsTrending field if non-nil, zero value otherwise.

### GetIsTrendingOk

`func (o *MicrosoftGraphItemActivityStat) GetIsTrendingOk() (*bool, bool)`

GetIsTrendingOk returns a tuple with the IsTrending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrending

`func (o *MicrosoftGraphItemActivityStat) SetIsTrending(v bool)`

SetIsTrending sets IsTrending field to given value.

### HasIsTrending

`func (o *MicrosoftGraphItemActivityStat) HasIsTrending() bool`

HasIsTrending returns a boolean if a field has been set.

### SetIsTrendingNil

`func (o *MicrosoftGraphItemActivityStat) SetIsTrendingNil(b bool)`

 SetIsTrendingNil sets the value for IsTrending to be an explicit nil

### UnsetIsTrending
`func (o *MicrosoftGraphItemActivityStat) UnsetIsTrending()`

UnsetIsTrending ensures that no value is present for IsTrending, not even an explicit nil
### GetMove

`func (o *MicrosoftGraphItemActivityStat) GetMove() AnyOfmicrosoftGraphItemActionStat`

GetMove returns the Move field if non-nil, zero value otherwise.

### GetMoveOk

`func (o *MicrosoftGraphItemActivityStat) GetMoveOk() (*AnyOfmicrosoftGraphItemActionStat, bool)`

GetMoveOk returns a tuple with the Move field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMove

`func (o *MicrosoftGraphItemActivityStat) SetMove(v AnyOfmicrosoftGraphItemActionStat)`

SetMove sets Move field to given value.

### HasMove

`func (o *MicrosoftGraphItemActivityStat) HasMove() bool`

HasMove returns a boolean if a field has been set.

### SetMoveNil

`func (o *MicrosoftGraphItemActivityStat) SetMoveNil(b bool)`

 SetMoveNil sets the value for Move to be an explicit nil

### UnsetMove
`func (o *MicrosoftGraphItemActivityStat) UnsetMove()`

UnsetMove ensures that no value is present for Move, not even an explicit nil
### GetStartDateTime

`func (o *MicrosoftGraphItemActivityStat) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *MicrosoftGraphItemActivityStat) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *MicrosoftGraphItemActivityStat) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *MicrosoftGraphItemActivityStat) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTimeNil

`func (o *MicrosoftGraphItemActivityStat) SetStartDateTimeNil(b bool)`

 SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil

### UnsetStartDateTime
`func (o *MicrosoftGraphItemActivityStat) UnsetStartDateTime()`

UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
### GetActivities

`func (o *MicrosoftGraphItemActivityStat) GetActivities() []MicrosoftGraphItemActivity`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *MicrosoftGraphItemActivityStat) GetActivitiesOk() (*[]MicrosoftGraphItemActivity, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivities

`func (o *MicrosoftGraphItemActivityStat) SetActivities(v []MicrosoftGraphItemActivity)`

SetActivities sets Activities field to given value.

### HasActivities

`func (o *MicrosoftGraphItemActivityStat) HasActivities() bool`

HasActivities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


