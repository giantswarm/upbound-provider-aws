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

type EndpointConfigurationObservation struct {
}

type EndpointConfigurationParameters struct {

	// Indicates whether client IP address preservation is enabled for an Application Load Balancer endpoint. See the AWS documentation for more details. The default value is false.
	// +kubebuilder:validation:Optional
	ClientIPPreservationEnabled *bool `json:"clientIpPreservationEnabled,omitempty" tf:"client_ip_preservation_enabled,omitempty"`

	// An ID for the endpoint. If the endpoint is a Network Load Balancer or Application Load Balancer, this is the Amazon Resource Name (ARN) of the resource. If the endpoint is an Elastic IP address, this is the Elastic IP address allocation ID.
	// +kubebuilder:validation:Optional
	EndpointID *string `json:"endpointId,omitempty" tf:"endpoint_id,omitempty"`

	// The weight associated with the endpoint. When you add weights to endpoints, you configure AWS Global Accelerator to route traffic based on proportions that you specify.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type EndpointGroupObservation struct {

	// The Amazon Resource Name (ARN) of the endpoint group.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The Amazon Resource Name (ARN) of the endpoint group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EndpointGroupParameters struct {

	// The list of endpoint objects. Fields documented below.
	// +kubebuilder:validation:Optional
	EndpointConfiguration []EndpointConfigurationParameters `json:"endpointConfiguration,omitempty" tf:"endpoint_configuration,omitempty"`

	// The name of the AWS Region where the endpoint group is located.
	// +kubebuilder:validation:Optional
	EndpointGroupRegion *string `json:"endpointGroupRegion,omitempty" tf:"endpoint_group_region,omitempty"`

	// The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
	// +kubebuilder:validation:Optional
	HealthCheckIntervalSeconds *float64 `json:"healthCheckIntervalSeconds,omitempty" tf:"health_check_interval_seconds,omitempty"`

	// If the protocol is HTTP/S, then this specifies the path that is the destination for health check targets. The default value is slash (/).
	// +kubebuilder:validation:Optional
	HealthCheckPath *string `json:"healthCheckPath,omitempty" tf:"health_check_path,omitempty"`

	// The port that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default port is the listener port that this endpoint group is associated with. If listener port is a list of ports, Global Accelerator uses the first port in the list.
	// +kubebuilder:validation:Optional
	HealthCheckPort *float64 `json:"healthCheckPort,omitempty" tf:"health_check_port,omitempty"`

	// The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
	// +kubebuilder:validation:Optional
	HealthCheckProtocol *string `json:"healthCheckProtocol,omitempty" tf:"health_check_protocol,omitempty"`

	// The Amazon Resource Name (ARN) of the listener.
	// +crossplane:generate:reference:type=Listener
	// +kubebuilder:validation:Optional
	ListenerArn *string `json:"listenerArn,omitempty" tf:"listener_arn,omitempty"`

	// Reference to a Listener to populate listenerArn.
	// +kubebuilder:validation:Optional
	ListenerArnRef *v1.Reference `json:"listenerArnRef,omitempty" tf:"-"`

	// Selector for a Listener to populate listenerArn.
	// +kubebuilder:validation:Optional
	ListenerArnSelector *v1.Selector `json:"listenerArnSelector,omitempty" tf:"-"`

	// Override specific listener ports used to route traffic to endpoints that are part of this endpoint group. Fields documented below.
	// +kubebuilder:validation:Optional
	PortOverride []PortOverrideParameters `json:"portOverride,omitempty" tf:"port_override,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
	// +kubebuilder:validation:Optional
	ThresholdCount *float64 `json:"thresholdCount,omitempty" tf:"threshold_count,omitempty"`

	// The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.
	// +kubebuilder:validation:Optional
	TrafficDialPercentage *float64 `json:"trafficDialPercentage,omitempty" tf:"traffic_dial_percentage,omitempty"`
}

type PortOverrideObservation struct {
}

type PortOverrideParameters struct {

	// The endpoint port that you want a listener port to be mapped to. This is the port on the endpoint, such as the Application Load Balancer or Amazon EC2 instance.
	// +kubebuilder:validation:Required
	EndpointPort *float64 `json:"endpointPort" tf:"endpoint_port,omitempty"`

	// The listener port that you want to map to a specific endpoint port. This is the port that user traffic arrives to the Global Accelerator on.
	// +kubebuilder:validation:Required
	ListenerPort *float64 `json:"listenerPort" tf:"listener_port,omitempty"`
}

// EndpointGroupSpec defines the desired state of EndpointGroup
type EndpointGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EndpointGroupParameters `json:"forProvider"`
}

// EndpointGroupStatus defines the observed state of EndpointGroup.
type EndpointGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EndpointGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointGroup is the Schema for the EndpointGroups API. Provides a Global Accelerator endpoint group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EndpointGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointGroupSpec   `json:"spec"`
	Status            EndpointGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointGroupList contains a list of EndpointGroups
type EndpointGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EndpointGroup `json:"items"`
}

// Repository type metadata.
var (
	EndpointGroup_Kind             = "EndpointGroup"
	EndpointGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EndpointGroup_Kind}.String()
	EndpointGroup_KindAPIVersion   = EndpointGroup_Kind + "." + CRDGroupVersion.String()
	EndpointGroup_GroupVersionKind = CRDGroupVersion.WithKind(EndpointGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&EndpointGroup{}, &EndpointGroupList{})
}
