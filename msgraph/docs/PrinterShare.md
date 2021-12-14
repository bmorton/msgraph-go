# PrinterShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowAllUsers** | Pointer to **bool** | If true, all users and groups will be granted access to this printer share. This supersedes the allow lists defined by the allowedUsers and allowedGroups navigation properties. | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | The DateTimeOffset when the printer share was created. Read-only. | [optional] 
**AllowedGroups** | Pointer to [**[]MicrosoftGraphGroup**](MicrosoftGraphGroup.md) | The groups whose users have access to print using the printer. | [optional] 
**AllowedUsers** | Pointer to [**[]MicrosoftGraphUser**](MicrosoftGraphUser.md) | The users who have access to print using the printer. | [optional] 
**Printer** | Pointer to [**AnyOfmicrosoftGraphPrinter**](anyOf&lt;microsoft.graph.printer&gt;.md) | The printer that this printer share is related to. | [optional] 

## Methods

### NewPrinterShare

`func NewPrinterShare() *PrinterShare`

NewPrinterShare instantiates a new PrinterShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrinterShareWithDefaults

`func NewPrinterShareWithDefaults() *PrinterShare`

NewPrinterShareWithDefaults instantiates a new PrinterShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowAllUsers

`func (o *PrinterShare) GetAllowAllUsers() bool`

GetAllowAllUsers returns the AllowAllUsers field if non-nil, zero value otherwise.

### GetAllowAllUsersOk

`func (o *PrinterShare) GetAllowAllUsersOk() (*bool, bool)`

GetAllowAllUsersOk returns a tuple with the AllowAllUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAllUsers

`func (o *PrinterShare) SetAllowAllUsers(v bool)`

SetAllowAllUsers sets AllowAllUsers field to given value.

### HasAllowAllUsers

`func (o *PrinterShare) HasAllowAllUsers() bool`

HasAllowAllUsers returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *PrinterShare) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *PrinterShare) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *PrinterShare) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *PrinterShare) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetAllowedGroups

`func (o *PrinterShare) GetAllowedGroups() []MicrosoftGraphGroup`

GetAllowedGroups returns the AllowedGroups field if non-nil, zero value otherwise.

### GetAllowedGroupsOk

`func (o *PrinterShare) GetAllowedGroupsOk() (*[]MicrosoftGraphGroup, bool)`

GetAllowedGroupsOk returns a tuple with the AllowedGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedGroups

`func (o *PrinterShare) SetAllowedGroups(v []MicrosoftGraphGroup)`

SetAllowedGroups sets AllowedGroups field to given value.

### HasAllowedGroups

`func (o *PrinterShare) HasAllowedGroups() bool`

HasAllowedGroups returns a boolean if a field has been set.

### GetAllowedUsers

`func (o *PrinterShare) GetAllowedUsers() []MicrosoftGraphUser`

GetAllowedUsers returns the AllowedUsers field if non-nil, zero value otherwise.

### GetAllowedUsersOk

`func (o *PrinterShare) GetAllowedUsersOk() (*[]MicrosoftGraphUser, bool)`

GetAllowedUsersOk returns a tuple with the AllowedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUsers

`func (o *PrinterShare) SetAllowedUsers(v []MicrosoftGraphUser)`

SetAllowedUsers sets AllowedUsers field to given value.

### HasAllowedUsers

`func (o *PrinterShare) HasAllowedUsers() bool`

HasAllowedUsers returns a boolean if a field has been set.

### GetPrinter

`func (o *PrinterShare) GetPrinter() AnyOfmicrosoftGraphPrinter`

GetPrinter returns the Printer field if non-nil, zero value otherwise.

### GetPrinterOk

`func (o *PrinterShare) GetPrinterOk() (*AnyOfmicrosoftGraphPrinter, bool)`

GetPrinterOk returns a tuple with the Printer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrinter

`func (o *PrinterShare) SetPrinter(v AnyOfmicrosoftGraphPrinter)`

SetPrinter sets Printer field to given value.

### HasPrinter

`func (o *PrinterShare) HasPrinter() bool`

HasPrinter returns a boolean if a field has been set.

### SetPrinterNil

`func (o *PrinterShare) SetPrinterNil(b bool)`

 SetPrinterNil sets the value for Printer to be an explicit nil

### UnsetPrinter
`func (o *PrinterShare) UnsetPrinter()`

UnsetPrinter ensures that no value is present for Printer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


