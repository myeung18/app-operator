package appservice

import (
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/gobuffalo/packr/v2"
	//"github.com/myeung18/operator-utils/pkg/webconsole"
	//creator "github.com/myeung18/operator-utils/pkg/webconsole/creator"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func LoadWebConsoleYamlSamplesLocal(path string, folder string) (map[string]string, error) {
	return loadfiles(path, folder)
}

func LoadFilesOnlyWithBox(boxname string, path string, folder string) ([]string, error) {
	fullpath := strings.Join([]string{path, folder}, "/")
	box := packr.New(boxname, fullpath)
	if  box.List() == nil {
		return nil, fmt.Errorf("%s not found ", fullpath)
	}

	var files []string
	for _, filename := range box.List() {
		yamlStr, err := box.FindString(filename)
		if err != nil {
			fmt.Println(err)
			continue
		}
		files = append(files, yamlStr)
	}
	return files, nil
}

func loadfiles(path string, folder string) (map[string]string, error) {
	fullpath := strings.Join([]string{path, folder}, "/")
	box := packr.New("consoleyamlsamples", "../../../deploy/examples/consoleyamlsamples")
	if !box.HasDir(folder) || box.List() == nil {
		return nil, fmt.Errorf("%s not found ", fullpath)
	}

	box = packr.New("webconsole", "../../../deploy/examples/webconsole")
	if !box.HasDir(folder) || box.List() == nil {
		return nil, fmt.Errorf("%s not found ", fullpath)
	}

	resMap := make(map[string]string)
	for _, filename := range box.List() {
		yamlStr, err := box.FindString(filename)
		if err != nil {
			resMap[filename] = err.Error()
		}
		//obj := nil //&creator.CustomResourceDefinition{}
		err = yaml.Unmarshal([]byte(yamlStr), nil)
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

func loadTestFiles(path string, folder string) []string {
	fullpath := strings.Join([]string{path, folder}, "/")

	fileList, err := ioutil.ReadDir(fullpath)
	if err != nil {
		fmt.Println( fmt.Errorf("%s not found with io ", fullpath))
	}

	var files []string
	for _, filename := range fileList {
		yamlStr, err := ioutil.ReadFile(fullpath + "/" + filename.Name())
		if err != nil {
			fmt.Println(err)
			continue
		}
		files = append(files, string(yamlStr))
	}
	//for _, f := range files {
	//	a := []rune(f)
	//	fmt.Println("filename: ", string(a[0: 10]))
	//
	//	err := ApplyWebConsoleYaml(f)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//}
	//
	//resMap, err := webconsole.ApplyMultipleWebConsoleYamls(files)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for k, v := range resMap {
	//	//fmt.Println(files[k])
	//	fmt.Println(k, " - ", v)
	//}
	return files
}
