// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/Newscatcher/newscatcher-go/internal"
	time "time"
)

type SearchSimilarGetRequest struct {
	Q                       string   `json:"-" url:"q"`
	SearchIn                *string  `json:"-" url:"search_in,omitempty"`
	IncludeSimilarDocuments *bool    `json:"-" url:"include_similar_documents,omitempty"`
	SimilarDocumentsNumber  *int     `json:"-" url:"similar_documents_number,omitempty"`
	SimilarDocumentsFields  *string  `json:"-" url:"similar_documents_fields,omitempty"`
	PredefinedSources       string   `json:"-" url:"predefined_sources"`
	Sources                 string   `json:"-" url:"sources"`
	NotSources              string   `json:"-" url:"not_sources"`
	Lang                    string   `json:"-" url:"lang"`
	NotLang                 string   `json:"-" url:"not_lang"`
	Countries               string   `json:"-" url:"countries"`
	NotCountries            string   `json:"-" url:"not_countries"`
	From                    *string  `json:"-" url:"from_,omitempty"`
	To                      *string  `json:"-" url:"to_,omitempty"`
	ByParseDate             *bool    `json:"-" url:"by_parse_date,omitempty"`
	PublishedDatePrecision  *string  `json:"-" url:"published_date_precision,omitempty"`
	SortBy                  *string  `json:"-" url:"sort_by,omitempty"`
	RankedOnly              *string  `json:"-" url:"ranked_only,omitempty"`
	FromRank                *int     `json:"-" url:"from_rank,omitempty"`
	ToRank                  *int     `json:"-" url:"to_rank,omitempty"`
	IsHeadline              *bool    `json:"-" url:"is_headline,omitempty"`
	IsOpinion               *bool    `json:"-" url:"is_opinion,omitempty"`
	IsPaidContent           *bool    `json:"-" url:"is_paid_content,omitempty"`
	ParentUrl               string   `json:"-" url:"parent_url"`
	AllLinks                string   `json:"-" url:"all_links"`
	AllDomainLinks          string   `json:"-" url:"all_domain_links"`
	WordCountMin            *int     `json:"-" url:"word_count_min,omitempty"`
	WordCountMax            *int     `json:"-" url:"word_count_max,omitempty"`
	Page                    *int     `json:"-" url:"page,omitempty"`
	PageSize                *int     `json:"-" url:"page_size,omitempty"`
	IncludeNlpData          *bool    `json:"-" url:"include_nlp_data,omitempty"`
	HasNlp                  *bool    `json:"-" url:"has_nlp,omitempty"`
	Theme                   *string  `json:"-" url:"theme,omitempty"`
	NotTheme                *string  `json:"-" url:"not_theme,omitempty"`
	TitleSentimentMin       *float64 `json:"-" url:"title_sentiment_min,omitempty"`
	TitleSentimentMax       *float64 `json:"-" url:"title_sentiment_max,omitempty"`
	ContentSentimentMin     *float64 `json:"-" url:"content_sentiment_min,omitempty"`
	ContentSentimentMax     *float64 `json:"-" url:"content_sentiment_max,omitempty"`
	IptcTags                string   `json:"-" url:"iptc_tags"`
	NotIptcTags             string   `json:"-" url:"not_iptc_tags"`
}

