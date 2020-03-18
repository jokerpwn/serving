package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type FuncPool struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec RevisionSpec `json:"spec,omitempty"`

	// +optional
	Status RevisionStatus `json:"status,omitempty"`
}

// FuncPoolTemplateSpec describes the data a FuncPool should have when created from a template.
type FuncPoolTemplateSpec struct {
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec RevisionSpec `json:"spec,omitempty"`
}

// FuncPoolSpec holds the desired state of the FuncPool (from the client).
type FuncPoolSpec struct {
	corev1.PodSpec `json:",inline"`

	// ContainerConcurrency specifies the maximum allowed in-flight (concurrent)
	// requests per container of the Revision.  Defaults to `0` which means
	// concurrency to the application is not limited, and the system decides the
	// target concurrency for the autoscaler.
	// +optional
	ContainerConcurrency *int64 `json:"containerConcurrency,omitempty"`

	// TimeoutSeconds holds the max duration the instance is allowed for
	// responding to a request.  If unspecified, a system default will
	// be provided.
	// +optional
	TimeoutSeconds *int64 `json:"timeoutSeconds,omitempty"`
}

const (
	// RevisionConditionReady is set when the revision is starting to materialize
	// runtime resources, and becomes true when those resources are ready.
	FuncPoolConditionReady = apis.ConditionReady

	// FuncPoolConditionResourcesAvailable is set when underlying
	// Kubernetes resources have been provisioned.
	FuncPoolConditionResourcesAvailable apis.ConditionType = "ResourcesAvailable"

	// FuncPoolConditionContainerHealthy is set when the revision readiness check completes.
	FuncPoolConditionContainerHealthy apis.ConditionType = "ContainerHealthy"

	// FuncPoolConditionActive is set when the revision is receiving traffic.
	FuncPoolConditionActive apis.ConditionType = "Active"
)
// IsRevisionCondition returns true if the ConditionType is a revision condition type
func IsFuncPoolCondition(t apis.ConditionType) bool {
	switch t {
	case
		FuncPoolConditionReady,
		FuncPoolConditionResourcesAvailable,
		FuncPoolConditionContainerHealthy,
		FuncPoolConditionActive:
		return true
	}
	return false
}
type FuncPoolStatus struct {
	duckv1.Status `json:",inline"`

	// ServiceName holds the name of a core Kubernetes Service resource that
	// load balances over the pods backing this FuncPool.
	// +optional
	ServiceName string `json:"serviceName,omitempty"`

	// LogURL specifies the generated logging url for this particular revision
	// based on the FuncPool url template specified in the controller's config.
	// +optional
	LogURL string `json:"logUrl,omitempty"`

	// ImageDigest holds the resolved digest for the image specified
	// within .Spec.Container.Image. The digest is resolved during the creation
	// of FuncPool. This field holds the digest value regardless of whether
	// a tag or digest was originally specified in the Container object. It
	// may be empty if the image comes from a registry listed to skip resolution.
	// +optional
	ImageDigest string `json:"imageDigest,omitempty"`
}
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RevisionList is a list of FuncPool resources
type FuncPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []FuncPool `json:"items"`
}
