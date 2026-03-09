# Reference
## Search
<details><summary><code>client.Search.Get() -> *v505.SearchGetResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Searches for articles based on specified criteria such as keyword, language, country, source, and more.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.SearchGetRequest{
        Q: `"supply chain" AND Amazon NOT China`,
        SearchIn: v505.String(
            "title_content, title_content_translated",
        ),
        IncludeTranslationFields: v505.Bool(
            true,
        ),
        PredefinedSources: v505.String(
            "top 100 US, top 5 GB",
        ),
        SourceName: v505.String(
            "sport",
        ),
        Sources: v505.String(
            "nytimes.com",
        ),
        NotSources: v505.String(
            "cnn.com",
        ),
        Lang: v505.String(
            "en",
        ),
        NotLang: v505.String(
            "fr",
        ),
        Countries: v505.String(
            "US",
        ),
        NotCountries: v505.String(
            "UK",
        ),
        NotAuthorName: v505.String(
            "John Doe",
        ),
        From: &v505.SearchGetRequestFrom{
            DateTime: v505.MustParseDateTime(
                "2024-07-01T00:00:00Z",
            ),
        },
        To: &v505.SearchGetRequestTo{
            DateTime: v505.MustParseDateTime(
                "2024-07-01T00:00:00Z",
            ),
        },
        ParentUrl: v505.String(
            "https://www.washingtonpost.com/politics",
        ),
        AllLinks: v505.String(
            "https://aiindex.stanford.edu/report",
        ),
        AllDomainLinks: v505.String(
            "nvidia.com",
        ),
        NewsType: v505.String(
            "General News Outlets",
        ),
        IncludeNlpData: v505.Bool(
            true,
        ),
        HasNlp: v505.Bool(
            true,
        ),
        Theme: v505.String(
            "Business,Finance",
        ),
        NotTheme: v505.String(
            "Crime",
        ),
        OrgEntityName: v505.String(
            "Apple",
        ),
        PerEntityName: v505.String(
            "Elon Musk",
        ),
        LocEntityName: v505.String(
            "California",
        ),
        MiscEntityName: v505.String(
            "Bitcoin",
        ),
        IptcTags: v505.String(
            "20000199,20000209",
        ),
        NotIptcTags: v505.String(
            "20000205,20000209",
        ),
        IabTags: v505.String(
            "Business,Events",
        ),
        NotIabTags: v505.String(
            "Agriculture,Metals",
        ),
        CustomTags: v505.String(
            "Tag1,Tag2,Tag3",
        ),
    }
client.Search.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**q:** `v505.Q` 
    
</dd>
</dl>

<dl>
<dd>

**searchIn:** `*v505.SearchIn` 
    
</dd>
</dl>

<dl>
<dd>

**includeTranslationFields:** `*v505.IncludeTranslationFields` 
    
</dd>
</dl>

<dl>
<dd>

**predefinedSources:** `*string` 

Predefined top news sources per country. 

Format: start with the word `top`, followed by the number of desired sources, and then the two-letter country code [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). Multiple countries with the number of top sources can be specified as a comma-separated string.

Examples: 
- `"top 100 US"`
- `"top 33 AT"`
- `"top 50 US, top 20 GB"`
- `"top 33 AT, top 50 IT"`
    
</dd>
</dl>

<dl>
<dd>

**sourceName:** `*string` 

Word or phrase to search within the source names. To specify multiple values, use a comma-separated string.

Example: `"sport, tech"`

**Note**: The search doesn't require an exact match and returns sources containing the specified terms in their names. You can use any word or phrase, like `"sport"` or `"new york times"`. For example, `"sport"` returns sources such as `"Motorsport"`, `"Dot Esport"`, and `"Tuttosport"`.
    
</dd>
</dl>

<dl>
<dd>

**sources:** `*string` 

One or more news sources to narrow down the search. The format must be a domain URL. Subdomains, such as `finance.yahoo.com`, are also acceptable.To specify multiple sources, use a comma-separated string.

Examples:
- `"nytimes.com"`
- `"theguardian.com, finance.yahoo.com"`
    
</dd>
</dl>

<dl>
<dd>

**notSources:** `*string` 

The news sources to exclude from the search. To exclude multiple sources, use a comma-separated string. 

Example: `"cnn.com, wsj.com"`
    
</dd>
</dl>

<dl>
<dd>

**lang:** `*string` 

