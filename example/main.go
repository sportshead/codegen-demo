package main

import (
	"context"
	"flag"
	"fmt"
	clientset "github.com/sportshead/codegen-demo/pkg/generated/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"strings"
)

func main() {
	var namespace = flag.String("namespace", "", "namespace to search for songs in")
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	musicClient, err := clientset.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	songs, err := musicClient.CoolMusicV1().Songs(*namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Found %d songs in namespace '%s'\n", len(songs.Items), *namespace)

	count := 0
	sum := 0

	for _, song := range songs.Items {
		spec := song.Spec
		rating := ""
		if spec.Rating != 0 {
			count++
			sum += spec.Rating

			rating = fmt.Sprintf("\nRating: %d/5", spec.Rating)
		}
		genres := ""
		if len(spec.Genres) > 0 {
			genres = "\nGenres: " + strings.Join(spec.Genres, ", ")
		}
		fmt.Printf("\n%s by %s%s%s\n", spec.Title, spec.Artist, rating, genres)
	}
	if count > 0 {
		fmt.Printf("\nThere are %d songs with ratings. Average rating: %d/5", count, sum/count)
	}
}
