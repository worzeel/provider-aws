/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// IdentityPoolParameters defines the desired state of IdentityPool
type IdentityPoolParameters struct {
	// Region is which region the IdentityPool will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// Enables or disables the Basic (Classic) authentication flow. For more information,
	// see Identity Pools (Federated Identities) Authentication Flow (https://docs.aws.amazon.com/cognito/latest/developerguide/authentication-flow.html)
	// in the Amazon Cognito Developer Guide.
	AllowClassicFlow *bool `json:"allowClassicFlow,omitempty"`
	// The "domain" by which Cognito will refer to your users. This name acts as
	// a placeholder that allows your backend and the Cognito service to communicate
	// about the developer provider. For the DeveloperProviderName, you can use
	// letters as well as period (.), underscore (_), and dash (-).
	//
	// Once you have set a developer provider name, you cannot change it. Please
	// take care in setting this parameter.
	DeveloperProviderName *string `json:"developerProviderName,omitempty"`
	// A string that you provide.
	// +kubebuilder:validation:Required
	IdentityPoolName *string `json:"identityPoolName"`
	// Tags to assign to the identity pool. A tag is a label that you can apply
	// to identity pools to categorize and manage them in different ways, such as
	// by purpose, owner, environment, or other criteria.
	IdentityPoolTags map[string]*string `json:"identityPoolTags,omitempty"`
	// An array of Amazon Resource Names (ARNs) of the SAML provider for your identity
	// pool.
	SamlProviderARNs []*string `json:"samlProviderARNs,omitempty"`
	// Optional key:value pairs mapping provider names to provider app IDs.
	SupportedLoginProviders      map[string]*string `json:"supportedLoginProviders,omitempty"`
	CustomIdentityPoolParameters `json:",inline"`
}

// IdentityPoolSpec defines the desired state of IdentityPool
type IdentityPoolSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IdentityPoolParameters `json:"forProvider"`
}

// IdentityPoolObservation defines the observed state of IdentityPool
type IdentityPoolObservation struct {
	// TRUE if the identity pool supports unauthenticated logins.
	AllowUnauthenticatedIDentities *bool `json:"allowUnauthenticatedIDentities,omitempty"`
	// A list representing an Amazon Cognito user pool and its client ID.
	CognitoIdentityProviders []*Provider `json:"cognitoIdentityProviders,omitempty"`
	// An identity pool ID in the format REGION:GUID.
	IdentityPoolID *string `json:"identityPoolID,omitempty"`
	// The ARNs of the OpenID Connect providers.
	OpenIDConnectProviderARNs []*string `json:"openIDConnectProviderARNs,omitempty"`
}

// IdentityPoolStatus defines the observed state of IdentityPool.
type IdentityPoolStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IdentityPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityPool is the Schema for the IdentityPools API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type IdentityPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IdentityPoolSpec   `json:"spec"`
	Status            IdentityPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityPoolList contains a list of IdentityPools
type IdentityPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdentityPool `json:"items"`
}

// Repository type metadata.
var (
	IdentityPoolKind             = "IdentityPool"
	IdentityPoolGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IdentityPoolKind}.String()
	IdentityPoolKindAPIVersion   = IdentityPoolKind + "." + GroupVersion.String()
	IdentityPoolGroupVersionKind = GroupVersion.WithKind(IdentityPoolKind)
)

func init() {
	SchemeBuilder.Register(&IdentityPool{}, &IdentityPoolList{})
}