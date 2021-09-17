// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an artifact. An artifact is a lineage tracking entity that represents a
// URI addressable object or data. Some examples are the S3 URI of a dataset and
// the ECR registry path of an image. For more information, see Amazon SageMaker ML
// Lineage Tracking
// (https://docs.aws.amazon.com/sagemaker/latest/dg/lineage-tracking.html).
func (c *Client) CreateArtifact(ctx context.Context, params *CreateArtifactInput, optFns ...func(*Options)) (*CreateArtifactOutput, error) {
	if params == nil {
		params = &CreateArtifactInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateArtifact", params, optFns, c.addOperationCreateArtifactMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateArtifactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateArtifactInput struct {

	// The artifact type.
	//
	// This member is required.
	ArtifactType *string

	// The ID, ID type, and URI of the source.
	//
	// This member is required.
	Source *types.ArtifactSource

	// The name of the artifact. Must be unique to your account in an Amazon Web
	// Services Region.
	ArtifactName *string

	// Metadata properties of the tracking entity, trial, or trial component.
	MetadataProperties *types.MetadataProperties

	// A list of properties to add to the artifact.
	Properties map[string]string

	// A list of tags to apply to the artifact.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateArtifactOutput struct {

	// The Amazon Resource Name (ARN) of the artifact.
	ArtifactArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateArtifactMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateArtifact{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateArtifact{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateArtifactValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateArtifact(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateArtifact(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateArtifact",
	}
}
