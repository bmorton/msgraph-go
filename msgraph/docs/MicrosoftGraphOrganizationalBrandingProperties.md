# MicrosoftGraphOrganizationalBrandingProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**BackgroundColor** | Pointer to **NullableString** | Color that will appear in place of the background image in low-bandwidth connections. We recommend that you use the primary color of your banner logo or your organization color. Specify this in hexadecimal format, for example, white is #FFFFFF. | [optional] 
**BackgroundImage** | Pointer to **NullableString** | Image that appears as the background of the sign-in page. The allowed types are PNG or JPEG not smaller than 300 KB and not larger than 1920 × 1080 pixels. A smaller image will reduce bandwidth requirements and make the page load faster. | [optional] 
**BackgroundImageRelativeUrl** | Pointer to **NullableString** | A relative URL for the backgroundImage property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only. | [optional] 
**BannerLogo** | Pointer to **NullableString** | A banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG no larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo. | [optional] 
**BannerLogoRelativeUrl** | Pointer to **NullableString** | A relative url for the bannerLogo property that is combined with a CDN base URL from the cdnList to provide the read-only version served by a CDN. Read-only. | [optional] 
**CdnList** | Pointer to **[]string** | A list of base URLs for all available CDN providers that are serving the assets of the current resource. Several CDN providers are used at the same time for high availability of read requests. Read-only. | [optional] 
**SignInPageText** | Pointer to **NullableString** | Text that appears at the bottom of the sign-in box. You can use this to communicate additional information, such as the phone number to your help desk or a legal statement. This text must be Unicode and not exceed 1024 characters. | [optional] 
**SquareLogo** | Pointer to **NullableString** | A square version of your company logo that appears in Windows 10 out-of-box experiences (OOBE) and when Windows Autopilot is enabled for deployment. Allowed types are PNG or JPEG no larger than 240 x 240 pixels and no more than 10 KB in size. We recommend using a transparent image with no padding around the logo. | [optional] 
**SquareLogoRelativeUrl** | Pointer to **NullableString** | A relative url for the squareLogo property that is combined with a CDN base URL from the cdnList to provide the version served by a CDN. Read-only. | [optional] 
**UsernameHintText** | Pointer to **NullableString** | String that shows as the hint in the username textbox on the sign-in screen. This text must be a Unicode, without links or code, and can&#39;t exceed 64 characters. | [optional] 

## Methods

### NewMicrosoftGraphOrganizationalBrandingProperties

`func NewMicrosoftGraphOrganizationalBrandingProperties() *MicrosoftGraphOrganizationalBrandingProperties`

NewMicrosoftGraphOrganizationalBrandingProperties instantiates a new MicrosoftGraphOrganizationalBrandingProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphOrganizationalBrandingPropertiesWithDefaults

`func NewMicrosoftGraphOrganizationalBrandingPropertiesWithDefaults() *MicrosoftGraphOrganizationalBrandingProperties`

NewMicrosoftGraphOrganizationalBrandingPropertiesWithDefaults instantiates a new MicrosoftGraphOrganizationalBrandingProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphOrganizationalBrandingProperties) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphOrganizationalBrandingProperties) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphOrganizationalBrandingProperties) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphOrganizationalBrandingProperties) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBackgroundColor

