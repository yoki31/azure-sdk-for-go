//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmaps

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// Client contains the methods for the Maps group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewClient creates a new instance of Client with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *Client {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &Client{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// ListOperations - List operations available for the Maps Resource Provider
// If the operation fails it returns an *azcore.ResponseError type.
// options - ClientListOperationsOptions contains the optional parameters for the Client.ListOperations method.
func (client *Client) ListOperations(options *ClientListOperationsOptions) *ClientListOperationsPager {
	return &ClientListOperationsPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listOperationsCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ClientListOperationsResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.Operations.NextLink)
		},
	}
}

// listOperationsCreateRequest creates the ListOperations request.
func (client *Client) listOperationsCreateRequest(ctx context.Context, options *ClientListOperationsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Maps/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listOperationsHandleResponse handles the ListOperations response.
func (client *Client) listOperationsHandleResponse(resp *http.Response) (ClientListOperationsResponse, error) {
	result := ClientListOperationsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Operations); err != nil {
		return ClientListOperationsResponse{}, err
	}
	return result, nil
}

// ListSubscriptionOperations - List operations available for the Maps Resource Provider
// If the operation fails it returns an *azcore.ResponseError type.
// options - ClientListSubscriptionOperationsOptions contains the optional parameters for the Client.ListSubscriptionOperations
// method.
func (client *Client) ListSubscriptionOperations(options *ClientListSubscriptionOperationsOptions) *ClientListSubscriptionOperationsPager {
	return &ClientListSubscriptionOperationsPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listSubscriptionOperationsCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ClientListSubscriptionOperationsResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.Operations.NextLink)
		},
	}
}

// listSubscriptionOperationsCreateRequest creates the ListSubscriptionOperations request.
func (client *Client) listSubscriptionOperationsCreateRequest(ctx context.Context, options *ClientListSubscriptionOperationsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Maps/operations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listSubscriptionOperationsHandleResponse handles the ListSubscriptionOperations response.
func (client *Client) listSubscriptionOperationsHandleResponse(resp *http.Response) (ClientListSubscriptionOperationsResponse, error) {
	result := ClientListSubscriptionOperationsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Operations); err != nil {
		return ClientListSubscriptionOperationsResponse{}, err
	}
	return result, nil
}
