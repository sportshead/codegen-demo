// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/sportshead/codegen-demo/pkg/apis/music.sportshead.dev/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SongLister helps list Songs.
// All objects returned here must be treated as read-only.
type SongLister interface {
	// List lists all Songs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Song, err error)
	// Songs returns an object that can list and get Songs.
	Songs(namespace string) SongNamespaceLister
	SongListerExpansion
}

// songLister implements the SongLister interface.
type songLister struct {
	indexer cache.Indexer
}

// NewSongLister returns a new SongLister.
func NewSongLister(indexer cache.Indexer) SongLister {
	return &songLister{indexer: indexer}
}

// List lists all Songs in the indexer.
func (s *songLister) List(selector labels.Selector) (ret []*v1.Song, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Song))
	})
	return ret, err
}

// Songs returns an object that can list and get Songs.
func (s *songLister) Songs(namespace string) SongNamespaceLister {
	return songNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SongNamespaceLister helps list and get Songs.
// All objects returned here must be treated as read-only.
type SongNamespaceLister interface {
	// List lists all Songs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Song, err error)
	// Get retrieves the Song from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Song, error)
	SongNamespaceListerExpansion
}

// songNamespaceLister implements the SongNamespaceLister
// interface.
type songNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Songs in the indexer for a given namespace.
func (s songNamespaceLister) List(selector labels.Selector) (ret []*v1.Song, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Song))
	})
	return ret, err
}

// Get retrieves the Song from the indexer for a given namespace and name.
func (s songNamespaceLister) Get(name string) (*v1.Song, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("song"), name)
	}
	return obj.(*v1.Song), nil
}
