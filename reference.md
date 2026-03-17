# Reference
## Search
<details><summary><code>client.Search.Get() -> *v2.GetSearchResponse</code></summary>
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
request := &v2.GetSearchRequest{
        Q: `"supply chain" AND Amazon NOT China`,
        SearchIn: v2.String(
            "title_content, title_content_translated",
        ),
        IncludeTranslationFields: v2.Bool(
            true,
        ),
        PredefinedSources: v2.String(
            "top 50 US, top 20 GB",
        ),
        SourceName: v2.String(
            "sport,tech",
        ),
        Sources: v2.String(
            "nytimes.com,finance.yahoo.com",
        ),
        NotSources: v2.String(
            "cnn.com,wsj.com",
        ),
        Lang: v2.String(
            "en,es",
        ),
        NotLang: v2.String(
            "fr,de",
        ),
        Countries: v2.String(
            "US,CA",
        ),
        NotCountries: v2.String(
            "UK,FR",
        ),
        NotAuthorName: v2.String(
            "John Doe, Jane Doe",
        ),
        From: &v2.From{
            DateTime: v2.MustParseDateTime(
                "2024-07-01T00:00:00Z",
            ),
        },
        To: &v2.To{
            DateTime: v2.MustParseDateTime(
                "2024-01-01T00:00:00Z",
            ),
        },
        PublishedDatePrecision: v2.String(
            "full",
        ),
        ByParseDate: v2.Bool(
            true,
        ),
        RankedOnly: v2.Bool(
            true,
        ),
        FromRank: v2.Int(
            100,
        ),
        ToRank: v2.Int(
            100,
        ),
        IsHeadline: v2.Bool(
            true,
        ),
        IsOpinion: v2.Bool(
            true,
        ),
        IsPaidContent: v2.Bool(
            false,
        ),
        ParentUrl: v2.String(
            "wsj.com/politics,wsj.com/tech",
        ),
        AllLinks: v2.String(
            "https://aiindex.stanford.edu/report,https://www.stateof.ai",
        ),
        AllDomainLinks: v2.String(
            "who.int,nih.gov",
        ),
        AllLinksText: v2.String(
            "Nvidia,Tesla",
        ),
        AdditionalDomainInfo: v2.Bool(
            true,
        ),
        IsNewsDomain: v2.Bool(
            true,
        ),
        NewsType: v2.String(
            "General News Outlets,Tech News and Updates",
        ),
        WordCountMin: v2.Int(
            300,
        ),
        WordCountMax: v2.Int(
            1000,
        ),
        Page: v2.Int(
            2,
        ),
        PageSize: v2.Int(
            50,
        ),
        ClusteringEnabled: v2.Bool(
            true,
        ),
        ClusteringThreshold: v2.Float64(
            0.6,
        ),
        IncludeNlpData: v2.Bool(
            true,
        ),
        HasNlp: v2.Bool(
            true,
        ),
        Theme: v2.String(
            "Finance,Tech",
        ),
        NotTheme: v2.String(
            "Crime,Sports",
        ),
        OrgEntityName: v2.String(
            `"Apple Inc" OR Microsoft`,
        ),
        PerEntityName: v2.String(
            `"Elon Musk" OR "Jeff Bezos"`,
        ),
        LocEntityName: v2.String(
            `"San Francisco" OR "New York City"`,
        ),
        MiscEntityName: v2.String(
            `AWS OR "Microsoft Azure"`,
        ),
        TitleSentimentMin: v2.Float64(
            -0.5,
        ),
        TitleSentimentMax: v2.Float64(
            0.5,
        ),
        ContentSentimentMin: v2.Float64(
            -0.5,
        ),
        ContentSentimentMax: v2.Float64(
            0.5,
        ),
        IptcTags: v2.String(
            "20000199,20000209",
        ),
        NotIptcTags: v2.String(
            "20000205,20000209",
        ),
        IabTags: v2.String(
            "Business,Events",
        ),
        NotIabTags: v2.String(
            "Agriculture,Metals",
        ),
        CustomTags: v2.String(
            "Tag1,Tag2",
        ),
        ExcludeDuplicates: v2.Bool(
            true,
        ),
        RobotsCompliant: v2.Bool(
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

**q:** `v2.Q` 
    
</dd>
</dl>

<dl>
<dd>

**searchIn:** `*v2.SearchIn` 
    
</dd>
</dl>

<dl>
<dd>

**includeTranslationFields:** `*v2.IncludeTranslationFields` 
    
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

**from:** `*v2.From` 
    
</dd>
</dl>

<dl>
<dd>

**to:** `*v2.To` 
    
</dd>
</dl>

<dl>
<dd>

**publishedDatePrecision:** `*v2.PublishedDatePrecision` 
    
</dd>
</dl>

<dl>
<dd>

**byParseDate:** `*v2.ByParseDate` 
    
</dd>
</dl>

<dl>
<dd>

**sortBy:** `*v2.SortBy` 
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*v2.RankedOnly` 
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*v2.FromRank` 
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*v2.ToRank` 
    
</dd>
</dl>

<dl>
<dd>

**isHeadline:** `*v2.IsHeadline` 
    
</dd>
</dl>

<dl>
<dd>

**isOpinion:** `*v2.IsOpinion` 
    
</dd>
</dl>

<dl>
<dd>

**isPaidContent:** `*v2.IsPaidContent` 
    
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

**additionalDomainInfo:** `*v2.AdditionalDomainInfo` 
    
</dd>
</dl>

<dl>
<dd>

**isNewsDomain:** `*v2.IsNewsDomain` 
    
</dd>
</dl>

<dl>
<dd>

**newsDomainType:** `*v2.NewsDomainType` 
    
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

**wordCountMin:** `*v2.WordCountMin` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMax:** `*v2.WordCountMax` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*v2.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*v2.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringEnabled:** `*v2.ClusteringEnabled` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringVariable:** `*v2.ClusteringVariable` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringThreshold:** `*v2.ClusteringThreshold` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*v2.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*v2.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*v2.Theme` 
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*v2.NotTheme` 
    
</dd>
</dl>

<dl>
<dd>

**orgEntityName:** `*v2.OrgEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**perEntityName:** `*v2.PerEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**locEntityName:** `*v2.LocEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**miscEntityName:** `*v2.MiscEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*v2.TitleSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*v2.TitleSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*v2.ContentSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*v2.ContentSentimentMax` 
    
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

**excludeDuplicates:** `*v2.ExcludeDuplicates` 
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*v2.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Search.Post(request) -> *v2.PostSearchResponse</code></summary>
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
request := &v2.PostSearchRequest{
        Q: `"supply chain" AND Amazon NOT China`,
        PageSize: v2.Int(
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

**q:** `v2.Q` 
    
</dd>
</dl>

<dl>
<dd>

**searchIn:** `*v2.SearchIn` 
    
</dd>
</dl>

<dl>
<dd>

**includeTranslationFields:** `*v2.IncludeTranslationFields` 
    
</dd>
</dl>

<dl>
<dd>

**predefinedSources:** `*v2.PredefinedSources` 
    
</dd>
</dl>

<dl>
<dd>

**sourceName:** `*v2.SourceName` 
    
</dd>
</dl>

<dl>
<dd>

**sources:** `*v2.Sources` 
    
</dd>
</dl>

<dl>
<dd>

**notSources:** `*v2.NotSources` 
    
</dd>
</dl>

<dl>
<dd>

**lang:** `*v2.Lang` 
    
</dd>
</dl>

<dl>
<dd>

**notLang:** `*v2.NotLang` 
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*v2.Countries` 
    
</dd>
</dl>

<dl>
<dd>

**notCountries:** `*v2.NotCountries` 
    
</dd>
</dl>

<dl>
<dd>

**notAuthorName:** `*v2.NotAuthorName` 
    
</dd>
</dl>

<dl>
<dd>

**from:** `*v2.From` 
    
</dd>
</dl>

<dl>
<dd>

**to:** `*v2.To` 
    
</dd>
</dl>

<dl>
<dd>

**publishedDatePrecision:** `*v2.PublishedDatePrecision` 
    
</dd>
</dl>

<dl>
<dd>

**byParseDate:** `*v2.ByParseDate` 
    
</dd>
</dl>

<dl>
<dd>

**sortBy:** `*v2.SortBy` 
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*v2.RankedOnly` 
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*v2.FromRank` 
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*v2.ToRank` 
    
</dd>
</dl>

<dl>
<dd>

**isHeadline:** `*v2.IsHeadline` 
    
</dd>
</dl>

<dl>
<dd>

**isOpinion:** `*v2.IsOpinion` 
    
</dd>
</dl>

<dl>
<dd>

**isPaidContent:** `*v2.IsPaidContent` 
    
</dd>
</dl>

<dl>
<dd>

**parentUrl:** `*v2.ParentUrl` 
    
</dd>
</dl>

<dl>
<dd>

**allLinks:** `*v2.AllLinks` 
    
</dd>
</dl>

<dl>
<dd>

**allDomainLinks:** `*v2.AllDomainLinks` 
    
</dd>
</dl>

<dl>
<dd>

**allLinksText:** `*v2.AllLinksText` 
    
</dd>
</dl>

<dl>
<dd>

**additionalDomainInfo:** `*v2.AdditionalDomainInfo` 
    
</dd>
</dl>

<dl>
<dd>

**isNewsDomain:** `*v2.IsNewsDomain` 
    
</dd>
</dl>

<dl>
<dd>

**newsDomainType:** `*v2.NewsDomainType` 
    
</dd>
</dl>

<dl>
<dd>

**newsType:** `*v2.NewsType` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMin:** `*v2.WordCountMin` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMax:** `*v2.WordCountMax` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*v2.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*v2.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringEnabled:** `*v2.ClusteringEnabled` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringVariable:** `*v2.ClusteringVariable` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringThreshold:** `*v2.ClusteringThreshold` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*v2.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*v2.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*v2.Theme` 
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*v2.NotTheme` 
    
</dd>
</dl>

<dl>
<dd>

**orgEntityName:** `*v2.OrgEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**perEntityName:** `*v2.PerEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**locEntityName:** `*v2.LocEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**miscEntityName:** `*v2.MiscEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*v2.TitleSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*v2.TitleSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*v2.ContentSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*v2.ContentSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**iptcTags:** `*v2.IptcTags` 
    
</dd>
</dl>

<dl>
<dd>

**notIptcTags:** `*v2.NotIptcTags` 
    
</dd>
</dl>

<dl>
<dd>

**iabTags:** `*v2.IabTags` 
    
</dd>
</dl>

<dl>
<dd>

**notIabTags:** `*v2.NotIabTags` 
    
</dd>
</dl>

<dl>
<dd>

**customTags:** `*v2.CustomTags` 
    
</dd>
</dl>

<dl>
<dd>

**excludeDuplicates:** `*v2.ExcludeDuplicates` 
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*v2.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## LatestHeadlines
<details><summary><code>client.LatestHeadlines.Get() -> *v2.GetLatestHeadlinesResponse</code></summary>
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
request := &v2.GetLatestHeadlinesRequest{
        When: v2.String(
            "7d",
        ),
        ByParseDate: v2.Bool(
            true,
        ),
        Lang: v2.String(
            "en,es",
        ),
        NotLang: v2.String(
            "fr,de",
        ),
        Countries: v2.String(
            "US,CA",
        ),
        NotCountries: v2.String(
            "UK,FR",
        ),
        PredefinedSources: v2.String(
            "top 50 US, top 20 GB",
        ),
        Sources: v2.String(
            "nytimes.com,finance.yahoo.com",
        ),
        NotSources: v2.String(
            "cnn.com,wsj.com",
        ),
        NotAuthorName: v2.String(
            "John Doe, Jane Doe",
        ),
        RankedOnly: v2.Bool(
            true,
        ),
        IsHeadline: v2.Bool(
            true,
        ),
        IsOpinion: v2.Bool(
            true,
        ),
        IsPaidContent: v2.Bool(
            false,
        ),
        ParentUrl: v2.String(
            "wsj.com/politics,wsj.com/tech",
        ),
        AllLinks: v2.String(
            "https://aiindex.stanford.edu/report,https://www.stateof.ai",
        ),
        AllDomainLinks: v2.String(
            "who.int,nih.gov",
        ),
        AllLinksText: v2.String(
            "Nvidia,Tesla",
        ),
        WordCountMin: v2.Int(
            300,
        ),
        WordCountMax: v2.Int(
            1000,
        ),
        Page: v2.Int(
            2,
        ),
        PageSize: v2.Int(
            50,
        ),
        ClusteringEnabled: v2.Bool(
            true,
        ),
        ClusteringThreshold: v2.Float64(
            0.6,
        ),
        IncludeTranslationFields: v2.Bool(
            true,
        ),
        IncludeNlpData: v2.Bool(
            true,
        ),
        HasNlp: v2.Bool(
            true,
        ),
        Theme: v2.String(
            "Finance,Tech",
        ),
        NotTheme: v2.String(
            "Crime,Sports",
        ),
        OrgEntityName: v2.String(
            `"Apple Inc" OR Microsoft`,
        ),
        PerEntityName: v2.String(
            `"Elon Musk" OR "Jeff Bezos"`,
        ),
        LocEntityName: v2.String(
            `"San Francisco" OR "New York City"`,
        ),
        MiscEntityName: v2.String(
            `AWS OR "Microsoft Azure"`,
        ),
        TitleSentimentMin: v2.Float64(
            -0.5,
        ),
        TitleSentimentMax: v2.Float64(
            0.5,
        ),
        ContentSentimentMin: v2.Float64(
            -0.5,
        ),
        ContentSentimentMax: v2.Float64(
            0.5,
        ),
        IptcTags: v2.String(
            "20000199,20000209",
        ),
        NotIptcTags: v2.String(
            "20000205,20000209",
        ),
        IabTags: v2.String(
            "Business,Events",
        ),
        NotIabTags: v2.String(
            "Agriculture,Metals",
        ),
        CustomTags: v2.String(
            "Tag1,Tag2",
        ),
        RobotsCompliant: v2.Bool(
            true,
        ),
    }
client.LatestHeadlines.Get(
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

**when:** `*v2.When` 
    
</dd>
</dl>

<dl>
<dd>

**byParseDate:** `*v2.ByParseDate` 
    
</dd>
</dl>

<dl>
<dd>

**sortBy:** `*v2.SortBy` 
    
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

**rankedOnly:** `*v2.RankedOnly` 
    
</dd>
</dl>

<dl>
<dd>

**isHeadline:** `*v2.IsHeadline` 
    
</dd>
</dl>

<dl>
<dd>

**isOpinion:** `*v2.IsOpinion` 
    
</dd>
</dl>

<dl>
<dd>

**isPaidContent:** `*v2.IsPaidContent` 
    
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

**wordCountMin:** `*v2.WordCountMin` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMax:** `*v2.WordCountMax` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*v2.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*v2.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringEnabled:** `*v2.ClusteringEnabled` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringVariable:** `*v2.ClusteringVariable` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringThreshold:** `*v2.ClusteringThreshold` 
    
</dd>
</dl>

<dl>
<dd>

**includeTranslationFields:** `*v2.IncludeTranslationFields` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*v2.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*v2.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*v2.Theme` 
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*v2.NotTheme` 
    
</dd>
</dl>

<dl>
<dd>

**orgEntityName:** `*v2.OrgEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**perEntityName:** `*v2.PerEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**locEntityName:** `*v2.LocEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**miscEntityName:** `*v2.MiscEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*v2.TitleSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*v2.TitleSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*v2.ContentSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*v2.ContentSentimentMax` 
    
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

**robotsCompliant:** `*v2.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LatestHeadlines.Post(request) -> *v2.PostLatestHeadlinesResponse</code></summary>
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
request := &v2.PostLatestHeadlinesRequest{
        When: v2.String(
            "7d",
        ),
        PageSize: v2.Int(
            1,
        ),
    }
client.LatestHeadlines.Post(
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

**when:** `*v2.When` 
    
</dd>
</dl>

<dl>
<dd>

**byParseDate:** `*v2.ByParseDate` 
    
</dd>
</dl>

<dl>
<dd>

**lang:** `*v2.Lang` 
    
</dd>
</dl>

<dl>
<dd>

**notLang:** `*v2.NotLang` 
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*v2.Countries` 
    
</dd>
</dl>

<dl>
<dd>

**notCountries:** `*v2.NotCountries` 
    
</dd>
</dl>

<dl>
<dd>

**predefinedSources:** `*v2.PredefinedSources` 
    
</dd>
</dl>

<dl>
<dd>

**sources:** `*v2.Sources` 
    
</dd>
</dl>

<dl>
<dd>

**notSources:** `*v2.NotSources` 
    
</dd>
</dl>

<dl>
<dd>

**notAuthorName:** `*v2.NotAuthorName` 
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*v2.RankedOnly` 
    
</dd>
</dl>

<dl>
<dd>

**isHeadline:** `*v2.IsHeadline` 
    
</dd>
</dl>

<dl>
<dd>

**isOpinion:** `*v2.IsOpinion` 
    
</dd>
</dl>

<dl>
<dd>

**isPaidContent:** `*v2.IsPaidContent` 
    
</dd>
</dl>

<dl>
<dd>

**parentUrl:** `*v2.ParentUrl` 
    
</dd>
</dl>

<dl>
<dd>

**allLinks:** `*v2.AllLinks` 
    
</dd>
</dl>

<dl>
<dd>

**allDomainLinks:** `*v2.AllDomainLinks` 
    
</dd>
</dl>

<dl>
<dd>

**allLinksText:** `*v2.AllLinksText` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMin:** `*v2.WordCountMin` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMax:** `*v2.WordCountMax` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*v2.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*v2.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringEnabled:** `*v2.ClusteringEnabled` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringVariable:** `*v2.ClusteringVariable` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringThreshold:** `*v2.ClusteringThreshold` 
    
</dd>
</dl>

<dl>
<dd>

**includeTranslationFields:** `*v2.IncludeTranslationFields` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*v2.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*v2.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*v2.Theme` 
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*v2.NotTheme` 
    
</dd>
</dl>

<dl>
<dd>

**orgEntityName:** `*v2.OrgEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**perEntityName:** `*v2.PerEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**locEntityName:** `*v2.LocEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**miscEntityName:** `*v2.MiscEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*v2.TitleSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*v2.TitleSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*v2.ContentSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*v2.ContentSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**iptcTags:** `*v2.IptcTags` 
    
</dd>
</dl>

<dl>
<dd>

**notIptcTags:** `*v2.NotIptcTags` 
    
</dd>
</dl>

<dl>
<dd>

**iabTags:** `*v2.IabTags` 
    
</dd>
</dl>

<dl>
<dd>

**notIabTags:** `*v2.NotIabTags` 
    
</dd>
</dl>

<dl>
<dd>

**customTags:** `*v2.CustomTags` 
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*v2.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## BreakingNews
<details><summary><code>client.BreakingNews.Get() -> *v2.BreakingNewsResponseDto</code></summary>
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
request := &v2.GetBreakingNewsRequest{
        RankedOnly: v2.Bool(
            true,
        ),
        FromRank: v2.Int(
            100,
        ),
        ToRank: v2.Int(
            100,
        ),
        Page: v2.Int(
            2,
        ),
        PageSize: v2.Int(
            50,
        ),
        TopNArticles: v2.Int(
            5,
        ),
        IncludeTranslationFields: v2.Bool(
            true,
        ),
        IncludeNlpData: v2.Bool(
            true,
        ),
        HasNlp: v2.Bool(
            true,
        ),
        Theme: v2.String(
            "Finance,Tech",
        ),
        NotTheme: v2.String(
            "Crime,Sports",
        ),
        OrgEntityName: v2.String(
            `"Apple Inc" OR Microsoft`,
        ),
        PerEntityName: v2.String(
            `"Elon Musk" OR "Jeff Bezos"`,
        ),
        LocEntityName: v2.String(
            `"San Francisco" OR "New York City"`,
        ),
        MiscEntityName: v2.String(
            `AWS OR "Microsoft Azure"`,
        ),
        TitleSentimentMin: v2.Float64(
            -0.5,
        ),
        TitleSentimentMax: v2.Float64(
            0.5,
        ),
        ContentSentimentMin: v2.Float64(
            -0.5,
        ),
        ContentSentimentMax: v2.Float64(
            0.5,
        ),
    }
client.BreakingNews.Get(
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

**sortBy:** `*v2.SortBy` 
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*v2.RankedOnly` 
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*v2.FromRank` 
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*v2.ToRank` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*v2.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*v2.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**topNArticles:** `*v2.TopNArticles` 
    
</dd>
</dl>

<dl>
<dd>

**includeTranslationFields:** `*v2.IncludeTranslationFields` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*v2.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*v2.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*v2.Theme` 
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*v2.NotTheme` 
    
</dd>
</dl>

<dl>
<dd>

**orgEntityName:** `*v2.OrgEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**perEntityName:** `*v2.PerEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**locEntityName:** `*v2.LocEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**miscEntityName:** `*v2.MiscEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*v2.TitleSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*v2.TitleSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*v2.ContentSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*v2.ContentSentimentMax` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BreakingNews.Post(request) -> *v2.BreakingNewsResponseDto</code></summary>
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
request := &v2.PostBreakingNewsRequest{
        SortBy: v2.SortByRelevancy.Ptr(),
        RankedOnly: v2.Bool(
            true,
        ),
        TopNArticles: v2.Int(
            1,
        ),
    }
client.BreakingNews.Post(
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

**sortBy:** `*v2.SortBy` 
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*v2.RankedOnly` 
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*v2.FromRank` 
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*v2.ToRank` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*v2.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*v2.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**topNArticles:** `*v2.TopNArticles` 
    
</dd>
</dl>

<dl>
<dd>

**includeTranslationFields:** `*v2.IncludeTranslationFields` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*v2.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*v2.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*v2.Theme` 
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*v2.NotTheme` 
    
</dd>
</dl>

<dl>
<dd>

**orgEntityName:** `*v2.OrgEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**perEntityName:** `*v2.PerEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**locEntityName:** `*v2.LocEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**miscEntityName:** `*v2.MiscEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*v2.TitleSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*v2.TitleSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*v2.ContentSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*v2.ContentSentimentMax` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Authors
<details><summary><code>client.Authors.Get() -> *v2.GetAuthorsResponse</code></summary>
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
request := &v2.GetAuthorsRequest{
        AuthorName: "Jane Smith",
        NotAuthorName: v2.String(
            "John Doe, Jane Doe",
        ),
        PredefinedSources: v2.String(
            "top 50 US, top 20 GB",
        ),
        Sources: v2.String(
            "nytimes.com,finance.yahoo.com",
        ),
        NotSources: v2.String(
            "cnn.com,wsj.com",
        ),
        Lang: v2.String(
            "en,es",
        ),
        NotLang: v2.String(
            "fr,de",
        ),
        Countries: v2.String(
            "US,CA",
        ),
        NotCountries: v2.String(
            "UK,FR",
        ),
        From: &v2.From{
            DateTime: v2.MustParseDateTime(
                "2024-07-01T00:00:00Z",
            ),
        },
        To: &v2.To{
            DateTime: v2.MustParseDateTime(
                "2024-01-01T00:00:00Z",
            ),
        },
        PublishedDatePrecision: v2.String(
            "full",
        ),
        ByParseDate: v2.Bool(
            true,
        ),
        RankedOnly: v2.Bool(
            true,
        ),
        FromRank: v2.Int(
            100,
        ),
        ToRank: v2.Int(
            100,
        ),
        IsHeadline: v2.Bool(
            true,
        ),
        IsOpinion: v2.Bool(
            true,
        ),
        IsPaidContent: v2.Bool(
            false,
        ),
        ParentUrl: v2.String(
            "wsj.com/politics,wsj.com/tech",
        ),
        AllLinks: v2.String(
            "https://aiindex.stanford.edu/report,https://www.stateof.ai",
        ),
        AllDomainLinks: v2.String(
            "who.int,nih.gov",
        ),
        AllLinksText: v2.String(
            "Nvidia,Tesla",
        ),
        WordCountMin: v2.Int(
            300,
        ),
        WordCountMax: v2.Int(
            1000,
        ),
        Page: v2.Int(
            2,
        ),
        PageSize: v2.Int(
            50,
        ),
        IncludeTranslationFields: v2.Bool(
            true,
        ),
        IncludeNlpData: v2.Bool(
            true,
        ),
        HasNlp: v2.Bool(
            true,
        ),
        Theme: v2.String(
            "Finance,Tech",
        ),
        NotTheme: v2.String(
            "Crime,Sports",
        ),
        NerName: v2.String(
            "Tesla,Amazon",
        ),
        TitleSentimentMin: v2.Float64(
            -0.5,
        ),
        TitleSentimentMax: v2.Float64(
            0.5,
        ),
        ContentSentimentMin: v2.Float64(
            -0.5,
        ),
        ContentSentimentMax: v2.Float64(
            0.5,
        ),
        IptcTags: v2.String(
            "20000199,20000209",
        ),
        NotIptcTags: v2.String(
            "20000205,20000209",
        ),
        IabTags: v2.String(
            "Business,Events",
        ),
        NotIabTags: v2.String(
            "Agriculture,Metals",
        ),
        CustomTags: v2.String(
            "Tag1,Tag2",
        ),
        RobotsCompliant: v2.Bool(
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

**authorName:** `v2.AuthorName` 
    
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

**from:** `*v2.From` 
    
</dd>
</dl>

<dl>
<dd>

**to:** `*v2.To` 
    
</dd>
</dl>

<dl>
<dd>

**publishedDatePrecision:** `*v2.PublishedDatePrecision` 
    
</dd>
</dl>

<dl>
<dd>

**byParseDate:** `*v2.ByParseDate` 
    
</dd>
</dl>

<dl>
<dd>

**sortBy:** `*v2.SortBy` 
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*v2.RankedOnly` 
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*v2.FromRank` 
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*v2.ToRank` 
    
</dd>
</dl>

<dl>
<dd>

**isHeadline:** `*v2.IsHeadline` 
    
</dd>
</dl>

<dl>
<dd>

**isOpinion:** `*v2.IsOpinion` 
    
</dd>
</dl>

<dl>
<dd>

**isPaidContent:** `*v2.IsPaidContent` 
    
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

**wordCountMin:** `*v2.WordCountMin` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMax:** `*v2.WordCountMax` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*v2.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*v2.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**includeTranslationFields:** `*v2.IncludeTranslationFields` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*v2.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*v2.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*v2.Theme` 
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*v2.NotTheme` 
    
</dd>
</dl>

<dl>
<dd>

**nerName:** `*v2.NerName` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*v2.TitleSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*v2.TitleSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*v2.ContentSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*v2.ContentSentimentMax` 
    
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

**robotsCompliant:** `*v2.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Authors.Post(request) -> *v2.PostAuthorsResponse</code></summary>
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
request := &v2.PostAuthorsRequest{
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

**authorName:** `v2.AuthorName` 
    
</dd>
</dl>

<dl>
<dd>

**notAuthorName:** `*v2.NotAuthorName` 
    
</dd>
</dl>

<dl>
<dd>

**predefinedSources:** `*v2.PredefinedSources` 
    
</dd>
</dl>

<dl>
<dd>

**sources:** `*v2.Sources` 
    
</dd>
</dl>

<dl>
<dd>

**notSources:** `*v2.NotSources` 
    
</dd>
</dl>

<dl>
<dd>

**lang:** `*v2.Lang` 
    
</dd>
</dl>

<dl>
<dd>

**notLang:** `*v2.NotLang` 
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*v2.Countries` 
    
</dd>
</dl>

<dl>
<dd>

**notCountries:** `*v2.NotCountries` 
    
</dd>
</dl>

<dl>
<dd>

**from:** `*v2.From` 
    
</dd>
</dl>

<dl>
<dd>

**to:** `*v2.To` 
    
</dd>
</dl>

<dl>
<dd>

**publishedDatePrecision:** `*v2.PublishedDatePrecision` 
    
</dd>
</dl>

<dl>
<dd>

**byParseDate:** `*v2.ByParseDate` 
    
</dd>
</dl>

<dl>
<dd>

**sortBy:** `*v2.SortBy` 
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*v2.RankedOnly` 
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*v2.FromRank` 
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*v2.ToRank` 
    
</dd>
</dl>

<dl>
<dd>

**isHeadline:** `*v2.IsHeadline` 
    
</dd>
</dl>

<dl>
<dd>

**isOpinion:** `*v2.IsOpinion` 
    
</dd>
</dl>

<dl>
<dd>

**isPaidContent:** `*v2.IsPaidContent` 
    
</dd>
</dl>

<dl>
<dd>

**parentUrl:** `*v2.ParentUrl` 
    
</dd>
</dl>

<dl>
<dd>

**allLinks:** `*v2.AllLinks` 
    
</dd>
</dl>

<dl>
<dd>

**allDomainLinks:** `*v2.AllDomainLinks` 
    
</dd>
</dl>

<dl>
<dd>

**allLinksText:** `*v2.AllLinksText` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMin:** `*v2.WordCountMin` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMax:** `*v2.WordCountMax` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*v2.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*v2.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**includeTranslationFields:** `*v2.IncludeTranslationFields` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*v2.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*v2.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*v2.Theme` 
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*v2.NotTheme` 
    
</dd>
</dl>

<dl>
<dd>

**nerName:** `*v2.NerName` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*v2.TitleSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*v2.TitleSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*v2.ContentSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*v2.ContentSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**iptcTags:** `*v2.IptcTags` 
    
</dd>
</dl>

<dl>
<dd>

**notIptcTags:** `*v2.NotIptcTags` 
    
</dd>
</dl>

<dl>
<dd>

**iabTags:** `*v2.IabTags` 
    
</dd>
</dl>

<dl>
<dd>

**notIabTags:** `*v2.NotIabTags` 
    
</dd>
</dl>

<dl>
<dd>

**customTags:** `*v2.CustomTags` 
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*v2.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## SearchByLink
<details><summary><code>client.SearchByLink.Get() -> *v2.SearchResponseDto</code></summary>
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
request := &v2.GetSearchByLinkRequest{
        Ids: v2.String(
            "5f8d0d55b6e45e00179c6e7e",
        ),
        Links: v2.String(
            "https://nytimes.com/article1,https://bbc.com/article2",
        ),
        From: &v2.From{
            DateTime: v2.MustParseDateTime(
                "2024-07-01T00:00:00Z",
            ),
        },
        To: &v2.To{
            DateTime: v2.MustParseDateTime(
                "2024-01-01T00:00:00Z",
            ),
        },
        Page: v2.Int(
            2,
        ),
        PageSize: v2.Int(
            50,
        ),
        RobotsCompliant: v2.Bool(
            true,
        ),
    }
client.SearchByLink.Get(
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

**from:** `*v2.From` 
    
</dd>
</dl>

<dl>
<dd>

**to:** `*v2.To` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*v2.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*v2.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*v2.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SearchByLink.Post(request) -> *v2.SearchResponseDto</code></summary>
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
request := &v2.PostSearchByLinkRequest{
        Links: &v2.Links{
            String: "https://www.reuters.com/business/energy/oil-prices-up-after-israeli-attacks-oversupply-caps-gains-2025-09-10/",
        },
    }
client.SearchByLink.Post(
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

**ids:** `*v2.Ids` 
    
</dd>
</dl>

<dl>
<dd>

**links:** `*v2.Links` 
    
</dd>
</dl>

<dl>
<dd>

**from:** `*v2.From` 

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

**to:** `*v2.To` 

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

**page:** `*v2.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*v2.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*v2.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Sources
<details><summary><code>client.Sources.Get() -> *v2.SourcesResponseDto</code></summary>
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
request := &v2.GetSourcesRequest{
        Lang: v2.String(
            "en,es",
        ),
        Countries: v2.String(
            "US,CA",
        ),
        PredefinedSources: v2.String(
            "top 50 US, top 20 GB",
        ),
        SourceName: v2.String(
            "sport,tech",
        ),
        SourceUrl: v2.String(
            "bbc.com",
        ),
        IncludeAdditionalInfo: v2.Bool(
            true,
        ),
        IsNewsDomain: v2.Bool(
            true,
        ),
        NewsType: v2.String(
            "General News Outlets,Tech News and Updates",
        ),
        FromRank: v2.Int(
            100,
        ),
        ToRank: v2.Int(
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

**includeAdditionalInfo:** `*v2.IncludeAdditionalInfo` 
    
</dd>
</dl>

<dl>
<dd>

**isNewsDomain:** `*v2.IsNewsDomain` 
    
</dd>
</dl>

<dl>
<dd>

**newsDomainType:** `*v2.NewsDomainType` 
    
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

**fromRank:** `*v2.FromRank` 
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*v2.ToRank` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sources.Post(request) -> *v2.SourcesResponseDto</code></summary>
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
request := &v2.PostSourcesRequest{
        PredefinedSources: &v2.PredefinedSources{
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

**lang:** `*v2.Lang` 
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*v2.Countries` 
    
</dd>
</dl>

<dl>
<dd>

**predefinedSources:** `*v2.PredefinedSources` 
    
</dd>
</dl>

<dl>
<dd>

**sourceName:** `*v2.SourceName` 
    
</dd>
</dl>

<dl>
<dd>

**sourceUrl:** `*v2.SourceUrl` 
    
</dd>
</dl>

<dl>
<dd>

**includeAdditionalInfo:** `*v2.IncludeAdditionalInfo` 
    
</dd>
</dl>

<dl>
<dd>

**isNewsDomain:** `*v2.IsNewsDomain` 
    
</dd>
</dl>

<dl>
<dd>

**newsDomainType:** `*v2.NewsDomainType` 
    
</dd>
</dl>

<dl>
<dd>

**newsType:** `*v2.NewsType` 
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*v2.FromRank` 
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*v2.ToRank` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## AggregationCount
<details><summary><code>client.AggregationCount.Get() -> *v2.GetAggregationCountResponse</code></summary>
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
request := &v2.GetAggregationCountRequest{
        Q: `"supply chain" AND Amazon NOT China`,
        SearchIn: v2.String(
            "title_content, title_content_translated",
        ),
        PredefinedSources: v2.String(
            "top 50 US, top 20 GB",
        ),
        Sources: v2.String(
            "nytimes.com,finance.yahoo.com",
        ),
        NotSources: v2.String(
            "cnn.com,wsj.com",
        ),
        Lang: v2.String(
            "en,es",
        ),
        NotLang: v2.String(
            "fr,de",
        ),
        Countries: v2.String(
            "US,CA",
        ),
        NotCountries: v2.String(
            "UK,FR",
        ),
        NotAuthorName: v2.String(
            "John Doe, Jane Doe",
        ),
        From: &v2.From{
            DateTime: v2.MustParseDateTime(
                "2024-07-01T00:00:00Z",
            ),
        },
        To: &v2.To{
            DateTime: v2.MustParseDateTime(
                "2024-01-01T00:00:00Z",
            ),
        },
        PublishedDatePrecision: v2.String(
            "full",
        ),
        ByParseDate: v2.Bool(
            true,
        ),
        RankedOnly: v2.Bool(
            true,
        ),
        FromRank: v2.Int(
            100,
        ),
        ToRank: v2.Int(
            100,
        ),
        IsHeadline: v2.Bool(
            true,
        ),
        IsOpinion: v2.Bool(
            true,
        ),
        IsPaidContent: v2.Bool(
            false,
        ),
        ParentUrl: v2.String(
            "wsj.com/politics,wsj.com/tech",
        ),
        AllLinks: v2.String(
            "https://aiindex.stanford.edu/report,https://www.stateof.ai",
        ),
        AllDomainLinks: v2.String(
            "who.int,nih.gov",
        ),
        AllLinksText: v2.String(
            "Nvidia,Tesla",
        ),
        WordCountMin: v2.Int(
            300,
        ),
        WordCountMax: v2.Int(
            1000,
        ),
        Page: v2.Int(
            2,
        ),
        PageSize: v2.Int(
            50,
        ),
        IncludeNlpData: v2.Bool(
            true,
        ),
        HasNlp: v2.Bool(
            true,
        ),
        Theme: v2.String(
            "Finance,Tech",
        ),
        NotTheme: v2.String(
            "Crime,Sports",
        ),
        OrgEntityName: v2.String(
            `"Apple Inc" OR Microsoft`,
        ),
        PerEntityName: v2.String(
            `"Elon Musk" OR "Jeff Bezos"`,
        ),
        LocEntityName: v2.String(
            `"San Francisco" OR "New York City"`,
        ),
        MiscEntityName: v2.String(
            `AWS OR "Microsoft Azure"`,
        ),
        TitleSentimentMin: v2.Float64(
            -0.5,
        ),
        TitleSentimentMax: v2.Float64(
            0.5,
        ),
        ContentSentimentMin: v2.Float64(
            -0.5,
        ),
        ContentSentimentMax: v2.Float64(
            0.5,
        ),
        IptcTags: v2.String(
            "20000199,20000209",
        ),
        NotIptcTags: v2.String(
            "20000205,20000209",
        ),
        RobotsCompliant: v2.Bool(
            true,
        ),
    }
client.AggregationCount.Get(
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

**q:** `v2.Q` 
    
</dd>
</dl>

<dl>
<dd>

**aggregationBy:** `*v2.AggregationBy` 
    
</dd>
</dl>

<dl>
<dd>

**searchIn:** `*v2.SearchIn` 
    
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

**from:** `*v2.From` 
    
</dd>
</dl>

<dl>
<dd>

**to:** `*v2.To` 
    
</dd>
</dl>

<dl>
<dd>

**publishedDatePrecision:** `*v2.PublishedDatePrecision` 
    
</dd>
</dl>

<dl>
<dd>

**byParseDate:** `*v2.ByParseDate` 
    
</dd>
</dl>

<dl>
<dd>

**sortBy:** `*v2.SortBy` 
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*v2.RankedOnly` 
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*v2.FromRank` 
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*v2.ToRank` 
    
</dd>
</dl>

<dl>
<dd>

**isHeadline:** `*v2.IsHeadline` 
    
</dd>
</dl>

<dl>
<dd>

**isOpinion:** `*v2.IsOpinion` 
    
</dd>
</dl>

<dl>
<dd>

**isPaidContent:** `*v2.IsPaidContent` 
    
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

**wordCountMin:** `*v2.WordCountMin` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMax:** `*v2.WordCountMax` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*v2.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*v2.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*v2.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*v2.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*v2.Theme` 
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*v2.NotTheme` 
    
</dd>
</dl>

<dl>
<dd>

**orgEntityName:** `*v2.OrgEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**perEntityName:** `*v2.PerEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**locEntityName:** `*v2.LocEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**miscEntityName:** `*v2.MiscEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*v2.TitleSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*v2.TitleSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*v2.ContentSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*v2.ContentSentimentMax` 
    
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

**robotsCompliant:** `*v2.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AggregationCount.Post(request) -> *v2.PostAggregationCountResponse</code></summary>
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
request := &v2.PostAggregationCountRequest{
        Q: `"supply chain" AND Amazon NOT China`,
        AggregationBy: v2.AggregationByDay.Ptr(),
    }
client.AggregationCount.Post(
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

**q:** `v2.Q` 
    
</dd>
</dl>

<dl>
<dd>

**aggregationBy:** `*v2.AggregationBy` 
    
</dd>
</dl>

<dl>
<dd>

**searchIn:** `*v2.SearchIn` 
    
</dd>
</dl>

<dl>
<dd>

**predefinedSources:** `*v2.PredefinedSources` 
    
</dd>
</dl>

<dl>
<dd>

**sources:** `*v2.Sources` 
    
</dd>
</dl>

<dl>
<dd>

**notSources:** `*v2.NotSources` 
    
</dd>
</dl>

<dl>
<dd>

**lang:** `*v2.Lang` 
    
</dd>
</dl>

<dl>
<dd>

**notLang:** `*v2.NotLang` 
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*v2.Countries` 
    
</dd>
</dl>

<dl>
<dd>

**notCountries:** `*v2.NotCountries` 
    
</dd>
</dl>

<dl>
<dd>

**notAuthorName:** `*v2.NotAuthorName` 
    
</dd>
</dl>

<dl>
<dd>

**from:** `*v2.From` 
    
</dd>
</dl>

<dl>
<dd>

**to:** `*v2.To` 
    
</dd>
</dl>

<dl>
<dd>

**publishedDatePrecision:** `*v2.PublishedDatePrecision` 
    
</dd>
</dl>

<dl>
<dd>

**byParseDate:** `*v2.ByParseDate` 
    
</dd>
</dl>

<dl>
<dd>

**sortBy:** `*v2.SortBy` 
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*v2.RankedOnly` 
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*v2.FromRank` 
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*v2.ToRank` 
    
</dd>
</dl>

<dl>
<dd>

**isHeadline:** `*v2.IsHeadline` 
    
</dd>
</dl>

<dl>
<dd>

**isOpinion:** `*v2.IsOpinion` 
    
</dd>
</dl>

<dl>
<dd>

**isPaidContent:** `*v2.IsPaidContent` 
    
</dd>
</dl>

<dl>
<dd>

**parentUrl:** `*v2.ParentUrl` 
    
</dd>
</dl>

<dl>
<dd>

**allLinks:** `*v2.AllLinks` 
    
</dd>
</dl>

<dl>
<dd>

**allDomainLinks:** `*v2.AllDomainLinks` 
    
</dd>
</dl>

<dl>
<dd>

**allLinksText:** `*v2.AllLinksText` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMin:** `*v2.WordCountMin` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMax:** `*v2.WordCountMax` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*v2.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*v2.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*v2.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*v2.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*v2.Theme` 
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*v2.NotTheme` 
    
</dd>
</dl>

<dl>
<dd>

**orgEntityName:** `*v2.OrgEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**perEntityName:** `*v2.PerEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**locEntityName:** `*v2.LocEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**miscEntityName:** `*v2.MiscEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*v2.TitleSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*v2.TitleSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*v2.ContentSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*v2.ContentSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**iptcTags:** `*v2.IptcTags` 
    
</dd>
</dl>

<dl>
<dd>

**notIptcTags:** `*v2.NotIptcTags` 
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*v2.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Subscription
<details><summary><code>client.Subscription.Get() -> *v2.SubscriptionResponseDto</code></summary>
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

<details><summary><code>client.Subscription.Post() -> *v2.SubscriptionResponseDto</code></summary>
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

