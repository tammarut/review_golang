package search

import (
	"fmt"
	"sync"
)

func searchFromFolder(keyword, folder string, wg *sync.WaitGroup) {
	fmt.Println(folder, "Now searching..")

	wg.Done()
}

func Run(keyword string) {
	var wg sync.WaitGroup
	folders := [3]string{"Document", "Image", "Library"}
	wg.Add(len(folders))
	for _, folder := range folders {
		go searchFromFolder(keyword, folder, &wg)
	}

	wg.Wait()
}
