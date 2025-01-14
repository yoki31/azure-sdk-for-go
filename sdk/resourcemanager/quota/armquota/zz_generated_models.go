//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armquota

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"time"
)

// ClientBeginCreateOrUpdateOptions contains the optional parameters for the Client.BeginCreateOrUpdate method.
type ClientBeginCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ClientBeginUpdateOptions contains the optional parameters for the Client.BeginUpdate method.
type ClientBeginUpdateOptions struct {
	// placeholder for future optional parameters
}

// ClientGetOptions contains the optional parameters for the Client.Get method.
type ClientGetOptions struct {
	// placeholder for future optional parameters
}

// ClientListOptions contains the optional parameters for the Client.List method.
type ClientListOptions struct {
	// placeholder for future optional parameters
}

// CommonResourceProperties - Resource properties.
type CommonResourceProperties struct {
	// READ-ONLY; Resource ID
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource type. Example: "Microsoft.Quota/quotas"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// CreateGenericQuotaRequestParameters - Quota change requests information.
type CreateGenericQuotaRequestParameters struct {
	// Quota change requests.
	Value []*CurrentQuotaLimitBase `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type CreateGenericQuotaRequestParameters.
func (c CreateGenericQuotaRequestParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", c.Value)
	return json.Marshal(objectMap)
}

// CurrentQuotaLimitBase - Quota limit.
type CurrentQuotaLimitBase struct {
	// Quota properties for the specified resource, based on the API called, Quotas or Usages.
	Properties *Properties `json:"properties,omitempty"`

	// READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type CurrentQuotaLimitBase.
func (c CurrentQuotaLimitBase) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", c.ID)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "properties", c.Properties)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// CurrentUsagesBase - Resource usage.
type CurrentUsagesBase struct {
	// Usage properties for the specified resource.
	Properties *UsagesProperties `json:"properties,omitempty"`

	// READ-ONLY; The resource ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The resource name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ExceptionResponse - Error.
type ExceptionResponse struct {
	// API error details.
	Error *ServiceError `json:"error,omitempty"`
}

// LimitJSONObjectClassification provides polymorphic access to related types.
// Call the interface's GetLimitJSONObject() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *LimitJSONObject, *LimitObject
type LimitJSONObjectClassification interface {
	// GetLimitJSONObject returns the LimitJSONObject content of the underlying type.
	GetLimitJSONObject() *LimitJSONObject
}

// LimitJSONObject - LimitJson abstract class.
type LimitJSONObject struct {
	// REQUIRED; The limit object type.
	LimitObjectType *LimitType `json:"limitObjectType,omitempty"`
}

// GetLimitJSONObject implements the LimitJSONObjectClassification interface for type LimitJSONObject.
func (l *LimitJSONObject) GetLimitJSONObject() *LimitJSONObject { return l }

// LimitObject - The resource quota limit value.
type LimitObject struct {
	// REQUIRED; The limit object type.
	LimitObjectType *LimitType `json:"limitObjectType,omitempty"`

	// REQUIRED; The quota/limit value
	Value *int32 `json:"value,omitempty"`

	// The quota or usages limit types.
	LimitType *QuotaLimitTypes `json:"limitType,omitempty"`
}

// GetLimitJSONObject implements the LimitJSONObjectClassification interface for type LimitObject.
func (l *LimitObject) GetLimitJSONObject() *LimitJSONObject {
	return &LimitJSONObject{
		LimitObjectType: l.LimitObjectType,
	}
}

// MarshalJSON implements the json.Marshaller interface for type LimitObject.
func (l LimitObject) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["limitObjectType"] = LimitTypeLimitValue
	populate(objectMap, "limitType", l.LimitType)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LimitObject.
func (l *LimitObject) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "limitObjectType":
			err = unpopulate(val, &l.LimitObjectType)
			delete(rawMsg, key)
		case "limitType":
			err = unpopulate(val, &l.LimitType)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, &l.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// Limits - Quota limits.
type Limits struct {
	// The URI used to fetch the next page of quota limits. When there are no more pages, this string is null.
	NextLink *string `json:"nextLink,omitempty"`

	// List of quota limits.
	Value []*CurrentQuotaLimitBase `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type Limits.
func (l Limits) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", l.NextLink)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// LimitsResponse - Quota limits request response.
type LimitsResponse struct {
	// The URI used to fetch the next page of quota limits. When there are no more pages, this is null.
	NextLink *string `json:"nextLink,omitempty"`

	// List of quota limits with the quota request status.
	Value []*CurrentQuotaLimitBase `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type LimitsResponse.
func (l LimitsResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", l.NextLink)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// OperationClientListOptions contains the optional parameters for the OperationClient.List method.
type OperationClientListOptions struct {
	// placeholder for future optional parameters
}

type OperationDisplay struct {
	// Operation description.
	Description *string `json:"description,omitempty"`

	// Operation name.
	Operation *string `json:"operation,omitempty"`

	// Provider name.
	Provider *string `json:"provider,omitempty"`

	// Resource name.
	Resource *string `json:"resource,omitempty"`
}

type OperationList struct {
	// URL to get the next page of items.
	NextLink *string              `json:"nextLink,omitempty"`
	Value    []*OperationResponse `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type OperationList.
func (o OperationList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

type OperationResponse struct {
	Display *OperationDisplay `json:"display,omitempty"`
	Name    *string           `json:"name,omitempty"`
	Origin  *string           `json:"origin,omitempty"`
}

// Properties - Quota properties for the specified resource.
type Properties struct {
	// Resource quota limit properties.
	Limit LimitJSONObjectClassification `json:"limit,omitempty"`

	// Resource name provided by the resource provider. Use this property name when requesting quota.
	Name *ResourceName `json:"name,omitempty"`

	// Additional properties for the specific resource provider.
	Properties map[string]interface{} `json:"properties,omitempty"`

	// Resource type name.
	ResourceType *string `json:"resourceType,omitempty"`

	// READ-ONLY; States if quota can be requested for this resource.
	IsQuotaApplicable *bool `json:"isQuotaApplicable,omitempty" azure:"ro"`

	// READ-ONLY; The time period over which the quota usage values are summarized. For example: *P1D (per one day) *PT1M (per
	// one minute) *PT1S (per one second). This parameter is optional because, for some resources
	// like compute, the period is irrelevant.
	QuotaPeriod *string `json:"quotaPeriod,omitempty" azure:"ro"`

	// READ-ONLY; The quota units, such as Count and Bytes. When requesting quota, use the unit value returned in the GET response
	// in the request body of your PUT operation.
	Unit *string `json:"unit,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type Properties.
func (p Properties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "isQuotaApplicable", p.IsQuotaApplicable)
	populate(objectMap, "limit", p.Limit)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "properties", p.Properties)
	populate(objectMap, "quotaPeriod", p.QuotaPeriod)
	populate(objectMap, "resourceType", p.ResourceType)
	populate(objectMap, "unit", p.Unit)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Properties.
func (p *Properties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "isQuotaApplicable":
			err = unpopulate(val, &p.IsQuotaApplicable)
			delete(rawMsg, key)
		case "limit":
			p.Limit, err = unmarshalLimitJSONObjectClassification(val)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &p.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, &p.Properties)
			delete(rawMsg, key)
		case "quotaPeriod":
			err = unpopulate(val, &p.QuotaPeriod)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, &p.ResourceType)
			delete(rawMsg, key)
		case "unit":
			err = unpopulate(val, &p.Unit)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// RequestDetails - List of quota requests with details.
type RequestDetails struct {
	// Quota request details.
	Properties *RequestProperties `json:"properties,omitempty"`

	// READ-ONLY; Quota request ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Quota request name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource type. "Microsoft.Quota/quotas".
	Type *string `json:"type,omitempty" azure:"ro"`
}

// RequestDetailsList - Quota request information.
type RequestDetailsList struct {
	// The URI for fetching the next page of quota limits. When there are no more pages, this string is null.
	NextLink *string `json:"nextLink,omitempty"`

	// Quota request details.
	Value []*RequestDetails `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type RequestDetailsList.
func (r RequestDetailsList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", r.NextLink)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// RequestOneResourceProperties - Quota request.
type RequestOneResourceProperties struct {
	// Error details of the quota request.
	Error *ServiceErrorDetail `json:"error,omitempty"`

	// Resource quota limit properties.
	Limit *LimitObject `json:"limit,omitempty"`

	// Resource name provided by the resource provider. Use this property name when requesting quota.
	Name *ResourceName `json:"name,omitempty"`

	// Additional properties for the specific resource provider.
	Properties map[string]interface{} `json:"properties,omitempty"`

	// Resource type name.
	ResourceType *string `json:"resourceType,omitempty"`

	// The quota limit units, such as Count and Bytes. When requesting quota, use the unit value returned in the GET response
	// in the request body of your PUT operation.
	Unit *string `json:"unit,omitempty"`

	// READ-ONLY; Usage information for the current resource.
	CurrentValue *int32 `json:"currentValue,omitempty" azure:"ro"`

	// READ-ONLY; States if quota can be requested for this resource.
	IsQuotaApplicable *bool `json:"isQuotaApplicable,omitempty" azure:"ro"`

	// READ-ONLY; User-friendly status message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; Quota request status.
	ProvisioningState *QuotaRequestState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; The time period over which the quota usage values are summarized. For example: *P1D (per one day) *PT1M (per
	// one minute) *PT1S (per one second). This parameter is optional because, for some resources
	// like compute, the period is irrelevant.
	QuotaPeriod *string `json:"quotaPeriod,omitempty" azure:"ro"`

	// READ-ONLY; Quota request submission time. The date conforms to the following ISO 8601 standard format: yyyy-MM-ddTHH:mm:ssZ.
	RequestSubmitTime *time.Time `json:"requestSubmitTime,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type RequestOneResourceProperties.
func (r RequestOneResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "currentValue", r.CurrentValue)
	populate(objectMap, "error", r.Error)
	populate(objectMap, "isQuotaApplicable", r.IsQuotaApplicable)
	populate(objectMap, "limit", r.Limit)
	populate(objectMap, "message", r.Message)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "properties", r.Properties)
	populate(objectMap, "provisioningState", r.ProvisioningState)
	populate(objectMap, "quotaPeriod", r.QuotaPeriod)
	populateTimeRFC3339(objectMap, "requestSubmitTime", r.RequestSubmitTime)
	populate(objectMap, "resourceType", r.ResourceType)
	populate(objectMap, "unit", r.Unit)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RequestOneResourceProperties.
func (r *RequestOneResourceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "currentValue":
			err = unpopulate(val, &r.CurrentValue)
			delete(rawMsg, key)
		case "error":
			err = unpopulate(val, &r.Error)
			delete(rawMsg, key)
		case "isQuotaApplicable":
			err = unpopulate(val, &r.IsQuotaApplicable)
			delete(rawMsg, key)
		case "limit":
			err = unpopulate(val, &r.Limit)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, &r.Message)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &r.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, &r.Properties)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &r.ProvisioningState)
			delete(rawMsg, key)
		case "quotaPeriod":
			err = unpopulate(val, &r.QuotaPeriod)
			delete(rawMsg, key)
		case "requestSubmitTime":
			err = unpopulateTimeRFC3339(val, &r.RequestSubmitTime)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, &r.ResourceType)
			delete(rawMsg, key)
		case "unit":
			err = unpopulate(val, &r.Unit)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// RequestOneResourceSubmitResponse - Quota request response.
type RequestOneResourceSubmitResponse struct {
	// Quota request details.
	Properties *RequestOneResourceProperties `json:"properties,omitempty"`

	// READ-ONLY; Quota request ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the quota request.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource type. "Microsoft.Quota/ServiceLimitRequests"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// RequestProperties - Quota request properties.
type RequestProperties struct {
	// Error details of the quota request.
	Error *ServiceErrorDetail `json:"error,omitempty"`

	// Quota request details.
	Value []*SubRequest `json:"value,omitempty"`

	// READ-ONLY; User-friendly status message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The quota request status.
	ProvisioningState *QuotaRequestState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; The quota request submission time. The date conforms to the following format specified by the ISO 8601 standard:
	// yyyy-MM-ddTHH:mm:ssZ
	RequestSubmitTime *time.Time `json:"requestSubmitTime,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type RequestProperties.
func (r RequestProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "error", r.Error)
	populate(objectMap, "message", r.Message)
	populate(objectMap, "provisioningState", r.ProvisioningState)
	populateTimeRFC3339(objectMap, "requestSubmitTime", r.RequestSubmitTime)
	populate(objectMap, "value", r.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RequestProperties.
func (r *RequestProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "error":
			err = unpopulate(val, &r.Error)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, &r.Message)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &r.ProvisioningState)
			delete(rawMsg, key)
		case "requestSubmitTime":
			err = unpopulateTimeRFC3339(val, &r.RequestSubmitTime)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, &r.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// RequestStatusClientGetOptions contains the optional parameters for the RequestStatusClient.Get method.
type RequestStatusClientGetOptions struct {
	// placeholder for future optional parameters
}

// RequestStatusClientListOptions contains the optional parameters for the RequestStatusClient.List method.
type RequestStatusClientListOptions struct {
	// FIELD SUPPORTED OPERATORS
	// requestSubmitTime ge, le, eq, gt, lt
	// provisioningState eq {QuotaRequestState}
	// resourceName eq {resourceName}
	Filter *string
	// The Skiptoken parameter is used only if a previous operation returned a partial result. If a previous response contains
	// a nextLink element, its value includes a skiptoken parameter that specifies a
	// starting point to use for subsequent calls.
	Skiptoken *string
	// Number of records to return.
	Top *int32
}

// RequestStatusDetails - Quota request status details.
type RequestStatusDetails struct {
	// Resource quota limit properties.
	Limit *LimitObject `json:"limit,omitempty"`

	// Resource name provided by the resource provider. Use this property name when requesting quota.
	Name *ResourceName `json:"name,omitempty"`

	// Additional properties for the specific resource provider.
	Properties map[string]interface{} `json:"properties,omitempty"`

	// Resource type name.
	ResourceType *string `json:"resourceType,omitempty"`

	// The quota limit units, such as Count and Bytes. When requesting quota, use the unit value returned in the GET response
	// in the request body of your PUT operation.
	Unit *string `json:"unit,omitempty"`

	// READ-ONLY; User-friendly message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; Quota request status.
	ProvisioningState *QuotaRequestState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; The time period over which the quota usage values are summarized. For example: *P1D (per one day) *PT1M (per
	// one minute) *PT1S (per one second). This parameter is optional because, for some resources
	// like compute, the period is irrelevant.
	QuotaPeriod *string `json:"quotaPeriod,omitempty" azure:"ro"`
}

// RequestSubmitResponse - Quota request response.
type RequestSubmitResponse struct {
	// Quota request details.
	Properties *RequestProperties `json:"properties,omitempty"`

	// READ-ONLY; Quota request ID.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Quota request name.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource type. "Microsoft.Quota/quotas".
	Type *string `json:"type,omitempty" azure:"ro"`
}

// RequestSubmitResponse202 - The quota request response with the quota request ID.
type RequestSubmitResponse202 struct {
	// Quota request status.
	Properties *RequestStatusDetails `json:"properties,omitempty"`

	// READ-ONLY; The quota request ID. To check the request status, use the id value in a Quota Request Status [https://docs.microsoft.com/en-us/rest/api/reserved-vm-instances/quotarequeststatus/get]
	// GET operation.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Operation ID.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ResourceName - Name of the resource provided by the resource Provider. When requesting quota, use this property name.
type ResourceName struct {
	// Resource name.
	Value *string `json:"value,omitempty"`

	// READ-ONLY; Resource display name.
	LocalizedValue *string `json:"localizedValue,omitempty" azure:"ro"`
}

// ServiceError - API error details.
type ServiceError struct {
	// Error code.
	Code *string `json:"code,omitempty"`

	// Error message.
	Message *string `json:"message,omitempty"`

	// READ-ONLY; List of error details.
	Details []*ServiceErrorDetail `json:"details,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ServiceError.
func (s ServiceError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", s.Code)
	populate(objectMap, "details", s.Details)
	populate(objectMap, "message", s.Message)
	return json.Marshal(objectMap)
}

// ServiceErrorDetail - Error details.
type ServiceErrorDetail struct {
	// READ-ONLY; Error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; Error message.
	Message *string `json:"message,omitempty" azure:"ro"`
}

// SubRequest - Request property.
type SubRequest struct {
	// Resource quota limit properties.
	Limit LimitJSONObjectClassification `json:"limit,omitempty"`

	// Resource name.
	Name *ResourceName `json:"name,omitempty"`

	// Quota limit units, such as Count and Bytes. When requesting quota, use the unit value returned in the GET response in the
	// request body of your PUT operation.
	Unit *string `json:"unit,omitempty"`

	// READ-ONLY; User-friendly status message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The quota request status.
	ProvisioningState *QuotaRequestState `json:"provisioningState,omitempty" azure:"ro"`

	// READ-ONLY; Resource type for which the quota properties were requested.
	ResourceType *string `json:"resourceType,omitempty" azure:"ro"`

	// READ-ONLY; Quota request ID.
	SubRequestID *string `json:"subRequestId,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type SubRequest.
func (s SubRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "limit", s.Limit)
	populate(objectMap, "message", s.Message)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "resourceType", s.ResourceType)
	populate(objectMap, "subRequestId", s.SubRequestID)
	populate(objectMap, "unit", s.Unit)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SubRequest.
func (s *SubRequest) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "limit":
			s.Limit, err = unmarshalLimitJSONObjectClassification(val)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, &s.Message)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, &s.Name)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, &s.ProvisioningState)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, &s.ResourceType)
			delete(rawMsg, key)
		case "subRequestId":
			err = unpopulate(val, &s.SubRequestID)
			delete(rawMsg, key)
		case "unit":
			err = unpopulate(val, &s.Unit)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// UsagesClientGetOptions contains the optional parameters for the UsagesClient.Get method.
type UsagesClientGetOptions struct {
	// placeholder for future optional parameters
}

// UsagesClientListOptions contains the optional parameters for the UsagesClient.List method.
type UsagesClientListOptions struct {
	// placeholder for future optional parameters
}

// UsagesLimits - Quota limits.
type UsagesLimits struct {
	// The URI used to fetch the next page of quota limits. When there are no more pages, this is null.
	NextLink *string `json:"nextLink,omitempty"`

	// List of quota limits.
	Value []*CurrentUsagesBase `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type UsagesLimits.
func (u UsagesLimits) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", u.NextLink)
	populate(objectMap, "value", u.Value)
	return json.Marshal(objectMap)
}

// UsagesObject - The resource usages value.
type UsagesObject struct {
	// REQUIRED; The usages value.
	Value *int32 `json:"value,omitempty"`

	// The quota or usages limit types.
	UsagesType *UsagesTypes `json:"usagesType,omitempty"`
}

// UsagesProperties - Usage properties for the specified resource.
type UsagesProperties struct {
	// Resource name provided by the resource provider. Use this property name when requesting quota.
	Name *ResourceName `json:"name,omitempty"`

	// Additional properties for the specific resource provider.
	Properties map[string]interface{} `json:"properties,omitempty"`

	// The name of the resource type.
	ResourceType *string `json:"resourceType,omitempty"`

	// The quota limit properties for this resource.
	Usages *UsagesObject `json:"usages,omitempty"`

	// READ-ONLY; States if quota can be requested for this resource.
	IsQuotaApplicable *bool `json:"isQuotaApplicable,omitempty" azure:"ro"`

	// READ-ONLY; The time period for the summary of the quota usage values. For example: *P1D (per one day) *PT1M (per one minute)
	// *PT1S (per one second). This parameter is optional because it is not relevant for all
	// resources such as compute.
	QuotaPeriod *string `json:"quotaPeriod,omitempty" azure:"ro"`

	// READ-ONLY; The units for the quota usage, such as Count and Bytes. When requesting quota, use the unit value returned in
	// the GET response in the request body of your PUT operation.
	Unit *string `json:"unit,omitempty" azure:"ro"`
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