type MoreLikeThisRequest struct {
	Q                       string                         `json:"q" url:"-"`
	SearchIn                *string                        `json:"search_in,omitempty" url:"-"`
	IncludeSimilarDocuments *bool                          `json:"include_similar_documents,omitempty" url:"-"`
	SimilarDocumentsNumber  *int                           `json:"similar_documents_number,omitempty" url:"-"`
	SimilarDocumentsFields  *string                        `json:"similar_documents_fields,omitempty" url:"-"`
	PredefinedSources       interface{}                    `json:"predefined_sources,omitempty" url:"-"`
	Sources                 interface{}                    `json:"sources,omitempty" url:"-"`
	NotSources              interface{}                    `json:"not_sources,omitempty" url:"-"`
	Lang                    interface{}                    `json:"lang,omitempty" url:"-"`
	NotLang                 interface{}                    `json:"not_lang,omitempty" url:"-"`
	Countries               interface{}                    `json:"countries,omitempty" url:"-"`
	NotCountries            interface{}                    `json:"not_countries,omitempty" url:"-"`
	From                    *MoreLikeThisRequestFrom       `json:"from_,omitempty" url:"-"`
	To                      *MoreLikeThisRequestTo         `json:"to_,omitempty" url:"-"`
	ByParseDate             *bool                          `json:"by_parse_date,omitempty" url:"-"`
	PublishedDatePrecision  *string                        `json:"published_date_precision,omitempty" url:"-"`
	SortBy                  *string                        `json:"sort_by,omitempty" url:"-"`
	RankedOnly              *MoreLikeThisRequestRankedOnly `json:"ranked_only,omitempty" url:"-"`
	FromRank                *int                           `json:"from_rank,omitempty" url:"-"`
	ToRank                  *int                           `json:"to_rank,omitempty" url:"-"`
	IsHeadline              *bool                          `json:"is_headline,omitempty" url:"-"`
	IsOpinion               *bool                          `json:"is_opinion,omitempty" url:"-"`
	IsPaidContent           *bool                          `json:"is_paid_content,omitempty" url:"-"`
	ParentUrl               interface{}                    `json:"parent_url,omitempty" url:"-"`
	AllLinks                interface{}                    `json:"all_links,omitempty" url:"-"`
	AllDomainLinks          interface{}                    `json:"all_domain_links,omitempty" url:"-"`
	WordCountMin            *int                           `json:"word_count_min,omitempty" url:"-"`
	WordCountMax            *int                           `json:"word_count_max,omitempty" url:"-"`
	Page                    *int                           `json:"page,omitempty" url:"-"`
	PageSize                *int                           `json:"page_size,omitempty" url:"-"`
	IncludeNlpData          *bool                          `json:"include_nlp_data,omitempty" url:"-"`
	HasNlp                  *bool                          `json:"has_nlp,omitempty" url:"-"`
	Theme                   *string                        `json:"theme,omitempty" url:"-"`
	NotTheme                *string                        `json:"not_theme,omitempty" url:"-"`
	TitleSentimentMin       *float64                       `json:"title_sentiment_min,omitempty" url:"-"`
	TitleSentimentMax       *float64                       `json:"title_sentiment_max,omitempty" url:"-"`
	ContentSentimentMin     *float64                       `json:"content_sentiment_min,omitempty" url:"-"`
	ContentSentimentMax     *float64                       `json:"content_sentiment_max,omitempty" url:"-"`
	IptcTags                interface{}                    `json:"iptc_tags,omitempty" url:"-"`
	NotIptcTags             interface{}                    `json:"not_iptc_tags,omitempty" url:"-"`
}

type MoreLikeThisRequestFrom struct {
	String   string
	DateTime time.Time

	typ string
}

func NewMoreLikeThisRequestFromFromString(value string) *MoreLikeThisRequestFrom {
	return &MoreLikeThisRequestFrom{typ: "String", String: value}
}

func NewMoreLikeThisRequestFromFromDateTime(value time.Time) *MoreLikeThisRequestFrom {
	return &MoreLikeThisRequestFrom{typ: "DateTime", DateTime: value}
}

func (m *MoreLikeThisRequestFrom) GetString() string {
	if m == nil {
		return ""
	}
	return m.String
}

func (m *MoreLikeThisRequestFrom) GetDateTime() time.Time {
	if m == nil {
		return time.Time{}
	}
	return m.DateTime
}

func (m *MoreLikeThisRequestFrom) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		m.typ = "String"
		m.String = valueString
		return nil
	}
	var valueDateTime *internal.DateTime
	if err := json.Unmarshal(data, &valueDateTime); err == nil {
		m.typ = "DateTime"
		m.DateTime = valueDateTime.Time()
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, m)
}

