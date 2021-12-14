# MicrosoftGraphSharingLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to [**AnyOfmicrosoftGraphIdentity**](anyOf&lt;microsoft.graph.identity&gt;.md) | The app the link is associated with. | [optional] 
**PreventsDownload** | Pointer to **NullableBool** | If true then the user can only use this link to view the item on the web, and cannot use it to download the contents of the item. Only for OneDrive for Business and SharePoint. | [optional] 
**Scope** | Pointer to **NullableString** | The scope of the link represented by this permission. Value anonymous indicates the link is usable by anyone, organization indicates the link is only usable for users signed into the same tenant. | [optional] 
**Type** | Pointer to **NullableString** | The type of the link created. | [optional] 
**WebHtml** | Pointer to **NullableString** | For embed links, this property contains the HTML code for an &lt;iframe&gt; element that will embed the item in a webpage. | [optional] 
**WebUrl** | Pointer to **NullableString** | A URL that opens the item in the browser on the OneDrive website. | [optional] 

## Methods

### NewMicrosoftGraphSharingLink

`func NewMicrosoftGraphSharingLink() *MicrosoftGraphSharingLink`

NewMicrosoftGraphSharingLink instantiates a new MicrosoftGraphSharingLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSharingLinkWithDefaults

`func NewMicrosoftGraphSharingLinkWithDefaults() *MicrosoftGraphSharingLink`

NewMicrosoftGraphSharingLinkWithDefaults instantiates a new MicrosoftGraphSharingLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *MicrosoftGraphSharingLink) GetApplication() AnyOfmicrosoftGraphIdentity`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *MicrosoftGraphSharingLink) GetApplicationOk() (*AnyOfmicrosoftGraphIdentity, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *MicrosoftGraphSharingLink) SetApplication(v AnyOfmicrosoftGraphIdentity)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *MicrosoftGraphSharingLink) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### SetApplicationNil

`func (o *MicrosoftGraphSharingLink) SetApplicationNil(b bool)`

 SetApplicationNil sets the value for Application to be an explicit nil

### UnsetApplication
`func (o *MicrosoftGraphSharingLink) UnsetApplication()`

UnsetApplication ensures that no value is present for Application, not even an explicit nil
### GetPreventsDownload

`func (o *MicrosoftGraphSharingLink) GetPreventsDownload() bool`

GetPreventsDownload returns the PreventsDownload field if non-nil, zero value otherwise.

### GetPreventsDownloadOk

`func (o *MicrosoftGraphSharingLink) GetPreventsDownloadOk() (*bool, bool)`

GetPreventsDownloadOk returns a tuple with the PreventsDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventsDownload

`func (o *MicrosoftGraphSharingLink) SetPreventsDownload(v bool)`

SetPreventsDownload sets PreventsDownload field to given value.

### HasPreventsDownload

`func (o *MicrosoftGraphSharingLink) HasPreventsDownload() bool`

HasPreventsDownload returns a boolean if a field has been set.

### SetPreventsDownloadNil

`func (o *MicrosoftGraphSharingLink) SetPreventsDownloadNil(b bool)`

 SetPreventsDownloadNil sets the value for PreventsDownload to be an explicit nil

### UnsetPreventsDownload
`func (o *MicrosoftGraphSharingLink) UnsetPreventsDownload()`

UnsetPreventsDownload ensures that no value is present for PreventsDownload, not even an explicit nil
### GetScope

`func (o *MicrosoftGraphSharingLink) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *MicrosoftGraphSharingLink) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *MicrosoftGraphSharingLink) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *MicrosoftGraphSharingLink) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *MicrosoftGraphSharingLink) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *MicrosoftGraphSharingLink) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetType

`func (o *MicrosoftGraphSharingLink) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MicrosoftGraphSharingLink) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MicrosoftGraphSharingLink) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MicrosoftGraphSharingLink) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MicrosoftGraphSharingLink) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MicrosoftGraphSharingLink) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetWebHtml

`func (o *MicrosoftGraphSharingLink) GetWebHtml() string`

GetWebHtml returns the WebHtml field if non-nil, zero value otherwise.

### GetWebHtmlOk

`func (o *MicrosoftGraphSharingLink) GetWebHtmlOk() (*string, bool)`

GetWebHtmlOk returns a tuple with the WebHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebHtml

`func (o *MicrosoftGraphSharingLink) SetWebHtml(v string)`

SetWebHtml sets WebHtml field to given value.

### HasWebHtml

`func (o *MicrosoftGraphSharingLink) HasWebHtml() bool`

HasWebHtml returns a boolean if a field has been set.

### SetWebHtmlNil

`func (o *MicrosoftGraphSharingLink) SetWebHtmlNil(b bool)`

 SetWebHtmlNil sets the value for WebHtml to be an explicit nil

### UnsetWebHtml
`func (o *MicrosoftGraphSharingLink) UnsetWebHtml()`

UnsetWebHtml ensures that no value is present for WebHtml, not even an explicit nil
### GetWebUrl

`func (o *MicrosoftGraphSharingLink) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *MicrosoftGraphSharingLink) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *MicrosoftGraphSharingLink) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.

### HasWebUrl

`func (o *MicrosoftGraphSharingLink) HasWebUrl() bool`

HasWebUrl returns a boolean if a field has been set.

### SetWebUrlNil

`func (o *MicrosoftGraphSharingLink) SetWebUrlNil(b bool)`

 SetWebUrlNil sets the value for WebUrl to be an explicit nil

### UnsetWebUrl
`func (o *MicrosoftGraphSharingLink) UnsetWebUrl()`

UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


