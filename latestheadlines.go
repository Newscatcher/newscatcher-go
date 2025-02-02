// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
)

type LatestHeadlinesGetRequest struct {
	// The time period for which you want to get the latest headlines.
	//
	// Format examples:
	// - `7d`: Last seven days
	// - `30d`: Last 30 days
	// - `1h`: Last hour
	// - `24h`: Last 24 hours
	When *string `json:"-" url:"when,omitempty"`
	// If true, the `from_` and `to_` parameters use article parse dates instead of published dates. Additionally, the `parse_date` variable is added to the output for each article object.
	ByParseDate *bool `json:"-" url:"by_parse_date,omitempty"`
	// The language(s) of the search. The only accepted format is the two-letter [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code. To select multiple languages, use a comma-separated string.
	//
	// Example: `"en, es"`
	//
	// To learn more, see [Enumerated parameters > Language](/docs/v3/api-reference/overview/enumerated-parameters#language-lang-and-not-lang).
	Lang *string `json:"-" url:"lang,omitempty"`
	// The language(s) to exclude from the search. The accepted format is the two-letter [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code. To exclude multiple languages, use a comma-separated string.
	//
	// Example: `"fr, de"`
	//
	// To learn more, see [Enumerated parameters > Language](/docs/v3/api-reference/overview/enumerated-parameters#language-lang-and-not-lang).
	NotLang *string `json:"-" url:"not_lang,omitempty"`
	// The countries where the news publisher is located. The accepted format is the two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. To select multiple countries, use a comma-separated string.
	//
	// Example: `"US, CA"`
	//
	// To learn more, see [Enumerated parameters > Country](/docs/v3/api-reference/overview/enumerated-parameters#country-country-and-not-country).
	Countries *string `json:"-" url:"countries,omitempty"`
	// The publisher location countries to exclude from the search. The accepted format is the two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. To exclude multiple countries, use a comma-separated string.
	//
	// Example:`"US, CA"`
	//
	// To learn more, see [Enumerated parameters > Country](/docs/v3/api-reference/overview/enumerated-parameters#country-country-and-not-country).
	NotCountries *string `json:"-" url:"not_countries,omitempty"`
	// Predefined top news sources per country.
	//
	// Format: start with the word `top`, followed by the number of desired sources, and then the two-letter country code [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). Multiple countries with the number of top sources can be specified as a comma-separated string.
	//
	// Examples:
	// - `"top 100 US"`
	// - `"top 33 AT"`
	// - `"top 50 US, top 20 GB"`
	// - `"top 33 AT, top 50 IT"`
	PredefinedSources *string `json:"-" url:"predefined_sources,omitempty"`
	// One or more news sources to narrow down the search. The format must be a domain URL. Subdomains, such as `finance.yahoo.com`, are also acceptable.To specify multiple sources, use a comma-separated string.
	//
	// Examples:
	// - `"nytimes.com"`
	// - `"theguardian.com, finance.yahoo.com"`
	Sources *string `json:"-" url:"sources,omitempty"`
	// The news sources to exclude from the search. To exclude multiple sources, use a comma-separated string.
	//
	// Example: `"cnn.com, wsj.com"`
	NotSources *string `json:"-" url:"not_sources,omitempty"`
	// The list of author names to exclude from your search. To exclude articles by specific authors, use a comma-separated string.
	//
	// Example: `"John Doe, Jane Doe"`
	NotAuthorName *string `json:"-" url:"not_author_name,omitempty"`
	// If true, limits the search to sources ranked in the top 1 million online websites. If false, includes unranked sources which are assigned a rank of 999999.
	RankedOnly *bool `json:"-" url:"ranked_only,omitempty"`
	// If true, only returns articles that were posted on the home page of a given news domain.
	IsHeadline *bool `json:"-" url:"is_headline,omitempty"`
	// If true, returns only opinion pieces. If false, excludes opinion-based articles and returns news only.
	IsOpinion *bool `json:"-" url:"is_opinion,omitempty"`
	// If false, returns only articles that have publicly available complete content. Some publishers partially block content, so this setting ensures that only full articles are retrieved.
	IsPaidContent *bool `json:"-" url:"is_paid_content,omitempty"`
	// The categorical URL(s) to filter your search. To filter your search by multiple categorical URLs, use a comma-separated string.
	//
	// Example: `"wsj.com/politics, wsj.com/tech"`
	ParentUrl *string `json:"-" url:"parent_url,omitempty"`
	// The complete URL(s) mentioned in the article. For multiple URLs, use a comma-separated string.
	//
	// Example: `"https://aiindex.stanford.edu/report, https://www.stateof.ai"`
	//
	// For more details, see [Search by URL](/docs/v3/documentation/how-to/search-by-url).
	AllLinks *string `json:"-" url:"all_links,omitempty"`
	// The domain(s) mentioned in the article. For multiple domains, use a comma-separated string.
	//
	// Example: `"who.int, nih.gov"`
	//
	// For more details, see [Search by URL](/docs/v3/documentation/how-to/search-by-url).
	AllDomainLinks *string `json:"-" url:"all_domain_links,omitempty"`
	// The minimum number of words an article must contain. To be used for avoiding articles with small content.
	WordCountMin *int `json:"-" url:"word_count_min,omitempty"`
	// The maximum number of words an article can contain. To be used for avoiding articles with large content.
	WordCountMax *int `json:"-" url:"word_count_max,omitempty"`
	// The page number to scroll through the results. Use for pagination, as a single API response can return up to 1,000 articles.
	//
	// For details, see [How to paginate large datasets](https://www.newscatcherapi.com/docs/v3/documentation/how-to/paginate-large-datasets).
	Page *int `json:"-" url:"page,omitempty"`
	// The number of articles to return per page.
	PageSize *int `json:"-" url:"page_size,omitempty"`
	// Determines whether to group similar articles into clusters. If true, the API returns clustered results.
	//
	// To learn more, see [Clustering news articles](/docs/v3/documentation/guides-and-concepts/clustering-news-articles).
	ClusteringEnabled *bool `json:"-" url:"clustering_enabled,omitempty"`
	// Specifies which part of the article to use for determining similarity when clustering.
	//
	// Possible values are:
	// - `content`: Uses the full article content (default).
	// - `title`: Uses only the article title.
	// - `summary`: Uses the article summary.
	//
	// To learn more, see [Clustering news articles](/docs/v3/documentation/guides-and-concepts/clustering-news-articles).
	ClusteringVariable *LatestHeadlinesGetRequestClusteringVariable `json:"-" url:"clustering_variable,omitempty"`
	// Sets the similarity threshold for grouping articles into clusters. A lower value creates more inclusive clusters, while a higher value requires greater similarity between articles.
	//
	// Examples:
	// - `0.3`: Results in larger, more diverse clusters.
	// - `0.6`: Balances cluster size and article similarity (default).
	// - `0.9`: Creates smaller, tightly related clusters.
	//
	// To learn more, see [Clustering news articles](/docs/v3/documentation/guides-and-concepts/clustering-news-articles).
	ClusteringThreshold *float64 `json:"-" url:"clustering_threshold,omitempty"`
	// If true, includes an NLP layer with each article in the response. This layer provides enhanced information such as theme classification, article summary, sentiment analysis, tags, and named entity recognition.
	//
	// The NLP layer includes:
	// - Theme: General topic of the article.
	// - Summary: A concise overview of the article content.
	// - Sentiment: Separate scores for title and content (range: -1 to 1).
	// - Named entities: Identified persons (PER), organizations (ORG), locations (LOC), and miscellaneous entities (MISC).
	// - IPTC tags: Standardized news category tags.
	// - IAB tags: Content categories for digital advertising.
	//
	// **Note**: The `include_nlp_data` parameter is only available if NLP is included in your subscription plan.
	//
	// To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
	IncludeNlpData *bool `json:"-" url:"include_nlp_data,omitempty"`
	// If true, filters the results to include only articles with an NLP layer. This allows you to focus on articles that have been processed with advanced NLP techniques.
	//
	// **Note**: The `has_nlp` parameter is only available if NLP is included in your subscription plan.
	//
	// To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
	HasNlp *bool `json:"-" url:"has_nlp,omitempty"`
	// Filters articles based on their general topic, as determined by NLP analysis. To select multiple themes, use a comma-separated string.
	//
	// Example: `"Finance, Tech"`
	//
	// **Note**: The `theme` parameter is only available if NLP is included in your subscription plan.
	//
	// To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
	//
	// Available options: `Business`, `Economics`, `Entertainment`, `Finance`, `Health`, `Politics`, `Science`, `Sports`, `Tech`, `Crime`, `Financial Crime`, `Lifestyle`, `Automotive`, `Travel`, `Weather`, `General`.
	Theme *string `json:"-" url:"theme,omitempty"`
	// Inverse of the `theme` parameter. Excludes articles based on their general topic, as determined by NLP analysis. To exclude multiple themes, use a comma-separated string.
	//
	// Example: `"Crime, Tech"`
	//
	// **Note**: The `not_theme` parameter is only available if NLP is included in your subscription plan.
	//
	// To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
	NotTheme *string `json:"-" url:"not_theme,omitempty"`
	// Filters articles that mention specific organization names, as identified by NLP analysis. To specify multiple organizations, use a comma-separated string.
	//
	// Example: `"Apple, Microsoft"`
	//
	// **Note**: The `ORG_entity_name` parameter is only available if NLP is included in your subscription plan.
	//
	// To learn more, see [Search by entity](/docs/v3/documentation/how-to/search-by-entity).
	OrgEntityName *string `json:"-" url:"ORG_entity_name,omitempty"`
	// Filters articles that mention specific person names, as identified by NLP analysis. To specify multiple names, use a comma-separated string.
	//
	// Example: `"Elon Musk, Jeff Bezos"`
	//
	// **Note**: The `PER_entity_name` parameter is only available if NLP is included in your subscription plan.
	//
	// To learn more, see [Search by entity](/docs/v3/documentation/how-to/search-by-entity).
	PerEntityName *string `json:"-" url:"PER_entity_name,omitempty"`
	// Filters articles that mention specific location names, as identified by NLP analysis. To specify multiple locations, use a comma-separated string.
	//
	// Example: `"California, New York"`
	//
	// **Note**: The `LOC_entity_name` parameter is only available if NLP is included in your subscription plan.
	//
	// To learn more, see [Search by entity](/docs/v3/documentation/how-to/search-by-entity).
	LocEntityName *string `json:"-" url:"LOC_entity_name,omitempty"`
	// Filters articles that mention other named entities not falling under person, organization, or location categories. Includes events, nationalities, products, works of art, and more. To specify multiple entities, use a comma-separated string.
	//
	// Example: `"Bitcoin, Blockchain"`
	//
	// **Note**: The `MISC_entity_name` parameter is only available if NLP is included in your subscription plan.
	//
	// To learn more, see [Search by entity](/docs/v3/documentation/how-to/search-by-entity).
	MiscEntityName *string `json:"-" url:"MISC_entity_name,omitempty"`
	// Filters articles based on the minimum sentiment score of their titles.
	//
	// Range is `-1.0` to `1.0`, where:
	// - Negative values indicate negative sentiment.
	// - Positive values indicate positive sentiment.
	// - Values close to 0 indicate neutral sentiment.
	//
	// **Note**: The `title_sentiment_min` parameter is only available if NLP is included in your subscription plan.
	//
	// To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
	TitleSentimentMin *float64 `json:"-" url:"title_sentiment_min,omitempty"`
	// Filters articles based on the maximum sentiment score of their titles.
	//
	// Range is `-1.0` to `1.0`, where:
	// - Negative values indicate negative sentiment.
	// - Positive values indicate positive sentiment.
	// - Values close to 0 indicate neutral sentiment.
	//
	// **Note**: The `title_sentiment_max` parameter is only available if NLP is included in your subscription plan.
	//
	// To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
	TitleSentimentMax *float64 `json:"-" url:"title_sentiment_max,omitempty"`
	// Filters articles based on the minimum sentiment score of their content.
	//
	// Range is `-1.0` to `1.0`, where:
	// - Negative values indicate negative sentiment.
	// - Positive values indicate positive sentiment.
	// - Values close to 0 indicate neutral sentiment.
	//
	// **Note**: The `content_sentiment_min` parameter is only available if NLP is included in your subscription plan.
	//
	// To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
	ContentSentimentMin *float64 `json:"-" url:"content_sentiment_min,omitempty"`
	// Filters articles based on the maximum sentiment score of their content.
	//
	// Range is `-1.0` to `1.0`, where:
	// - Negative values indicate negative sentiment.
	// - Positive values indicate positive sentiment.
	// - Values close to 0 indicate neutral sentiment.
	//
	// **Note**: The `content_sentiment_max` parameter is only available if NLP is included in your subscription plan.
	//
	// To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
	ContentSentimentMax *float64 `json:"-" url:"content_sentiment_max,omitempty"`
	// Filters articles based on International Press Telecommunications Council (IPTC) media topic tags. To specify multiple IPTC tags, use a comma-separated string of tag IDs.
	//
	// Example: `"20000199, 20000209"`
	//
	// **Note**: The `iptc_tags` parameter is only available if tags are included in your subscription plan.
	//
	// To learn more, see [IPTC Media Topic NewsCodes](https://www.iptc.org/std/NewsCodes/treeview/mediatopic/mediatopic-en-GB.html).
	IptcTags *string `json:"-" url:"iptc_tags,omitempty"`
	// Inverse of the `iptc_tags` parameter. Excludes articles based on International Press Telecommunications Council (IPTC) media topic tags. To specify multiple IPTC tags to exclude, use a comma-separated string of tag IDs.
	//
	// Example: `"20000205, 20000209"`
	//
	// **Note**: The `not_iptc_tags` parameter is only available if tags are included in your subscription plan.
	//
	// To learn more, see [IPTC Media Topic NewsCodes](https://www.iptc.org/std/NewsCodes/treeview/mediatopic/mediatopic-en-GB.html).
	NotIptcTags *string `json:"-" url:"not_iptc_tags,omitempty"`
	// Filters articles based on Interactive Advertising Bureau (IAB) content categories. These tags provide a standardized taxonomy for digital advertising content categorization. To specify multiple IAB categories, use a comma-separated string.
	//
	// Example: `"Business, Events"`
	//
	// **Note**: The `iab_tags` parameter is only available if tags are included in your subscription plan.
	//
	// To learn more, see the [IAB Content taxonomy](https://iabtechlab.com/standards/content-taxonomy/).
	IabTags *string `json:"-" url:"iab_tags,omitempty"`
	// Inverse of the `iab_tags` parameter. Excludes articles based on Interactive Advertising Bureau (IAB) content categories. These tags provide a standardized taxonomy for digital advertising content categorization. To specify multiple IAB categories to exclude, use a comma-separated string.
	//
	// Example: `"Agriculture, Metals"`
	//
	// **Note**: The `not_iab_tags` parameter is only available if tags are included in your subscription plan.
	//
	// To learn more, see the [IAB Content taxonomy](https://iabtechlab.com/standards/content-taxonomy/).
	NotIabTags *string `json:"-" url:"not_iab_tags,omitempty"`
	// Filters articles based on provided taxonomy that is tailored to your specific needs and is accessible only with your API key. To specify tags, use the following pattern:
	//
	// - `custom_tags.taxonomy=Tag1,Tag2,Tag3`, where `taxonomy` is the taxonomy name and `Tag1,Tag2,Tag3` is a comma-separated list of tags.
	//
	// Example: `custom_tags.industry="Manufacturing, Supply Chain, Logistics"`
	//
	// To learn more, see the [Custom tags](/docs/v3/documentation/guides-and-concepts/custom-tags).
	CustomTags *string `json:"-" url:"custom_tags,omitempty"`
}