func (m MoreLikeThisRequestFrom) MarshalJSON() ([]byte, error) {
	if m.typ == "String" || m.String != "" {
		return json.Marshal(m.String)
	}
	if m.typ == "DateTime" || !m.DateTime.IsZero() {
		return json.Marshal(internal.NewDateTime(m.DateTime))
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", m)
}

type MoreLikeThisRequestFromVisitor interface {
	VisitString(string) error
	VisitDateTime(time.Time) error
}

func (m *MoreLikeThisRequestFrom) Accept(visitor MoreLikeThisRequestFromVisitor) error {
	if m.typ == "String" || m.String != "" {
		return visitor.VisitString(m.String)
	}
	if m.typ == "DateTime" || !m.DateTime.IsZero() {
		return visitor.VisitDateTime(m.DateTime)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", m)
}

type MoreLikeThisRequestRankedOnly struct {
	String  string
	Boolean bool

	typ string
}

func NewMoreLikeThisRequestRankedOnlyFromString(value string) *MoreLikeThisRequestRankedOnly {
	return &MoreLikeThisRequestRankedOnly{typ: "String", String: value}
}

func NewMoreLikeThisRequestRankedOnlyFromBoolean(value bool) *MoreLikeThisRequestRankedOnly {
	return &MoreLikeThisRequestRankedOnly{typ: "Boolean", Boolean: value}
}

func (m *MoreLikeThisRequestRankedOnly) GetString() string {
	if m == nil {
		return ""
	}
	return m.String
}

func (m *MoreLikeThisRequestRankedOnly) GetBoolean() bool {
	if m == nil {
		return false
	}
	return m.Boolean
}

func (m *MoreLikeThisRequestRankedOnly) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		m.typ = "String"
		m.String = valueString
		return nil
	}
	var valueBoolean bool
	if err := json.Unmarshal(data, &valueBoolean); err == nil {
		m.typ = "Boolean"
		m.Boolean = valueBoolean
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, m)
}

func (m MoreLikeThisRequestRankedOnly) MarshalJSON() ([]byte, error) {
	if m.typ == "String" || m.String != "" {
		return json.Marshal(m.String)
	}
	if m.typ == "Boolean" || m.Boolean != false {
		return json.Marshal(m.Boolean)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", m)
}

type MoreLikeThisRequestRankedOnlyVisitor interface {
	VisitString(string) error
	VisitBoolean(bool) error
}

func (m *MoreLikeThisRequestRankedOnly) Accept(visitor MoreLikeThisRequestRankedOnlyVisitor) error {
	if m.typ == "String" || m.String != "" {
		return visitor.VisitString(m.String)
	}
	if m.typ == "Boolean" || m.Boolean != false {
		return visitor.VisitBoolean(m.Boolean)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", m)
}

type MoreLikeThisRequestTo struct {
	String   string
	DateTime time.Time

	typ string
}

func NewMoreLikeThisRequestToFromString(value string) *MoreLikeThisRequestTo {
	return &MoreLikeThisRequestTo{typ: "String", String: value}
}

func NewMoreLikeThisRequestToFromDateTime(value time.Time) *MoreLikeThisRequestTo {
	return &MoreLikeThisRequestTo{typ: "DateTime", DateTime: value}
}

func (m *MoreLikeThisRequestTo) GetString() string {
	if m == nil {
		return ""
	}
	return m.String
}

func (m *MoreLikeThisRequestTo) GetDateTime() time.Time {
	if m == nil {
		return time.Time{}
	}
	return m.DateTime
}

func (m *MoreLikeThisRequestTo) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		m.typ = "String"
		m.String = valueString
		return nil
	}
	var valueDateTime *internal.DateTime
	if err := json.Unmarshal(data, &valueDateTime); err == nil {
		m.typ = "DateTime"
		m.DateTime = valueDateTime.Time()
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, m)
}

