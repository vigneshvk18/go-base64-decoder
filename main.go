package main

import (
	"archive/zip"
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

func main() {
	// Step 1: Get the Base64 encoded ZIP file string from user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Base64 encoded ZIP file string:")
	base64Zip, _ := reader.ReadString('\n') // Read until newline

	// Trim any extra whitespace or newline characters
	base64Zip = base64Zip[:len(base64Zip)-1]

	// Step 2: Decode the Base64 string
	zipData, err := base64.StdEncoding.DecodeString(base64Zip)
	if err != nil {
		fmt.Println("Error decoding Base64 string:", err)
		return
	}
	fmt.Println("Successfully decoded Base64 string.")

	// Step 3: Create a ZIP file
	zipFilePath := "decoded_file.zip"
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		fmt.Println("Error creating ZIP file:", err)
		return
	}
	defer zipFile.Close()
	fmt.Println("Created ZIP file:", zipFilePath)

	// Step 4: Write the decoded data to the ZIP file
	if _, err := zipFile.Write(zipData); err != nil {
		fmt.Println("Error writing to ZIP file:", err)
		return
	}
	fmt.Println("Wrote data to ZIP file.")

	// Step 5: Open and extract the ZIP file
	if err := extractZipFile(zipFilePath); err != nil {
		fmt.Println("Error extracting ZIP file:", err)
		return
	}
	fmt.Println("ZIP file extracted successfully.")
}

// extractZipFile extracts the contents of a ZIP file
func extractZipFile(zipFilePath string) error {
	// Step 6: Open the ZIP file
	r, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return err
	}
	defer r.Close()
	fmt.Println("Opened ZIP file for extraction.")

	// Step 7: Iterate through the files in the ZIP archive
	for _, f := range r.File {
		fmt.Printf("Extracting: %s\n", f.Name)

		// Create the destination file
		dstFile, err := os.Create(f.Name)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		// Open the ZIP file for reading
		srcFile, err := f.Open()
		if err != nil {
			return err
		}
		defer srcFile.Close()

		// Copy the contents from the ZIP file to the destination file
		if _, err := io.Copy(dstFile, srcFile); err != nil {
			return err
		}
		fmt.Printf("Extracted: %s\n", f.Name)
	}

	return nil
}
