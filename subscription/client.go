// This file was auto-generated by Fern from our API Definition.

package subscription

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

// Retrieves information about your subscription plan.
func (c *Client) Get(
	ctx context.Context,
	opts ...option.RequestOption,
) (*newscatchergo.SubscriptionResponseDto, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://v3-api.newscatcherapi.com",
	)
	endpointURL := baseURL + "/api/subscription"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &newscatchergo.BadRequestError{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &newscatchergo.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &newscatchergo.ForbiddenError{
				APIError: apiError,
			}
		},
		408: func(apiError *core.APIError) error {
			return &newscatchergo.RequestTimeoutError{
				APIError: apiError,
			}
		},
		422: func(apiError *core.APIError) error {
			return &newscatchergo.UnprocessableEntityError{
				APIError: apiError,
			}
		},
		429: func(apiError *core.APIError) error {
			return &newscatchergo.TooManyRequestsError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &newscatchergo.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *newscatchergo.SubscriptionResponseDto
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

// Retrieves information about your subscription plan.
func (c *Client) Post(
	ctx context.Context,
	opts ...option.RequestOption,
) (*newscatchergo.SubscriptionResponseDto, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://v3-api.newscatcherapi.com",
	)
	endpointURL := baseURL + "/api/subscription"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &newscatchergo.BadRequestError{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &newscatchergo.UnauthorizedError{
				APIError: apiError,
			}
		},
		403: func(apiError *core.APIError) error {
			return &newscatchergo.ForbiddenError{
				APIError: apiError,
			}
		},
		408: func(apiError *core.APIError) error {
			return &newscatchergo.RequestTimeoutError{
				APIError: apiError,
			}
		},
		422: func(apiError *core.APIError) error {
			return &newscatchergo.UnprocessableEntityError{
				APIError: apiError,
			}
		},
		429: func(apiError *core.APIError) error {
			return &newscatchergo.TooManyRequestsError{
				APIError: apiError,
			}
		},
		500: func(apiError *core.APIError) error {
			return &newscatchergo.InternalServerError{
				APIError: apiError,
			}
		},
	}

	var response *newscatchergo.SubscriptionResponseDto
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
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