`func (o *MicrosoftGraphOrganizationalBrandingProperties) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *MicrosoftGraphOrganizationalBrandingProperties) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *MicrosoftGraphOrganizationalBrandingProperties) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *MicrosoftGraphOrganizationalBrandingProperties) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### SetBackgroundColorNil

`func (o *MicrosoftGraphOrganizationalBrandingProperties) SetBackgroundColorNil(b bool)`

 SetBackgroundColorNil sets the value for BackgroundColor to be an explicit nil

### UnsetBackgroundColor
`func (o *MicrosoftGraphOrganizationalBrandingProperties) UnsetBackgroundColor()`

UnsetBackgroundColor ensures that no value is present for BackgroundColor, not even an explicit nil
### GetBackgroundImage

`func (o *MicrosoftGraphOrganizationalBrandingProperties) GetBackgroundImage() string`

GetBackgroundImage returns the BackgroundImage field if non-nil, zero value otherwise.

### GetBackgroundImageOk

`func (o *MicrosoftGraphOrganizationalBrandingProperties) GetBackgroundImageOk() (*string, bool)`

GetBackgroundImageOk returns a tuple with the BackgroundImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundImage

`func (o *MicrosoftGraphOrganizationalBrandingProperties) SetBackgroundImage(v string)`

SetBackgroundImage sets BackgroundImage field to given value.

### HasBackgroundImage

`func (o *MicrosoftGraphOrganizationalBrandingProperties) HasBackgroundImage() bool`

HasBackgroundImage returns a boolean if a field has been set.

### SetBackgroundImageNil

`func (o *MicrosoftGraphOrganizationalBrandingProperties) SetBackgroundImageNil(b bool)`

 SetBackgroundImageNil sets the value for BackgroundImage to be an explicit nil

### UnsetBackgroundImage
`func (o *MicrosoftGraphOrganizationalBrandingProperties) UnsetBackgroundImage()`

UnsetBackgroundImage ensures that no value is present for BackgroundImage, not even an explicit nil
### GetBackgroundImageRelativeUrl

`func (o *MicrosoftGraphOrganizationalBrandingProperties) GetBackgroundImageRelativeUrl() string`

GetBackgroundImageRelativeUrl returns the BackgroundImageRelativeUrl field if non-nil, zero value otherwise.

### GetBackgroundImageRelativeUrlOk

`func (o *MicrosoftGraphOrganizationalBrandingProperties) GetBackgroundImageRelativeUrlOk() (*string, bool)`

GetBackgroundImageRelativeUrlOk returns a tuple with the BackgroundImageRelativeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundImageRelativeUrl

`func (o *MicrosoftGraphOrganizationalBrandingProperties) SetBackgroundImageRelativeUrl(v string)`

SetBackgroundImageRelativeUrl sets BackgroundImageRelativeUrl field to given value.

### HasBackgroundImageRelativeUrl

`func (o *MicrosoftGraphOrganizationalBrandingProperties) HasBackgroundImageRelativeUrl() bool`

HasBackgroundImageRelativeUrl returns a boolean if a field has been set.

### SetBackgroundImageRelativeUrlNil

`func (o *MicrosoftGraphOrganizationalBrandingProperties) SetBackgroundImageRelativeUrlNil(b bool)`

 SetBackgroundImageRelativeUrlNil sets the value for BackgroundImageRelativeUrl to be an explicit nil

### UnsetBackgroundImageRelativeUrl
`func (o *MicrosoftGraphOrganizationalBrandingProperties) UnsetBackgroundImageRelativeUrl()`

UnsetBackgroundImageRelativeUrl ensures that no value is present for BackgroundImageRelativeUrl, not even an explicit nil
### GetBannerLogo

`func (o *MicrosoftGraphOrganizationalBrandingProperties) GetBannerLogo() string`

GetBannerLogo returns the BannerLogo field if non-nil, zero value otherwise.

### GetBannerLogoOk

`func (o *MicrosoftGraphOrganizationalBrandingProperties) GetBannerLogoOk() (*string, bool)`

GetBannerLogoOk returns a tuple with the BannerLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerLogo

`func (o *MicrosoftGraphOrganizationalBrandingProperties) SetBannerLogo(v string)`

SetBannerLogo sets BannerLogo field to given value.

### HasBannerLogo

`func (o *MicrosoftGraphOrganizationalBrandingProperties) HasBannerLogo() bool`

HasBannerLogo returns a boolean if a field has been set.

### SetBannerLogoNil

`func (o *MicrosoftGraphOrganizationalBrandingProperties) SetBannerLogoNil(b bool)`

 SetBannerLogoNil sets the value for BannerLogo to be an explicit nil

### UnsetBannerLogo
`func (o *MicrosoftGraphOrganizationalBrandingProperties) UnsetBannerLogo()`

UnsetBannerLogo ensures that no value is present for BannerLogo, not even an explicit nil
### GetBannerLogoRelativeUrl

`func (o *MicrosoftGraphOrganizationalBrandingProperties) GetBannerLogoRelativeUrl() string`

GetBannerLogoRelativeUrl returns the BannerLogoRelativeUrl field if non-nil, zero value otherwise.

### GetBannerLogoRelativeUrlOk

`func (o *MicrosoftGraphOrganizationalBrandingProperties) GetBannerLogoRelativeUrlOk() (*string, bool)`

GetBannerLogoRelativeUrlOk returns a tuple with the BannerLogoRelativeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerLogoRelativeUrl

`func (o *MicrosoftGraphOrganizationalBrandingProperties) SetBannerLogoRelativeUrl(v string)`

SetBannerLogoRelativeUrl sets BannerLogoRelativeUrl field to given value.

### HasBannerLogoRelativeUrl

`func (o *MicrosoftGraphOrganizationalBrandingProperties) HasBannerLogoRelativeUrl() bool`

HasBannerLogoRelativeUrl returns a boolean if a field has been set.

### SetBannerLogoRelativeUrlNil

`func (o *MicrosoftGraphOrganizationalBrandingProperties) SetBannerLogoRelativeUrlNil(b bool)`

 SetBannerLogoRelativeUrlNil sets the value for BannerLogoRelativeUrl to be an explicit nil

### UnsetBannerLogoRelativeUrl
`func (o *MicrosoftGraphOrganizationalBrandingProperties) UnsetBannerLogoRelativeUrl()`

UnsetBannerLogoRelativeUrl ensures that no value is present for BannerLogoRelativeUrl, not even an explicit nil
### GetCdnList

`func (o *MicrosoftGraphOrganizationalBrandingProperties) GetCdnList() []*string`

GetCdnList returns the CdnList field if non-nil, zero value otherwise.

### GetCdnListOk

`func (o *MicrosoftGraphOrganizationalBrandingProperties) GetCdnListOk() (*[]*string, bool)`

GetCdnListOk returns a tuple with the CdnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnList

`func (o *MicrosoftGraphOrganizationalBrandingProperties) SetCdnList(v []*string)`

SetCdnList sets CdnList field to given value.

### HasCdnList

`func (o *MicrosoftGraphOrganizationalBrandingProperties) HasCdnList() bool`

HasCdnList returns a boolean if a field has been set.

### GetSignInPageText

`func (o *MicrosoftGraphOrganizationalBrandingProperties) GetSignInPageText() string`

GetSignInPageText returns the SignInPageText field if non-nil, zero value otherwise.

### GetSignInPageTextOk

`func (o *MicrosoftGraphOrganizationalBrandingProperties) GetSignInPageTextOk() (*string, bool)`

GetSignInPageTextOk returns a tuple with the SignInPageText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInPageText

`func (o *MicrosoftGraphOrganizationalBrandingProperties) SetSignInPageText(v string)`

SetSignInPageText sets SignInPageText field to given value.

### HasSignInPageText

`func (o *MicrosoftGraphOrganizationalBrandingProperties) HasSignInPageText() bool`

HasSignInPageText returns a boolean if a field has been set.

### SetSignInPageTextNil

`func (o *MicrosoftGraphOrganizationalBrandingProperties) SetSignInPageTextNil(b bool)`

 SetSignInPageTextNil sets the value for SignInPageText to be an explicit nil

### UnsetSignInPageText
`func (o *MicrosoftGraphOrganizationalBrandingProperties) UnsetSignInPageText()`

UnsetSignInPageText ensures that no value is present for SignInPageText, not even an explicit nil
### GetSquareLogo

`func (o *MicrosoftGraphOrganizationalBrandingProperties) GetSquareLogo() string`

GetSquareLogo returns the SquareLogo field if non-nil, zero value otherwise.

### GetSquareLogoOk

`func (o *MicrosoftGraphOrganizationalBrandingProperties) GetSquareLogoOk() (*string, bool)`

GetSquareLogoOk returns a tuple with the SquareLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquareLogo

`func (o *MicrosoftGraphOrganizationalBrandingProperties) SetSquareLogo(v string)`

SetSquareLogo sets SquareLogo field to given value.

### HasSquareLogo

`func (o *MicrosoftGraphOrganizationalBrandingProperties) HasSquareLogo() bool`

HasSquareLogo returns a boolean if a field has been set.

### SetSquareLogoNil

`func (o *MicrosoftGraphOrganizationalBrandingProperties) SetSquareLogoNil(b bool)`

 SetSquareLogoNil sets the value for SquareLogo to be an explicit nil

### UnsetSquareLogo
`func (o *MicrosoftGraphOrganizationalBrandingProperties) UnsetSquareLogo()`

UnsetSquareLogo ensures that no value is present for SquareLogo, not even an explicit nil
### GetSquareLogoRelativeUrl

`func (o *MicrosoftGraphOrganizationalBrandingProperties) GetSquareLogoRelativeUrl() string`

GetSquareLogoRelativeUrl returns the SquareLogoRelativeUrl field if non-nil, zero value otherwise.

### GetSquareLogoRelativeUrlOk

`func (o *MicrosoftGraphOrganizationalBrandingProperties) GetSquareLogoRelativeUrlOk() (*string, bool)`

GetSquareLogoRelativeUrlOk returns a tuple with the SquareLogoRelativeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquareLogoRelativeUrl

`func (o *MicrosoftGraphOrganizationalBrandingProperties) SetSquareLogoRelativeUrl(v string)`

SetSquareLogoRelativeUrl sets SquareLogoRelativeUrl field to given value.

### HasSquareLogoRelativeUrl

`func (o *MicrosoftGraphOrganizationalBrandingProperties) HasSquareLogoRelativeUrl() bool`

HasSquareLogoRelativeUrl returns a boolean if a field has been set.

### SetSquareLogoRelativeUrlNil

`func (o *MicrosoftGraphOrganizationalBrandingProperties) SetSquareLogoRelativeUrlNil(b bool)`

 SetSquareLogoRelativeUrlNil sets the value for SquareLogoRelativeUrl to be an explicit nil

### UnsetSquareLogoRelativeUrl
`func (o *MicrosoftGraphOrganizationalBrandingProperties) UnsetSquareLogoRelativeUrl()`

UnsetSquareLogoRelativeUrl ensures that no value is present for SquareLogoRelativeUrl, not even an explicit nil
### GetUsernameHintText

`func (o *MicrosoftGraphOrganizationalBrandingProperties) GetUsernameHintText() string`

GetUsernameHintText returns the UsernameHintText field if non-nil, zero value otherwise.

### GetUsernameHintTextOk

`func (o *MicrosoftGraphOrganizationalBrandingProperties) GetUsernameHintTextOk() (*string, bool)`

GetUsernameHintTextOk returns a tuple with the UsernameHintText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameHintText

`func (o *MicrosoftGraphOrganizationalBrandingProperties) SetUsernameHintText(v string)`

SetUsernameHintText sets UsernameHintText field to given value.

### HasUsernameHintText

`func (o *MicrosoftGraphOrganizationalBrandingProperties) HasUsernameHintText() bool`

HasUsernameHintText returns a boolean if a field has been set.

### SetUsernameHintTextNil

`func (o *MicrosoftGraphOrganizationalBrandingProperties) SetUsernameHintTextNil(b bool)`

 SetUsernameHintTextNil sets the value for UsernameHintText to be an explicit nil

### UnsetUsernameHintText
`func (o *MicrosoftGraphOrganizationalBrandingProperties) UnsetUsernameHintText()`

UnsetUsernameHintText ensures that no value is present for UsernameHintText, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


