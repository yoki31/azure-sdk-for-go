//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-10-01/examples/ReplicationProtectionIntents_List.json
func ExampleReplicationProtectionIntentsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationProtectionIntentsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	pager := client.List(&armrecoveryservicessiterecovery.ReplicationProtectionIntentsListOptions{SkipToken: nil,
		TakeToken: nil,
	})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("ReplicationProtectionIntent.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-10-01/examples/ReplicationProtectionIntents_Get.json
func ExampleReplicationProtectionIntentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationProtectionIntentsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<intent-object-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ReplicationProtectionIntent.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2021-10-01/examples/ReplicationProtectionIntents_Create.json
func ExampleReplicationProtectionIntentsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservicessiterecovery.NewReplicationProtectionIntentsClient("<resource-name>",
		"<resource-group-name>",
		"<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<intent-object-name>",
		armrecoveryservicessiterecovery.CreateProtectionIntentInput{
			Properties: &armrecoveryservicessiterecovery.CreateProtectionIntentProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.A2ACreateProtectionIntentInput{
					CreateProtectionIntentProviderSpecificDetails: armrecoveryservicessiterecovery.CreateProtectionIntentProviderSpecificDetails{
						InstanceType: to.StringPtr("<instance-type>"),
					},
					FabricObjectID:           to.StringPtr("<fabric-object-id>"),
					PrimaryLocation:          to.StringPtr("<primary-location>"),
					RecoveryAvailabilityType: armrecoveryservicessiterecovery.A2ARecoveryAvailabilityTypeSingle.ToPtr(),
					RecoveryLocation:         to.StringPtr("<recovery-location>"),
					RecoveryResourceGroupID:  to.StringPtr("<recovery-resource-group-id>"),
					RecoverySubscriptionID:   to.StringPtr("<recovery-subscription-id>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ReplicationProtectionIntent.ID: %s\n", *res.ID)
}