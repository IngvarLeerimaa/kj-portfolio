package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func getCode() string {
	return fmt.Sprintf(`
	package main
	
	import (
		"embed"
		"io"
		"log"
		"os"
		"os/exec"
		"path/filepath"
	)
	
	
	//go:embed %v
	var bin1Data embed.FS
	
	
	//go:embed %v
	var bin2Data embed.FS
	
	func main() {
		// Extract and execute bin
		bin1Path := extractFile("%v", bin1Data)
		executeFile(bin1Path)
		defer rmFile(bin1Path)
	
		// Extract and execute bin2
		bin2Path := extractFile("%v", bin2Data)
		executeFile(bin2Path)
		defer rmFile(bin2Path)
	}
	
	func extractFile(fileName string, data embed.FS) string {
		filePath := filepath.Join(fileName)
	
		file, err := data.Open(filePath)
		if err != nil {
			log.Fatal("Failed to open embedded file: ", err)
		}
		defer file.Close()
	
		fileContent, err := io.ReadAll(file)
		if err != nil {
			log.Fatal("Failed to read embedded file: ", err)
		}
	
		err = os.WriteFile(fileName, fileContent, 0o777)
		if err != nil {
			log.Fatal("Failed to extract embedded file: ", err)
		}
	
		return fileName
	}
	
	func executeFile(filePath string) {
		err := os.Chmod(filePath, 0755) // Set execution permission to 755
		if err != nil {
		log.Fatal("Failed to set execution permission: ", err)
		}
		cmd := exec.Command("./" + filePath) // Provide the full path to the extracted file
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	
		err = cmd.Run()
		if err != nil {
			log.Fatal("Failed to execute file: ", err)
		}
	}
	
	func rmFile(filePath string) {
		err := os.Remove(filePath)
		if err != nil {
			log.Fatal("Failed to remove file: ", err)
		}
	}
	`, os.Args[1], os.Args[2], os.Args[1], os.Args[2])
}

func createDirectory(dirPath string) error {
	return os.MkdirAll(dirPath, 0o755)
}

func cpToResult(dirPath string, files ...string) error {
	for _, file := range files {
		if err := copyFile(file, filepath.Join(dirPath, file)); err != nil {
			return err
		}
	}
	return nil
}

func writeGo(dirPath, mainGoCode string) error {
	return os.WriteFile(filepath.Join(dirPath, "main.go"), []byte(mainGoCode), 0o644)
}

func buildGo(goFilePath string) error {
	cmd := exec.Command("go", "build", "-o", "result/result", goFilePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func copyFile(sourceFile, destinationFile string) error {
	sourceData, err := os.ReadFile(sourceFile)
	if err != nil {
		return err
	}
	err = os.WriteFile(destinationFile, sourceData, 0o777)
	if err != nil {
		return err
	}
	return nil
}

func rmFile(files ...string) {
	for _, file := range files {
		os.Remove(file)
	}
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Missing arguments\nUsage: ./injector <target_file> <file_to_be_injected>")
	}
	dirPath := "result/"
	if err := createDirectory(dirPath); err != nil {
		log.Fatal("Failed to create directory: ", err)
	}
	if err := cpToResult(dirPath, os.Args[1], os.Args[2]); err != nil {
		log.Fatal("Failed to copy files: ", err)
	}
	if err := writeGo(dirPath, getCode()); err != nil {
		log.Fatal("Failed to create main.go file: ", err)
	}
	if err := buildGo(filepath.Join(dirPath, "main.go")); err != nil {
		log.Fatal("Failed to compile Go program: ", err)
	}
	rmFile(filepath.Join(dirPath, "main.go"), filepath.Join(dirPath, os.Args[1]), filepath.Join(dirPath, os.Args[2]))
}
