module github.com/sthales/go-whatsapp-dpc

require (
	github.com/Rhymen/go-whatsapp v0.1.0
	github.com/golang/protobuf v1.3.5
	github.com/gorilla/websocket v1.4.2
	github.com/pkg/errors v0.9.1
	golang.org/x/crypto v0.0.0-20200323165209-0ec3e9974c59
)

replace github.com/sthales/go-whatsapp-dpc => ./

replace github.com/sthales/go-whatsapp-dpc/examples/echo => ./examples/echo

replace github.com/sthales/go-whatsapp-dpc/examples/restoreSession => ./examples/restoreSession

replace github.com/sthales/go-whatsapp-dpc/examples/sendImage => ./examples/sendImage

replace github.com/sthales/go-whatsapp-dpc/examples/sendTextMessages => ./examples/sendTextMessages

go 1.13
