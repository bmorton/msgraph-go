# MicrosoftGraphAppConsentApprovalRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**AppConsentRequests** | Pointer to [**[]MicrosoftGraphAppConsentRequest**](MicrosoftGraphAppConsentRequest.md) |  | [optional] 

## Methods

### NewMicrosoftGraphAppConsentApprovalRoute

`func NewMicrosoftGraphAppConsentApprovalRoute() *MicrosoftGraphAppConsentApprovalRoute`

NewMicrosoftGraphAppConsentApprovalRoute instantiates a new MicrosoftGraphAppConsentApprovalRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAppConsentApprovalRouteWithDefaults

`func NewMicrosoftGraphAppConsentApprovalRouteWithDefaults() *MicrosoftGraphAppConsentApprovalRoute`

NewMicrosoftGraphAppConsentApprovalRouteWithDefaults instantiates a new MicrosoftGraphAppConsentApprovalRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphAppConsentApprovalRoute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAppConsentApprovalRoute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphAppConsentApprovalRoute) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphAppConsentApprovalRoute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAppConsentRequests

`func (o *MicrosoftGraphAppConsentApprovalRoute) GetAppConsentRequests() []MicrosoftGraphAppConsentRequest`

GetAppConsentRequests returns the AppConsentRequests field if non-nil, zero value otherwise.

### GetAppConsentRequestsOk

`func (o *MicrosoftGraphAppConsentApprovalRoute) GetAppConsentRequestsOk() (*[]MicrosoftGraphAppConsentRequest, bool)`

GetAppConsentRequestsOk returns a tuple with the AppConsentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppConsentRequests

`func (o *MicrosoftGraphAppConsentApprovalRoute) SetAppConsentRequests(v []MicrosoftGraphAppConsentRequest)`

SetAppConsentRequests sets AppConsentRequests field to given value.

### HasAppConsentRequests

`func (o *MicrosoftGraphAppConsentApprovalRoute) HasAppConsentRequests() bool`

HasAppConsentRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