type LatestHeadlinesPostRequest struct {
	When                *When                `json:"when,omitempty" url:"-"`
	ByParseDate         *ByParseDate         `json:"by_parse_date,omitempty" url:"-"`
	Lang                *Lang                `json:"lang,omitempty" url:"-"`
	NotLang             *NotLang             `json:"not_lang,omitempty" url:"-"`
	Countries           *Countries           `json:"countries,omitempty" url:"-"`
	NotCountries        *NotCountries        `json:"not_countries,omitempty" url:"-"`
	PredefinedSources   *PredefinedSources   `json:"predefined_sources,omitempty" url:"-"`
	Sources             *Sources             `json:"sources,omitempty" url:"-"`
	NotSources          *NotSources          `json:"not_sources,omitempty" url:"-"`
	NotAuthorName       *NotAuthorName       `json:"not_author_name,omitempty" url:"-"`
	RankedOnly          *RankedOnly          `json:"ranked_only,omitempty" url:"-"`
	IsHeadline          *IsHeadline          `json:"is_headline,omitempty" url:"-"`
	IsOpinion           *IsOpinion           `json:"is_opinion,omitempty" url:"-"`
	IsPaidContent       *IsPaidContent       `json:"is_paid_content,omitempty" url:"-"`
	ParentUrl           *ParentUrl           `json:"parent_url,omitempty" url:"-"`
	AllLinks            *AllLinks            `json:"all_links,omitempty" url:"-"`
	AllDomainLinks      *AllDomainLinks      `json:"all_domain_links,omitempty" url:"-"`
	WordCountMin        *WordCountMin        `json:"word_count_min,omitempty" url:"-"`
	WordCountMax        *WordCountMax        `json:"word_count_max,omitempty" url:"-"`
	Page                *Page                `json:"page,omitempty" url:"-"`
	PageSize            *PageSize            `json:"page_size,omitempty" url:"-"`
	ClusteringEnabled   *ClusteringEnabled   `json:"clustering_enabled,omitempty" url:"-"`
	ClusteringVariable  *ClusteringVariable  `json:"clustering_variable,omitempty" url:"-"`
	ClusteringThreshold *ClusteringThreshold `json:"clustering_threshold,omitempty" url:"-"`
	IncludeNlpData      *IncludeNlpData      `json:"include_nlp_data,omitempty" url:"-"`
	HasNlp              *HasNlp              `json:"has_nlp,omitempty" url:"-"`
	Theme               *Theme               `json:"theme,omitempty" url:"-"`
	NotTheme            *NotTheme            `json:"not_theme,omitempty" url:"-"`
	OrgEntityName       *OrgEntityName       `json:"ORG_entity_name,omitempty" url:"-"`
	PerEntityName       *PerEntityName       `json:"PER_entity_name,omitempty" url:"-"`
	LocEntityName       *LocEntityName       `json:"LOC_entity_name,omitempty" url:"-"`
	MiscEntityName      *MiscEntityName      `json:"MISC_entity_name,omitempty" url:"-"`
	TitleSentimentMin   *TitleSentimentMin   `json:"title_sentiment_min,omitempty" url:"-"`
	TitleSentimentMax   *TitleSentimentMax   `json:"title_sentiment_max,omitempty" url:"-"`
	ContentSentimentMin *ContentSentimentMin `json:"content_sentiment_min,omitempty" url:"-"`
	ContentSentimentMax *ContentSentimentMax `json:"content_sentiment_max,omitempty" url:"-"`
	IptcTags            *IptcTags            `json:"iptc_tags,omitempty" url:"-"`
	NotIptcTags         *NotIptcTags         `json:"not_iptc_tags,omitempty" url:"-"`
	IabTags             *IabTags             `json:"iab_tags,omitempty" url:"-"`
	NotIabTags          *NotIabTags          `json:"not_iab_tags,omitempty" url:"-"`
	CustomTags          *CustomTags          `json:"custom_tags,omitempty" url:"-"`
}

