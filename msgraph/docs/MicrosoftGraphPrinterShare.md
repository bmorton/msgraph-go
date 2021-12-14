# MicrosoftGraphPrinterShare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Capabilities** | Pointer to [**AnyOfmicrosoftGraphPrinterCapabilities**](anyOf&lt;microsoft.graph.printerCapabilities&gt;.md) | The capabilities of the printer/printerShare. | [optional] 
**Defaults** | Pointer to [**AnyOfmicrosoftGraphPrinterDefaults**](anyOf&lt;microsoft.graph.printerDefaults&gt;.md) | The default print settings of printer/printerShare. | [optional] 
**DisplayName** | Pointer to **string** | The name of the printer/printerShare. | [optional] 
**IsAcceptingJobs** | Pointer to **NullableBool** | Whether the printer/printerShare is currently accepting new print jobs. | [optional] 
**Location** | Pointer to [**AnyOfmicrosoftGraphPrinterLocation**](anyOf&lt;microsoft.graph.printerLocation&gt;.md) | The physical and/or organizational location of the printer/printerShare. | [optional] 
**Manufacturer** | Pointer to **NullableString** | The manufacturer of the printer/printerShare. | [optional] 
**Model** | Pointer to **NullableString** | The model name of the printer/printerShare. | [optional] 
**Status** | Pointer to [**MicrosoftGraphPrinterStatus**](MicrosoftGraphPrinterStatus.md) |  | [optional] 
**Jobs** | Pointer to [**[]MicrosoftGraphPrintJob**](MicrosoftGraphPrintJob.md) | The list of jobs that are queued for printing by the printer/printerShare. | [optional] 
**AllowAllUsers** | Pointer to **bool** | If true, all users and groups will be granted access to this printer share. This supersedes the allow lists defined by the allowedUsers and allowedGroups navigation properties. | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | The DateTimeOffset when the printer share was created. Read-only. | [optional] 
**AllowedGroups** | Pointer to [**[]MicrosoftGraphGroup**](MicrosoftGraphGroup.md) | The groups whose users have access to print using the printer. | [optional] 
**AllowedUsers** | Pointer to [**[]MicrosoftGraphUser**](MicrosoftGraphUser.md) | The users who have access to print using the printer. | [optional] 
**Printer** | Pointer to [**AnyOfmicrosoftGraphPrinter**](anyOf&lt;microsoft.graph.printer&gt;.md) | The printer that this printer share is related to. | [optional] 

## Methods

### NewMicrosoftGraphPrinterShare

`func NewMicrosoftGraphPrinterShare() *MicrosoftGraphPrinterShare`

NewMicrosoftGraphPrinterShare instantiates a new MicrosoftGraphPrinterShare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPrinterShareWithDefaults

`func NewMicrosoftGraphPrinterShareWithDefaults() *MicrosoftGraphPrinterShare`

NewMicrosoftGraphPrinterShareWithDefaults instantiates a new MicrosoftGraphPrinterShare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphPrinterShare) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPrinterShare) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphPrinterShare) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphPrinterShare) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCapabilities

`func (o *MicrosoftGraphPrinterShare) GetCapabilities() AnyOfmicrosoftGraphPrinterCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *MicrosoftGraphPrinterShare) GetCapabilitiesOk() (*AnyOfmicrosoftGraphPrinterCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *MicrosoftGraphPrinterShare) SetCapabilities(v AnyOfmicrosoftGraphPrinterCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *MicrosoftGraphPrinterShare) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *MicrosoftGraphPrinterShare) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *MicrosoftGraphPrinterShare) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetDefaults

`func (o *MicrosoftGraphPrinterShare) GetDefaults() AnyOfmicrosoftGraphPrinterDefaults`

GetDefaults returns the Defaults field if non-nil, zero value otherwise.

### GetDefaultsOk

`func (o *MicrosoftGraphPrinterShare) GetDefaultsOk() (*AnyOfmicrosoftGraphPrinterDefaults, bool)`

GetDefaultsOk returns a tuple with the Defaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaults

`func (o *MicrosoftGraphPrinterShare) SetDefaults(v AnyOfmicrosoftGraphPrinterDefaults)`

SetDefaults sets Defaults field to given value.

### HasDefaults

`func (o *MicrosoftGraphPrinterShare) HasDefaults() bool`

HasDefaults returns a boolean if a field has been set.

### SetDefaultsNil

`func (o *MicrosoftGraphPrinterShare) SetDefaultsNil(b bool)`

 SetDefaultsNil sets the value for Defaults to be an explicit nil

### UnsetDefaults
`func (o *MicrosoftGraphPrinterShare) UnsetDefaults()`

UnsetDefaults ensures that no value is present for Defaults, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphPrinterShare) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphPrinterShare) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphPrinterShare) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphPrinterShare) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIsAcceptingJobs

`func (o *MicrosoftGraphPrinterShare) GetIsAcceptingJobs() bool`

GetIsAcceptingJobs returns the IsAcceptingJobs field if non-nil, zero value otherwise.

### GetIsAcceptingJobsOk

`func (o *MicrosoftGraphPrinterShare) GetIsAcceptingJobsOk() (*bool, bool)`

GetIsAcceptingJobsOk returns a tuple with the IsAcceptingJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAcceptingJobs

`func (o *MicrosoftGraphPrinterShare) SetIsAcceptingJobs(v bool)`

SetIsAcceptingJobs sets IsAcceptingJobs field to given value.

### HasIsAcceptingJobs

`func (o *MicrosoftGraphPrinterShare) HasIsAcceptingJobs() bool`

HasIsAcceptingJobs returns a boolean if a field has been set.

### SetIsAcceptingJobsNil

`func (o *MicrosoftGraphPrinterShare) SetIsAcceptingJobsNil(b bool)`

 SetIsAcceptingJobsNil sets the value for IsAcceptingJobs to be an explicit nil

### UnsetIsAcceptingJobs
`func (o *MicrosoftGraphPrinterShare) UnsetIsAcceptingJobs()`

UnsetIsAcceptingJobs ensures that no value is present for IsAcceptingJobs, not even an explicit nil
### GetLocation

`func (o *MicrosoftGraphPrinterShare) GetLocation() AnyOfmicrosoftGraphPrinterLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MicrosoftGraphPrinterShare) GetLocationOk() (*AnyOfmicrosoftGraphPrinterLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MicrosoftGraphPrinterShare) SetLocation(v AnyOfmicrosoftGraphPrinterLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *MicrosoftGraphPrinterShare) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *MicrosoftGraphPrinterShare) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *MicrosoftGraphPrinterShare) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetManufacturer

