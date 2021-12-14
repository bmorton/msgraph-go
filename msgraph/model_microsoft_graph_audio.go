/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MicrosoftGraphAudio struct for MicrosoftGraphAudio
type MicrosoftGraphAudio struct {
	// The title of the album for this audio file.
	Album NullableString `json:"album,omitempty"`
	// The artist named on the album for the audio file.
	AlbumArtist NullableString `json:"albumArtist,omitempty"`
	// The performing artist for the audio file.
	Artist NullableString `json:"artist,omitempty"`
	// Bitrate expressed in kbps.
	Bitrate NullableInt64 `json:"bitrate,omitempty"`
	// The name of the composer of the audio file.
	Composers NullableString `json:"composers,omitempty"`
	// Copyright information for the audio file.
	Copyright NullableString `json:"copyright,omitempty"`
	// The number of the disc this audio file came from.
	Disc NullableInt32 `json:"disc,omitempty"`
	// The total number of discs in this album.
	DiscCount NullableInt32 `json:"discCount,omitempty"`
	// Duration of the audio file, expressed in milliseconds
	Duration NullableInt64 `json:"duration,omitempty"`
	// The genre of this audio file.
	Genre NullableString `json:"genre,omitempty"`
	// Indicates if the file is protected with digital rights management.
	HasDrm NullableBool `json:"hasDrm,omitempty"`
	// Indicates if the file is encoded with a variable bitrate.
	IsVariableBitrate NullableBool `json:"isVariableBitrate,omitempty"`
	// The title of the audio file.
	Title NullableString `json:"title,omitempty"`
	// The number of the track on the original disc for this audio file.
	Track NullableInt32 `json:"track,omitempty"`
	// The total number of tracks on the original disc for this audio file.
	TrackCount NullableInt32 `json:"trackCount,omitempty"`
	// The year the audio file was recorded.
	Year NullableInt32 `json:"year,omitempty"`
}

// NewMicrosoftGraphAudio instantiates a new MicrosoftGraphAudio object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphAudio() *MicrosoftGraphAudio {
	this := MicrosoftGraphAudio{}
	return &this
}

// NewMicrosoftGraphAudioWithDefaults instantiates a new MicrosoftGraphAudio object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphAudioWithDefaults() *MicrosoftGraphAudio {
	this := MicrosoftGraphAudio{}
	return &this
}

// GetAlbum returns the Album field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAudio) GetAlbum() string {
	if o == nil || o.Album.Get() == nil {
		var ret string
		return ret
	}
	return *o.Album.Get()
}

// GetAlbumOk returns a tuple with the Album field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAudio) GetAlbumOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Album.Get(), o.Album.IsSet()
}

// HasAlbum returns a boolean if a field has been set.
func (o *MicrosoftGraphAudio) HasAlbum() bool {
	if o != nil && o.Album.IsSet() {
		return true
	}

	return false
}

// SetAlbum gets a reference to the given NullableString and assigns it to the Album field.
func (o *MicrosoftGraphAudio) SetAlbum(v string) {
	o.Album.Set(&v)
}
// SetAlbumNil sets the value for Album to be an explicit nil
func (o *MicrosoftGraphAudio) SetAlbumNil() {
	o.Album.Set(nil)
}

// UnsetAlbum ensures that no value is present for Album, not even an explicit nil
func (o *MicrosoftGraphAudio) UnsetAlbum() {
	o.Album.Unset()
}

// GetAlbumArtist returns the AlbumArtist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAudio) GetAlbumArtist() string {
	if o == nil || o.AlbumArtist.Get() == nil {
		var ret string
		return ret
	}
	return *o.AlbumArtist.Get()
}

// GetAlbumArtistOk returns a tuple with the AlbumArtist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAudio) GetAlbumArtistOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AlbumArtist.Get(), o.AlbumArtist.IsSet()
}

// HasAlbumArtist returns a boolean if a field has been set.
func (o *MicrosoftGraphAudio) HasAlbumArtist() bool {
	if o != nil && o.AlbumArtist.IsSet() {
		return true
	}

	return false
}

