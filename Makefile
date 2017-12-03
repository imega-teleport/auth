NAME = auth
PGROUP= imega-teleport
DGROUP= imegateleport
IMG = $(DGROUP)/$(NAME)
CWD = /go/src/github.com/$(PGROUP)/$(NAME)
LINTER_FLAGS = --fast
TAG = latest

release: build acceptance
	@docker login --username $(DOCKER_USER) --password $(DOCKER_PASS)
	@docker push $(IMG):$(TAG)
	@curl -s -X POST -H "TOKEN: $(DEPLOY_TOKEN)" https://d.imega.ru -d '{"namespace":"$(PGROUP)", "project_name":"$(NAME)", "tag":"$(TAG)"}'

test: clean build acceptance

build: unit
	@docker build --build-arg CWD=$(CWD) -t $(IMG):$(TAG) .

unit: pretest
	@docker run --rm -v $(CURDIR):$(CWD) -w $(CWD) golang:1.8-alpine \
		sh -c "go list ./... | grep -v 'vendor\|acceptance' | xargs go test"

pretest:
	@docker run --rm -v $(CURDIR):$(CWD) -w $(CWD) dnephin/gometalinter \
		$(LINTER_FLAGS) --vendor --deadline=600s --disable=gotype --disable=gocyclo --disable=gas \
		--exclude=/usr --exclude='api' ./...

acceptance:
	@touch $(CURDIR)/mysql.log
	@TAG=$(TAG) IMG=$(IMG) docker-compose up -d
	@docker run --rm \
		-v $(CURDIR):$(CWD) \
		-w $(CWD) \
		-e TELEPORTDB_HOST=testdb \
		-e TELEPORTDB_PORT=3306 \
		-e TELEPORTDB_USER=root \
		-e TELEPORTDB_PASS=qwerty \
		-e TELEPORTDB_NAME=teleport \
		--network auth_default \
		golang:1.8-alpine sh -c "go test -v -tags=acceptance github.com/imega-teleport/auth/acceptance"

clean:
	@-rm $(CURDIR)/mysql.log
	@TAG=$(TAG) IMG=$(IMG) docker-compose rm -sfv
	@docker images --quiet --filter=dangling=true | xargs docker rmi

.PHONY: acceptance
