package main

import (
	"cdk.tf/go/stack/generated/hashicorp/aws"
	"cdk.tf/go/stack/generated/hashicorp/aws/ec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	aws.NewAwsProvider(stack, jsii.String("AWS"), &aws.AwsProviderConfig{
		Region: jsii.String("us-east-1"),
	})

	instance := ec2.NewInstance(stack, jsii.String("compute"), &ec2.InstanceConfig{
		Ami:          jsii.String("ami-0cff7528ff583bf9a"),
		InstanceType: jsii.String("t2.micro"),
		Tags: &map[string]*string{
			"Name": jsii.String("CDKTF-Demo"),
		},
	})

	cdktf.NewTerraformOutput(stack, jsii.String("public_ip"), &cdktf.TerraformOutputConfig{
		Value: instance.PublicIp(),
	})

	return stack
}

func main() {
	app := cdktf.NewApp(nil)
	stack := NewMyStack(app, "aws_instance")

	cdktf.NewS3Backend(stack, &cdktf.S3BackendProps{
		Bucket: jsii.String("fomiller-cdktf-state"),
		Key:    jsii.String("ec2"),
		Region: jsii.String("us-east-1"),
	})

	app.Synth()
}
