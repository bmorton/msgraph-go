# ItemActivityStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewItemActivityStat

`func NewItemActivityStat() *ItemActivityStat`

NewItemActivityStat instantiates a new ItemActivityStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemActivityStatWithDefaults

`func NewItemActivityStatWithDefaults() *ItemActivityStat`

NewItemActivityStatWithDefaults instantiates a new ItemActivityStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *ItemActivityStat) GetAccess() AnyOfmicrosoftGraphItemActionStat`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ItemActivityStat) GetAccessOk() (*AnyOfmicrosoftGraphItemActionStat, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ItemActivityStat) SetAccess(v AnyOfmicrosoftGraphItemActionStat)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ItemActivityStat) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### SetAccessNil

`func (o *ItemActivityStat) SetAccessNil(b bool)`

 SetAccessNil sets the value for Access to be an explicit nil

### UnsetAccess
`func (o *ItemActivityStat) UnsetAccess()`

UnsetAccess ensures that no value is present for Access, not even an explicit nil
### GetCreate

`func (o *ItemActivityStat) GetCreate() AnyOfmicrosoftGraphItemActionStat`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ItemActivityStat) GetCreateOk() (*AnyOfmicrosoftGraphItemActionStat, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ItemActivityStat) SetCreate(v AnyOfmicrosoftGraphItemActionStat)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ItemActivityStat) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### SetCreateNil

`func (o *ItemActivityStat) SetCreateNil(b bool)`

 SetCreateNil sets the value for Create to be an explicit nil

### UnsetCreate
`func (o *ItemActivityStat) UnsetCreate()`

UnsetCreate ensures that no value is present for Create, not even an explicit nil
### GetDelete

`func (o *ItemActivityStat) GetDelete() AnyOfmicrosoftGraphItemActionStat`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *ItemActivityStat) GetDeleteOk() (*AnyOfmicrosoftGraphItemActionStat, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *ItemActivityStat) SetDelete(v AnyOfmicrosoftGraphItemActionStat)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *ItemActivityStat) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### SetDeleteNil

`func (o *ItemActivityStat) SetDeleteNil(b bool)`

 SetDeleteNil sets the value for Delete to be an explicit nil

### UnsetDelete
`func (o *ItemActivityStat) UnsetDelete()`

UnsetDelete ensures that no value is present for Delete, not even an explicit nil
### GetEdit

`func (o *ItemActivityStat) GetEdit() AnyOfmicrosoftGraphItemActionStat`

GetEdit returns the Edit field if non-nil, zero value otherwise.

### GetEditOk

`func (o *ItemActivityStat) GetEditOk() (*AnyOfmicrosoftGraphItemActionStat, bool)`

GetEditOk returns a tuple with the Edit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdit

`func (o *ItemActivityStat) SetEdit(v AnyOfmicrosoftGraphItemActionStat)`

SetEdit sets Edit field to given value.

### HasEdit

`func (o *ItemActivityStat) HasEdit() bool`

HasEdit returns a boolean if a field has been set.

### SetEditNil

`func (o *ItemActivityStat) SetEditNil(b bool)`

 SetEditNil sets the value for Edit to be an explicit nil

### UnsetEdit
`func (o *ItemActivityStat) UnsetEdit()`

UnsetEdit ensures that no value is present for Edit, not even an explicit nil
### GetEndDateTime

`func (o *ItemActivityStat) GetEndDateTime() time.Time`

GetEndDateTime returns the EndDateTime field if non-nil, zero value otherwise.

### GetEndDateTimeOk

`func (o *ItemActivityStat) GetEndDateTimeOk() (*time.Time, bool)`

GetEndDateTimeOk returns a tuple with the EndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDateTime

`func (o *ItemActivityStat) SetEndDateTime(v time.Time)`

SetEndDateTime sets EndDateTime field to given value.

### HasEndDateTime

`func (o *ItemActivityStat) HasEndDateTime() bool`

HasEndDateTime returns a boolean if a field has been set.

### SetEndDateTimeNil

`func (o *ItemActivityStat) SetEndDateTimeNil(b bool)`

 SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil

### UnsetEndDateTime
`func (o *ItemActivityStat) UnsetEndDateTime()`

UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
### GetIncompleteData

`func (o *ItemActivityStat) GetIncompleteData() AnyOfmicrosoftGraphIncompleteData`

GetIncompleteData returns the IncompleteData field if non-nil, zero value otherwise.

### GetIncompleteDataOk

`func (o *ItemActivityStat) GetIncompleteDataOk() (*AnyOfmicrosoftGraphIncompleteData, bool)`

GetIncompleteDataOk returns a tuple with the IncompleteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompleteData

`func (o *ItemActivityStat) SetIncompleteData(v AnyOfmicrosoftGraphIncompleteData)`

SetIncompleteData sets IncompleteData field to given value.

### HasIncompleteData

`func (o *ItemActivityStat) HasIncompleteData() bool`

HasIncompleteData returns a boolean if a field has been set.

### SetIncompleteDataNil

`func (o *ItemActivityStat) SetIncompleteDataNil(b bool)`

 SetIncompleteDataNil sets the value for IncompleteData to be an explicit nil

### UnsetIncompleteData
`func (o *ItemActivityStat) UnsetIncompleteData()`

UnsetIncompleteData ensures that no value is present for IncompleteData, not even an explicit nil
### GetIsTrending

`func (o *ItemActivityStat) GetIsTrending() bool`

GetIsTrending returns the IsTrending field if non-nil, zero value otherwise.

### GetIsTrendingOk

`func (o *ItemActivityStat) GetIsTrendingOk() (*bool, bool)`

GetIsTrendingOk returns a tuple with the IsTrending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrending

`func (o *ItemActivityStat) SetIsTrending(v bool)`

SetIsTrending sets IsTrending field to given value.

### HasIsTrending

`func (o *ItemActivityStat) HasIsTrending() bool`

HasIsTrending returns a boolean if a field has been set.

### SetIsTrendingNil

`func (o *ItemActivityStat) SetIsTrendingNil(b bool)`

 SetIsTrendingNil sets the value for IsTrending to be an explicit nil

### UnsetIsTrending
`func (o *ItemActivityStat) UnsetIsTrending()`

UnsetIsTrending ensures that no value is present for IsTrending, not even an explicit nil
### GetMove

`func (o *ItemActivityStat) GetMove() AnyOfmicrosoftGraphItemActionStat`

GetMove returns the Move field if non-nil, zero value otherwise.

### GetMoveOk

`func (o *ItemActivityStat) GetMoveOk() (*AnyOfmicrosoftGraphItemActionStat, bool)`

GetMoveOk returns a tuple with the Move field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMove

`func (o *ItemActivityStat) SetMove(v AnyOfmicrosoftGraphItemActionStat)`

SetMove sets Move field to given value.

### HasMove

`func (o *ItemActivityStat) HasMove() bool`

HasMove returns a boolean if a field has been set.

### SetMoveNil

`func (o *ItemActivityStat) SetMoveNil(b bool)`

 SetMoveNil sets the value for Move to be an explicit nil

### UnsetMove
`func (o *ItemActivityStat) UnsetMove()`

UnsetMove ensures that no value is present for Move, not even an explicit nil
### GetStartDateTime

`func (o *ItemActivityStat) GetStartDateTime() time.Time`

GetStartDateTime returns the StartDateTime field if non-nil, zero value otherwise.

### GetStartDateTimeOk

`func (o *ItemActivityStat) GetStartDateTimeOk() (*time.Time, bool)`

GetStartDateTimeOk returns a tuple with the StartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDateTime

`func (o *ItemActivityStat) SetStartDateTime(v time.Time)`

SetStartDateTime sets StartDateTime field to given value.

### HasStartDateTime

`func (o *ItemActivityStat) HasStartDateTime() bool`

HasStartDateTime returns a boolean if a field has been set.

### SetStartDateTimeNil

`func (o *ItemActivityStat) SetStartDateTimeNil(b bool)`

 SetStartDateTimeNil sets the value for StartDateTime to be an explicit nil

### UnsetStartDateTime
`func (o *ItemActivityStat) UnsetStartDateTime()`

UnsetStartDateTime ensures that no value is present for StartDateTime, not even an explicit nil
### GetActivities

`func (o *ItemActivityStat) GetActivities() []MicrosoftGraphItemActivity`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *ItemActivityStat) GetActivitiesOk() (*[]MicrosoftGraphItemActivity, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivities

`func (o *ItemActivityStat) SetActivities(v []MicrosoftGraphItemActivity)`

SetActivities sets Activities field to given value.

### HasActivities

`func (o *ItemActivityStat) HasActivities() bool`

HasActivities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


