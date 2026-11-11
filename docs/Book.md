# Book

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** |  | [optional] [readonly] 
**Isbn10** | Pointer to **NullableString** |  | [optional] 
**Isbn13** | Pointer to **NullableString** |  | [optional] 
**WorkId** | Pointer to **NullableInt64** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Subtitle** | Pointer to **NullableString** |  | [optional] 
**NumberOfPages** | Pointer to **NullableInt64** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Publishers** | Pointer to [**[]EntityRef**](EntityRef.md) |  | [optional] 
**Imprint** | Pointer to [**[]EntityRef**](EntityRef.md) |  | [optional] 
**PublishDate** | Pointer to **NullableTime** |  | [optional] 
**PhysicalFormat** | Pointer to **NullableString** |  | [optional] 
**Covers** | Pointer to **[]string** |  | [optional] 
**Authors** | Pointer to [**[]EntityRef**](EntityRef.md) |  | [optional] 
**Contributors** | Pointer to [**[]Contributor**](Contributor.md) |  | [optional] 
**Subjects** | Pointer to [**[]SubjectRef**](SubjectRef.md) |  | [optional] 
**Series** | Pointer to [**NullableSeries**](Series.md) |  | [optional] 
**Genres** | Pointer to [**[]EntityRef**](EntityRef.md) |  | [optional] 
**OtherEditions** | Pointer to [**[]EditionRef**](EditionRef.md) |  | [optional] 
**PreviousEdition** | Pointer to [**NullableEditionRef**](EditionRef.md) |  | [optional] 
**NextEdition** | Pointer to [**NullableEditionRef**](EditionRef.md) |  | [optional] 
**PhysicalDimensions** | Pointer to **NullableString** |  | [optional] 
**Ratings** | Pointer to [**NullableRatings**](Ratings.md) |  | [optional] 
**Language** | Pointer to **NullableString** |  | [optional] 
**SaleStatus** | Pointer to **NullableString** |  | [optional] 
**Excerpt** | Pointer to **NullableString** |  | [optional] 
**ReviewQuotes** | Pointer to **NullableString** |  | [optional] 
**RetailLinks** | Pointer to [**[]RetailLink**](RetailLink.md) |  | [optional] 
**Categories** | Pointer to [**[]Category**](Category.md) |  | [optional] 
**AiSummary** | Pointer to **NullableString** |  | [optional] 
**CanonicalEditionId** | Pointer to **NullableInt64** |  | [optional] 
**LanguageCode** | Pointer to **NullableString** |  | [optional] 
**PublishDatePrecision** | Pointer to **NullableString** |  | [optional] 
**EstimatedReadTime** | Pointer to **NullableInt64** |  | [optional] 
**LanguagesCount** | Pointer to **NullableInt64** |  | [optional] 
**Link** | Pointer to **NullableString** |  | [optional] 
**IsbnModel** | Pointer to [**NullableIsbnDetails**](IsbnDetails.md) |  | [optional] 

## Methods

### NewBook

`func NewBook() *Book`

NewBook instantiates a new Book object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookWithDefaults

`func NewBookWithDefaults() *Book`

NewBookWithDefaults instantiates a new Book object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Book) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Book) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Book) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Book) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Book) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Book) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsbn10

`func (o *Book) GetIsbn10() string`

GetIsbn10 returns the Isbn10 field if non-nil, zero value otherwise.

### GetIsbn10Ok

`func (o *Book) GetIsbn10Ok() (*string, bool)`

GetIsbn10Ok returns a tuple with the Isbn10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbn10

`func (o *Book) SetIsbn10(v string)`

SetIsbn10 sets Isbn10 field to given value.

### HasIsbn10

`func (o *Book) HasIsbn10() bool`

HasIsbn10 returns a boolean if a field has been set.

### SetIsbn10Nil

`func (o *Book) SetIsbn10Nil(b bool)`

 SetIsbn10Nil sets the value for Isbn10 to be an explicit nil

### UnsetIsbn10
`func (o *Book) UnsetIsbn10()`

UnsetIsbn10 ensures that no value is present for Isbn10, not even an explicit nil
### GetIsbn13

`func (o *Book) GetIsbn13() string`

GetIsbn13 returns the Isbn13 field if non-nil, zero value otherwise.

