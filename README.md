# cfn-custom-resource-deployer
AWS Lambda for deploying  and managing custom CloudFormation resources.

CloudFormation supports the deployment/management of custom resources. This means that CloudFormation can be used to deploy effectively any resource, even those outside of the AWS ecosystem.

This repository includes the definition of a Lambda function that handles creating/updating/deleting custom resources used by El Toro. However, it can be easily extended to suit any needs.

For more information on custom resources, please read the [AWS documentation on CloudFormation custom resources](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-custom-resources.html).
