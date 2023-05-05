package utility

import (
	"fmt"
	"github.com/tebeka/selenium"
)

func ElementClickByTagName(wd *selenium.WebDriver, path *string) {
	ElementEventsToClick(
		FindElementToByTagName(*wd, *path))
}
func ElementSendKeysByTagName(wd *selenium.WebDriver, path, value *string) {
	ElementEventsToSendKeys(
		FindElementToByTagName(*wd, *path),
		*value)
}
func FindElementToByTagName(wd selenium.WebDriver, lementPath string) selenium.WebElement {
	elem, err := wd.FindElement(selenium.ByTagName, lementPath)
	if err != nil {
		panic(err)
	}
	return elem
}

func ElementClickLinkText(wd *selenium.WebDriver, path *string) {
	ElementEventsToClick(
		FindElementToLinkText(*wd, *path))
}
func ElementSendKeysLinkText(wd *selenium.WebDriver, path, value *string) {
	ElementEventsToSendKeys(
		FindElementToLinkText(*wd, *path),
		*value)
}
func FindElementToLinkText(wd selenium.WebDriver, lementPath string) selenium.WebElement {
	elem, err := wd.FindElement(selenium.ByLinkText, lementPath)
	if err != nil {
		panic(err)
	}
	return elem
}
func ElementClickXPATH(wd *selenium.WebDriver, path *string) {
	ElementEventsToClick(
		FindElementToXPATH(*wd, *path))
}
func ElementSendKeysXPATH(wd *selenium.WebDriver, path, value *string) {
	ElementEventsToSendKeys(
		FindElementToXPATH(*wd, *path),
		*value)
}
func FindElementToXPATH(wd selenium.WebDriver, lementPath string) selenium.WebElement {
	elem, err := wd.FindElement(selenium.ByXPATH, lementPath)
	if err != nil {
		panic(err)
	}
	return elem
}
func ElementClickCSS(wd *selenium.WebDriver, path *string) {
	ElementEventsToClick(
		FindElementToByCSSSelector(*wd, *path))
}
func ElementSendKeysCSS(wd *selenium.WebDriver, path, value *string) {
	ElementEventsToSendKeys(
		FindElementToByCSSSelector(*wd, *path),
		*value)
}
func FindElementToByCSSSelector(wd selenium.WebDriver, lementPath string) selenium.WebElement {
	elem, err := wd.FindElement(selenium.ByCSSSelector, lementPath)
	if err != nil {
		panic(err)
	}
	return elem
}

func ElementEventsToClick(elem selenium.WebElement) {
	if err := elem.Click(); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
func ElementEventsToSendKeys(elem selenium.WebElement, inputValue string) {
	if err := elem.SendKeys(inputValue); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