`func (o *MicrosoftGraphPrinterShare) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *MicrosoftGraphPrinterShare) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *MicrosoftGraphPrinterShare) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *MicrosoftGraphPrinterShare) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *MicrosoftGraphPrinterShare) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *MicrosoftGraphPrinterShare) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetModel

`func (o *MicrosoftGraphPrinterShare) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MicrosoftGraphPrinterShare) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *MicrosoftGraphPrinterShare) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *MicrosoftGraphPrinterShare) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *MicrosoftGraphPrinterShare) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *MicrosoftGraphPrinterShare) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphPrinterShare) GetStatus() MicrosoftGraphPrinterStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphPrinterShare) GetStatusOk() (*MicrosoftGraphPrinterStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphPrinterShare) SetStatus(v MicrosoftGraphPrinterStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphPrinterShare) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetJobs

`func (o *MicrosoftGraphPrinterShare) GetJobs() []MicrosoftGraphPrintJob`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *MicrosoftGraphPrinterShare) GetJobsOk() (*[]MicrosoftGraphPrintJob, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *MicrosoftGraphPrinterShare) SetJobs(v []MicrosoftGraphPrintJob)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *MicrosoftGraphPrinterShare) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetAllowAllUsers

`func (o *MicrosoftGraphPrinterShare) GetAllowAllUsers() bool`

GetAllowAllUsers returns the AllowAllUsers field if non-nil, zero value otherwise.

### GetAllowAllUsersOk

`func (o *MicrosoftGraphPrinterShare) GetAllowAllUsersOk() (*bool, bool)`

GetAllowAllUsersOk returns a tuple with the AllowAllUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAllUsers

`func (o *MicrosoftGraphPrinterShare) SetAllowAllUsers(v bool)`

SetAllowAllUsers sets AllowAllUsers field to given value.

### HasAllowAllUsers

`func (o *MicrosoftGraphPrinterShare) HasAllowAllUsers() bool`

HasAllowAllUsers returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *MicrosoftGraphPrinterShare) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphPrinterShare) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphPrinterShare) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphPrinterShare) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetAllowedGroups

`func (o *MicrosoftGraphPrinterShare) GetAllowedGroups() []MicrosoftGraphGroup`

GetAllowedGroups returns the AllowedGroups field if non-nil, zero value otherwise.

### GetAllowedGroupsOk

`func (o *MicrosoftGraphPrinterShare) GetAllowedGroupsOk() (*[]MicrosoftGraphGroup, bool)`

GetAllowedGroupsOk returns a tuple with the AllowedGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedGroups

`func (o *MicrosoftGraphPrinterShare) SetAllowedGroups(v []MicrosoftGraphGroup)`

SetAllowedGroups sets AllowedGroups field to given value.

### HasAllowedGroups

`func (o *MicrosoftGraphPrinterShare) HasAllowedGroups() bool`

HasAllowedGroups returns a boolean if a field has been set.

### GetAllowedUsers

`func (o *MicrosoftGraphPrinterShare) GetAllowedUsers() []MicrosoftGraphUser`

GetAllowedUsers returns the AllowedUsers field if non-nil, zero value otherwise.

### GetAllowedUsersOk

`func (o *MicrosoftGraphPrinterShare) GetAllowedUsersOk() (*[]MicrosoftGraphUser, bool)`

GetAllowedUsersOk returns a tuple with the AllowedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUsers

`func (o *MicrosoftGraphPrinterShare) SetAllowedUsers(v []MicrosoftGraphUser)`

SetAllowedUsers sets AllowedUsers field to given value.

### HasAllowedUsers

`func (o *MicrosoftGraphPrinterShare) HasAllowedUsers() bool`

HasAllowedUsers returns a boolean if a field has been set.

### GetPrinter

`func (o *MicrosoftGraphPrinterShare) GetPrinter() AnyOfmicrosoftGraphPrinter`

GetPrinter returns the Printer field if non-nil, zero value otherwise.

### GetPrinterOk

`func (o *MicrosoftGraphPrinterShare) GetPrinterOk() (*AnyOfmicrosoftGraphPrinter, bool)`

GetPrinterOk returns a tuple with the Printer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrinter

`func (o *MicrosoftGraphPrinterShare) SetPrinter(v AnyOfmicrosoftGraphPrinter)`

SetPrinter sets Printer field to given value.

### HasPrinter

`func (o *MicrosoftGraphPrinterShare) HasPrinter() bool`

HasPrinter returns a boolean if a field has been set.

### SetPrinterNil

`func (o *MicrosoftGraphPrinterShare) SetPrinterNil(b bool)`

 SetPrinterNil sets the value for Printer to be an explicit nil

### UnsetPrinter
`func (o *MicrosoftGraphPrinterShare) UnsetPrinter()`

UnsetPrinter ensures that no value is present for Printer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