func (m MoreLikeThisRequestTo) MarshalJSON() ([]byte, error) {
	if m.typ == "String" || m.String != "" {
		return json.Marshal(m.String)
	}
	if m.typ == "DateTime" || !m.DateTime.IsZero() {
		return json.Marshal(internal.NewDateTime(m.DateTime))
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", m)
}

type MoreLikeThisRequestToVisitor interface {
	VisitString(string) error
	VisitDateTime(time.Time) error
}

func (m *MoreLikeThisRequestTo) Accept(visitor MoreLikeThisRequestToVisitor) error {
	if m.typ == "String" || m.String != "" {
		return visitor.VisitString(m.String)
	}
	if m.typ == "DateTime" || !m.DateTime.IsZero() {
		return visitor.VisitDateTime(m.DateTime)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", m)
}

type SearchSimilarGetResponse struct {
	SearchResponse       *SearchResponse
	FailedSearchResponse *FailedSearchResponse

	typ string
}

func NewSearchSimilarGetResponseFromSearchResponse(value *SearchResponse) *SearchSimilarGetResponse {
	return &SearchSimilarGetResponse{typ: "SearchResponse", SearchResponse: value}
}

func NewSearchSimilarGetResponseFromFailedSearchResponse(value *FailedSearchResponse) *SearchSimilarGetResponse {
	return &SearchSimilarGetResponse{typ: "FailedSearchResponse", FailedSearchResponse: value}
}

func (s *SearchSimilarGetResponse) GetSearchResponse() *SearchResponse {
	if s == nil {
		return nil
	}
	return s.SearchResponse
}

func (s *SearchSimilarGetResponse) GetFailedSearchResponse() *FailedSearchResponse {
	if s == nil {
		return nil
	}
	return s.FailedSearchResponse
}

func (s *SearchSimilarGetResponse) UnmarshalJSON(data []byte) error {
	valueSearchResponse := new(SearchResponse)
	if err := json.Unmarshal(data, &valueSearchResponse); err == nil {
		s.typ = "SearchResponse"
		s.SearchResponse = valueSearchResponse
		return nil
	}
	valueFailedSearchResponse := new(FailedSearchResponse)
	if err := json.Unmarshal(data, &valueFailedSearchResponse); err == nil {
		s.typ = "FailedSearchResponse"
		s.FailedSearchResponse = valueFailedSearchResponse
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, s)
}

func (s SearchSimilarGetResponse) MarshalJSON() ([]byte, error) {
	if s.typ == "SearchResponse" || s.SearchResponse != nil {
		return json.Marshal(s.SearchResponse)
	}
	if s.typ == "FailedSearchResponse" || s.FailedSearchResponse != nil {
		return json.Marshal(s.FailedSearchResponse)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", s)
}

type SearchSimilarGetResponseVisitor interface {
	VisitSearchResponse(*SearchResponse) error
	VisitFailedSearchResponse(*FailedSearchResponse) error
}

func (s *SearchSimilarGetResponse) Accept(visitor SearchSimilarGetResponseVisitor) error {
	if s.typ == "SearchResponse" || s.SearchResponse != nil {
		return visitor.VisitSearchResponse(s.SearchResponse)
	}
	if s.typ == "FailedSearchResponse" || s.FailedSearchResponse != nil {
		return visitor.VisitFailedSearchResponse(s.FailedSearchResponse)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", s)
}

type SearchSimilarPostResponse struct {
	SearchResponse       *SearchResponse
	FailedSearchResponse *FailedSearchResponse

	typ string
}

func NewSearchSimilarPostResponseFromSearchResponse(value *SearchResponse) *SearchSimilarPostResponse {
	return &SearchSimilarPostResponse{typ: "SearchResponse", SearchResponse: value}
}

func NewSearchSimilarPostResponseFromFailedSearchResponse(value *FailedSearchResponse) *SearchSimilarPostResponse {
	return &SearchSimilarPostResponse{typ: "FailedSearchResponse", FailedSearchResponse: value}
}

func (s *SearchSimilarPostResponse) GetSearchResponse() *SearchResponse {
	if s == nil {
		return nil
	}
	return s.SearchResponse
}

func (s *SearchSimilarPostResponse) GetFailedSearchResponse() *FailedSearchResponse {
	if s == nil {
		return nil
	}
	return s.FailedSearchResponse
}

func (s *SearchSimilarPostResponse) UnmarshalJSON(data []byte) error {
	valueSearchResponse := new(SearchResponse)
	if err := json.Unmarshal(data, &valueSearchResponse); err == nil {
		s.typ = "SearchResponse"
		s.SearchResponse = valueSearchResponse
		return nil
	}
	valueFailedSearchResponse := new(FailedSearchResponse)
	if err := json.Unmarshal(data, &valueFailedSearchResponse); err == nil {
		s.typ = "FailedSearchResponse"
		s.FailedSearchResponse = valueFailedSearchResponse
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, s)
}

func (s SearchSimilarPostResponse) MarshalJSON() ([]byte, error) {
	if s.typ == "SearchResponse" || s.SearchResponse != nil {
		return json.Marshal(s.SearchResponse)
	}
	if s.typ == "FailedSearchResponse" || s.FailedSearchResponse != nil {
		return json.Marshal(s.FailedSearchResponse)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", s)
}

type SearchSimilarPostResponseVisitor interface {
	VisitSearchResponse(*SearchResponse) error
	VisitFailedSearchResponse(*FailedSearchResponse) error
}

func (s *SearchSimilarPostResponse) Accept(visitor SearchSimilarPostResponseVisitor) error {
	if s.typ == "SearchResponse" || s.SearchResponse != nil {
		return visitor.VisitSearchResponse(s.SearchResponse)
	}
	if s.typ == "FailedSearchResponse" || s.FailedSearchResponse != nil {
		return visitor.VisitFailedSearchResponse(s.FailedSearchResponse)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", s)
}
