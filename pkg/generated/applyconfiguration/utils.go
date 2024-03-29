// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1 "github.com/sportshead/codegen-demo/pkg/apis/music.sportshead.dev/v1"
	musicsportsheaddevv1 "github.com/sportshead/codegen-demo/pkg/generated/applyconfiguration/music.sportshead.dev/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=music.sportshead.dev, Version=v1
	case v1.SchemeGroupVersion.WithKind("Song"):
		return &musicsportsheaddevv1.SongApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SongSpec"):
		return &musicsportsheaddevv1.SongSpecApplyConfiguration{}

	}
	return nil
}
