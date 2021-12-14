# AppConsentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppDisplayName** | Pointer to **NullableString** | The display name of the app for which consent is requested. Required. Supports $filter (eq only) and $orderby. | [optional] 
**AppId** | Pointer to **string** | The identifier of the application. Required. Supports $filter (eq only) and $orderby. | [optional] 
**PendingScopes** | Pointer to [**[]AnyOfmicrosoftGraphAppConsentRequestScope**](AnyOfmicrosoftGraphAppConsentRequestScope.md) | A list of pending scopes waiting for approval. Required. | [optional] 
**UserConsentRequests** | Pointer to [**[]MicrosoftGraphUserConsentRequest**](MicrosoftGraphUserConsentRequest.md) | A list of pending user consent requests. | [optional] 

## Methods

### NewAppConsentRequest

`func NewAppConsentRequest() *AppConsentRequest`

NewAppConsentRequest instantiates a new AppConsentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppConsentRequestWithDefaults

`func NewAppConsentRequestWithDefaults() *AppConsentRequest`

NewAppConsentRequestWithDefaults instantiates a new AppConsentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppDisplayName

`func (o *AppConsentRequest) GetAppDisplayName() string`

GetAppDisplayName returns the AppDisplayName field if non-nil, zero value otherwise.

### GetAppDisplayNameOk

`func (o *AppConsentRequest) GetAppDisplayNameOk() (*string, bool)`

GetAppDisplayNameOk returns a tuple with the AppDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDisplayName

`func (o *AppConsentRequest) SetAppDisplayName(v string)`

SetAppDisplayName sets AppDisplayName field to given value.

### HasAppDisplayName

`func (o *AppConsentRequest) HasAppDisplayName() bool`

HasAppDisplayName returns a boolean if a field has been set.

### SetAppDisplayNameNil

`func (o *AppConsentRequest) SetAppDisplayNameNil(b bool)`

 SetAppDisplayNameNil sets the value for AppDisplayName to be an explicit nil

### UnsetAppDisplayName
`func (o *AppConsentRequest) UnsetAppDisplayName()`

UnsetAppDisplayName ensures that no value is present for AppDisplayName, not even an explicit nil
### GetAppId

`func (o *AppConsentRequest) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AppConsentRequest) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AppConsentRequest) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *AppConsentRequest) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetPendingScopes

`func (o *AppConsentRequest) GetPendingScopes() []*AnyOfmicrosoftGraphAppConsentRequestScope`

GetPendingScopes returns the PendingScopes field if non-nil, zero value otherwise.

### GetPendingScopesOk

`func (o *AppConsentRequest) GetPendingScopesOk() (*[]*AnyOfmicrosoftGraphAppConsentRequestScope, bool)`

GetPendingScopesOk returns a tuple with the PendingScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingScopes

`func (o *AppConsentRequest) SetPendingScopes(v []*AnyOfmicrosoftGraphAppConsentRequestScope)`

SetPendingScopes sets PendingScopes field to given value.

### HasPendingScopes

`func (o *AppConsentRequest) HasPendingScopes() bool`

HasPendingScopes returns a boolean if a field has been set.

### GetUserConsentRequests

`func (o *AppConsentRequest) GetUserConsentRequests() []MicrosoftGraphUserConsentRequest`

GetUserConsentRequests returns the UserConsentRequests field if non-nil, zero value otherwise.

### GetUserConsentRequestsOk

`func (o *AppConsentRequest) GetUserConsentRequestsOk() (*[]MicrosoftGraphUserConsentRequest, bool)`

GetUserConsentRequestsOk returns a tuple with the UserConsentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserConsentRequests

`func (o *AppConsentRequest) SetUserConsentRequests(v []MicrosoftGraphUserConsentRequest)`

SetUserConsentRequests sets UserConsentRequests field to given value.

### HasUserConsentRequests

`func (o *AppConsentRequest) HasUserConsentRequests() bool`

HasUserConsentRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


