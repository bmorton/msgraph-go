# MicrosoftGraphPrintConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**AppVersion** | Pointer to **string** | The connector&#39;s version. | [optional] 
**DisplayName** | Pointer to **string** | The name of the connector. | [optional] 
**FullyQualifiedDomainName** | Pointer to **string** | The connector machine&#39;s hostname. | [optional] 
**Location** | Pointer to [**AnyOfmicrosoftGraphPrinterLocation**](anyOf&lt;microsoft.graph.printerLocation&gt;.md) | The physical and/or organizational location of the connector. | [optional] 
**OperatingSystem** | Pointer to **string** | The connector machine&#39;s operating system version. | [optional] 
**RegisteredDateTime** | Pointer to **time.Time** | The DateTimeOffset when the connector was registered. | [optional] 

## Methods

### NewMicrosoftGraphPrintConnector

`func NewMicrosoftGraphPrintConnector() *MicrosoftGraphPrintConnector`

NewMicrosoftGraphPrintConnector instantiates a new MicrosoftGraphPrintConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPrintConnectorWithDefaults

`func NewMicrosoftGraphPrintConnectorWithDefaults() *MicrosoftGraphPrintConnector`

NewMicrosoftGraphPrintConnectorWithDefaults instantiates a new MicrosoftGraphPrintConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphPrintConnector) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPrintConnector) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphPrintConnector) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphPrintConnector) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAppVersion

`func (o *MicrosoftGraphPrintConnector) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *MicrosoftGraphPrintConnector) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *MicrosoftGraphPrintConnector) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *MicrosoftGraphPrintConnector) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### GetDisplayName

`func (o *MicrosoftGraphPrintConnector) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphPrintConnector) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphPrintConnector) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphPrintConnector) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFullyQualifiedDomainName

`func (o *MicrosoftGraphPrintConnector) GetFullyQualifiedDomainName() string`

GetFullyQualifiedDomainName returns the FullyQualifiedDomainName field if non-nil, zero value otherwise.

### GetFullyQualifiedDomainNameOk

`func (o *MicrosoftGraphPrintConnector) GetFullyQualifiedDomainNameOk() (*string, bool)`

GetFullyQualifiedDomainNameOk returns a tuple with the FullyQualifiedDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyQualifiedDomainName

`func (o *MicrosoftGraphPrintConnector) SetFullyQualifiedDomainName(v string)`

SetFullyQualifiedDomainName sets FullyQualifiedDomainName field to given value.

### HasFullyQualifiedDomainName

`func (o *MicrosoftGraphPrintConnector) HasFullyQualifiedDomainName() bool`

HasFullyQualifiedDomainName returns a boolean if a field has been set.

### GetLocation

`func (o *MicrosoftGraphPrintConnector) GetLocation() AnyOfmicrosoftGraphPrinterLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *MicrosoftGraphPrintConnector) GetLocationOk() (*AnyOfmicrosoftGraphPrinterLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *MicrosoftGraphPrintConnector) SetLocation(v AnyOfmicrosoftGraphPrinterLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *MicrosoftGraphPrintConnector) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *MicrosoftGraphPrintConnector) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *MicrosoftGraphPrintConnector) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetOperatingSystem

`func (o *MicrosoftGraphPrintConnector) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *MicrosoftGraphPrintConnector) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *MicrosoftGraphPrintConnector) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *MicrosoftGraphPrintConnector) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetRegisteredDateTime

`func (o *MicrosoftGraphPrintConnector) GetRegisteredDateTime() time.Time`

GetRegisteredDateTime returns the RegisteredDateTime field if non-nil, zero value otherwise.

### GetRegisteredDateTimeOk

`func (o *MicrosoftGraphPrintConnector) GetRegisteredDateTimeOk() (*time.Time, bool)`

GetRegisteredDateTimeOk returns a tuple with the RegisteredDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDateTime

`func (o *MicrosoftGraphPrintConnector) SetRegisteredDateTime(v time.Time)`

SetRegisteredDateTime sets RegisteredDateTime field to given value.

### HasRegisteredDateTime

`func (o *MicrosoftGraphPrintConnector) HasRegisteredDateTime() bool`

HasRegisteredDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