// The time period for which you want to get the latest headlines.
//
// Format examples:
// - `7d`: Last seven days
// - `30d`: Last 30 days
// - `1h`: Last hour
// - `24h`: Last 24 hours
type When = string

type LatestHeadlinesGetRequestClusteringVariable string

const (
	LatestHeadlinesGetRequestClusteringVariableContent LatestHeadlinesGetRequestClusteringVariable = "content"
	LatestHeadlinesGetRequestClusteringVariableTitle   LatestHeadlinesGetRequestClusteringVariable = "title"
	LatestHeadlinesGetRequestClusteringVariableSummary LatestHeadlinesGetRequestClusteringVariable = "summary"
)

func NewLatestHeadlinesGetRequestClusteringVariableFromString(s string) (LatestHeadlinesGetRequestClusteringVariable, error) {
	switch s {
	case "content":
		return LatestHeadlinesGetRequestClusteringVariableContent, nil
	case "title":
		return LatestHeadlinesGetRequestClusteringVariableTitle, nil
	case "summary":
		return LatestHeadlinesGetRequestClusteringVariableSummary, nil
	}
	var t LatestHeadlinesGetRequestClusteringVariable
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l LatestHeadlinesGetRequestClusteringVariable) Ptr() *LatestHeadlinesGetRequestClusteringVariable {
	return &l
}

