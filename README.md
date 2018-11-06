# ioctl

`import "acln.ro/ioctl"`

Package ioctl provides facilities to define ioctl numbers and perform ioctls
against file descriptors.

See the documentation at https://godoc.org/acln.ro/ioctl.

For the time being, this package is only tested and used on `linux/amd64`,
but it should work on other architectures too.

Package `ioctl` is pure Go, but running tests requires a C compiler and
the appropriate C headers. See `ioctl_test.go` and `ioctl_cgo.go`. To run
the tests, use the `cgotest` tag: `go test -tags cgotest`.

### License

Package ioctl is distributed under the ISC license. A copy of the license
can be found in the LICENSE file.
