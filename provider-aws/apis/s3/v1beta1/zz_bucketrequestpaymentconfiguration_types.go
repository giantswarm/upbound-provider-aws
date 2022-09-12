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

type BucketRequestPaymentConfigurationObservation struct {

	// The bucket or bucket and expected_bucket_owner separated by a comma (,) if the latter is provided.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BucketRequestPaymentConfigurationParameters struct {

	// The name of the bucket.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/s3/v1beta1.Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// The account ID of the expected bucket owner.
	// +kubebuilder:validation:Optional
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// Specifies who pays for the download and request fees. Valid values: BucketOwner, Requester.
	// +kubebuilder:validation:Required
	Payer *string `json:"payer" tf:"payer,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// BucketRequestPaymentConfigurationSpec defines the desired state of BucketRequestPaymentConfiguration
type BucketRequestPaymentConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketRequestPaymentConfigurationParameters `json:"forProvider"`
}

// BucketRequestPaymentConfigurationStatus defines the observed state of BucketRequestPaymentConfiguration.
type BucketRequestPaymentConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketRequestPaymentConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BucketRequestPaymentConfiguration is the Schema for the BucketRequestPaymentConfigurations API. Provides an S3 bucket request payment configuration resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type BucketRequestPaymentConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketRequestPaymentConfigurationSpec   `json:"spec"`
	Status            BucketRequestPaymentConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketRequestPaymentConfigurationList contains a list of BucketRequestPaymentConfigurations
type BucketRequestPaymentConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketRequestPaymentConfiguration `json:"items"`
}

// Repository type metadata.
var (
	BucketRequestPaymentConfiguration_Kind             = "BucketRequestPaymentConfiguration"
	BucketRequestPaymentConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketRequestPaymentConfiguration_Kind}.String()
	BucketRequestPaymentConfiguration_KindAPIVersion   = BucketRequestPaymentConfiguration_Kind + "." + CRDGroupVersion.String()
	BucketRequestPaymentConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(BucketRequestPaymentConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketRequestPaymentConfiguration{}, &BucketRequestPaymentConfigurationList{})
}
