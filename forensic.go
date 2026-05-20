// ==============================================
// Educational Digital Forensics Toolkit (Go)
// ==============================================
// Features:
// - Deleted File Finder
// - Log Analyzer
// - File Metadata Viewer
//
// Run:
// go run forensic.go
//
// Build:
// go build forensic.go
//
// Educational Use Only
// ==============================================

package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// ==============================================
// FILE HASHER
// ==============================================

func hashFile(path string) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer file.Close()

	hash := md5.New()

	_, err = io.Copy(hash, file)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("\nMD5: %x\n", hash.Sum(nil))
}

// ==============================================
// FILE METADATA
// ==============================================

func fileMetadata(path string) {
	info, err := os.Stat(path)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\n[+] File Metadata")
	fmt.Println("------------------------")
	fmt.Println("Name:", info.Name())
	fmt.Println("Size:", info.Size(), "bytes")
	fmt.Println("Modified:", info.ModTime())
	fmt.Println("Mode:", info.Mode())
}

// ==============================================
// LOG ANALYZER
// ==============================================

func analyzeLogs(path string) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	fmt.Println("\n[+] Suspicious Log Entries")
	fmt.Println("-----------------------------")

	keywords := []string{
		"failed",
		"error",
		"unauthorized",
		"denied",
		"invalid",
		"attack",
	}

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())

		for _, keyword := range keywords {
			if strings.Contains(line, keyword) {
				fmt.Println(scanner.Text())
				break
			}
		}
	}
}

// ==============================================
// RECOVER DELETED-LIKE FILES
// ==============================================

func findTempFiles(directory string) {
	fmt.Println("\n[+] Searching Temporary / Backup Files")
	fmt.Println("----------------------------------------")

	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return nil
		}

		extensions := []string{
			".tmp",
			".bak",
			".old",
			".cache",
		}

		for _, ext := range extensions {
			if strings.HasSuffix(strings.ToLower(info.Name()), ext) {
				fmt.Println(path)
			}
		}

		return nil
	})
}

// ==============================================
// MAIN MENU
// ==============================================

func main() {

	for {
		fmt.Println(`
====================================
 Educational Forensics Toolkit
====================================
1. File Metadata
2. File Hash (MD5)
3. Analyze Logs
4. Find Backup/Temp Files
5. Exit
====================================
`)

		var choice int
		fmt.Print("Select Option: ")
		fmt.Scanln(&choice)

		switch choice {

		case 1:
			var path string

			fmt.Print("Enter File Path: ")
			fmt.Scanln(&path)

			fileMetadata(path)

		case 2:
			var path string

			fmt.Print("Enter File Path: ")
			fmt.Scanln(&path)

			hashFile(path)

		case 3:
			var path string

			fmt.Print("Enter Log File Path: ")
			fmt.Scanln(&path)

			analyzeLogs(path)

		case 4:
			var dir string

			fmt.Print("Enter Directory: ")
			fmt.Scanln(&dir)

			findTempFiles(dir)

		case 5:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid Option")
		}

		time.Sleep(1 * time.Second)
	}
}
