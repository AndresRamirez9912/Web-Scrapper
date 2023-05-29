package mocks

import (
	"errors"

	"github.com/gocolly/colly"
)

type CollectorMock struct {
	Fail bool
}

func (CollectorMock) OnError(callback colly.ErrorCallback) {}
func (CollectorMock) OnRequest(colly.RequestCallback)      {}
func (CollectorMock) OnResponse(colly.ResponseCallback)    {}
func (CollectorMock) OnHTML(string, colly.HTMLCallback)    {}
func (CollectorMock) Wait()                                {}
func (c CollectorMock) Visit(string) error {
	if c.Fail {
		return errors.New("Error generated for testing")
	}
	return nil
}
