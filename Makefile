get-wrk:
	git clone git@github.com:layer5io/wrk2.git

docker-build-base:
	docker image prune -f
	DOCKER_BUILDKIT=1 docker build -t wrk2 -f Dockerfile.wrk2 --no-cache wrk2

docker-run:
	docker rm -f gowrk2 || true
	docker volume prune -f
	docker run --rm --name gowrk2 -it -v `pwd`:/github.com/layer5io/gowrk2 --workdir=/github.com/layer5io/gowrk2 wrk2 go run main.go