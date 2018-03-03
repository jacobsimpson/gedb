
# I tried golint but it doesn't recognize when it's suggestions can't actually
# compile. I checked the doc page and it says there is no knobs to suppress
# warnings. False alarms make it valueless.
build::
	goimports -w `find . -name "*.go"`
	go test github.com/jacobsimpson/gedb
	go vet
	go build -o gedbcmd ./cmd/gedbcmd

clean::
	rm -Rf gedbcmd

