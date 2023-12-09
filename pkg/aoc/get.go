package aoc

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

const (
	host = "https://adventofcode.com"
	// /2023/day/9"
)

func createDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeToFile(filename, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}

func Get(year, day, part int) error {
	dataPath := fmt.Sprintf("data/%d/%d", year, day)
	err := createDirIfNotExist(dataPath)
	if err != nil {
		log.Println(err)
		return err
	}
	testPath := fmt.Sprintf("data/%d/%d/example_input.txt", year, day)
	_, err = os.Stat(testPath)
	if os.IsNotExist(err) {
		err = writeToFile(testPath, "")
		if err != nil {
			log.Println(err)
			return err
		}
	}
	c := os.Getenv("AOC_SESSION")
	if c == "" {
		return fmt.Errorf("AOC_SESSION env var not set")
	}
	cookie := &http.Cookie{Name: "session", Value: c}
	client := &http.Client{Timeout: time.Second}

	dayUrl := fmt.Sprintf("%s/%d/day/%d", host, year, day)
	dayReq, err := http.NewRequest("GET", dayUrl, nil)
	if err != nil {
		return err
	}
	dayReq.AddCookie(cookie)
	dayResp, err := client.Do(dayReq)
	if err != nil {
		return err
	}
	defer dayResp.Body.Close()

	dayBody, err := io.ReadAll(dayResp.Body)
	if err != nil {
		return err
	}

	dayParsed := ParseDayHtml(string(dayBody), year, day)
	err = writeToFile(fmt.Sprintf("data/%d/%d/prompt.md", year, day), dayParsed)
	if err != nil {
		return err
	}

	inputUrl := fmt.Sprintf("%s/%d/day/%d/input", host, year, day)
	inputRequest, err := http.NewRequest("GET", inputUrl, nil)
	if err != nil {
		return err
	}
	inputRequest.AddCookie(cookie)

	inputResponse, err := client.Do(inputRequest)
	if err != nil {
		return err
	}
	defer inputResponse.Body.Close()

	inputBody, err := io.ReadAll(inputResponse.Body)
	if err != nil {
		return err
	}
	err = writeToFile(fmt.Sprintf("data/%d/%d/input.txt", year, day), string(inputBody))
	if err != nil {
		return err
	}

	return nil
}

func ParseDayHtml(html string, year, day int) string {
	scanner := bufio.NewScanner(strings.NewReader(html))
	var out string
	var main bool
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		if !main && strings.Contains(line, "<main>") {
			main = true
			continue
		}
		if main && strings.Contains(line, "</main>") {
			out += line
			break
		}
		if main {
			out += line + "\n"
		}
	}

	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(out)
	if err != nil {
		log.Fatal(err)
	}
	return markdown
}
