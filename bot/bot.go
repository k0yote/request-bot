package bot

type Bot struct {
	*RequestBot
}

func NewBot(requestURL string, requestCount int, customHeader *CustomHeader) (*Bot, error) {
	rbot, err := newRequestBot(&RequestBotOptions{
		CustomHeader: customHeader,
		requestURL:   requestURL,
		requestCount: requestCount,
	})

	if err != nil {
		return nil, err
	}

	return &Bot{
		rbot,
	}, nil
}
