package is

import "os"

func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else if err != nil {
		return false
	}
	return true
}

func PathNotExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	} else if err != nil {
		return true
	}
	return false
}

func PathFile(path string) bool {
	if info, err := os.Stat(path); err != nil {
		return false
	} else if info.IsDir() {
		return false
	}
	return true
}

func PathDir(path string) bool {
	if info, err := os.Stat(path); err != nil {
		return false
	} else if !info.IsDir() {
		return false
	}
	return true
}
