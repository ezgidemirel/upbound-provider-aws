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

type AuthObservation struct {
}

type AuthParameters struct {

	// +kubebuilder:validation:Optional
	AuthScheme *string `json:"authScheme,omitempty" tf:"auth_scheme,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	IAMAuth *string `json:"iamAuth,omitempty" tf:"iam_auth,omitempty"`

	// +kubebuilder:validation:Optional
	SecretArn *string `json:"secretArn,omitempty" tf:"secret_arn,omitempty"`

	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type ProxyObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ProxyParameters struct {

	// +kubebuilder:validation:Required
	Auth []AuthParameters `json:"auth" tf:"auth,omitempty"`

	// +kubebuilder:validation:Optional
	DebugLogging *bool `json:"debugLogging,omitempty" tf:"debug_logging,omitempty"`

	// +kubebuilder:validation:Required
	EngineFamily *string `json:"engineFamily" tf:"engine_family,omitempty"`

	// +kubebuilder:validation:Optional
	IdleClientTimeout *float64 `json:"idleClientTimeout,omitempty" tf:"idle_client_timeout,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RequireTLS *bool `json:"requireTls,omitempty" tf:"require_tls,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	VPCSecurityGroupIdRefs []v1.Reference `json:"vpcSecurityGroupIdRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VPCSecurityGroupIdSelector *v1.Selector `json:"vpcSecurityGroupIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=VPCSecurityGroupIdRefs
	// +crossplane:generate:reference:selectorFieldName=VPCSecurityGroupIdSelector
	// +kubebuilder:validation:Optional
	VPCSecurityGroupIds []*string `json:"vpcSecurityGroupIds,omitempty" tf:"vpc_security_group_ids,omitempty"`

	// +kubebuilder:validation:Required
	VPCSubnetIds []*string `json:"vpcSubnetIds" tf:"vpc_subnet_ids,omitempty"`
}

// ProxySpec defines the desired state of Proxy
type ProxySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProxyParameters `json:"forProvider"`
}

// ProxyStatus defines the observed state of Proxy.
type ProxyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProxyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Proxy is the Schema for the Proxys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Proxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProxySpec   `json:"spec"`
	Status            ProxyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProxyList contains a list of Proxys
type ProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Proxy `json:"items"`
}

// Repository type metadata.
var (
	Proxy_Kind             = "Proxy"
	Proxy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Proxy_Kind}.String()
	Proxy_KindAPIVersion   = Proxy_Kind + "." + CRDGroupVersion.String()
	Proxy_GroupVersionKind = CRDGroupVersion.WithKind(Proxy_Kind)
)

func init() {
	SchemeBuilder.Register(&Proxy{}, &ProxyList{})
}
