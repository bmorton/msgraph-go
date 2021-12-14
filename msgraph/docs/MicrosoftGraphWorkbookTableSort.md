# MicrosoftGraphWorkbookTableSort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Fields** | Pointer to [**[]AnyOfmicrosoftGraphWorkbookSortField**](AnyOfmicrosoftGraphWorkbookSortField.md) | Represents the current conditions used to last sort the table. Read-only. | [optional] 
**MatchCase** | Pointer to **bool** | Represents whether the casing impacted the last sort of the table. Read-only. | [optional] 
**Method** | Pointer to **string** | Represents Chinese character ordering method last used to sort the table. The possible values are: PinYin, StrokeCount. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphWorkbookTableSort

`func NewMicrosoftGraphWorkbookTableSort() *MicrosoftGraphWorkbookTableSort`

NewMicrosoftGraphWorkbookTableSort instantiates a new MicrosoftGraphWorkbookTableSort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookTableSortWithDefaults

`func NewMicrosoftGraphWorkbookTableSortWithDefaults() *MicrosoftGraphWorkbookTableSort`

NewMicrosoftGraphWorkbookTableSortWithDefaults instantiates a new MicrosoftGraphWorkbookTableSort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphWorkbookTableSort) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphWorkbookTableSort) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphWorkbookTableSort) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphWorkbookTableSort) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFields

`func (o *MicrosoftGraphWorkbookTableSort) GetFields() []*AnyOfmicrosoftGraphWorkbookSortField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MicrosoftGraphWorkbookTableSort) GetFieldsOk() (*[]*AnyOfmicrosoftGraphWorkbookSortField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MicrosoftGraphWorkbookTableSort) SetFields(v []*AnyOfmicrosoftGraphWorkbookSortField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *MicrosoftGraphWorkbookTableSort) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetMatchCase

`func (o *MicrosoftGraphWorkbookTableSort) GetMatchCase() bool`

GetMatchCase returns the MatchCase field if non-nil, zero value otherwise.

### GetMatchCaseOk

`func (o *MicrosoftGraphWorkbookTableSort) GetMatchCaseOk() (*bool, bool)`

GetMatchCaseOk returns a tuple with the MatchCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCase

`func (o *MicrosoftGraphWorkbookTableSort) SetMatchCase(v bool)`

SetMatchCase sets MatchCase field to given value.

### HasMatchCase

`func (o *MicrosoftGraphWorkbookTableSort) HasMatchCase() bool`

HasMatchCase returns a boolean if a field has been set.

### GetMethod

`func (o *MicrosoftGraphWorkbookTableSort) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *MicrosoftGraphWorkbookTableSort) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *MicrosoftGraphWorkbookTableSort) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *MicrosoftGraphWorkbookTableSort) HasMethod() bool`

HasMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


