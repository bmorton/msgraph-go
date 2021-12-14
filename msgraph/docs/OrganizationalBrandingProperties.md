# OrganizationalBrandingProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewOrganizationalBrandingProperties

`func NewOrganizationalBrandingProperties() *OrganizationalBrandingProperties`

NewOrganizationalBrandingProperties instantiates a new OrganizationalBrandingProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationalBrandingPropertiesWithDefaults

`func NewOrganizationalBrandingPropertiesWithDefaults() *OrganizationalBrandingProperties`

NewOrganizationalBrandingPropertiesWithDefaults instantiates a new OrganizationalBrandingProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackgroundColor

`func (o *OrganizationalBrandingProperties) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *OrganizationalBrandingProperties) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *OrganizationalBrandingProperties) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *OrganizationalBrandingProperties) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### SetBackgroundColorNil

`func (o *OrganizationalBrandingProperties) SetBackgroundColorNil(b bool)`

 SetBackgroundColorNil sets the value for BackgroundColor to be an explicit nil

### UnsetBackgroundColor
`func (o *OrganizationalBrandingProperties) UnsetBackgroundColor()`

UnsetBackgroundColor ensures that no value is present for BackgroundColor, not even an explicit nil
### GetBackgroundImage

`func (o *OrganizationalBrandingProperties) GetBackgroundImage() string`

GetBackgroundImage returns the BackgroundImage field if non-nil, zero value otherwise.

### GetBackgroundImageOk

`func (o *OrganizationalBrandingProperties) GetBackgroundImageOk() (*string, bool)`

GetBackgroundImageOk returns a tuple with the BackgroundImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundImage

`func (o *OrganizationalBrandingProperties) SetBackgroundImage(v string)`

SetBackgroundImage sets BackgroundImage field to given value.

### HasBackgroundImage

`func (o *OrganizationalBrandingProperties) HasBackgroundImage() bool`

HasBackgroundImage returns a boolean if a field has been set.

### SetBackgroundImageNil

`func (o *OrganizationalBrandingProperties) SetBackgroundImageNil(b bool)`

 SetBackgroundImageNil sets the value for BackgroundImage to be an explicit nil

### UnsetBackgroundImage
`func (o *OrganizationalBrandingProperties) UnsetBackgroundImage()`

UnsetBackgroundImage ensures that no value is present for BackgroundImage, not even an explicit nil
### GetBackgroundImageRelativeUrl

`func (o *OrganizationalBrandingProperties) GetBackgroundImageRelativeUrl() string`

GetBackgroundImageRelativeUrl returns the BackgroundImageRelativeUrl field if non-nil, zero value otherwise.

### GetBackgroundImageRelativeUrlOk

`func (o *OrganizationalBrandingProperties) GetBackgroundImageRelativeUrlOk() (*string, bool)`

GetBackgroundImageRelativeUrlOk returns a tuple with the BackgroundImageRelativeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundImageRelativeUrl

`func (o *OrganizationalBrandingProperties) SetBackgroundImageRelativeUrl(v string)`

SetBackgroundImageRelativeUrl sets BackgroundImageRelativeUrl field to given value.

### HasBackgroundImageRelativeUrl

`func (o *OrganizationalBrandingProperties) HasBackgroundImageRelativeUrl() bool`

HasBackgroundImageRelativeUrl returns a boolean if a field has been set.

### SetBackgroundImageRelativeUrlNil

`func (o *OrganizationalBrandingProperties) SetBackgroundImageRelativeUrlNil(b bool)`

 SetBackgroundImageRelativeUrlNil sets the value for BackgroundImageRelativeUrl to be an explicit nil

### UnsetBackgroundImageRelativeUrl
`func (o *OrganizationalBrandingProperties) UnsetBackgroundImageRelativeUrl()`

UnsetBackgroundImageRelativeUrl ensures that no value is present for BackgroundImageRelativeUrl, not even an explicit nil
### GetBannerLogo

`func (o *OrganizationalBrandingProperties) GetBannerLogo() string`

GetBannerLogo returns the BannerLogo field if non-nil, zero value otherwise.

### GetBannerLogoOk

`func (o *OrganizationalBrandingProperties) GetBannerLogoOk() (*string, bool)`

GetBannerLogoOk returns a tuple with the BannerLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerLogo

`func (o *OrganizationalBrandingProperties) SetBannerLogo(v string)`

SetBannerLogo sets BannerLogo field to given value.

### HasBannerLogo

`func (o *OrganizationalBrandingProperties) HasBannerLogo() bool`

HasBannerLogo returns a boolean if a field has been set.

### SetBannerLogoNil

`func (o *OrganizationalBrandingProperties) SetBannerLogoNil(b bool)`

 SetBannerLogoNil sets the value for BannerLogo to be an explicit nil

### UnsetBannerLogo
`func (o *OrganizationalBrandingProperties) UnsetBannerLogo()`

UnsetBannerLogo ensures that no value is present for BannerLogo, not even an explicit nil
### GetBannerLogoRelativeUrl

`func (o *OrganizationalBrandingProperties) GetBannerLogoRelativeUrl() string`

GetBannerLogoRelativeUrl returns the BannerLogoRelativeUrl field if non-nil, zero value otherwise.

### GetBannerLogoRelativeUrlOk

`func (o *OrganizationalBrandingProperties) GetBannerLogoRelativeUrlOk() (*string, bool)`

GetBannerLogoRelativeUrlOk returns a tuple with the BannerLogoRelativeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerLogoRelativeUrl

`func (o *OrganizationalBrandingProperties) SetBannerLogoRelativeUrl(v string)`

SetBannerLogoRelativeUrl sets BannerLogoRelativeUrl field to given value.

### HasBannerLogoRelativeUrl

`func (o *OrganizationalBrandingProperties) HasBannerLogoRelativeUrl() bool`

HasBannerLogoRelativeUrl returns a boolean if a field has been set.

### SetBannerLogoRelativeUrlNil

`func (o *OrganizationalBrandingProperties) SetBannerLogoRelativeUrlNil(b bool)`

 SetBannerLogoRelativeUrlNil sets the value for BannerLogoRelativeUrl to be an explicit nil

### UnsetBannerLogoRelativeUrl
`func (o *OrganizationalBrandingProperties) UnsetBannerLogoRelativeUrl()`

UnsetBannerLogoRelativeUrl ensures that no value is present for BannerLogoRelativeUrl, not even an explicit nil
### GetCdnList

`func (o *OrganizationalBrandingProperties) GetCdnList() []*string`

GetCdnList returns the CdnList field if non-nil, zero value otherwise.

### GetCdnListOk

`func (o *OrganizationalBrandingProperties) GetCdnListOk() (*[]*string, bool)`

GetCdnListOk returns a tuple with the CdnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnList

`func (o *OrganizationalBrandingProperties) SetCdnList(v []*string)`

SetCdnList sets CdnList field to given value.

### HasCdnList

`func (o *OrganizationalBrandingProperties) HasCdnList() bool`

HasCdnList returns a boolean if a field has been set.

### GetSignInPageText

`func (o *OrganizationalBrandingProperties) GetSignInPageText() string`

GetSignInPageText returns the SignInPageText field if non-nil, zero value otherwise.

### GetSignInPageTextOk

`func (o *OrganizationalBrandingProperties) GetSignInPageTextOk() (*string, bool)`

GetSignInPageTextOk returns a tuple with the SignInPageText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInPageText

`func (o *OrganizationalBrandingProperties) SetSignInPageText(v string)`

SetSignInPageText sets SignInPageText field to given value.

### HasSignInPageText

`func (o *OrganizationalBrandingProperties) HasSignInPageText() bool`

HasSignInPageText returns a boolean if a field has been set.

### SetSignInPageTextNil

`func (o *OrganizationalBrandingProperties) SetSignInPageTextNil(b bool)`

 SetSignInPageTextNil sets the value for SignInPageText to be an explicit nil

### UnsetSignInPageText
`func (o *OrganizationalBrandingProperties) UnsetSignInPageText()`

UnsetSignInPageText ensures that no value is present for SignInPageText, not even an explicit nil
### GetSquareLogo

`func (o *OrganizationalBrandingProperties) GetSquareLogo() string`

GetSquareLogo returns the SquareLogo field if non-nil, zero value otherwise.

### GetSquareLogoOk

`func (o *OrganizationalBrandingProperties) GetSquareLogoOk() (*string, bool)`

GetSquareLogoOk returns a tuple with the SquareLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquareLogo

`func (o *OrganizationalBrandingProperties) SetSquareLogo(v string)`

SetSquareLogo sets SquareLogo field to given value.

### HasSquareLogo

`func (o *OrganizationalBrandingProperties) HasSquareLogo() bool`

HasSquareLogo returns a boolean if a field has been set.

### SetSquareLogoNil

`func (o *OrganizationalBrandingProperties) SetSquareLogoNil(b bool)`

 SetSquareLogoNil sets the value for SquareLogo to be an explicit nil

### UnsetSquareLogo
`func (o *OrganizationalBrandingProperties) UnsetSquareLogo()`

UnsetSquareLogo ensures that no value is present for SquareLogo, not even an explicit nil
### GetSquareLogoRelativeUrl

`func (o *OrganizationalBrandingProperties) GetSquareLogoRelativeUrl() string`

GetSquareLogoRelativeUrl returns the SquareLogoRelativeUrl field if non-nil, zero value otherwise.

### GetSquareLogoRelativeUrlOk

`func (o *OrganizationalBrandingProperties) GetSquareLogoRelativeUrlOk() (*string, bool)`

GetSquareLogoRelativeUrlOk returns a tuple with the SquareLogoRelativeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquareLogoRelativeUrl

`func (o *OrganizationalBrandingProperties) SetSquareLogoRelativeUrl(v string)`

SetSquareLogoRelativeUrl sets SquareLogoRelativeUrl field to given value.

### HasSquareLogoRelativeUrl

`func (o *OrganizationalBrandingProperties) HasSquareLogoRelativeUrl() bool`

HasSquareLogoRelativeUrl returns a boolean if a field has been set.

### SetSquareLogoRelativeUrlNil

`func (o *OrganizationalBrandingProperties) SetSquareLogoRelativeUrlNil(b bool)`

 SetSquareLogoRelativeUrlNil sets the value for SquareLogoRelativeUrl to be an explicit nil

### UnsetSquareLogoRelativeUrl
`func (o *OrganizationalBrandingProperties) UnsetSquareLogoRelativeUrl()`

UnsetSquareLogoRelativeUrl ensures that no value is present for SquareLogoRelativeUrl, not even an explicit nil
### GetUsernameHintText

`func (o *OrganizationalBrandingProperties) GetUsernameHintText() string`

GetUsernameHintText returns the UsernameHintText field if non-nil, zero value otherwise.

### GetUsernameHintTextOk

`func (o *OrganizationalBrandingProperties) GetUsernameHintTextOk() (*string, bool)`

GetUsernameHintTextOk returns a tuple with the UsernameHintText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameHintText

`func (o *OrganizationalBrandingProperties) SetUsernameHintText(v string)`

SetUsernameHintText sets UsernameHintText field to given value.

### HasUsernameHintText

`func (o *OrganizationalBrandingProperties) HasUsernameHintText() bool`

HasUsernameHintText returns a boolean if a field has been set.

### SetUsernameHintTextNil

`func (o *OrganizationalBrandingProperties) SetUsernameHintTextNil(b bool)`

 SetUsernameHintTextNil sets the value for UsernameHintText to be an explicit nil

### UnsetUsernameHintText
`func (o *OrganizationalBrandingProperties) UnsetUsernameHintText()`

UnsetUsernameHintText ensures that no value is present for UsernameHintText, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


