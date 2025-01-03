// This file was auto-generated by Fern from our API Definition.

package searchsimilar

import (
	context "context"
	newscatchergo "github.com/Newscatcher/newscatcher-go"
	core "github.com/Newscatcher/newscatcher-go/core"
	internal "github.com/Newscatcher/newscatcher-go/internal"
	option "github.com/Newscatcher/newscatcher-go/option"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

// This endpoint returns a list of articles that are similar to the query provided. You also have the option to get similar articles for the results of a search.
func (c *Client) Get(
	ctx context.Context,
	request *newscatchergo.SearchSimilarGetRequest,
	opts ...option.RequestOption,
) (*newscatchergo.SearchSimilarGetResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://v3-api.newscatcherapi.com",
	)
	endpointURL := baseURL + "/api/search_similar"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		422: func(apiError *core.APIError) error {
			return &newscatchergo.UnprocessableEntityError{
				APIError: apiError,
			}
		},
	}

	var response *newscatchergo.SearchSimilarGetResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// This endpoint returns a list of articles that are similar to the query provided. You also have the option to get similar articles for the results of a search.
func (c *Client) Post(
	ctx context.Context,
	request *newscatchergo.MoreLikeThisRequest,
	opts ...option.RequestOption,
) (*newscatchergo.SearchSimilarPostResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://v3-api.newscatcherapi.com",
	)
	endpointURL := baseURL + "/api/search_similar"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")
	errorCodes := internal.ErrorCodes{
		422: func(apiError *core.APIError) error {
			return &newscatchergo.UnprocessableEntityError{
				APIError: apiError,
			}
		},
	}

	var response *newscatchergo.SearchSimilarPostResponse
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
