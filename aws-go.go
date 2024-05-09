package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type AwsGoStackProps struct {
	awscdk.StackProps
}

func NewAwsGoStack(scope constructs.Construct, id string, props *AwsGoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	//create a table
	table := awsdynamodb.NewTable(stack, jsii.String("myUsersTable"), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("username"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		TableName: jsii.String("users"),
	})

	//define a lambda function to invoke our backend
	myLambda := awslambda.NewFunction(stack, jsii.String("myFirstCDKLambda"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_PROVIDED_AL2023(),
		Handler: jsii.String("main"),
		Code:    awslambda.Code_FromAsset(jsii.String("lambda/function.zip"), nil),
	})

	api := awsapigateway.NewRestApi(stack, jsii.String("myFirstCDKApi"), &awsapigateway.RestApiProps{
		RestApiName: jsii.String("myCDKApi"),
		DefaultCorsPreflightOptions: &awsapigateway.CorsOptions{
			AllowHeaders: jsii.Strings("Content-Type", "Authorization"),
			AllowMethods: jsii.Strings("GET", "POST", "DELETE", "PUT", "OPTIONS"),
			AllowOrigins: jsii.Strings("*"),
		},
	})

	integration := awsapigateway.NewLambdaIntegration(myLambda, nil)

	//now the api routes
	registerResource := api.Root().AddResource(jsii.String("register"), nil)
	registerResource.AddMethod(jsii.String("POST"), integration, nil)

	loginResource := api.Root().AddResource(jsii.String("login"), nil)
	loginResource.AddMethod(jsii.String("POST"), integration, nil)

	//give permision to lambda on the dynamodb table
	table.GrantReadWriteData(myLambda)
	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewAwsGoStack(app, "AwsGoStack", &AwsGoStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil
}
