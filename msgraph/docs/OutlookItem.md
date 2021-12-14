# OutlookItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to **[]string** | The categories associated with the item | [optional] 
**ChangeKey** | Pointer to **NullableString** | Identifies the version of the item. Every time the item is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 

## Methods

### NewOutlookItem

`func NewOutlookItem() *OutlookItem`

NewOutlookItem instantiates a new OutlookItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutlookItemWithDefaults

`func NewOutlookItemWithDefaults() *OutlookItem`

NewOutlookItemWithDefaults instantiates a new OutlookItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *OutlookItem) GetCategories() []*string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *OutlookItem) GetCategoriesOk() (*[]*string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *OutlookItem) SetCategories(v []*string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *OutlookItem) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetChangeKey

`func (o *OutlookItem) GetChangeKey() string`

GetChangeKey returns the ChangeKey field if non-nil, zero value otherwise.

### GetChangeKeyOk

`func (o *OutlookItem) GetChangeKeyOk() (*string, bool)`

GetChangeKeyOk returns a tuple with the ChangeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeKey

`func (o *OutlookItem) SetChangeKey(v string)`

SetChangeKey sets ChangeKey field to given value.

### HasChangeKey

`func (o *OutlookItem) HasChangeKey() bool`

HasChangeKey returns a boolean if a field has been set.

### SetChangeKeyNil

`func (o *OutlookItem) SetChangeKeyNil(b bool)`

 SetChangeKeyNil sets the value for ChangeKey to be an explicit nil

### UnsetChangeKey
`func (o *OutlookItem) UnsetChangeKey()`

UnsetChangeKey ensures that no value is present for ChangeKey, not even an explicit nil
### GetCreatedDateTime

`func (o *OutlookItem) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *OutlookItem) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *OutlookItem) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *OutlookItem) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *OutlookItem) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *OutlookItem) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetLastModifiedDateTime

`func (o *OutlookItem) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *OutlookItem) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *OutlookItem) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *OutlookItem) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *OutlookItem) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *OutlookItem) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