// SetAlbumArtist gets a reference to the given NullableString and assigns it to the AlbumArtist field.
func (o *MicrosoftGraphAudio) SetAlbumArtist(v string) {
	o.AlbumArtist.Set(&v)
}
// SetAlbumArtistNil sets the value for AlbumArtist to be an explicit nil
func (o *MicrosoftGraphAudio) SetAlbumArtistNil() {
	o.AlbumArtist.Set(nil)
}

// UnsetAlbumArtist ensures that no value is present for AlbumArtist, not even an explicit nil
func (o *MicrosoftGraphAudio) UnsetAlbumArtist() {
	o.AlbumArtist.Unset()
}

// GetArtist returns the Artist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAudio) GetArtist() string {
	if o == nil || o.Artist.Get() == nil {
		var ret string
		return ret
	}
	return *o.Artist.Get()
}

// GetArtistOk returns a tuple with the Artist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAudio) GetArtistOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Artist.Get(), o.Artist.IsSet()
}

// HasArtist returns a boolean if a field has been set.
func (o *MicrosoftGraphAudio) HasArtist() bool {
	if o != nil && o.Artist.IsSet() {
		return true
	}

	return false
}

// SetArtist gets a reference to the given NullableString and assigns it to the Artist field.
func (o *MicrosoftGraphAudio) SetArtist(v string) {
	o.Artist.Set(&v)
}
// SetArtistNil sets the value for Artist to be an explicit nil
func (o *MicrosoftGraphAudio) SetArtistNil() {
	o.Artist.Set(nil)
}

// UnsetArtist ensures that no value is present for Artist, not even an explicit nil
func (o *MicrosoftGraphAudio) UnsetArtist() {
	o.Artist.Unset()
}

// GetBitrate returns the Bitrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAudio) GetBitrate() int64 {
	if o == nil || o.Bitrate.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Bitrate.Get()
}

// GetBitrateOk returns a tuple with the Bitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAudio) GetBitrateOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Bitrate.Get(), o.Bitrate.IsSet()
}

// HasBitrate returns a boolean if a field has been set.
func (o *MicrosoftGraphAudio) HasBitrate() bool {
	if o != nil && o.Bitrate.IsSet() {
		return true
	}

	return false
}

// SetBitrate gets a reference to the given NullableInt64 and assigns it to the Bitrate field.
func (o *MicrosoftGraphAudio) SetBitrate(v int64) {
	o.Bitrate.Set(&v)
}
// SetBitrateNil sets the value for Bitrate to be an explicit nil
func (o *MicrosoftGraphAudio) SetBitrateNil() {
	o.Bitrate.Set(nil)
}

// UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
func (o *MicrosoftGraphAudio) UnsetBitrate() {
	o.Bitrate.Unset()
}

// GetComposers returns the Composers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAudio) GetComposers() string {
	if o == nil || o.Composers.Get() == nil {
		var ret string
		return ret
	}
	return *o.Composers.Get()
}

// GetComposersOk returns a tuple with the Composers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAudio) GetComposersOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Composers.Get(), o.Composers.IsSet()
}

// HasComposers returns a boolean if a field has been set.
func (o *MicrosoftGraphAudio) HasComposers() bool {
	if o != nil && o.Composers.IsSet() {
		return true
	}

	return false
}

// SetComposers gets a reference to the given NullableString and assigns it to the Composers field.
func (o *MicrosoftGraphAudio) SetComposers(v string) {
	o.Composers.Set(&v)
}
// SetComposersNil sets the value for Composers to be an explicit nil
func (o *MicrosoftGraphAudio) SetComposersNil() {
	o.Composers.Set(nil)
}

// UnsetComposers ensures that no value is present for Composers, not even an explicit nil
func (o *MicrosoftGraphAudio) UnsetComposers() {
	o.Composers.Unset()
}

// GetCopyright returns the Copyright field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAudio) GetCopyright() string {
	if o == nil || o.Copyright.Get() == nil {
		var ret string
		return ret
	}
	return *o.Copyright.Get()
}

// GetCopyrightOk returns a tuple with the Copyright field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAudio) GetCopyrightOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Copyright.Get(), o.Copyright.IsSet()
}

// HasCopyright returns a boolean if a field has been set.
func (o *MicrosoftGraphAudio) HasCopyright() bool {
	if o != nil && o.Copyright.IsSet() {
		return true
	}

	return false
}

