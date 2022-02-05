package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	fp "path/filepath"
)

type PodTracker struct {
	Path string
}

func NewPodTracker(path string) PodTracker {
	return PodTracker{path}
}

// does the store exist?
func (pd PodTracker) StoreExists() bool {
	b, err := exists(pd.Path)
	logif(err)
	return b
}

// create store
func (pd PodTracker) CreateStore() error {
	if pd.StoreExists() {
		return Err(nil, "Store already exists")
	} else {
		os.Mkdir(pd.Path, os.ModePerm)
		return nil
	}
}

// store record
func (pd PodTracker) StoreRecord(edx_anon_id, ipaddr string) error {
	errmsg := fmt.Sprintf("Could not store record for: %s @ %s", edx_anon_id, ipaddr)

	if !pd.StoreExists() {
		err := pd.CreateStore()
		if err != nil {
			return Err(err, "Store did not exist and could not create a new one!")
		}
	}

	filename := fmt.Sprintf("%s%c%s", pd.Path, fp.Separator, edx_anon_id)
	outfile, err := os.Create(filename)
	if err != nil {
		return Err(err, errmsg)
	}

	_, err = outfile.Write([]byte(ipaddr))
	if err != nil {
		return Err(err, errmsg)
	}
	outfile.Close()

	log.Printf("wrote record for: %s @ %s\n", edx_anon_id, ipaddr)
	return nil
}

// check the record
func (pd PodTracker) CheckRecord(edx_anon_id, ipaddr string) (bool, error) {
	if !pd.StoreExists() {
		return false, nil
	}

	filename := fmt.Sprintf("%s%c%s", pd.Path, fp.Separator, edx_anon_id)
	filebytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return false, Err(err, "Could not real record file")
	}

	storedIp := string(filebytes)
	if ipaddr == storedIp {
		return true, nil
	} else {
		log.Printf("check fails for edx_anon_id: %s\n", edx_anon_id)
		log.Printf("stored ip was              : %s\n", storedIp)
		log.Printf("submitted ip was           : %s\n", ipaddr)
		return false, nil
	}
}
