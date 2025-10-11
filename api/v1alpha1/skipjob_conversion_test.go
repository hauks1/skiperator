package v1alpha1

import (
	"testing"

	"github.com/kartverket/skiperator/api/v1alpha1/istiotypes"
	"github.com/kartverket/skiperator/api/v1alpha1/podtypes"
	"github.com/kartverket/skiperator/api/v1beta1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestSKIPJobConvertTo(t *testing.T) {
	tests := []struct {
		name    string
		src     *SKIPJob
		wantErr bool
	}{
		{
			name: "basic skipjob conversion",
			src: &SKIPJob{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-job",
					Namespace: "test-namespace",
				},
				Spec: SKIPJobSpec{
					Job: &JobSettings{
						ActiveDeadlineSeconds:   int64Ptr(3600),
						BackoffLimit:            int32Ptr(3),
						Suspend:                 boolPtr(false),
						TTLSecondsAfterFinished: int32Ptr(86400),
					},
					Container: ContainerSettings{
						Image:    "test-image:latest",
						Priority: "medium",
						Command:  []string{"echo", "hello"},
						Env: []corev1.EnvVar{
							{Name: "TEST_ENV", Value: "test_value"},
						},
						RestartPolicy: restartPolicyPtr(corev1.RestartPolicyNever),
					},
				},
				Status: SkiperatorStatus{
					Summary: Status{
						Status:  SYNCED,
						Message: "All good",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "skipjob with cron settings",
			src: &SKIPJob{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-cron-job",
					Namespace: "test-namespace",
				},
				Spec: SKIPJobSpec{
					Job: &JobSettings{
						BackoffLimit: int32Ptr(6),
					},
					Cron: &CronSettings{
						Schedule:                "0 * * * *",
						ConcurrencyPolicy:       batchv1.AllowConcurrent,
						StartingDeadlineSeconds: int64Ptr(300),
						Suspend:                 boolPtr(false),
						TimeZone:                stringPtr("Europe/Oslo"),
					},
					Container: ContainerSettings{
						Image:    "cron-image:latest",
						Priority: "high",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "skipjob with access policy",
			src: &SKIPJob{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-access-job",
					Namespace: "test-namespace",
				},
				Spec: SKIPJobSpec{
					Container: ContainerSettings{
						Image: "test-image:latest",
						AccessPolicy: &podtypes.AccessPolicy{
							Outbound: &podtypes.OutboundPolicy{
								External: []podtypes.ExternalRule{
									{Host: "example.com"},
								},
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "skipjob with resources and probes",
			src: &SKIPJob{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-resources-job",
					Namespace: "test-namespace",
				},
				Spec: SKIPJobSpec{
					Container: ContainerSettings{
						Image: "test-image:latest",
						Resources: &podtypes.ResourceRequirements{
							Limits: &podtypes.ResourceList{
								CPU:    "1000m",
								Memory: "1Gi",
							},
							Requests: &podtypes.ResourceList{
								CPU:    "100m",
								Memory: "128Mi",
							},
						},
						Liveness: &podtypes.Probe{
							Path: "/healthz",
							Port: 8080,
						},
						Readiness: &podtypes.Probe{
							Path: "/ready",
							Port: 8080,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "skipjob with gcp settings",
			src: &SKIPJob{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-gcp-job",
					Namespace: "test-namespace",
				},
				Spec: SKIPJobSpec{
					Container: ContainerSettings{
						Image: "test-image:latest",
						GCP: &podtypes.GCP{
							Auth: &podtypes.Auth{
								ServiceAccount: "test-sa@project.iam.gserviceaccount.com",
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "skipjob with istio settings",
			src: &SKIPJob{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-istio-job",
					Namespace: "test-namespace",
				},
				Spec: SKIPJobSpec{
					IstioSettings: &istiotypes.IstioSettingsBase{
						Telemetry: &istiotypes.Telemetry{
							Tracing: []*istiotypes.Tracing{
								{RandomSamplingPercentage: 50},
							},
						},
					},
					Container: ContainerSettings{
						Image: "test-image:latest",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "skipjob with env and files",
			src: &SKIPJob{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-env-files-job",
					Namespace: "test-namespace",
				},
				Spec: SKIPJobSpec{
					Container: ContainerSettings{
						Image: "test-image:latest",
						Env: []corev1.EnvVar{
							{Name: "VAR1", Value: "value1"},
							{Name: "VAR2", Value: "value2"},
						},
						EnvFrom: []podtypes.EnvFrom{
							{ConfigMap: "test-configmap"},
							{Secret: "test-secret"},
						},
						FilesFrom: []podtypes.FilesFrom{
							{
								MountPath: "/etc/config",
								ConfigMap: "config-map",
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "skipjob with additional ports",
			src: &SKIPJob{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-ports-job",
					Namespace: "test-namespace",
				},
				Spec: SKIPJobSpec{
					Container: ContainerSettings{
						Image: "test-image:latest",
						AdditionalPorts: []podtypes.InternalPort{
							{Name: "metrics", Port: 9090, Protocol: "TCP"},
							{Name: "admin", Port: 8081, Protocol: "TCP"},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "skipjob with pod settings",
			src: &SKIPJob{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-pod-settings-job",
					Namespace: "test-namespace",
				},
				Spec: SKIPJobSpec{
					Container: ContainerSettings{
						Image: "test-image:latest",
						PodSettings: &podtypes.PodSettings{
							DisablePodSpreadTopologyConstraints: boolPtr(true),
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dst := &v1beta1.SKIPJob{}
			err := tt.src.ConvertTo(dst)
			
			if (err != nil) != tt.wantErr {
				t.Errorf("SKIPJob.ConvertTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			// Verify metadata is copied
			if dst.Name != tt.src.Name {
				t.Errorf("Name not copied correctly: got %v, want %v", dst.Name, tt.src.Name)
			}
			if dst.Namespace != tt.src.Namespace {
				t.Errorf("Namespace not copied correctly: got %v, want %v", dst.Namespace, tt.src.Namespace)
			}

			// Verify basic fields are converted
			if tt.src.Spec.Container.Image != "" && dst.Spec.Image != tt.src.Spec.Container.Image {
				t.Errorf("Image not converted correctly: got %v, want %v", dst.Spec.Image, tt.src.Spec.Container.Image)
			}
			if tt.src.Spec.Container.Priority != "" && dst.Spec.Priority != tt.src.Spec.Container.Priority {
				t.Errorf("Priority not converted correctly: got %v, want %v", dst.Spec.Priority, tt.src.Spec.Container.Priority)
			}

			// Verify Job settings if present
			if tt.src.Spec.Job != nil {
				if dst.Spec.Job == nil {
					t.Error("Job settings not converted")
				} else {
					if tt.src.Spec.Job.BackoffLimit != nil && *dst.Spec.Job.BackoffLimit != *tt.src.Spec.Job.BackoffLimit {
						t.Errorf("BackoffLimit not converted correctly: got %v, want %v", *dst.Spec.Job.BackoffLimit, *tt.src.Spec.Job.BackoffLimit)
					}
				}
			}

			// Verify Cron settings if present
			if tt.src.Spec.Cron != nil {
				if dst.Spec.Cron == nil {
					t.Error("Cron settings not converted")
				} else {
					if dst.Spec.Cron.Schedule != tt.src.Spec.Cron.Schedule {
						t.Errorf("Schedule not converted correctly: got %v, want %v", dst.Spec.Cron.Schedule, tt.src.Spec.Cron.Schedule)
					}
				}
			}

			// Verify Command is converted
			if len(tt.src.Spec.Container.Command) > 0 {
				if len(dst.Spec.Command) != len(tt.src.Spec.Container.Command) {
					t.Errorf("Command length not converted correctly: got %v, want %v", len(dst.Spec.Command), len(tt.src.Spec.Container.Command))
				}
			}

			// Verify Env is converted
			if len(tt.src.Spec.Container.Env) > 0 {
				if len(dst.Spec.Env) != len(tt.src.Spec.Container.Env) {
					t.Errorf("Env length not converted correctly: got %v, want %v", len(dst.Spec.Env), len(tt.src.Spec.Container.Env))
				}
			}

			// Verify EnvFrom is converted
			if len(tt.src.Spec.Container.EnvFrom) > 0 {
				if len(dst.Spec.EnvFrom) != len(tt.src.Spec.Container.EnvFrom) {
					t.Errorf("EnvFrom length not converted correctly: got %v, want %v", len(dst.Spec.EnvFrom), len(tt.src.Spec.Container.EnvFrom))
				}
			}

			// Verify FilesFrom is converted
			if len(tt.src.Spec.Container.FilesFrom) > 0 {
				if len(dst.Spec.FilesFrom) != len(tt.src.Spec.Container.FilesFrom) {
					t.Errorf("FilesFrom length not converted correctly: got %v, want %v", len(dst.Spec.FilesFrom), len(tt.src.Spec.Container.FilesFrom))
				}
			}

			// Verify AdditionalPorts is converted
			if len(tt.src.Spec.Container.AdditionalPorts) > 0 {
				if len(dst.Spec.AdditionalPorts) != len(tt.src.Spec.Container.AdditionalPorts) {
					t.Errorf("AdditionalPorts length not converted correctly: got %v, want %v", len(dst.Spec.AdditionalPorts), len(tt.src.Spec.Container.AdditionalPorts))
				}
			}

			// Verify AccessPolicy is converted
			if tt.src.Spec.Container.AccessPolicy != nil {
				if dst.Spec.AccessPolicy == nil {
					t.Error("AccessPolicy not converted")
				}
			}

			// Verify GCP is converted
			if tt.src.Spec.Container.GCP != nil {
				if dst.Spec.GCP == nil {
					t.Error("GCP not converted")
				}
			}

			// Verify Resources is converted
			if tt.src.Spec.Container.Resources != nil {
				if dst.Spec.Resources == nil {
					t.Error("Resources not converted")
				}
			}

			// Verify Probes are converted
			if tt.src.Spec.Container.Liveness != nil {
				if dst.Spec.Liveness == nil {
					t.Error("Liveness probe not converted")
				}
			}
			if tt.src.Spec.Container.Readiness != nil {
				if dst.Spec.Readiness == nil {
					t.Error("Readiness probe not converted")
				}
			}
			if tt.src.Spec.Container.Startup != nil {
				if dst.Spec.Startup == nil {
					t.Error("Startup probe not converted")
				}
			}

			// Verify PodSettings is converted
			if tt.src.Spec.Container.PodSettings != nil {
				if dst.Spec.PodSettings == nil {
					t.Error("PodSettings not converted")
				}
			}

			// Verify RestartPolicy is converted
			if tt.src.Spec.Container.RestartPolicy != nil {
				if dst.Spec.RestartPolicy == nil {
					t.Error("RestartPolicy not converted")
				} else if *dst.Spec.RestartPolicy != *tt.src.Spec.Container.RestartPolicy {
					t.Errorf("RestartPolicy not converted correctly: got %v, want %v", *dst.Spec.RestartPolicy, *tt.src.Spec.Container.RestartPolicy)
				}
			}

			// Verify IstioSettings is converted
			if tt.src.Spec.IstioSettings != nil {
				if dst.Spec.IstioSettings == nil {
					t.Error("IstioSettings not converted")
				}
			}
		})
	}
}

// Helper functions for pointer creation
func int32Ptr(i int32) *int32 {
	return &i
}

func int64Ptr(i int64) *int64 {
	return &i
}

func boolPtr(b bool) *bool {
	return &b
}

func stringPtr(s string) *string {
	return &s
}

func restartPolicyPtr(rp corev1.RestartPolicy) *corev1.RestartPolicy {
	return &rp
}
