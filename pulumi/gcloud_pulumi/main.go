package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/cloudbuild"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/cloudrun"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a Google Cloud Build Trigger.
		buildTrigger, err := cloudbuild.NewTrigger(ctx, "spcvBuildTrigger", &cloudbuild.TriggerArgs{
			Name: pulumi.String("spcv-sharp-griffin"),
			Github: cloudbuild.TriggerGithubArgs{
				Owner: pulumi.String("spatel96"), // Replace with your GitHub username.
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
			Name: pulumi.String("spcv-sharp-griffin"),
			Location: pulumi.String("us-west1"), // Replace with your preferred Google Cloud region.
			Template: &cloudrun.ServiceTemplateArgs{
				Spec: &cloudrun.ServiceTemplateSpecArgs{
					// Assuming the Cloud Build process tags the image with 'latest'.
					Containers: cloudrun.ServiceTemplateSpecContainerArray{
						&cloudrun.ServiceTemplateSpecContainerArgs{
							Image: pulumi.Sprintf("us-west1-docker.pkg.dev/spc-cv-sharp-griffin/cloud-run-source-deploy/spcv/spcv:latest"),
						},
					},
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{buildTrigger}))
		if err != nil {
			return err
		}
		
		return nil
	})
}