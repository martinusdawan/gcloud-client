package gcloud

import (
	"io/ioutil"
	"os"
)

type Instance struct {
	Name              string             `json:"name"`
	Zone              string             `json:"zone"`
	NetworkInterfaces []NetworkInterface `json:"networkInterfaces"`
	Status            string             `json:"status"`
}

func (i Instance) IP() string {
	if len(i.NetworkInterfaces) == 0 {
		return ""
	}
	return i.NetworkInterfaces[0].NetworkIP
}

type NetworkInterface struct {
	NetworkIP string `json:"networkIP"`
}

func createTempFile(keys []SSHKey) (*os.File, error) {
	f, err := ioutil.TempFile("", "ssh-key")
	if err != nil {
		return nil, err
	}
	for _, k := range keys {
		f.WriteString(k.String() + "\n")
	}
	f.Sync()
	return f, nil
}