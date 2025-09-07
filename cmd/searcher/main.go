package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"log"
	"bufio"
	"strings"
	"path/filepath"
)

func main() {
	var term *string = flag.String("")
	var path *string = flag.String("p", "./", "-p /path/to/start/search")
	var types *string = flag.String("t", "", "-t comma,seperated,file,extensions e.g. -t .txt,.bin,.docx")
	flag.Parse()
	
	basePath := *path
	filesystem := os.DirFS(basePath)
	
	fs.WalkDir(filesystem, ".", func(path string, d fs.DirEntry, err error) error {
		if(!d.IsDir()){
			file, err := os.Open(basePath+path)
			if(!strings.Contains(*types ,filepath.Ext(path))){
				return err
			}
			if(err != nil){
				log.Fatal(err)
			}
			defer file.Close()
	
			reader := bufio.NewReader(file)
			for {
			    line, err := reader.ReadString('\n')
			    if strings.Contains(line, "hello") {
			        fmt.Println("Found in", path, "->", line)
			    }
			    if err != nil {
			        if err.Error() == "EOF" {
			            break
			        }
			        log.Fatal(err)
			    }
			}
		}
		return err;
	})

}
