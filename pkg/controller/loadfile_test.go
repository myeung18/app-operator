package controller

import (
	"fmt"
	"github.com/example-inc/app-operator/pkg/controller/appservice"
	"strings"
	"testing"
)

func TestLoadYaml(t *testing.T) {
	filename := strings.Join([]string{"abc", "def--"}, "/")
	fmt.Println("filename:", filename)
	resMap, err := appservice.LoadWebConsoleYamlSamplesLocal("../../examples", "consoleyamlsamples")
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range resMap {
		fmt.Println(k, " - ", v)
	}
}

