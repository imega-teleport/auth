NAME = auth
PGROUP= imegateleport
IMG = $(PGROUP)/$(NAME)
CWD = /go/src/github.com/$(PGROUP)
LINTER_FLAGS = --fast
TAG = latest

release: build
	@docker login --username $(DOCKER_USER) --password $(DOCKER_PASS)
	@docker push $(IMG):$(TAG)
	@curl -s -X POST -H "TOKEN: $(DEPLOY_TOKEN)" https://d.imega.ru -d '{"namespace":"imega-teleport", "project_name":"$(NAME)", "tag":"$(TAG)"}'

build: unit
	@docker build --build-arg CWD=$(CWD) -t $(IMG):$(TAG) .

unit: pretest
	@docker run --rm -v $(CURDIR):$(CWD) -w $(CWD) golang:1.8-alpine \
		sh -c "go list ./... | grep -v 'vendor\|acceptance' | xargs go test"

pretest:
	@docker run --rm -v $(CURDIR):$(CWD) -w $(CWD) dnephin/gometalinter \
		$(LINTER_FLAGS) --vendor --deadline=600s --disable=gotype --disable=gocyclo \
		--exclude=/usr ./...
