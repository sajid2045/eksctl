package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSDMSCertificate AWS CloudFormation Resource (AWS::DMS::Certificate)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-certificate.html
type AWSDMSCertificate struct {

	// CertificateIdentifier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-certificate.html#cfn-dms-certificate-certificateidentifier
	CertificateIdentifier *Value `json:"CertificateIdentifier,omitempty"`

	// CertificatePem AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-certificate.html#cfn-dms-certificate-certificatepem
	CertificatePem *Value `json:"CertificatePem,omitempty"`

	// CertificateWallet AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-certificate.html#cfn-dms-certificate-certificatewallet
	CertificateWallet *Value `json:"CertificateWallet,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDMSCertificate) AWSCloudFormationType() string {
	return "AWS::DMS::Certificate"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSDMSCertificate) MarshalJSON() ([]byte, error) {
	type Properties AWSDMSCertificate
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSDMSCertificate) UnmarshalJSON(b []byte) error {
	type Properties AWSDMSCertificate
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSDMSCertificate(*res.Properties)
	}

	return nil
}

// GetAllAWSDMSCertificateResources retrieves all AWSDMSCertificate items from an AWS CloudFormation template
func (t *Template) GetAllAWSDMSCertificateResources() map[string]AWSDMSCertificate {
	results := map[string]AWSDMSCertificate{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSDMSCertificate:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::DMS::Certificate" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSDMSCertificate{}
						if err := result.UnmarshalJSON(b); err == nil {
							results[name] = *result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSDMSCertificateWithName retrieves all AWSDMSCertificate items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDMSCertificateWithName(name string) (AWSDMSCertificate, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSDMSCertificate:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::DMS::Certificate" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSDMSCertificate{}
						if err := result.UnmarshalJSON(b); err == nil {
							return *result, nil
						}
					}
				}
			}
		}
	}
	return AWSDMSCertificate{}, errors.New("resource not found")
}
