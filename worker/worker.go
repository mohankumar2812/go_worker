package worker

import (
	"sync"

	"github.com/mohankumar2812/go_worker/model"
	"github.com/mohankumar2812/go_worker/utils"
)

func StartWorker(channel chan model.Input, outch chan model.Output,wg *sync.WaitGroup) {
	
	for {
		select {
			case request := <-channel:
				outch <- ConvertRequest(request)
				wg.Done()
		}
	}

}

func ConvertRequest(request model.Input) model.Output {

	converted := model.Output{
		Event:           request.Ev,
		EventType:       request.Et,
		AppID:           request.ID,
		UserID:          request.UID,
		MessageID:       request.MID,
		PageTitle:       request.T,
		PageURL:         request.P,
		BrowserLanguage: request.L,
		ScreenSize:      request.SC,
		Attributes:      getAttribute(request),
		Traits:          getTraits(request),
	}

	return converted
}

func getAttribute(input interface{}) interface{} {
	return utils.GetParams(input, "attributes")
}

func getTraits(input interface{}) interface{} {
	return utils.GetParams(input, "traits")
}
