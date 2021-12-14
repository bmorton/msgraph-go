# MicrosoftGraphUriClickSecurityState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClickAction** | Pointer to **NullableString** |  | [optional] 
**ClickDateTime** | Pointer to **NullableTime** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**SourceId** | Pointer to **NullableString** |  | [optional] 
**UriDomain** | Pointer to **NullableString** |  | [optional] 
**Verdict** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMicrosoftGraphUriClickSecurityState

`func NewMicrosoftGraphUriClickSecurityState() *MicrosoftGraphUriClickSecurityState`

NewMicrosoftGraphUriClickSecurityState instantiates a new MicrosoftGraphUriClickSecurityState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphUriClickSecurityStateWithDefaults

`func NewMicrosoftGraphUriClickSecurityStateWithDefaults() *MicrosoftGraphUriClickSecurityState`

NewMicrosoftGraphUriClickSecurityStateWithDefaults instantiates a new MicrosoftGraphUriClickSecurityState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClickAction

`func (o *MicrosoftGraphUriClickSecurityState) GetClickAction() string`

GetClickAction returns the ClickAction field if non-nil, zero value otherwise.

### GetClickActionOk

`func (o *MicrosoftGraphUriClickSecurityState) GetClickActionOk() (*string, bool)`

GetClickActionOk returns a tuple with the ClickAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickAction

`func (o *MicrosoftGraphUriClickSecurityState) SetClickAction(v string)`

SetClickAction sets ClickAction field to given value.

### HasClickAction

`func (o *MicrosoftGraphUriClickSecurityState) HasClickAction() bool`

HasClickAction returns a boolean if a field has been set.

### SetClickActionNil

`func (o *MicrosoftGraphUriClickSecurityState) SetClickActionNil(b bool)`

 SetClickActionNil sets the value for ClickAction to be an explicit nil

### UnsetClickAction
`func (o *MicrosoftGraphUriClickSecurityState) UnsetClickAction()`

UnsetClickAction ensures that no value is present for ClickAction, not even an explicit nil
### GetClickDateTime

`func (o *MicrosoftGraphUriClickSecurityState) GetClickDateTime() time.Time`

GetClickDateTime returns the ClickDateTime field if non-nil, zero value otherwise.

### GetClickDateTimeOk

`func (o *MicrosoftGraphUriClickSecurityState) GetClickDateTimeOk() (*time.Time, bool)`

GetClickDateTimeOk returns a tuple with the ClickDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickDateTime

`func (o *MicrosoftGraphUriClickSecurityState) SetClickDateTime(v time.Time)`

SetClickDateTime sets ClickDateTime field to given value.

### HasClickDateTime

`func (o *MicrosoftGraphUriClickSecurityState) HasClickDateTime() bool`

HasClickDateTime returns a boolean if a field has been set.

### SetClickDateTimeNil

`func (o *MicrosoftGraphUriClickSecurityState) SetClickDateTimeNil(b bool)`

 SetClickDateTimeNil sets the value for ClickDateTime to be an explicit nil

### UnsetClickDateTime
`func (o *MicrosoftGraphUriClickSecurityState) UnsetClickDateTime()`

UnsetClickDateTime ensures that no value is present for ClickDateTime, not even an explicit nil
### GetId

`func (o *MicrosoftGraphUriClickSecurityState) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphUriClickSecurityState) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphUriClickSecurityState) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphUriClickSecurityState) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MicrosoftGraphUriClickSecurityState) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MicrosoftGraphUriClickSecurityState) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSourceId

`func (o *MicrosoftGraphUriClickSecurityState) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *MicrosoftGraphUriClickSecurityState) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *MicrosoftGraphUriClickSecurityState) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *MicrosoftGraphUriClickSecurityState) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *MicrosoftGraphUriClickSecurityState) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *MicrosoftGraphUriClickSecurityState) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetUriDomain

`func (o *MicrosoftGraphUriClickSecurityState) GetUriDomain() string`

GetUriDomain returns the UriDomain field if non-nil, zero value otherwise.

### GetUriDomainOk

`func (o *MicrosoftGraphUriClickSecurityState) GetUriDomainOk() (*string, bool)`

GetUriDomainOk returns a tuple with the UriDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUriDomain

`func (o *MicrosoftGraphUriClickSecurityState) SetUriDomain(v string)`

SetUriDomain sets UriDomain field to given value.

### HasUriDomain

`func (o *MicrosoftGraphUriClickSecurityState) HasUriDomain() bool`

HasUriDomain returns a boolean if a field has been set.

### SetUriDomainNil

`func (o *MicrosoftGraphUriClickSecurityState) SetUriDomainNil(b bool)`

 SetUriDomainNil sets the value for UriDomain to be an explicit nil

### UnsetUriDomain
`func (o *MicrosoftGraphUriClickSecurityState) UnsetUriDomain()`

UnsetUriDomain ensures that no value is present for UriDomain, not even an explicit nil
### GetVerdict

`func (o *MicrosoftGraphUriClickSecurityState) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *MicrosoftGraphUriClickSecurityState) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *MicrosoftGraphUriClickSecurityState) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.

### HasVerdict

`func (o *MicrosoftGraphUriClickSecurityState) HasVerdict() bool`

HasVerdict returns a boolean if a field has been set.

### SetVerdictNil

`func (o *MicrosoftGraphUriClickSecurityState) SetVerdictNil(b bool)`

 SetVerdictNil sets the value for Verdict to be an explicit nil

### UnsetVerdict
`func (o *MicrosoftGraphUriClickSecurityState) UnsetVerdict()`

UnsetVerdict ensures that no value is present for Verdict, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


