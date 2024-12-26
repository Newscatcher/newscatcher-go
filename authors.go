// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/Newscatcher/newscatcher-go/internal"
	time "time"
)

type AuthorsGetRequest struct {
	AuthorName             string   `json:"-" url:"author_name"`
	NotAuthorName          *string  `json:"-" url:"not_author_name,omitempty"`
	Sources                string   `json:"-" url:"sources"`
	PredefinedSources      string   `json:"-" url:"predefined_sources"`
	NotSources             string   `json:"-" url:"not_sources"`
	Lang                   string   `json:"-" url:"lang"`
	NotLang                string   `json:"-" url:"not_lang"`
	Countries              string   `json:"-" url:"countries"`
	NotCountries           string   `json:"-" url:"not_countries"`
	From                   *string  `json:"-" url:"from_,omitempty"`
	To                     *string  `json:"-" url:"to_,omitempty"`
	PublishedDatePrecision *string  `json:"-" url:"published_date_precision,omitempty"`
	ByParseDate            *bool    `json:"-" url:"by_parse_date,omitempty"`
	SortBy                 *string  `json:"-" url:"sort_by,omitempty"`
	RankedOnly             *string  `json:"-" url:"ranked_only,omitempty"`
	FromRank               *int     `json:"-" url:"from_rank,omitempty"`
	ToRank                 *int     `json:"-" url:"to_rank,omitempty"`
	IsHeadline             *bool    `json:"-" url:"is_headline,omitempty"`
	IsOpinion              *bool    `json:"-" url:"is_opinion,omitempty"`
	IsPaidContent          *bool    `json:"-" url:"is_paid_content,omitempty"`
	ParentUrl              string   `json:"-" url:"parent_url"`
	AllLinks               string   `json:"-" url:"all_links"`
	AllDomainLinks         string   `json:"-" url:"all_domain_links"`
	WordCountMin           *int     `json:"-" url:"word_count_min,omitempty"`
	WordCountMax           *int     `json:"-" url:"word_count_max,omitempty"`
	Page                   *int     `json:"-" url:"page,omitempty"`
	PageSize               *int     `json:"-" url:"page_size,omitempty"`
	IncludeNlpData         *bool    `json:"-" url:"include_nlp_data,omitempty"`
	HasNlp                 *bool    `json:"-" url:"has_nlp,omitempty"`
	Theme                  *string  `json:"-" url:"theme,omitempty"`
	NotTheme               *string  `json:"-" url:"not_theme,omitempty"`
	TitleSentimentMin      *float64 `json:"-" url:"title_sentiment_min,omitempty"`
	TitleSentimentMax      *float64 `json:"-" url:"title_sentiment_max,omitempty"`
	ContentSentimentMin    *float64 `json:"-" url:"content_sentiment_min,omitempty"`
	ContentSentimentMax    *float64 `json:"-" url:"content_sentiment_max,omitempty"`
	IptcTags               string   `json:"-" url:"iptc_tags"`
	NotIptcTags            string   `json:"-" url:"not_iptc_tags"`
	IabTags                string   `json:"-" url:"iab_tags"`
	NotIabTags             string   `json:"-" url:"not_iab_tags"`
}

