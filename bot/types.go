package bot

type RequestCh struct {
	Endpoint string
	Data     []byte
}

type RequestBotOptions struct {
	*CustomHeader
	requestURL   string
	requestCount int
}

type CustomHeader struct {
	Key   string
	Value string
}
