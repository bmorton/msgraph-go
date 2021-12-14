# MicrosoftGraphAppConsentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**AppDisplayName** | Pointer to **NullableString** | The display name of the app for which consent is requested. Required. Supports $filter (eq only) and $orderby. | [optional] 
**AppId** | Pointer to **string** | The identifier of the application. Required. Supports $filter (eq only) and $orderby. | [optional] 
**PendingScopes** | Pointer to [**[]AnyOfmicrosoftGraphAppConsentRequestScope**](AnyOfmicrosoftGraphAppConsentRequestScope.md) | A list of pending scopes waiting for approval. Required. | [optional] 
**UserConsentRequests** | Pointer to [**[]MicrosoftGraphUserConsentRequest**](MicrosoftGraphUserConsentRequest.md) | A list of pending user consent requests. | [optional] 

## Methods

### NewMicrosoftGraphAppConsentRequest

`func NewMicrosoftGraphAppConsentRequest() *MicrosoftGraphAppConsentRequest`

NewMicrosoftGraphAppConsentRequest instantiates a new MicrosoftGraphAppConsentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAppConsentRequestWithDefaults

`func NewMicrosoftGraphAppConsentRequestWithDefaults() *MicrosoftGraphAppConsentRequest`

NewMicrosoftGraphAppConsentRequestWithDefaults instantiates a new MicrosoftGraphAppConsentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphAppConsentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAppConsentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphAppConsentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphAppConsentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAppDisplayName

`func (o *MicrosoftGraphAppConsentRequest) GetAppDisplayName() string`

GetAppDisplayName returns the AppDisplayName field if non-nil, zero value otherwise.

### GetAppDisplayNameOk

`func (o *MicrosoftGraphAppConsentRequest) GetAppDisplayNameOk() (*string, bool)`

GetAppDisplayNameOk returns a tuple with the AppDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDisplayName

`func (o *MicrosoftGraphAppConsentRequest) SetAppDisplayName(v string)`

SetAppDisplayName sets AppDisplayName field to given value.

### HasAppDisplayName

`func (o *MicrosoftGraphAppConsentRequest) HasAppDisplayName() bool`

HasAppDisplayName returns a boolean if a field has been set.

### SetAppDisplayNameNil

`func (o *MicrosoftGraphAppConsentRequest) SetAppDisplayNameNil(b bool)`

 SetAppDisplayNameNil sets the value for AppDisplayName to be an explicit nil

### UnsetAppDisplayName
`func (o *MicrosoftGraphAppConsentRequest) UnsetAppDisplayName()`

UnsetAppDisplayName ensures that no value is present for AppDisplayName, not even an explicit nil
### GetAppId

`func (o *MicrosoftGraphAppConsentRequest) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *MicrosoftGraphAppConsentRequest) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *MicrosoftGraphAppConsentRequest) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *MicrosoftGraphAppConsentRequest) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetPendingScopes

`func (o *MicrosoftGraphAppConsentRequest) GetPendingScopes() []*AnyOfmicrosoftGraphAppConsentRequestScope`

GetPendingScopes returns the PendingScopes field if non-nil, zero value otherwise.

### GetPendingScopesOk

`func (o *MicrosoftGraphAppConsentRequest) GetPendingScopesOk() (*[]*AnyOfmicrosoftGraphAppConsentRequestScope, bool)`

GetPendingScopesOk returns a tuple with the PendingScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingScopes

`func (o *MicrosoftGraphAppConsentRequest) SetPendingScopes(v []*AnyOfmicrosoftGraphAppConsentRequestScope)`

SetPendingScopes sets PendingScopes field to given value.

### HasPendingScopes

`func (o *MicrosoftGraphAppConsentRequest) HasPendingScopes() bool`

HasPendingScopes returns a boolean if a field has been set.

### GetUserConsentRequests

`func (o *MicrosoftGraphAppConsentRequest) GetUserConsentRequests() []MicrosoftGraphUserConsentRequest`

GetUserConsentRequests returns the UserConsentRequests field if non-nil, zero value otherwise.

### GetUserConsentRequestsOk

`func (o *MicrosoftGraphAppConsentRequest) GetUserConsentRequestsOk() (*[]MicrosoftGraphUserConsentRequest, bool)`

GetUserConsentRequestsOk returns a tuple with the UserConsentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserConsentRequests

`func (o *MicrosoftGraphAppConsentRequest) SetUserConsentRequests(v []MicrosoftGraphUserConsentRequest)`

SetUserConsentRequests sets UserConsentRequests field to given value.

### HasUserConsentRequests

`func (o *MicrosoftGraphAppConsentRequest) HasUserConsentRequests() bool`

HasUserConsentRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


