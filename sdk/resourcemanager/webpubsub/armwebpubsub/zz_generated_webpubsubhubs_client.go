//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armwebpubsub

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// WebPubSubHubsClient contains the methods for the WebPubSubHubs group.
// Don't use this type directly, use NewWebPubSubHubsClient() instead.
type WebPubSubHubsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewWebPubSubHubsClient creates a new instance of WebPubSubHubsClient with the specified values.
func NewWebPubSubHubsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *WebPubSubHubsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &WebPubSubHubsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - Create or update a hub setting.
// If the operation fails it returns the *ErrorResponse error type.
func (client *WebPubSubHubsClient) BeginCreateOrUpdate(ctx context.Context, hubName string, resourceGroupName string, resourceName string, parameters WebPubSubHub, options *WebPubSubHubsBeginCreateOrUpdateOptions) (WebPubSubHubsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, hubName, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return WebPubSubHubsCreateOrUpdatePollerResponse{}, err
	}
	result := WebPubSubHubsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("WebPubSubHubsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return WebPubSubHubsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &WebPubSubHubsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create or update a hub setting.
// If the operation fails it returns the *ErrorResponse error type.
func (client *WebPubSubHubsClient) createOrUpdate(ctx context.Context, hubName string, resourceGroupName string, resourceName string, parameters WebPubSubHub, options *WebPubSubHubsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, hubName, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WebPubSubHubsClient) createOrUpdateCreateRequest(ctx context.Context, hubName string, resourceGroupName string, resourceName string, parameters WebPubSubHub, options *WebPubSubHubsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/webPubSub/{resourceName}/hubs/{hubName}"
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *WebPubSubHubsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Delete a hub setting.
// If the operation fails it returns the *ErrorResponse error type.
func (client *WebPubSubHubsClient) BeginDelete(ctx context.Context, hubName string, resourceGroupName string, resourceName string, options *WebPubSubHubsBeginDeleteOptions) (WebPubSubHubsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, hubName, resourceGroupName, resourceName, options)
	if err != nil {
		return WebPubSubHubsDeletePollerResponse{}, err
	}
	result := WebPubSubHubsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("WebPubSubHubsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return WebPubSubHubsDeletePollerResponse{}, err
	}
	result.Poller = &WebPubSubHubsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete a hub setting.
// If the operation fails it returns the *ErrorResponse error type.
func (client *WebPubSubHubsClient) deleteOperation(ctx context.Context, hubName string, resourceGroupName string, resourceName string, options *WebPubSubHubsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, hubName, resourceGroupName, resourceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WebPubSubHubsClient) deleteCreateRequest(ctx context.Context, hubName string, resourceGroupName string, resourceName string, options *WebPubSubHubsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/webPubSub/{resourceName}/hubs/{hubName}"
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *WebPubSubHubsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get a hub setting.
// If the operation fails it returns the *ErrorResponse error type.
func (client *WebPubSubHubsClient) Get(ctx context.Context, hubName string, resourceGroupName string, resourceName string, options *WebPubSubHubsGetOptions) (WebPubSubHubsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, hubName, resourceGroupName, resourceName, options)
	if err != nil {
		return WebPubSubHubsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WebPubSubHubsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WebPubSubHubsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WebPubSubHubsClient) getCreateRequest(ctx context.Context, hubName string, resourceGroupName string, resourceName string, options *WebPubSubHubsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/webPubSub/{resourceName}/hubs/{hubName}"
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WebPubSubHubsClient) getHandleResponse(resp *http.Response) (WebPubSubHubsGetResponse, error) {
	result := WebPubSubHubsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebPubSubHub); err != nil {
		return WebPubSubHubsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *WebPubSubHubsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - List hub settings.
// If the operation fails it returns the *ErrorResponse error type.
func (client *WebPubSubHubsClient) List(resourceGroupName string, resourceName string, options *WebPubSubHubsListOptions) *WebPubSubHubsListPager {
	return &WebPubSubHubsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
		},
		advancer: func(ctx context.Context, resp WebPubSubHubsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WebPubSubHubList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *WebPubSubHubsClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *WebPubSubHubsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/webPubSub/{resourceName}/hubs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WebPubSubHubsClient) listHandleResponse(resp *http.Response) (WebPubSubHubsListResponse, error) {
	result := WebPubSubHubsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebPubSubHubList); err != nil {
		return WebPubSubHubsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *WebPubSubHubsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