### GetIsbn13Ok

`func (o *Book) GetIsbn13Ok() (*string, bool)`

GetIsbn13Ok returns a tuple with the Isbn13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbn13

`func (o *Book) SetIsbn13(v string)`

SetIsbn13 sets Isbn13 field to given value.

### HasIsbn13

`func (o *Book) HasIsbn13() bool`

HasIsbn13 returns a boolean if a field has been set.

### SetIsbn13Nil

`func (o *Book) SetIsbn13Nil(b bool)`

 SetIsbn13Nil sets the value for Isbn13 to be an explicit nil

### UnsetIsbn13
`func (o *Book) UnsetIsbn13()`

UnsetIsbn13 ensures that no value is present for Isbn13, not even an explicit nil
### GetWorkId

`func (o *Book) GetWorkId() int64`

GetWorkId returns the WorkId field if non-nil, zero value otherwise.

### GetWorkIdOk

`func (o *Book) GetWorkIdOk() (*int64, bool)`

GetWorkIdOk returns a tuple with the WorkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkId

`func (o *Book) SetWorkId(v int64)`

SetWorkId sets WorkId field to given value.

### HasWorkId

`func (o *Book) HasWorkId() bool`

HasWorkId returns a boolean if a field has been set.

### SetWorkIdNil

`func (o *Book) SetWorkIdNil(b bool)`

 SetWorkIdNil sets the value for WorkId to be an explicit nil

### UnsetWorkId
`func (o *Book) UnsetWorkId()`

UnsetWorkId ensures that no value is present for WorkId, not even an explicit nil
### GetTitle

`func (o *Book) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Book) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Book) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Book) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *Book) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *Book) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetSubtitle

`func (o *Book) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *Book) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *Book) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *Book) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### SetSubtitleNil

`func (o *Book) SetSubtitleNil(b bool)`

 SetSubtitleNil sets the value for Subtitle to be an explicit nil

### UnsetSubtitle
`func (o *Book) UnsetSubtitle()`

UnsetSubtitle ensures that no value is present for Subtitle, not even an explicit nil
### GetNumberOfPages

`func (o *Book) GetNumberOfPages() int64`

GetNumberOfPages returns the NumberOfPages field if non-nil, zero value otherwise.

### GetNumberOfPagesOk

`func (o *Book) GetNumberOfPagesOk() (*int64, bool)`

GetNumberOfPagesOk returns a tuple with the NumberOfPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPages

`func (o *Book) SetNumberOfPages(v int64)`

SetNumberOfPages sets NumberOfPages field to given value.

### HasNumberOfPages

`func (o *Book) HasNumberOfPages() bool`

HasNumberOfPages returns a boolean if a field has been set.

### SetNumberOfPagesNil

`func (o *Book) SetNumberOfPagesNil(b bool)`

 SetNumberOfPagesNil sets the value for NumberOfPages to be an explicit nil

### UnsetNumberOfPages
`func (o *Book) UnsetNumberOfPages()`

UnsetNumberOfPages ensures that no value is present for NumberOfPages, not even an explicit nil
### GetDescription

