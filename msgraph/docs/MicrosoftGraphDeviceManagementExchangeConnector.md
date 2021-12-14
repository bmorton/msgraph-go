# MicrosoftGraphDeviceManagementExchangeConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**ConnectorServerName** | Pointer to **NullableString** | The name of the server hosting the Exchange Connector. | [optional] 
**ExchangeAlias** | Pointer to **NullableString** | An alias assigned to the Exchange server | [optional] 
**ExchangeConnectorType** | Pointer to [**AnyOfmicrosoftGraphDeviceManagementExchangeConnectorType**](anyOf&lt;microsoft.graph.deviceManagementExchangeConnectorType&gt;.md) | The type of Exchange Connector Configured. Possible values are: onPremises, hosted, serviceToService, dedicated. | [optional] 
**ExchangeOrganization** | Pointer to **NullableString** | Exchange Organization to the Exchange server | [optional] 
**LastSyncDateTime** | Pointer to **time.Time** | Last sync time for the Exchange Connector | [optional] 
**PrimarySmtpAddress** | Pointer to **NullableString** | Email address used to configure the Service To Service Exchange Connector. | [optional] 
**ServerName** | Pointer to **NullableString** | The name of the Exchange server. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphDeviceManagementExchangeConnectorStatus**](anyOf&lt;microsoft.graph.deviceManagementExchangeConnectorStatus&gt;.md) | Exchange Connector Status. Possible values are: none, connectionPending, connected, disconnected. | [optional] 
**Version** | Pointer to **NullableString** | The version of the ExchangeConnectorAgent | [optional] 

## Methods

### NewMicrosoftGraphDeviceManagementExchangeConnector

`func NewMicrosoftGraphDeviceManagementExchangeConnector() *MicrosoftGraphDeviceManagementExchangeConnector`

NewMicrosoftGraphDeviceManagementExchangeConnector instantiates a new MicrosoftGraphDeviceManagementExchangeConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDeviceManagementExchangeConnectorWithDefaults

`func NewMicrosoftGraphDeviceManagementExchangeConnectorWithDefaults() *MicrosoftGraphDeviceManagementExchangeConnector`

NewMicrosoftGraphDeviceManagementExchangeConnectorWithDefaults instantiates a new MicrosoftGraphDeviceManagementExchangeConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConnectorServerName

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetConnectorServerName() string`

GetConnectorServerName returns the ConnectorServerName field if non-nil, zero value otherwise.

### GetConnectorServerNameOk

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetConnectorServerNameOk() (*string, bool)`

GetConnectorServerNameOk returns a tuple with the ConnectorServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorServerName

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetConnectorServerName(v string)`

SetConnectorServerName sets ConnectorServerName field to given value.

### HasConnectorServerName

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) HasConnectorServerName() bool`

HasConnectorServerName returns a boolean if a field has been set.

### SetConnectorServerNameNil

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetConnectorServerNameNil(b bool)`

 SetConnectorServerNameNil sets the value for ConnectorServerName to be an explicit nil

### UnsetConnectorServerName
`func (o *MicrosoftGraphDeviceManagementExchangeConnector) UnsetConnectorServerName()`

UnsetConnectorServerName ensures that no value is present for ConnectorServerName, not even an explicit nil
### GetExchangeAlias

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetExchangeAlias() string`

GetExchangeAlias returns the ExchangeAlias field if non-nil, zero value otherwise.

### GetExchangeAliasOk

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetExchangeAliasOk() (*string, bool)`

GetExchangeAliasOk returns a tuple with the ExchangeAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeAlias

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetExchangeAlias(v string)`

SetExchangeAlias sets ExchangeAlias field to given value.

### HasExchangeAlias

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) HasExchangeAlias() bool`

HasExchangeAlias returns a boolean if a field has been set.

### SetExchangeAliasNil

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetExchangeAliasNil(b bool)`

 SetExchangeAliasNil sets the value for ExchangeAlias to be an explicit nil

### UnsetExchangeAlias
`func (o *MicrosoftGraphDeviceManagementExchangeConnector) UnsetExchangeAlias()`

UnsetExchangeAlias ensures that no value is present for ExchangeAlias, not even an explicit nil
### GetExchangeConnectorType

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetExchangeConnectorType() AnyOfmicrosoftGraphDeviceManagementExchangeConnectorType`

GetExchangeConnectorType returns the ExchangeConnectorType field if non-nil, zero value otherwise.

### GetExchangeConnectorTypeOk

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetExchangeConnectorTypeOk() (*AnyOfmicrosoftGraphDeviceManagementExchangeConnectorType, bool)`

