package database
import (
	"log"
	"os"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func errorDirectoryAccess() {
	var err error
	var value string
	var directoryPath string
	fmt.Println("Create a new directory? [Y/n]")
	fmt.Scan(&value)
	if (value=="Y") {
		fmt.Println("Type the desired path to your directory: ")
		fmt.Scan(&directoryPath)
		err = os.MkdirAll(directoryPath, 0755)
		if err != nil {
			log.Fatalf("Error trying to create directory: %v", err)
		}
	}
}

func errorDatabaseAccess() {
	var filePath string
	var err error
	var value string
	fmt.Println("Create a new database file? [Y/n]")
	fmt.Scan(&value)
	if (value=="Y") {
		fmt.Println("Type the path to where your file is going to be created: ")
		fmt.Scan(&filePath)

		filePath = filePath + "/database.db"

		err = os.WriteFile(filePath, nil, 0644)  
		if err != nil {
			log.Fatalf("Error trying to find path %v", err)
		}
	}
}
