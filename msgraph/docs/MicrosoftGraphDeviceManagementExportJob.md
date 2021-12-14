# MicrosoftGraphDeviceManagementExportJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ExpirationDateTime** | Pointer to **time.Time** | Time that the exported report expires | [optional] 
**Filter** | Pointer to **NullableString** | Filters applied on the report | [optional] 
**Format** | Pointer to [**AnyOfmicrosoftGraphDeviceManagementReportFileFormat**](anyOf&lt;microsoft.graph.deviceManagementReportFileFormat&gt;.md) | Format of the exported report. Possible values are: csv, pdf. | [optional] 
**LocalizationType** | Pointer to [**AnyOfmicrosoftGraphDeviceManagementExportJobLocalizationType**](anyOf&lt;microsoft.graph.deviceManagementExportJobLocalizationType&gt;.md) | Configures how the requested export job is localized. Possible values are: localizedValuesAsAdditionalColumn, replaceLocalizableValues. | [optional] 
**ReportName** | Pointer to **string** | Name of the report | [optional] 
**RequestDateTime** | Pointer to **time.Time** | Time that the exported report was requested | [optional] 
**Select** | Pointer to **[]string** | Columns selected from the report | [optional] 
**SnapshotId** | Pointer to **NullableString** | A snapshot is an identifiable subset of the dataset represented by the ReportName. A sessionId or CachedReportConfiguration id can be used here. If a sessionId is specified, Filter, Select, and OrderBy are applied to the data represented by the sessionId. Filter, Select, and OrderBy cannot be specified together with a CachedReportConfiguration id. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphDeviceManagementReportStatus**](anyOf&lt;microsoft.graph.deviceManagementReportStatus&gt;.md) | Status of the export job. Possible values are: unknown, notStarted, inProgress, completed, failed. | [optional] 
**Url** | Pointer to **NullableString** | Temporary location of the exported report | [optional] 

## Methods

### NewMicrosoftGraphDeviceManagementExportJob

`func NewMicrosoftGraphDeviceManagementExportJob() *MicrosoftGraphDeviceManagementExportJob`

NewMicrosoftGraphDeviceManagementExportJob instantiates a new MicrosoftGraphDeviceManagementExportJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDeviceManagementExportJobWithDefaults

`func NewMicrosoftGraphDeviceManagementExportJobWithDefaults() *MicrosoftGraphDeviceManagementExportJob`

NewMicrosoftGraphDeviceManagementExportJobWithDefaults instantiates a new MicrosoftGraphDeviceManagementExportJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDeviceManagementExportJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceManagementExportJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDeviceManagementExportJob) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDeviceManagementExportJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExpirationDateTime

`func (o *MicrosoftGraphDeviceManagementExportJob) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *MicrosoftGraphDeviceManagementExportJob) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *MicrosoftGraphDeviceManagementExportJob) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *MicrosoftGraphDeviceManagementExportJob) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### GetFilter

`func (o *MicrosoftGraphDeviceManagementExportJob) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MicrosoftGraphDeviceManagementExportJob) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MicrosoftGraphDeviceManagementExportJob) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *MicrosoftGraphDeviceManagementExportJob) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *MicrosoftGraphDeviceManagementExportJob) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *MicrosoftGraphDeviceManagementExportJob) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetFormat

