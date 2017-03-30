all: lint vet codgen

FILES := $$(find . -name '*.go' | grep -vE 'vendor') 

golint:
	go get github.com/golang/lint/golint  

lint: golint
	@for path in $(SOURCE_PATH); do echo "golint $$path"; golint $$path; done;

clean:
	@rm -rf bin

vet:
	go tool vet $(FILES) 2>&1
	go tool vet --shadow $(FILES) 2>&1

codgen:
	go install ./cmd/codgen.go 


