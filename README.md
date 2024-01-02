# codegen-demo
Code for my blog post ["Generating a Kubernetes Go client using a CustomResourceDefinition (CRD)"](https://sportshead.dev/2024/01/02/generating-k8s-go-client-types).

## Example
```shell
$ git clone https://github.com/sportshead/codegen-demo.git
$ cd codegen-demo
$ kubectl apply -f crd.yaml
$ kubectl apply -f example/songs.yaml
$ go run example/main.go
```
