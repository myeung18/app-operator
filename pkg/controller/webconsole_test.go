package controller

import (
	"fmt"
	"github.com/myeung18/operator-utils/pkg/webconsole"
	"testing"
)

func TestLoadYaml(t *testing.T) {
	resMap, err := webconsole.LoadWebConsoleYamlSamples("../../examples", "consoleyamlsamples")
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range resMap {
		fmt.Println(k, " - ", v)
	}
}

func TestWebConsoleEnrichment(t *testing.T) {
	resMap, err := webconsole.LoadWebConsoleEnrichment("./examples", "console")
	if err != nil {
		fmt.Println(err)
	}
	for k, v := range resMap {
		fmt.Println(k, " - ", v)
	}
}
