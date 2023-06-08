package main

import (
	"context"
	"errors"
	"fmt"
	"gambituser/awsgo"
	"gambituser/bd"
	"gambituser/models"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(cxt context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.StartedAws()

	if !ValidateParameters() {
		fmt.Println("Failed parameters. must send 'SecretName'")
		err := errors.New("failed to validate parameters SecretName")
		return event, err
	}

	var data models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = att
			fmt.Println("Email = " + data.UserEmail)
		case "sub":
			data.UserUUID = att
			fmt.Println("sub = " + data.UserUUID)
		}
	}

	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("Error reading secret: ", err.Error())
		return event, err
	}

	err = bd.SignUp(data)
	return event, err

}

func ValidateParameters() bool {
	var bringParams bool
	_, bringParams = os.LookupEnv("SecretName")
	return bringParams
}