`func (o *Book) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Book) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Book) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Book) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Book) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Book) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPublishers

`func (o *Book) GetPublishers() []EntityRef`

GetPublishers returns the Publishers field if non-nil, zero value otherwise.

### GetPublishersOk

`func (o *Book) GetPublishersOk() (*[]EntityRef, bool)`

GetPublishersOk returns a tuple with the Publishers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishers

`func (o *Book) SetPublishers(v []EntityRef)`

SetPublishers sets Publishers field to given value.

### HasPublishers

`func (o *Book) HasPublishers() bool`

HasPublishers returns a boolean if a field has been set.

### SetPublishersNil

`func (o *Book) SetPublishersNil(b bool)`

 SetPublishersNil sets the value for Publishers to be an explicit nil

### UnsetPublishers
`func (o *Book) UnsetPublishers()`

UnsetPublishers ensures that no value is present for Publishers, not even an explicit nil
### GetImprint

`func (o *Book) GetImprint() []EntityRef`

GetImprint returns the Imprint field if non-nil, zero value otherwise.

### GetImprintOk

`func (o *Book) GetImprintOk() (*[]EntityRef, bool)`

GetImprintOk returns a tuple with the Imprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImprint

`func (o *Book) SetImprint(v []EntityRef)`

SetImprint sets Imprint field to given value.

### HasImprint

`func (o *Book) HasImprint() bool`

HasImprint returns a boolean if a field has been set.

### SetImprintNil

`func (o *Book) SetImprintNil(b bool)`

 SetImprintNil sets the value for Imprint to be an explicit nil

### UnsetImprint
`func (o *Book) UnsetImprint()`

UnsetImprint ensures that no value is present for Imprint, not even an explicit nil
### GetPublishDate

`func (o *Book) GetPublishDate() time.Time`

GetPublishDate returns the PublishDate field if non-nil, zero value otherwise.

### GetPublishDateOk

`func (o *Book) GetPublishDateOk() (*time.Time, bool)`

GetPublishDateOk returns a tuple with the PublishDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishDate

`func (o *Book) SetPublishDate(v time.Time)`

SetPublishDate sets PublishDate field to given value.

### HasPublishDate

`func (o *Book) HasPublishDate() bool`

HasPublishDate returns a boolean if a field has been set.

### SetPublishDateNil

`func (o *Book) SetPublishDateNil(b bool)`

 SetPublishDateNil sets the value for PublishDate to be an explicit nil

### UnsetPublishDate
`func (o *Book) UnsetPublishDate()`

UnsetPublishDate ensures that no value is present for PublishDate, not even an explicit nil
### GetPhysicalFormat

`func (o *Book) GetPhysicalFormat() string`

GetPhysicalFormat returns the PhysicalFormat field if non-nil, zero value otherwise.

### GetPhysicalFormatOk

`func (o *Book) GetPhysicalFormatOk() (*string, bool)`

GetPhysicalFormatOk returns a tuple with the PhysicalFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalFormat

`func (o *Book) SetPhysicalFormat(v string)`

SetPhysicalFormat sets PhysicalFormat field to given value.

### HasPhysicalFormat

`func (o *Book) HasPhysicalFormat() bool`

HasPhysicalFormat returns a boolean if a field has been set.

### SetPhysicalFormatNil

`func (o *Book) SetPhysicalFormatNil(b bool)`

 SetPhysicalFormatNil sets the value for PhysicalFormat to be an explicit nil

### UnsetPhysicalFormat
`func (o *Book) UnsetPhysicalFormat()`

UnsetPhysicalFormat ensures that no value is present for PhysicalFormat, not even an explicit nil
### GetCovers

`func (o *Book) GetCovers() []string`

GetCovers returns the Covers field if non-nil, zero value otherwise.

### GetCoversOk

`func (o *Book) GetCoversOk() (*[]string, bool)`

GetCoversOk returns a tuple with the Covers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCovers

`func (o *Book) SetCovers(v []string)`

SetCovers sets Covers field to given value.

### HasCovers

`func (o *Book) HasCovers() bool`

HasCovers returns a boolean if a field has been set.

### SetCoversNil

`func (o *Book) SetCoversNil(b bool)`

 SetCoversNil sets the value for Covers to be an explicit nil

### UnsetCovers
`func (o *Book) UnsetCovers()`

UnsetCovers ensures that no value is present for Covers, not even an explicit nil
### GetAuthors

`func (o *Book) GetAuthors() []EntityRef`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *Book) GetAuthorsOk() (*[]EntityRef, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *Book) SetAuthors(v []EntityRef)`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *Book) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### SetAuthorsNil

`func (o *Book) SetAuthorsNil(b bool)`

 SetAuthorsNil sets the value for Authors to be an explicit nil

### UnsetAuthors
`func (o *Book) UnsetAuthors()`

UnsetAuthors ensures that no value is present for Authors, not even an explicit nil
### GetContributors

