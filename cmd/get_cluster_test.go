package cmd

import (
	"bytes"
	"context"
	"net/http"
	"testing"

	"github.com/neticdk-k8s/ic/internal/apiclient"
	"github.com/neticdk-k8s/ic/internal/logger"
	"github.com/neticdk-k8s/ic/internal/oidc"
	"github.com/neticdk-k8s/ic/internal/usecases/authentication"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_GetClusterCommand(t *testing.T) {
	got := new(bytes.Buffer)
	in := ExecutionContextInput{
		Stdout: got,
		Stderr: got,
	}
	ec := NewExecutionContext(in)
	mockAuthenticator := authentication.NewMockAuthenticator(t)
	mockAuthenticator.EXPECT().
		SetLogger(mock.Anything).
		Run(func(_ logger.Logger) {}).
		Return()
	mockAuthenticator.EXPECT().
		Login(mock.Anything, mock.Anything).
		Run(func(_ context.Context, in authentication.LoginInput) {}).
		Return(&oidc.TokenSet{
			IDToken:      "YOUR_ID_TOKEN",
			RefreshToken: "YOUR_REFRESH_TOKEN",
		}, nil)
	ec.Authenticator = mockAuthenticator
	included := []map[string]interface{}{
		{
			"@id":   "my-provider-id",
			"@type": "Provider",
			"name":  "my-provider",
		},
		{
			"@id":   "my-rz-id",
			"@type": "ResilienceZone",
			"name":  "my-resilience-zone",
		},
		{
			"@id":             "my-cluster-id",
			"@type":           "Cluster",
			"name":            "my-cluster",
			"clusterType":     "dedicated",
			"environmentName": "testing",
			"provider":        "my-provider-id",
			"resilienceZone":  "my-rz-id",
			"kubernetesVersion": map[string]interface{}{
				"version": "v1.2.3",
			},
		},
	}
	name := "my-cluster"
	providerId := "my-provider-id"
	mockClientWithResponsesInterface := apiclient.NewMockClientWithResponsesInterface(t)
	mockClientWithResponsesInterface.EXPECT().
		GetClusterWithResponse(mock.Anything, mock.Anything).
		Return(
			&apiclient.GetClusterResponse{
				Body: make([]byte, 0),
				HTTPResponse: &http.Response{
					Status:     "200 OK",
					StatusCode: 200,
				},
				ApplicationldJSONDefault: &apiclient.Cluster{
					Name:     &name,
					Provider: &providerId,
					Included: &included,
				},
			}, nil)
	apiClient := mockClientWithResponsesInterface
	ec.APIClient = apiClient

	cmd := NewRootCmd(ec)

	t.Run("get cluster my-cluster.my-provider", func(t *testing.T) {
		cmd.SetArgs([]string{"get", "cluster", "my-cluster.my-provider"})
		err := cmd.ExecuteContext(context.Background())
		assert.NoError(t, err)
		assert.Contains(t, got.String(), "Logging in")
		assert.Contains(t, got.String(), "Getting cluster")
		assert.Contains(t, got.String(), "my-cluster")
		assert.Contains(t, got.String(), "my-provider")
	})

	t.Run("get cluster my-cluster -o json", func(t *testing.T) {
		cmd.SetArgs([]string{"get", "cluster", "my-cluster", "-o", "json"})
		err := cmd.ExecuteContext(context.Background())
		assert.NoError(t, err)
		assert.Contains(t, got.String(), "\"name\": \"my-cluster\"")
		assert.Contains(t, got.String(), "\"provider_name\": \"my-provider\"")
	})
}