type AuthorSearchRequest struct {
	AuthorName             string                         `json:"author_name" url:"-"`
	NotAuthorName          *string                        `json:"not_author_name,omitempty" url:"-"`
	Sources                interface{}                    `json:"sources,omitempty" url:"-"`
	PredefinedSources      interface{}                    `json:"predefined_sources,omitempty" url:"-"`
	NotSources             interface{}                    `json:"not_sources,omitempty" url:"-"`
	Lang                   interface{}                    `json:"lang,omitempty" url:"-"`
	NotLang                interface{}                    `json:"not_lang,omitempty" url:"-"`
	Countries              interface{}                    `json:"countries,omitempty" url:"-"`
	NotCountries           interface{}                    `json:"not_countries,omitempty" url:"-"`
	From                   *AuthorSearchRequestFrom       `json:"from_,omitempty" url:"-"`
	To                     *AuthorSearchRequestTo         `json:"to_,omitempty" url:"-"`
	PublishedDatePrecision *string                        `json:"published_date_precision,omitempty" url:"-"`
	ByParseDate            *bool                          `json:"by_parse_date,omitempty" url:"-"`
	SortBy                 *string                        `json:"sort_by,omitempty" url:"-"`
	RankedOnly             *AuthorSearchRequestRankedOnly `json:"ranked_only,omitempty" url:"-"`
	FromRank               *int                           `json:"from_rank,omitempty" url:"-"`
	ToRank                 *int                           `json:"to_rank,omitempty" url:"-"`
	IsHeadline             *bool                          `json:"is_headline,omitempty" url:"-"`
	IsOpinion              *bool                          `json:"is_opinion,omitempty" url:"-"`
	IsPaidContent          *bool                          `json:"is_paid_content,omitempty" url:"-"`
	ParentUrl              interface{}                    `json:"parent_url,omitempty" url:"-"`
	AllLinks               interface{}                    `json:"all_links,omitempty" url:"-"`
	AllDomainLinks         interface{}                    `json:"all_domain_links,omitempty" url:"-"`
	WordCountMin           *int                           `json:"word_count_min,omitempty" url:"-"`
	WordCountMax           *int                           `json:"word_count_max,omitempty" url:"-"`
	Page                   *int                           `json:"page,omitempty" url:"-"`
	PageSize               *int                           `json:"page_size,omitempty" url:"-"`
	IncludeNlpData         *bool                          `json:"include_nlp_data,omitempty" url:"-"`
	HasNlp                 *bool                          `json:"has_nlp,omitempty" url:"-"`
	Theme                  *string                        `json:"theme,omitempty" url:"-"`
	NotTheme               *string                        `json:"not_theme,omitempty" url:"-"`
	TitleSentimentMin      *float64                       `json:"title_sentiment_min,omitempty" url:"-"`
	TitleSentimentMax      *float64                       `json:"title_sentiment_max,omitempty" url:"-"`
	ContentSentimentMin    *float64                       `json:"content_sentiment_min,omitempty" url:"-"`
	ContentSentimentMax    *float64                       `json:"content_sentiment_max,omitempty" url:"-"`
	IptcTags               interface{}                    `json:"iptc_tags,omitempty" url:"-"`
	NotIptcTags            interface{}                    `json:"not_iptc_tags,omitempty" url:"-"`
	IabTags                interface{}                    `json:"iab_tags,omitempty" url:"-"`
	NotIabTags             interface{}                    `json:"not_iab_tags,omitempty" url:"-"`
}

type AuthorSearchRequestFrom struct {
	String   string
	DateTime time.Time

	typ string
}

func NewAuthorSearchRequestFromFromString(value string) *AuthorSearchRequestFrom {
	return &AuthorSearchRequestFrom{typ: "String", String: value}
}

func NewAuthorSearchRequestFromFromDateTime(value time.Time) *AuthorSearchRequestFrom {
	return &AuthorSearchRequestFrom{typ: "DateTime", DateTime: value}
}

func (a *AuthorSearchRequestFrom) GetString() string {
	if a == nil {
		return ""
	}
	return a.String
}

func (a *AuthorSearchRequestFrom) GetDateTime() time.Time {
	if a == nil {
		return time.Time{}
	}
	return a.DateTime
}

