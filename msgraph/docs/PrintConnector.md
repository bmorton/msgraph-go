# PrintConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppVersion** | Pointer to **string** | The connector&#39;s version. | [optional] 
**DisplayName** | Pointer to **string** | The name of the connector. | [optional] 
**FullyQualifiedDomainName** | Pointer to **string** | The connector machine&#39;s hostname. | [optional] 
**Location** | Pointer to [**AnyOfmicrosoftGraphPrinterLocation**](anyOf&lt;microsoft.graph.printerLocation&gt;.md) | The physical and/or organizational location of the connector. | [optional] 
**OperatingSystem** | Pointer to **string** | The connector machine&#39;s operating system version. | [optional] 
**RegisteredDateTime** | Pointer to **time.Time** | The DateTimeOffset when the connector was registered. | [optional] 

## Methods

### NewPrintConnector

`func NewPrintConnector() *PrintConnector`

NewPrintConnector instantiates a new PrintConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrintConnectorWithDefaults

`func NewPrintConnectorWithDefaults() *PrintConnector`

NewPrintConnectorWithDefaults instantiates a new PrintConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppVersion

`func (o *PrintConnector) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *PrintConnector) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *PrintConnector) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *PrintConnector) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### GetDisplayName

`func (o *PrintConnector) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PrintConnector) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PrintConnector) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PrintConnector) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFullyQualifiedDomainName

`func (o *PrintConnector) GetFullyQualifiedDomainName() string`

GetFullyQualifiedDomainName returns the FullyQualifiedDomainName field if non-nil, zero value otherwise.

### GetFullyQualifiedDomainNameOk

`func (o *PrintConnector) GetFullyQualifiedDomainNameOk() (*string, bool)`

GetFullyQualifiedDomainNameOk returns a tuple with the FullyQualifiedDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyQualifiedDomainName

`func (o *PrintConnector) SetFullyQualifiedDomainName(v string)`

SetFullyQualifiedDomainName sets FullyQualifiedDomainName field to given value.

### HasFullyQualifiedDomainName

`func (o *PrintConnector) HasFullyQualifiedDomainName() bool`

HasFullyQualifiedDomainName returns a boolean if a field has been set.

### GetLocation

`func (o *PrintConnector) GetLocation() AnyOfmicrosoftGraphPrinterLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PrintConnector) GetLocationOk() (*AnyOfmicrosoftGraphPrinterLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PrintConnector) SetLocation(v AnyOfmicrosoftGraphPrinterLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PrintConnector) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *PrintConnector) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *PrintConnector) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetOperatingSystem

`func (o *PrintConnector) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *PrintConnector) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *PrintConnector) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *PrintConnector) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetRegisteredDateTime

`func (o *PrintConnector) GetRegisteredDateTime() time.Time`

GetRegisteredDateTime returns the RegisteredDateTime field if non-nil, zero value otherwise.

### GetRegisteredDateTimeOk

`func (o *PrintConnector) GetRegisteredDateTimeOk() (*time.Time, bool)`

GetRegisteredDateTimeOk returns a tuple with the RegisteredDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDateTime

`func (o *PrintConnector) SetRegisteredDateTime(v time.Time)`

SetRegisteredDateTime sets RegisteredDateTime field to given value.

### HasRegisteredDateTime

`func (o *PrintConnector) HasRegisteredDateTime() bool`

HasRegisteredDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


