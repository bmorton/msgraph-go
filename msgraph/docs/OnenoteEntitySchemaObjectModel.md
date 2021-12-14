# OnenoteEntitySchemaObjectModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateTime** | Pointer to **NullableTime** | The date and time when the page was created. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. | [optional] 

## Methods

### NewOnenoteEntitySchemaObjectModel

`func NewOnenoteEntitySchemaObjectModel() *OnenoteEntitySchemaObjectModel`

NewOnenoteEntitySchemaObjectModel instantiates a new OnenoteEntitySchemaObjectModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnenoteEntitySchemaObjectModelWithDefaults

`func NewOnenoteEntitySchemaObjectModelWithDefaults() *OnenoteEntitySchemaObjectModel`

NewOnenoteEntitySchemaObjectModelWithDefaults instantiates a new OnenoteEntitySchemaObjectModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDateTime

`func (o *OnenoteEntitySchemaObjectModel) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *OnenoteEntitySchemaObjectModel) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *OnenoteEntitySchemaObjectModel) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *OnenoteEntitySchemaObjectModel) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *OnenoteEntitySchemaObjectModel) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *OnenoteEntitySchemaObjectModel) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