// SetCopyright gets a reference to the given NullableString and assigns it to the Copyright field.
func (o *MicrosoftGraphAudio) SetCopyright(v string) {
	o.Copyright.Set(&v)
}
// SetCopyrightNil sets the value for Copyright to be an explicit nil
func (o *MicrosoftGraphAudio) SetCopyrightNil() {
	o.Copyright.Set(nil)
}

// UnsetCopyright ensures that no value is present for Copyright, not even an explicit nil
func (o *MicrosoftGraphAudio) UnsetCopyright() {
	o.Copyright.Unset()
}

// GetDisc returns the Disc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAudio) GetDisc() int32 {
	if o == nil || o.Disc.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Disc.Get()
}

// GetDiscOk returns a tuple with the Disc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAudio) GetDiscOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Disc.Get(), o.Disc.IsSet()
}

// HasDisc returns a boolean if a field has been set.
func (o *MicrosoftGraphAudio) HasDisc() bool {
	if o != nil && o.Disc.IsSet() {
		return true
	}

	return false
}

// SetDisc gets a reference to the given NullableInt32 and assigns it to the Disc field.
func (o *MicrosoftGraphAudio) SetDisc(v int32) {
	o.Disc.Set(&v)
}
// SetDiscNil sets the value for Disc to be an explicit nil
func (o *MicrosoftGraphAudio) SetDiscNil() {
	o.Disc.Set(nil)
}

// UnsetDisc ensures that no value is present for Disc, not even an explicit nil
func (o *MicrosoftGraphAudio) UnsetDisc() {
	o.Disc.Unset()
}

// GetDiscCount returns the DiscCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAudio) GetDiscCount() int32 {
	if o == nil || o.DiscCount.Get() == nil {
		var ret int32
		return ret
	}
	return *o.DiscCount.Get()
}

// GetDiscCountOk returns a tuple with the DiscCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAudio) GetDiscCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DiscCount.Get(), o.DiscCount.IsSet()
}

// HasDiscCount returns a boolean if a field has been set.
func (o *MicrosoftGraphAudio) HasDiscCount() bool {
	if o != nil && o.DiscCount.IsSet() {
		return true
	}

	return false
}

// SetDiscCount gets a reference to the given NullableInt32 and assigns it to the DiscCount field.
func (o *MicrosoftGraphAudio) SetDiscCount(v int32) {
	o.DiscCount.Set(&v)
}
// SetDiscCountNil sets the value for DiscCount to be an explicit nil
func (o *MicrosoftGraphAudio) SetDiscCountNil() {
	o.DiscCount.Set(nil)
}

// UnsetDiscCount ensures that no value is present for DiscCount, not even an explicit nil
func (o *MicrosoftGraphAudio) UnsetDiscCount() {
	o.DiscCount.Unset()
}

// GetDuration returns the Duration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAudio) GetDuration() int64 {
	if o == nil || o.Duration.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Duration.Get()
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAudio) GetDurationOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Duration.Get(), o.Duration.IsSet()
}

// HasDuration returns a boolean if a field has been set.
func (o *MicrosoftGraphAudio) HasDuration() bool {
	if o != nil && o.Duration.IsSet() {
		return true
	}

	return false
}

// SetDuration gets a reference to the given NullableInt64 and assigns it to the Duration field.
func (o *MicrosoftGraphAudio) SetDuration(v int64) {
	o.Duration.Set(&v)
}
// SetDurationNil sets the value for Duration to be an explicit nil
func (o *MicrosoftGraphAudio) SetDurationNil() {
	o.Duration.Set(nil)
}

// UnsetDuration ensures that no value is present for Duration, not even an explicit nil
func (o *MicrosoftGraphAudio) UnsetDuration() {
	o.Duration.Unset()
}

// GetGenre returns the Genre field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAudio) GetGenre() string {
	if o == nil || o.Genre.Get() == nil {
		var ret string
		return ret
	}
	return *o.Genre.Get()
}

// GetGenreOk returns a tuple with the Genre field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAudio) GetGenreOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Genre.Get(), o.Genre.IsSet()
}

// HasGenre returns a boolean if a field has been set.
func (o *MicrosoftGraphAudio) HasGenre() bool {
	if o != nil && o.Genre.IsSet() {
		return true
	}

	return false
}

