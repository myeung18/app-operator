package controller

import (
	"fmt"
	"io/ioutil"
	"strings"

	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var log = logf.Log.WithName("controller_appservice")

func LoadWebConsoleYamlSamples_local(path string, folder string) (map[string]string, error) {
	return loadFilesWithIO(path, folder)
}

func loadFilesWithIO(path string, folder string) (map[string]string, error) {
	fullpath := strings.Join([]string{path, folder}, "/")
	fileList, err := ioutil.ReadDir(fullpath)
	if err != nil {
		return nil, fmt.Errorf("%s not found with io ", fullpath)
	}

	resMap := make(map[string]string)
	for _, filename := range fileList {
		process(fullpath, filename.Name(), resMap)
	}
	return resMap, nil
}

func process(fullpath string, filename string, resMap map[string]string) {
	yamlStr, err := ioutil.ReadFile(fullpath + "/" + filename)
	if err != nil {
		resMap[filename] = err.Error()
	}

	fmt.Println("content ", "File: " , string(yamlStr))
	log.Info("content ", "File: " , yamlStr);
}
