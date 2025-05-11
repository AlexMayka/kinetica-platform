package device

type Device interface {
	Connect() error
	Disconnect() error
	SendData(command string) (err error)
}
