# PrintOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateTime** | Pointer to **time.Time** | The DateTimeOffset when the operation was created. Read-only. | [optional] 
**Status** | Pointer to [**MicrosoftGraphPrintOperationStatus**](MicrosoftGraphPrintOperationStatus.md) |  | [optional] 

## Methods

### NewPrintOperation

`func NewPrintOperation() *PrintOperation`

NewPrintOperation instantiates a new PrintOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrintOperationWithDefaults

`func NewPrintOperationWithDefaults() *PrintOperation`

NewPrintOperationWithDefaults instantiates a new PrintOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDateTime

`func (o *PrintOperation) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *PrintOperation) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *PrintOperation) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *PrintOperation) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetStatus

`func (o *PrintOperation) GetStatus() MicrosoftGraphPrintOperationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrintOperation) GetStatusOk() (*MicrosoftGraphPrintOperationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrintOperation) SetStatus(v MicrosoftGraphPrintOperationStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PrintOperation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


