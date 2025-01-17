// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
)

type SearchUrlPostRequest struct {
	Ids   *Ids   `json:"ids,omitempty" url:"-"`
	Links *Links `json:"links,omitempty" url:"-"`
	// The starting point in time to search from. Accepts date-time strings in ISO 8601 format and plain text strings. The default time zone is UTC.
	//
	// Formats with examples:
	// - YYYY-mm-ddTHH:MM:SS: `2024-07-01T00:00:00`
	// - YYYY-MM-dd: `2024-07-01`
	// - YYYY/mm/dd HH:MM:SS: `2024/07/01 00:00:00`
	// - YYYY/mm/dd: `2024/07/01`
	// - English phrases: `1 day ago`, `today`
	From *From `json:"from_,omitempty" url:"-"`
	// The ending point in time to search up to. Accepts date-time strings in ISO 8601 format and plain text strings. The default time zone is UTC.
	//
	// Formats with examples:
	// - YYYY-mm-ddTHH:MM:SS: `2024-07-01T00:00:00`
	// - YYYY-MM-dd: `2024-07-01`
	// - YYYY/mm/dd HH:MM:SS: `2024/07/01 00:00:00`
	// - YYYY/mm/dd: `2024/07/01`
	// - English phrases: `1 day ago`, `today`
	To       *To       `json:"to_,omitempty" url:"-"`
	Page     *Page     `json:"page,omitempty" url:"-"`
	PageSize *PageSize `json:"page_size,omitempty" url:"-"`
}

type SearchUrlGetRequest struct {
	// The Newscatcher article ID (corresponds to the `_id` field in API response) or a list of article IDs to search for. To specify multiple IDs, use a comma-separated string.
	//
	// Example: `"1234567890abcdef, abcdef1234567890"`
	//
	// **Caution**: You can use either the `links` or the `ids` parameter, but not both at the same time.
	Ids *string `json:"-" url:"ids,omitempty"`
	// The article link or list of article links to search for. To specify multiple links, use a comma-separated string.
	//
	// Example: `"https://example.com/article1, https://example.com/article2"`
	//
	// **Caution**: You can use either the `links` or the `ids` parameter, but not both at the same time.
	Links *string `json:"-" url:"links,omitempty"`
	From  *From   `json:"-" url:"from_,omitempty"`
	To    *To     `json:"-" url:"to_,omitempty"`
	// The page number to scroll through the results. Use for pagination, as a single API response can return up to 1,000 articles.
	//
	// For details, see [How to paginate large datasets](https://www.newscatcherapi.com/docs/v3/documentation/how-to/paginate-large-datasets).
	Page *int `json:"-" url:"page,omitempty"`
	// The number of articles to return per page.
	PageSize *int `json:"-" url:"page_size,omitempty"`
}

// The Newscatcher article ID (corresponds to the `_id` field in API response) or a list of article IDs to search for. To specify multiple IDs, use a comma-separated string or an array of strings.
//
// Examples:
// - `"5f8d0d55b6e45e00179c6e7e,5f8d0d55b6e45e00179c6e7f"`
// - `["5f8d0d55b6e45e00179c6e7e","5f8d0d55b6e45e00179c6e7f"]`
//
// **Caution**: You can use either the `links` or the `ids` parameter, but not both at the same time.
type Ids struct {
	String     string
	StringList []string

	typ string
}

func NewIdsFromString(value string) *Ids {
	return &Ids{typ: "String", String: value}
}

func NewIdsFromStringList(value []string) *Ids {
	return &Ids{typ: "StringList", StringList: value}
}

func (i *Ids) GetString() string {
	if i == nil {
		return ""
	}
	return i.String
}

func (i *Ids) GetStringList() []string {
	if i == nil {
		return nil
	}
	return i.StringList
}

func (i *Ids) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		i.typ = "String"
		i.String = valueString
		return nil
	}
	var valueStringList []string
	if err := json.Unmarshal(data, &valueStringList); err == nil {
		i.typ = "StringList"
		i.StringList = valueStringList
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, i)
}

func (i Ids) MarshalJSON() ([]byte, error) {
	if i.typ == "String" || i.String != "" {
		return json.Marshal(i.String)
	}
	if i.typ == "StringList" || i.StringList != nil {
		return json.Marshal(i.StringList)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", i)
}

type IdsVisitor interface {
	VisitString(string) error
	VisitStringList([]string) error
}

func (i *Ids) Accept(visitor IdsVisitor) error {
	if i.typ == "String" || i.String != "" {
		return visitor.VisitString(i.String)
	}
	if i.typ == "StringList" || i.StringList != nil {
		return visitor.VisitStringList(i.StringList)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", i)
}

// The article link or list of article links to search for. To specify multiple links, use a comma-separated string or an array of strings.
//
// Examples:
// - `"https://nytimes.com/article1, https://bbc.com/article2"`
// - `["https://nytimes.com/article1", "https://bbc.com/article2"]`
//
// **Caution**: You can use either the `links` or the `ids` parameter, but not both at the same time.
type Links struct {
	String     string
	StringList []string

	typ string
}

func NewLinksFromString(value string) *Links {
	return &Links{typ: "String", String: value}
}

func NewLinksFromStringList(value []string) *Links {
	return &Links{typ: "StringList", StringList: value}
}

func (l *Links) GetString() string {
	if l == nil {
		return ""
	}
	return l.String
}

func (l *Links) GetStringList() []string {
	if l == nil {
		return nil
	}
	return l.StringList
}

func (l *Links) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		l.typ = "String"
		l.String = valueString
		return nil
	}
	var valueStringList []string
	if err := json.Unmarshal(data, &valueStringList); err == nil {
		l.typ = "StringList"
		l.StringList = valueStringList
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, l)
}

func (l Links) MarshalJSON() ([]byte, error) {
	if l.typ == "String" || l.String != "" {
		return json.Marshal(l.String)
	}
	if l.typ == "StringList" || l.StringList != nil {
		return json.Marshal(l.StringList)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", l)
}

type LinksVisitor interface {
	VisitString(string) error
	VisitStringList([]string) error
}

func (l *Links) Accept(visitor LinksVisitor) error {
	if l.typ == "String" || l.String != "" {
		return visitor.VisitString(l.String)
	}
	if l.typ == "StringList" || l.StringList != nil {
		return visitor.VisitStringList(l.StringList)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", l)
}
