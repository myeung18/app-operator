package controller

import (
	"fmt"
	"strings"
	"testing"
)

func TestLoadYaml(t *testing.T) {
	filename := strings.Join([]string{"abc", "def--"}, "/")
	fmt.Println("filename:", filename)

	//files, err := appservice.LoadFilesOnlyWithBox("consoleyamlsamples", "../../deploy/examples", "consoleyamlsamples")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//resMap, err := webconsole.ApplyMultipleWebConsoleYamls(files)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for k, v := range resMap {
	//	fmt.Println(k, " - ", v)
	//}
}

