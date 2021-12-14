# MicrosoftGraphOnenoteEntitySchemaObjectModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Self** | Pointer to **NullableString** | The endpoint where you can get details about the page. Read-only. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The date and time when the page was created. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphOnenoteEntitySchemaObjectModel

`func NewMicrosoftGraphOnenoteEntitySchemaObjectModel() *MicrosoftGraphOnenoteEntitySchemaObjectModel`

NewMicrosoftGraphOnenoteEntitySchemaObjectModel instantiates a new MicrosoftGraphOnenoteEntitySchemaObjectModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphOnenoteEntitySchemaObjectModelWithDefaults

`func NewMicrosoftGraphOnenoteEntitySchemaObjectModelWithDefaults() *MicrosoftGraphOnenoteEntitySchemaObjectModel`

NewMicrosoftGraphOnenoteEntitySchemaObjectModelWithDefaults instantiates a new MicrosoftGraphOnenoteEntitySchemaObjectModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphOnenoteEntitySchemaObjectModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphOnenoteEntitySchemaObjectModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphOnenoteEntitySchemaObjectModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphOnenoteEntitySchemaObjectModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSelf

`func (o *MicrosoftGraphOnenoteEntitySchemaObjectModel) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *MicrosoftGraphOnenoteEntitySchemaObjectModel) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *MicrosoftGraphOnenoteEntitySchemaObjectModel) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *MicrosoftGraphOnenoteEntitySchemaObjectModel) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### SetSelfNil

`func (o *MicrosoftGraphOnenoteEntitySchemaObjectModel) SetSelfNil(b bool)`

 SetSelfNil sets the value for Self to be an explicit nil

### UnsetSelf
`func (o *MicrosoftGraphOnenoteEntitySchemaObjectModel) UnsetSelf()`

UnsetSelf ensures that no value is present for Self, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphOnenoteEntitySchemaObjectModel) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphOnenoteEntitySchemaObjectModel) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphOnenoteEntitySchemaObjectModel) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphOnenoteEntitySchemaObjectModel) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphOnenoteEntitySchemaObjectModel) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphOnenoteEntitySchemaObjectModel) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