`func (o *Book) GetContributors() []Contributor`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *Book) GetContributorsOk() (*[]Contributor, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *Book) SetContributors(v []Contributor)`

SetContributors sets Contributors field to given value.

### HasContributors

`func (o *Book) HasContributors() bool`

HasContributors returns a boolean if a field has been set.

### SetContributorsNil

`func (o *Book) SetContributorsNil(b bool)`

 SetContributorsNil sets the value for Contributors to be an explicit nil

### UnsetContributors
`func (o *Book) UnsetContributors()`

UnsetContributors ensures that no value is present for Contributors, not even an explicit nil
### GetSubjects

`func (o *Book) GetSubjects() []SubjectRef`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *Book) GetSubjectsOk() (*[]SubjectRef, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *Book) SetSubjects(v []SubjectRef)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *Book) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.

### SetSubjectsNil

`func (o *Book) SetSubjectsNil(b bool)`

 SetSubjectsNil sets the value for Subjects to be an explicit nil

### UnsetSubjects
`func (o *Book) UnsetSubjects()`

UnsetSubjects ensures that no value is present for Subjects, not even an explicit nil
### GetSeries

`func (o *Book) GetSeries() Series`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *Book) GetSeriesOk() (*Series, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *Book) SetSeries(v Series)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *Book) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### SetSeriesNil

`func (o *Book) SetSeriesNil(b bool)`

 SetSeriesNil sets the value for Series to be an explicit nil

### UnsetSeries
`func (o *Book) UnsetSeries()`

UnsetSeries ensures that no value is present for Series, not even an explicit nil
### GetGenres

`func (o *Book) GetGenres() []EntityRef`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *Book) GetGenresOk() (*[]EntityRef, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *Book) SetGenres(v []EntityRef)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *Book) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### SetGenresNil

`func (o *Book) SetGenresNil(b bool)`

 SetGenresNil sets the value for Genres to be an explicit nil

### UnsetGenres
`func (o *Book) UnsetGenres()`

UnsetGenres ensures that no value is present for Genres, not even an explicit nil
### GetOtherEditions

`func (o *Book) GetOtherEditions() []EditionRef`

GetOtherEditions returns the OtherEditions field if non-nil, zero value otherwise.

### GetOtherEditionsOk

`func (o *Book) GetOtherEditionsOk() (*[]EditionRef, bool)`

GetOtherEditionsOk returns a tuple with the OtherEditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherEditions

`func (o *Book) SetOtherEditions(v []EditionRef)`

SetOtherEditions sets OtherEditions field to given value.

### HasOtherEditions

`func (o *Book) HasOtherEditions() bool`

HasOtherEditions returns a boolean if a field has been set.

### SetOtherEditionsNil

`func (o *Book) SetOtherEditionsNil(b bool)`

 SetOtherEditionsNil sets the value for OtherEditions to be an explicit nil

### UnsetOtherEditions
`func (o *Book) UnsetOtherEditions()`

UnsetOtherEditions ensures that no value is present for OtherEditions, not even an explicit nil
### GetPreviousEdition

`func (o *Book) GetPreviousEdition() EditionRef`

GetPreviousEdition returns the PreviousEdition field if non-nil, zero value otherwise.

### GetPreviousEditionOk

`func (o *Book) GetPreviousEditionOk() (*EditionRef, bool)`

GetPreviousEditionOk returns a tuple with the PreviousEdition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousEdition

`func (o *Book) SetPreviousEdition(v EditionRef)`

SetPreviousEdition sets PreviousEdition field to given value.

### HasPreviousEdition

`func (o *Book) HasPreviousEdition() bool`

HasPreviousEdition returns a boolean if a field has been set.

### SetPreviousEditionNil

`func (o *Book) SetPreviousEditionNil(b bool)`

 SetPreviousEditionNil sets the value for PreviousEdition to be an explicit nil

### UnsetPreviousEdition
`func (o *Book) UnsetPreviousEdition()`

UnsetPreviousEdition ensures that no value is present for PreviousEdition, not even an explicit nil
### GetNextEdition

`func (o *Book) GetNextEdition() EditionRef`

GetNextEdition returns the NextEdition field if non-nil, zero value otherwise.

### GetNextEditionOk

`func (o *Book) GetNextEditionOk() (*EditionRef, bool)`

GetNextEditionOk returns a tuple with the NextEdition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextEdition

`func (o *Book) SetNextEdition(v EditionRef)`

SetNextEdition sets NextEdition field to given value.

### HasNextEdition

`func (o *Book) HasNextEdition() bool`

HasNextEdition returns a boolean if a field has been set.

### SetNextEditionNil

`func (o *Book) SetNextEditionNil(b bool)`

 SetNextEditionNil sets the value for NextEdition to be an explicit nil

### UnsetNextEdition
`func (o *Book) UnsetNextEdition()`

UnsetNextEdition ensures that no value is present for NextEdition, not even an explicit nil
### GetPhysicalDimensions

`func (o *Book) GetPhysicalDimensions() string`

GetPhysicalDimensions returns the PhysicalDimensions field if non-nil, zero value otherwise.

### GetPhysicalDimensionsOk

`func (o *Book) GetPhysicalDimensionsOk() (*string, bool)`

GetPhysicalDimensionsOk returns a tuple with the PhysicalDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalDimensions

`func (o *Book) SetPhysicalDimensions(v string)`

SetPhysicalDimensions sets PhysicalDimensions field to given value.

### HasPhysicalDimensions

`func (o *Book) HasPhysicalDimensions() bool`

HasPhysicalDimensions returns a boolean if a field has been set.

### SetPhysicalDimensionsNil

`func (o *Book) SetPhysicalDimensionsNil(b bool)`

 SetPhysicalDimensionsNil sets the value for PhysicalDimensions to be an explicit nil

### UnsetPhysicalDimensions
`func (o *Book) UnsetPhysicalDimensions()`

UnsetPhysicalDimensions ensures that no value is present for PhysicalDimensions, not even an explicit nil
### GetRatings

`func (o *Book) GetRatings() Ratings`

GetRatings returns the Ratings field if non-nil, zero value otherwise.

### GetRatingsOk

`func (o *Book) GetRatingsOk() (*Ratings, bool)`

GetRatingsOk returns a tuple with the Ratings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatings

`func (o *Book) SetRatings(v Ratings)`

SetRatings sets Ratings field to given value.

### HasRatings

`func (o *Book) HasRatings() bool`

HasRatings returns a boolean if a field has been set.

### SetRatingsNil

`func (o *Book) SetRatingsNil(b bool)`

 SetRatingsNil sets the value for Ratings to be an explicit nil

### UnsetRatings
`func (o *Book) UnsetRatings()`

UnsetRatings ensures that no value is present for Ratings, not even an explicit nil
### GetLanguage

`func (o *Book) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Book) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Book) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Book) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *Book) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *Book) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetSaleStatus

