# cloud rules ------------------------------------------------------------------

GO=$(shell which go)

build:
	$(GO) build	

# cloud rules are only relevent in the context of the hosting cloud
# platform, don't use these on your dev machine.

cloud.pull: ## pull the current main version of this repo
	git pull

cloud.apply: ## apply kubernetes config files
	kubectl apply -f deployment.yaml

cloud.test: ## cloud testing
	@echo how to test?

cloud.delete: ## delete the deployment
	kubectl delete deploy submitter -n edxj || true
	kubectl delete service submitter -n edxj

kubectl.get.all: ## kubectl get all -n edxj
	kubectl get all -n edxj

# dev rules ------------------------------------------------------------------

# rules for development

dev.clean: ## clean all the things
	$(GO) clean
	echo implement clean makefile rule

dev.work: ## open all files in editor
	emacs -nw *.go *.yaml README.md *.bash


# dockerize ------------------------------------------------------------------
TAG=v1.8
DOCKER_USER=rhodesd
IMAGE=$(DOCKER_USER)/submitter:$(TAG)

docker.build: ## build the docker image
	docker build -t $(IMAGE) .

# this rule is only for local testing
docker.run: docker.build ## run the docker image
    # docker-env.bash looks like,
    # STAFF_SUBMITTER_USERID="staff"
    # STAFF_SUBMITTER_PASSWD="[this needs to be changed]" 
    # ANSWER_SERVER="www.example.com:5001"
    # ANSWER_SERVER_USERID=""
    # ANSWER_SERVER_PASSWD=""

	docker run -it -p 3000:3000 --expose=3000 --env-file ./docker-env.bash $(IMAGE)

docker.shell: ## get a shell in the container
	docker exec -it $(shell docker ps | grep submitter | cut -d " " -f 1) /bin/sh

docker.push: docker.build ## push the image to dockerhub
	docker push $(IMAGE)

# http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help:
	@grep -E '^[\\.a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk \
	'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

FORCE:

