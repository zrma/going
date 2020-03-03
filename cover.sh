GO111MODULE=on go test -coverprofile=coverage.out -covermode=count $(go list ./... | grep -v /cmd/)
