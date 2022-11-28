package main

import(
	"fmt"
	"os"
	"time"
)

func main() {

	file, err := os.OpenFile("text.txt", os.Q_APPEND| os.O_CREATE| os.WRONLY, 0644)

	if err!=  nil{
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for{
		name, err := os.Hostname()
		if err !=  nil {
			panic(err)
		}
		dataWriter.WriteString(time.Now(), name + "\n")
		time.Sleep(60 * time.Second)
	}
	dataWriter.Flush()
	file.Close()
}




