package main

import (
    "context"
    "fmt"
    "log"

    "github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
    "github.com/Azure/azure-sdk-for-go/sdk/azidentity"
    "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
    "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
)

func main() {
    // Autentisering
    cred, err := azidentity.NewDefaultAzureCredential(nil)
    if err != nil {
        log.Fatalf("Kunne ikke autentisere: %v", err)
    }
    ctx := context.Background()

    subscriptionID := "061e6f78-04b6-4cac-ace8-cdbd4d5206b5" // ← din Azure subscription ID

    // Opprett ny Resource Group
    rgClient, err := armresources.NewResourceGroupsClient(subscriptionID, cred, nil)
    if err != nil {
        log.Fatalf("Kunne ikke lage ResourceGroupsClient: %v", err)
    }

    _, err = rgClient.CreateOrUpdate(ctx, "aks-demo-rg2", armresources.ResourceGroup{
        Location: to.Ptr("Norway East"),
    }, nil)
    if err != nil {
        log.Fatalf("Kunne ikke opprette Resource Group: %v", err)
    }
    fmt.Println("✅ Resource Group 'aks-demo-rg2' opprettet!")

    // Opprett nytt AKS Cluster i aks-demo-rg2
    aksClient, err := armcontainerservice.NewManagedClustersClient(subscriptionID, cred, nil)
    if err != nil {
        log.Fatalf("Kunne ikke lage ManagedClustersClient: %v", err)
    }

    cluster := armcontainerservice.ManagedCluster{
        Location: to.Ptr("Norway East"),
        Properties: &armcontainerservice.ManagedClusterProperties{
            DNSPrefix: to.Ptr("aksdemocluster2"),
            AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
                {
                    Name:   to.Ptr("systempool"),
                    Count:  to.Ptr(int32(1)),
                    VMSize: to.Ptr("Standard_D2as_v5"),
                    Mode:   to.Ptr(armcontainerservice.AgentPoolModeSystem), // 👈 viktig!
                },
            },
        },
        Identity: &armcontainerservice.ManagedClusterIdentity{
            Type: to.Ptr(armcontainerservice.ResourceIdentityTypeSystemAssigned),
        },
    }

    poller, err := aksClient.BeginCreateOrUpdate(ctx, "aks-demo-rg2", "aks-demo-cluster2", cluster, nil)
    if err != nil {
        log.Fatalf("Kunne ikke opprette AKS cluster: %v", err)
    }

    resp, err := poller.PollUntilDone(ctx, nil)
    if err != nil {
        log.Fatalf("Feil under oppretting av cluster: %v", err)
    }

    fmt.Printf("✅ AKS Cluster opprettet: %s\n", *resp.Name)
}
