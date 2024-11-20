package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/cloudbuild"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/cloudrun"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		// Replace 'my-service' with the name of your Cloud Run service.
		serviceName := "spcv-sharp-griffin"
		// Replace 'your-gcp-project' with your actual GCP project ID.
		projectID := "spc-cv-sharp-griffin"
		// Replace 'your-cloudrun-service-location' with the location of your Cloud Run service (e.g., 'us-central1').
		locationID := "us-west1"
		// The name of the custom domain to map to the Cloud Run service.
		// customDomain := "spctechnology.co.uk"

		// Create a Google Cloud Build Trigger.
		buildTrigger, err := cloudbuild.NewTrigger(ctx, "spcvBuildTrigger", &cloudbuild.TriggerArgs{
			Name: pulumi.String(serviceName),
			Github: cloudbuild.TriggerGithubArgs{
				Owner: pulumi.String("SPC-CLOUD"), // Replace with your GitHub username.
				Name:  pulumi.String("spcv"),      // Replace with your GitHub repository name.
				Push: cloudbuild.TriggerGithubPushArgs{
					Branch: pulumi.String("main"), // Replace with your branch name if different.
				},
			},
			// Replace the filename if your build configuration file has a different name or path.
			Filename: pulumi.String("cloudbuild.yaml"),
		})
		if err != nil {
			return err
		}
		// Create the Cloud Run service.
		_, err = cloudrun.NewService(ctx, "spcv", &cloudrun.ServiceArgs{
			Name: pulumi.String(projectID),
			Location: pulumi.String(locationID),
			Template: &cloudrun.ServiceTemplateArgs{
				Spec: &cloudrun.ServiceTemplateSpecArgs{
					// Assuming the Cloud Build process tags the image with 'latest'.
					Containers: cloudrun.ServiceTemplateSpecContainerArray{
						&cloudrun.ServiceTemplateSpecContainerArgs{
							Image: pulumi.Sprintf(locationID+"-docker.pkg.dev/"+projectID+"/cloud-run-source-deploy/spcv/spcv:latest"),
						},
					},
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{buildTrigger}))
		if err != nil {
			return err
		}
		
		noauth, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
			Bindings: []organizations.GetIAMPolicyBinding{
				{
					Role: "roles/run.invoker",
					Members: []string{
						"allUsers",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = cloudrun.NewIamPolicy(ctx, "noauth", &cloudrun.IamPolicyArgs{
			Location:   _default.Location,
			Project:    _default.Project,
			Service:    _default.Name,
			PolicyData: pulumi.String(noauth.PolicyData),
		})
		
		return nil
	})
}