// SetGenre gets a reference to the given NullableString and assigns it to the Genre field.
func (o *MicrosoftGraphAudio) SetGenre(v string) {
	o.Genre.Set(&v)
}
// SetGenreNil sets the value for Genre to be an explicit nil
func (o *MicrosoftGraphAudio) SetGenreNil() {
	o.Genre.Set(nil)
}

// UnsetGenre ensures that no value is present for Genre, not even an explicit nil
func (o *MicrosoftGraphAudio) UnsetGenre() {
	o.Genre.Unset()
}

// GetHasDrm returns the HasDrm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAudio) GetHasDrm() bool {
	if o == nil || o.HasDrm.Get() == nil {
		var ret bool
		return ret
	}
	return *o.HasDrm.Get()
}

// GetHasDrmOk returns a tuple with the HasDrm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAudio) GetHasDrmOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HasDrm.Get(), o.HasDrm.IsSet()
}

// HasHasDrm returns a boolean if a field has been set.
func (o *MicrosoftGraphAudio) HasHasDrm() bool {
	if o != nil && o.HasDrm.IsSet() {
		return true
	}

	return false
}

// SetHasDrm gets a reference to the given NullableBool and assigns it to the HasDrm field.
func (o *MicrosoftGraphAudio) SetHasDrm(v bool) {
	o.HasDrm.Set(&v)
}
// SetHasDrmNil sets the value for HasDrm to be an explicit nil
func (o *MicrosoftGraphAudio) SetHasDrmNil() {
	o.HasDrm.Set(nil)
}

// UnsetHasDrm ensures that no value is present for HasDrm, not even an explicit nil
func (o *MicrosoftGraphAudio) UnsetHasDrm() {
	o.HasDrm.Unset()
}

// GetIsVariableBitrate returns the IsVariableBitrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAudio) GetIsVariableBitrate() bool {
	if o == nil || o.IsVariableBitrate.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsVariableBitrate.Get()
}

// GetIsVariableBitrateOk returns a tuple with the IsVariableBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAudio) GetIsVariableBitrateOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsVariableBitrate.Get(), o.IsVariableBitrate.IsSet()
}

// HasIsVariableBitrate returns a boolean if a field has been set.
func (o *MicrosoftGraphAudio) HasIsVariableBitrate() bool {
	if o != nil && o.IsVariableBitrate.IsSet() {
		return true
	}

	return false
}

// SetIsVariableBitrate gets a reference to the given NullableBool and assigns it to the IsVariableBitrate field.
func (o *MicrosoftGraphAudio) SetIsVariableBitrate(v bool) {
	o.IsVariableBitrate.Set(&v)
}
// SetIsVariableBitrateNil sets the value for IsVariableBitrate to be an explicit nil
func (o *MicrosoftGraphAudio) SetIsVariableBitrateNil() {
	o.IsVariableBitrate.Set(nil)
}

// UnsetIsVariableBitrate ensures that no value is present for IsVariableBitrate, not even an explicit nil
func (o *MicrosoftGraphAudio) UnsetIsVariableBitrate() {
	o.IsVariableBitrate.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAudio) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAudio) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *MicrosoftGraphAudio) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *MicrosoftGraphAudio) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *MicrosoftGraphAudio) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *MicrosoftGraphAudio) UnsetTitle() {
	o.Title.Unset()
}

// GetTrack returns the Track field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAudio) GetTrack() int32 {
	if o == nil || o.Track.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Track.Get()
}

// GetTrackOk returns a tuple with the Track field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAudio) GetTrackOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Track.Get(), o.Track.IsSet()
}

// HasTrack returns a boolean if a field has been set.
func (o *MicrosoftGraphAudio) HasTrack() bool {
	if o != nil && o.Track.IsSet() {
		return true
	}

	return false
}

// SetTrack gets a reference to the given NullableInt32 and assigns it to the Track field.
func (o *MicrosoftGraphAudio) SetTrack(v int32) {
	o.Track.Set(&v)
}
// SetTrackNil sets the value for Track to be an explicit nil
func (o *MicrosoftGraphAudio) SetTrackNil() {
	o.Track.Set(nil)
}

// UnsetTrack ensures that no value is present for Track, not even an explicit nil
func (o *MicrosoftGraphAudio) UnsetTrack() {
	o.Track.Unset()
}

