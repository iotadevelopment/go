package tcpprotocol

type IntConsumer = func(number int)

type DataConsumer = func(data []byte)

type ErrorConsumer = func(err error)
