module greeter_server

go 1.13

replace aplum.com/helloworld => ../helloworld

require (
	aplum.com/helloworld v0.0.0-00010101000000-000000000000 // indirect
	google.golang.org/grpc v1.26.0 // indirect
)
