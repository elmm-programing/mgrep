package main


import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func searchFileRecursive(directory string, files *[]string,wg *sync.WaitGroup) {
  dir,err := os.ReadDir(directory)
  if err != nil {
    log.Fatal(err)
  }
  defer wg.Done()
for _, v := range dir {
		if v.IsDir() {
      wg.Add(1)
		 go searchFileRecursive(fmt.Sprintf("%v%v/", directory,v.Name()), files,wg)
		} else {
			*files = append(*files, directory + v.Name())  
		}
		}
}

func readFile(path string,substring string,wg *sync.WaitGroup)  {
  defer wg.Done()
  file,err := os.Open(path)
  if err != nil {
    log.Fatal(err)
  }
  scanner := bufio.NewScanner(file)
  var i int
  for scanner.Scan(){
    if strings.Contains(scanner.Text(), substring) {
      fmt.Println("Path:",path)
    fmt.Println(i,scanner.Text())
      
    }
    i++
    
  }
  
}

func main() {
	subString := os.Args[1]
	directory := os.Args[2]
	files, err := os.ReadDir(directory)
  var wg sync.WaitGroup
	if err != nil {
		fmt.Println("Directory not found")
	}
	 var allFiles []string
	for _, v := range files {
		if v.IsDir() {
      wg.Add(1)
			go searchFileRecursive(fmt.Sprintf("%v%v/", directory,v.Name()), &allFiles,&wg)
		} else {
			allFiles = append(allFiles, directory + v.Name())  
		}

		}
  	
  wg.Wait()
  for _, v := range allFiles {
    wg.Add(1)
    go readFile(v,subString,&wg)
  }
  wg.Wait()

}
