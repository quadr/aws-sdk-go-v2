// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a batch job to transcribe medical speech to text.
func (c *Client) StartMedicalTranscriptionJob(ctx context.Context, params *StartMedicalTranscriptionJobInput, optFns ...func(*Options)) (*StartMedicalTranscriptionJobOutput, error) {
	if params == nil {
		params = &StartMedicalTranscriptionJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartMedicalTranscriptionJob", params, optFns, c.addOperationStartMedicalTranscriptionJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartMedicalTranscriptionJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartMedicalTranscriptionJobInput struct {

	// The language code for the language spoken in the input media file. US English
	// (en-US) is the valid value for medical transcription jobs. Any other value you
	// enter for language code results in a BadRequestException error.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// Describes the input media file in a transcription request.
	//
	// This member is required.
	Media *types.Media

	// The name of the medical transcription job. You can't use the strings "." or ".."
	// by themselves as the job name. The name must also be unique within an Amazon Web
	// Services account. If you try to create a medical transcription job with the same
	// name as a previous medical transcription job, you get a ConflictException error.
	//
	// This member is required.
	MedicalTranscriptionJobName *string

	// The Amazon S3 location where the transcription is stored. You must set
	// OutputBucketName for Amazon Transcribe Medical to store the transcription
	// results. Your transcript appears in the S3 location you specify. When you call
	// the GetMedicalTranscriptionJob, the operation returns this location in the
	// TranscriptFileUri field. The S3 bucket must have permissions that allow Amazon
	// Transcribe Medical to put files in the bucket. For more information, see
	// Permissions Required for IAM User Roles
	// (https://docs.aws.amazon.com/transcribe/latest/dg/security_iam_id-based-policy-examples.html#auth-role-iam-user).
	// You can specify an Amazon Web Services Key Management Service (KMS) key to
	// encrypt the output of your transcription using the OutputEncryptionKMSKeyId
	// parameter. If you don't specify a KMS key, Amazon Transcribe Medical uses the
	// default Amazon S3 key for server-side encryption of transcripts that are placed
	// in your S3 bucket.
	//
	// This member is required.
	OutputBucketName *string

	// The medical specialty of any clinician speaking in the input media.
	//
	// This member is required.
	Specialty types.Specialty

	// The type of speech in the input audio. CONVERSATION refers to conversations
	// between two or more speakers, e.g., a conversations between doctors and
	// patients. DICTATION refers to single-speaker dictated speech, such as clinical
	// notes.
	//
	// This member is required.
	Type types.Type

	// You can configure Amazon Transcribe Medical to label content in the
	// transcription output. If you specify PHI, Amazon Transcribe Medical labels the
	// personal health information (PHI) that it identifies in the transcription
	// output.
	ContentIdentificationType types.MedicalContentIdentificationType

	// A map of plain text, non-secret key:value pairs, known as encryption context
	// pairs, that provide an added layer of security for your data.
	KMSEncryptionContext map[string]string

	// The audio format of the input media file.
	MediaFormat types.MediaFormat

	// The sample rate, in Hertz, of the audio track in the input media file. If you do
	// not specify the media sample rate, Amazon Transcribe Medical determines the
	// sample rate. If you specify the sample rate, it must match the rate detected by
	// Amazon Transcribe Medical. In most cases, you should leave the
	// MediaSampleRateHertz field blank and let Amazon Transcribe Medical determine the
	// sample rate.
	MediaSampleRateHertz *int32

	// The Amazon Resource Name (ARN) of the Amazon Web Services Key Management Service
	// (KMS) key used to encrypt the output of the transcription job. The user calling
	// the StartMedicalTranscriptionJob operation must have permission to use the
	// specified KMS key. You use either of the following to identify a KMS key in the
	// current account:
	//
	// * KMS Key ID: "1234abcd-12ab-34cd-56ef-1234567890ab"
	//
	// * KMS
	// Key Alias: "alias/ExampleAlias"
	//
	// You can use either of the following to identify
	// a KMS key in the current account or another account:
	//
	// * Amazon Resource Name
	// (ARN) of a KMS key in the current account or another account:
	// "arn:aws:kms:region:account ID:key/1234abcd-12ab-34cd-56ef-1234567890ab"
	//
	// * ARN
	// of a KMS Key Alias: "arn:aws:kms:region:account ID:alias/ExampleAlias"
	//
	// If you
	// don't specify an encryption key, the output of the medical transcription job is
	// encrypted with the default Amazon S3 key (SSE-S3). If you specify a KMS key to
	// encrypt your output, you must also specify an output location in the
	// OutputBucketName parameter.
	OutputEncryptionKMSKeyId *string

	// You can specify a location in an Amazon S3 bucket to store the output of your
	// medical transcription job. If you don't specify an output key, Amazon Transcribe
	// Medical stores the output of your transcription job in the Amazon S3 bucket you
	// specified. By default, the object key is "your-transcription-job-name.json". You
	// can use output keys to specify the Amazon S3 prefix and file name of the
	// transcription output. For example, specifying the Amazon S3 prefix,
	// "folder1/folder2/", as an output key would lead to the output being stored as
	// "folder1/folder2/your-transcription-job-name.json". If you specify
	// "my-other-job-name.json" as the output key, the object key is changed to
	// "my-other-job-name.json". You can use an output key to change both the prefix
	// and the file name, for example "folder/my-other-job-name.json". If you specify
	// an output key, you must also specify an S3 bucket in the OutputBucketName
	// parameter.
	OutputKey *string

	// Optional settings for the medical transcription job.
	Settings *types.MedicalTranscriptionSetting

	// Add tags to an Amazon Transcribe medical transcription job.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type StartMedicalTranscriptionJobOutput struct {

	// A batch job submitted to transcribe medical speech to text.
	MedicalTranscriptionJob *types.MedicalTranscriptionJob

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartMedicalTranscriptionJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartMedicalTranscriptionJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartMedicalTranscriptionJob{}, middleware.After)
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
	if err = addOpStartMedicalTranscriptionJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartMedicalTranscriptionJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartMedicalTranscriptionJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "StartMedicalTranscriptionJob",
	}
}