`func (o *MicrosoftGraphDeviceManagementExportJob) GetFormat() AnyOfmicrosoftGraphDeviceManagementReportFileFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *MicrosoftGraphDeviceManagementExportJob) GetFormatOk() (*AnyOfmicrosoftGraphDeviceManagementReportFileFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *MicrosoftGraphDeviceManagementExportJob) SetFormat(v AnyOfmicrosoftGraphDeviceManagementReportFileFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *MicrosoftGraphDeviceManagementExportJob) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *MicrosoftGraphDeviceManagementExportJob) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *MicrosoftGraphDeviceManagementExportJob) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetLocalizationType

`func (o *MicrosoftGraphDeviceManagementExportJob) GetLocalizationType() AnyOfmicrosoftGraphDeviceManagementExportJobLocalizationType`

GetLocalizationType returns the LocalizationType field if non-nil, zero value otherwise.

### GetLocalizationTypeOk

`func (o *MicrosoftGraphDeviceManagementExportJob) GetLocalizationTypeOk() (*AnyOfmicrosoftGraphDeviceManagementExportJobLocalizationType, bool)`

GetLocalizationTypeOk returns a tuple with the LocalizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizationType

`func (o *MicrosoftGraphDeviceManagementExportJob) SetLocalizationType(v AnyOfmicrosoftGraphDeviceManagementExportJobLocalizationType)`

SetLocalizationType sets LocalizationType field to given value.

### HasLocalizationType

`func (o *MicrosoftGraphDeviceManagementExportJob) HasLocalizationType() bool`

HasLocalizationType returns a boolean if a field has been set.

### SetLocalizationTypeNil

`func (o *MicrosoftGraphDeviceManagementExportJob) SetLocalizationTypeNil(b bool)`

 SetLocalizationTypeNil sets the value for LocalizationType to be an explicit nil

### UnsetLocalizationType
`func (o *MicrosoftGraphDeviceManagementExportJob) UnsetLocalizationType()`

UnsetLocalizationType ensures that no value is present for LocalizationType, not even an explicit nil
### GetReportName

`func (o *MicrosoftGraphDeviceManagementExportJob) GetReportName() string`

GetReportName returns the ReportName field if non-nil, zero value otherwise.

### GetReportNameOk

`func (o *MicrosoftGraphDeviceManagementExportJob) GetReportNameOk() (*string, bool)`

GetReportNameOk returns a tuple with the ReportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportName

`func (o *MicrosoftGraphDeviceManagementExportJob) SetReportName(v string)`

SetReportName sets ReportName field to given value.

### HasReportName

`func (o *MicrosoftGraphDeviceManagementExportJob) HasReportName() bool`

HasReportName returns a boolean if a field has been set.

### GetRequestDateTime

`func (o *MicrosoftGraphDeviceManagementExportJob) GetRequestDateTime() time.Time`

GetRequestDateTime returns the RequestDateTime field if non-nil, zero value otherwise.

### GetRequestDateTimeOk

`func (o *MicrosoftGraphDeviceManagementExportJob) GetRequestDateTimeOk() (*time.Time, bool)`

GetRequestDateTimeOk returns a tuple with the RequestDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDateTime

`func (o *MicrosoftGraphDeviceManagementExportJob) SetRequestDateTime(v time.Time)`

SetRequestDateTime sets RequestDateTime field to given value.

### HasRequestDateTime

`func (o *MicrosoftGraphDeviceManagementExportJob) HasRequestDateTime() bool`

HasRequestDateTime returns a boolean if a field has been set.

### GetSelect

`func (o *MicrosoftGraphDeviceManagementExportJob) GetSelect() []*string`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *MicrosoftGraphDeviceManagementExportJob) GetSelectOk() (*[]*string, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *MicrosoftGraphDeviceManagementExportJob) SetSelect(v []*string)`

SetSelect sets Select field to given value.

### HasSelect

`func (o *MicrosoftGraphDeviceManagementExportJob) HasSelect() bool`

HasSelect returns a boolean if a field has been set.

### GetSnapshotId

`func (o *MicrosoftGraphDeviceManagementExportJob) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *MicrosoftGraphDeviceManagementExportJob) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *MicrosoftGraphDeviceManagementExportJob) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *MicrosoftGraphDeviceManagementExportJob) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### SetSnapshotIdNil

`func (o *MicrosoftGraphDeviceManagementExportJob) SetSnapshotIdNil(b bool)`

 SetSnapshotIdNil sets the value for SnapshotId to be an explicit nil

### UnsetSnapshotId
`func (o *MicrosoftGraphDeviceManagementExportJob) UnsetSnapshotId()`

UnsetSnapshotId ensures that no value is present for SnapshotId, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphDeviceManagementExportJob) GetStatus() AnyOfmicrosoftGraphDeviceManagementReportStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphDeviceManagementExportJob) GetStatusOk() (*AnyOfmicrosoftGraphDeviceManagementReportStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphDeviceManagementExportJob) SetStatus(v AnyOfmicrosoftGraphDeviceManagementReportStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphDeviceManagementExportJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphDeviceManagementExportJob) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphDeviceManagementExportJob) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUrl

`func (o *MicrosoftGraphDeviceManagementExportJob) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MicrosoftGraphDeviceManagementExportJob) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MicrosoftGraphDeviceManagementExportJob) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MicrosoftGraphDeviceManagementExportJob) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *MicrosoftGraphDeviceManagementExportJob) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *MicrosoftGraphDeviceManagementExportJob) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


