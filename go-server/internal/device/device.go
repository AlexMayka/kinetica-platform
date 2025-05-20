package device

type Device interface {
	SendData(command string) (err error)
	GetData() (data interface{}, err error)
}
