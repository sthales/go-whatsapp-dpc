module github.com/sthales/go-whatsapp

require (

	github.com/golang/protobuf v1.3.0
	github.com/gorilla/websocket v1.4.0
	github.com/pkg/errors v0.8.1
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
)

replace github.com/sthales/go-whatsapp => ./

replace github.com/sthales/go-whatsapp/examples/echo => ./examples/echo

replace github.com/sthales/go-whatsapp/examples/restoreSession => ./examples/restoreSession

replace github.com/sthales/go-whatsapp/examples/sendImage => ./examples/sendImage

replace github.com/sthales/go-whatsapp/examples/sendTextMessages => ./examples/sendTextMessages