`func (o *Book) GetSaleStatus() string`

GetSaleStatus returns the SaleStatus field if non-nil, zero value otherwise.

### GetSaleStatusOk

`func (o *Book) GetSaleStatusOk() (*string, bool)`

GetSaleStatusOk returns a tuple with the SaleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleStatus

`func (o *Book) SetSaleStatus(v string)`

SetSaleStatus sets SaleStatus field to given value.

### HasSaleStatus

`func (o *Book) HasSaleStatus() bool`

HasSaleStatus returns a boolean if a field has been set.

### SetSaleStatusNil

`func (o *Book) SetSaleStatusNil(b bool)`

 SetSaleStatusNil sets the value for SaleStatus to be an explicit nil

### UnsetSaleStatus
`func (o *Book) UnsetSaleStatus()`

UnsetSaleStatus ensures that no value is present for SaleStatus, not even an explicit nil
### GetExcerpt

`func (o *Book) GetExcerpt() string`

GetExcerpt returns the Excerpt field if non-nil, zero value otherwise.

### GetExcerptOk

`func (o *Book) GetExcerptOk() (*string, bool)`

GetExcerptOk returns a tuple with the Excerpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcerpt

`func (o *Book) SetExcerpt(v string)`

SetExcerpt sets Excerpt field to given value.

### HasExcerpt

`func (o *Book) HasExcerpt() bool`

HasExcerpt returns a boolean if a field has been set.

### SetExcerptNil

`func (o *Book) SetExcerptNil(b bool)`

 SetExcerptNil sets the value for Excerpt to be an explicit nil

### UnsetExcerpt
`func (o *Book) UnsetExcerpt()`

UnsetExcerpt ensures that no value is present for Excerpt, not even an explicit nil
### GetReviewQuotes

`func (o *Book) GetReviewQuotes() string`

GetReviewQuotes returns the ReviewQuotes field if non-nil, zero value otherwise.

### GetReviewQuotesOk

`func (o *Book) GetReviewQuotesOk() (*string, bool)`

