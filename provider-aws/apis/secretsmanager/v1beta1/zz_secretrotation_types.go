/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SecretRotationObservation struct {

	// Amazon Resource Name (ARN) of the secret.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies whether automatic rotation is enabled for this secret.
	RotationEnabled *bool `json:"rotationEnabled,omitempty" tf:"rotation_enabled,omitempty"`
}

type SecretRotationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Specifies the ARN of the Lambda function that can rotate the secret.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/lambda/v1beta1.Function
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	RotationLambdaArn *string `json:"rotationLambdaArn,omitempty" tf:"rotation_lambda_arn,omitempty"`

	// Reference to a Function in lambda to populate rotationLambdaArn.
	// +kubebuilder:validation:Optional
	RotationLambdaArnRef *v1.Reference `json:"rotationLambdaArnRef,omitempty" tf:"-"`

	// Selector for a Function in lambda to populate rotationLambdaArn.
	// +kubebuilder:validation:Optional
	RotationLambdaArnSelector *v1.Selector `json:"rotationLambdaArnSelector,omitempty" tf:"-"`

	// A structure that defines the rotation configuration for this secret. Defined below.
	// +kubebuilder:validation:Required
	RotationRules []SecretRotationRotationRulesParameters `json:"rotationRules" tf:"rotation_rules,omitempty"`

	// Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/secretsmanager/v1beta1.Secret
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SecretID *string `json:"secretId,omitempty" tf:"secret_id,omitempty"`

	// Reference to a Secret in secretsmanager to populate secretId.
	// +kubebuilder:validation:Optional
	SecretIDRef *v1.Reference `json:"secretIdRef,omitempty" tf:"-"`

	// Selector for a Secret in secretsmanager to populate secretId.
	// +kubebuilder:validation:Optional
	SecretIDSelector *v1.Selector `json:"secretIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SecretRotationRotationRulesObservation struct {
}

type SecretRotationRotationRulesParameters struct {

	// Specifies the number of days between automatic scheduled rotations of the secret.
	// +kubebuilder:validation:Required
	AutomaticallyAfterDays *float64 `json:"automaticallyAfterDays" tf:"automatically_after_days,omitempty"`
}

// SecretRotationSpec defines the desired state of SecretRotation
type SecretRotationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretRotationParameters `json:"forProvider"`
}

// SecretRotationStatus defines the observed state of SecretRotation.
type SecretRotationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretRotationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecretRotation is the Schema for the SecretRotations API. Provides a resource to manage AWS Secrets Manager secret rotation
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type SecretRotation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretRotationSpec   `json:"spec"`
	Status            SecretRotationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretRotationList contains a list of SecretRotations
type SecretRotationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretRotation `json:"items"`
}

// Repository type metadata.
var (
	SecretRotation_Kind             = "SecretRotation"
	SecretRotation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretRotation_Kind}.String()
	SecretRotation_KindAPIVersion   = SecretRotation_Kind + "." + CRDGroupVersion.String()
	SecretRotation_GroupVersionKind = CRDGroupVersion.WithKind(SecretRotation_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretRotation{}, &SecretRotationList{})
}
