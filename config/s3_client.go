package config


import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetObjectStorageClientConnection() (*s3.S3, error) {
	defaultResolver := endpoints.DefaultResolver()
	s3CustomResolverFn := func(service, region string, optFns ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
		if service == "s3" {
			return endpoints.ResolvedEndpoint{
				URL: ObjectStorageEndpoint,
			}, nil
		}
		return defaultResolver.EndpointFor(service, region, optFns...)
	}

	creds := credentials.NewStaticCredentials(ObjectStorageAccessKey,ObjectStorageSecretKey, "")
	pathStyle := true
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String("us-east-1"),
		Credentials:      creds,
		S3ForcePathStyle: &pathStyle,
		EndpointResolver: endpoints.ResolverFunc(s3CustomResolverFn),
	})
	if err != nil {
		return nil, err
	}
	s3Client := s3.New(sess)
	return s3Client, nil
}
