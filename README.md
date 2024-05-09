# Deploy resources on aws using CDK(and go)!

A basic setup for deploying lambda , apigateway and a dynamodb database using amazon cdk in golang.

Setup:

- AWS account (a free account will work)
- Setup aws cli and cdk.
- cd into the lambda directory.
- `make build`
- cd back to root of the project.
- `cdk deploy`

A lambda function to register and login users will be deployed. This lambda function sits behind amazon api gateway and stores data in a dynamodb database.

## Useful commands

- `cdk deploy` deploy this stack to your default AWS account/region
- `cdk diff` compare deployed stack with current state
- `cdk synth` emits the synthesized CloudFormation template
- `go test` run unit tests
