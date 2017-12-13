EXECUTABLE_NAME := photo-upload-example

.PHONY: $(EXECUTABLE_NAME) $(EXECUTABLE_NAME)-native
.PHONY: image

REPO_URL=github.com/mockallthethings/$(EXECUTABLE_NAME)
GOLANG_IMAGE=mockallthethings/golang:1.9.2
CONTAINER_GOPATH=/go
CONTAINER_SOURCE_DIR=$(CONTAINER_GOPATH)/src/$(REPO_URL)

IMAGE_NAME=mockallthethings/photo-upload-example

# Default way to build the executable for use inside a docker image
$(EXECUTABLE_NAME):
	docker run \
		-e "CGO_ENABLED=0" \
		-v $(CURDIR):$(CONTAINER_SOURCE_DIR) \
		-w $(CONTAINER_SOURCE_DIR) \
		--rm $(GOLANG_IMAGE) \
		/bin/bash -c "glide install -v && go build -o $(EXECUTABLE_NAME)"

# Use this for a fast build on Mac OS X
$(EXECUTABLE_NAME)-native:
	go build -o photo-upload-example

image:
	docker build -t $(IMAGE_NAME) .
