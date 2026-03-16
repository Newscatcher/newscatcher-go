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

Searches for articles based on specified criteria such as keywords, language, country, source, and more.
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
            "top 50 US, top 20 GB",
        ),
        SourceName: v505.String(
            "sport,tech",
        ),
        Sources: v505.String(
            "nytimes.com,finance.yahoo.com",
        ),
        NotSources: v505.String(
            "cnn.com,wsj.com",
        ),
        Lang: v505.String(
            "en,es",
        ),
        NotLang: v505.String(
            "fr,de",
        ),
        Countries: v505.String(
            "US,CA",
        ),
        NotCountries: v505.String(
            "UK,FR",
        ),
        NotAuthorName: v505.String(
            "John Doe, Jane Doe",
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
        PublishedDatePrecision: v505.String(
            "full",
        ),
        ByParseDate: v505.Bool(
            true,
        ),
        RankedOnly: v505.Bool(
            true,
        ),
        FromRank: v505.Int(
            100,
        ),
        ToRank: v505.Int(
            100,
        ),
        IsHeadline: v505.Bool(
            true,
        ),
        IsOpinion: v505.Bool(
            true,
        ),
        IsPaidContent: v505.Bool(
            false,
        ),
        ParentUrl: v505.String(
            "wsj.com/politics,wsj.com/tech",
        ),
        AllLinks: v505.String(
            "https://aiindex.stanford.edu/report,https://www.stateof.ai",
        ),
        AllDomainLinks: v505.String(
            "who.int,nih.gov",
        ),
        AllLinksText: v505.String(
            "Nvidia,Tesla",
        ),
        AdditionalDomainInfo: v505.Bool(
            true,
        ),
        IsNewsDomain: v505.Bool(
            true,
        ),
        NewsType: v505.String(
            "General News Outlets,Tech News and Updates",
        ),
        WordCountMin: v505.Int(
            300,
        ),
        WordCountMax: v505.Int(
            1000,
        ),
        Page: v505.Int(
            2,
        ),
        PageSize: v505.Int(
            50,
        ),
        ClusteringEnabled: v505.Bool(
            true,
        ),
        ClusteringThreshold: v505.Float64(
            0.6,
        ),
        IncludeNlpData: v505.Bool(
            true,
        ),
        HasNlp: v505.Bool(
            true,
        ),
        Theme: v505.String(
            "Finance,Tech",
        ),
        NotTheme: v505.String(
            "Crime,Sports",
        ),
        OrgEntityName: v505.String(
            `"Apple Inc" OR Microsoft`,
        ),
        PerEntityName: v505.String(
            `"Elon Musk" OR "Jeff Bezos"`,
        ),
        LocEntityName: v505.String(
            `"San Francisco" OR "New York City"`,
        ),
        MiscEntityName: v505.String(
            `AWS OR "Microsoft Azure"`,
        ),
        TitleSentimentMin: v505.Float64(
            -0.5,
        ),
        TitleSentimentMax: v505.Float64(
            0.5,
        ),
        ContentSentimentMin: v505.Float64(
            -0.5,
        ),
        ContentSentimentMax: v505.Float64(
            0.5,
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
            "Tag1,Tag2",
        ),
        ExcludeDuplicates: v505.Bool(
            true,
        ),
        RobotsCompliant: v505.Bool(
            true,
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

Format: start with the word `top`, followed by the number of desired sources, and then the two-letter country code [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). 

Multiple countries with the number of top sources can be specified as a comma-separated string.
    
</dd>
</dl>

<dl>
<dd>

**sourceName:** `*string` 

Word or phrase to search within the source names. To specify multiple values, use a comma-separated string.

**Note**: The search doesn't require an exact match and returns sources containing the specified terms in their names. You can use any word or phrase, like `"sport"` or `"new york times"`. For example, `"sport"` returns sources such as `"Motorsport"`, `"Dot Esport"`, and `"Tuttosport"`.
    
</dd>
</dl>

<dl>
<dd>

**sources:** `*string` — One or more news sources to narrow down the search. The format must be a domain URL. Subdomains, such as `finance.yahoo.com`, are also acceptable.To specify multiple sources, use a comma-separated string.
    
</dd>
</dl>

<dl>
<dd>

**notSources:** `*string` — The news sources to exclude from the search. To exclude multiple sources, use a comma-separated string.
    
</dd>
</dl>

<dl>
<dd>

**lang:** `*string` 

The language(s) of the search. The only accepted format is the two-letter [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code. To select multiple languages, use a comma-separated string.

To learn more, see [Enumerated parameters > Language](https://www.newscatcherapi.com/docs/news-api/api-reference/enumerated-parameters#language-lang-and-not-lang).
    
</dd>
</dl>

<dl>
<dd>

**notLang:** `*string` 

The language(s) to exclude from the search. The accepted format is the two-letter [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code. To exclude multiple languages, use a comma-separated string. 

To learn more, see [Enumerated parameters > Language](https://www.newscatcherapi.com/docs/news-api/api-reference/enumerated-parameters#language-lang-and-not-lang).
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*string` 

The countries where the news publisher is located. The accepted format is the two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. To select multiple countries, use a comma-separated string.

To learn more, see [Enumerated parameters > Country](https://www.newscatcherapi.com/docs/news-api/api-reference/enumerated-parameters#country-country-and-not-country).
    
</dd>
</dl>

<dl>
<dd>

**notCountries:** `*string` 

The publisher location countries to exclude from the search. The accepted format is the two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. To exclude multiple countries, use a comma-separated string. 

To learn more, see [Enumerated parameters > Country](https://www.newscatcherapi.com/docs/news-api/api-reference/enumerated-parameters#country-country-and-not-country).
    
</dd>
</dl>

<dl>
<dd>

**notAuthorName:** `*string` — The list of author names to exclude from your search. To exclude articles by specific authors, use a comma-separated string.
    
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

**parentUrl:** `*string` — The categorical URL(s) to filter your search. To filter your search by multiple categorical URLs, use a comma-separated string.
    
</dd>
</dl>

<dl>
<dd>

**allLinks:** `*string` 

The complete URL(s) mentioned in the article. For multiple URLs, use a comma-separated string.

For more details, see [Search by URL](https://www.newscatcherapi.com/docs/news-api/how-to/search-by-url).
    
</dd>
</dl>

<dl>
<dd>

**allDomainLinks:** `*string` 

The domain(s) mentioned in the article. For multiple domains, use a comma-separated string.

For more details, see [Search by URL](https://www.newscatcherapi.com/docs/news-api/how-to/search-by-url).
    
</dd>
</dl>

<dl>
<dd>

**allLinksText:** `*string` 

The text content of links mentioned in the article. Searches for links where the anchor text contains the specified terms. For multiple terms, use a comma-separated string.

**Note**: When this parameter is used, the response includes the `all_links_data` field with detailed link information.

To learn more, see [Search by URL](https://www.newscatcherapi.com/docs/news-api/how-to/search-by-url).
    
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

**newsType:** `*string` 

Filters results based on the news type. Multiple types can be specified using a comma-separated string.

For a complete list of available news types, see [Enumerated parameters > News type](https://www.newscatcherapi.com/docs/news-api/api-reference/enumerated-parameters#news-type-news-type).
    
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

**iptcTags:** `*string` 

Filters articles based on International Press Telecommunications Council (IPTC) media topic tags. To specify multiple IPTC tags, use a comma-separated string of tag IDs. 

**Note**: The `iptc_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see [IPTC Media Topic NewsCodes](https://www.iptc.org/std/NewsCodes/treeview/mediatopic/mediatopic-en-GB.html).
    
</dd>
</dl>

<dl>
<dd>

**notIptcTags:** `*string` 

Inverse of the `iptc_tags` parameter. Excludes articles based on International Press Telecommunications Council (IPTC) media topic tags. To specify multiple IPTC tags to exclude, use a comma-separated string of tag IDs.

**Note**: The `not_iptc_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see [IPTC Media Topic NewsCodes](https://www.iptc.org/std/NewsCodes/treeview/mediatopic/mediatopic-en-GB.html).
    
</dd>
</dl>

<dl>
<dd>

**iabTags:** `*string` 

Filters articles based on Interactive Advertising Bureau (IAB) content categories. These tags provide a standardized taxonomy for digital advertising content categorization. To specify multiple IAB categories, use a comma-separated string.

**Note**: The `iab_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see the [IAB Content taxonomy](https://iabtechlab.com/standards/content-taxonomy/).
    
</dd>
</dl>

<dl>
<dd>

**notIabTags:** `*string` 

Inverse of the `iab_tags` parameter. Excludes articles based on Interactive Advertising Bureau (IAB) content categories. These tags provide a standardized taxonomy for digital advertising content categorization. To specify multiple IAB categories to exclude, use a comma-separated string.

**Note**: The `not_iab_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see the [IAB Content taxonomy](https://iabtechlab.com/standards/content-taxonomy/).
    
</dd>
</dl>

<dl>
<dd>

**customTags:** `*string` 

Filters articles based on provided taxonomy that is tailored to your specific needs and is accessible only with your API key. To specify tags, use the following pattern: 

- `custom_tags.taxonomy=Tag1,Tag2`, where `taxonomy` is the taxonomy name and `Tag1,Tag2` is a comma-separated list of tag names.

Example: `custom_tags.industry="Manufacturing,Logistics"`

To learn more, see the [Custom tags](https://www.newscatcherapi.com/docs/news-api/guides-and-concepts/custom-tags).
    
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

<details><summary><code>client.Search.Post(request) -> *v505.SearchPostResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Searches for articles based on specified criteria such as keywords, language, country, source, and more.
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

**allLinksText:** `*v505.AllLinksText` 
    
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
<details><summary><code>client.LatestHeadlines.LatestHeadlinesGet() -> *v505.LatestHeadlinesGetResponse</code></summary>
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
        ByParseDate: v505.Bool(
            true,
        ),
        Lang: v505.String(
            "en,es",
        ),
        NotLang: v505.String(
            "fr,de",
        ),
        Countries: v505.String(
            "US,CA",
        ),
        NotCountries: v505.String(
            "UK,FR",
        ),
        PredefinedSources: v505.String(
            "top 50 US, top 20 GB",
        ),
        Sources: v505.String(
            "nytimes.com,finance.yahoo.com",
        ),
        NotSources: v505.String(
            "cnn.com,wsj.com",
        ),
        NotAuthorName: v505.String(
            "John Doe, Jane Doe",
        ),
        RankedOnly: v505.Bool(
            true,
        ),
        IsHeadline: v505.Bool(
            true,
        ),
        IsOpinion: v505.Bool(
            true,
        ),
        IsPaidContent: v505.Bool(
            false,
        ),
        ParentUrl: v505.String(
            "wsj.com/politics,wsj.com/tech",
        ),
        AllLinks: v505.String(
            "https://aiindex.stanford.edu/report,https://www.stateof.ai",
        ),
        AllDomainLinks: v505.String(
            "who.int,nih.gov",
        ),
        AllLinksText: v505.String(
            "Nvidia,Tesla",
        ),
        WordCountMin: v505.Int(
            300,
        ),
        WordCountMax: v505.Int(
            1000,
        ),
        Page: v505.Int(
            2,
        ),
        PageSize: v505.Int(
            50,
        ),
        ClusteringEnabled: v505.Bool(
            true,
        ),
        ClusteringThreshold: v505.Float64(
            0.6,
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
            "Finance,Tech",
        ),
        NotTheme: v505.String(
            "Crime,Sports",
        ),
        OrgEntityName: v505.String(
            `"Apple Inc" OR Microsoft`,
        ),
        PerEntityName: v505.String(
            `"Elon Musk" OR "Jeff Bezos"`,
        ),
        LocEntityName: v505.String(
            `"San Francisco" OR "New York City"`,
        ),
        MiscEntityName: v505.String(
            `AWS OR "Microsoft Azure"`,
        ),
        TitleSentimentMin: v505.Float64(
            -0.5,
        ),
        TitleSentimentMax: v505.Float64(
            0.5,
        ),
        ContentSentimentMin: v505.Float64(
            -0.5,
        ),
        ContentSentimentMax: v505.Float64(
            0.5,
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
            "Tag1,Tag2",
        ),
        RobotsCompliant: v505.Bool(
            true,
        ),
    }
client.LatestHeadlines.LatestHeadlinesGet(
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

**sortBy:** `*v505.SortBy` 
    
</dd>
</dl>

<dl>
<dd>

**lang:** `*string` 

The language(s) of the search. The only accepted format is the two-letter [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code. To select multiple languages, use a comma-separated string.

To learn more, see [Enumerated parameters > Language](https://www.newscatcherapi.com/docs/news-api/api-reference/enumerated-parameters#language-lang-and-not-lang).
    
</dd>
</dl>

<dl>
<dd>

**notLang:** `*string` 

The language(s) to exclude from the search. The accepted format is the two-letter [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code. To exclude multiple languages, use a comma-separated string. 

To learn more, see [Enumerated parameters > Language](https://www.newscatcherapi.com/docs/news-api/api-reference/enumerated-parameters#language-lang-and-not-lang).
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*string` 

The countries where the news publisher is located. The accepted format is the two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. To select multiple countries, use a comma-separated string.

To learn more, see [Enumerated parameters > Country](https://www.newscatcherapi.com/docs/news-api/api-reference/enumerated-parameters#country-country-and-not-country).
    
</dd>
</dl>

<dl>
<dd>

**notCountries:** `*string` 

The publisher location countries to exclude from the search. The accepted format is the two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. To exclude multiple countries, use a comma-separated string. 

To learn more, see [Enumerated parameters > Country](https://www.newscatcherapi.com/docs/news-api/api-reference/enumerated-parameters#country-country-and-not-country).
    
</dd>
</dl>

<dl>
<dd>

**predefinedSources:** `*string` 

Predefined top news sources per country. 

Format: start with the word `top`, followed by the number of desired sources, and then the two-letter country code [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). 

Multiple countries with the number of top sources can be specified as a comma-separated string.
    
</dd>
</dl>

<dl>
<dd>

**sources:** `*string` — One or more news sources to narrow down the search. The format must be a domain URL. Subdomains, such as `finance.yahoo.com`, are also acceptable.To specify multiple sources, use a comma-separated string.
    
</dd>
</dl>

<dl>
<dd>

**notSources:** `*string` — The news sources to exclude from the search. To exclude multiple sources, use a comma-separated string.
    
</dd>
</dl>

<dl>
<dd>

**notAuthorName:** `*string` — The list of author names to exclude from your search. To exclude articles by specific authors, use a comma-separated string.
    
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

**parentUrl:** `*string` — The categorical URL(s) to filter your search. To filter your search by multiple categorical URLs, use a comma-separated string.
    
</dd>
</dl>

<dl>
<dd>

**allLinks:** `*string` 

The complete URL(s) mentioned in the article. For multiple URLs, use a comma-separated string.

For more details, see [Search by URL](https://www.newscatcherapi.com/docs/news-api/how-to/search-by-url).
    
</dd>
</dl>

<dl>
<dd>

**allDomainLinks:** `*string` 

The domain(s) mentioned in the article. For multiple domains, use a comma-separated string.

For more details, see [Search by URL](https://www.newscatcherapi.com/docs/news-api/how-to/search-by-url).
    
</dd>
</dl>

<dl>
<dd>

**allLinksText:** `*string` 

The text content of links mentioned in the article. Searches for links where the anchor text contains the specified terms. For multiple terms, use a comma-separated string.

**Note**: When this parameter is used, the response includes the `all_links_data` field with detailed link information.

To learn more, see [Search by URL](https://www.newscatcherapi.com/docs/news-api/how-to/search-by-url).
    
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

**iptcTags:** `*string` 

Filters articles based on International Press Telecommunications Council (IPTC) media topic tags. To specify multiple IPTC tags, use a comma-separated string of tag IDs. 

**Note**: The `iptc_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see [IPTC Media Topic NewsCodes](https://www.iptc.org/std/NewsCodes/treeview/mediatopic/mediatopic-en-GB.html).
    
</dd>
</dl>

<dl>
<dd>

**notIptcTags:** `*string` 

Inverse of the `iptc_tags` parameter. Excludes articles based on International Press Telecommunications Council (IPTC) media topic tags. To specify multiple IPTC tags to exclude, use a comma-separated string of tag IDs.

**Note**: The `not_iptc_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see [IPTC Media Topic NewsCodes](https://www.iptc.org/std/NewsCodes/treeview/mediatopic/mediatopic-en-GB.html).
    
</dd>
</dl>

<dl>
<dd>

**iabTags:** `*string` 

Filters articles based on Interactive Advertising Bureau (IAB) content categories. These tags provide a standardized taxonomy for digital advertising content categorization. To specify multiple IAB categories, use a comma-separated string.

**Note**: The `iab_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see the [IAB Content taxonomy](https://iabtechlab.com/standards/content-taxonomy/).
    
</dd>
</dl>

<dl>
<dd>

**notIabTags:** `*string` 

Inverse of the `iab_tags` parameter. Excludes articles based on Interactive Advertising Bureau (IAB) content categories. These tags provide a standardized taxonomy for digital advertising content categorization. To specify multiple IAB categories to exclude, use a comma-separated string.

**Note**: The `not_iab_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see the [IAB Content taxonomy](https://iabtechlab.com/standards/content-taxonomy/).
    
</dd>
</dl>

<dl>
<dd>

**customTags:** `*string` 

Filters articles based on provided taxonomy that is tailored to your specific needs and is accessible only with your API key. To specify tags, use the following pattern: 

- `custom_tags.taxonomy=Tag1,Tag2`, where `taxonomy` is the taxonomy name and `Tag1,Tag2` is a comma-separated list of tag names.

Example: `custom_tags.industry="Manufacturing,Logistics"`

To learn more, see the [Custom tags](https://www.newscatcherapi.com/docs/news-api/guides-and-concepts/custom-tags).
    
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

<details><summary><code>client.LatestHeadlines.LatestHeadlinesPost(request) -> *v505.LatestHeadlinesPostResponse</code></summary>
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
client.LatestHeadlines.LatestHeadlinesPost(
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

**allLinksText:** `*v505.AllLinksText` 
    
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

## BreakingNews
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
        RankedOnly: v505.Bool(
            true,
        ),
        FromRank: v505.Int(
            100,
        ),
        ToRank: v505.Int(
            100,
        ),
        Page: v505.Int(
            2,
        ),
        PageSize: v505.Int(
            50,
        ),
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
            "Finance,Tech",
        ),
        NotTheme: v505.String(
            "Crime,Sports",
        ),
        OrgEntityName: v505.String(
            `"Apple Inc" OR Microsoft`,
        ),
        PerEntityName: v505.String(
            `"Elon Musk" OR "Jeff Bezos"`,
        ),
        LocEntityName: v505.String(
            `"San Francisco" OR "New York City"`,
        ),
        MiscEntityName: v505.String(
            `AWS OR "Microsoft Azure"`,
        ),
        TitleSentimentMin: v505.Float64(
            -0.5,
        ),
        TitleSentimentMax: v505.Float64(
            0.5,
        ),
        ContentSentimentMin: v505.Float64(
            -0.5,
        ),
        ContentSentimentMax: v505.Float64(
            0.5,
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

**contentSentimentMax:** `*v505.ContentSentimentMax` 
    
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

**contentSentimentMax:** `*v505.ContentSentimentMax` 
    
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
            "John Doe, Jane Doe",
        ),
        PredefinedSources: v505.String(
            "top 50 US, top 20 GB",
        ),
        Sources: v505.String(
            "nytimes.com,finance.yahoo.com",
        ),
        NotSources: v505.String(
            "cnn.com,wsj.com",
        ),
        Lang: v505.String(
            "en,es",
        ),
        NotLang: v505.String(
            "fr,de",
        ),
        Countries: v505.String(
            "US,CA",
        ),
        NotCountries: v505.String(
            "UK,FR",
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
        PublishedDatePrecision: v505.String(
            "full",
        ),
        ByParseDate: v505.Bool(
            true,
        ),
        RankedOnly: v505.Bool(
            true,
        ),
        FromRank: v505.Int(
            100,
        ),
        ToRank: v505.Int(
            100,
        ),
        IsHeadline: v505.Bool(
            true,
        ),
        IsOpinion: v505.Bool(
            true,
        ),
        IsPaidContent: v505.Bool(
            false,
        ),
        ParentUrl: v505.String(
            "wsj.com/politics,wsj.com/tech",
        ),
        AllLinks: v505.String(
            "https://aiindex.stanford.edu/report,https://www.stateof.ai",
        ),
        AllDomainLinks: v505.String(
            "who.int,nih.gov",
        ),
        AllLinksText: v505.String(
            "Nvidia,Tesla",
        ),
        WordCountMin: v505.Int(
            300,
        ),
        WordCountMax: v505.Int(
            1000,
        ),
        Page: v505.Int(
            2,
        ),
        PageSize: v505.Int(
            50,
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
            "Finance,Tech",
        ),
        NotTheme: v505.String(
            "Crime,Sports",
        ),
        NerName: v505.String(
            "Tesla,Amazon",
        ),
        TitleSentimentMin: v505.Float64(
            -0.5,
        ),
        TitleSentimentMax: v505.Float64(
            0.5,
        ),
        ContentSentimentMin: v505.Float64(
            -0.5,
        ),
        ContentSentimentMax: v505.Float64(
            0.5,
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
            "Tag1,Tag2",
        ),
        RobotsCompliant: v505.Bool(
            true,
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

**authorName:** `v505.AuthorName` 
    
</dd>
</dl>

<dl>
<dd>

**notAuthorName:** `*string` — The list of author names to exclude from your search. To exclude articles by specific authors, use a comma-separated string.
    
</dd>
</dl>

<dl>
<dd>

**predefinedSources:** `*string` 

Predefined top news sources per country. 

Format: start with the word `top`, followed by the number of desired sources, and then the two-letter country code [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). 

Multiple countries with the number of top sources can be specified as a comma-separated string.
    
</dd>
</dl>

<dl>
<dd>

**sources:** `*string` — One or more news sources to narrow down the search. The format must be a domain URL. Subdomains, such as `finance.yahoo.com`, are also acceptable.To specify multiple sources, use a comma-separated string.
    
</dd>
</dl>

<dl>
<dd>

**notSources:** `*string` — The news sources to exclude from the search. To exclude multiple sources, use a comma-separated string.
    
</dd>
</dl>

<dl>
<dd>

**lang:** `*string` 

The language(s) of the search. The only accepted format is the two-letter [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code. To select multiple languages, use a comma-separated string.

To learn more, see [Enumerated parameters > Language](https://www.newscatcherapi.com/docs/news-api/api-reference/enumerated-parameters#language-lang-and-not-lang).
    
</dd>
</dl>

<dl>
<dd>

**notLang:** `*string` 

The language(s) to exclude from the search. The accepted format is the two-letter [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code. To exclude multiple languages, use a comma-separated string. 

To learn more, see [Enumerated parameters > Language](https://www.newscatcherapi.com/docs/news-api/api-reference/enumerated-parameters#language-lang-and-not-lang).
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*string` 

The countries where the news publisher is located. The accepted format is the two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. To select multiple countries, use a comma-separated string.

To learn more, see [Enumerated parameters > Country](https://www.newscatcherapi.com/docs/news-api/api-reference/enumerated-parameters#country-country-and-not-country).
    
</dd>
</dl>

<dl>
<dd>

**notCountries:** `*string` 

The publisher location countries to exclude from the search. The accepted format is the two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. To exclude multiple countries, use a comma-separated string. 

To learn more, see [Enumerated parameters > Country](https://www.newscatcherapi.com/docs/news-api/api-reference/enumerated-parameters#country-country-and-not-country).
    
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

**parentUrl:** `*string` — The categorical URL(s) to filter your search. To filter your search by multiple categorical URLs, use a comma-separated string.
    
</dd>
</dl>

<dl>
<dd>

**allLinks:** `*string` 

The complete URL(s) mentioned in the article. For multiple URLs, use a comma-separated string.

For more details, see [Search by URL](https://www.newscatcherapi.com/docs/news-api/how-to/search-by-url).
    
</dd>
</dl>

<dl>
<dd>

**allDomainLinks:** `*string` 

The domain(s) mentioned in the article. For multiple domains, use a comma-separated string.

For more details, see [Search by URL](https://www.newscatcherapi.com/docs/news-api/how-to/search-by-url).
    
</dd>
</dl>

<dl>
<dd>

**allLinksText:** `*string` 

The text content of links mentioned in the article. Searches for links where the anchor text contains the specified terms. For multiple terms, use a comma-separated string.

**Note**: When this parameter is used, the response includes the `all_links_data` field with detailed link information.

To learn more, see [Search by URL](https://www.newscatcherapi.com/docs/news-api/how-to/search-by-url).
    
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

**iptcTags:** `*string` 

Filters articles based on International Press Telecommunications Council (IPTC) media topic tags. To specify multiple IPTC tags, use a comma-separated string of tag IDs. 

**Note**: The `iptc_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see [IPTC Media Topic NewsCodes](https://www.iptc.org/std/NewsCodes/treeview/mediatopic/mediatopic-en-GB.html).
    
</dd>
</dl>

<dl>
<dd>

**notIptcTags:** `*string` 

Inverse of the `iptc_tags` parameter. Excludes articles based on International Press Telecommunications Council (IPTC) media topic tags. To specify multiple IPTC tags to exclude, use a comma-separated string of tag IDs.

**Note**: The `not_iptc_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see [IPTC Media Topic NewsCodes](https://www.iptc.org/std/NewsCodes/treeview/mediatopic/mediatopic-en-GB.html).
    
</dd>
</dl>

<dl>
<dd>

**iabTags:** `*string` 

Filters articles based on Interactive Advertising Bureau (IAB) content categories. These tags provide a standardized taxonomy for digital advertising content categorization. To specify multiple IAB categories, use a comma-separated string.

**Note**: The `iab_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see the [IAB Content taxonomy](https://iabtechlab.com/standards/content-taxonomy/).
    
</dd>
</dl>

<dl>
<dd>

**notIabTags:** `*string` 

Inverse of the `iab_tags` parameter. Excludes articles based on Interactive Advertising Bureau (IAB) content categories. These tags provide a standardized taxonomy for digital advertising content categorization. To specify multiple IAB categories to exclude, use a comma-separated string.

**Note**: The `not_iab_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see the [IAB Content taxonomy](https://iabtechlab.com/standards/content-taxonomy/).
    
</dd>
</dl>

<dl>
<dd>

**customTags:** `*string` 

Filters articles based on provided taxonomy that is tailored to your specific needs and is accessible only with your API key. To specify tags, use the following pattern: 

- `custom_tags.taxonomy=Tag1,Tag2`, where `taxonomy` is the taxonomy name and `Tag1,Tag2` is a comma-separated list of tag names.

Example: `custom_tags.industry="Manufacturing,Logistics"`

To learn more, see the [Custom tags](https://www.newscatcherapi.com/docs/news-api/guides-and-concepts/custom-tags).
    
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

**allLinksText:** `*v505.AllLinksText` 
    
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

## SearchByLink
<details><summary><code>client.SearchByLink.SearchByLinkGet() -> *v505.SearchResponseDto</code></summary>
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
request := &v505.SearchByLinkGetRequest{
        Ids: v505.String(
            "5f8d0d55b6e45e00179c6e7e",
        ),
        Links: v505.String(
            "https://nytimes.com/article1,https://bbc.com/article2",
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
        Page: v505.Int(
            2,
        ),
        PageSize: v505.Int(
            50,
        ),
        RobotsCompliant: v505.Bool(
            true,
        ),
    }
client.SearchByLink.SearchByLinkGet(
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

**Caution**: You can use either the `links` or the `ids` parameter, but not both at the same time.
    
</dd>
</dl>

<dl>
<dd>

**links:** `*string` 

The article link or list of article links to search for. To specify multiple links, use a comma-separated string.

**Caution**: You can use either the `links` or the `ids` parameter, but not both at the same time.
    
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

<details><summary><code>client.SearchByLink.SearchByLinkPost(request) -> *v505.SearchResponseDto</code></summary>
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
request := &v505.SearchByLinkPostRequest{
        Links: &v505.Links{
            String: "https://www.reuters.com/business/energy/oil-prices-up-after-israeli-attacks-oversupply-caps-gains-2025-09-10/",
        },
    }
client.SearchByLink.SearchByLinkPost(
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
            "en,es",
        ),
        Countries: v505.String(
            "US,CA",
        ),
        PredefinedSources: v505.String(
            "top 50 US, top 20 GB",
        ),
        SourceName: v505.String(
            "sport,tech",
        ),
        SourceUrl: v505.String(
            "bbc.com",
        ),
        IncludeAdditionalInfo: v505.Bool(
            true,
        ),
        IsNewsDomain: v505.Bool(
            true,
        ),
        NewsType: v505.String(
            "General News Outlets,Tech News and Updates",
        ),
        FromRank: v505.Int(
            100,
        ),
        ToRank: v505.Int(
            100,
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

To learn more, see [Enumerated parameters > Language](https://www.newscatcherapi.com/docs/news-api/api-reference/enumerated-parameters#language-lang-and-not-lang).
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*string` 

The countries where the news publisher is located. The accepted format is the two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. To select multiple countries, use a comma-separated string.

To learn more, see [Enumerated parameters > Country](https://www.newscatcherapi.com/docs/news-api/api-reference/enumerated-parameters#country-country-and-not-country).
    
</dd>
</dl>

<dl>
<dd>

**predefinedSources:** `*string` 

Predefined top news sources per country. 

Format: start with the word `top`, followed by the number of desired sources, and then the two-letter country code [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). 

Multiple countries with the number of top sources can be specified as a comma-separated string.
    
</dd>
</dl>

<dl>
<dd>

**sourceName:** `*string` 

Word or phrase to search within the source names. To specify multiple values, use a comma-separated string.

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

**newsType:** `*string` 

Filters results based on the news type. Multiple types can be specified using a comma-separated string.

For a complete list of available news types, see [Enumerated parameters > News type](https://www.newscatcherapi.com/docs/news-api/api-reference/enumerated-parameters#news-type-news-type).
    
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
<details><summary><code>client.Aggregation.CountGet() -> *v505.AggregationCountGetResponse</code></summary>
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
request := &v505.AggregationCountGetRequest{
        Q: `"supply chain" AND Amazon NOT China`,
        SearchIn: v505.String(
            "title_content, title_content_translated",
        ),
        PredefinedSources: v505.String(
            "top 50 US, top 20 GB",
        ),
        Sources: v505.String(
            "nytimes.com,finance.yahoo.com",
        ),
        NotSources: v505.String(
            "cnn.com,wsj.com",
        ),
        Lang: v505.String(
            "en,es",
        ),
        NotLang: v505.String(
            "fr,de",
        ),
        Countries: v505.String(
            "US,CA",
        ),
        NotCountries: v505.String(
            "UK,FR",
        ),
        NotAuthorName: v505.String(
            "John Doe, Jane Doe",
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
        PublishedDatePrecision: v505.String(
            "full",
        ),
        ByParseDate: v505.Bool(
            true,
        ),
        RankedOnly: v505.Bool(
            true,
        ),
        FromRank: v505.Int(
            100,
        ),
        ToRank: v505.Int(
            100,
        ),
        IsHeadline: v505.Bool(
            true,
        ),
        IsOpinion: v505.Bool(
            true,
        ),
        IsPaidContent: v505.Bool(
            false,
        ),
        ParentUrl: v505.String(
            "wsj.com/politics,wsj.com/tech",
        ),
        AllLinks: v505.String(
            "https://aiindex.stanford.edu/report,https://www.stateof.ai",
        ),
        AllDomainLinks: v505.String(
            "who.int,nih.gov",
        ),
        AllLinksText: v505.String(
            "Nvidia,Tesla",
        ),
        WordCountMin: v505.Int(
            300,
        ),
        WordCountMax: v505.Int(
            1000,
        ),
        Page: v505.Int(
            2,
        ),
        PageSize: v505.Int(
            50,
        ),
        IncludeNlpData: v505.Bool(
            true,
        ),
        HasNlp: v505.Bool(
            true,
        ),
        Theme: v505.String(
            "Finance,Tech",
        ),
        NotTheme: v505.String(
            "Crime,Sports",
        ),
        OrgEntityName: v505.String(
            `"Apple Inc" OR Microsoft`,
        ),
        PerEntityName: v505.String(
            `"Elon Musk" OR "Jeff Bezos"`,
        ),
        LocEntityName: v505.String(
            `"San Francisco" OR "New York City"`,
        ),
        MiscEntityName: v505.String(
            `AWS OR "Microsoft Azure"`,
        ),
        TitleSentimentMin: v505.Float64(
            -0.5,
        ),
        TitleSentimentMax: v505.Float64(
            0.5,
        ),
        ContentSentimentMin: v505.Float64(
            -0.5,
        ),
        ContentSentimentMax: v505.Float64(
            0.5,
        ),
        IptcTags: v505.String(
            "20000199,20000209",
        ),
        NotIptcTags: v505.String(
            "20000205,20000209",
        ),
        RobotsCompliant: v505.Bool(
            true,
        ),
    }
client.Aggregation.CountGet(
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

Format: start with the word `top`, followed by the number of desired sources, and then the two-letter country code [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2). 

Multiple countries with the number of top sources can be specified as a comma-separated string.
    
</dd>
</dl>

<dl>
<dd>

**sources:** `*string` — One or more news sources to narrow down the search. The format must be a domain URL. Subdomains, such as `finance.yahoo.com`, are also acceptable.To specify multiple sources, use a comma-separated string.
    
</dd>
</dl>

<dl>
<dd>

**notSources:** `*string` — The news sources to exclude from the search. To exclude multiple sources, use a comma-separated string.
    
</dd>
</dl>

<dl>
<dd>

**lang:** `*string` 

The language(s) of the search. The only accepted format is the two-letter [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code. To select multiple languages, use a comma-separated string.

To learn more, see [Enumerated parameters > Language](https://www.newscatcherapi.com/docs/news-api/api-reference/enumerated-parameters#language-lang-and-not-lang).
    
</dd>
</dl>

<dl>
<dd>

**notLang:** `*string` 

The language(s) to exclude from the search. The accepted format is the two-letter [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) code. To exclude multiple languages, use a comma-separated string. 

To learn more, see [Enumerated parameters > Language](https://www.newscatcherapi.com/docs/news-api/api-reference/enumerated-parameters#language-lang-and-not-lang).
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*string` 

The countries where the news publisher is located. The accepted format is the two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. To select multiple countries, use a comma-separated string.

To learn more, see [Enumerated parameters > Country](https://www.newscatcherapi.com/docs/news-api/api-reference/enumerated-parameters#country-country-and-not-country).
    
</dd>
</dl>

<dl>
<dd>

**notCountries:** `*string` 

The publisher location countries to exclude from the search. The accepted format is the two-letter [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code. To exclude multiple countries, use a comma-separated string. 

To learn more, see [Enumerated parameters > Country](https://www.newscatcherapi.com/docs/news-api/api-reference/enumerated-parameters#country-country-and-not-country).
    
</dd>
</dl>

<dl>
<dd>

**notAuthorName:** `*string` — The list of author names to exclude from your search. To exclude articles by specific authors, use a comma-separated string.
    
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

**parentUrl:** `*string` — The categorical URL(s) to filter your search. To filter your search by multiple categorical URLs, use a comma-separated string.
    
</dd>
</dl>

<dl>
<dd>

**allLinks:** `*string` 

The complete URL(s) mentioned in the article. For multiple URLs, use a comma-separated string.

For more details, see [Search by URL](https://www.newscatcherapi.com/docs/news-api/how-to/search-by-url).
    
</dd>
</dl>

<dl>
<dd>

**allDomainLinks:** `*string` 

The domain(s) mentioned in the article. For multiple domains, use a comma-separated string.

For more details, see [Search by URL](https://www.newscatcherapi.com/docs/news-api/how-to/search-by-url).
    
</dd>
</dl>

<dl>
<dd>

**allLinksText:** `*string` 

The text content of links mentioned in the article. Searches for links where the anchor text contains the specified terms. For multiple terms, use a comma-separated string.

**Note**: When this parameter is used, the response includes the `all_links_data` field with detailed link information.

To learn more, see [Search by URL](https://www.newscatcherapi.com/docs/news-api/how-to/search-by-url).
    
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

**iptcTags:** `*string` 

Filters articles based on International Press Telecommunications Council (IPTC) media topic tags. To specify multiple IPTC tags, use a comma-separated string of tag IDs. 

**Note**: The `iptc_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see [IPTC Media Topic NewsCodes](https://www.iptc.org/std/NewsCodes/treeview/mediatopic/mediatopic-en-GB.html).
    
</dd>
</dl>

<dl>
<dd>

**notIptcTags:** `*string` 

Inverse of the `iptc_tags` parameter. Excludes articles based on International Press Telecommunications Council (IPTC) media topic tags. To specify multiple IPTC tags to exclude, use a comma-separated string of tag IDs.

**Note**: The `not_iptc_tags` parameter is only available in the `v3_nlp_iptc_tags` subscription plan.

To learn more, see [IPTC Media Topic NewsCodes](https://www.iptc.org/std/NewsCodes/treeview/mediatopic/mediatopic-en-GB.html).
    
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

<details><summary><code>client.Aggregation.CountPost(request) -> *v505.AggregationCountPostResponse</code></summary>
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
request := &v505.AggregationCountPostRequest{
        Q: `"supply chain" AND Amazon NOT China`,
        AggregationBy: v505.AggregationByDay.Ptr(),
    }
client.Aggregation.CountPost(
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

**allLinksText:** `*v505.AllLinksText` 
    
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

