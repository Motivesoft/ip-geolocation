package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	var ipAddress string

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Data is being piped to stdin
		reader := bufio.NewReader(os.Stdin)
		ipAddress, _ = reader.ReadString('\n')

		// if ReadString captured the \n as part of the input, trim it
		ipAddress = strings.TrimSuffix(ipAddress, "\n")
	} else if len(os.Args) > 1 {
		// Use command-line argument
		ipAddress = os.Args[1]
	} else {
		// No piped input and no command-line argument
		fmt.Println("Pass an IP address as command line parameter or as piped input")
		os.Exit(1)
	}

	MakeRequest(ipAddress)
}

func MakeRequest(ipAddress string) {
	env, err := readHeadersFromDotfile(".env")
	if err != nil {
		log.Fatalln(err)
	}

	apiKey := env["api_key"]
	if apiKey == "" {
		log.Fatalln(fmt.Errorf("missing API key in .env"))
	}

	resp, err := http.Get(fmt.Sprintf("https://ipgeolocation.abstractapi.com/v1/?api_key=%s&ip_address=%s", apiKey, ipAddress))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	print(string(body))
}

func print(input string) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(input), "", "    "); err != nil {
		fmt.Println("Error prettifying JSON:", err)
		return
	}
	fmt.Println(prettyJSON.String())
}

func readHeadersFromDotfile(filename string) (map[string]string, error) {
	headers := make(map[string]string)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Ignore empty lines and comment lines (starting with #)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			headers[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return headers, nil
}
