# Reference
## Search
<details><summary><code>client.Search.Get() -> *newscatchergo.GetSearchResponse</code></summary>
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
request := &newscatchergo.GetSearchRequest{
        Q: `"supply chain" AND Amazon NOT China`,
        SearchIn: newscatchergo.String(
            "title_content, title_content_translated",
        ),
        IncludeTranslationFields: newscatchergo.Bool(
            true,
        ),
        PredefinedSources: newscatchergo.String(
            "top 50 US, top 20 GB",
        ),
        SourceName: newscatchergo.String(
            "sport,tech",
        ),
        Sources: newscatchergo.String(
            "nytimes.com,finance.yahoo.com",
        ),
        NotSources: newscatchergo.String(
            "cnn.com,wsj.com",
        ),
        Lang: newscatchergo.String(
            "en,es",
        ),
        NotLang: newscatchergo.String(
            "fr,de",
        ),
        Countries: newscatchergo.String(
            "US,CA",
        ),
        NotCountries: newscatchergo.String(
            "UK,FR",
        ),
        NotAuthorName: newscatchergo.String(
            "John Doe, Jane Doe",
        ),
        From: &newscatchergo.From{
            DateTime: newscatchergo.MustParseDateTime(
                "1 day ago",
            ),
        },
        To: &newscatchergo.To{
            String: "now",
        },
        PublishedDatePrecision: newscatchergo.String(
            "full",
        ),
        ByParseDate: newscatchergo.Bool(
            true,
        ),
        RankedOnly: newscatchergo.Bool(
            true,
        ),
        FromRank: newscatchergo.Int(
            100,
        ),
        ToRank: newscatchergo.Int(
            100,
        ),
        IsHeadline: newscatchergo.Bool(
            true,
        ),
        IsOpinion: newscatchergo.Bool(
            true,
        ),
        IsPaidContent: newscatchergo.Bool(
            false,
        ),
        ParentUrl: newscatchergo.String(
            "wsj.com/politics,wsj.com/tech",
        ),
        AllLinks: newscatchergo.String(
            "https://aiindex.stanford.edu/report,https://www.stateof.ai",
        ),
        AllDomainLinks: newscatchergo.String(
            "who.int,nih.gov",
        ),
        AllLinksText: newscatchergo.String(
            "Nvidia,Tesla",
        ),
        AdditionalDomainInfo: newscatchergo.Bool(
            true,
        ),
        IsNewsDomain: newscatchergo.Bool(
            true,
        ),
        NewsType: newscatchergo.String(
            "General News Outlets,Tech News and Updates",
        ),
        WordCountMin: newscatchergo.Int(
            300,
        ),
        WordCountMax: newscatchergo.Int(
            1000,
        ),
        Page: newscatchergo.Int(
            2,
        ),
        PageSize: newscatchergo.Int(
            50,
        ),
        ClusteringEnabled: newscatchergo.Bool(
            true,
        ),
        ClusteringThreshold: newscatchergo.Float64(
            0.7,
        ),
        IncludeNlpData: newscatchergo.Bool(
            true,
        ),
        HasNlp: newscatchergo.Bool(
            true,
        ),
        Theme: newscatchergo.String(
            "Finance,Tech",
        ),
        NotTheme: newscatchergo.String(
            "Crime,Sports",
        ),
        OrgEntityName: newscatchergo.String(
            `"Apple Inc" OR Microsoft`,
        ),
        PerEntityName: newscatchergo.String(
            `"Elon Musk" OR "Jeff Bezos"`,
        ),
        LocEntityName: newscatchergo.String(
            `"San Francisco" OR "New York City"`,
        ),
        MiscEntityName: newscatchergo.String(
            `AWS OR "Microsoft Azure"`,
        ),
        TitleSentimentMin: newscatchergo.Float64(
            -0.5,
        ),
        TitleSentimentMax: newscatchergo.Float64(
            0.5,
        ),
        ContentSentimentMin: newscatchergo.Float64(
            -0.5,
        ),
        ContentSentimentMax: newscatchergo.Float64(
            0.5,
        ),
        IptcTags: newscatchergo.String(
            "20000199,20000209",
        ),
        NotIptcTags: newscatchergo.String(
            "20000205,20000209",
        ),
        IabTags: newscatchergo.String(
            "Business,Events",
        ),
        NotIabTags: newscatchergo.String(
            "Agriculture,Metals",
        ),
        CustomTags: newscatchergo.String(
            "Tag1,Tag2",
        ),
        ExcludeDuplicates: newscatchergo.Bool(
            true,
        ),
        RobotsCompliant: newscatchergo.Bool(
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

**q:** `newscatchergo.Q` 
    
</dd>
</dl>

<dl>
<dd>

**searchIn:** `*newscatchergo.SearchIn` 
    
</dd>
</dl>

<dl>
<dd>

**includeTranslationFields:** `*newscatchergo.IncludeTranslationFields` 
    
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

**Note**: The search doesn't require an exact match and returns sources containing the specified terms in their names. You can use any word or phrase, like `"sport"` or `"new york times"`. 

For example, `"sport"` returns sources such as `"Motorsport"`, `"Dot Esport"`, and `"Tuttosport"`.
    
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

**from:** `*newscatchergo.From` 
    
</dd>
</dl>

<dl>
<dd>

**to:** `*newscatchergo.To` 
    
</dd>
</dl>

<dl>
<dd>

**publishedDatePrecision:** `*newscatchergo.PublishedDatePrecision` 
    
</dd>
</dl>

<dl>
<dd>

**byParseDate:** `*newscatchergo.ByParseDate` 
    
</dd>
</dl>

<dl>
<dd>

**sortBy:** `*newscatchergo.SortBy` 
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*newscatchergo.RankedOnly` 
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*newscatchergo.FromRank` 
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*newscatchergo.ToRank` 
    
</dd>
</dl>

<dl>
<dd>

**isHeadline:** `*newscatchergo.IsHeadline` 
    
</dd>
</dl>

<dl>
<dd>

**isOpinion:** `*newscatchergo.IsOpinion` 
    
</dd>
</dl>

<dl>
<dd>

**isPaidContent:** `*newscatchergo.IsPaidContent` 
    
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

**additionalDomainInfo:** `*newscatchergo.AdditionalDomainInfo` 
    
</dd>
</dl>

<dl>
<dd>

**isNewsDomain:** `*newscatchergo.IsNewsDomain` 
    
</dd>
</dl>

<dl>
<dd>

**newsDomainType:** `*newscatchergo.NewsDomainType` 
    
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

**wordCountMin:** `*newscatchergo.WordCountMin` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMax:** `*newscatchergo.WordCountMax` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*newscatchergo.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*newscatchergo.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringEnabled:** `*newscatchergo.ClusteringEnabled` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringVariable:** `*newscatchergo.ClusteringVariable` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringThreshold:** `*newscatchergo.ClusteringThreshold` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*newscatchergo.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*newscatchergo.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*newscatchergo.Theme` 
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*newscatchergo.NotTheme` 
    
</dd>
</dl>

<dl>
<dd>

**orgEntityName:** `*newscatchergo.OrgEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**perEntityName:** `*newscatchergo.PerEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**locEntityName:** `*newscatchergo.LocEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**miscEntityName:** `*newscatchergo.MiscEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*newscatchergo.TitleSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*newscatchergo.TitleSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*newscatchergo.ContentSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*newscatchergo.ContentSentimentMax` 
    
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

**excludeDuplicates:** `*newscatchergo.ExcludeDuplicates` 
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*newscatchergo.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Search.Post(request) -> *newscatchergo.PostSearchResponse</code></summary>
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
request := &newscatchergo.PostSearchRequest{
        Q: `"supply chain" AND Amazon NOT China`,
        PageSize: newscatchergo.Int(
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

**q:** `newscatchergo.Q` 
    
</dd>
</dl>

<dl>
<dd>

**searchIn:** `*newscatchergo.SearchIn` 
    
</dd>
</dl>

<dl>
<dd>

**includeTranslationFields:** `*newscatchergo.IncludeTranslationFields` 
    
</dd>
</dl>

<dl>
<dd>

**predefinedSources:** `*newscatchergo.PredefinedSources` 
    
</dd>
</dl>

<dl>
<dd>

**sourceName:** `*newscatchergo.SourceName` 
    
</dd>
</dl>

<dl>
<dd>

**sources:** `*newscatchergo.Sources` 
    
</dd>
</dl>

<dl>
<dd>

**notSources:** `*newscatchergo.NotSources` 
    
</dd>
</dl>

<dl>
<dd>

**lang:** `*newscatchergo.Lang` 
    
</dd>
</dl>

<dl>
<dd>

**notLang:** `*newscatchergo.NotLang` 
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*newscatchergo.Countries` 
    
</dd>
</dl>

<dl>
<dd>

**notCountries:** `*newscatchergo.NotCountries` 
    
</dd>
</dl>

<dl>
<dd>

**notAuthorName:** `*newscatchergo.NotAuthorName` 
    
</dd>
</dl>

<dl>
<dd>

**from:** `*newscatchergo.From` 
    
</dd>
</dl>

<dl>
<dd>

**to:** `*newscatchergo.To` 
    
</dd>
</dl>

<dl>
<dd>

**publishedDatePrecision:** `*newscatchergo.PublishedDatePrecision` 
    
</dd>
</dl>

<dl>
<dd>

**byParseDate:** `*newscatchergo.ByParseDate` 
    
</dd>
</dl>

<dl>
<dd>

**sortBy:** `*newscatchergo.SortBy` 
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*newscatchergo.RankedOnly` 
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*newscatchergo.FromRank` 
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*newscatchergo.ToRank` 
    
</dd>
</dl>

<dl>
<dd>

**isHeadline:** `*newscatchergo.IsHeadline` 
    
</dd>
</dl>

<dl>
<dd>

**isOpinion:** `*newscatchergo.IsOpinion` 
    
</dd>
</dl>

<dl>
<dd>

**isPaidContent:** `*newscatchergo.IsPaidContent` 
    
</dd>
</dl>

<dl>
<dd>

**parentUrl:** `*newscatchergo.ParentUrl` 
    
</dd>
</dl>

<dl>
<dd>

**allLinks:** `*newscatchergo.AllLinks` 
    
</dd>
</dl>

<dl>
<dd>

**allDomainLinks:** `*newscatchergo.AllDomainLinks` 
    
</dd>
</dl>

<dl>
<dd>

**allLinksText:** `*newscatchergo.AllLinksText` 
    
</dd>
</dl>

<dl>
<dd>

**additionalDomainInfo:** `*newscatchergo.AdditionalDomainInfo` 
    
</dd>
</dl>

<dl>
<dd>

**isNewsDomain:** `*newscatchergo.IsNewsDomain` 
    
</dd>
</dl>

<dl>
<dd>

**newsDomainType:** `*newscatchergo.NewsDomainType` 
    
</dd>
</dl>

<dl>
<dd>

**newsType:** `*newscatchergo.NewsType` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMin:** `*newscatchergo.WordCountMin` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMax:** `*newscatchergo.WordCountMax` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*newscatchergo.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*newscatchergo.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringEnabled:** `*newscatchergo.ClusteringEnabled` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringVariable:** `*newscatchergo.ClusteringVariable` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringThreshold:** `*newscatchergo.ClusteringThreshold` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*newscatchergo.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*newscatchergo.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*newscatchergo.Theme` 
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*newscatchergo.NotTheme` 
    
</dd>
</dl>

<dl>
<dd>

**orgEntityName:** `*newscatchergo.OrgEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**perEntityName:** `*newscatchergo.PerEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**locEntityName:** `*newscatchergo.LocEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**miscEntityName:** `*newscatchergo.MiscEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*newscatchergo.TitleSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*newscatchergo.TitleSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*newscatchergo.ContentSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*newscatchergo.ContentSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**iptcTags:** `*newscatchergo.IptcTags` 
    
</dd>
</dl>

<dl>
<dd>

**notIptcTags:** `*newscatchergo.NotIptcTags` 
    
</dd>
</dl>

<dl>
<dd>

**iabTags:** `*newscatchergo.IabTags` 
    
</dd>
</dl>

<dl>
<dd>

**notIabTags:** `*newscatchergo.NotIabTags` 
    
</dd>
</dl>

<dl>
<dd>

**customTags:** `*newscatchergo.CustomTags` 
    
</dd>
</dl>

<dl>
<dd>

**excludeDuplicates:** `*newscatchergo.ExcludeDuplicates` 
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*newscatchergo.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## LatestHeadlines
<details><summary><code>client.LatestHeadlines.Get() -> *newscatchergo.GetLatestHeadlinesResponse</code></summary>
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
request := &newscatchergo.GetLatestHeadlinesRequest{
        When: newscatchergo.String(
            "7d",
        ),
        ByParseDate: newscatchergo.Bool(
            true,
        ),
        Lang: newscatchergo.String(
            "en,es",
        ),
        NotLang: newscatchergo.String(
            "fr,de",
        ),
        Countries: newscatchergo.String(
            "US,CA",
        ),
        NotCountries: newscatchergo.String(
            "UK,FR",
        ),
        PredefinedSources: newscatchergo.String(
            "top 50 US, top 20 GB",
        ),
        Sources: newscatchergo.String(
            "nytimes.com,finance.yahoo.com",
        ),
        NotSources: newscatchergo.String(
            "cnn.com,wsj.com",
        ),
        NotAuthorName: newscatchergo.String(
            "John Doe, Jane Doe",
        ),
        RankedOnly: newscatchergo.Bool(
            true,
        ),
        IsHeadline: newscatchergo.Bool(
            true,
        ),
        IsOpinion: newscatchergo.Bool(
            true,
        ),
        IsPaidContent: newscatchergo.Bool(
            false,
        ),
        ParentUrl: newscatchergo.String(
            "wsj.com/politics,wsj.com/tech",
        ),
        AllLinks: newscatchergo.String(
            "https://aiindex.stanford.edu/report,https://www.stateof.ai",
        ),
        AllDomainLinks: newscatchergo.String(
            "who.int,nih.gov",
        ),
        AllLinksText: newscatchergo.String(
            "Nvidia,Tesla",
        ),
        WordCountMin: newscatchergo.Int(
            300,
        ),
        WordCountMax: newscatchergo.Int(
            1000,
        ),
        Page: newscatchergo.Int(
            2,
        ),
        PageSize: newscatchergo.Int(
            50,
        ),
        ClusteringEnabled: newscatchergo.Bool(
            true,
        ),
        ClusteringThreshold: newscatchergo.Float64(
            0.7,
        ),
        IncludeTranslationFields: newscatchergo.Bool(
            true,
        ),
        IncludeNlpData: newscatchergo.Bool(
            true,
        ),
        HasNlp: newscatchergo.Bool(
            true,
        ),
        Theme: newscatchergo.String(
            "Finance,Tech",
        ),
        NotTheme: newscatchergo.String(
            "Crime,Sports",
        ),
        OrgEntityName: newscatchergo.String(
            `"Apple Inc" OR Microsoft`,
        ),
        PerEntityName: newscatchergo.String(
            `"Elon Musk" OR "Jeff Bezos"`,
        ),
        LocEntityName: newscatchergo.String(
            `"San Francisco" OR "New York City"`,
        ),
        MiscEntityName: newscatchergo.String(
            `AWS OR "Microsoft Azure"`,
        ),
        TitleSentimentMin: newscatchergo.Float64(
            -0.5,
        ),
        TitleSentimentMax: newscatchergo.Float64(
            0.5,
        ),
        ContentSentimentMin: newscatchergo.Float64(
            -0.5,
        ),
        ContentSentimentMax: newscatchergo.Float64(
            0.5,
        ),
        IptcTags: newscatchergo.String(
            "20000199,20000209",
        ),
        NotIptcTags: newscatchergo.String(
            "20000205,20000209",
        ),
        IabTags: newscatchergo.String(
            "Business,Events",
        ),
        NotIabTags: newscatchergo.String(
            "Agriculture,Metals",
        ),
        CustomTags: newscatchergo.String(
            "Tag1,Tag2",
        ),
        RobotsCompliant: newscatchergo.Bool(
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

**when:** `*newscatchergo.When` 
    
</dd>
</dl>

<dl>
<dd>

**byParseDate:** `*newscatchergo.ByParseDate` 
    
</dd>
</dl>

<dl>
<dd>

**sortBy:** `*newscatchergo.SortBy` 
    
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

**rankedOnly:** `*newscatchergo.RankedOnly` 
    
</dd>
</dl>

<dl>
<dd>

**isHeadline:** `*newscatchergo.IsHeadline` 
    
</dd>
</dl>

<dl>
<dd>

**isOpinion:** `*newscatchergo.IsOpinion` 
    
</dd>
</dl>

<dl>
<dd>

**isPaidContent:** `*newscatchergo.IsPaidContent` 
    
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

**wordCountMin:** `*newscatchergo.WordCountMin` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMax:** `*newscatchergo.WordCountMax` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*newscatchergo.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*newscatchergo.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringEnabled:** `*newscatchergo.ClusteringEnabled` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringVariable:** `*newscatchergo.ClusteringVariable` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringThreshold:** `*newscatchergo.ClusteringThreshold` 
    
</dd>
</dl>

<dl>
<dd>

**includeTranslationFields:** `*newscatchergo.IncludeTranslationFields` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*newscatchergo.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*newscatchergo.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*newscatchergo.Theme` 
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*newscatchergo.NotTheme` 
    
</dd>
</dl>

<dl>
<dd>

**orgEntityName:** `*newscatchergo.OrgEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**perEntityName:** `*newscatchergo.PerEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**locEntityName:** `*newscatchergo.LocEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**miscEntityName:** `*newscatchergo.MiscEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*newscatchergo.TitleSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*newscatchergo.TitleSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*newscatchergo.ContentSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*newscatchergo.ContentSentimentMax` 
    
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

**robotsCompliant:** `*newscatchergo.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.LatestHeadlines.Post(request) -> *newscatchergo.PostLatestHeadlinesResponse</code></summary>
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
request := &newscatchergo.PostLatestHeadlinesRequest{
        When: newscatchergo.String(
            "7d",
        ),
        PageSize: newscatchergo.Int(
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

**when:** `*newscatchergo.When` 
    
</dd>
</dl>

<dl>
<dd>

**byParseDate:** `*newscatchergo.ByParseDate` 
    
</dd>
</dl>

<dl>
<dd>

**lang:** `*newscatchergo.Lang` 
    
</dd>
</dl>

<dl>
<dd>

**notLang:** `*newscatchergo.NotLang` 
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*newscatchergo.Countries` 
    
</dd>
</dl>

<dl>
<dd>

**notCountries:** `*newscatchergo.NotCountries` 
    
</dd>
</dl>

<dl>
<dd>

**predefinedSources:** `*newscatchergo.PredefinedSources` 
    
</dd>
</dl>

<dl>
<dd>

**sources:** `*newscatchergo.Sources` 
    
</dd>
</dl>

<dl>
<dd>

**notSources:** `*newscatchergo.NotSources` 
    
</dd>
</dl>

<dl>
<dd>

**notAuthorName:** `*newscatchergo.NotAuthorName` 
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*newscatchergo.RankedOnly` 
    
</dd>
</dl>

<dl>
<dd>

**isHeadline:** `*newscatchergo.IsHeadline` 
    
</dd>
</dl>

<dl>
<dd>

**isOpinion:** `*newscatchergo.IsOpinion` 
    
</dd>
</dl>

<dl>
<dd>

**isPaidContent:** `*newscatchergo.IsPaidContent` 
    
</dd>
</dl>

<dl>
<dd>

**parentUrl:** `*newscatchergo.ParentUrl` 
    
</dd>
</dl>

<dl>
<dd>

**allLinks:** `*newscatchergo.AllLinks` 
    
</dd>
</dl>

<dl>
<dd>

**allDomainLinks:** `*newscatchergo.AllDomainLinks` 
    
</dd>
</dl>

<dl>
<dd>

**allLinksText:** `*newscatchergo.AllLinksText` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMin:** `*newscatchergo.WordCountMin` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMax:** `*newscatchergo.WordCountMax` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*newscatchergo.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*newscatchergo.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringEnabled:** `*newscatchergo.ClusteringEnabled` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringVariable:** `*newscatchergo.ClusteringVariable` 
    
</dd>
</dl>

<dl>
<dd>

**clusteringThreshold:** `*newscatchergo.ClusteringThreshold` 
    
</dd>
</dl>

<dl>
<dd>

**includeTranslationFields:** `*newscatchergo.IncludeTranslationFields` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*newscatchergo.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*newscatchergo.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*newscatchergo.Theme` 
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*newscatchergo.NotTheme` 
    
</dd>
</dl>

<dl>
<dd>

**orgEntityName:** `*newscatchergo.OrgEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**perEntityName:** `*newscatchergo.PerEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**locEntityName:** `*newscatchergo.LocEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**miscEntityName:** `*newscatchergo.MiscEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*newscatchergo.TitleSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*newscatchergo.TitleSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*newscatchergo.ContentSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*newscatchergo.ContentSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**iptcTags:** `*newscatchergo.IptcTags` 
    
</dd>
</dl>

<dl>
<dd>

**notIptcTags:** `*newscatchergo.NotIptcTags` 
    
</dd>
</dl>

<dl>
<dd>

**iabTags:** `*newscatchergo.IabTags` 
    
</dd>
</dl>

<dl>
<dd>

**notIabTags:** `*newscatchergo.NotIabTags` 
    
</dd>
</dl>

<dl>
<dd>

**customTags:** `*newscatchergo.CustomTags` 
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*newscatchergo.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## BreakingNews
<details><summary><code>client.BreakingNews.Get() -> *newscatchergo.BreakingNewsResponseDto</code></summary>
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
request := &newscatchergo.GetBreakingNewsRequest{
        RankedOnly: newscatchergo.Bool(
            true,
        ),
        FromRank: newscatchergo.Int(
            100,
        ),
        ToRank: newscatchergo.Int(
            100,
        ),
        Page: newscatchergo.Int(
            2,
        ),
        PageSize: newscatchergo.Int(
            50,
        ),
        TopNArticles: newscatchergo.Int(
            5,
        ),
        IncludeTranslationFields: newscatchergo.Bool(
            true,
        ),
        IncludeNlpData: newscatchergo.Bool(
            true,
        ),
        HasNlp: newscatchergo.Bool(
            true,
        ),
        Theme: newscatchergo.String(
            "Finance,Tech",
        ),
        NotTheme: newscatchergo.String(
            "Crime,Sports",
        ),
        OrgEntityName: newscatchergo.String(
            `"Apple Inc" OR Microsoft`,
        ),
        PerEntityName: newscatchergo.String(
            `"Elon Musk" OR "Jeff Bezos"`,
        ),
        LocEntityName: newscatchergo.String(
            `"San Francisco" OR "New York City"`,
        ),
        MiscEntityName: newscatchergo.String(
            `AWS OR "Microsoft Azure"`,
        ),
        TitleSentimentMin: newscatchergo.Float64(
            -0.5,
        ),
        TitleSentimentMax: newscatchergo.Float64(
            0.5,
        ),
        ContentSentimentMin: newscatchergo.Float64(
            -0.5,
        ),
        ContentSentimentMax: newscatchergo.Float64(
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

**sortBy:** `*newscatchergo.SortBy` 
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*newscatchergo.RankedOnly` 
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*newscatchergo.FromRank` 
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*newscatchergo.ToRank` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*newscatchergo.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*newscatchergo.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**topNArticles:** `*newscatchergo.TopNArticles` 
    
</dd>
</dl>

<dl>
<dd>

**includeTranslationFields:** `*newscatchergo.IncludeTranslationFields` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*newscatchergo.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*newscatchergo.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*newscatchergo.Theme` 
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*newscatchergo.NotTheme` 
    
</dd>
</dl>

<dl>
<dd>

**orgEntityName:** `*newscatchergo.OrgEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**perEntityName:** `*newscatchergo.PerEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**locEntityName:** `*newscatchergo.LocEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**miscEntityName:** `*newscatchergo.MiscEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*newscatchergo.TitleSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*newscatchergo.TitleSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*newscatchergo.ContentSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*newscatchergo.ContentSentimentMax` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BreakingNews.Post(request) -> *newscatchergo.BreakingNewsResponseDto</code></summary>
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
request := &newscatchergo.PostBreakingNewsRequest{
        SortBy: newscatchergo.SortByRelevancy.Ptr(),
        RankedOnly: newscatchergo.Bool(
            true,
        ),
        TopNArticles: newscatchergo.Int(
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

**sortBy:** `*newscatchergo.SortBy` 
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*newscatchergo.RankedOnly` 
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*newscatchergo.FromRank` 
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*newscatchergo.ToRank` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*newscatchergo.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*newscatchergo.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**topNArticles:** `*newscatchergo.TopNArticles` 
    
</dd>
</dl>

<dl>
<dd>

**includeTranslationFields:** `*newscatchergo.IncludeTranslationFields` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*newscatchergo.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*newscatchergo.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*newscatchergo.Theme` 
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*newscatchergo.NotTheme` 
    
</dd>
</dl>

<dl>
<dd>

**orgEntityName:** `*newscatchergo.OrgEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**perEntityName:** `*newscatchergo.PerEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**locEntityName:** `*newscatchergo.LocEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**miscEntityName:** `*newscatchergo.MiscEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*newscatchergo.TitleSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*newscatchergo.TitleSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*newscatchergo.ContentSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*newscatchergo.ContentSentimentMax` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Authors
<details><summary><code>client.Authors.Get() -> *newscatchergo.GetAuthorsResponse</code></summary>
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
request := &newscatchergo.GetAuthorsRequest{
        AuthorName: "Jane Smith",
        NotAuthorName: newscatchergo.String(
            "John Doe, Jane Doe",
        ),
        PredefinedSources: newscatchergo.String(
            "top 50 US, top 20 GB",
        ),
        Sources: newscatchergo.String(
            "nytimes.com,finance.yahoo.com",
        ),
        NotSources: newscatchergo.String(
            "cnn.com,wsj.com",
        ),
        Lang: newscatchergo.String(
            "en,es",
        ),
        NotLang: newscatchergo.String(
            "fr,de",
        ),
        Countries: newscatchergo.String(
            "US,CA",
        ),
        NotCountries: newscatchergo.String(
            "UK,FR",
        ),
        From: &newscatchergo.From{
            DateTime: newscatchergo.MustParseDateTime(
                "1 day ago",
            ),
        },
        To: &newscatchergo.To{
            String: "now",
        },
        PublishedDatePrecision: newscatchergo.String(
            "full",
        ),
        ByParseDate: newscatchergo.Bool(
            true,
        ),
        RankedOnly: newscatchergo.Bool(
            true,
        ),
        FromRank: newscatchergo.Int(
            100,
        ),
        ToRank: newscatchergo.Int(
            100,
        ),
        IsHeadline: newscatchergo.Bool(
            true,
        ),
        IsOpinion: newscatchergo.Bool(
            true,
        ),
        IsPaidContent: newscatchergo.Bool(
            false,
        ),
        ParentUrl: newscatchergo.String(
            "wsj.com/politics,wsj.com/tech",
        ),
        AllLinks: newscatchergo.String(
            "https://aiindex.stanford.edu/report,https://www.stateof.ai",
        ),
        AllDomainLinks: newscatchergo.String(
            "who.int,nih.gov",
        ),
        AllLinksText: newscatchergo.String(
            "Nvidia,Tesla",
        ),
        WordCountMin: newscatchergo.Int(
            300,
        ),
        WordCountMax: newscatchergo.Int(
            1000,
        ),
        Page: newscatchergo.Int(
            2,
        ),
        PageSize: newscatchergo.Int(
            50,
        ),
        IncludeTranslationFields: newscatchergo.Bool(
            true,
        ),
        IncludeNlpData: newscatchergo.Bool(
            true,
        ),
        HasNlp: newscatchergo.Bool(
            true,
        ),
        Theme: newscatchergo.String(
            "Finance,Tech",
        ),
        NotTheme: newscatchergo.String(
            "Crime,Sports",
        ),
        NerName: newscatchergo.String(
            "Tesla,Amazon",
        ),
        TitleSentimentMin: newscatchergo.Float64(
            -0.5,
        ),
        TitleSentimentMax: newscatchergo.Float64(
            0.5,
        ),
        ContentSentimentMin: newscatchergo.Float64(
            -0.5,
        ),
        ContentSentimentMax: newscatchergo.Float64(
            0.5,
        ),
        IptcTags: newscatchergo.String(
            "20000199,20000209",
        ),
        NotIptcTags: newscatchergo.String(
            "20000205,20000209",
        ),
        IabTags: newscatchergo.String(
            "Business,Events",
        ),
        NotIabTags: newscatchergo.String(
            "Agriculture,Metals",
        ),
        CustomTags: newscatchergo.String(
            "Tag1,Tag2",
        ),
        RobotsCompliant: newscatchergo.Bool(
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

**authorName:** `newscatchergo.AuthorName` 
    
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

**from:** `*newscatchergo.From` 
    
</dd>
</dl>

<dl>
<dd>

**to:** `*newscatchergo.To` 
    
</dd>
</dl>

<dl>
<dd>

**publishedDatePrecision:** `*newscatchergo.PublishedDatePrecision` 
    
</dd>
</dl>

<dl>
<dd>

**byParseDate:** `*newscatchergo.ByParseDate` 
    
</dd>
</dl>

<dl>
<dd>

**sortBy:** `*newscatchergo.SortBy` 
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*newscatchergo.RankedOnly` 
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*newscatchergo.FromRank` 
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*newscatchergo.ToRank` 
    
</dd>
</dl>

<dl>
<dd>

**isHeadline:** `*newscatchergo.IsHeadline` 
    
</dd>
</dl>

<dl>
<dd>

**isOpinion:** `*newscatchergo.IsOpinion` 
    
</dd>
</dl>

<dl>
<dd>

**isPaidContent:** `*newscatchergo.IsPaidContent` 
    
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

**wordCountMin:** `*newscatchergo.WordCountMin` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMax:** `*newscatchergo.WordCountMax` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*newscatchergo.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*newscatchergo.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**includeTranslationFields:** `*newscatchergo.IncludeTranslationFields` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*newscatchergo.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*newscatchergo.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*newscatchergo.Theme` 
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*newscatchergo.NotTheme` 
    
</dd>
</dl>

<dl>
<dd>

**nerName:** `*newscatchergo.NerName` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*newscatchergo.TitleSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*newscatchergo.TitleSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*newscatchergo.ContentSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*newscatchergo.ContentSentimentMax` 
    
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

**robotsCompliant:** `*newscatchergo.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Authors.Post(request) -> *newscatchergo.PostAuthorsResponse</code></summary>
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
request := &newscatchergo.PostAuthorsRequest{
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

**authorName:** `newscatchergo.AuthorName` 
    
</dd>
</dl>

<dl>
<dd>

**notAuthorName:** `*newscatchergo.NotAuthorName` 
    
</dd>
</dl>

<dl>
<dd>

**predefinedSources:** `*newscatchergo.PredefinedSources` 
    
</dd>
</dl>

<dl>
<dd>

**sources:** `*newscatchergo.Sources` 
    
</dd>
</dl>

<dl>
<dd>

**notSources:** `*newscatchergo.NotSources` 
    
</dd>
</dl>

<dl>
<dd>

**lang:** `*newscatchergo.Lang` 
    
</dd>
</dl>

<dl>
<dd>

**notLang:** `*newscatchergo.NotLang` 
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*newscatchergo.Countries` 
    
</dd>
</dl>

<dl>
<dd>

**notCountries:** `*newscatchergo.NotCountries` 
    
</dd>
</dl>

<dl>
<dd>

**from:** `*newscatchergo.From` 
    
</dd>
</dl>

<dl>
<dd>

**to:** `*newscatchergo.To` 
    
</dd>
</dl>

<dl>
<dd>

**publishedDatePrecision:** `*newscatchergo.PublishedDatePrecision` 
    
</dd>
</dl>

<dl>
<dd>

**byParseDate:** `*newscatchergo.ByParseDate` 
    
</dd>
</dl>

<dl>
<dd>

**sortBy:** `*newscatchergo.SortBy` 
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*newscatchergo.RankedOnly` 
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*newscatchergo.FromRank` 
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*newscatchergo.ToRank` 
    
</dd>
</dl>

<dl>
<dd>

**isHeadline:** `*newscatchergo.IsHeadline` 
    
</dd>
</dl>

<dl>
<dd>

**isOpinion:** `*newscatchergo.IsOpinion` 
    
</dd>
</dl>

<dl>
<dd>

**isPaidContent:** `*newscatchergo.IsPaidContent` 
    
</dd>
</dl>

<dl>
<dd>

**parentUrl:** `*newscatchergo.ParentUrl` 
    
</dd>
</dl>

<dl>
<dd>

**allLinks:** `*newscatchergo.AllLinks` 
    
</dd>
</dl>

<dl>
<dd>

**allDomainLinks:** `*newscatchergo.AllDomainLinks` 
    
</dd>
</dl>

<dl>
<dd>

**allLinksText:** `*newscatchergo.AllLinksText` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMin:** `*newscatchergo.WordCountMin` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMax:** `*newscatchergo.WordCountMax` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*newscatchergo.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*newscatchergo.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**includeTranslationFields:** `*newscatchergo.IncludeTranslationFields` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*newscatchergo.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*newscatchergo.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*newscatchergo.Theme` 
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*newscatchergo.NotTheme` 
    
</dd>
</dl>

<dl>
<dd>

**nerName:** `*newscatchergo.NerName` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*newscatchergo.TitleSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*newscatchergo.TitleSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*newscatchergo.ContentSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*newscatchergo.ContentSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**iptcTags:** `*newscatchergo.IptcTags` 
    
</dd>
</dl>

<dl>
<dd>

**notIptcTags:** `*newscatchergo.NotIptcTags` 
    
</dd>
</dl>

<dl>
<dd>

**iabTags:** `*newscatchergo.IabTags` 
    
</dd>
</dl>

<dl>
<dd>

**notIabTags:** `*newscatchergo.NotIabTags` 
    
</dd>
</dl>

<dl>
<dd>

**customTags:** `*newscatchergo.CustomTags` 
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*newscatchergo.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## SearchByLink
<details><summary><code>client.SearchByLink.Get() -> *newscatchergo.SearchResponseDto</code></summary>
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
request := &newscatchergo.GetSearchByLinkRequest{
        Ids: newscatchergo.String(
            "5f8d0d55b6e45e00179c6e7e",
        ),
        Links: newscatchergo.String(
            "https://nytimes.com/article1,https://bbc.com/article2",
        ),
        From: &newscatchergo.From{
            DateTime: newscatchergo.MustParseDateTime(
                "1 day ago",
            ),
        },
        To: &newscatchergo.To{
            String: "now",
        },
        Page: newscatchergo.Int(
            2,
        ),
        PageSize: newscatchergo.Int(
            50,
        ),
        RobotsCompliant: newscatchergo.Bool(
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

**from:** `*newscatchergo.From` 
    
</dd>
</dl>

<dl>
<dd>

**to:** `*newscatchergo.To` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*newscatchergo.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*newscatchergo.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*newscatchergo.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.SearchByLink.Post(request) -> *newscatchergo.SearchResponseDto</code></summary>
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
request := &newscatchergo.PostSearchByLinkRequest{
        Links: &newscatchergo.Links{
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

**ids:** `*newscatchergo.Ids` 
    
</dd>
</dl>

<dl>
<dd>

**links:** `*newscatchergo.Links` 
    
</dd>
</dl>

<dl>
<dd>

**from:** `*newscatchergo.From` 
    
</dd>
</dl>

<dl>
<dd>

**to:** `*newscatchergo.To` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*newscatchergo.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*newscatchergo.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*newscatchergo.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Sources
<details><summary><code>client.Sources.Get() -> *newscatchergo.SourcesResponseDto</code></summary>
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
request := &newscatchergo.GetSourcesRequest{
        Lang: newscatchergo.String(
            "en,es",
        ),
        Countries: newscatchergo.String(
            "US,CA",
        ),
        PredefinedSources: newscatchergo.String(
            "top 50 US, top 20 GB",
        ),
        SourceName: newscatchergo.String(
            "sport,tech",
        ),
        SourceUrl: newscatchergo.String(
            "bbc.com",
        ),
        IncludeAdditionalInfo: newscatchergo.Bool(
            true,
        ),
        IsNewsDomain: newscatchergo.Bool(
            true,
        ),
        NewsType: newscatchergo.String(
            "General News Outlets,Tech News and Updates",
        ),
        FromRank: newscatchergo.Int(
            100,
        ),
        ToRank: newscatchergo.Int(
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

**Note**: The search doesn't require an exact match and returns sources containing the specified terms in their names. You can use any word or phrase, like `"sport"` or `"new york times"`. 

For example, `"sport"` returns sources such as `"Motorsport"`, `"Dot Esport"`, and `"Tuttosport"`.
    
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

**includeAdditionalInfo:** `*newscatchergo.IncludeAdditionalInfo` 
    
</dd>
</dl>

<dl>
<dd>

**isNewsDomain:** `*newscatchergo.IsNewsDomain` 
    
</dd>
</dl>

<dl>
<dd>

**newsDomainType:** `*newscatchergo.NewsDomainType` 
    
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

**fromRank:** `*newscatchergo.FromRank` 
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*newscatchergo.ToRank` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Sources.Post(request) -> *newscatchergo.SourcesResponseDto</code></summary>
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
request := &newscatchergo.PostSourcesRequest{
        PredefinedSources: &newscatchergo.PredefinedSources{
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

**lang:** `*newscatchergo.Lang` 
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*newscatchergo.Countries` 
    
</dd>
</dl>

<dl>
<dd>

**predefinedSources:** `*newscatchergo.PredefinedSources` 
    
</dd>
</dl>

<dl>
<dd>

**sourceName:** `*newscatchergo.SourceName` 
    
</dd>
</dl>

<dl>
<dd>

**sourceUrl:** `*newscatchergo.SourceUrl` 
    
</dd>
</dl>

<dl>
<dd>

**includeAdditionalInfo:** `*newscatchergo.IncludeAdditionalInfo` 
    
</dd>
</dl>

<dl>
<dd>

**isNewsDomain:** `*newscatchergo.IsNewsDomain` 
    
</dd>
</dl>

<dl>
<dd>

**newsDomainType:** `*newscatchergo.NewsDomainType` 
    
</dd>
</dl>

<dl>
<dd>

**newsType:** `*newscatchergo.NewsType` 
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*newscatchergo.FromRank` 
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*newscatchergo.ToRank` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## AggregationCount
<details><summary><code>client.AggregationCount.Get() -> *newscatchergo.GetAggregationCountResponse</code></summary>
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
request := &newscatchergo.GetAggregationCountRequest{
        Q: `"supply chain" AND Amazon NOT China`,
        SearchIn: newscatchergo.String(
            "title_content, title_content_translated",
        ),
        PredefinedSources: newscatchergo.String(
            "top 50 US, top 20 GB",
        ),
        Sources: newscatchergo.String(
            "nytimes.com,finance.yahoo.com",
        ),
        NotSources: newscatchergo.String(
            "cnn.com,wsj.com",
        ),
        Lang: newscatchergo.String(
            "en,es",
        ),
        NotLang: newscatchergo.String(
            "fr,de",
        ),
        Countries: newscatchergo.String(
            "US,CA",
        ),
        NotCountries: newscatchergo.String(
            "UK,FR",
        ),
        NotAuthorName: newscatchergo.String(
            "John Doe, Jane Doe",
        ),
        From: &newscatchergo.From{
            DateTime: newscatchergo.MustParseDateTime(
                "1 day ago",
            ),
        },
        To: &newscatchergo.To{
            String: "now",
        },
        PublishedDatePrecision: newscatchergo.String(
            "full",
        ),
        ByParseDate: newscatchergo.Bool(
            true,
        ),
        RankedOnly: newscatchergo.Bool(
            true,
        ),
        FromRank: newscatchergo.Int(
            100,
        ),
        ToRank: newscatchergo.Int(
            100,
        ),
        IsHeadline: newscatchergo.Bool(
            true,
        ),
        IsOpinion: newscatchergo.Bool(
            true,
        ),
        IsPaidContent: newscatchergo.Bool(
            false,
        ),
        ParentUrl: newscatchergo.String(
            "wsj.com/politics,wsj.com/tech",
        ),
        AllLinks: newscatchergo.String(
            "https://aiindex.stanford.edu/report,https://www.stateof.ai",
        ),
        AllDomainLinks: newscatchergo.String(
            "who.int,nih.gov",
        ),
        AllLinksText: newscatchergo.String(
            "Nvidia,Tesla",
        ),
        WordCountMin: newscatchergo.Int(
            300,
        ),
        WordCountMax: newscatchergo.Int(
            1000,
        ),
        Page: newscatchergo.Int(
            2,
        ),
        PageSize: newscatchergo.Int(
            50,
        ),
        IncludeNlpData: newscatchergo.Bool(
            true,
        ),
        HasNlp: newscatchergo.Bool(
            true,
        ),
        Theme: newscatchergo.String(
            "Finance,Tech",
        ),
        NotTheme: newscatchergo.String(
            "Crime,Sports",
        ),
        OrgEntityName: newscatchergo.String(
            `"Apple Inc" OR Microsoft`,
        ),
        PerEntityName: newscatchergo.String(
            `"Elon Musk" OR "Jeff Bezos"`,
        ),
        LocEntityName: newscatchergo.String(
            `"San Francisco" OR "New York City"`,
        ),
        MiscEntityName: newscatchergo.String(
            `AWS OR "Microsoft Azure"`,
        ),
        TitleSentimentMin: newscatchergo.Float64(
            -0.5,
        ),
        TitleSentimentMax: newscatchergo.Float64(
            0.5,
        ),
        ContentSentimentMin: newscatchergo.Float64(
            -0.5,
        ),
        ContentSentimentMax: newscatchergo.Float64(
            0.5,
        ),
        IptcTags: newscatchergo.String(
            "20000199,20000209",
        ),
        NotIptcTags: newscatchergo.String(
            "20000205,20000209",
        ),
        RobotsCompliant: newscatchergo.Bool(
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

**q:** `newscatchergo.Q` 
    
</dd>
</dl>

<dl>
<dd>

**aggregationBy:** `*newscatchergo.AggregationBy` 
    
</dd>
</dl>

<dl>
<dd>

**searchIn:** `*newscatchergo.SearchIn` 
    
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

**from:** `*newscatchergo.From` 
    
</dd>
</dl>

<dl>
<dd>

**to:** `*newscatchergo.To` 
    
</dd>
</dl>

<dl>
<dd>

**publishedDatePrecision:** `*newscatchergo.PublishedDatePrecision` 
    
</dd>
</dl>

<dl>
<dd>

**byParseDate:** `*newscatchergo.ByParseDate` 
    
</dd>
</dl>

<dl>
<dd>

**sortBy:** `*newscatchergo.SortBy` 
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*newscatchergo.RankedOnly` 
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*newscatchergo.FromRank` 
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*newscatchergo.ToRank` 
    
</dd>
</dl>

<dl>
<dd>

**isHeadline:** `*newscatchergo.IsHeadline` 
    
</dd>
</dl>

<dl>
<dd>

**isOpinion:** `*newscatchergo.IsOpinion` 
    
</dd>
</dl>

<dl>
<dd>

**isPaidContent:** `*newscatchergo.IsPaidContent` 
    
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

**wordCountMin:** `*newscatchergo.WordCountMin` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMax:** `*newscatchergo.WordCountMax` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*newscatchergo.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*newscatchergo.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*newscatchergo.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*newscatchergo.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*newscatchergo.Theme` 
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*newscatchergo.NotTheme` 
    
</dd>
</dl>

<dl>
<dd>

**orgEntityName:** `*newscatchergo.OrgEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**perEntityName:** `*newscatchergo.PerEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**locEntityName:** `*newscatchergo.LocEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**miscEntityName:** `*newscatchergo.MiscEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*newscatchergo.TitleSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*newscatchergo.TitleSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*newscatchergo.ContentSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*newscatchergo.ContentSentimentMax` 
    
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

**robotsCompliant:** `*newscatchergo.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.AggregationCount.Post(request) -> *newscatchergo.PostAggregationCountResponse</code></summary>
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
request := &newscatchergo.PostAggregationCountRequest{
        Q: `"supply chain" AND Amazon NOT China`,
        AggregationBy: newscatchergo.AggregationByDay.Ptr(),
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

**q:** `newscatchergo.Q` 
    
</dd>
</dl>

<dl>
<dd>

**aggregationBy:** `*newscatchergo.AggregationBy` 
    
</dd>
</dl>

<dl>
<dd>

**searchIn:** `*newscatchergo.SearchIn` 
    
</dd>
</dl>

<dl>
<dd>

**predefinedSources:** `*newscatchergo.PredefinedSources` 
    
</dd>
</dl>

<dl>
<dd>

**sources:** `*newscatchergo.Sources` 
    
</dd>
</dl>

<dl>
<dd>

**notSources:** `*newscatchergo.NotSources` 
    
</dd>
</dl>

<dl>
<dd>

**lang:** `*newscatchergo.Lang` 
    
</dd>
</dl>

<dl>
<dd>

**notLang:** `*newscatchergo.NotLang` 
    
</dd>
</dl>

<dl>
<dd>

**countries:** `*newscatchergo.Countries` 
    
</dd>
</dl>

<dl>
<dd>

**notCountries:** `*newscatchergo.NotCountries` 
    
</dd>
</dl>

<dl>
<dd>

**notAuthorName:** `*newscatchergo.NotAuthorName` 
    
</dd>
</dl>

<dl>
<dd>

**from:** `*newscatchergo.From` 
    
</dd>
</dl>

<dl>
<dd>

**to:** `*newscatchergo.To` 
    
</dd>
</dl>

<dl>
<dd>

**publishedDatePrecision:** `*newscatchergo.PublishedDatePrecision` 
    
</dd>
</dl>

<dl>
<dd>

**byParseDate:** `*newscatchergo.ByParseDate` 
    
</dd>
</dl>

<dl>
<dd>

**sortBy:** `*newscatchergo.SortBy` 
    
</dd>
</dl>

<dl>
<dd>

**rankedOnly:** `*newscatchergo.RankedOnly` 
    
</dd>
</dl>

<dl>
<dd>

**fromRank:** `*newscatchergo.FromRank` 
    
</dd>
</dl>

<dl>
<dd>

**toRank:** `*newscatchergo.ToRank` 
    
</dd>
</dl>

<dl>
<dd>

**isHeadline:** `*newscatchergo.IsHeadline` 
    
</dd>
</dl>

<dl>
<dd>

**isOpinion:** `*newscatchergo.IsOpinion` 
    
</dd>
</dl>

<dl>
<dd>

**isPaidContent:** `*newscatchergo.IsPaidContent` 
    
</dd>
</dl>

<dl>
<dd>

**parentUrl:** `*newscatchergo.ParentUrl` 
    
</dd>
</dl>

<dl>
<dd>

**allLinks:** `*newscatchergo.AllLinks` 
    
</dd>
</dl>

<dl>
<dd>

**allDomainLinks:** `*newscatchergo.AllDomainLinks` 
    
</dd>
</dl>

<dl>
<dd>

**allLinksText:** `*newscatchergo.AllLinksText` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMin:** `*newscatchergo.WordCountMin` 
    
</dd>
</dl>

<dl>
<dd>

**wordCountMax:** `*newscatchergo.WordCountMax` 
    
</dd>
</dl>

<dl>
<dd>

**page:** `*newscatchergo.Page` 
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*newscatchergo.PageSize` 
    
</dd>
</dl>

<dl>
<dd>

**includeNlpData:** `*newscatchergo.IncludeNlpData` 
    
</dd>
</dl>

<dl>
<dd>

**hasNlp:** `*newscatchergo.HasNlp` 
    
</dd>
</dl>

<dl>
<dd>

**theme:** `*newscatchergo.Theme` 
    
</dd>
</dl>

<dl>
<dd>

**notTheme:** `*newscatchergo.NotTheme` 
    
</dd>
</dl>

<dl>
<dd>

**orgEntityName:** `*newscatchergo.OrgEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**perEntityName:** `*newscatchergo.PerEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**locEntityName:** `*newscatchergo.LocEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**miscEntityName:** `*newscatchergo.MiscEntityName` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMin:** `*newscatchergo.TitleSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**titleSentimentMax:** `*newscatchergo.TitleSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMin:** `*newscatchergo.ContentSentimentMin` 
    
</dd>
</dl>

<dl>
<dd>

**contentSentimentMax:** `*newscatchergo.ContentSentimentMax` 
    
</dd>
</dl>

<dl>
<dd>

**iptcTags:** `*newscatchergo.IptcTags` 
    
</dd>
</dl>

<dl>
<dd>

**notIptcTags:** `*newscatchergo.NotIptcTags` 
    
</dd>
</dl>

<dl>
<dd>

**robotsCompliant:** `*newscatchergo.RobotsCompliant` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Subscription
<details><summary><code>client.Subscription.Get() -> *newscatchergo.SubscriptionResponseDto</code></summary>
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

<details><summary><code>client.Subscription.Post() -> *newscatchergo.SubscriptionResponseDto</code></summary>
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

