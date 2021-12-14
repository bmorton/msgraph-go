# MicrosoftGraphEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DeletedDateTime** | Pointer to **NullableTime** |  | [optional] 
**Capability** | Pointer to **string** | Describes the capability that is associated with this resource. (e.g. Messages, Conversations, etc.)  Not nullable. Read-only. | [optional] 
**ProviderId** | Pointer to **NullableString** | Application id of the publishing underlying service. Not nullable. Read-only. | [optional] 
**ProviderName** | Pointer to **NullableString** | Name of the publishing underlying service. Read-only. | [optional] 
**ProviderResourceId** | Pointer to **NullableString** | For Microsoft 365 groups, this is set to a well-known name for the resource (e.g. Yammer.FeedURL etc.). Not nullable. Read-only. | [optional] 
**Uri** | Pointer to **string** | URL of the published resource. Not nullable. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphEndpoint

`func NewMicrosoftGraphEndpoint() *MicrosoftGraphEndpoint`

NewMicrosoftGraphEndpoint instantiates a new MicrosoftGraphEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphEndpointWithDefaults

`func NewMicrosoftGraphEndpointWithDefaults() *MicrosoftGraphEndpoint`

NewMicrosoftGraphEndpointWithDefaults instantiates a new MicrosoftGraphEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphEndpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphEndpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphEndpoint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphEndpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeletedDateTime

`func (o *MicrosoftGraphEndpoint) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *MicrosoftGraphEndpoint) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *MicrosoftGraphEndpoint) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *MicrosoftGraphEndpoint) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.

### SetDeletedDateTimeNil

`func (o *MicrosoftGraphEndpoint) SetDeletedDateTimeNil(b bool)`

 SetDeletedDateTimeNil sets the value for DeletedDateTime to be an explicit nil

### UnsetDeletedDateTime
`func (o *MicrosoftGraphEndpoint) UnsetDeletedDateTime()`

UnsetDeletedDateTime ensures that no value is present for DeletedDateTime, not even an explicit nil
### GetCapability

`func (o *MicrosoftGraphEndpoint) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *MicrosoftGraphEndpoint) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *MicrosoftGraphEndpoint) SetCapability(v string)`

SetCapability sets Capability field to given value.

### HasCapability

`func (o *MicrosoftGraphEndpoint) HasCapability() bool`

HasCapability returns a boolean if a field has been set.

### GetProviderId

`func (o *MicrosoftGraphEndpoint) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *MicrosoftGraphEndpoint) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *MicrosoftGraphEndpoint) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *MicrosoftGraphEndpoint) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### SetProviderIdNil

`func (o *MicrosoftGraphEndpoint) SetProviderIdNil(b bool)`

 SetProviderIdNil sets the value for ProviderId to be an explicit nil

### UnsetProviderId
`func (o *MicrosoftGraphEndpoint) UnsetProviderId()`

UnsetProviderId ensures that no value is present for ProviderId, not even an explicit nil
### GetProviderName

`func (o *MicrosoftGraphEndpoint) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *MicrosoftGraphEndpoint) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *MicrosoftGraphEndpoint) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *MicrosoftGraphEndpoint) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *MicrosoftGraphEndpoint) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *MicrosoftGraphEndpoint) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetProviderResourceId

`func (o *MicrosoftGraphEndpoint) GetProviderResourceId() string`

GetProviderResourceId returns the ProviderResourceId field if non-nil, zero value otherwise.

### GetProviderResourceIdOk

`func (o *MicrosoftGraphEndpoint) GetProviderResourceIdOk() (*string, bool)`

GetProviderResourceIdOk returns a tuple with the ProviderResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderResourceId

`func (o *MicrosoftGraphEndpoint) SetProviderResourceId(v string)`

SetProviderResourceId sets ProviderResourceId field to given value.

### HasProviderResourceId

`func (o *MicrosoftGraphEndpoint) HasProviderResourceId() bool`

HasProviderResourceId returns a boolean if a field has been set.

### SetProviderResourceIdNil

`func (o *MicrosoftGraphEndpoint) SetProviderResourceIdNil(b bool)`

 SetProviderResourceIdNil sets the value for ProviderResourceId to be an explicit nil

### UnsetProviderResourceId
`func (o *MicrosoftGraphEndpoint) UnsetProviderResourceId()`

UnsetProviderResourceId ensures that no value is present for ProviderResourceId, not even an explicit nil
### GetUri

`func (o *MicrosoftGraphEndpoint) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *MicrosoftGraphEndpoint) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *MicrosoftGraphEndpoint) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *MicrosoftGraphEndpoint) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


