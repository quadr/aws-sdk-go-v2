// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The certificate authority certificate you are importing does not comply with
// conditions specified in the certificate that signed it.
type CertificateMismatchException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *CertificateMismatchException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CertificateMismatchException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CertificateMismatchException) ErrorCode() string             { return "CertificateMismatchException" }
func (e *CertificateMismatchException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A previous update to your private CA is still ongoing.
type ConcurrentModificationException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ConcurrentModificationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentModificationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentModificationException) ErrorCode() string {
	return "ConcurrentModificationException"
}
func (e *ConcurrentModificationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// One or more of the specified arguments was not valid.
type InvalidArgsException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidArgsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidArgsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidArgsException) ErrorCode() string             { return "InvalidArgsException" }
func (e *InvalidArgsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested Amazon Resource Name (ARN) does not refer to an existing resource.
type InvalidArnException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidArnException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidArnException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidArnException) ErrorCode() string             { return "InvalidArnException" }
func (e *InvalidArnException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The token specified in the NextToken argument is not valid. Use the token
// returned from your previous call to ListCertificateAuthorities
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_ListCertificateAuthorities.html).
type InvalidNextTokenException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidNextTokenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidNextTokenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidNextTokenException) ErrorCode() string             { return "InvalidNextTokenException" }
func (e *InvalidNextTokenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource policy is invalid or is missing a required statement. For general
// information about IAM policy and statement structure, see Overview of JSON
// Policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#access_policies-json).
type InvalidPolicyException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidPolicyException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidPolicyException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidPolicyException) ErrorCode() string             { return "InvalidPolicyException" }
func (e *InvalidPolicyException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request action cannot be performed or is prohibited.
type InvalidRequestException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRequestException) ErrorCode() string             { return "InvalidRequestException" }
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The state of the private CA does not allow this action to occur.
type InvalidStateException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidStateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidStateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidStateException) ErrorCode() string             { return "InvalidStateException" }
func (e *InvalidStateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The tag associated with the CA is not valid. The invalid argument is contained
// in the message field.
type InvalidTagException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidTagException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidTagException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidTagException) ErrorCode() string             { return "InvalidTagException" }
func (e *InvalidTagException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An Amazon Web Services Private CA quota has been exceeded. See the exception
// message returned to determine the quota that was exceeded.
type LimitExceededException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The current action was prevented because it would lock the caller out from
// performing subsequent actions. Verify that the specified parameters would not
// result in the caller being denied access to the resource.
type LockoutPreventedException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *LockoutPreventedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LockoutPreventedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LockoutPreventedException) ErrorCode() string             { return "LockoutPreventedException" }
func (e *LockoutPreventedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// One or more fields in the certificate are invalid.
type MalformedCertificateException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *MalformedCertificateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MalformedCertificateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MalformedCertificateException) ErrorCode() string             { return "MalformedCertificateException" }
func (e *MalformedCertificateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The certificate signing request is invalid.
type MalformedCSRException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *MalformedCSRException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MalformedCSRException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MalformedCSRException) ErrorCode() string             { return "MalformedCSRException" }
func (e *MalformedCSRException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The designated permission has already been given to the user.
type PermissionAlreadyExistsException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *PermissionAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PermissionAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PermissionAlreadyExistsException) ErrorCode() string {
	return "PermissionAlreadyExistsException"
}
func (e *PermissionAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Your request has already been completed.
type RequestAlreadyProcessedException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *RequestAlreadyProcessedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RequestAlreadyProcessedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RequestAlreadyProcessedException) ErrorCode() string {
	return "RequestAlreadyProcessedException"
}
func (e *RequestAlreadyProcessedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request has failed for an unspecified reason.
type RequestFailedException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *RequestFailedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RequestFailedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RequestFailedException) ErrorCode() string             { return "RequestFailedException" }
func (e *RequestFailedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Your request is already in progress.
type RequestInProgressException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *RequestInProgressException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RequestInProgressException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RequestInProgressException) ErrorCode() string             { return "RequestInProgressException" }
func (e *RequestInProgressException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A resource such as a private CA, S3 bucket, certificate, audit report, or policy
// cannot be found.
type ResourceNotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You can associate up to 50 tags with a private CA. Exception information is
// contained in the exception message field.
type TooManyTagsException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *TooManyTagsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyTagsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyTagsException) ErrorCode() string             { return "TooManyTagsException" }
func (e *TooManyTagsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