The language(s) of the search. The only accepted format is the two-letter [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code. To select multiple languages, use a comma-separated string. 

Example: `"en, es"`

To learn more, see [Enumerated parameters > Language](/docs/v3/api-reference/overview/enumerated-parameters#language-lang-and-not-lang).
    
</dd>
</dl>

<dl>
<dd>

**notLang:** `*string` 

The language(s) to exclude from the search. The accepted format is the two-letter [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code. To exclude multiple languages, use a comma-separated string. 

Example: `"fr, de"`

To learn more, see [Enumerated parameters > Language](/docs/v3/api-reference/overview/enumerated-parameters#language-lang-and-not-lang).
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*string` 

The countries where the news publisher is located. The accepted format is the two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. To select multiple countries, use a comma-separated string.

Example: `"US, CA"`

To learn more, see [Enumerated parameters > Country](/docs/v3/api-reference/overview/enumerated-parameters#country-country-and-not-country).
    
</dd>
</dl>

<dl>
<dd>

**notCountries:** `*string` 

The publisher location countries to exclude from the search. The accepted format is the two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. To exclude multiple countries, use a comma-separated string. 

Example:`"US, CA"`

To learn more, see [Enumerated parameters > Country](/docs/v3/api-reference/overview/enumerated-parameters#country-country-and-not-country).
    
</dd>
</dl>

<dl>
<dd>

**notAuthorName:** `*string` 

The list of author names to exclude from your search. To exclude articles by specific authors, use a comma-separated string.

Example: `"John Doe, Jane Doe"`
    
</dd>
</dl>

<dl>
<dd>

**from:** `*v505.SearchGetRequestFrom` 

The starting point in time to search from. Accepts date-time strings in ISO 8601 format and plain text. The default time zone is UTC. 

Formats with examples:
- YYYY-mm-ddTHH:MM:SS: `2024-07-01T00:00:00`
- YYYY-MM-dd: `2024-07-01`
- YYYY/mm/dd HH:MM:SS: `2024/07/01 00:00:00`
- YYYY/mm/dd: `2024/07/01`
- English phrases: `7 day ago`, `today`

**Note**: By default, applied to the publication date of the article. To use the article's parse date instead, set the `by_parse_date` parameter to `true`.
    
</dd>
</dl>

<dl>
<dd>

**to:** `*v505.SearchGetRequestTo` 

The ending point in time to search up to. Accepts date-time strings in ISO 8601 format and plain text. The default time zone is UTC. 

Formats with examples:
- YYYY-mm-ddTHH:MM:SS: `2024-07-01T00:00:00`
- YYYY-MM-dd: `2024-07-01`
- YYYY/mm/dd HH:MM:SS: `2024/07/01 00:00:00`
- YYYY/mm/dd: `2024/07/01`
- English phrases: `1 day ago`, `now`

**Note**: By default, applied to the publication date of the article. To use the article's parse date instead, set the `by_parse_date` parameter to `true`.
    
</dd>
</dl>

<dl>
<dd>

**publishedDatePrecision:** `*v505.SearchGetRequestPublishedDatePrecision` 

The precision of the published date. There are three types:
- `full`: The day and time of an article is correctly identified with the appropriate timezone.
- `timezone unknown`: The day and time of an article is correctly identified without timezone.
- `date`: Only the day is identified without an exact time.
    
</dd>
</dl>

<dl>
<dd>

**byParseDate:** `*bool` — If true, the `from_` and `to_` parameters use article parse dates instead of published dates. Additionally, the `parse_date` variable is added to the output for each article object.
    
</dd>
</dl>

<dl>
<dd>

**sortBy:** `*v505.SearchGetRequestSortBy` 

The sorting order of the results. Possible values are:
- `relevancy`: The most relevant results first.
- `date`: The most recently published results first.
- `rank`: The results from the highest-ranked sources first.
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*bool` — If true, limits the search to sources ranked in the top 1 million online websites. If false, includes unranked sources which are assigned a rank of 999999.
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*int` — The lowest boundary of the rank of a news website to filter by. A lower rank indicates a more popular source.
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*int` — The highest boundary of the rank of a news website to filter by. A lower rank indicates a more popular source.
    
</dd>
</dl>

<dl>
<dd>

**isHeadline:** `*bool` — If true, only returns articles that were posted on the home page of a given news domain.
    
</dd>
</dl>

<dl>
<dd>

**isOpinion:** `*bool` — If true, returns only opinion pieces. If false, excludes opinion-based articles and returns news only.
    
</dd>
</dl>

<dl>
<dd>

**isPaidContent:** `*bool` — If false, returns only articles that have publicly available complete content. Some publishers partially block content, so this setting ensures that only full articles are retrieved.
    
</dd>
</dl>

<dl>
<dd>

**parentUrl:** `*string` 

The categorical URL(s) to filter your search. To filter your search by multiple categorical URLs, use a comma-separated string.

Example: `"wsj.com/politics, wsj.com/tech"`
    
</dd>
</dl>

<dl>
<dd>

**allLinks:** `*string` 

The complete URL(s) mentioned in the article. For multiple URLs, use a comma-separated string.

Example: `"https://aiindex.stanford.edu/report, https://www.stateof.ai"`

For more details, see [Search by URL](/docs/v3/documentation/how-to/search-by-url).
    
</dd>
</dl>

<dl>
<dd>

**allDomainLinks:** `*string` 

The domain(s) mentioned in the article. For multiple domains, use a comma-separated string.

Example: `"who.int, nih.gov"`

For more details, see [Search by URL](/docs/v3/documentation/how-to/search-by-url).
    
</dd>
</dl>

<dl>
<dd>

**additionalDomainInfo:** `*bool` 

If true, includes additional domain information in the response for each article:
- `is_news_domain`: Boolean indicating if the source is a news domain.
- `news_domain_type`: Type of news domain (e.g., `"Original Content"`).
- `news_type`: Category of news (e.g., `"News and Blogs"`).
    
</dd>
</dl>

<dl>
<dd>

**isNewsDomain:** `*bool` — If true, filters results to include only news domains.
    
</dd>
</dl>

<dl>
<dd>

**newsDomainType:** `*v505.SearchGetRequestNewsDomainType` 

Filters results based on the news domain type. Possible values are:
- `Original Content`: Sources that produce their own content.
- `Aggregator`: Sources that collect content from various other sources.
- `Press Releases`: Sources primarily publishing press releases.
- `Republisher`: Sources that republish content from other sources.
- `Other`: Sources that don't fit into main categories.
    
</dd>
</dl>

<dl>
<dd>

**newsType:** `*string` 

Filters results based on the news type. Multiple types can be specified using a comma-separated string.

Example: `"General News Outlets,Tech News and Updates"`

For a complete list of available news types, see [Enumerated parameters > News type](/docs/v3/api-reference/overview/enumerated-parameters#news-type-news-type).
    
</dd>
</dl>

<dl>
<dd>

**wordCountMin:** `*int` — The minimum number of words an article must contain. To be used for avoiding articles with small content.
    
</dd>
</dl>

<dl>
<dd>

**wordCountMax:** `*int` — The maximum number of words an article can contain. To be used for avoiding articles with large content.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 

The page number to scroll through the results. Use for pagination, as a single API response can return up to 1,000 articles. 

For details, see [How to paginate large datasets](https://www.newscatcherapi.com/docs/v3/documentation/how-to/paginate-large-datasets).
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The number of articles to return per page.
    
</dd>
</dl>

<dl>
<dd>

**clusteringEnabled:** `*bool` 

Determines whether to group similar articles into clusters. If true, the API returns clustered results.

To learn more, see [Clustering news articles](/docs/v3/documentation/guides-and-concepts/clustering-news-articles).
    
</dd>
</dl>

<dl>
<dd>

**clusteringVariable:** `*v505.SearchGetRequestClusteringVariable` 

Specifies which part of the article to use for determining similarity when clustering.

Possible values are:
- `content`: Uses the full article content (default).
- `title`: Uses only the article title.
- `summary`: Uses the article summary.

To learn more, see [Clustering news articles](/docs/v3/documentation/guides-and-concepts/clustering-news-articles).
    
</dd>
</dl>

<dl>
<dd>

**clusteringThreshold:** `*float64` 

Sets the similarity threshold for grouping articles into clusters. A lower value creates more inclusive clusters, while a higher value requires greater similarity between articles.

Examples:
- `0.3`: Results in larger, more diverse clusters.
- `0.6`: Balances cluster size and article similarity (default).
- `0.9`: Creates smaller, tightly related clusters.

To learn more, see [Clustering news articles](/docs/v3/documentation/guides-and-concepts/clustering-news-articles).
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*v505.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*v505.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*string` 

Filters articles based on their general topic, as determined by NLP analysis. To select multiple themes, use a comma-separated string.

Example: `"Finance, Tech"`

**Note**: The `theme` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).

Available options: `Business`, `Economics`, `Entertainment`, `Finance`, `Health`, `Politics`, `Science`, `Sports`, `Tech`, `Crime`, `Financial Crime`, `Lifestyle`, `Automotive`, `Travel`, `Weather`, `General`.
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*string` 

Inverse of the `theme` parameter. Excludes articles based on their general topic, as determined by NLP analysis. To exclude multiple themes, use a comma-separated string. 

Example: `"Crime, Tech"`

**Note**: The `not_theme` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**orgEntityName:** `*string` 

Filters articles that mention specific organization names, as identified by NLP analysis. To specify multiple organizations, use a comma-separated string. To search named entities in translations, combine with the translation options of the `search_in` parameter (e.g., `title_content_translated`).

Example: `"Apple, Microsoft"`

**Note**: The `ORG_entity_name` parameter is only available if NLP is included in your subscription plan.

To learn more, see [Search by entity](/docs/v3/documentation/how-to/search-by-entity).
    
</dd>
</dl>

<dl>
<dd>

**perEntityName:** `*string` 

Filters articles that mention specific person names, as identified by NLP analysis. To specify multiple names, use a comma-separated string. To search named entities in translations, combine with the translation options of the `search_in` parameter (e.g., `title_content_translated`).

Example: `"Elon Musk, Jeff Bezos"`

**Note**: The `PER_entity_name` parameter is only available if NLP is included in your subscription plan.

To learn more, see [Search by entity](/docs/v3/documentation/how-to/search-by-entity).
    
</dd>
</dl>

<dl>
<dd>

**locEntityName:** `*string` 

Filters articles that mention specific location names, as identified by NLP analysis. To specify multiple locations, use a comma-separated string. To search named entities in translations, combine with the translation options of the `search_in` parameter (e.g., `title_content_translated`).

Example: `"California, New York"`

**Note**: The `LOC_entity_name` parameter is only available if NLP is included in your subscription plan.

To learn more, see [Search by entity](/docs/v3/documentation/how-to/search-by-entity).
    
</dd>
</dl>

<dl>
<dd>

**miscEntityName:** `*string` 

Filters articles that mention other named entities not falling under person, organization, or location categories. Includes events, nationalities, products, works of art, and more. To specify multiple entities, use a comma-separated string. To search named entities in translations, combine with the translation options of the `search_in` parameter (e.g., `title_content_translated`).

Example: `"Bitcoin, Blockchain"`

**Note**: The `MISC_entity_name` parameter is only available if NLP is included in your subscription plan.

To learn more, see [Search by entity](/docs/v3/documentation/how-to/search-by-entity).
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*float64` 

Filters articles based on the minimum sentiment score of their titles.

Range is `-1.0` to `1.0`, where:
- Negative values indicate negative sentiment.
- Positive values indicate positive sentiment.
- Values close to 0 indicate neutral sentiment.

**Note**: The `title_sentiment_min` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*float64` 

Filters articles based on the maximum sentiment score of their titles.

Range is `-1.0` to `1.0`, where:
- Negative values indicate negative sentiment.
- Positive values indicate positive sentiment.
- Values close to 0 indicate neutral sentiment.

**Note**: The `title_sentiment_max` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*float64` 

Filters articles based on the minimum sentiment score of their content.

Range is `-1.0` to `1.0`, where:
- Negative values indicate negative sentiment.
- Positive values indicate positive sentiment.
- Values close to 0 indicate neutral sentiment.

**Note**: The `content_sentiment_min` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*float64` 

Filters articles based on the maximum sentiment score of their content.

Range is `-1.0` to `1.0`, where:
- Negative values indicate negative sentiment.
- Positive values indicate positive sentiment.
- Values close to 0 indicate neutral sentiment.

**Note**: The `content_sentiment_max` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**iptcTags:** `*string` 

Filters articles based on International Press Telecommunications Council (IPTC) media topic tags. To specify multiple IPTC tags, use a comma-separated string of tag IDs. 

Example: `"20000199, 20000209"`

**Note**: The `iptc_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see [IPTC Media Topic NewsCodes](https://www.iptc.org/std/NewsCodes/treeview/mediatopic/mediatopic-en-GB.html).
    
</dd>
</dl>

<dl>
<dd>

**notIptcTags:** `*string` 

Inverse of the `iptc_tags` parameter. Excludes articles based on International Press Telecommunications Council (IPTC) media topic tags. To specify multiple IPTC tags to exclude, use a comma-separated string of tag IDs. 

Example: `"20000205, 20000209"`

**Note**: The `not_iptc_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see [IPTC Media Topic NewsCodes](https://www.iptc.org/std/NewsCodes/treeview/mediatopic/mediatopic-en-GB.html).
    
</dd>
</dl>

<dl>
<dd>

**iabTags:** `*string` 

Filters articles based on Interactive Advertising Bureau (IAB) content categories. These tags provide a standardized taxonomy for digital advertising content categorization. To specify multiple IAB categories, use a comma-separated string. 

Example: `"Business, Events"`

**Note**: The `iab_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see the [IAB Content taxonomy](https://iabtechlab.com/standards/content-taxonomy/).
    
</dd>
</dl>

<dl>
<dd>

**notIabTags:** `*string` 

Inverse of the `iab_tags` parameter. Excludes articles based on Interactive Advertising Bureau (IAB) content categories. These tags provide a standardized taxonomy for digital advertising content categorization. To specify multiple IAB categories to exclude, use a comma-separated string. 

Example: `"Agriculture, Metals"`

**Note**: The `not_iab_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see the [IAB Content taxonomy](https://iabtechlab.com/standards/content-taxonomy/).
    
</dd>
</dl>

<dl>
<dd>

**customTags:** `*string` 

Filters articles based on provided taxonomy that is tailored to your specific needs and is accessible only with your API key. To specify tags, use the following pattern: 

- `custom_tags.taxonomy=Tag1,Tag2,Tag3`, where `taxonomy` is the taxonomy name and `Tag1,Tag2,Tag3` is a comma-separated list of tags.

Example: `custom_tags.industry="Manufacturing, Supply Chain, Logistics"`

To learn more, see the [Custom tags](/docs/v3/documentation/guides-and-concepts/custom-tags).
    
</dd>
</dl>

<dl>
<dd>

**excludeDuplicates:** `*bool` 

If true, excludes duplicate and highly similar articles from the search results. If false, returns all relevant articles, including duplicates. 

To learn more, see [Articles deduplication](/docs/v3/documentation/guides-and-concepts/articles-deduplication).
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*bool` — If true, returns only articles/sources that comply with the publisher's robots.txt rules. If false, returns only articles/sources that do not comply with robots.txt rules. If omitted, returns all articles/sources regardless of compliance status.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Search.Post(request) -> *v505.SearchPostResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Searches for articles based on specified criteria such as keyword, language, country, source, and more.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.SearchPostRequest{
        Q: `"supply chain" AND Amazon NOT China`,
        PageSize: v505.Int(
            1,
        ),
    }
client.Search.Post(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**q:** `v505.Q` 
    
</dd>
</dl>

<dl>
<dd>

**searchIn:** `*v505.SearchIn` 
    
</dd>
</dl>

<dl>
<dd>

**includeTranslationFields:** `*v505.IncludeTranslationFields` 
    
</dd>
</dl>

<dl>
<dd>

**predefinedSources:** `*v505.PredefinedSources` 
    
</dd>
</dl>

<dl>
<dd>

**sourceName:** `*v505.SourceName` 
    
</dd>
</dl>

<dl>
<dd>

**sources:** `*v505.Sources` 
    
</dd>
</dl>

<dl>
<dd>

**notSources:** `*v505.NotSources` 
    
</dd>
</dl>

<dl>
<dd>

**lang:** `*v505.Lang` 
    
</dd>
</dl>

<dl>
<dd>

**notLang:** `*v505.NotLang` 
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*v505.Countries` 
    
</dd>
</dl>

<dl>
<dd>

**notCountries:** `*v505.NotCountries` 
    
</dd>
</dl>

<dl>
<dd>

**notAuthorName:** `*v505.NotAuthorName` 
    
</dd>
</dl>

<dl>
<dd>

**from:** `*v505.From` 
    
</dd>
</dl>

<dl>
<dd>

**to:** `*v505.To` 
    
</dd>
</dl>

<dl>
<dd>

**publishedDatePrecision:** `*v505.PublishedDatePrecision` 
    
</dd>
</dl>

<dl>
<dd>

**byParseDate:** `*v505.ByParseDate` 
    
</dd>
</dl>

<dl>
<dd>

**sortBy:** `*v505.SortBy` 
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*v505.RankedOnly` 
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*v505.FromRank` 
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*v505.ToRank` 
    
</dd>
</dl>

<dl>
<dd>

**isHeadline:** `*v505.IsHeadline` 
    
</dd>
</dl>

<dl>
<dd>

**isOpinion:** `*v505.IsOpinion` 
    
</dd>
</dl>

<dl>
<dd>

**isPaidContent:** `*v505.IsPaidContent` 
    
</dd>
</dl>

<dl>
<dd>

**parentUrl:** `*v505.ParentUrl` 
    
</dd>
</dl>

<dl>
<dd>

**allLinks:** `*v505.AllLinks` 
    
</dd>
</dl>

<dl>
<dd>

**allDomainLinks:** `*v505.AllDomainLinks` 
    
</dd>
</dl>

<dl>
<dd>

**additionalDomainInfo:** `*v505.AdditionalDomainInfo` 
    
</dd>
</dl>

<dl>
<dd>

**isNewsDomain:** `*v505.IsNewsDomain` 
    
</dd>
</dl>

<dl>
<dd>

**newsDomainType:** `*v505.NewsDomainType` 
    
</dd>
</dl>

<dl>
<dd>

**newsType:** `*v505.NewsType` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMin:** `*v505.WordCountMin` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMax:** `*v505.WordCountMax` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*v505.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*v505.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringEnabled:** `*v505.ClusteringEnabled` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringVariable:** `*v505.ClusteringVariable` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringThreshold:** `*v505.ClusteringThreshold` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*v505.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*v505.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*v505.Theme` 
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*v505.NotTheme` 
    
</dd>
</dl>

<dl>
<dd>

**orgEntityName:** `*v505.OrgEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**perEntityName:** `*v505.PerEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**locEntityName:** `*v505.LocEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**miscEntityName:** `*v505.MiscEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*v505.TitleSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*v505.TitleSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*v505.ContentSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*v505.ContentSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**iptcTags:** `*v505.IptcTags` 
    
</dd>
</dl>

<dl>
<dd>

**notIptcTags:** `*v505.NotIptcTags` 
    
</dd>
</dl>

<dl>
<dd>

**iabTags:** `*v505.IabTags` 
    
</dd>
</dl>

<dl>
<dd>

**notIabTags:** `*v505.NotIabTags` 
    
</dd>
</dl>

<dl>
<dd>

**customTags:** `*v505.CustomTags` 
    
</dd>
</dl>

<dl>
<dd>

**excludeDuplicates:** `*v505.ExcludeDuplicates` 
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*v505.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## LatestHeadlines
<details><summary><code>client.Latestheadlines.Get() -> *v505.LatestHeadlinesGetResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the latest headlines for the specified time period. You can filter results by language, country, source, and more.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.LatestHeadlinesGetRequest{
        When: v505.String(
            "7d",
        ),
        Lang: v505.String(
            "en",
        ),
        NotLang: v505.String(
            "fr",
        ),
        Countries: v505.String(
            "US",
        ),
        NotCountries: v505.String(
            "UK",
        ),
        PredefinedSources: v505.String(
            "top 100 US, top 5 GB",
        ),
        Sources: v505.String(
            "nytimes.com",
        ),
        NotSources: v505.String(
            "cnn.com",
        ),
        NotAuthorName: v505.String(
            "John Doe",
        ),
        ParentUrl: v505.String(
            "https://www.washingtonpost.com/politics",
        ),
        AllLinks: v505.String(
            "https://aiindex.stanford.edu/report",
        ),
        AllDomainLinks: v505.String(
            "nvidia.com",
        ),
        IncludeTranslationFields: v505.Bool(
            true,
        ),
        IncludeNlpData: v505.Bool(
            true,
        ),
        HasNlp: v505.Bool(
            true,
        ),
        Theme: v505.String(
            "Business,Finance",
        ),
        NotTheme: v505.String(
            "Crime",
        ),
        OrgEntityName: v505.String(
            "Apple",
        ),
        PerEntityName: v505.String(
            "Elon Musk",
        ),
        LocEntityName: v505.String(
            "California",
        ),
        MiscEntityName: v505.String(
            "Bitcoin",
        ),
        IptcTags: v505.String(
            "20000199,20000209",
        ),
        NotIptcTags: v505.String(
            "20000205,20000209",
        ),
        IabTags: v505.String(
            "Business,Events",
        ),
        NotIabTags: v505.String(
            "Agriculture,Metals",
        ),
        CustomTags: v505.String(
            "Tag1,Tag2,Tag3",
        ),
    }
client.Latestheadlines.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**when:** `*string` 

The time period for which you want to get the latest headlines.

Format examples:
- `7d`: Last seven days
- `30d`: Last 30 days
- `1h`: Last hour
- `24h`: Last 24 hours
    
</dd>
</dl>

<dl>
<dd>

**byParseDate:** `*bool` — If true, the `from_` and `to_` parameters use article parse dates instead of published dates. Additionally, the `parse_date` variable is added to the output for each article object.
    
</dd>
</dl>

<dl>
<dd>

**lang:** `*string` 

The language(s) of the search. The only accepted format is the two-letter [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code. To select multiple languages, use a comma-separated string. 

Example: `"en, es"`

To learn more, see [Enumerated parameters > Language](/docs/v3/api-reference/overview/enumerated-parameters#language-lang-and-not-lang).
    
</dd>
</dl>

<dl>
<dd>

**notLang:** `*string` 

The language(s) to exclude from the search. The accepted format is the two-letter [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code. To exclude multiple languages, use a comma-separated string. 

Example: `"fr, de"`

To learn more, see [Enumerated parameters > Language](/docs/v3/api-reference/overview/enumerated-parameters#language-lang-and-not-lang).
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*string` 

The countries where the news publisher is located. The accepted format is the two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. To select multiple countries, use a comma-separated string.

Example: `"US, CA"`

To learn more, see [Enumerated parameters > Country](/docs/v3/api-reference/overview/enumerated-parameters#country-country-and-not-country).
    
</dd>
</dl>

<dl>
<dd>

**notCountries:** `*string` 

The publisher location countries to exclude from the search. The accepted format is the two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. To exclude multiple countries, use a comma-separated string. 

Example:`"US, CA"`

To learn more, see [Enumerated parameters > Country](/docs/v3/api-reference/overview/enumerated-parameters#country-country-and-not-country).
    
</dd>
</dl>

<dl>
<dd>

**predefinedSources:** `*string` 

Predefined top news sources per country. 

Format: start with the word `top`, followed by the number of desired sources, and then the two-letter country code [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). Multiple countries with the number of top sources can be specified as a comma-separated string.

Examples: 
- `"top 100 US"`
- `"top 33 AT"`
- `"top 50 US, top 20 GB"`
- `"top 33 AT, top 50 IT"`
    
</dd>
</dl>

<dl>
<dd>

**sources:** `*string` 

One or more news sources to narrow down the search. The format must be a domain URL. Subdomains, such as `finance.yahoo.com`, are also acceptable.To specify multiple sources, use a comma-separated string.

Examples:
- `"nytimes.com"`
- `"theguardian.com, finance.yahoo.com"`
    
</dd>
</dl>

<dl>
<dd>

**notSources:** `*string` 

The news sources to exclude from the search. To exclude multiple sources, use a comma-separated string. 

Example: `"cnn.com, wsj.com"`
    
</dd>
</dl>

<dl>
<dd>

**notAuthorName:** `*string` 

The list of author names to exclude from your search. To exclude articles by specific authors, use a comma-separated string.

Example: `"John Doe, Jane Doe"`
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*bool` — If true, limits the search to sources ranked in the top 1 million online websites. If false, includes unranked sources which are assigned a rank of 999999.
    
</dd>
</dl>

<dl>
<dd>

**isHeadline:** `*bool` — If true, only returns articles that were posted on the home page of a given news domain.
    
</dd>
</dl>

<dl>
<dd>

**isOpinion:** `*bool` — If true, returns only opinion pieces. If false, excludes opinion-based articles and returns news only.
    
</dd>
</dl>

<dl>
<dd>

**isPaidContent:** `*bool` — If false, returns only articles that have publicly available complete content. Some publishers partially block content, so this setting ensures that only full articles are retrieved.
    
</dd>
</dl>

<dl>
<dd>

**parentUrl:** `*string` 

The categorical URL(s) to filter your search. To filter your search by multiple categorical URLs, use a comma-separated string.

Example: `"wsj.com/politics, wsj.com/tech"`
    
</dd>
</dl>

<dl>
<dd>

**allLinks:** `*string` 

The complete URL(s) mentioned in the article. For multiple URLs, use a comma-separated string.

Example: `"https://aiindex.stanford.edu/report, https://www.stateof.ai"`

For more details, see [Search by URL](/docs/v3/documentation/how-to/search-by-url).
    
</dd>
</dl>

<dl>
<dd>

**allDomainLinks:** `*string` 

The domain(s) mentioned in the article. For multiple domains, use a comma-separated string.

Example: `"who.int, nih.gov"`

For more details, see [Search by URL](/docs/v3/documentation/how-to/search-by-url).
    
</dd>
</dl>

<dl>
<dd>

**wordCountMin:** `*int` — The minimum number of words an article must contain. To be used for avoiding articles with small content.
    
</dd>
</dl>

<dl>
<dd>

**wordCountMax:** `*int` — The maximum number of words an article can contain. To be used for avoiding articles with large content.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 

The page number to scroll through the results. Use for pagination, as a single API response can return up to 1,000 articles. 

For details, see [How to paginate large datasets](https://www.newscatcherapi.com/docs/v3/documentation/how-to/paginate-large-datasets).
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The number of articles to return per page.
    
</dd>
</dl>

<dl>
<dd>

**clusteringEnabled:** `*bool` 

Determines whether to group similar articles into clusters. If true, the API returns clustered results.

To learn more, see [Clustering news articles](/docs/v3/documentation/guides-and-concepts/clustering-news-articles).
    
</dd>
</dl>

<dl>
<dd>

**clusteringVariable:** `*v505.LatestHeadlinesGetRequestClusteringVariable` 

Specifies which part of the article to use for determining similarity when clustering.

Possible values are:
- `content`: Uses the full article content (default).
- `title`: Uses only the article title.
- `summary`: Uses the article summary.

To learn more, see [Clustering news articles](/docs/v3/documentation/guides-and-concepts/clustering-news-articles).
    
</dd>
</dl>

<dl>
<dd>

**clusteringThreshold:** `*float64` 

Sets the similarity threshold for grouping articles into clusters. A lower value creates more inclusive clusters, while a higher value requires greater similarity between articles.

Examples:
- `0.3`: Results in larger, more diverse clusters.
- `0.6`: Balances cluster size and article similarity (default).
- `0.9`: Creates smaller, tightly related clusters.

To learn more, see [Clustering news articles](/docs/v3/documentation/guides-and-concepts/clustering-news-articles).
    
</dd>
</dl>

<dl>
<dd>

**includeTranslationFields:** `*v505.IncludeTranslationFields` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*v505.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*v505.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*string` 

Filters articles based on their general topic, as determined by NLP analysis. To select multiple themes, use a comma-separated string.

Example: `"Finance, Tech"`

**Note**: The `theme` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).

Available options: `Business`, `Economics`, `Entertainment`, `Finance`, `Health`, `Politics`, `Science`, `Sports`, `Tech`, `Crime`, `Financial Crime`, `Lifestyle`, `Automotive`, `Travel`, `Weather`, `General`.
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*string` 

Inverse of the `theme` parameter. Excludes articles based on their general topic, as determined by NLP analysis. To exclude multiple themes, use a comma-separated string. 

Example: `"Crime, Tech"`

**Note**: The `not_theme` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**orgEntityName:** `*string` 

Filters articles that mention specific organization names, as identified by NLP analysis. To specify multiple organizations, use a comma-separated string. To search named entities in translations, combine with the translation options of the `search_in` parameter (e.g., `title_content_translated`).

Example: `"Apple, Microsoft"`

**Note**: The `ORG_entity_name` parameter is only available if NLP is included in your subscription plan.

To learn more, see [Search by entity](/docs/v3/documentation/how-to/search-by-entity).
    
</dd>
</dl>

<dl>
<dd>

**perEntityName:** `*string` 

Filters articles that mention specific person names, as identified by NLP analysis. To specify multiple names, use a comma-separated string. To search named entities in translations, combine with the translation options of the `search_in` parameter (e.g., `title_content_translated`).

Example: `"Elon Musk, Jeff Bezos"`

**Note**: The `PER_entity_name` parameter is only available if NLP is included in your subscription plan.

To learn more, see [Search by entity](/docs/v3/documentation/how-to/search-by-entity).
    
</dd>
</dl>

<dl>
<dd>

**locEntityName:** `*string` 

Filters articles that mention specific location names, as identified by NLP analysis. To specify multiple locations, use a comma-separated string. To search named entities in translations, combine with the translation options of the `search_in` parameter (e.g., `title_content_translated`).

Example: `"California, New York"`

**Note**: The `LOC_entity_name` parameter is only available if NLP is included in your subscription plan.

To learn more, see [Search by entity](/docs/v3/documentation/how-to/search-by-entity).
    
</dd>
</dl>

<dl>
<dd>

**miscEntityName:** `*string` 

Filters articles that mention other named entities not falling under person, organization, or location categories. Includes events, nationalities, products, works of art, and more. To specify multiple entities, use a comma-separated string. To search named entities in translations, combine with the translation options of the `search_in` parameter (e.g., `title_content_translated`).

Example: `"Bitcoin, Blockchain"`

**Note**: The `MISC_entity_name` parameter is only available if NLP is included in your subscription plan.

To learn more, see [Search by entity](/docs/v3/documentation/how-to/search-by-entity).
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*float64` 

Filters articles based on the minimum sentiment score of their titles.

Range is `-1.0` to `1.0`, where:
- Negative values indicate negative sentiment.
- Positive values indicate positive sentiment.
- Values close to 0 indicate neutral sentiment.

**Note**: The `title_sentiment_min` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*float64` 

Filters articles based on the maximum sentiment score of their titles.

Range is `-1.0` to `1.0`, where:
- Negative values indicate negative sentiment.
- Positive values indicate positive sentiment.
- Values close to 0 indicate neutral sentiment.

**Note**: The `title_sentiment_max` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*float64` 

Filters articles based on the minimum sentiment score of their content.

Range is `-1.0` to `1.0`, where:
- Negative values indicate negative sentiment.
- Positive values indicate positive sentiment.
- Values close to 0 indicate neutral sentiment.

**Note**: The `content_sentiment_min` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*float64` 

Filters articles based on the maximum sentiment score of their content.

Range is `-1.0` to `1.0`, where:
- Negative values indicate negative sentiment.
- Positive values indicate positive sentiment.
- Values close to 0 indicate neutral sentiment.

**Note**: The `content_sentiment_max` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**iptcTags:** `*string` 

Filters articles based on International Press Telecommunications Council (IPTC) media topic tags. To specify multiple IPTC tags, use a comma-separated string of tag IDs. 

Example: `"20000199, 20000209"`

**Note**: The `iptc_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see [IPTC Media Topic NewsCodes](https://www.iptc.org/std/NewsCodes/treeview/mediatopic/mediatopic-en-GB.html).
    
</dd>
</dl>

<dl>
<dd>

**notIptcTags:** `*string` 

Inverse of the `iptc_tags` parameter. Excludes articles based on International Press Telecommunications Council (IPTC) media topic tags. To specify multiple IPTC tags to exclude, use a comma-separated string of tag IDs. 

Example: `"20000205, 20000209"`

**Note**: The `not_iptc_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see [IPTC Media Topic NewsCodes](https://www.iptc.org/std/NewsCodes/treeview/mediatopic/mediatopic-en-GB.html).
    
</dd>
</dl>

<dl>
<dd>

**iabTags:** `*string` 

Filters articles based on Interactive Advertising Bureau (IAB) content categories. These tags provide a standardized taxonomy for digital advertising content categorization. To specify multiple IAB categories, use a comma-separated string. 

Example: `"Business, Events"`

**Note**: The `iab_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see the [IAB Content taxonomy](https://iabtechlab.com/standards/content-taxonomy/).
    
</dd>
</dl>

<dl>
<dd>

**notIabTags:** `*string` 

Inverse of the `iab_tags` parameter. Excludes articles based on Interactive Advertising Bureau (IAB) content categories. These tags provide a standardized taxonomy for digital advertising content categorization. To specify multiple IAB categories to exclude, use a comma-separated string. 

Example: `"Agriculture, Metals"`

**Note**: The `not_iab_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see the [IAB Content taxonomy](https://iabtechlab.com/standards/content-taxonomy/).
    
</dd>
</dl>

<dl>
<dd>

**customTags:** `*string` 

Filters articles based on provided taxonomy that is tailored to your specific needs and is accessible only with your API key. To specify tags, use the following pattern: 

- `custom_tags.taxonomy=Tag1,Tag2,Tag3`, where `taxonomy` is the taxonomy name and `Tag1,Tag2,Tag3` is a comma-separated list of tags.

Example: `custom_tags.industry="Manufacturing, Supply Chain, Logistics"`

To learn more, see the [Custom tags](/docs/v3/documentation/guides-and-concepts/custom-tags).
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*bool` — If true, returns only articles/sources that comply with the publisher's robots.txt rules. If false, returns only articles/sources that do not comply with robots.txt rules. If omitted, returns all articles/sources regardless of compliance status.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Latestheadlines.Post(request) -> *v505.LatestHeadlinesPostResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the latest headlines for the specified time period. You can filter results by language, country, source, and more.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.LatestHeadlinesPostRequest{
        When: v505.String(
            "7d",
        ),
        PageSize: v505.Int(
            1,
        ),
    }
client.Latestheadlines.Post(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**when:** `*v505.When` 
    
</dd>
</dl>

<dl>
<dd>

**byParseDate:** `*v505.ByParseDate` 
    
</dd>
</dl>

<dl>
<dd>

**lang:** `*v505.Lang` 
    
</dd>
</dl>

<dl>
<dd>

**notLang:** `*v505.NotLang` 
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*v505.Countries` 
    
</dd>
</dl>

<dl>
<dd>

**notCountries:** `*v505.NotCountries` 
    
</dd>
</dl>

<dl>
<dd>

**predefinedSources:** `*v505.PredefinedSources` 
    
</dd>
</dl>

<dl>
<dd>

**sources:** `*v505.Sources` 
    
</dd>
</dl>

<dl>
<dd>

**notSources:** `*v505.NotSources` 
    
</dd>
</dl>

<dl>
<dd>

**notAuthorName:** `*v505.NotAuthorName` 
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*v505.RankedOnly` 
    
</dd>
</dl>

<dl>
<dd>

**isHeadline:** `*v505.IsHeadline` 
    
</dd>
</dl>

<dl>
<dd>

**isOpinion:** `*v505.IsOpinion` 
    
</dd>
</dl>

<dl>
<dd>

**isPaidContent:** `*v505.IsPaidContent` 
    
</dd>
</dl>

<dl>
<dd>

**parentUrl:** `*v505.ParentUrl` 
    
</dd>
</dl>

<dl>
<dd>

**allLinks:** `*v505.AllLinks` 
    
</dd>
</dl>

<dl>
<dd>

**allDomainLinks:** `*v505.AllDomainLinks` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMin:** `*v505.WordCountMin` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMax:** `*v505.WordCountMax` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*v505.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*v505.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringEnabled:** `*v505.ClusteringEnabled` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringVariable:** `*v505.ClusteringVariable` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringThreshold:** `*v505.ClusteringThreshold` 
    
</dd>
</dl>

<dl>
<dd>

**includeTranslationFields:** `*v505.IncludeTranslationFields` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*v505.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*v505.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*v505.Theme` 
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*v505.NotTheme` 
    
</dd>
</dl>

<dl>
<dd>

**orgEntityName:** `*v505.OrgEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**perEntityName:** `*v505.PerEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**locEntityName:** `*v505.LocEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**miscEntityName:** `*v505.MiscEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*v505.TitleSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*v505.TitleSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*v505.ContentSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*v505.ContentSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**iptcTags:** `*v505.IptcTags` 
    
</dd>
</dl>

<dl>
<dd>

**notIptcTags:** `*v505.NotIptcTags` 
    
</dd>
</dl>

<dl>
<dd>

**iabTags:** `*v505.IabTags` 
    
</dd>
</dl>

<dl>
<dd>

**notIabTags:** `*v505.NotIabTags` 
    
</dd>
</dl>

<dl>
<dd>

**customTags:** `*v505.CustomTags` 
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*v505.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Breaking News
<details><summary><code>client.BreakingNews.BreakingNewsGet() -> *v505.BreakingNewsResponseDto</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves breaking news articles and sorts them based on specified criteria.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.BreakingNewsGetRequest{
        TopNArticles: v505.Int(
            5,
        ),
        IncludeTranslationFields: v505.Bool(
            true,
        ),
        IncludeNlpData: v505.Bool(
            true,
        ),
        HasNlp: v505.Bool(
            true,
        ),
        Theme: v505.String(
            "Business,Finance",
        ),
        NotTheme: v505.String(
            "Crime",
        ),
        OrgEntityName: v505.String(
            "Apple",
        ),
        PerEntityName: v505.String(
            "Elon Musk",
        ),
        LocEntityName: v505.String(
            "California",
        ),
        MiscEntityName: v505.String(
            "Bitcoin",
        ),
    }
client.BreakingNews.BreakingNewsGet(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sortBy:** `*v505.BreakingNewsGetRequestSortBy` 

The sorting order of the results. Possible values are:
- `relevancy`: The most relevant results first.
- `date`: The most recently published results first.
- `rank`: The results from the highest-ranked sources first.
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*bool` — If true, limits the search to sources ranked in the top 1 million online websites. If false, includes unranked sources which are assigned a rank of 999999.
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*int` — The lowest boundary of the rank of a news website to filter by. A lower rank indicates a more popular source.
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*int` — The highest boundary of the rank of a news website to filter by. A lower rank indicates a more popular source.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 

The page number to scroll through the results. Use for pagination, as a single API response can return up to 1,000 articles. 

For details, see [How to paginate large datasets](https://www.newscatcherapi.com/docs/v3/documentation/how-to/paginate-large-datasets).
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The number of articles to return per page.
    
</dd>
</dl>

<dl>
<dd>

**topNArticles:** `*v505.TopNArticles` 
    
</dd>
</dl>

<dl>
<dd>

**includeTranslationFields:** `*v505.IncludeTranslationFields` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*v505.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*v505.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*string` 

Filters articles based on their general topic, as determined by NLP analysis. To select multiple themes, use a comma-separated string.

Example: `"Finance, Tech"`

**Note**: The `theme` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).

Available options: `Business`, `Economics`, `Entertainment`, `Finance`, `Health`, `Politics`, `Science`, `Sports`, `Tech`, `Crime`, `Financial Crime`, `Lifestyle`, `Automotive`, `Travel`, `Weather`, `General`.
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*string` 

Inverse of the `theme` parameter. Excludes articles based on their general topic, as determined by NLP analysis. To exclude multiple themes, use a comma-separated string. 

Example: `"Crime, Tech"`

**Note**: The `not_theme` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**orgEntityName:** `*string` 

Filters articles that mention specific organization names, as identified by NLP analysis. To specify multiple organizations, use a comma-separated string. To search named entities in translations, combine with the translation options of the `search_in` parameter (e.g., `title_content_translated`).

Example: `"Apple, Microsoft"`

**Note**: The `ORG_entity_name` parameter is only available if NLP is included in your subscription plan.

To learn more, see [Search by entity](/docs/v3/documentation/how-to/search-by-entity).
    
</dd>
</dl>

<dl>
<dd>

**perEntityName:** `*string` 

Filters articles that mention specific person names, as identified by NLP analysis. To specify multiple names, use a comma-separated string. To search named entities in translations, combine with the translation options of the `search_in` parameter (e.g., `title_content_translated`).

Example: `"Elon Musk, Jeff Bezos"`

**Note**: The `PER_entity_name` parameter is only available if NLP is included in your subscription plan.

To learn more, see [Search by entity](/docs/v3/documentation/how-to/search-by-entity).
    
</dd>
</dl>

<dl>
<dd>

**locEntityName:** `*string` 

Filters articles that mention specific location names, as identified by NLP analysis. To specify multiple locations, use a comma-separated string. To search named entities in translations, combine with the translation options of the `search_in` parameter (e.g., `title_content_translated`).

Example: `"California, New York"`

**Note**: The `LOC_entity_name` parameter is only available if NLP is included in your subscription plan.

To learn more, see [Search by entity](/docs/v3/documentation/how-to/search-by-entity).
    
</dd>
</dl>

<dl>
<dd>

**miscEntityName:** `*string` 

Filters articles that mention other named entities not falling under person, organization, or location categories. Includes events, nationalities, products, works of art, and more. To specify multiple entities, use a comma-separated string. To search named entities in translations, combine with the translation options of the `search_in` parameter (e.g., `title_content_translated`).

Example: `"Bitcoin, Blockchain"`

**Note**: The `MISC_entity_name` parameter is only available if NLP is included in your subscription plan.

To learn more, see [Search by entity](/docs/v3/documentation/how-to/search-by-entity).
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*float64` 

Filters articles based on the minimum sentiment score of their titles.

Range is `-1.0` to `1.0`, where:
- Negative values indicate negative sentiment.
- Positive values indicate positive sentiment.
- Values close to 0 indicate neutral sentiment.

**Note**: The `title_sentiment_min` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*float64` 

Filters articles based on the maximum sentiment score of their titles.

Range is `-1.0` to `1.0`, where:
- Negative values indicate negative sentiment.
- Positive values indicate positive sentiment.
- Values close to 0 indicate neutral sentiment.

**Note**: The `title_sentiment_max` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*float64` 

Filters articles based on the minimum sentiment score of their content.

Range is `-1.0` to `1.0`, where:
- Negative values indicate negative sentiment.
- Positive values indicate positive sentiment.
- Values close to 0 indicate neutral sentiment.

**Note**: The `content_sentiment_min` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*float64` 

Filters articles based on the maximum sentiment score of their content.

Range is `-1.0` to `1.0`, where:
- Negative values indicate negative sentiment.
- Positive values indicate positive sentiment.
- Values close to 0 indicate neutral sentiment.

**Note**: The `content_sentiment_max` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*bool` — If true, returns only articles/sources that comply with the publisher's robots.txt rules. If false, returns only articles/sources that do not comply with robots.txt rules. If omitted, returns all articles/sources regardless of compliance status.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BreakingNews.BreakingNewsPost(request) -> *v505.BreakingNewsResponseDto</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves breaking news articles and sorts them based on specified criteria.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.BreakingNewsPostRequest{
        SortBy: v505.SortByRelevancy.Ptr(),
        RankedOnly: v505.Bool(
            true,
        ),
        TopNArticles: v505.Int(
            1,
        ),
    }
client.BreakingNews.BreakingNewsPost(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sortBy:** `*v505.SortBy` 
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*v505.RankedOnly` 
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*v505.FromRank` 
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*v505.ToRank` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*v505.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*v505.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**topNArticles:** `*v505.TopNArticles` 
    
</dd>
</dl>

<dl>
<dd>

**includeTranslationFields:** `*v505.IncludeTranslationFields` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*v505.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*v505.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*v505.Theme` 
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*v505.NotTheme` 
    
</dd>
</dl>

<dl>
<dd>

**orgEntityName:** `*v505.OrgEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**perEntityName:** `*v505.PerEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**locEntityName:** `*v505.LocEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**miscEntityName:** `*v505.MiscEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*v505.TitleSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*v505.TitleSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*v505.ContentSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentientMax:** `*v505.ContentSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*v505.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Authors
<details><summary><code>client.Authors.Get() -> *v505.AuthorsGetResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Searches for articles written by a specified author. You can filter results by language, country, source, and more.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.AuthorsGetRequest{
        AuthorName: "Jane Smith",
        NotAuthorName: v505.String(
            "John Doe",
        ),
        PredefinedSources: v505.String(
            "top 100 US, top 5 GB",
        ),
        Sources: v505.String(
            "nytimes.com",
        ),
        NotSources: v505.String(
            "cnn.com",
        ),
        Lang: v505.String(
            "en",
        ),
        NotLang: v505.String(
            "fr",
        ),
        Countries: v505.String(
            "US",
        ),
        NotCountries: v505.String(
            "UK",
        ),
        From: &v505.AuthorsGetRequestFrom{
            DateTime: v505.MustParseDateTime(
                "2024-07-01T00:00:00Z",
            ),
        },
        To: &v505.AuthorsGetRequestTo{
            DateTime: v505.MustParseDateTime(
                "2024-07-01T00:00:00Z",
            ),
        },
        ParentUrl: v505.String(
            "https://www.washingtonpost.com/politics",
        ),
        AllLinks: v505.String(
            "https://aiindex.stanford.edu/report",
        ),
        AllDomainLinks: v505.String(
            "nvidia.com",
        ),
        IncludeTranslationFields: v505.Bool(
            true,
        ),
        IncludeNlpData: v505.Bool(
            true,
        ),
        HasNlp: v505.Bool(
            true,
        ),
        Theme: v505.String(
            "Business,Finance",
        ),
        NotTheme: v505.String(
            "Crime",
        ),
        NerName: v505.String(
            "Tesla",
        ),
        IptcTags: v505.String(
            "20000199,20000209",
        ),
        NotIptcTags: v505.String(
            "20000205,20000209",
        ),
        IabTags: v505.String(
            "Business,Events",
        ),
        NotIabTags: v505.String(
            "Agriculture,Metals",
        ),
        CustomTags: v505.String(
            "Tag1,Tag2,Tag3",
        ),
    }
client.Authors.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authorName:** `string` — The name of the author to search for. This parameter returns exact matches only.
    
</dd>
</dl>

<dl>
<dd>

**notAuthorName:** `*string` 

The list of author names to exclude from your search. To exclude articles by specific authors, use a comma-separated string.

Example: `"John Doe, Jane Doe"`
    
</dd>
</dl>

<dl>
<dd>

**predefinedSources:** `*string` 

Predefined top news sources per country. 

Format: start with the word `top`, followed by the number of desired sources, and then the two-letter country code [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). Multiple countries with the number of top sources can be specified as a comma-separated string.

Examples: 
- `"top 100 US"`
- `"top 33 AT"`
- `"top 50 US, top 20 GB"`
- `"top 33 AT, top 50 IT"`
    
</dd>
</dl>

<dl>
<dd>

**sources:** `*string` 

One or more news sources to narrow down the search. The format must be a domain URL. Subdomains, such as `finance.yahoo.com`, are also acceptable.To specify multiple sources, use a comma-separated string.

Examples:
- `"nytimes.com"`
- `"theguardian.com, finance.yahoo.com"`
    
</dd>
</dl>

<dl>
<dd>

**notSources:** `*string` 

The news sources to exclude from the search. To exclude multiple sources, use a comma-separated string. 

Example: `"cnn.com, wsj.com"`
    
</dd>
</dl>

<dl>
<dd>

**lang:** `*string` 

The language(s) of the search. The only accepted format is the two-letter [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code. To select multiple languages, use a comma-separated string. 

Example: `"en, es"`

To learn more, see [Enumerated parameters > Language](/docs/v3/api-reference/overview/enumerated-parameters#language-lang-and-not-lang).
    
</dd>
</dl>

<dl>
<dd>

**notLang:** `*string` 

The language(s) to exclude from the search. The accepted format is the two-letter [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code. To exclude multiple languages, use a comma-separated string. 

Example: `"fr, de"`

To learn more, see [Enumerated parameters > Language](/docs/v3/api-reference/overview/enumerated-parameters#language-lang-and-not-lang).
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*string` 

The countries where the news publisher is located. The accepted format is the two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. To select multiple countries, use a comma-separated string.

Example: `"US, CA"`

To learn more, see [Enumerated parameters > Country](/docs/v3/api-reference/overview/enumerated-parameters#country-country-and-not-country).
    
</dd>
</dl>

<dl>
<dd>

**notCountries:** `*string` 

The publisher location countries to exclude from the search. The accepted format is the two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. To exclude multiple countries, use a comma-separated string. 

Example:`"US, CA"`

To learn more, see [Enumerated parameters > Country](/docs/v3/api-reference/overview/enumerated-parameters#country-country-and-not-country).
    
</dd>
</dl>

<dl>
<dd>

**from:** `*v505.AuthorsGetRequestFrom` 

The starting point in time to search from. Accepts date-time strings in ISO 8601 format and plain text. The default time zone is UTC. 

Formats with examples:
- YYYY-mm-ddTHH:MM:SS: `2024-07-01T00:00:00`
- YYYY-MM-dd: `2024-07-01`
- YYYY/mm/dd HH:MM:SS: `2024/07/01 00:00:00`
- YYYY/mm/dd: `2024/07/01`
- English phrases: `7 day ago`, `today`

**Note**: By default, applied to the publication date of the article. To use the article's parse date instead, set the `by_parse_date` parameter to `true`.
    
</dd>
</dl>

<dl>
<dd>

**to:** `*v505.AuthorsGetRequestTo` 

The ending point in time to search up to. Accepts date-time strings in ISO 8601 format and plain text. The default time zone is UTC. 

Formats with examples:
- YYYY-mm-ddTHH:MM:SS: `2024-07-01T00:00:00`
- YYYY-MM-dd: `2024-07-01`
- YYYY/mm/dd HH:MM:SS: `2024/07/01 00:00:00`
- YYYY/mm/dd: `2024/07/01`
- English phrases: `1 day ago`, `now`

**Note**: By default, applied to the publication date of the article. To use the article's parse date instead, set the `by_parse_date` parameter to `true`.
    
</dd>
</dl>

<dl>
<dd>

**publishedDatePrecision:** `*v505.AuthorsGetRequestPublishedDatePrecision` 

The precision of the published date. There are three types:
- `full`: The day and time of an article is correctly identified with the appropriate timezone.
- `timezone unknown`: The day and time of an article is correctly identified without timezone.
- `date`: Only the day is identified without an exact time.
    
</dd>
</dl>

<dl>
<dd>

**byParseDate:** `*bool` — If true, the `from_` and `to_` parameters use article parse dates instead of published dates. Additionally, the `parse_date` variable is added to the output for each article object.
    
</dd>
</dl>

<dl>
<dd>

**sortBy:** `*v505.AuthorsGetRequestSortBy` 

The sorting order of the results. Possible values are:
- `relevancy`: The most relevant results first.
- `date`: The most recently published results first.
- `rank`: The results from the highest-ranked sources first.
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*bool` — If true, limits the search to sources ranked in the top 1 million online websites. If false, includes unranked sources which are assigned a rank of 999999.
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*int` — The lowest boundary of the rank of a news website to filter by. A lower rank indicates a more popular source.
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*int` — The highest boundary of the rank of a news website to filter by. A lower rank indicates a more popular source.
    
</dd>
</dl>

<dl>
<dd>

**isHeadline:** `*bool` — If true, only returns articles that were posted on the home page of a given news domain.
    
</dd>
</dl>

<dl>
<dd>

**isOpinion:** `*bool` — If true, returns only opinion pieces. If false, excludes opinion-based articles and returns news only.
    
</dd>
</dl>

<dl>
<dd>

**isPaidContent:** `*bool` — If false, returns only articles that have publicly available complete content. Some publishers partially block content, so this setting ensures that only full articles are retrieved.
    
</dd>
</dl>

<dl>
<dd>

**parentUrl:** `*string` 

The categorical URL(s) to filter your search. To filter your search by multiple categorical URLs, use a comma-separated string.

Example: `"wsj.com/politics, wsj.com/tech"`
    
</dd>
</dl>

<dl>
<dd>

**allLinks:** `*string` 

The complete URL(s) mentioned in the article. For multiple URLs, use a comma-separated string.

Example: `"https://aiindex.stanford.edu/report, https://www.stateof.ai"`

For more details, see [Search by URL](/docs/v3/documentation/how-to/search-by-url).
    
</dd>
</dl>

<dl>
<dd>

**allDomainLinks:** `*string` 

The domain(s) mentioned in the article. For multiple domains, use a comma-separated string.

Example: `"who.int, nih.gov"`

For more details, see [Search by URL](/docs/v3/documentation/how-to/search-by-url).
    
</dd>
</dl>

<dl>
<dd>

**wordCountMin:** `*int` — The minimum number of words an article must contain. To be used for avoiding articles with small content.
    
</dd>
</dl>

<dl>
<dd>

**wordCountMax:** `*int` — The maximum number of words an article can contain. To be used for avoiding articles with large content.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 

The page number to scroll through the results. Use for pagination, as a single API response can return up to 1,000 articles. 

For details, see [How to paginate large datasets](https://www.newscatcherapi.com/docs/v3/documentation/how-to/paginate-large-datasets).
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The number of articles to return per page.
    
</dd>
</dl>

<dl>
<dd>

**includeTranslationFields:** `*v505.IncludeTranslationFields` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*v505.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*v505.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*string` 

Filters articles based on their general topic, as determined by NLP analysis. To select multiple themes, use a comma-separated string.

Example: `"Finance, Tech"`

**Note**: The `theme` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).

Available options: `Business`, `Economics`, `Entertainment`, `Finance`, `Health`, `Politics`, `Science`, `Sports`, `Tech`, `Crime`, `Financial Crime`, `Lifestyle`, `Automotive`, `Travel`, `Weather`, `General`.
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*string` 

Inverse of the `theme` parameter. Excludes articles based on their general topic, as determined by NLP analysis. To exclude multiple themes, use a comma-separated string. 

Example: `"Crime, Tech"`

**Note**: The `not_theme` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**nerName:** `*string` 

The name of person, organization, location, product or other named entity to search for. To specify multiple names use a comma-separated string. 

Example: `"Tesla, Amazon"`
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*float64` 

Filters articles based on the minimum sentiment score of their titles.

Range is `-1.0` to `1.0`, where:
- Negative values indicate negative sentiment.
- Positive values indicate positive sentiment.
- Values close to 0 indicate neutral sentiment.

**Note**: The `title_sentiment_min` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*float64` 

Filters articles based on the maximum sentiment score of their titles.

Range is `-1.0` to `1.0`, where:
- Negative values indicate negative sentiment.
- Positive values indicate positive sentiment.
- Values close to 0 indicate neutral sentiment.

**Note**: The `title_sentiment_max` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*float64` 

Filters articles based on the minimum sentiment score of their content.

Range is `-1.0` to `1.0`, where:
- Negative values indicate negative sentiment.
- Positive values indicate positive sentiment.
- Values close to 0 indicate neutral sentiment.

**Note**: The `content_sentiment_min` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*float64` 

Filters articles based on the maximum sentiment score of their content.

Range is `-1.0` to `1.0`, where:
- Negative values indicate negative sentiment.
- Positive values indicate positive sentiment.
- Values close to 0 indicate neutral sentiment.

**Note**: The `content_sentiment_max` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**iptcTags:** `*string` 

Filters articles based on International Press Telecommunications Council (IPTC) media topic tags. To specify multiple IPTC tags, use a comma-separated string of tag IDs. 

Example: `"20000199, 20000209"`

**Note**: The `iptc_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see [IPTC Media Topic NewsCodes](https://www.iptc.org/std/NewsCodes/treeview/mediatopic/mediatopic-en-GB.html).
    
</dd>
</dl>

<dl>
<dd>

**notIptcTags:** `*string` 

Inverse of the `iptc_tags` parameter. Excludes articles based on International Press Telecommunications Council (IPTC) media topic tags. To specify multiple IPTC tags to exclude, use a comma-separated string of tag IDs. 

Example: `"20000205, 20000209"`

**Note**: The `not_iptc_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see [IPTC Media Topic NewsCodes](https://www.iptc.org/std/NewsCodes/treeview/mediatopic/mediatopic-en-GB.html).
    
</dd>
</dl>

<dl>
<dd>

**iabTags:** `*string` 

Filters articles based on Interactive Advertising Bureau (IAB) content categories. These tags provide a standardized taxonomy for digital advertising content categorization. To specify multiple IAB categories, use a comma-separated string. 

Example: `"Business, Events"`

**Note**: The `iab_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see the [IAB Content taxonomy](https://iabtechlab.com/standards/content-taxonomy/).
    
</dd>
</dl>

<dl>
<dd>

**notIabTags:** `*string` 

Inverse of the `iab_tags` parameter. Excludes articles based on Interactive Advertising Bureau (IAB) content categories. These tags provide a standardized taxonomy for digital advertising content categorization. To specify multiple IAB categories to exclude, use a comma-separated string. 

Example: `"Agriculture, Metals"`

**Note**: The `not_iab_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see the [IAB Content taxonomy](https://iabtechlab.com/standards/content-taxonomy/).
    
</dd>
</dl>

<dl>
<dd>

**customTags:** `*string` 

Filters articles based on provided taxonomy that is tailored to your specific needs and is accessible only with your API key. To specify tags, use the following pattern: 

- `custom_tags.taxonomy=Tag1,Tag2,Tag3`, where `taxonomy` is the taxonomy name and `Tag1,Tag2,Tag3` is a comma-separated list of tags.

Example: `custom_tags.industry="Manufacturing, Supply Chain, Logistics"`

To learn more, see the [Custom tags](/docs/v3/documentation/guides-and-concepts/custom-tags).
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*bool` — If true, returns only articles/sources that comply with the publisher's robots.txt rules. If false, returns only articles/sources that do not comply with robots.txt rules. If omitted, returns all articles/sources regardless of compliance status.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Authors.Post(request) -> *v505.AuthorsPostResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Searches for articles by author. You can filter results by language, country, source, and more.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.AuthorsPostRequest{
        AuthorName: "David Muir",
    }
client.Authors.Post(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**authorName:** `v505.AuthorName` 
    
</dd>
</dl>

<dl>
<dd>

**notAuthorName:** `*v505.NotAuthorName` 
    
</dd>
</dl>

<dl>
<dd>

**predefinedSources:** `*v505.PredefinedSources` 
    
</dd>
</dl>

<dl>
<dd>

**sources:** `*v505.Sources` 
    
</dd>
</dl>

<dl>
<dd>

**notSources:** `*v505.NotSources` 
    
</dd>
</dl>

<dl>
<dd>

**lang:** `*v505.Lang` 
    
</dd>
</dl>

<dl>
<dd>

**notLang:** `*v505.NotLang` 
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*v505.Countries` 
    
</dd>
</dl>

<dl>
<dd>

**notCountries:** `*v505.NotCountries` 
    
</dd>
</dl>

<dl>
<dd>

**from:** `*v505.From` 
    
</dd>
</dl>

<dl>
<dd>

**to:** `*v505.To` 
    
</dd>
</dl>

<dl>
<dd>

**publishedDatePrecision:** `*v505.PublishedDatePrecision` 
    
</dd>
</dl>

<dl>
<dd>

**byParseDate:** `*v505.ByParseDate` 
    
</dd>
</dl>

<dl>
<dd>

**sortBy:** `*v505.SortBy` 
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*v505.RankedOnly` 
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*v505.FromRank` 
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*v505.ToRank` 
    
</dd>
</dl>

<dl>
<dd>

**isHeadline:** `*v505.IsHeadline` 
    
</dd>
</dl>

<dl>
<dd>

**isOpinion:** `*v505.IsOpinion` 
    
</dd>
</dl>

<dl>
<dd>

**isPaidContent:** `*v505.IsPaidContent` 
    
</dd>
</dl>

<dl>
<dd>

**parentUrl:** `*v505.ParentUrl` 
    
</dd>
</dl>

<dl>
<dd>

**allLinks:** `*v505.AllLinks` 
    
</dd>
</dl>

<dl>
<dd>

**allDomainLinks:** `*v505.AllDomainLinks` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMin:** `*v505.WordCountMin` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMax:** `*v505.WordCountMax` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*v505.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*v505.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**includeTranslationFields:** `*v505.IncludeTranslationFields` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*v505.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*v505.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*v505.Theme` 
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*v505.NotTheme` 
    
</dd>
</dl>

<dl>
<dd>

**nerName:** `*v505.NerName` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*v505.TitleSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*v505.TitleSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*v505.ContentSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*v505.ContentSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**iptcTags:** `*v505.IptcTags` 
    
</dd>
</dl>

<dl>
<dd>

**notIptcTags:** `*v505.NotIptcTags` 
    
</dd>
</dl>

<dl>
<dd>

**iabTags:** `*v505.IabTags` 
    
</dd>
</dl>

<dl>
<dd>

**notIabTags:** `*v505.NotIabTags` 
    
</dd>
</dl>

<dl>
<dd>

**customTags:** `*v505.CustomTags` 
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*v505.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## SearchLink
<details><summary><code>client.SearchLink.SearchUrlGet() -> *v505.SearchResponseDto</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Searches for articles based on specified links or IDs. You can filter results by date range.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.SearchUrlGetRequest{
        Ids: v505.String(
            "5f8d0d55b6e45e00179c6e7e",
        ),
        Links: v505.String(
            "https://nytimes.com/article1",
        ),
        Source: v505.String(
            "articles.id,articles.title,articles.link,articles.published_date",
        ),
        From: &v505.From{
            DateTime: v505.MustParseDateTime(
                "2024-07-01T00:00:00Z",
            ),
        },
        To: &v505.To{
            DateTime: v505.MustParseDateTime(
                "2024-01-01T00:00:00Z",
            ),
        },
    }
client.SearchLink.SearchUrlGet(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `*string` 

The Newscatcher article ID (corresponds to the `_id` field in API response) or a list of article IDs to search for. To specify multiple IDs, use a comma-separated string. 

Example: `"1234567890abcdef, abcdef1234567890"`

**Caution**: You can use either the `links` or the `ids` parameter, but not both at the same time.
    
</dd>
</dl>

<dl>
<dd>

**links:** `*string` 

The article link or list of article links to search for. To specify multiple links, use a comma-separated string. 

Example: `"https://example.com/article1, https://example.com/article2"`

**Caution**: You can use either the `links` or the `ids` parameter, but not both at the same time.
    
</dd>
</dl>

<dl>
<dd>

**source:** `*v505.Source` 
    
</dd>
</dl>

<dl>
<dd>

**from:** `*v505.From` 
    
</dd>
</dl>

<dl>
<dd>

**to:** `*v505.To` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 

The page number to scroll through the results. Use for pagination, as a single API response can return up to 1,000 articles. 

For details, see [How to paginate large datasets](https://www.newscatcherapi.com/docs/v3/documentation/how-to/paginate-large-datasets).
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The number of articles to return per page.
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*bool` — If true, returns only articles/sources that comply with the publisher's robots.txt rules. If false, returns only articles/sources that do not comply with robots.txt rules. If omitted, returns all articles/sources regardless of compliance status.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SearchLink.SearchUrlPost(request) -> *v505.SearchResponseDto</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Searches for articles using their ID(s) or link(s).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.SearchUrlPostRequest{
        Links: &v505.Links{
            String: "https://www.reuters.com/business/energy/oil-prices-up-after-israeli-attacks-oversupply-caps-gains-2025-09-10/",
        },
        Source: v505.String(
            "articles.id,articles.title,articles.link,articles.canonical_url",
        ),
    }
client.SearchLink.SearchUrlPost(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ids:** `*v505.Ids` 
    
</dd>
</dl>

<dl>
<dd>

**links:** `*v505.Links` 
    
</dd>
</dl>

<dl>
<dd>

**source:** `*v505.Source` 
    
</dd>
</dl>

<dl>
<dd>

**from:** `*v505.From` 

The starting point in time to search from. Accepts date-time strings in ISO 8601 format and plain text strings. The default time zone is UTC. 

Formats with examples:
- YYYY-mm-ddTHH:MM:SS: `2024-07-01T00:00:00`
- YYYY-MM-dd: `2024-07-01`
- YYYY/mm/dd HH:MM:SS: `2024/07/01 00:00:00`
- YYYY/mm/dd: `2024/07/01`
- English phrases: `1 day ago`, `today`
    
</dd>
</dl>

<dl>
<dd>

**to:** `*v505.To` 

The ending point in time to search up to. Accepts date-time strings in ISO 8601 format and plain text strings. The default time zone is UTC. 

Formats with examples:
- YYYY-mm-ddTHH:MM:SS: `2024-07-01T00:00:00`
- YYYY-MM-dd: `2024-07-01`
- YYYY/mm/dd HH:MM:SS: `2024/07/01 00:00:00`
- YYYY/mm/dd: `2024/07/01`
- English phrases: `1 day ago`, `now`
    
</dd>
</dl>

<dl>
<dd>

**page:** `*v505.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*v505.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*v505.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## SearchSimilar
<details><summary><code>client.Searchsimilar.Get() -> *v505.SearchSimilarGetResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Searches for articles similar to a specified query.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.SearchSimilarGetRequest{
        Q: `"supply chain" AND Amazon NOT China`,
        SearchIn: v505.String(
            "title_content, title_content_translated",
        ),
        IncludeTranslationFields: v505.Bool(
            true,
        ),
        SimilarDocumentsFields: v505.String(
            "title,summary",
        ),
        PredefinedSources: v505.String(
            "top 100 US, top 5 GB",
        ),
        Sources: v505.String(
            "nytimes.com",
        ),
        NotSources: v505.String(
            "cnn.com",
        ),
        Lang: v505.String(
            "en",
        ),
        NotLang: v505.String(
            "fr",
        ),
        Countries: v505.String(
            "US",
        ),
        NotCountries: v505.String(
            "UK",
        ),
        From: &v505.SearchSimilarGetRequestFrom{
            DateTime: v505.MustParseDateTime(
                "2024-07-01T00:00:00Z",
            ),
        },
        To: &v505.SearchSimilarGetRequestTo{
            DateTime: v505.MustParseDateTime(
                "2024-07-01T00:00:00Z",
            ),
        },
        ParentUrl: v505.String(
            "https://www.washingtonpost.com/politics",
        ),
        AllLinks: v505.String(
            "https://aiindex.stanford.edu/report",
        ),
        AllDomainLinks: v505.String(
            "nvidia.com",
        ),
        IncludeNlpData: v505.Bool(
            true,
        ),
        HasNlp: v505.Bool(
            true,
        ),
        Theme: v505.String(
            "Business,Finance",
        ),
        NotTheme: v505.String(
            "Crime",
        ),
        NerName: v505.String(
            "Tesla",
        ),
        IptcTags: v505.String(
            "20000199,20000209",
        ),
        NotIptcTags: v505.String(
            "20000205,20000209",
        ),
        CustomTags: v505.String(
            "Tag1,Tag2,Tag3",
        ),
    }
client.Searchsimilar.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**q:** `v505.Q` 
    
</dd>
</dl>

<dl>
<dd>

**searchIn:** `*v505.SearchIn` 
    
</dd>
</dl>

<dl>
<dd>

**includeTranslationFields:** `*v505.IncludeTranslationFields` 
    
</dd>
</dl>

<dl>
<dd>

**includeSimilarDocuments:** `*bool` — If true, includes similar documents in the response.
    
</dd>
</dl>

<dl>
<dd>

**similarDocumentsNumber:** `*int` — The number of similar documents to return.
    
</dd>
</dl>

<dl>
<dd>

**similarDocumentsFields:** `*string` — The fields to consider for finding similar documents.
    
</dd>
</dl>

<dl>
<dd>

**predefinedSources:** `*string` 

Predefined top news sources per country. 

Format: start with the word `top`, followed by the number of desired sources, and then the two-letter country code [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). Multiple countries with the number of top sources can be specified as a comma-separated string.

Examples: 
- `"top 100 US"`
- `"top 33 AT"`
- `"top 50 US, top 20 GB"`
- `"top 33 AT, top 50 IT"`
    
</dd>
</dl>

<dl>
<dd>

**sources:** `*string` 

One or more news sources to narrow down the search. The format must be a domain URL. Subdomains, such as `finance.yahoo.com`, are also acceptable.To specify multiple sources, use a comma-separated string.

Examples:
- `"nytimes.com"`
- `"theguardian.com, finance.yahoo.com"`
    
</dd>
</dl>

<dl>
<dd>

**notSources:** `*string` 

The news sources to exclude from the search. To exclude multiple sources, use a comma-separated string. 

Example: `"cnn.com, wsj.com"`
    
</dd>
</dl>

<dl>
<dd>

**lang:** `*string` 

The language(s) of the search. The only accepted format is the two-letter [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code. To select multiple languages, use a comma-separated string. 

Example: `"en, es"`

To learn more, see [Enumerated parameters > Language](/docs/v3/api-reference/overview/enumerated-parameters#language-lang-and-not-lang).
    
</dd>
</dl>

<dl>
<dd>

**notLang:** `*string` 

The language(s) to exclude from the search. The accepted format is the two-letter [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code. To exclude multiple languages, use a comma-separated string. 

Example: `"fr, de"`

To learn more, see [Enumerated parameters > Language](/docs/v3/api-reference/overview/enumerated-parameters#language-lang-and-not-lang).
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*string` 

The countries where the news publisher is located. The accepted format is the two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. To select multiple countries, use a comma-separated string.

Example: `"US, CA"`

To learn more, see [Enumerated parameters > Country](/docs/v3/api-reference/overview/enumerated-parameters#country-country-and-not-country).
    
</dd>
</dl>

<dl>
<dd>

**notCountries:** `*string` 

The publisher location countries to exclude from the search. The accepted format is the two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. To exclude multiple countries, use a comma-separated string. 

Example:`"US, CA"`

To learn more, see [Enumerated parameters > Country](/docs/v3/api-reference/overview/enumerated-parameters#country-country-and-not-country).
    
</dd>
</dl>

<dl>
<dd>

**from:** `*v505.SearchSimilarGetRequestFrom` 

The starting point in time to search from. Accepts date-time strings in ISO 8601 format and plain text. The default time zone is UTC. 

Formats with examples:
- YYYY-mm-ddTHH:MM:SS: `2024-07-01T00:00:00`
- YYYY-MM-dd: `2024-07-01`
- YYYY/mm/dd HH:MM:SS: `2024/07/01 00:00:00`
- YYYY/mm/dd: `2024/07/01`
- English phrases: `7 day ago`, `today`

**Note**: By default, applied to the publication date of the article. To use the article's parse date instead, set the `by_parse_date` parameter to `true`.
    
</dd>
</dl>

<dl>
<dd>

**to:** `*v505.SearchSimilarGetRequestTo` 

The ending point in time to search up to. Accepts date-time strings in ISO 8601 format and plain text. The default time zone is UTC. 

Formats with examples:
- YYYY-mm-ddTHH:MM:SS: `2024-07-01T00:00:00`
- YYYY-MM-dd: `2024-07-01`
- YYYY/mm/dd HH:MM:SS: `2024/07/01 00:00:00`
- YYYY/mm/dd: `2024/07/01`
- English phrases: `1 day ago`, `now`

**Note**: By default, applied to the publication date of the article. To use the article's parse date instead, set the `by_parse_date` parameter to `true`.
    
</dd>
</dl>

<dl>
<dd>

**byParseDate:** `*bool` — If true, the `from_` and `to_` parameters use article parse dates instead of published dates. Additionally, the `parse_date` variable is added to the output for each article object.
    
</dd>
</dl>

<dl>
<dd>

**publishedDatePrecision:** `*v505.SearchSimilarGetRequestPublishedDatePrecision` 

The precision of the published date. There are three types:
- `full`: The day and time of an article is correctly identified with the appropriate timezone.
- `timezone unknown`: The day and time of an article is correctly identified without timezone.
- `date`: Only the day is identified without an exact time.
    
</dd>
</dl>

<dl>
<dd>

**sortBy:** `*v505.SearchSimilarGetRequestSortBy` 

The sorting order of the results. Possible values are:
- `relevancy`: The most relevant results first.
- `date`: The most recently published results first.
- `rank`: The results from the highest-ranked sources first.
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*bool` — If true, limits the search to sources ranked in the top 1 million online websites. If false, includes unranked sources which are assigned a rank of 999999.
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*int` — The lowest boundary of the rank of a news website to filter by. A lower rank indicates a more popular source.
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*int` — The highest boundary of the rank of a news website to filter by. A lower rank indicates a more popular source.
    
</dd>
</dl>

<dl>
<dd>

**isHeadline:** `*bool` — If true, only returns articles that were posted on the home page of a given news domain.
    
</dd>
</dl>

<dl>
<dd>

**isOpinion:** `*bool` — If true, returns only opinion pieces. If false, excludes opinion-based articles and returns news only.
    
</dd>
</dl>

<dl>
<dd>

**isPaidContent:** `*bool` — If false, returns only articles that have publicly available complete content. Some publishers partially block content, so this setting ensures that only full articles are retrieved.
    
</dd>
</dl>

<dl>
<dd>

**parentUrl:** `*string` 

The categorical URL(s) to filter your search. To filter your search by multiple categorical URLs, use a comma-separated string.

Example: `"wsj.com/politics, wsj.com/tech"`
    
</dd>
</dl>

<dl>
<dd>

**allLinks:** `*string` 

The complete URL(s) mentioned in the article. For multiple URLs, use a comma-separated string.

Example: `"https://aiindex.stanford.edu/report, https://www.stateof.ai"`

For more details, see [Search by URL](/docs/v3/documentation/how-to/search-by-url).
    
</dd>
</dl>

<dl>
<dd>

**allDomainLinks:** `*string` 

The domain(s) mentioned in the article. For multiple domains, use a comma-separated string.

Example: `"who.int, nih.gov"`

For more details, see [Search by URL](/docs/v3/documentation/how-to/search-by-url).
    
</dd>
</dl>

<dl>
<dd>

**wordCountMin:** `*int` — The minimum number of words an article must contain. To be used for avoiding articles with small content.
    
</dd>
</dl>

<dl>
<dd>

**wordCountMax:** `*int` — The maximum number of words an article can contain. To be used for avoiding articles with large content.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 

The page number to scroll through the results. Use for pagination, as a single API response can return up to 1,000 articles. 

For details, see [How to paginate large datasets](https://www.newscatcherapi.com/docs/v3/documentation/how-to/paginate-large-datasets).
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The number of articles to return per page.
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*v505.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*v505.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*string` 

Filters articles based on their general topic, as determined by NLP analysis. To select multiple themes, use a comma-separated string.

Example: `"Finance, Tech"`

**Note**: The `theme` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).

Available options: `Business`, `Economics`, `Entertainment`, `Finance`, `Health`, `Politics`, `Science`, `Sports`, `Tech`, `Crime`, `Financial Crime`, `Lifestyle`, `Automotive`, `Travel`, `Weather`, `General`.
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*string` 

Inverse of the `theme` parameter. Excludes articles based on their general topic, as determined by NLP analysis. To exclude multiple themes, use a comma-separated string. 

Example: `"Crime, Tech"`

**Note**: The `not_theme` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**nerName:** `*string` 

The name of person, organization, location, product or other named entity to search for. To specify multiple names use a comma-separated string. 

Example: `"Tesla, Amazon"`
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*float64` 

Filters articles based on the minimum sentiment score of their titles.

Range is `-1.0` to `1.0`, where:
- Negative values indicate negative sentiment.
- Positive values indicate positive sentiment.
- Values close to 0 indicate neutral sentiment.

**Note**: The `title_sentiment_min` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*float64` 

Filters articles based on the maximum sentiment score of their titles.

Range is `-1.0` to `1.0`, where:
- Negative values indicate negative sentiment.
- Positive values indicate positive sentiment.
- Values close to 0 indicate neutral sentiment.

**Note**: The `title_sentiment_max` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*float64` 

Filters articles based on the minimum sentiment score of their content.

Range is `-1.0` to `1.0`, where:
- Negative values indicate negative sentiment.
- Positive values indicate positive sentiment.
- Values close to 0 indicate neutral sentiment.

**Note**: The `content_sentiment_min` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*float64` 

Filters articles based on the maximum sentiment score of their content.

Range is `-1.0` to `1.0`, where:
- Negative values indicate negative sentiment.
- Positive values indicate positive sentiment.
- Values close to 0 indicate neutral sentiment.

**Note**: The `content_sentiment_max` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**iptcTags:** `*string` 

Filters articles based on International Press Telecommunications Council (IPTC) media topic tags. To specify multiple IPTC tags, use a comma-separated string of tag IDs. 

Example: `"20000199, 20000209"`

**Note**: The `iptc_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see [IPTC Media Topic NewsCodes](https://www.iptc.org/std/NewsCodes/treeview/mediatopic/mediatopic-en-GB.html).
    
</dd>
</dl>

<dl>
<dd>

**notIptcTags:** `*string` 

Inverse of the `iptc_tags` parameter. Excludes articles based on International Press Telecommunications Council (IPTC) media topic tags. To specify multiple IPTC tags to exclude, use a comma-separated string of tag IDs. 

Example: `"20000205, 20000209"`

**Note**: The `not_iptc_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see [IPTC Media Topic NewsCodes](https://www.iptc.org/std/NewsCodes/treeview/mediatopic/mediatopic-en-GB.html).
    
</dd>
</dl>

<dl>
<dd>

**customTags:** `*string` 

Filters articles based on provided taxonomy that is tailored to your specific needs and is accessible only with your API key. To specify tags, use the following pattern: 

- `custom_tags.taxonomy=Tag1,Tag2,Tag3`, where `taxonomy` is the taxonomy name and `Tag1,Tag2,Tag3` is a comma-separated list of tags.

Example: `custom_tags.industry="Manufacturing, Supply Chain, Logistics"`

To learn more, see the [Custom tags](/docs/v3/documentation/guides-and-concepts/custom-tags).
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*bool` — If true, returns only articles/sources that comply with the publisher's robots.txt rules. If false, returns only articles/sources that do not comply with robots.txt rules. If omitted, returns all articles/sources regardless of compliance status.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Searchsimilar.Post(request) -> *v505.SearchSimilarPostResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Searches for articles similar to the specified query. You can filter results by language, country, source, and more.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.SearchSimilarPostRequest{
        Q: `"supply chain" AND Amazon NOT China`,
        IncludeSimilarDocuments: v505.Bool(
            true,
        ),
        SimilarDocumentsNumber: v505.Int(
            5,
        ),
        PageSize: v505.Int(
            10,
        ),
    }
client.Searchsimilar.Post(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**q:** `v505.Q` 
    
</dd>
</dl>

<dl>
<dd>

**searchIn:** `*v505.SearchIn` 
    
</dd>
</dl>

<dl>
<dd>

**includeTranslationFields:** `*v505.IncludeTranslationFields` 
    
</dd>
</dl>

<dl>
<dd>

**includeSimilarDocuments:** `*v505.IncludeSimilarDocuments` 
    
</dd>
</dl>

<dl>
<dd>

**similarDocumentsNumber:** `*v505.SimilarDocumentsNumber` 
    
</dd>
</dl>

<dl>
<dd>

**similarDocumentsFields:** `*v505.SimilarDocumentsFields` 
    
</dd>
</dl>

<dl>
<dd>

**predefinedSources:** `*v505.PredefinedSources` 
    
</dd>
</dl>

<dl>
<dd>

**sources:** `*v505.Sources` 
    
</dd>
</dl>

<dl>
<dd>

**notSources:** `*v505.NotSources` 
    
</dd>
</dl>

<dl>
<dd>

**lang:** `*v505.Lang` 
    
</dd>
</dl>

<dl>
<dd>

**notLang:** `*v505.NotLang` 
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*v505.Countries` 
    
</dd>
</dl>

<dl>
<dd>

**notCountries:** `*v505.NotCountries` 
    
</dd>
</dl>

<dl>
<dd>

**from:** `*v505.From` 
    
</dd>
</dl>

<dl>
<dd>

**to:** `*v505.To` 
    
</dd>
</dl>

<dl>
<dd>

**byParseDate:** `*v505.ByParseDate` 
    
</dd>
</dl>

<dl>
<dd>

**publishedDatePrecision:** `*v505.PublishedDatePrecision` 
    
</dd>
</dl>

<dl>
<dd>

**sortBy:** `*v505.SortBy` 
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*v505.RankedOnly` 
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*v505.FromRank` 
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*v505.ToRank` 
    
</dd>
</dl>

<dl>
<dd>

**isHeadline:** `*v505.IsHeadline` 
    
</dd>
</dl>

<dl>
<dd>

**isOpinion:** `*v505.IsOpinion` 
    
</dd>
</dl>

<dl>
<dd>

**isPaidContent:** `*v505.IsPaidContent` 
    
</dd>
</dl>

<dl>
<dd>

**parentUrl:** `*v505.ParentUrl` 
    
</dd>
</dl>

<dl>
<dd>

**allLinks:** `*v505.AllLinks` 
    
</dd>
</dl>

<dl>
<dd>

**allDomainLinks:** `*v505.AllDomainLinks` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMin:** `*v505.WordCountMin` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMax:** `*v505.WordCountMax` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*v505.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*v505.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*v505.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*v505.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*v505.Theme` 
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*v505.NotTheme` 
    
</dd>
</dl>

<dl>
<dd>

**nerName:** `*v505.NerName` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*v505.TitleSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*v505.TitleSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*v505.ContentSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*v505.ContentSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**iptcTags:** `*v505.IptcTags` 
    
</dd>
</dl>

<dl>
<dd>

**notIptcTags:** `*v505.NotIptcTags` 
    
</dd>
</dl>

<dl>
<dd>

**customTags:** `*v505.CustomTags` 
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*v505.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Sources
<details><summary><code>client.Sources.Get() -> *v505.SourcesResponseDto</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves a list of sources based on specified criteria such as language, country, rank, and more.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.SourcesGetRequest{
        Lang: v505.String(
            "en",
        ),
        Countries: v505.String(
            "US",
        ),
        PredefinedSources: v505.String(
            "top 100 US, top 5 GB",
        ),
        SourceName: v505.String(
            "sport",
        ),
        SourceUrl: v505.String(
            "bbc.com",
        ),
        NewsType: v505.String(
            "General News Outlets",
        ),
    }
client.Sources.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**lang:** `*string` 

The language(s) of the search. The only accepted format is the two-letter [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code. To select multiple languages, use a comma-separated string. 

Example: `"en, es"`

To learn more, see [Enumerated parameters > Language](/docs/v3/api-reference/overview/enumerated-parameters#language-lang-and-not-lang).
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*string` 

The countries where the news publisher is located. The accepted format is the two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. To select multiple countries, use a comma-separated string.

Example: `"US, CA"`

To learn more, see [Enumerated parameters > Country](/docs/v3/api-reference/overview/enumerated-parameters#country-country-and-not-country).
    
</dd>
</dl>

<dl>
<dd>

**predefinedSources:** `*string` 

Predefined top news sources per country. 

Format: start with the word `top`, followed by the number of desired sources, and then the two-letter country code [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). Multiple countries with the number of top sources can be specified as a comma-separated string.

Examples: 
- `"top 100 US"`
- `"top 33 AT"`
- `"top 50 US, top 20 GB"`
- `"top 33 AT, top 50 IT"`
    
</dd>
</dl>

<dl>
<dd>

**sourceName:** `*string` 

Word or phrase to search within the source names. To specify multiple values, use a comma-separated string.

Example: `"sport, tech"`

**Note**: The search doesn't require an exact match and returns sources containing the specified terms in their names. You can use any word or phrase, like `"sport"` or `"new york times"`. For example, `"sport"` returns sources such as `"Motorsport"`, `"Dot Esport"`, and `"Tuttosport"`.
    
</dd>
</dl>

<dl>
<dd>

**sourceUrl:** `*string` 

The domain(s) of the news publication to search for. 

**Caution**:  When specifying the `source_url` parameter, 
you can only use `include_additional_info` as an extra parameter.
    
</dd>
</dl>

<dl>
<dd>

**includeAdditionalInfo:** `*bool` 

If true, returns the following additional datapoints about each news source:
- `nb_articles_for_7d`: The number of articles published by the source in the last week.
- `country`: Source country of origin.
- `rank`: SEO rank.
- `is_news_domain`: Boolean indicating if the source is a news domain.
- `news_domain_type`: Type of news domain (e.g., "Original Content").
- `news_type`: Category of news (e.g., "General News Outlets").
    
</dd>
</dl>

<dl>
<dd>

**isNewsDomain:** `*bool` — If true, filters results to include only news domains.
    
</dd>
</dl>

<dl>
<dd>

**newsDomainType:** `*v505.SourcesGetRequestNewsDomainType` 

Filters results based on the news domain type. Possible values are:
- `Original Content`: Sources that produce their own content.
- `Aggregator`: Sources that collect content from various other sources.
- `Press Releases`: Sources primarily publishing press releases.
- `Republisher`: Sources that republish content from other sources.
- `Other`: Sources that don't fit into main categories.
    
</dd>
</dl>

<dl>
<dd>

**newsType:** `*string` 

Filters results based on the news type. Multiple types can be specified using a comma-separated string.

Example: `"General News Outlets,Tech News and Updates"`

For a complete list of available news types, see [Enumerated parameters > News type](/docs/v3/api-reference/overview/enumerated-parameters#news-type-news-type).
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*int` — The lowest boundary of the rank of a news website to filter by. A lower rank indicates a more popular source.
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*int` — The highest boundary of the rank of a news website to filter by. A lower rank indicates a more popular source.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sources.Post(request) -> *v505.SourcesResponseDto</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the list of sources available in the database. You can filter the sources by language, country, and more.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.SourcesPostRequest{
        PredefinedSources: &v505.PredefinedSources{
            String: "top 10 US",
        },
    }
client.Sources.Post(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**lang:** `*v505.Lang` 
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*v505.Countries` 
    
</dd>
</dl>

<dl>
<dd>

**predefinedSources:** `*v505.PredefinedSources` 
    
</dd>
</dl>

<dl>
<dd>

**sourceName:** `*v505.SourceName` 
    
</dd>
</dl>

<dl>
<dd>

**sourceUrl:** `*v505.SourceUrl` 
    
</dd>
</dl>

<dl>
<dd>

**includeAdditionalInfo:** `*v505.IncludeAdditionalInfo` 
    
</dd>
</dl>

<dl>
<dd>

**isNewsDomain:** `*v505.IsNewsDomain` 
    
</dd>
</dl>

<dl>
<dd>

**newsDomainType:** `*v505.NewsDomainType` 
    
</dd>
</dl>

<dl>
<dd>

**newsType:** `*v505.NewsType` 
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*v505.FromRank` 
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*v505.ToRank` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Aggregation
<details><summary><code>client.Aggregation.Get() -> *v505.AggregationGetResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the count of articles aggregated by day or hour based on various search criteria, such as keyword, language, country, and source.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.AggregationGetRequest{
        Q: `"supply chain" AND Amazon NOT China`,
        SearchIn: v505.String(
            "title_content, title_content_translated",
        ),
        PredefinedSources: v505.String(
            "top 100 US, top 5 GB",
        ),
        Sources: v505.String(
            "nytimes.com",
        ),
        NotSources: v505.String(
            "cnn.com",
        ),
        Lang: v505.String(
            "en",
        ),
        NotLang: v505.String(
            "fr",
        ),
        Countries: v505.String(
            "US",
        ),
        NotCountries: v505.String(
            "UK",
        ),
        NotAuthorName: v505.String(
            "John Doe",
        ),
        From: &v505.AggregationGetRequestFrom{
            DateTime: v505.MustParseDateTime(
                "2024-07-01T00:00:00Z",
            ),
        },
        To: &v505.AggregationGetRequestTo{
            DateTime: v505.MustParseDateTime(
                "2024-07-01T00:00:00Z",
            ),
        },
        ParentUrl: v505.String(
            "https://www.washingtonpost.com/politics",
        ),
        AllLinks: v505.String(
            "https://aiindex.stanford.edu/report",
        ),
        AllDomainLinks: v505.String(
            "nvidia.com",
        ),
        IncludeNlpData: v505.Bool(
            true,
        ),
        HasNlp: v505.Bool(
            true,
        ),
        Theme: v505.String(
            "Business,Finance",
        ),
        NotTheme: v505.String(
            "Crime",
        ),
        OrgEntityName: v505.String(
            "Apple",
        ),
        PerEntityName: v505.String(
            "Elon Musk",
        ),
        LocEntityName: v505.String(
            "California",
        ),
        MiscEntityName: v505.String(
            "Bitcoin",
        ),
        IptcTags: v505.String(
            "20000199,20000209",
        ),
        NotIptcTags: v505.String(
            "20000205,20000209",
        ),
    }
client.Aggregation.Get(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**q:** `v505.Q` 
    
</dd>
</dl>

<dl>
<dd>

**aggregationBy:** `*v505.AggregationBy` 
    
</dd>
</dl>

<dl>
<dd>

**searchIn:** `*v505.SearchIn` 
    
</dd>
</dl>

<dl>
<dd>

**predefinedSources:** `*string` 

Predefined top news sources per country. 

Format: start with the word `top`, followed by the number of desired sources, and then the two-letter country code [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). Multiple countries with the number of top sources can be specified as a comma-separated string.

Examples: 
- `"top 100 US"`
- `"top 33 AT"`
- `"top 50 US, top 20 GB"`
- `"top 33 AT, top 50 IT"`
    
</dd>
</dl>

<dl>
<dd>

**sources:** `*string` 

One or more news sources to narrow down the search. The format must be a domain URL. Subdomains, such as `finance.yahoo.com`, are also acceptable.To specify multiple sources, use a comma-separated string.

Examples:
- `"nytimes.com"`
- `"theguardian.com, finance.yahoo.com"`
    
</dd>
</dl>

<dl>
<dd>

**notSources:** `*string` 

The news sources to exclude from the search. To exclude multiple sources, use a comma-separated string. 

Example: `"cnn.com, wsj.com"`
    
</dd>
</dl>

<dl>
<dd>

**lang:** `*string` 

The language(s) of the search. The only accepted format is the two-letter [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code. To select multiple languages, use a comma-separated string. 

Example: `"en, es"`

To learn more, see [Enumerated parameters > Language](/docs/v3/api-reference/overview/enumerated-parameters#language-lang-and-not-lang).
    
</dd>
</dl>

<dl>
<dd>

**notLang:** `*string` 

The language(s) to exclude from the search. The accepted format is the two-letter [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code. To exclude multiple languages, use a comma-separated string. 

Example: `"fr, de"`

To learn more, see [Enumerated parameters > Language](/docs/v3/api-reference/overview/enumerated-parameters#language-lang-and-not-lang).
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*string` 

The countries where the news publisher is located. The accepted format is the two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. To select multiple countries, use a comma-separated string.

Example: `"US, CA"`

To learn more, see [Enumerated parameters > Country](/docs/v3/api-reference/overview/enumerated-parameters#country-country-and-not-country).
    
</dd>
</dl>

<dl>
<dd>

**notCountries:** `*string` 

The publisher location countries to exclude from the search. The accepted format is the two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. To exclude multiple countries, use a comma-separated string. 

Example:`"US, CA"`

To learn more, see [Enumerated parameters > Country](/docs/v3/api-reference/overview/enumerated-parameters#country-country-and-not-country).
    
</dd>
</dl>

<dl>
<dd>

**notAuthorName:** `*string` 

The list of author names to exclude from your search. To exclude articles by specific authors, use a comma-separated string.

Example: `"John Doe, Jane Doe"`
    
</dd>
</dl>

<dl>
<dd>

**from:** `*v505.AggregationGetRequestFrom` 

The starting point in time to search from. Accepts date-time strings in ISO 8601 format and plain text. The default time zone is UTC. 

Formats with examples:
- YYYY-mm-ddTHH:MM:SS: `2024-07-01T00:00:00`
- YYYY-MM-dd: `2024-07-01`
- YYYY/mm/dd HH:MM:SS: `2024/07/01 00:00:00`
- YYYY/mm/dd: `2024/07/01`
- English phrases: `7 day ago`, `today`

**Note**: By default, applied to the publication date of the article. To use the article's parse date instead, set the `by_parse_date` parameter to `true`.
    
</dd>
</dl>

<dl>
<dd>

**to:** `*v505.AggregationGetRequestTo` 

The ending point in time to search up to. Accepts date-time strings in ISO 8601 format and plain text. The default time zone is UTC. 

Formats with examples:
- YYYY-mm-ddTHH:MM:SS: `2024-07-01T00:00:00`
- YYYY-MM-dd: `2024-07-01`
- YYYY/mm/dd HH:MM:SS: `2024/07/01 00:00:00`
- YYYY/mm/dd: `2024/07/01`
- English phrases: `1 day ago`, `now`

**Note**: By default, applied to the publication date of the article. To use the article's parse date instead, set the `by_parse_date` parameter to `true`.
    
</dd>
</dl>

<dl>
<dd>

**publishedDatePrecision:** `*v505.AggregationGetRequestPublishedDatePrecision` 

The precision of the published date. There are three types:
- `full`: The day and time of an article is correctly identified with the appropriate timezone.
- `timezone unknown`: The day and time of an article is correctly identified without timezone.
- `date`: Only the day is identified without an exact time.
    
</dd>
</dl>

<dl>
<dd>

**byParseDate:** `*bool` — If true, the `from_` and `to_` parameters use article parse dates instead of published dates. Additionally, the `parse_date` variable is added to the output for each article object.
    
</dd>
</dl>

<dl>
<dd>

**sortBy:** `*v505.AggregationGetRequestSortBy` 

The sorting order of the results. Possible values are:
- `relevancy`: The most relevant results first.
- `date`: The most recently published results first.
- `rank`: The results from the highest-ranked sources first.
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*bool` — If true, limits the search to sources ranked in the top 1 million online websites. If false, includes unranked sources which are assigned a rank of 999999.
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*int` — The lowest boundary of the rank of a news website to filter by. A lower rank indicates a more popular source.
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*int` — The highest boundary of the rank of a news website to filter by. A lower rank indicates a more popular source.
    
</dd>
</dl>

<dl>
<dd>

**isHeadline:** `*bool` — If true, only returns articles that were posted on the home page of a given news domain.
    
</dd>
</dl>

<dl>
<dd>

**isOpinion:** `*bool` — If true, returns only opinion pieces. If false, excludes opinion-based articles and returns news only.
    
</dd>
</dl>

<dl>
<dd>

**isPaidContent:** `*bool` — If false, returns only articles that have publicly available complete content. Some publishers partially block content, so this setting ensures that only full articles are retrieved.
    
</dd>
</dl>

<dl>
<dd>

**parentUrl:** `*string` 

The categorical URL(s) to filter your search. To filter your search by multiple categorical URLs, use a comma-separated string.

Example: `"wsj.com/politics, wsj.com/tech"`
    
</dd>
</dl>

<dl>
<dd>

**allLinks:** `*string` 

The complete URL(s) mentioned in the article. For multiple URLs, use a comma-separated string.

Example: `"https://aiindex.stanford.edu/report, https://www.stateof.ai"`

For more details, see [Search by URL](/docs/v3/documentation/how-to/search-by-url).
    
</dd>
</dl>

<dl>
<dd>

**allDomainLinks:** `*string` 

The domain(s) mentioned in the article. For multiple domains, use a comma-separated string.

Example: `"who.int, nih.gov"`

For more details, see [Search by URL](/docs/v3/documentation/how-to/search-by-url).
    
</dd>
</dl>

<dl>
<dd>

**wordCountMin:** `*int` — The minimum number of words an article must contain. To be used for avoiding articles with small content.
    
</dd>
</dl>

<dl>
<dd>

**wordCountMax:** `*int` — The maximum number of words an article can contain. To be used for avoiding articles with large content.
    
</dd>
</dl>

<dl>
<dd>

**page:** `*int` 

The page number to scroll through the results. Use for pagination, as a single API response can return up to 1,000 articles. 

For details, see [How to paginate large datasets](https://www.newscatcherapi.com/docs/v3/documentation/how-to/paginate-large-datasets).
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — The number of articles to return per page.
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*v505.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*v505.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*string` 

Filters articles based on their general topic, as determined by NLP analysis. To select multiple themes, use a comma-separated string.

Example: `"Finance, Tech"`

**Note**: The `theme` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).

Available options: `Business`, `Economics`, `Entertainment`, `Finance`, `Health`, `Politics`, `Science`, `Sports`, `Tech`, `Crime`, `Financial Crime`, `Lifestyle`, `Automotive`, `Travel`, `Weather`, `General`.
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*string` 

Inverse of the `theme` parameter. Excludes articles based on their general topic, as determined by NLP analysis. To exclude multiple themes, use a comma-separated string. 

Example: `"Crime, Tech"`

**Note**: The `not_theme` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**orgEntityName:** `*string` 

Filters articles that mention specific organization names, as identified by NLP analysis. To specify multiple organizations, use a comma-separated string. To search named entities in translations, combine with the translation options of the `search_in` parameter (e.g., `title_content_translated`).

Example: `"Apple, Microsoft"`

**Note**: The `ORG_entity_name` parameter is only available if NLP is included in your subscription plan.

To learn more, see [Search by entity](/docs/v3/documentation/how-to/search-by-entity).
    
</dd>
</dl>

<dl>
<dd>

**perEntityName:** `*string` 

Filters articles that mention specific person names, as identified by NLP analysis. To specify multiple names, use a comma-separated string. To search named entities in translations, combine with the translation options of the `search_in` parameter (e.g., `title_content_translated`).

Example: `"Elon Musk, Jeff Bezos"`

**Note**: The `PER_entity_name` parameter is only available if NLP is included in your subscription plan.

To learn more, see [Search by entity](/docs/v3/documentation/how-to/search-by-entity).
    
</dd>
</dl>

<dl>
<dd>

**locEntityName:** `*string` 

Filters articles that mention specific location names, as identified by NLP analysis. To specify multiple locations, use a comma-separated string. To search named entities in translations, combine with the translation options of the `search_in` parameter (e.g., `title_content_translated`).

Example: `"California, New York"`

**Note**: The `LOC_entity_name` parameter is only available if NLP is included in your subscription plan.

To learn more, see [Search by entity](/docs/v3/documentation/how-to/search-by-entity).
    
</dd>
</dl>

<dl>
<dd>

**miscEntityName:** `*string` 

Filters articles that mention other named entities not falling under person, organization, or location categories. Includes events, nationalities, products, works of art, and more. To specify multiple entities, use a comma-separated string. To search named entities in translations, combine with the translation options of the `search_in` parameter (e.g., `title_content_translated`).

Example: `"Bitcoin, Blockchain"`

**Note**: The `MISC_entity_name` parameter is only available if NLP is included in your subscription plan.

To learn more, see [Search by entity](/docs/v3/documentation/how-to/search-by-entity).
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*float64` 

Filters articles based on the minimum sentiment score of their titles.

Range is `-1.0` to `1.0`, where:
- Negative values indicate negative sentiment.
- Positive values indicate positive sentiment.
- Values close to 0 indicate neutral sentiment.

**Note**: The `title_sentiment_min` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*float64` 

Filters articles based on the maximum sentiment score of their titles.

Range is `-1.0` to `1.0`, where:
- Negative values indicate negative sentiment.
- Positive values indicate positive sentiment.
- Values close to 0 indicate neutral sentiment.

**Note**: The `title_sentiment_max` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*float64` 

Filters articles based on the minimum sentiment score of their content.

Range is `-1.0` to `1.0`, where:
- Negative values indicate negative sentiment.
- Positive values indicate positive sentiment.
- Values close to 0 indicate neutral sentiment.

**Note**: The `content_sentiment_min` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*float64` 

Filters articles based on the maximum sentiment score of their content.

Range is `-1.0` to `1.0`, where:
- Negative values indicate negative sentiment.
- Positive values indicate positive sentiment.
- Values close to 0 indicate neutral sentiment.

**Note**: The `content_sentiment_max` parameter is only available if NLP is included in your subscription plan.

To learn more, see [NLP features](/docs/v3/documentation/guides-and-concepts/nlp-features).
    
</dd>
</dl>

<dl>
<dd>

**iptcTags:** `*string` 

Filters articles based on International Press Telecommunications Council (IPTC) media topic tags. To specify multiple IPTC tags, use a comma-separated string of tag IDs. 

Example: `"20000199, 20000209"`

**Note**: The `iptc_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see [IPTC Media Topic NewsCodes](https://www.iptc.org/std/NewsCodes/treeview/mediatopic/mediatopic-en-GB.html).
    
</dd>
</dl>

<dl>
<dd>

**notIptcTags:** `*string` 

Inverse of the `iptc_tags` parameter. Excludes articles based on International Press Telecommunications Council (IPTC) media topic tags. To specify multiple IPTC tags to exclude, use a comma-separated string of tag IDs. 

Example: `"20000205, 20000209"`

**Note**: The `not_iptc_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see [IPTC Media Topic NewsCodes](https://www.iptc.org/std/NewsCodes/treeview/mediatopic/mediatopic-en-GB.html).
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*bool` — If true, returns only articles/sources that comply with the publisher's robots.txt rules. If false, returns only articles/sources that do not comply with robots.txt rules. If omitted, returns all articles/sources regardless of compliance status.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Aggregation.Post(request) -> *v505.AggregationPostResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves the count of articles aggregated by day or hour based on various search criteria, such as keyword, language, country, and source.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &v505.AggregationPostRequest{
        Q: `"supply chain" AND Amazon NOT China`,
        AggregationBy: v505.AggregationByDay.Ptr(),
    }
client.Aggregation.Post(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**q:** `v505.Q` 
    
</dd>
</dl>

<dl>
<dd>

**aggregationBy:** `*v505.AggregationBy` 
    
</dd>
</dl>

<dl>
<dd>

**searchIn:** `*v505.SearchIn` 
    
</dd>
</dl>

<dl>
<dd>

**predefinedSources:** `*v505.PredefinedSources` 
    
</dd>
</dl>

<dl>
<dd>

**sources:** `*v505.Sources` 
    
</dd>
</dl>

<dl>
<dd>

**notSources:** `*v505.NotSources` 
    
</dd>
</dl>

<dl>
<dd>

**lang:** `*v505.Lang` 
    
</dd>
</dl>

<dl>
<dd>

**notLang:** `*v505.NotLang` 
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*v505.Countries` 
    
</dd>
</dl>

<dl>
<dd>

**notCountries:** `*v505.NotCountries` 
    
</dd>
</dl>

<dl>
<dd>

**notAuthorName:** `*v505.NotAuthorName` 
    
</dd>
</dl>

<dl>
<dd>

**from:** `*v505.From` 
    
</dd>
</dl>

<dl>
<dd>

**to:** `*v505.To` 
    
</dd>
</dl>

<dl>
<dd>

**publishedDatePrecision:** `*v505.PublishedDatePrecision` 
    
</dd>
</dl>

<dl>
<dd>

**byParseDate:** `*v505.ByParseDate` 
    
</dd>
</dl>

<dl>
<dd>

**sortBy:** `*v505.SortBy` 
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*v505.RankedOnly` 
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*v505.FromRank` 
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*v505.ToRank` 
    
</dd>
</dl>

<dl>
<dd>

**isHeadline:** `*v505.IsHeadline` 
    
</dd>
</dl>

<dl>
<dd>

**isOpinion:** `*v505.IsOpinion` 
    
</dd>
</dl>

<dl>
<dd>

**isPaidContent:** `*v505.IsPaidContent` 
    
</dd>
</dl>

<dl>
<dd>

**parentUrl:** `*v505.ParentUrl` 
    
</dd>
</dl>

<dl>
<dd>

**allLinks:** `*v505.AllLinks` 
    
</dd>
</dl>

<dl>
<dd>

**allDomainLinks:** `*v505.AllDomainLinks` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMin:** `*v505.WordCountMin` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMax:** `*v505.WordCountMax` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*v505.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*v505.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*v505.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*v505.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*v505.Theme` 
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*v505.NotTheme` 
    
</dd>
</dl>

<dl>
<dd>

**orgEntityName:** `*v505.OrgEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**perEntityName:** `*v505.PerEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**locEntityName:** `*v505.LocEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**miscEntityName:** `*v505.MiscEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*v505.TitleSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*v505.TitleSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*v505.ContentSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*v505.ContentSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**iptcTags:** `*v505.IptcTags` 
    
</dd>
</dl>

<dl>
<dd>

**notIptcTags:** `*v505.NotIptcTags` 
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*v505.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Subscription
<details><summary><code>client.Subscription.Get() -> *v505.SubscriptionResponseDto</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves information about your subscription plan.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Subscription.Get(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Subscription.Post() -> *v505.SubscriptionResponseDto</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves information about your subscription plan.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Subscription.Post(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

