package bot

import (
	"fmt"
	"log"
	"math"
	"time"
)

type RequestBot struct {
	requestBotHelper *requestBotHelper
}

func newRequestBot(opts *RequestBotOptions) (*RequestBot, error) {

	helper, err := newRequestBotHelper(opts)
	if err != nil {
		return nil, err
	}

	return &RequestBot{
		requestBotHelper: helper,
	}, nil
}

func (d RequestBot) BotRequest(reqBody []byte) error {
	maxPage := d.GetRequestMaxPage()

	botRequestCount := 0
	for i := 1; i <= maxPage; i++ {
		endpoints, _, _ := pagination(d.requestBotHelper.endpoints, i, fetchRequestLimit)

		botRequestList, err := botRequest(endpoints, d.requestBotHelper.customHeader, reqBody)
		if err != nil {
			return err
		}

		for _, attack := range botRequestList {
			fmt.Println(string(attack.Endpoint))
			fmt.Println(string(attack.Data))
		}

		botRequestCount += len(botRequestList)
		log.Println("attacked count: ", botRequestCount)
		time.Sleep(waitTime)
	}

	return nil
}

func (d RequestBot) GetRequestMaxPage() int {
	_, _, maxPage := pagination(d.requestBotHelper.endpoints, 1, fetchRequestLimit)
	return maxPage
}

func (d RequestBot) GetRequestMaxCount() int {
	return len(d.requestBotHelper.endpoints)
}

func pagination[T comparable](x []T, page int, perPage int) (data []T, currentPage int, lastPage int) {
	lastPage = int(math.Ceil(float64(len(x)) / float64(perPage)))
	currentPage = page

	if page < 1 {
		page = 1
	} else if lastPage < page {
		page = lastPage
	}

	if page == lastPage {
		data = x[(page-1)*perPage:]
	} else {
		data = x[(page-1)*perPage : page*perPage]
	}

	return
}
