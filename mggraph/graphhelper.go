// package mggraphhelper

// import (
// 	"context"
// 	"os"
// 	"log"

// 	"github.com/joho/godotenv"
// 	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
// 	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
// 	auth "github.com/microsoft/kiota-authentication-azure-go"
// 	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
// 	"github.com/microsoftgraph/msgraph-sdk-go/models"
// 	"github.com/microsoftgraph/msgraph-sdk-go/users"
// )

// type GraphHelper struct {
// 	clientSecretCredential *azidentity.ClientSecretCredential
// 	appClient			   *msgraphsdk.GraphServiceClient
// }

// func NewGraphHelpeer() *GraphHelper {
// 	g := &GraphHelper{}
// 	return g
// }

// func (g *GraphHelper) InitializeGraphForAppOnlyAuth() error {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	clientId := os.Getenv("CLIENT_ID")
// 	tenantId := os.Getenv("TENANT_ID")
// 	clientSecret := os.Getenv("CLIENT_SECRET")
// 	credential, err := azidentity.NewClientSecretCredential(tenantId, clientId, clientSecret, nil)
// 	if err != nil {
// 		log.Panicf("Error creating new client secret credential: %v\n", err)
// 		return err
// 	}

// 	g.clientSecretCredential = credential

// 	authProvider, err := auth.NewAzureIdentityAuthenticationProviderWithScopes(g.clientSecretCredential, []string{"https://graph.microsoft.com/.default"})
// 	if err != nil {
// 		log.Panicf("Error creating authentication provider: %v\n", err)
// 		return err
// 	}

// 	adapter, err := msgraphsdk.NewGraphRequestAdapter(authProvider)
// 	if err != nil {
// 		log.Panicf("Error creating graph request adapter: %v\n", err)
// 		return err
// 	}

// 	client := msgraphsdk.NewGraphServiceClient(adapter)
// 	g.appClient = client

// 	return nil
// }

// func (g *GraphHelper) GetUsers() (models.UserCollectionResponseable, error) {
// 	var topValue int32 = 25
// 	query := users.UsersRequestBuilderGetQueryParameters{
// 		// Only request specific properties
// 		Select: []string{"displayName", "id", "mail"},
// 		// Get at most 25 results
// 		Top: &topValue,
// 		// Sort by display name
// 		Orderby: []string{"displayName"},
// 	}

// 	return g.appClient.Users().
// 		Get(context.Background(),
// 			&users.UsersRequestBuilderGetRequestConfiguration{
// 				QueryParameters: &query,
// 			})
// }