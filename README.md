# Educational Digital Forensics Toolkit (Go)

A beginner-friendly digital forensics toolkit written in Go.

This project demonstrates basic forensic investigation concepts such as:

- File metadata analysis
- File hashing
- Log analysis
- Temporary/backup file discovery

⚠️ Educational use only.

---

# Features

## File Metadata Viewer
Displays:

- File name
- File size
- Modification time
- File permissions

---

## File Hashing
Generates MD5 hashes for integrity verification.

---

## Log Analyzer
Searches log files for suspicious keywords such as:

- failed
- error
- unauthorized
- denied
- invalid
- attack

---

## Temp/Backup File Discovery
Finds files such as:

```text
.tmp
.bak
.old
.cache
```

---

# Requirements

Install Go:

https://go.dev/dl/

Verify installation:

```bash
go version
```

---

# Project Structure

```text
Educational-Forensics-Toolkit/
│
├── forensic.go
├── README.md
```

---

# How To Run

## Run Directly

```bash
go run forensic.go
```

---

## Build Executable

```bash
go build forensic.go
```

---

## Run Built Executable

### Linux / macOS

```bash
./forensic
```

### Windows

```bash
forensic.exe
```

---

# Menu

```text
====================================
 Educational Forensics Toolkit
====================================
1. File Metadata
2. File Hash (MD5)
3. Analyze Logs
4. Find Backup/Temp Files
5. Exit
====================================
```

---

# Example Usage

## File Metadata

Input:

```text
/path/to/file.txt
```

Output:

```text
Name: file.txt
Size: 1024 bytes
Modified: 2026-05-20
```

---

## File Hash

Output Example:

```text
MD5: 5d41402abc4b2a76b9719d911017c592
```

---

## Log Analysis

Suspicious log entries:

```text
Failed login attempt
Unauthorized access detected
```

---

# Learning Goals

This project helps beginners understand:

- Digital forensics basics
- File integrity checking
- Log investigation
- Incident response concepts
- File system analysis

---

# Recommended Tools To Learn

- Autopsy
- FTK
- Volatility
- Sleuth Kit
- Wireshark

---

# Disclaimer

This project is intended for:

- Education
- Lab environments
- Authorized investigations

Do NOT use for unauthorized access or illegal activities.

---

# License

MIT License
