package appservice

import (
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/gobuffalo/packr/v2"
	creator "github.com/myeung18/operator-utils/pkg/webconsole/creator"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func LoadWebConsoleYamlSamplesLocal(path string, folder string) (map[string]string, error) {
	return loadfiles(path, folder)
}

func loadfiles(path string, folder string) (map[string]string, error) {
	fullpath := strings.Join([]string{path, folder}, "/")
	box := packr.New("webconsoleyaml", "../../../deploy/examples")
	if !box.HasDir(folder) || box.List() == nil {
		return nil, fmt.Errorf("%s not found ", fullpath)
	}

	resMap := make(map[string]string)
	for _, filename := range box.List() {
		yamlStr, err := box.FindString(filename)
		if err != nil {
			resMap[filename] = err.Error()
		}
		obj := &creator.CustomResourceDefinition{}
		err = yaml.Unmarshal([]byte(yamlStr), obj)
		if err != nil {
			resMap[filename] = err.Error()
		}
		fmt.Println("content ", filename, "File: ", len(yamlStr))
		log.Info("content ", "name", filename, "File: ", len(yamlStr));
	}
	return resMap, nil
}

func loadFilesWithIO(path string, folder string) (map[string]string, error) {
	fullpath := strings.Join([]string{path, folder}, "/")

	fileList, err := ioutil.ReadDir(fullpath)
	if err != nil {
		log.Info("content ", "F:" , err);
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

	fmt.Println("content ", "File: " , len(yamlStr))
	log.Info("content ", "File: " , yamlStr);
}

func listDir(path string, folder string) {
	fullpath := strings.Join([]string{path, folder}, "/")
	fmt.Println("path: ", fullpath, filepath.Dir(fullpath));

	filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error { //annonmous func
			if err != nil {
				return err
			}
			fmt.Println("file: ", path, info.Size())
			return nil
		})
}

