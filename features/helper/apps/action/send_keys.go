package appsaction

import "github.com/golang-automation/features/helper/errors"

// SendKeysByXpath : input text by Xpath selector
func (s *Page) SendKeysByXpath(locator string, text string) error {
	element := s.device().FindByXPath(locator).SendKeys(text)
	errors.LogPanicln(element)

	return element
}

// SendKeysByButton : input text by button
func (s *Page) SendKeysByButton(locator string, text string) error {
	element := s.device().FindByButton(locator).SendKeys(text)
	errors.LogPanicln(element)

	return element
}

// SendKeysByClass : input text by class
func (s *Page) SendKeysByClass(locator string, text string) error {
	element := s.device().FindByClass(locator).SendKeys(text)
	errors.LogPanicln(element)

	return element
}

// SendKeysByID : input text by ID
func (s *Page) SendKeysByID(locator string, text string) error {
	element := s.device().FindByID(locator).SendKeys(text)
	errors.LogPanicln(element)

	return element
}

// SendKeysByLabel : input text by label
func (s *Page) SendKeysByLabel(locator string, text string) error {
	element := s.device().FindByLabel(locator).SendKeys(text)
	errors.LogPanicln(element)

	return element
}

// SendKeysByLink : input text by link
func (s *Page) SendKeysByLink(locator string, text string) error {
	element := s.device().FindByLink(locator).SendKeys(text)
	errors.LogPanicln(element)

	return element
}

// SendKeysByName : input text by class name
func (s *Page) SendKeysByName(locator string, text string) error {
	element := s.device().FindByName(locator).SendKeys(text)
	errors.LogPanicln(element)

	return element
}

// SendKeysByText : input text by text in xpath
func (s *Page) SendKeysByText(locator string, text string) error {
	element := s.device().FindByXPath("//*[contains(@text, '" + locator + "')]").SendKeys(text)
	errors.LogPanicln(element)

	return element
}