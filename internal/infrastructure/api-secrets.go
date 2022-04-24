package infrastructure

import (
	"context"
	"fmt"
	"log"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/pkg/errors"
	"google.golang.org/api/cloudresourcemanager/v1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

func getProjectNumber() (int64, error) {
	crmService, err := cloudresourcemanager.NewService(context.Background())
	if err != nil {
		return 0, errors.Wrap(err, "Failed to get cloud resource manager service")
	}

	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	project, err := crmService.Projects.Get(projectID).Do()
	if err != nil {
		return 0, errors.Wrap(err, "Failed to get project number from ID")
	}

	log.Print("Project number for project: ", project.ProjectId, project.ProjectNumber)
	return project.ProjectNumber, nil
}

func GetSecret(secretId string) (string, error) {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return "", errors.Wrap(err, "Failed to setup secret manager client")
	}

	projectNumber, err := getProjectNumber()
	if err != nil {
		return "", err
	}

	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%d/secrets/%s/versions/%d", projectNumber, secretId, 1),
	}

	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		return "", errors.Wrap(err, "Failed to get secret")
	}

	return string(result.Payload.Data), nil
}