// GetTrackCount returns the TrackCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAudio) GetTrackCount() int32 {
	if o == nil || o.TrackCount.Get() == nil {
		var ret int32
		return ret
	}
	return *o.TrackCount.Get()
}

// GetTrackCountOk returns a tuple with the TrackCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAudio) GetTrackCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TrackCount.Get(), o.TrackCount.IsSet()
}

// HasTrackCount returns a boolean if a field has been set.
func (o *MicrosoftGraphAudio) HasTrackCount() bool {
	if o != nil && o.TrackCount.IsSet() {
		return true
	}

	return false
}

// SetTrackCount gets a reference to the given NullableInt32 and assigns it to the TrackCount field.
func (o *MicrosoftGraphAudio) SetTrackCount(v int32) {
	o.TrackCount.Set(&v)
}
// SetTrackCountNil sets the value for TrackCount to be an explicit nil
func (o *MicrosoftGraphAudio) SetTrackCountNil() {
	o.TrackCount.Set(nil)
}

// UnsetTrackCount ensures that no value is present for TrackCount, not even an explicit nil
func (o *MicrosoftGraphAudio) UnsetTrackCount() {
	o.TrackCount.Unset()
}

// GetYear returns the Year field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAudio) GetYear() int32 {
	if o == nil || o.Year.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Year.Get()
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAudio) GetYearOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Year.Get(), o.Year.IsSet()
}

// HasYear returns a boolean if a field has been set.
func (o *MicrosoftGraphAudio) HasYear() bool {
	if o != nil && o.Year.IsSet() {
		return true
	}

	return false
}

// SetYear gets a reference to the given NullableInt32 and assigns it to the Year field.
func (o *MicrosoftGraphAudio) SetYear(v int32) {
	o.Year.Set(&v)
}
// SetYearNil sets the value for Year to be an explicit nil
func (o *MicrosoftGraphAudio) SetYearNil() {
	o.Year.Set(nil)
}

// UnsetYear ensures that no value is present for Year, not even an explicit nil
func (o *MicrosoftGraphAudio) UnsetYear() {
	o.Year.Unset()
}

func (o MicrosoftGraphAudio) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Album.IsSet() {
		toSerialize["album"] = o.Album.Get()
	}
	if o.AlbumArtist.IsSet() {
		toSerialize["albumArtist"] = o.AlbumArtist.Get()
	}
	if o.Artist.IsSet() {
		toSerialize["artist"] = o.Artist.Get()
	}
	if o.Bitrate.IsSet() {
		toSerialize["bitrate"] = o.Bitrate.Get()
	}
	if o.Composers.IsSet() {
		toSerialize["composers"] = o.Composers.Get()
	}
	if o.Copyright.IsSet() {
		toSerialize["copyright"] = o.Copyright.Get()
	}
	if o.Disc.IsSet() {
		toSerialize["disc"] = o.Disc.Get()
	}
	if o.DiscCount.IsSet() {
		toSerialize["discCount"] = o.DiscCount.Get()
	}
	if o.Duration.IsSet() {
		toSerialize["duration"] = o.Duration.Get()
	}
	if o.Genre.IsSet() {
		toSerialize["genre"] = o.Genre.Get()
	}
	if o.HasDrm.IsSet() {
		toSerialize["hasDrm"] = o.HasDrm.Get()
	}
	if o.IsVariableBitrate.IsSet() {
		toSerialize["isVariableBitrate"] = o.IsVariableBitrate.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Track.IsSet() {
		toSerialize["track"] = o.Track.Get()
	}
	if o.TrackCount.IsSet() {
		toSerialize["trackCount"] = o.TrackCount.Get()
	}
	if o.Year.IsSet() {
		toSerialize["year"] = o.Year.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphAudio struct {
	value *MicrosoftGraphAudio
	isSet bool
}

func (v NullableMicrosoftGraphAudio) Get() *MicrosoftGraphAudio {
	return v.value
}

func (v *NullableMicrosoftGraphAudio) Set(val *MicrosoftGraphAudio) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAudio) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAudio) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAudio(val *MicrosoftGraphAudio) *NullableMicrosoftGraphAudio {
	return &NullableMicrosoftGraphAudio{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAudio) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAudio) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

