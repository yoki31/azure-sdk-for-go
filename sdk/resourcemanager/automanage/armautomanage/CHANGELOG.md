# Release History

## 0.2.0 (2022-01-13)
### Breaking Changes

- Function `*ConfigurationProfileAssignmentsClient.ListBySubscription` parameter(s) have been changed from `(context.Context, *ConfigurationProfileAssignmentsListBySubscriptionOptions)` to `(context.Context, *ConfigurationProfileAssignmentsClientListBySubscriptionOptions)`
- Function `*ConfigurationProfileAssignmentsClient.ListBySubscription` return value(s) have been changed from `(ConfigurationProfileAssignmentsListBySubscriptionResponse, error)` to `(ConfigurationProfileAssignmentsClientListBySubscriptionResponse, error)`
- Function `*ConfigurationProfileAssignmentsClient.Get` parameter(s) have been changed from `(context.Context, string, string, string, *ConfigurationProfileAssignmentsGetOptions)` to `(context.Context, string, string, string, *ConfigurationProfileAssignmentsClientGetOptions)`
- Function `*ConfigurationProfileAssignmentsClient.Get` return value(s) have been changed from `(ConfigurationProfileAssignmentsGetResponse, error)` to `(ConfigurationProfileAssignmentsClientGetResponse, error)`
- Function `*OperationsClient.List` parameter(s) have been changed from `(context.Context, *OperationsListOptions)` to `(context.Context, *OperationsClientListOptions)`
- Function `*OperationsClient.List` return value(s) have been changed from `(OperationsListResponse, error)` to `(OperationsClientListResponse, error)`
- Function `*ConfigurationProfileAssignmentsClient.List` parameter(s) have been changed from `(context.Context, string, *ConfigurationProfileAssignmentsListOptions)` to `(context.Context, string, *ConfigurationProfileAssignmentsClientListOptions)`
- Function `*ConfigurationProfileAssignmentsClient.List` return value(s) have been changed from `(ConfigurationProfileAssignmentsListResponse, error)` to `(ConfigurationProfileAssignmentsClientListResponse, error)`
- Function `*ConfigurationProfileAssignmentsClient.Delete` parameter(s) have been changed from `(context.Context, string, string, string, *ConfigurationProfileAssignmentsDeleteOptions)` to `(context.Context, string, string, string, *ConfigurationProfileAssignmentsClientDeleteOptions)`
- Function `*ConfigurationProfileAssignmentsClient.Delete` return value(s) have been changed from `(ConfigurationProfileAssignmentsDeleteResponse, error)` to `(ConfigurationProfileAssignmentsClientDeleteResponse, error)`
- Type of `Operation.IsDataAction` has been changed from `*string` to `*bool`
- Type of `ConfigurationProfileAssignmentProperties.ConfigurationProfile` has been changed from `*ConfigurationProfile` to `*string`
- Const `ScanTypeFull` has been removed
- Const `UpdateStatusCreated` has been removed
- Const `EnableRealTimeProtectionTrue` has been removed
- Const `RunScheduledScanTrue` has been removed
- Const `ScanTypeQuick` has been removed
- Const `ConfigurationProfileAzureVirtualMachineBestPracticesProduction` has been removed
- Const `ResourceIdentityTypeSystemAssigned` has been removed
- Const `RunScheduledScanFalse` has been removed
- Const `ProvisioningStateSucceeded` has been removed
- Const `ProvisioningStateCreated` has been removed
- Const `ProvisioningStateFailed` has been removed
- Const `UpdateStatusFailed` has been removed
- Const `UpdateStatusSucceeded` has been removed
- Const `EnableRealTimeProtectionFalse` has been removed
- Const `ResourceIdentityTypeNone` has been removed
- Const `ConfigurationProfileAzureVirtualMachineBestPracticesDevTest` has been removed
- Function `EnableRealTimeProtection.ToPtr` has been removed
- Function `*ConfigurationProfilePreferencesClient.CreateOrUpdate` has been removed
- Function `PossibleEnableRealTimeProtectionValues` has been removed
- Function `ErrorResponse.Error` has been removed
- Function `PossibleScanTypeValues` has been removed
- Function `PossibleResourceIdentityTypeValues` has been removed
- Function `ResourceIdentityType.ToPtr` has been removed
- Function `RunScheduledScan.ToPtr` has been removed
- Function `ConfigurationProfilePreference.MarshalJSON` has been removed
- Function `ConfigurationProfile.ToPtr` has been removed
- Function `PossibleRunScheduledScanValues` has been removed
- Function `PossibleConfigurationProfileValues` has been removed
- Function `*ConfigurationProfileAssignmentsCreateOrUpdatePollerResponse.Resume` has been removed
- Function `UpdateStatus.ToPtr` has been removed
- Function `Account.MarshalJSON` has been removed
- Function `ConfigurationProfilePreferenceUpdate.MarshalJSON` has been removed
- Function `*AccountsClient.ListByResourceGroup` has been removed
- Function `*ConfigurationProfileAssignmentsCreateOrUpdatePoller.FinalResponse` has been removed
- Function `*ConfigurationProfilePreferencesClient.ListBySubscription` has been removed
- Function `*ConfigurationProfilePreferencesClient.Get` has been removed
- Function `Resource.MarshalJSON` has been removed
- Function `OperationList.MarshalJSON` has been removed
- Function `*ConfigurationProfileAssignmentsCreateOrUpdatePoller.Done` has been removed
- Function `*AccountsClient.ListBySubscription` has been removed
- Function `ConfigurationProfileAssignment.MarshalJSON` has been removed
- Function `*ConfigurationProfilePreferencesClient.Update` has been removed
- Function `*AccountsClient.Get` has been removed
- Function `PossibleProvisioningStateValues` has been removed
- Function `ConfigurationProfileAssignmentsCreateOrUpdatePollerResponse.PollUntilDone` has been removed
- Function `NewAccountsClient` has been removed
- Function `*AccountsClient.CreateOrUpdate` has been removed
- Function `*ConfigurationProfileAssignmentsCreateOrUpdatePoller.Poll` has been removed
- Function `PossibleUpdateStatusValues` has been removed
- Function `*ConfigurationProfileAssignmentsCreateOrUpdatePoller.ResumeToken` has been removed
- Function `AccountUpdate.MarshalJSON` has been removed
- Function `AccountList.MarshalJSON` has been removed
- Function `*AccountsClient.Update` has been removed
- Function `ProvisioningState.ToPtr` has been removed
- Function `*ConfigurationProfilePreferencesClient.Delete` has been removed
- Function `*AccountsClient.Delete` has been removed
- Function `NewConfigurationProfilePreferencesClient` has been removed
- Function `*ConfigurationProfilePreferencesClient.ListByResourceGroup` has been removed
- Function `ConfigurationProfilePreferenceList.MarshalJSON` has been removed
- Function `*ConfigurationProfileAssignmentsClient.BeginCreateOrUpdate` has been removed
- Function `ScanType.ToPtr` has been removed
- Struct `Account` has been removed
- Struct `AccountIdentity` has been removed
- Struct `AccountList` has been removed
- Struct `AccountUpdate` has been removed
- Struct `AccountsClient` has been removed
- Struct `AccountsCreateOrUpdateOptions` has been removed
- Struct `AccountsCreateOrUpdateResponse` has been removed
- Struct `AccountsCreateOrUpdateResult` has been removed
- Struct `AccountsDeleteOptions` has been removed
- Struct `AccountsDeleteResponse` has been removed
- Struct `AccountsGetOptions` has been removed
- Struct `AccountsGetResponse` has been removed
- Struct `AccountsGetResult` has been removed
- Struct `AccountsListByResourceGroupOptions` has been removed
- Struct `AccountsListByResourceGroupResponse` has been removed
- Struct `AccountsListByResourceGroupResult` has been removed
- Struct `AccountsListBySubscriptionOptions` has been removed
- Struct `AccountsListBySubscriptionResponse` has been removed
- Struct `AccountsListBySubscriptionResult` has been removed
- Struct `AccountsUpdateOptions` has been removed
- Struct `AccountsUpdateResponse` has been removed
- Struct `AccountsUpdateResult` has been removed
- Struct `ConfigurationProfileAssignmentCompliance` has been removed
- Struct `ConfigurationProfileAssignmentsBeginCreateOrUpdateOptions` has been removed
- Struct `ConfigurationProfileAssignmentsCreateOrUpdatePoller` has been removed
- Struct `ConfigurationProfileAssignmentsCreateOrUpdatePollerResponse` has been removed
- Struct `ConfigurationProfileAssignmentsCreateOrUpdateResponse` has been removed
- Struct `ConfigurationProfileAssignmentsCreateOrUpdateResult` has been removed
- Struct `ConfigurationProfileAssignmentsDeleteOptions` has been removed
- Struct `ConfigurationProfileAssignmentsDeleteResponse` has been removed
- Struct `ConfigurationProfileAssignmentsGetOptions` has been removed
- Struct `ConfigurationProfileAssignmentsGetResponse` has been removed
- Struct `ConfigurationProfileAssignmentsGetResult` has been removed
- Struct `ConfigurationProfileAssignmentsListBySubscriptionOptions` has been removed
- Struct `ConfigurationProfileAssignmentsListBySubscriptionResponse` has been removed
- Struct `ConfigurationProfileAssignmentsListBySubscriptionResult` has been removed
- Struct `ConfigurationProfileAssignmentsListOptions` has been removed
- Struct `ConfigurationProfileAssignmentsListResponse` has been removed
- Struct `ConfigurationProfileAssignmentsListResult` has been removed
- Struct `ConfigurationProfilePreference` has been removed
- Struct `ConfigurationProfilePreferenceAntiMalware` has been removed
- Struct `ConfigurationProfilePreferenceList` has been removed
- Struct `ConfigurationProfilePreferenceProperties` has been removed
- Struct `ConfigurationProfilePreferenceUpdate` has been removed
- Struct `ConfigurationProfilePreferenceVMBackup` has been removed
- Struct `ConfigurationProfilePreferencesClient` has been removed
- Struct `ConfigurationProfilePreferencesCreateOrUpdateOptions` has been removed
- Struct `ConfigurationProfilePreferencesCreateOrUpdateResponse` has been removed
- Struct `ConfigurationProfilePreferencesCreateOrUpdateResult` has been removed
- Struct `ConfigurationProfilePreferencesDeleteOptions` has been removed
- Struct `ConfigurationProfilePreferencesDeleteResponse` has been removed
- Struct `ConfigurationProfilePreferencesGetOptions` has been removed
- Struct `ConfigurationProfilePreferencesGetResponse` has been removed
- Struct `ConfigurationProfilePreferencesGetResult` has been removed
- Struct `ConfigurationProfilePreferencesListByResourceGroupOptions` has been removed
- Struct `ConfigurationProfilePreferencesListByResourceGroupResponse` has been removed
- Struct `ConfigurationProfilePreferencesListByResourceGroupResult` has been removed
- Struct `ConfigurationProfilePreferencesListBySubscriptionOptions` has been removed
- Struct `ConfigurationProfilePreferencesListBySubscriptionResponse` has been removed
- Struct `ConfigurationProfilePreferencesListBySubscriptionResult` has been removed
- Struct `ConfigurationProfilePreferencesUpdateOptions` has been removed
- Struct `ConfigurationProfilePreferencesUpdateResponse` has been removed
- Struct `ConfigurationProfilePreferencesUpdateResult` has been removed
- Struct `OperationList` has been removed
- Struct `OperationProperties` has been removed
- Struct `OperationsListOptions` has been removed
- Struct `OperationsListResponse` has been removed
- Struct `OperationsListResult` has been removed
- Field `Resource` of struct `ProxyResource` has been removed
- Field `Resource` of struct `TrackedResource` has been removed
- Field `Properties` of struct `Operation` has been removed
- Field `ProxyResource` of struct `ConfigurationProfileAssignment` has been removed
- Field `InnerError` of struct `ErrorResponse` has been removed
- Field `AccountID` of struct `ConfigurationProfileAssignmentProperties` has been removed
- Field `Compliance` of struct `ConfigurationProfileAssignmentProperties` has been removed
- Field `ConfigurationProfilePreferenceID` of struct `ConfigurationProfileAssignmentProperties` has been removed
- Field `ProvisioningState` of struct `ConfigurationProfileAssignmentProperties` has been removed

