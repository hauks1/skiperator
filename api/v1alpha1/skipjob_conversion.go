package v1alpha1

import (
	"github.com/kartverket/skiperator/api/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (src *SKIPJob) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.SKIPJob)

	// Copy metadata
	dst.ObjectMeta = src.ObjectMeta

	dst.Spec = src.Spec.toHost()
	dst.Status = src.Status.toHost()

	return nil
}

// func (dst *SKIPJob) ConvertFrom(srcRaw conversion.Hub) error {
// 	src := srcRaw.(*v1beta1.SKIPJob)

// 	// Copy metadata
// 	dst.ObjectMeta = src.ObjectMeta

// 	// Copy Job and Cron settings
// 	dst.Spec.Job = src.Spec.Job
// 	dst.Spec.Cron = src.Spec.Cron

// 	// Copy IstioSettings and Prometheus
// 	dst.Spec.IstioSettings = src.Spec.IstioSettings
// 	dst.Spec.Prometheus = src.Spec.Prometheus

// 	// Copy primitive fields to Container
// 	dst.Spec.Container.Image = src.Spec.Image
// 	dst.Spec.Container.Priority = src.Spec.Priority
// 	dst.Spec.Container.Command = src.Spec.Command
// 	dst.Spec.Container.Env = src.Spec.Env
// 	dst.Spec.Container.RestartPolicy = src.Spec.RestartPolicy

// 	// Convert custom struct fields
// 	if src.Spec.Resources != nil {
// 		dst.Spec.Container.Resources = convertResourceRequirementsReverse(src.Spec.Resources)
// 	}

// 	if src.Spec.AccessPolicy != nil {
// 		dst.Spec.Container.AccessPolicy = convertAccessPolicyReverse(src.Spec.AccessPolicy)
// 	}

// 	if src.Spec.GCP != nil {
// 		dst.Spec.Container.GCP = convertGCPReverse(src.Spec.GCP)
// 	}

// 	if src.Spec.PodSettings != nil {
// 		dst.Spec.Container.PodSettings = convertPodSettingsReverse(src.Spec.PodSettings)
// 	}

// 	if src.Spec.Liveness != nil {
// 		dst.Spec.Container.Liveness = convertProbeReverse(src.Spec.Liveness)
// 	}

// 	if src.Spec.Readiness != nil {
// 		dst.Spec.Container.Readiness = convertProbeReverse(src.Spec.Readiness)
// 	}

// 	if src.Spec.Startup != nil {
// 		dst.Spec.Container.Startup = convertProbeReverse(src.Spec.Startup)
// 	}

// 	dst.Spec.Container.EnvFrom = convertEnvFromSliceReverse(src.Spec.EnvFrom)
// 	dst.Spec.Container.FilesFrom = convertFilesFromSliceReverse(src.Spec.FilesFrom)
// 	dst.Spec.Container.AdditionalPorts = convertInternalPortSliceReverse(src.Spec.AdditionalPorts)

// 	// Copy status
// 	dst.Status = convertSkiperatorStatus(src.Status)

// 	src.status

// 	return nil
// }