GetReviewQuotesOk returns a tuple with the ReviewQuotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewQuotes

`func (o *Book) SetReviewQuotes(v string)`

SetReviewQuotes sets ReviewQuotes field to given value.

### HasReviewQuotes

`func (o *Book) HasReviewQuotes() bool`

HasReviewQuotes returns a boolean if a field has been set.

### SetReviewQuotesNil

`func (o *Book) SetReviewQuotesNil(b bool)`

 SetReviewQuotesNil sets the value for ReviewQuotes to be an explicit nil

### UnsetReviewQuotes
`func (o *Book) UnsetReviewQuotes()`

UnsetReviewQuotes ensures that no value is present for ReviewQuotes, not even an explicit nil
### GetRetailLinks

`func (o *Book) GetRetailLinks() []RetailLink`

GetRetailLinks returns the RetailLinks field if non-nil, zero value otherwise.

### GetRetailLinksOk

`func (o *Book) GetRetailLinksOk() (*[]RetailLink, bool)`

GetRetailLinksOk returns a tuple with the RetailLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetailLinks

`func (o *Book) SetRetailLinks(v []RetailLink)`

SetRetailLinks sets RetailLinks field to given value.

### HasRetailLinks

`func (o *Book) HasRetailLinks() bool`

HasRetailLinks returns a boolean if a field has been set.

### SetRetailLinksNil

`func (o *Book) SetRetailLinksNil(b bool)`

 SetRetailLinksNil sets the value for RetailLinks to be an explicit nil

### UnsetRetailLinks
`func (o *Book) UnsetRetailLinks()`

UnsetRetailLinks ensures that no value is present for RetailLinks, not even an explicit nil
### GetCategories

`func (o *Book) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Book) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Book) SetCategories(v []Category)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *Book) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategoriesNil

`func (o *Book) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *Book) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil
### GetAiSummary

`func (o *Book) GetAiSummary() string`

GetAiSummary returns the AiSummary field if non-nil, zero value otherwise.

### GetAiSummaryOk

`func (o *Book) GetAiSummaryOk() (*string, bool)`

GetAiSummaryOk returns a tuple with the AiSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiSummary

`func (o *Book) SetAiSummary(v string)`

SetAiSummary sets AiSummary field to given value.

### HasAiSummary

`func (o *Book) HasAiSummary() bool`

HasAiSummary returns a boolean if a field has been set.

### SetAiSummaryNil

`func (o *Book) SetAiSummaryNil(b bool)`

 SetAiSummaryNil sets the value for AiSummary to be an explicit nil

### UnsetAiSummary
`func (o *Book) UnsetAiSummary()`

UnsetAiSummary ensures that no value is present for AiSummary, not even an explicit nil
### GetCanonicalEditionId

`func (o *Book) GetCanonicalEditionId() int64`

GetCanonicalEditionId returns the CanonicalEditionId field if non-nil, zero value otherwise.

### GetCanonicalEditionIdOk

`func (o *Book) GetCanonicalEditionIdOk() (*int64, bool)`

GetCanonicalEditionIdOk returns a tuple with the CanonicalEditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalEditionId

`func (o *Book) SetCanonicalEditionId(v int64)`

SetCanonicalEditionId sets CanonicalEditionId field to given value.

### HasCanonicalEditionId

`func (o *Book) HasCanonicalEditionId() bool`

HasCanonicalEditionId returns a boolean if a field has been set.

### SetCanonicalEditionIdNil

`func (o *Book) SetCanonicalEditionIdNil(b bool)`

 SetCanonicalEditionIdNil sets the value for CanonicalEditionId to be an explicit nil

### UnsetCanonicalEditionId
`func (o *Book) UnsetCanonicalEditionId()`

UnsetCanonicalEditionId ensures that no value is present for CanonicalEditionId, not even an explicit nil
### GetLanguageCode

