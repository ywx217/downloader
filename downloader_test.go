package downloader

import (
	"fmt"
	"log"
	"sync"
	"testing"
)

func Test_NewFile(t *testing.T) {
	fileDl, err := NewFileDl("", "http://packages.linuxdeepin.com/ubuntu/dists/devel/main/binary-amd64/Packages.bz2", -1, "/tmp", "")
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v\n", fileDl)

	fileDl.OnStart(func(id string) {
		fmt.Println(GetDownloader(id).File.Name, "download started")
	})

	fileDl.Start()

	var wg sync.WaitGroup
	wg.Add(1)
	fileDl.OnFinish(func(id string) {
		fmt.Println(GetDownloader(id).File.Name, "download finished")
		wg.Done()
	})
	wg.Wait()
}
