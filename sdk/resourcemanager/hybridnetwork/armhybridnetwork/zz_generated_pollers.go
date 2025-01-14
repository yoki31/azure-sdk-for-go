//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// DevicesClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type DevicesClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *DevicesClientCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *DevicesClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final DevicesClientCreateOrUpdateResponse will be returned.
func (p *DevicesClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (DevicesClientCreateOrUpdateResponse, error) {
	respType := DevicesClientCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.Device)
	if err != nil {
		return DevicesClientCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *DevicesClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// DevicesClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type DevicesClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *DevicesClientDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *DevicesClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final DevicesClientDeleteResponse will be returned.
func (p *DevicesClientDeletePoller) FinalResponse(ctx context.Context) (DevicesClientDeleteResponse, error) {
	respType := DevicesClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return DevicesClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *DevicesClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// NetworkFunctionsClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type NetworkFunctionsClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *NetworkFunctionsClientCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *NetworkFunctionsClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final NetworkFunctionsClientCreateOrUpdateResponse will be returned.
func (p *NetworkFunctionsClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (NetworkFunctionsClientCreateOrUpdateResponse, error) {
	respType := NetworkFunctionsClientCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.NetworkFunction)
	if err != nil {
		return NetworkFunctionsClientCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *NetworkFunctionsClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// NetworkFunctionsClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type NetworkFunctionsClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *NetworkFunctionsClientDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *NetworkFunctionsClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final NetworkFunctionsClientDeleteResponse will be returned.
func (p *NetworkFunctionsClientDeletePoller) FinalResponse(ctx context.Context) (NetworkFunctionsClientDeleteResponse, error) {
	respType := NetworkFunctionsClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return NetworkFunctionsClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *NetworkFunctionsClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// RoleInstancesClientRestartPoller provides polling facilities until the operation reaches a terminal state.
type RoleInstancesClientRestartPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *RoleInstancesClientRestartPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *RoleInstancesClientRestartPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final RoleInstancesClientRestartResponse will be returned.
func (p *RoleInstancesClientRestartPoller) FinalResponse(ctx context.Context) (RoleInstancesClientRestartResponse, error) {
	respType := RoleInstancesClientRestartResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return RoleInstancesClientRestartResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *RoleInstancesClientRestartPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// RoleInstancesClientStartPoller provides polling facilities until the operation reaches a terminal state.
type RoleInstancesClientStartPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *RoleInstancesClientStartPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *RoleInstancesClientStartPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final RoleInstancesClientStartResponse will be returned.
func (p *RoleInstancesClientStartPoller) FinalResponse(ctx context.Context) (RoleInstancesClientStartResponse, error) {
	respType := RoleInstancesClientStartResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return RoleInstancesClientStartResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *RoleInstancesClientStartPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// RoleInstancesClientStopPoller provides polling facilities until the operation reaches a terminal state.
type RoleInstancesClientStopPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *RoleInstancesClientStopPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *RoleInstancesClientStopPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final RoleInstancesClientStopResponse will be returned.
func (p *RoleInstancesClientStopPoller) FinalResponse(ctx context.Context) (RoleInstancesClientStopResponse, error) {
	respType := RoleInstancesClientStopResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return RoleInstancesClientStopResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *RoleInstancesClientStopPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VendorNetworkFunctionsClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type VendorNetworkFunctionsClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VendorNetworkFunctionsClientCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VendorNetworkFunctionsClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VendorNetworkFunctionsClientCreateOrUpdateResponse will be returned.
func (p *VendorNetworkFunctionsClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (VendorNetworkFunctionsClientCreateOrUpdateResponse, error) {
	respType := VendorNetworkFunctionsClientCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.VendorNetworkFunction)
	if err != nil {
		return VendorNetworkFunctionsClientCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VendorNetworkFunctionsClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VendorSKUPreviewClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type VendorSKUPreviewClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VendorSKUPreviewClientCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VendorSKUPreviewClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VendorSKUPreviewClientCreateOrUpdateResponse will be returned.
func (p *VendorSKUPreviewClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (VendorSKUPreviewClientCreateOrUpdateResponse, error) {
	respType := VendorSKUPreviewClientCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.PreviewSubscription)
	if err != nil {
		return VendorSKUPreviewClientCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VendorSKUPreviewClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VendorSKUPreviewClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type VendorSKUPreviewClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VendorSKUPreviewClientDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VendorSKUPreviewClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VendorSKUPreviewClientDeleteResponse will be returned.
func (p *VendorSKUPreviewClientDeletePoller) FinalResponse(ctx context.Context) (VendorSKUPreviewClientDeleteResponse, error) {
	respType := VendorSKUPreviewClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return VendorSKUPreviewClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VendorSKUPreviewClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VendorSKUsClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type VendorSKUsClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VendorSKUsClientCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VendorSKUsClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VendorSKUsClientCreateOrUpdateResponse will be returned.
func (p *VendorSKUsClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (VendorSKUsClientCreateOrUpdateResponse, error) {
	respType := VendorSKUsClientCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.VendorSKU)
	if err != nil {
		return VendorSKUsClientCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VendorSKUsClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VendorSKUsClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type VendorSKUsClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VendorSKUsClientDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VendorSKUsClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VendorSKUsClientDeleteResponse will be returned.
func (p *VendorSKUsClientDeletePoller) FinalResponse(ctx context.Context) (VendorSKUsClientDeleteResponse, error) {
	respType := VendorSKUsClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return VendorSKUsClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VendorSKUsClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VendorsClientCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type VendorsClientCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VendorsClientCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VendorsClientCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VendorsClientCreateOrUpdateResponse will be returned.
func (p *VendorsClientCreateOrUpdatePoller) FinalResponse(ctx context.Context) (VendorsClientCreateOrUpdateResponse, error) {
	respType := VendorsClientCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.Vendor)
	if err != nil {
		return VendorsClientCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VendorsClientCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// VendorsClientDeletePoller provides polling facilities until the operation reaches a terminal state.
type VendorsClientDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *VendorsClientDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *VendorsClientDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final VendorsClientDeleteResponse will be returned.
func (p *VendorsClientDeletePoller) FinalResponse(ctx context.Context) (VendorsClientDeleteResponse, error) {
	respType := VendorsClientDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return VendorsClientDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *VendorsClientDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}
