# BookCover

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Isbn13** | Pointer to **NullableString** |  | [optional] 
**Covers** | Pointer to **[]string** |  | [optional] 
**Sizes** | Pointer to [**[]CoverSize**](CoverSize.md) |  | [optional] 

## Methods

### NewBookCover

`func NewBookCover() *BookCover`

NewBookCover instantiates a new BookCover object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookCoverWithDefaults

`func NewBookCoverWithDefaults() *BookCover`

NewBookCoverWithDefaults instantiates a new BookCover object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsbn13

`func (o *BookCover) GetIsbn13() string`

GetIsbn13 returns the Isbn13 field if non-nil, zero value otherwise.

### GetIsbn13Ok

`func (o *BookCover) GetIsbn13Ok() (*string, bool)`

GetIsbn13Ok returns a tuple with the Isbn13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbn13

`func (o *BookCover) SetIsbn13(v string)`

SetIsbn13 sets Isbn13 field to given value.

### HasIsbn13

`func (o *BookCover) HasIsbn13() bool`

HasIsbn13 returns a boolean if a field has been set.

### SetIsbn13Nil

`func (o *BookCover) SetIsbn13Nil(b bool)`

 SetIsbn13Nil sets the value for Isbn13 to be an explicit nil

### UnsetIsbn13
`func (o *BookCover) UnsetIsbn13()`

UnsetIsbn13 ensures that no value is present for Isbn13, not even an explicit nil
### GetCovers

`func (o *BookCover) GetCovers() []string`

GetCovers returns the Covers field if non-nil, zero value otherwise.

### GetCoversOk

`func (o *BookCover) GetCoversOk() (*[]string, bool)`

GetCoversOk returns a tuple with the Covers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCovers

`func (o *BookCover) SetCovers(v []string)`

SetCovers sets Covers field to given value.

### HasCovers

`func (o *BookCover) HasCovers() bool`

HasCovers returns a boolean if a field has been set.

### SetCoversNil

`func (o *BookCover) SetCoversNil(b bool)`

 SetCoversNil sets the value for Covers to be an explicit nil

### UnsetCovers
`func (o *BookCover) UnsetCovers()`

UnsetCovers ensures that no value is present for Covers, not even an explicit nil
### GetSizes

`func (o *BookCover) GetSizes() []CoverSize`

GetSizes returns the Sizes field if non-nil, zero value otherwise.

### GetSizesOk

`func (o *BookCover) GetSizesOk() (*[]CoverSize, bool)`

GetSizesOk returns a tuple with the Sizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizes

`func (o *BookCover) SetSizes(v []CoverSize)`

SetSizes sets Sizes field to given value.

### HasSizes

`func (o *BookCover) HasSizes() bool`

HasSizes returns a boolean if a field has been set.

### SetSizesNil

`func (o *BookCover) SetSizesNil(b bool)`

 SetSizesNil sets the value for Sizes to be an explicit nil

### UnsetSizes
`func (o *BookCover) UnsetSizes()`

UnsetSizes ensures that no value is present for Sizes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


