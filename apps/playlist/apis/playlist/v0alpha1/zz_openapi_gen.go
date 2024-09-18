//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by grafana-app-sdk. DO NOT EDIT.

package v0alpha1

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/grafana/grafana/apps/playlist/apis/playlist/v0alpha1.Playlist":                    schema_playlist_apis_playlist_v0alpha1_Playlist(ref),
		"github.com/grafana/grafana/apps/playlist/apis/playlist/v0alpha1.PlaylistItem":                schema_playlist_apis_playlist_v0alpha1_PlaylistItem(ref),
		"github.com/grafana/grafana/apps/playlist/apis/playlist/v0alpha1.PlaylistList":                schema_playlist_apis_playlist_v0alpha1_PlaylistList(ref),
		"github.com/grafana/grafana/apps/playlist/apis/playlist/v0alpha1.PlaylistOperatorState":       schema_playlist_apis_playlist_v0alpha1_PlaylistOperatorState(ref),
		"github.com/grafana/grafana/apps/playlist/apis/playlist/v0alpha1.PlaylistSpec":                schema_playlist_apis_playlist_v0alpha1_PlaylistSpec(ref),
		"github.com/grafana/grafana/apps/playlist/apis/playlist/v0alpha1.PlaylistStatus":              schema_playlist_apis_playlist_v0alpha1_PlaylistStatus(ref),
		"github.com/grafana/grafana/apps/playlist/apis/playlist/v0alpha1.PlayliststatusOperatorState": schema_playlist_apis_playlist_v0alpha1_PlayliststatusOperatorState(ref),
	}
}

func schema_playlist_apis_playlist_v0alpha1_Playlist(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/grafana/grafana/apps/playlist/apis/playlist/v0alpha1.PlaylistSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/grafana/grafana/apps/playlist/apis/playlist/v0alpha1.PlaylistStatus"),
						},
					},
				},
				Required: []string{"metadata", "spec", "status"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/apps/playlist/apis/playlist/v0alpha1.PlaylistSpec", "github.com/grafana/grafana/apps/playlist/apis/playlist/v0alpha1.PlaylistStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_playlist_apis_playlist_v0alpha1_PlaylistItem(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PlaylistItem defines model for PlaylistItem.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "type of the item.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"value": {
						SchemaProps: spec.SchemaProps{
							Description: "Value depends on type and describes the playlist item.\n - dashboard_by_id: The value is an internal numerical identifier set by Grafana. This\n is not portable as the numerical identifier is non-deterministic between different instances.\n Will be replaced by dashboard_by_uid in the future. (deprecated)\n - dashboard_by_tag: The value is a tag which is set on any number of dashboards. All\n dashboards behind the tag will be added to the playlist.\n - dashboard_by_uid: The value is the dashboard UID",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"type", "value"},
			},
		},
	}
}

func schema_playlist_apis_playlist_v0alpha1_PlaylistList(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"),
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/apps/playlist/apis/playlist/v0alpha1.Playlist"),
									},
								},
							},
						},
					},
				},
				Required: []string{"metadata", "items"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/apps/playlist/apis/playlist/v0alpha1.Playlist", "k8s.io/apimachinery/pkg/apis/meta/v1.ListMeta"},
	}
}

func schema_playlist_apis_playlist_v0alpha1_PlaylistOperatorState(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PlaylistOperatorState defines model for PlaylistOperatorState.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"descriptiveState": {
						SchemaProps: spec.SchemaProps{
							Description: "descriptiveState is an optional more descriptive state field which has no requirements on format",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"details": {
						SchemaProps: spec.SchemaProps{
							Description: "details contains any extra information that is operator-specific",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"object"},
										Format: "",
									},
								},
							},
						},
					},
					"lastEvaluation": {
						SchemaProps: spec.SchemaProps{
							Description: "lastEvaluation is the ResourceVersion last evaluated",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"state": {
						SchemaProps: spec.SchemaProps{
							Description: "state describes the state of the lastEvaluation. It is limited to three possible states for machine evaluation.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"lastEvaluation", "state"},
			},
		},
	}
}

func schema_playlist_apis_playlist_v0alpha1_PlaylistSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PlaylistSpec defines model for PlaylistSpec.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"interval": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"items": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/apps/playlist/apis/playlist/v0alpha1.PlaylistItem"),
									},
								},
							},
						},
					},
					"title": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
				},
				Required: []string{"interval", "items", "title"},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/apps/playlist/apis/playlist/v0alpha1.PlaylistItem"},
	}
}

func schema_playlist_apis_playlist_v0alpha1_PlaylistStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PlaylistStatus defines model for PlaylistStatus.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"additionalFields": {
						SchemaProps: spec.SchemaProps{
							Description: "additionalFields is reserved for future use",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"object"},
										Format: "",
									},
								},
							},
						},
					},
					"operatorStates": {
						SchemaProps: spec.SchemaProps{
							Description: "operatorStates is a map of operator ID to operator state evaluations. Any operator which consumes this kind SHOULD add its state evaluation information to this field.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("github.com/grafana/grafana/apps/playlist/apis/playlist/v0alpha1.PlayliststatusOperatorState"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/grafana/grafana/apps/playlist/apis/playlist/v0alpha1.PlayliststatusOperatorState"},
	}
}

func schema_playlist_apis_playlist_v0alpha1_PlayliststatusOperatorState(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PlayliststatusOperatorState defines model for Playliststatus.#OperatorState.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"descriptiveState": {
						SchemaProps: spec.SchemaProps{
							Description: "descriptiveState is an optional more descriptive state field which has no requirements on format",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"details": {
						SchemaProps: spec.SchemaProps{
							Description: "details contains any extra information that is operator-specific",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"object"},
										Format: "",
									},
								},
							},
						},
					},
					"lastEvaluation": {
						SchemaProps: spec.SchemaProps{
							Description: "lastEvaluation is the ResourceVersion last evaluated",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"state": {
						SchemaProps: spec.SchemaProps{
							Description: "state describes the state of the lastEvaluation. It is limited to three possible states for machine evaluation.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"lastEvaluation", "state"},
			},
		},
	}
}
