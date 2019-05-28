package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func main() {
	profile := flag.String("profile", "default", "Name of the AWS profile to use.")
	serial := flag.String("serial-number", "", "The serial number of your MFA device.")
	duration := flag.Int64("duration-seconds", 43200, "The amount of time your generated STS token is valid for.")
	flag.Parse()

	code := flag.Arg(0)

	os.Setenv("AWS_PROFILE", *profile)

	sess := session.Must(session.NewSession())
	_, err := sess.Config.Credentials.Get()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error retrieving credentials.")
		os.Exit(1)
	}

	svc := sts.New(session.New())
	input := &sts.GetSessionTokenInput{
		DurationSeconds: aws.Int64(*duration),
		SerialNumber:    aws.String(*serial),
		TokenCode:       aws.String(code),
	}

	result, err := svc.GetSessionToken(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case sts.ErrCodeRegionDisabledException:
				fmt.Println(sts.ErrCodeRegionDisabledException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println("aws_access_key_id = ", *result.Credentials.AccessKeyId)
	fmt.Println("aws_secret_access_key = ", *result.Credentials.SecretAccessKey)
	fmt.Println("aws_session_token = ", *result.Credentials.SessionToken)
}
