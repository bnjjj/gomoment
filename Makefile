# -----------------------------------------------------------------
#
#        ENV VARIABLE
#
# -----------------------------------------------------------------

# go env vars
GO=$(firstword $(subst :, ,$(GOPATH)))
MAIN_FILE=moment.go
export GO15VENDOREXPERIMENT=1

test:
	go test -v