GetExchangeConnectorTypeOk returns a tuple with the ExchangeConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeConnectorType

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetExchangeConnectorType(v AnyOfmicrosoftGraphDeviceManagementExchangeConnectorType)`

SetExchangeConnectorType sets ExchangeConnectorType field to given value.

### HasExchangeConnectorType

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) HasExchangeConnectorType() bool`

HasExchangeConnectorType returns a boolean if a field has been set.

### SetExchangeConnectorTypeNil

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetExchangeConnectorTypeNil(b bool)`

 SetExchangeConnectorTypeNil sets the value for ExchangeConnectorType to be an explicit nil

### UnsetExchangeConnectorType
`func (o *MicrosoftGraphDeviceManagementExchangeConnector) UnsetExchangeConnectorType()`

UnsetExchangeConnectorType ensures that no value is present for ExchangeConnectorType, not even an explicit nil
### GetExchangeOrganization

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetExchangeOrganization() string`

GetExchangeOrganization returns the ExchangeOrganization field if non-nil, zero value otherwise.

### GetExchangeOrganizationOk

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetExchangeOrganizationOk() (*string, bool)`

GetExchangeOrganizationOk returns a tuple with the ExchangeOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeOrganization

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetExchangeOrganization(v string)`

SetExchangeOrganization sets ExchangeOrganization field to given value.

### HasExchangeOrganization

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) HasExchangeOrganization() bool`

HasExchangeOrganization returns a boolean if a field has been set.

### SetExchangeOrganizationNil

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetExchangeOrganizationNil(b bool)`

 SetExchangeOrganizationNil sets the value for ExchangeOrganization to be an explicit nil

### UnsetExchangeOrganization
`func (o *MicrosoftGraphDeviceManagementExchangeConnector) UnsetExchangeOrganization()`

UnsetExchangeOrganization ensures that no value is present for ExchangeOrganization, not even an explicit nil
### GetLastSyncDateTime

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetLastSyncDateTime() time.Time`

GetLastSyncDateTime returns the LastSyncDateTime field if non-nil, zero value otherwise.

### GetLastSyncDateTimeOk

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetLastSyncDateTimeOk() (*time.Time, bool)`

GetLastSyncDateTimeOk returns a tuple with the LastSyncDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDateTime

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetLastSyncDateTime(v time.Time)`

SetLastSyncDateTime sets LastSyncDateTime field to given value.

### HasLastSyncDateTime

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) HasLastSyncDateTime() bool`

HasLastSyncDateTime returns a boolean if a field has been set.

### GetPrimarySmtpAddress

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetPrimarySmtpAddress() string`

GetPrimarySmtpAddress returns the PrimarySmtpAddress field if non-nil, zero value otherwise.

### GetPrimarySmtpAddressOk

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetPrimarySmtpAddressOk() (*string, bool)`

GetPrimarySmtpAddressOk returns a tuple with the PrimarySmtpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySmtpAddress

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetPrimarySmtpAddress(v string)`

SetPrimarySmtpAddress sets PrimarySmtpAddress field to given value.

### HasPrimarySmtpAddress

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) HasPrimarySmtpAddress() bool`

HasPrimarySmtpAddress returns a boolean if a field has been set.

### SetPrimarySmtpAddressNil

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetPrimarySmtpAddressNil(b bool)`

 SetPrimarySmtpAddressNil sets the value for PrimarySmtpAddress to be an explicit nil

### UnsetPrimarySmtpAddress
`func (o *MicrosoftGraphDeviceManagementExchangeConnector) UnsetPrimarySmtpAddress()`

UnsetPrimarySmtpAddress ensures that no value is present for PrimarySmtpAddress, not even an explicit nil
### GetServerName

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### SetServerNameNil

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetServerNameNil(b bool)`

 SetServerNameNil sets the value for ServerName to be an explicit nil

### UnsetServerName
`func (o *MicrosoftGraphDeviceManagementExchangeConnector) UnsetServerName()`

UnsetServerName ensures that no value is present for ServerName, not even an explicit nil
### GetStatus

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetStatus() AnyOfmicrosoftGraphDeviceManagementExchangeConnectorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetStatusOk() (*AnyOfmicrosoftGraphDeviceManagementExchangeConnectorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetStatus(v AnyOfmicrosoftGraphDeviceManagementExchangeConnectorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MicrosoftGraphDeviceManagementExchangeConnector) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetVersion

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *MicrosoftGraphDeviceManagementExchangeConnector) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *MicrosoftGraphDeviceManagementExchangeConnector) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


