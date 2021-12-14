# MicrosoftGraphLookupColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowMultipleValues** | Pointer to **NullableBool** | Indicates whether multiple values can be selected from the source. | [optional] 
**AllowUnlimitedLength** | Pointer to **NullableBool** | Indicates whether values in the column should be able to exceed the standard limit of 255 characters. | [optional] 
**ColumnName** | Pointer to **NullableString** | The name of the lookup source column. | [optional] 
**ListId** | Pointer to **NullableString** | The unique identifier of the lookup source list. | [optional] 
**PrimaryLookupColumnId** | Pointer to **NullableString** | If specified, this column is a secondary lookup, pulling an additional field from the list item looked up by the primary lookup. Use the list item looked up by the primary as the source for the column named here. | [optional] 

## Methods

### NewMicrosoftGraphLookupColumn

`func NewMicrosoftGraphLookupColumn() *MicrosoftGraphLookupColumn`

NewMicrosoftGraphLookupColumn instantiates a new MicrosoftGraphLookupColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphLookupColumnWithDefaults

`func NewMicrosoftGraphLookupColumnWithDefaults() *MicrosoftGraphLookupColumn`

NewMicrosoftGraphLookupColumnWithDefaults instantiates a new MicrosoftGraphLookupColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowMultipleValues

`func (o *MicrosoftGraphLookupColumn) GetAllowMultipleValues() bool`

GetAllowMultipleValues returns the AllowMultipleValues field if non-nil, zero value otherwise.

### GetAllowMultipleValuesOk

`func (o *MicrosoftGraphLookupColumn) GetAllowMultipleValuesOk() (*bool, bool)`

GetAllowMultipleValuesOk returns a tuple with the AllowMultipleValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleValues

`func (o *MicrosoftGraphLookupColumn) SetAllowMultipleValues(v bool)`

SetAllowMultipleValues sets AllowMultipleValues field to given value.

### HasAllowMultipleValues

`func (o *MicrosoftGraphLookupColumn) HasAllowMultipleValues() bool`

HasAllowMultipleValues returns a boolean if a field has been set.

### SetAllowMultipleValuesNil

`func (o *MicrosoftGraphLookupColumn) SetAllowMultipleValuesNil(b bool)`

 SetAllowMultipleValuesNil sets the value for AllowMultipleValues to be an explicit nil

### UnsetAllowMultipleValues
`func (o *MicrosoftGraphLookupColumn) UnsetAllowMultipleValues()`

UnsetAllowMultipleValues ensures that no value is present for AllowMultipleValues, not even an explicit nil
### GetAllowUnlimitedLength

`func (o *MicrosoftGraphLookupColumn) GetAllowUnlimitedLength() bool`

GetAllowUnlimitedLength returns the AllowUnlimitedLength field if non-nil, zero value otherwise.

### GetAllowUnlimitedLengthOk

`func (o *MicrosoftGraphLookupColumn) GetAllowUnlimitedLengthOk() (*bool, bool)`

GetAllowUnlimitedLengthOk returns a tuple with the AllowUnlimitedLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnlimitedLength

`func (o *MicrosoftGraphLookupColumn) SetAllowUnlimitedLength(v bool)`

SetAllowUnlimitedLength sets AllowUnlimitedLength field to given value.

### HasAllowUnlimitedLength

`func (o *MicrosoftGraphLookupColumn) HasAllowUnlimitedLength() bool`

HasAllowUnlimitedLength returns a boolean if a field has been set.

### SetAllowUnlimitedLengthNil

`func (o *MicrosoftGraphLookupColumn) SetAllowUnlimitedLengthNil(b bool)`

 SetAllowUnlimitedLengthNil sets the value for AllowUnlimitedLength to be an explicit nil

### UnsetAllowUnlimitedLength
`func (o *MicrosoftGraphLookupColumn) UnsetAllowUnlimitedLength()`

UnsetAllowUnlimitedLength ensures that no value is present for AllowUnlimitedLength, not even an explicit nil
### GetColumnName

`func (o *MicrosoftGraphLookupColumn) GetColumnName() string`

GetColumnName returns the ColumnName field if non-nil, zero value otherwise.

### GetColumnNameOk

`func (o *MicrosoftGraphLookupColumn) GetColumnNameOk() (*string, bool)`

GetColumnNameOk returns a tuple with the ColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnName

`func (o *MicrosoftGraphLookupColumn) SetColumnName(v string)`

SetColumnName sets ColumnName field to given value.

### HasColumnName

`func (o *MicrosoftGraphLookupColumn) HasColumnName() bool`

HasColumnName returns a boolean if a field has been set.

### SetColumnNameNil

`func (o *MicrosoftGraphLookupColumn) SetColumnNameNil(b bool)`

 SetColumnNameNil sets the value for ColumnName to be an explicit nil

### UnsetColumnName
`func (o *MicrosoftGraphLookupColumn) UnsetColumnName()`

UnsetColumnName ensures that no value is present for ColumnName, not even an explicit nil
### GetListId

`func (o *MicrosoftGraphLookupColumn) GetListId() string`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *MicrosoftGraphLookupColumn) GetListIdOk() (*string, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *MicrosoftGraphLookupColumn) SetListId(v string)`

SetListId sets ListId field to given value.

### HasListId

`func (o *MicrosoftGraphLookupColumn) HasListId() bool`

HasListId returns a boolean if a field has been set.

### SetListIdNil

`func (o *MicrosoftGraphLookupColumn) SetListIdNil(b bool)`

 SetListIdNil sets the value for ListId to be an explicit nil

### UnsetListId
`func (o *MicrosoftGraphLookupColumn) UnsetListId()`

UnsetListId ensures that no value is present for ListId, not even an explicit nil
### GetPrimaryLookupColumnId

`func (o *MicrosoftGraphLookupColumn) GetPrimaryLookupColumnId() string`

GetPrimaryLookupColumnId returns the PrimaryLookupColumnId field if non-nil, zero value otherwise.

### GetPrimaryLookupColumnIdOk

`func (o *MicrosoftGraphLookupColumn) GetPrimaryLookupColumnIdOk() (*string, bool)`

GetPrimaryLookupColumnIdOk returns a tuple with the PrimaryLookupColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryLookupColumnId

`func (o *MicrosoftGraphLookupColumn) SetPrimaryLookupColumnId(v string)`

SetPrimaryLookupColumnId sets PrimaryLookupColumnId field to given value.

### HasPrimaryLookupColumnId

`func (o *MicrosoftGraphLookupColumn) HasPrimaryLookupColumnId() bool`

HasPrimaryLookupColumnId returns a boolean if a field has been set.

### SetPrimaryLookupColumnIdNil

`func (o *MicrosoftGraphLookupColumn) SetPrimaryLookupColumnIdNil(b bool)`

 SetPrimaryLookupColumnIdNil sets the value for PrimaryLookupColumnId to be an explicit nil

### UnsetPrimaryLookupColumnId
`func (o *MicrosoftGraphLookupColumn) UnsetPrimaryLookupColumnId()`

UnsetPrimaryLookupColumnId ensures that no value is present for PrimaryLookupColumnId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


