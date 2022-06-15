export tag=v1
root:
	export ROOT=.

build:
	echo "building server binary"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o .

release:
	echo "building httserverpserver container"
	docker build -t zheng11581/server:${tag} .

push: release
	echo "pushing zheng11581/server"
	docker push zheng11581/server:${tag}