package podtypes

import (
	host "github.com/kartverket/skiperator/api/v1beta1/podtypes"
	corev1 "k8s.io/api/core/v1"
)

type InternalPort struct {
	//+kubebuilder:validation:Required
	Name string `json:"name"`
	//+kubebuilder:validation:Required
	Port int32 `json:"port"`
	//+kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=TCP;UDP;SCTP
	// +kubebuilder:default:TCP
	Protocol corev1.Protocol `json:"protocol"`
}

func (src *InternalPort) toHost() *host.InternalPort {
	return &host.InternalPort{
		Name:     src.Name,
		Port:     src.Port,
		Protocol: src.Protocol,
	}
}
