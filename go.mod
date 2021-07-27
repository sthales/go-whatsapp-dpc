module github.com/Rhymen/go-whatsapp

require (
	github.com/golang/protobuf v1.4.1
	github.com/gorilla/websocket v1.4.2
	github.com/pkg/errors v0.9.1
	golang.org/x/crypto v0.0.0-20200323165209-0ec3e9974c59
	google.golang.org/protobuf v1.22.0
)

replace github.com/Rhymen/go-whatsapp => ./

replace github.com/Rhymen/go-whatsapp/examples/echo => ./examples/echo

replace github.com/Rhymen/go-whatsapp/examples/restoreSession => ./examples/restoreSession

replace github.com/Rhymen/go-whatsapp/examples/sendImage => ./examples/sendImage

replace github.com/Rhymen/go-whatsapp/examples/sendTextMessages => ./examples/sendTextMessages

go 1.13
