package controller

import (
	"fmt"
	"github.com/myeung18/operator-utils/pkg/webconsole"
	"strings"
	"testing"
)

func TestLoadYaml(t *testing.T) {
	filename := strings.Join([]string{"abc", "def--"}, "/")
	fmt.Println("filename:", filename)
	resMap, err := webconsole.LoadWebConsoleYamlSamples("xxxxxxxxx", "../../examples", "consoleyamlsamples")
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range resMap {
		fmt.Println(k, " - ", v)
	}
}

