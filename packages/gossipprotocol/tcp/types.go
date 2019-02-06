package tcp

type DataConsumer = func(data []byte)

type ErrorConsumer = func(err error)
