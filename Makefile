TEST?=$$(go list ./... | grep -v 'vendor')
BINARY=terraform-provider-dynadot
PLUGINS_DIR=~/.terraform.d/plugins/

default: install

build:
	go build -o ${BINARY}

install: build
	mkdir -p ${PLUGINS_DIR}
	mv ${BINARY} ${PLUGINS_DIR}

test: 
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
	xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m
