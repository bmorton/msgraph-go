# MicrosoftGraphDomainDnsRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**IsOptional** | Pointer to **bool** | If false, this record must be configured by the customer at the DNS host for Microsoft Online Services to operate correctly with the domain. | [optional] 
**Label** | Pointer to **string** | Value used when configuring the name of the DNS record at the DNS host. | [optional] 
**RecordType** | Pointer to **NullableString** | Indicates what type of DNS record this entity represents.The value can be one of the following: CName, Mx, Srv, TxtKey | [optional] 
**SupportedService** | Pointer to **string** | Microsoft Online Service or feature that has a dependency on this DNS record.Can be one of the following values: null, Email, Sharepoint, EmailInternalRelayOnly, OfficeCommunicationsOnline, SharePointDefaultDomain, FullRedelegation, SharePointPublic, OrgIdAuthentication, Yammer, Intune | [optional] 
**Ttl** | Pointer to **int32** | Value to use when configuring the time-to-live (ttl) property of the DNS record at the DNS host. Not nullable | [optional] 

## Methods

### NewMicrosoftGraphDomainDnsRecord

`func NewMicrosoftGraphDomainDnsRecord() *MicrosoftGraphDomainDnsRecord`

NewMicrosoftGraphDomainDnsRecord instantiates a new MicrosoftGraphDomainDnsRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDomainDnsRecordWithDefaults

`func NewMicrosoftGraphDomainDnsRecordWithDefaults() *MicrosoftGraphDomainDnsRecord`

NewMicrosoftGraphDomainDnsRecordWithDefaults instantiates a new MicrosoftGraphDomainDnsRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphDomainDnsRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphDomainDnsRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphDomainDnsRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphDomainDnsRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsOptional

`func (o *MicrosoftGraphDomainDnsRecord) GetIsOptional() bool`

GetIsOptional returns the IsOptional field if non-nil, zero value otherwise.

### GetIsOptionalOk

`func (o *MicrosoftGraphDomainDnsRecord) GetIsOptionalOk() (*bool, bool)`

GetIsOptionalOk returns a tuple with the IsOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOptional

`func (o *MicrosoftGraphDomainDnsRecord) SetIsOptional(v bool)`

SetIsOptional sets IsOptional field to given value.

### HasIsOptional

`func (o *MicrosoftGraphDomainDnsRecord) HasIsOptional() bool`

HasIsOptional returns a boolean if a field has been set.

### GetLabel

`func (o *MicrosoftGraphDomainDnsRecord) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MicrosoftGraphDomainDnsRecord) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MicrosoftGraphDomainDnsRecord) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *MicrosoftGraphDomainDnsRecord) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRecordType

`func (o *MicrosoftGraphDomainDnsRecord) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *MicrosoftGraphDomainDnsRecord) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *MicrosoftGraphDomainDnsRecord) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *MicrosoftGraphDomainDnsRecord) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### SetRecordTypeNil

`func (o *MicrosoftGraphDomainDnsRecord) SetRecordTypeNil(b bool)`

 SetRecordTypeNil sets the value for RecordType to be an explicit nil

### UnsetRecordType
`func (o *MicrosoftGraphDomainDnsRecord) UnsetRecordType()`

UnsetRecordType ensures that no value is present for RecordType, not even an explicit nil
### GetSupportedService

`func (o *MicrosoftGraphDomainDnsRecord) GetSupportedService() string`

GetSupportedService returns the SupportedService field if non-nil, zero value otherwise.

### GetSupportedServiceOk

`func (o *MicrosoftGraphDomainDnsRecord) GetSupportedServiceOk() (*string, bool)`

GetSupportedServiceOk returns a tuple with the SupportedService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedService

`func (o *MicrosoftGraphDomainDnsRecord) SetSupportedService(v string)`

SetSupportedService sets SupportedService field to given value.

### HasSupportedService

`func (o *MicrosoftGraphDomainDnsRecord) HasSupportedService() bool`

HasSupportedService returns a boolean if a field has been set.

### GetTtl

`func (o *MicrosoftGraphDomainDnsRecord) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *MicrosoftGraphDomainDnsRecord) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *MicrosoftGraphDomainDnsRecord) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *MicrosoftGraphDomainDnsRecord) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


