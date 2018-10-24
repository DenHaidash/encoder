package file

import "io/ioutil"

func Write(path string, content string) error {
	return ioutil.WriteFile(path, []byte(content), 0644)
}