### Features Added

- New const `ActionTypeInternal`
- New const `CreatedByTypeApplication`
- New const `CreatedByTypeKey`
- New const `CreatedByTypeManagedIdentity`
- New const `CreatedByTypeUser`
- New const `OriginSystem`
- New const `OriginUser`
- New const `OriginUserSystem`
- New function `Origin.ToPtr() *Origin`
- New function `BestPracticeList.MarshalJSON() ([]byte, error)`
- New function `PossibleOriginValues() []Origin`
- New function `*ConfigurationProfilesVersionsClient.ListChildResources(context.Context, string, string, *ConfigurationProfilesVersionsClientListChildResourcesOptions) (ConfigurationProfilesVersionsClientListChildResourcesResponse, error)`
- New function `timeRFC3339.MarshalText() ([]byte, error)`
- New function `*timeRFC3339.UnmarshalText([]byte) error`
- New function `*BestPracticesVersionsClient.Get(context.Context, string, string, *BestPracticesVersionsClientGetOptions) (BestPracticesVersionsClientGetResponse, error)`
- New function `*ConfigurationProfilesClient.ListByResourceGroup(context.Context, string, *ConfigurationProfilesClientListByResourceGroupOptions) (ConfigurationProfilesClientListByResourceGroupResponse, error)`
- New function `*ConfigurationProfileAssignmentsClient.CreateOrUpdate(context.Context, string, string, string, ConfigurationProfileAssignment, *ConfigurationProfileAssignmentsClientCreateOrUpdateOptions) (ConfigurationProfileAssignmentsClientCreateOrUpdateResponse, error)`
- New function `CreatedByType.ToPtr() *CreatedByType`
- New function `*SystemData.UnmarshalJSON([]byte) error`
- New function `*ConfigurationProfilesClient.ListBySubscription(context.Context, *ConfigurationProfilesClientListBySubscriptionOptions) (ConfigurationProfilesClientListBySubscriptionResponse, error)`
- New function `*ConfigurationProfilesClient.CreateOrUpdate(context.Context, string, string, ConfigurationProfile, *ConfigurationProfilesClientCreateOrUpdateOptions) (ConfigurationProfilesClientCreateOrUpdateResponse, error)`
- New function `*ConfigurationProfilesClient.Delete(context.Context, string, string, *ConfigurationProfilesClientDeleteOptions) (ConfigurationProfilesClientDeleteResponse, error)`
- New function `ReportList.MarshalJSON() ([]byte, error)`
- New function `*ConfigurationProfilesVersionsClient.Get(context.Context, string, string, string, *ConfigurationProfilesVersionsClientGetOptions) (ConfigurationProfilesVersionsClientGetResponse, error)`
- New function `*ConfigurationProfilesClient.Update(context.Context, string, string, ConfigurationProfileUpdate, *ConfigurationProfilesClientUpdateOptions) (ConfigurationProfilesClientUpdateResponse, error)`
- New function `*ConfigurationProfilesClient.Get(context.Context, string, string, *ConfigurationProfilesClientGetOptions) (ConfigurationProfilesClientGetResponse, error)`
- New function `NewConfigurationProfilesClient(string, azcore.TokenCredential, *arm.ClientOptions) *ConfigurationProfilesClient`
- New function `NewBestPracticesClient(azcore.TokenCredential, *arm.ClientOptions) *BestPracticesClient`
- New function `ConfigurationProfileList.MarshalJSON() ([]byte, error)`
- New function `*ReportsClient.ListByConfigurationProfileAssignments(context.Context, string, string, string, *ReportsClientListByConfigurationProfileAssignmentsOptions) (ReportsClientListByConfigurationProfileAssignmentsResponse, error)`
- New function `*BestPracticesVersionsClient.ListByTenant(context.Context, string, *BestPracticesVersionsClientListByTenantOptions) (BestPracticesVersionsClientListByTenantResponse, error)`
- New function `*ReportsClient.Get(context.Context, string, string, string, string, *ReportsClientGetOptions) (ReportsClientGetResponse, error)`
- New function `*BestPracticesClient.ListByTenant(context.Context, *BestPracticesClientListByTenantOptions) (BestPracticesClientListByTenantResponse, error)`
- New function `*timeRFC3339.Parse(string) error`
- New function `PossibleActionTypeValues() []ActionType`
- New function `OperationListResult.MarshalJSON() ([]byte, error)`
- New function `ConfigurationProfile.MarshalJSON() ([]byte, error)`
- New function `NewBestPracticesVersionsClient(azcore.TokenCredential, *arm.ClientOptions) *BestPracticesVersionsClient`
- New function `*ConfigurationProfilesVersionsClient.Update(context.Context, string, string, string, ConfigurationProfileUpdate, *ConfigurationProfilesVersionsClientUpdateOptions) (ConfigurationProfilesVersionsClientUpdateResponse, error)`
- New function `NewReportsClient(string, azcore.TokenCredential, *arm.ClientOptions) *ReportsClient`
- New function `*ConfigurationProfilesVersionsClient.CreateOrUpdate(context.Context, string, string, string, ConfigurationProfile, *ConfigurationProfilesVersionsClientCreateOrUpdateOptions) (ConfigurationProfilesVersionsClientCreateOrUpdateResponse, error)`
- New function `ActionType.ToPtr() *ActionType`
- New function `*BestPracticesClient.Get(context.Context, string, *BestPracticesClientGetOptions) (BestPracticesClientGetResponse, error)`
- New function `NewConfigurationProfilesVersionsClient(string, azcore.TokenCredential, *arm.ClientOptions) *ConfigurationProfilesVersionsClient`
- New function `timeRFC3339.MarshalJSON() ([]byte, error)`
- New function `*ConfigurationProfilesVersionsClient.Delete(context.Context, string, string, string, *ConfigurationProfilesVersionsClientDeleteOptions) (ConfigurationProfilesVersionsClientDeleteResponse, error)`
- New function `PossibleCreatedByTypeValues() []CreatedByType`
- New function `ConfigurationProfileProperties.MarshalJSON() ([]byte, error)`
- New function `ConfigurationProfileUpdate.MarshalJSON() ([]byte, error)`
- New function `SystemData.MarshalJSON() ([]byte, error)`
- New function `ConfigurationProfileAssignmentProperties.MarshalJSON() ([]byte, error)`
- New function `AssignmentReportProperties.MarshalJSON() ([]byte, error)`
- New function `*timeRFC3339.UnmarshalJSON([]byte) error`
- New struct `AssignmentReportProperties`
- New struct `BestPractice`
- New struct `BestPracticeList`
- New struct `BestPracticesClient`
- New struct `BestPracticesClientGetOptions`
- New struct `BestPracticesClientGetResponse`
- New struct `BestPracticesClientGetResult`
- New struct `BestPracticesClientListByTenantOptions`
- New struct `BestPracticesClientListByTenantResponse`
- New struct `BestPracticesClientListByTenantResult`
- New struct `BestPracticesVersionsClient`
- New struct `BestPracticesVersionsClientGetOptions`
- New struct `BestPracticesVersionsClientGetResponse`
- New struct `BestPracticesVersionsClientGetResult`
- New struct `BestPracticesVersionsClientListByTenantOptions`
- New struct `BestPracticesVersionsClientListByTenantResponse`
- New struct `BestPracticesVersionsClientListByTenantResult`
- New struct `ConfigurationProfile`
- New struct `ConfigurationProfileAssignmentsClientCreateOrUpdateOptions`
- New struct `ConfigurationProfileAssignmentsClientCreateOrUpdateResponse`
- New struct `ConfigurationProfileAssignmentsClientCreateOrUpdateResult`
- New struct `ConfigurationProfileAssignmentsClientDeleteOptions`
- New struct `ConfigurationProfileAssignmentsClientDeleteResponse`
- New struct `ConfigurationProfileAssignmentsClientGetOptions`
- New struct `ConfigurationProfileAssignmentsClientGetResponse`
- New struct `ConfigurationProfileAssignmentsClientGetResult`
- New struct `ConfigurationProfileAssignmentsClientListBySubscriptionOptions`
- New struct `ConfigurationProfileAssignmentsClientListBySubscriptionResponse`
- New struct `ConfigurationProfileAssignmentsClientListBySubscriptionResult`
- New struct `ConfigurationProfileAssignmentsClientListOptions`
- New struct `ConfigurationProfileAssignmentsClientListResponse`
- New struct `ConfigurationProfileAssignmentsClientListResult`
- New struct `ConfigurationProfileList`
- New struct `ConfigurationProfileProperties`
- New struct `ConfigurationProfileUpdate`
- New struct `ConfigurationProfilesClient`
- New struct `ConfigurationProfilesClientCreateOrUpdateOptions`
- New struct `ConfigurationProfilesClientCreateOrUpdateResponse`
- New struct `ConfigurationProfilesClientCreateOrUpdateResult`
- New struct `ConfigurationProfilesClientDeleteOptions`
- New struct `ConfigurationProfilesClientDeleteResponse`
- New struct `ConfigurationProfilesClientGetOptions`
- New struct `ConfigurationProfilesClientGetResponse`
- New struct `ConfigurationProfilesClientGetResult`
- New struct `ConfigurationProfilesClientListByResourceGroupOptions`
- New struct `ConfigurationProfilesClientListByResourceGroupResponse`
- New struct `ConfigurationProfilesClientListByResourceGroupResult`
- New struct `ConfigurationProfilesClientListBySubscriptionOptions`
- New struct `ConfigurationProfilesClientListBySubscriptionResponse`
- New struct `ConfigurationProfilesClientListBySubscriptionResult`
- New struct `ConfigurationProfilesClientUpdateOptions`
- New struct `ConfigurationProfilesClientUpdateResponse`
- New struct `ConfigurationProfilesClientUpdateResult`
- New struct `ConfigurationProfilesVersionsClient`
- New struct `ConfigurationProfilesVersionsClientCreateOrUpdateOptions`
- New struct `ConfigurationProfilesVersionsClientCreateOrUpdateResponse`
- New struct `ConfigurationProfilesVersionsClientCreateOrUpdateResult`
- New struct `ConfigurationProfilesVersionsClientDeleteOptions`
- New struct `ConfigurationProfilesVersionsClientDeleteResponse`
- New struct `ConfigurationProfilesVersionsClientGetOptions`
- New struct `ConfigurationProfilesVersionsClientGetResponse`
- New struct `ConfigurationProfilesVersionsClientGetResult`
- New struct `ConfigurationProfilesVersionsClientListChildResourcesOptions`
- New struct `ConfigurationProfilesVersionsClientListChildResourcesResponse`
- New struct `ConfigurationProfilesVersionsClientListChildResourcesResult`
- New struct `ConfigurationProfilesVersionsClientUpdateOptions`
- New struct `ConfigurationProfilesVersionsClientUpdateResponse`
- New struct `ConfigurationProfilesVersionsClientUpdateResult`
- New struct `OperationListResult`
- New struct `OperationsClientListOptions`
- New struct `OperationsClientListResponse`
- New struct `OperationsClientListResult`
- New struct `Report`
- New struct `ReportList`
- New struct `ReportResource`
- New struct `ReportsClient`
- New struct `ReportsClientGetOptions`
- New struct `ReportsClientGetResponse`
- New struct `ReportsClientGetResult`
- New struct `ReportsClientListByConfigurationProfileAssignmentsOptions`
- New struct `ReportsClientListByConfigurationProfileAssignmentsResponse`
- New struct `ReportsClientListByConfigurationProfileAssignmentsResult`
- New struct `SystemData`
- New field `ID` in struct `ConfigurationProfileAssignment`
- New field `Name` in struct `ConfigurationProfileAssignment`
- New field `SystemData` in struct `ConfigurationProfileAssignment`
- New field `Type` in struct `ConfigurationProfileAssignment`
- New field `ID` in struct `ProxyResource`
- New field `Name` in struct `ProxyResource`
- New field `Type` in struct `ProxyResource`
- New field `Error` in struct `ErrorResponse`
- New field `ID` in struct `TrackedResource`
- New field `Name` in struct `TrackedResource`
- New field `Type` in struct `TrackedResource`
- New field `ActionType` in struct `Operation`
- New field `Origin` in struct `Operation`
- New field `Status` in struct `ConfigurationProfileAssignmentProperties`
- New field `ProfileOverrides` in struct `ConfigurationProfileAssignmentProperties`


## 0.1.0 (2021-11-30)

- Initial preview release.