`func (o *Book) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *Book) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *Book) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *Book) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.

### SetLanguageCodeNil

`func (o *Book) SetLanguageCodeNil(b bool)`

 SetLanguageCodeNil sets the value for LanguageCode to be an explicit nil

### UnsetLanguageCode
`func (o *Book) UnsetLanguageCode()`

UnsetLanguageCode ensures that no value is present for LanguageCode, not even an explicit nil
### GetPublishDatePrecision

`func (o *Book) GetPublishDatePrecision() string`

GetPublishDatePrecision returns the PublishDatePrecision field if non-nil, zero value otherwise.

### GetPublishDatePrecisionOk

`func (o *Book) GetPublishDatePrecisionOk() (*string, bool)`

GetPublishDatePrecisionOk returns a tuple with the PublishDatePrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishDatePrecision

`func (o *Book) SetPublishDatePrecision(v string)`

SetPublishDatePrecision sets PublishDatePrecision field to given value.

### HasPublishDatePrecision

`func (o *Book) HasPublishDatePrecision() bool`

HasPublishDatePrecision returns a boolean if a field has been set.

### SetPublishDatePrecisionNil

`func (o *Book) SetPublishDatePrecisionNil(b bool)`

 SetPublishDatePrecisionNil sets the value for PublishDatePrecision to be an explicit nil

### UnsetPublishDatePrecision
`func (o *Book) UnsetPublishDatePrecision()`

UnsetPublishDatePrecision ensures that no value is present for PublishDatePrecision, not even an explicit nil
### GetEstimatedReadTime

`func (o *Book) GetEstimatedReadTime() int64`

GetEstimatedReadTime returns the EstimatedReadTime field if non-nil, zero value otherwise.

### GetEstimatedReadTimeOk

`func (o *Book) GetEstimatedReadTimeOk() (*int64, bool)`

GetEstimatedReadTimeOk returns a tuple with the EstimatedReadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedReadTime

`func (o *Book) SetEstimatedReadTime(v int64)`

SetEstimatedReadTime sets EstimatedReadTime field to given value.

### HasEstimatedReadTime

`func (o *Book) HasEstimatedReadTime() bool`

HasEstimatedReadTime returns a boolean if a field has been set.

### SetEstimatedReadTimeNil

`func (o *Book) SetEstimatedReadTimeNil(b bool)`

 SetEstimatedReadTimeNil sets the value for EstimatedReadTime to be an explicit nil

### UnsetEstimatedReadTime
`func (o *Book) UnsetEstimatedReadTime()`

UnsetEstimatedReadTime ensures that no value is present for EstimatedReadTime, not even an explicit nil
### GetLanguagesCount

`func (o *Book) GetLanguagesCount() int64`

GetLanguagesCount returns the LanguagesCount field if non-nil, zero value otherwise.

### GetLanguagesCountOk

`func (o *Book) GetLanguagesCountOk() (*int64, bool)`

GetLanguagesCountOk returns a tuple with the LanguagesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguagesCount

`func (o *Book) SetLanguagesCount(v int64)`

SetLanguagesCount sets LanguagesCount field to given value.

### HasLanguagesCount

`func (o *Book) HasLanguagesCount() bool`

HasLanguagesCount returns a boolean if a field has been set.

### SetLanguagesCountNil

`func (o *Book) SetLanguagesCountNil(b bool)`

 SetLanguagesCountNil sets the value for LanguagesCount to be an explicit nil

### UnsetLanguagesCount
`func (o *Book) UnsetLanguagesCount()`

UnsetLanguagesCount ensures that no value is present for LanguagesCount, not even an explicit nil
### GetLink

`func (o *Book) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *Book) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *Book) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *Book) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *Book) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *Book) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil
### GetIsbnModel

`func (o *Book) GetIsbnModel() IsbnDetails`

GetIsbnModel returns the IsbnModel field if non-nil, zero value otherwise.

### GetIsbnModelOk

`func (o *Book) GetIsbnModelOk() (*IsbnDetails, bool)`

GetIsbnModelOk returns a tuple with the IsbnModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbnModel

`func (o *Book) SetIsbnModel(v IsbnDetails)`

SetIsbnModel sets IsbnModel field to given value.

### HasIsbnModel

`func (o *Book) HasIsbnModel() bool`

HasIsbnModel returns a boolean if a field has been set.

### SetIsbnModelNil

`func (o *Book) SetIsbnModelNil(b bool)`

 SetIsbnModelNil sets the value for IsbnModel to be an explicit nil

### UnsetIsbnModel
`func (o *Book) UnsetIsbnModel()`

UnsetIsbnModel ensures that no value is present for IsbnModel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


