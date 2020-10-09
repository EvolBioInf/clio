all : clio

clio: clio.go
	go build clio.go
clio.go: clio.org
	awk -f scripts/preTangle.awk clio.org | bash scripts/org2nw | notangle -Rclio.go | gofmt > clio.go

.PHONY:	doc
doc:
	make -C doc

clean:
	rm -f *.go
	make clean -C doc
