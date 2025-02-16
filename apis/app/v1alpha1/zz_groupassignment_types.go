/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type GroupAssignmentObservation struct {

	// App to associate group with
	AppID *string `json:"appId,omitempty" tf:"app_id,omitempty"`

	// Group associated with the application
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Priority of group assignment.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// JSON document containing [application profile](https://developer.okta.com/docs/reference/api/apps/#profile-object)
	Profile *string `json:"profile,omitempty" tf:"profile,omitempty"`

	// Retain the group assignment on destroy. If set to true, the resource will be removed from state but not from the Okta app.
	RetainAssignment *bool `json:"retainAssignment,omitempty" tf:"retain_assignment,omitempty"`
}

type GroupAssignmentParameters struct {

	// App to associate group with
	// +kubebuilder:validation:Optional
	AppID *string `json:"appId,omitempty" tf:"app_id,omitempty"`

	// Group associated with the application
	// +crossplane:generate:reference:type=github.com/healthcarecom/provider-okta/apis/group/v1alpha1.Group
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Reference to a Group in group to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDRef *v1.Reference `json:"groupIdRef,omitempty" tf:"-"`

	// Selector for a Group in group to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDSelector *v1.Selector `json:"groupIdSelector,omitempty" tf:"-"`

	// Priority of group assignment.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// JSON document containing [application profile](https://developer.okta.com/docs/reference/api/apps/#profile-object)
	// +kubebuilder:validation:Optional
	Profile *string `json:"profile,omitempty" tf:"profile,omitempty"`

	// Retain the group assignment on destroy. If set to true, the resource will be removed from state but not from the Okta app.
	// +kubebuilder:validation:Optional
	RetainAssignment *bool `json:"retainAssignment,omitempty" tf:"retain_assignment,omitempty"`
}

// GroupAssignmentSpec defines the desired state of GroupAssignment
type GroupAssignmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupAssignmentParameters `json:"forProvider"`
}

// GroupAssignmentStatus defines the observed state of GroupAssignment.
type GroupAssignmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupAssignmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GroupAssignment is the Schema for the GroupAssignments API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,okta}
type GroupAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.appId)",message="appId is a required parameter"
	Spec   GroupAssignmentSpec   `json:"spec"`
	Status GroupAssignmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupAssignmentList contains a list of GroupAssignments
type GroupAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GroupAssignment `json:"items"`
}

// Repository type metadata.
var (
	GroupAssignment_Kind             = "GroupAssignment"
	GroupAssignment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GroupAssignment_Kind}.String()
	GroupAssignment_KindAPIVersion   = GroupAssignment_Kind + "." + CRDGroupVersion.String()
	GroupAssignment_GroupVersionKind = CRDGroupVersion.WithKind(GroupAssignment_Kind)
)

func init() {
	SchemeBuilder.Register(&GroupAssignment{}, &GroupAssignmentList{})
}
