package html_link_parser

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseFirst(t *testing.T) {
		r := strings.NewReader(firstLink)
		links, err := Parse(r)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", links)
}

func TestParseSecond(t *testing.T) {
	r := strings.NewReader(secondLink)
	links, err := Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}

func TestParseThird(t *testing.T) {
	r := strings.NewReader(thirdLink)
	links, err := Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}

func TestParseFourth(t *testing.T) {
	r := strings.NewReader(fourthLink)
	links, err := Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
