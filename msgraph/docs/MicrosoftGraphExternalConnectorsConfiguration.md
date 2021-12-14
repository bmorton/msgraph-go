# MicrosoftGraphExternalConnectorsConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizedAppIds** | Pointer to **[]string** | A collection of application IDs for registered Azure Active Directory apps that are allowed to manage the externalConnection and to index content in the externalConnection. | [optional] 

## Methods

### NewMicrosoftGraphExternalConnectorsConfiguration

`func NewMicrosoftGraphExternalConnectorsConfiguration() *MicrosoftGraphExternalConnectorsConfiguration`

NewMicrosoftGraphExternalConnectorsConfiguration instantiates a new MicrosoftGraphExternalConnectorsConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphExternalConnectorsConfigurationWithDefaults

`func NewMicrosoftGraphExternalConnectorsConfigurationWithDefaults() *MicrosoftGraphExternalConnectorsConfiguration`

NewMicrosoftGraphExternalConnectorsConfigurationWithDefaults instantiates a new MicrosoftGraphExternalConnectorsConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizedAppIds

`func (o *MicrosoftGraphExternalConnectorsConfiguration) GetAuthorizedAppIds() []*string`

GetAuthorizedAppIds returns the AuthorizedAppIds field if non-nil, zero value otherwise.

### GetAuthorizedAppIdsOk

`func (o *MicrosoftGraphExternalConnectorsConfiguration) GetAuthorizedAppIdsOk() (*[]*string, bool)`

GetAuthorizedAppIdsOk returns a tuple with the AuthorizedAppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedAppIds

`func (o *MicrosoftGraphExternalConnectorsConfiguration) SetAuthorizedAppIds(v []*string)`

SetAuthorizedAppIds sets AuthorizedAppIds field to given value.

### HasAuthorizedAppIds

`func (o *MicrosoftGraphExternalConnectorsConfiguration) HasAuthorizedAppIds() bool`

HasAuthorizedAppIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