func (a *AuthorSearchRequestFrom) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		a.typ = "String"
		a.String = valueString
		return nil
	}
	var valueDateTime *internal.DateTime
	if err := json.Unmarshal(data, &valueDateTime); err == nil {
		a.typ = "DateTime"
		a.DateTime = valueDateTime.Time()
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a AuthorSearchRequestFrom) MarshalJSON() ([]byte, error) {
	if a.typ == "String" || a.String != "" {
		return json.Marshal(a.String)
	}
	if a.typ == "DateTime" || !a.DateTime.IsZero() {
		return json.Marshal(internal.NewDateTime(a.DateTime))
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", a)
}

type AuthorSearchRequestFromVisitor interface {
	VisitString(string) error
	VisitDateTime(time.Time) error
}

func (a *AuthorSearchRequestFrom) Accept(visitor AuthorSearchRequestFromVisitor) error {
	if a.typ == "String" || a.String != "" {
		return visitor.VisitString(a.String)
	}
	if a.typ == "DateTime" || !a.DateTime.IsZero() {
		return visitor.VisitDateTime(a.DateTime)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", a)
}

type AuthorSearchRequestRankedOnly struct {
	String  string
	Boolean bool

	typ string
}

func NewAuthorSearchRequestRankedOnlyFromString(value string) *AuthorSearchRequestRankedOnly {
	return &AuthorSearchRequestRankedOnly{typ: "String", String: value}
}

func NewAuthorSearchRequestRankedOnlyFromBoolean(value bool) *AuthorSearchRequestRankedOnly {
	return &AuthorSearchRequestRankedOnly{typ: "Boolean", Boolean: value}
}

func (a *AuthorSearchRequestRankedOnly) GetString() string {
	if a == nil {
		return ""
	}
	return a.String
}

func (a *AuthorSearchRequestRankedOnly) GetBoolean() bool {
	if a == nil {
		return false
	}
	return a.Boolean
}

func (a *AuthorSearchRequestRankedOnly) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		a.typ = "String"
		a.String = valueString
		return nil
	}
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		a.typ = "Boolean"
		a.Boolean = valueBoolean
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a AuthorSearchRequestRankedOnly) MarshalJSON() ([]byte, error) {
	if a.typ == "String" || a.String != "" {
		return json.Marshal(a.String)
	}
	if a.typ == "Boolean" || a.Boolean != false {
		return json.Marshal(a.Boolean)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", a)
}

type AuthorSearchRequestRankedOnlyVisitor interface {
	VisitString(string) error
	VisitBoolean(bool) error
}

func (a *AuthorSearchRequestRankedOnly) Accept(visitor AuthorSearchRequestRankedOnlyVisitor) error {
	if a.typ == "String" || a.String != "" {
		return visitor.VisitString(a.String)
	}
	if a.typ == "Boolean" || a.Boolean != false {
		return visitor.VisitBoolean(a.Boolean)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", a)
}

type AuthorSearchRequestTo struct {
	String   string
	DateTime time.Time

	typ string
}

func NewAuthorSearchRequestToFromString(value string) *AuthorSearchRequestTo {
	return &AuthorSearchRequestTo{typ: "String", String: value}
}

func NewAuthorSearchRequestToFromDateTime(value time.Time) *AuthorSearchRequestTo {
	return &AuthorSearchRequestTo{typ: "DateTime", DateTime: value}
}

func (a *AuthorSearchRequestTo) GetString() string {
	if a == nil {
		return ""
	}
	return a.String
}

func (a *AuthorSearchRequestTo) GetDateTime() time.Time {
	if a == nil {
		return time.Time{}
	}
	return a.DateTime
}

func (a *AuthorSearchRequestTo) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		a.typ = "String"
		a.String = valueString
		return nil
	}
	var valueDateTime *internal.DateTime
	if err := json.Unmarshal(data, &valueDateTime); err == nil {
		a.typ = "DateTime"
		a.DateTime = valueDateTime.Time()
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a AuthorSearchRequestTo) MarshalJSON() ([]byte, error) {
	if a.typ == "String" || a.String != "" {
		return json.Marshal(a.String)
	}
	if a.typ == "DateTime" || !a.DateTime.IsZero() {
		return json.Marshal(internal.NewDateTime(a.DateTime))
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", a)
}

type AuthorSearchRequestToVisitor interface {
	VisitString(string) error
	VisitDateTime(time.Time) error
}

func (a *AuthorSearchRequestTo) Accept(visitor AuthorSearchRequestToVisitor) error {
	if a.typ == "String" || a.String != "" {
		return visitor.VisitString(a.String)
	}
	if a.typ == "DateTime" || !a.DateTime.IsZero() {
		return visitor.VisitDateTime(a.DateTime)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", a)
}

type AuthorsGetResponse struct {
	SearchResponse       *SearchResponse
	FailedSearchResponse *FailedSearchResponse

	typ string
}

func NewAuthorsGetResponseFromSearchResponse(value *SearchResponse) *AuthorsGetResponse {
	return &AuthorsGetResponse{typ: "SearchResponse", SearchResponse: value}
}

func NewAuthorsGetResponseFromFailedSearchResponse(value *FailedSearchResponse) *AuthorsGetResponse {
	return &AuthorsGetResponse{typ: "FailedSearchResponse", FailedSearchResponse: value}
}

func (a *AuthorsGetResponse) GetSearchResponse() *SearchResponse {
	if a == nil {
		return nil
	}
	return a.SearchResponse
}

func (a *AuthorsGetResponse) GetFailedSearchResponse() *FailedSearchResponse {
	if a == nil {
		return nil
	}
	return a.FailedSearchResponse
}

func (a *AuthorsGetResponse) UnmarshalJSON(data []byte) error {
	valueSearchResponse := new(SearchResponse)
	if err := json.Unmarshal(data, &valueSearchResponse); err == nil {
		a.typ = "SearchResponse"
		a.SearchResponse = valueSearchResponse
		return nil
	}
	valueFailedSearchResponse := new(FailedSearchResponse)
	if err := json.Unmarshal(data, &valueFailedSearchResponse); err == nil {
		a.typ = "FailedSearchResponse"
		a.FailedSearchResponse = valueFailedSearchResponse
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a AuthorsGetResponse) MarshalJSON() ([]byte, error) {
	if a.typ == "SearchResponse" || a.SearchResponse != nil {
		return json.Marshal(a.SearchResponse)
	}
	if a.typ == "FailedSearchResponse" || a.FailedSearchResponse != nil {
		return json.Marshal(a.FailedSearchResponse)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", a)
}

type AuthorsGetResponseVisitor interface {
	VisitSearchResponse(*SearchResponse) error
	VisitFailedSearchResponse(*FailedSearchResponse) error
}

func (a *AuthorsGetResponse) Accept(visitor AuthorsGetResponseVisitor) error {
	if a.typ == "SearchResponse" || a.SearchResponse != nil {
		return visitor.VisitSearchResponse(a.SearchResponse)
	}
	if a.typ == "FailedSearchResponse" || a.FailedSearchResponse != nil {
		return visitor.VisitFailedSearchResponse(a.FailedSearchResponse)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", a)
}

type AuthorsPostResponse struct {
	SearchResponse       *SearchResponse
	FailedSearchResponse *FailedSearchResponse

	typ string
}

func NewAuthorsPostResponseFromSearchResponse(value *SearchResponse) *AuthorsPostResponse {
	return &AuthorsPostResponse{typ: "SearchResponse", SearchResponse: value}
}

func NewAuthorsPostResponseFromFailedSearchResponse(value *FailedSearchResponse) *AuthorsPostResponse {
	return &AuthorsPostResponse{typ: "FailedSearchResponse", FailedSearchResponse: value}
}

func (a *AuthorsPostResponse) GetSearchResponse() *SearchResponse {
	if a == nil {
		return nil
	}
	return a.SearchResponse
}

func (a *AuthorsPostResponse) GetFailedSearchResponse() *FailedSearchResponse {
	if a == nil {
		return nil
	}
	return a.FailedSearchResponse
}

func (a *AuthorsPostResponse) UnmarshalJSON(data []byte) error {
	valueSearchResponse := new(SearchResponse)
	if err := json.Unmarshal(data, &valueSearchResponse); err == nil {
		a.typ = "SearchResponse"
		a.SearchResponse = valueSearchResponse
		return nil
	}
	valueFailedSearchResponse := new(FailedSearchResponse)
	if err := json.Unmarshal(data, &valueFailedSearchResponse); err == nil {
		a.typ = "FailedSearchResponse"
		a.FailedSearchResponse = valueFailedSearchResponse
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a AuthorsPostResponse) MarshalJSON() ([]byte, error) {
	if a.typ == "SearchResponse" || a.SearchResponse != nil {
		return json.Marshal(a.SearchResponse)
	}
	if a.typ == "FailedSearchResponse" || a.FailedSearchResponse != nil {
		return json.Marshal(a.FailedSearchResponse)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", a)
}

type AuthorsPostResponseVisitor interface {
	VisitSearchResponse(*SearchResponse) error
	VisitFailedSearchResponse(*FailedSearchResponse) error
}

func (a *AuthorsPostResponse) Accept(visitor AuthorsPostResponseVisitor) error {
	if a.typ == "SearchResponse" || a.SearchResponse != nil {
		return visitor.VisitSearchResponse(a.SearchResponse)
	}
	if a.typ == "FailedSearchResponse" || a.FailedSearchResponse != nil {
		return visitor.VisitFailedSearchResponse(a.FailedSearchResponse)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", a)
}