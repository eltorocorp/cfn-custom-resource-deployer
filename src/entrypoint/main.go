package main

import (
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/eltorocorp/cfn-response/cfnhelper"
)

// Handler is the entrypoint for the deployer lambda.
func Handler(context *lambdacontext.LambdaContext, request *cfnhelper.Request) error {
	return nil
}

func main() {
	panic("This function is only supplied to squelch possible main.main runtime errors.")
}