type LatestHeadlinesGetResponse struct {
	SearchResponseDto          *SearchResponseDto
	ClusteredSearchResponseDto *ClusteredSearchResponseDto

	typ string
}

func NewLatestHeadlinesGetResponseFromSearchResponseDto(value *SearchResponseDto) *LatestHeadlinesGetResponse {
	return &LatestHeadlinesGetResponse{typ: "SearchResponseDto", SearchResponseDto: value}
}

func NewLatestHeadlinesGetResponseFromClusteredSearchResponseDto(value *ClusteredSearchResponseDto) *LatestHeadlinesGetResponse {
	return &LatestHeadlinesGetResponse{typ: "ClusteredSearchResponseDto", ClusteredSearchResponseDto: value}
}

func (l *LatestHeadlinesGetResponse) GetSearchResponseDto() *SearchResponseDto {
	if l == nil {
		return nil
	}
	return l.SearchResponseDto
}

func (l *LatestHeadlinesGetResponse) GetClusteredSearchResponseDto() *ClusteredSearchResponseDto {
	if l == nil {
		return nil
	}
	return l.ClusteredSearchResponseDto
}

func (l *LatestHeadlinesGetResponse) UnmarshalJSON(data []byte) error {
	valueSearchResponseDto := new(SearchResponseDto)
	if err := json.Unmarshal(data, &valueSearchResponseDto); err == nil {
		l.typ = "SearchResponseDto"
		l.SearchResponseDto = valueSearchResponseDto
		return nil
	}
	valueClusteredSearchResponseDto := new(ClusteredSearchResponseDto)
	if err := json.Unmarshal(data, &valueClusteredSearchResponseDto); err == nil {
		l.typ = "ClusteredSearchResponseDto"
		l.ClusteredSearchResponseDto = valueClusteredSearchResponseDto
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, l)
}

