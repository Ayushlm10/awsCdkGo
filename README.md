# Deploy resources on aws using CDK(and go)!

A basic setup for deploying lambda , apigateway and a dynamodb database using amazon cdk in golang.

This is what we deploy:

Cloudformation stack:

![image](https://github.com/Ayushlm10/awsCdkGo/assets/51413362/446ff2b8-1384-40ea-a4c5-042d1f053659)

A simple user dynamodb table:

![image](https://github.com/Ayushlm10/awsCdkGo/assets/51413362/a30dac13-76e5-43b5-9e35-765b26232300)

API gateway with the following API routes:

![image](https://github.com/Ayushlm10/awsCdkGo/assets/51413362/cf7ad394-11a4-4784-8a0b-1f059e406d6a)

.. and a few lambda functions (check out lambda/api.go).

![image](https://github.com/Ayushlm10/awsCdkGo/assets/51413362/eceffd6e-82da-4d1a-b7d4-6444b9d62d6d)


Setup:

- AWS account (a free account will work)
- Setup aws cli and cdk.
- cd into the lambda directory.
- `make build`
- cd back to root of the project.
- `cdk diff` to check what is being deployed.
- `cdk deploy`

A lambda function to register and login users will be deployed. This lambda function sits behind amazon api gateway and stores data in a dynamodb database.

## Useful commands

- `cdk deploy` deploy this stack to your default AWS account/region
- `cdk diff` compare deployed stack with current state
- `cdk synth` emits the synthesized CloudFormation template
- `go test` run unit tests
