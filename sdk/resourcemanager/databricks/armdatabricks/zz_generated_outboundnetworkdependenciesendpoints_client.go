//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatabricks

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

// OutboundNetworkDependenciesEndpointsClient contains the methods for the OutboundNetworkDependenciesEndpoints group.
// Don't use this type directly, use NewOutboundNetworkDependenciesEndpointsClient() instead.
type OutboundNetworkDependenciesEndpointsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewOutboundNetworkDependenciesEndpointsClient creates a new instance of OutboundNetworkDependenciesEndpointsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOutboundNetworkDependenciesEndpointsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *OutboundNetworkDependenciesEndpointsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &OutboundNetworkDependenciesEndpointsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// List - Gets the list of endpoints that VNET Injected Workspace calls Azure Databricks Control Plane. You must configure
// outbound access with these endpoints. For more information, see
// https://docs.microsoft.com/en-us/azure/databricks/administration-guide/cloud-configurations/azure/udr
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// options - OutboundNetworkDependenciesEndpointsClientListOptions contains the optional parameters for the OutboundNetworkDependenciesEndpointsClient.List
// method.
func (client *OutboundNetworkDependenciesEndpointsClient) List(ctx context.Context, resourceGroupName string, workspaceName string, options *OutboundNetworkDependenciesEndpointsClientListOptions) (OutboundNetworkDependenciesEndpointsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return OutboundNetworkDependenciesEndpointsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OutboundNetworkDependenciesEndpointsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OutboundNetworkDependenciesEndpointsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *OutboundNetworkDependenciesEndpointsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *OutboundNetworkDependenciesEndpointsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Databricks/workspaces/{workspaceName}/outboundNetworkDependenciesEndpoints"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OutboundNetworkDependenciesEndpointsClient) listHandleResponse(resp *http.Response) (OutboundNetworkDependenciesEndpointsClientListResponse, error) {
	result := OutboundNetworkDependenciesEndpointsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.OutboundEnvironmentEndpointArray); err != nil {
		return OutboundNetworkDependenciesEndpointsClientListResponse{}, err
	}
	return result, nil
}
