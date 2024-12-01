package support

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var notReady string = "Please don't repeatedly request this endpoint before it unlocks! The calendar countdown is synchronized with the server time; the link will be enabled on the calendar the instant this puzzle becomes available.\n"

func getYearDay() (int, int, error) {
	cwd, _ := os.Getwd()
	dayS := filepath.Base(cwd)
	if strings.HasPrefix(dayS, "part") {
		cwd = filepath.Dir(cwd)
		dayS = filepath.Base(cwd)
	}
	yearS := filepath.Base(filepath.Dir(cwd))
	if !(strings.HasPrefix(dayS, "day") && strings.HasPrefix(yearS, "aoc")) {
		return 0, 0, fmt.Errorf("unexpected working dir %s\n", cwd)
	}
	year, err := strconv.Atoi(yearS[len("aoc"):])
	if err != nil {
		return 0, 0, err
	}
	day, err := strconv.Atoi(dayS[len("day"):])
	if err != nil {
		return 0, 0, err
	}
	return year, day, nil
}

func applyCookieHeaders(req *http.Request) {
	_, currentFilePath, _, _ := runtime.Caller(0)
	dir := filepath.Dir(currentFilePath)
	dir = filepath.Dir(dir)
	cookiePath := filepath.Join(dir, ".env")
	contents, err := os.ReadFile(cookiePath)
	if err != nil {
		log.Fatalln(err)
	}
	cookie := strings.TrimSpace(string(contents))
	req.Header.Set("User-Agent", "sqrtxander")
	req.Header.Set("Cookie", cookie)
}

func getInputBytes(year int, day int) ([]byte, error) {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	applyCookieHeaders(req)

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	if string(body) == notReady {
		return nil, fmt.Errorf("%s\n", string(body))
	}

	return body, nil
}

func DownloadInput() {
	year, day, err := getYearDay()
	if err != nil {
		log.Fatalln(err)
	}
	var s []byte
	success := false
	for range 2 {
		s, err = getInputBytes(year, day)
		if err != nil {
			fmt.Println(err)
			time.Sleep(1 * time.Second)
		} else {
			success = true
			break
		}
	}
	if !success {
		log.Fatalln("Timed out after attempting many times")
	}
	path := "input.in"
	cwd, _ := os.Getwd()
	if strings.HasPrefix(filepath.Base(cwd), "part") {
		path = filepath.Join("..", path)
	}
	err = os.WriteFile(path, s, 0400)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(s), "\n")
	if len(lines) > 10 {
		for _, line := range lines[:10] {
			fmt.Println(line)
		}
		fmt.Println("...")
	} else if len(lines[0]) > 80 {
		fmt.Printf("%s...\n", lines[0][:80])
	} else {
		fmt.Println(lines[0])
	}
}

func postAnswer(year int, day int, part int, answer int) (string, error) {
	formData := url.Values{
		"level":  {strconv.Itoa(part)},
		"answer": {strconv.Itoa(answer)},
	}
	formDataStr := formData.Encode()
	requestBody := strings.NewReader(formDataStr)
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", year, day)
	req, err := http.NewRequest(http.MethodPost, url, requestBody)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	applyCookieHeaders(req)

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()
	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func SubmitSolution(part int) {
	year, day, err := getYearDay()
	if err != nil {
		log.Fatalln(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Err() != nil {
		log.Fatalln(err)
	}

	answer, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatalln(err)
	}

	contents, err := postAnswer(year, day, part, answer)
	if err != nil {
		log.Fatalln(err)
	}

	tooQuick, _ := regexp.Compile(`You gave an answer too recently.*to wait.`)
	wrong, _ := regexp.Compile(`That's not the right answer.*?\.`)
	alreadyDone, _ := regexp.Compile(`You don't seem to be solving.*\?`)
	right := "That's the right answer!"

	for _, errorPattern := range []*regexp.Regexp{wrong, tooQuick, alreadyDone} {
		errorMatch := errorPattern.Find([]byte(contents))
		if errorMatch != nil {
			fmt.Printf("\033[41m%s\033[m\n", string(errorMatch))
			return
		}
	}

	if strings.Contains(contents, right) {
		fmt.Printf("\033[42m%s\033[m\n", right)
	} else {
		fmt.Printf("%s\n", contents)
	}
}
