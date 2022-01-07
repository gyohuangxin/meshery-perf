check:
	golangci-lint run

protoc-setup:
	cd meshes
	wget https://raw.githubusercontent.com/layer5io/meshery/master/meshes/meshops.proto

proto:	
	protoc -I meshes/ meshes/meshops.proto --go_out=plugins=grpc:./meshes/

docker:
	DOCKER_BUILDKIT=1 docker build -t gyohuang/meshery-perf .

docker-run:
	(docker rm -f meshery-perf) || true
	docker run --name meshery-perf -d \
	-p 9999:9999\
	-e DEBUG=true \
	gyohuang/meshery-perf:edge-latest

run:
	DEBUG=true GOPROXY=direct GOSUMDB=off go run main.go

run-force-dynamic-reg:
	FORCE_DYNAMIC_REG=true DEBUG=true GOPROXY=direct GOSUMDB=off go run main.go
error:
	go run github.com/layer5io/meshkit/cmd/errorutil -d . analyze -i ./helpers -o ./helpers

test:
	export CURRENTCONTEXT="$(kubectl config current-context)" 
	echo "current-context:" ${CURRENTCONTEXT} 
	export KUBECONFIG="${HOME}/.kube/config"
	echo "environment-kubeconfig:" ${KUBECONFIG}
	GOPROXY=direct GOSUMDB=off GO111MODULE=on go test -v ./...
