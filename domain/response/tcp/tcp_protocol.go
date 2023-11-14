package tcp

type TcpResponseProtocol interface {
	BuildSuc() []byte
	BuildFailed(code string) []byte
}
