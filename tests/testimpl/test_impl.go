package testimpl

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
)

func TestComposableComplete(t *testing.T, ctx types.TestContext) {

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	if len(subscriptionID) == 0 {
		t.Fatalf("AZURE_SUBSCRIPTION_ID is not set in the environment variables")
	}

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		t.Fatalf("Unable to get credentials: %v\n", err)
	}

	options := arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Cloud: cloud.AzurePublic,
		},
	}

	clientFactory, err := armmonitor.NewClientFactory(subscriptionID, credential, &options)
	if err != nil {
		t.Fatalf("Unable to get clientFactory: %v\n", err)
	}

	expectedRgName := terraform.Output(t, ctx.TerratestTerraformOptions(), "resource_group_name")
	expectedMonitorWorkspaceName := terraform.Output(t, ctx.TerratestTerraformOptions(), "monitor_workspace_name")
	expectedMonitorWorkspaceId := terraform.Output(t, ctx.TerratestTerraformOptions(), "monitor_workspace_id")

	workspacesClient := clientFactory.NewAzureMonitorWorkspacesClient()

	res, err := workspacesClient.Get(context.Background(), expectedRgName, expectedMonitorWorkspaceName, nil)
	if err != nil {
		t.Fatalf("Error occurred while getting resource: %v\n", err)

	}

	t.Run("TestWorkspaceClientExists", func(t *testing.T) {
		assert.Equal(t, strings.ToLower(expectedMonitorWorkspaceId), strings.ToLower(*res.ID), "Ids must match")
	})
}