func (l LatestHeadlinesGetResponse) MarshalJSON() ([]byte, error) {
	if l.typ == "SearchResponseDto" || l.SearchResponseDto != nil {
		return json.Marshal(l.SearchResponseDto)
	}
	if l.typ == "ClusteredSearchResponseDto" || l.ClusteredSearchResponseDto != nil {
		return json.Marshal(l.ClusteredSearchResponseDto)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", l)
}

type LatestHeadlinesGetResponseVisitor interface {
	VisitSearchResponseDto(*SearchResponseDto) error
	VisitClusteredSearchResponseDto(*ClusteredSearchResponseDto) error
}

func (l *LatestHeadlinesGetResponse) Accept(visitor LatestHeadlinesGetResponseVisitor) error {
	if l.typ == "SearchResponseDto" || l.SearchResponseDto != nil {
		return visitor.VisitSearchResponseDto(l.SearchResponseDto)
	}
	if l.typ == "ClusteredSearchResponseDto" || l.ClusteredSearchResponseDto != nil {
		return visitor.VisitClusteredSearchResponseDto(l.ClusteredSearchResponseDto)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", l)
}

type LatestHeadlinesPostResponse struct {
	SearchResponseDto          *SearchResponseDto
	ClusteredSearchResponseDto *ClusteredSearchResponseDto

	typ string
}

func NewLatestHeadlinesPostResponseFromSearchResponseDto(value *SearchResponseDto) *LatestHeadlinesPostResponse {
	return &LatestHeadlinesPostResponse{typ: "SearchResponseDto", SearchResponseDto: value}
}

func NewLatestHeadlinesPostResponseFromClusteredSearchResponseDto(value *ClusteredSearchResponseDto) *LatestHeadlinesPostResponse {
	return &LatestHeadlinesPostResponse{typ: "ClusteredSearchResponseDto", ClusteredSearchResponseDto: value}
}

func (l *LatestHeadlinesPostResponse) GetSearchResponseDto() *SearchResponseDto {
	if l == nil {
		return nil
	}
	return l.SearchResponseDto
}

func (l *LatestHeadlinesPostResponse) GetClusteredSearchResponseDto() *ClusteredSearchResponseDto {
	if l == nil {
		return nil
	}
	return l.ClusteredSearchResponseDto
}

func (l *LatestHeadlinesPostResponse) UnmarshalJSON(data []byte) error {
	valueSearchResponseDto := new(SearchResponseDto)
	if err := json.Unmarshal(data, &valueSearchResponseDto); err == nil {
		l.typ = "SearchResponseDto"
		l.SearchResponseDto = valueSearchResponseDto
		return nil
	}
	valueClusteredSearchResponseDto := new(ClusteredSearchResponseDto)
	if err := json.Unmarshal(data, &valueClusteredSearchResponseDto); err == nil {
		l.typ = "ClusteredSearchResponseDto"
		l.ClusteredSearchResponseDto = valueClusteredSearchResponseDto
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, l)
}

func (l LatestHeadlinesPostResponse) MarshalJSON() ([]byte, error) {
	if l.typ == "SearchResponseDto" || l.SearchResponseDto != nil {
		return json.Marshal(l.SearchResponseDto)
	}
	if l.typ == "ClusteredSearchResponseDto" || l.ClusteredSearchResponseDto != nil {
		return json.Marshal(l.ClusteredSearchResponseDto)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", l)
}

type LatestHeadlinesPostResponseVisitor interface {
	VisitSearchResponseDto(*SearchResponseDto) error
	VisitClusteredSearchResponseDto(*ClusteredSearchResponseDto) error
}

func (l *LatestHeadlinesPostResponse) Accept(visitor LatestHeadlinesPostResponseVisitor) error {
	if l.typ == "SearchResponseDto" || l.SearchResponseDto != nil {
		return visitor.VisitSearchResponseDto(l.SearchResponseDto)
	}
	if l.typ == "ClusteredSearchResponseDto" || l.ClusteredSearchResponseDto != nil {
		return visitor.VisitClusteredSearchResponseDto(l.ClusteredSearchResponseDto)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", l)
}
