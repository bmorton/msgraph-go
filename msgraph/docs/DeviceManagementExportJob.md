# DeviceManagementExportJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewDeviceManagementExportJob

`func NewDeviceManagementExportJob() *DeviceManagementExportJob`

NewDeviceManagementExportJob instantiates a new DeviceManagementExportJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceManagementExportJobWithDefaults

`func NewDeviceManagementExportJobWithDefaults() *DeviceManagementExportJob`

NewDeviceManagementExportJobWithDefaults instantiates a new DeviceManagementExportJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationDateTime

`func (o *DeviceManagementExportJob) GetExpirationDateTime() time.Time`

GetExpirationDateTime returns the ExpirationDateTime field if non-nil, zero value otherwise.

### GetExpirationDateTimeOk

`func (o *DeviceManagementExportJob) GetExpirationDateTimeOk() (*time.Time, bool)`

GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDateTime

`func (o *DeviceManagementExportJob) SetExpirationDateTime(v time.Time)`

SetExpirationDateTime sets ExpirationDateTime field to given value.

### HasExpirationDateTime

`func (o *DeviceManagementExportJob) HasExpirationDateTime() bool`

HasExpirationDateTime returns a boolean if a field has been set.

### GetFilter

`func (o *DeviceManagementExportJob) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *DeviceManagementExportJob) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *DeviceManagementExportJob) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *DeviceManagementExportJob) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *DeviceManagementExportJob) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *DeviceManagementExportJob) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetFormat

`func (o *DeviceManagementExportJob) GetFormat() AnyOfmicrosoftGraphDeviceManagementReportFileFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *DeviceManagementExportJob) GetFormatOk() (*AnyOfmicrosoftGraphDeviceManagementReportFileFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *DeviceManagementExportJob) SetFormat(v AnyOfmicrosoftGraphDeviceManagementReportFileFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *DeviceManagementExportJob) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *DeviceManagementExportJob) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *DeviceManagementExportJob) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetLocalizationType

`func (o *DeviceManagementExportJob) GetLocalizationType() AnyOfmicrosoftGraphDeviceManagementExportJobLocalizationType`

GetLocalizationType returns the LocalizationType field if non-nil, zero value otherwise.

### GetLocalizationTypeOk

`func (o *DeviceManagementExportJob) GetLocalizationTypeOk() (*AnyOfmicrosoftGraphDeviceManagementExportJobLocalizationType, bool)`

GetLocalizationTypeOk returns a tuple with the LocalizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizationType

`func (o *DeviceManagementExportJob) SetLocalizationType(v AnyOfmicrosoftGraphDeviceManagementExportJobLocalizationType)`

SetLocalizationType sets LocalizationType field to given value.

### HasLocalizationType

`func (o *DeviceManagementExportJob) HasLocalizationType() bool`

HasLocalizationType returns a boolean if a field has been set.

### SetLocalizationTypeNil

`func (o *DeviceManagementExportJob) SetLocalizationTypeNil(b bool)`

 SetLocalizationTypeNil sets the value for LocalizationType to be an explicit nil

### UnsetLocalizationType
`func (o *DeviceManagementExportJob) UnsetLocalizationType()`

UnsetLocalizationType ensures that no value is present for LocalizationType, not even an explicit nil
### GetReportName

`func (o *DeviceManagementExportJob) GetReportName() string`

GetReportName returns the ReportName field if non-nil, zero value otherwise.

### GetReportNameOk

`func (o *DeviceManagementExportJob) GetReportNameOk() (*string, bool)`

GetReportNameOk returns a tuple with the ReportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportName

`func (o *DeviceManagementExportJob) SetReportName(v string)`

SetReportName sets ReportName field to given value.

### HasReportName

`func (o *DeviceManagementExportJob) HasReportName() bool`

HasReportName returns a boolean if a field has been set.

### GetRequestDateTime

`func (o *DeviceManagementExportJob) GetRequestDateTime() time.Time`

GetRequestDateTime returns the RequestDateTime field if non-nil, zero value otherwise.

### GetRequestDateTimeOk

`func (o *DeviceManagementExportJob) GetRequestDateTimeOk() (*time.Time, bool)`

GetRequestDateTimeOk returns a tuple with the RequestDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDateTime

`func (o *DeviceManagementExportJob) SetRequestDateTime(v time.Time)`

SetRequestDateTime sets RequestDateTime field to given value.

### HasRequestDateTime

`func (o *DeviceManagementExportJob) HasRequestDateTime() bool`

HasRequestDateTime returns a boolean if a field has been set.

### GetSelect

`func (o *DeviceManagementExportJob) GetSelect() []*string`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *DeviceManagementExportJob) GetSelectOk() (*[]*string, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *DeviceManagementExportJob) SetSelect(v []*string)`

SetSelect sets Select field to given value.

### HasSelect

`func (o *DeviceManagementExportJob) HasSelect() bool`

HasSelect returns a boolean if a field has been set.

### GetSnapshotId

`func (o *DeviceManagementExportJob) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *DeviceManagementExportJob) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *DeviceManagementExportJob) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *DeviceManagementExportJob) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### SetSnapshotIdNil

`func (o *DeviceManagementExportJob) SetSnapshotIdNil(b bool)`

 SetSnapshotIdNil sets the value for SnapshotId to be an explicit nil

### UnsetSnapshotId
`func (o *DeviceManagementExportJob) UnsetSnapshotId()`

UnsetSnapshotId ensures that no value is present for SnapshotId, not even an explicit nil
### GetStatus

`func (o *DeviceManagementExportJob) GetStatus() AnyOfmicrosoftGraphDeviceManagementReportStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceManagementExportJob) GetStatusOk() (*AnyOfmicrosoftGraphDeviceManagementReportStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceManagementExportJob) SetStatus(v AnyOfmicrosoftGraphDeviceManagementReportStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceManagementExportJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *DeviceManagementExportJob) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *DeviceManagementExportJob) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUrl

`func (o *DeviceManagementExportJob) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DeviceManagementExportJob) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DeviceManagementExportJob) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DeviceManagementExportJob) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *DeviceManagementExportJob) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *DeviceManagementExportJob) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


