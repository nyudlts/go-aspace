package lib

import "fmt"

func(f FileVersion) toString() string {
	s := fmt.Sprintf("File URI: %s\nPublish %v\n", f.FileURI, f.Publish)
    return s;
}