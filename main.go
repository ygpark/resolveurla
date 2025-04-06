package main

import (
	"bufio"
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) >= 1 && (args[0] == "-h" || args[0] == "--help") {
		printHelp()
		os.Exit(0)
	}

	var scanner *bufio.Scanner

	if len(args) >= 1 {
		// 파일에서 입력
		filename := args[0]
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("파일 열기 오류: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		// stdin에서 입력
		fmt.Fprintln(os.Stderr, "파일 입력이 없으므로 표준 입력을 사용합니다. (Ctrl+D로 종료)")
		scanner = bufio.NewScanner(os.Stdin)
	}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		// URL에서 호스트 추출
		host := extractHost(line)
		if host == "" {
			fmt.Printf("호스트 추출 실패: %s\n", line)
			continue
		}

		ips, err := net.LookupIP(host)
		if err != nil {
			fmt.Printf("IP 조회 실패: %s (%v)\n", host, err)
			continue
		}

		for _, ip := range ips {
			fmt.Printf("%s => %s\n", host, ip.String())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("입력 처리 중 오류 발생: %v\n", err)
	}
}

func extractHost(raw string) string {
	if !strings.HasPrefix(raw, "http") {
		raw = "http://" + raw
	}
	u, err := url.Parse(raw)
	if err != nil {
		return ""
	}
	return u.Hostname()
}

func printHelp() {
	fmt.Println(`사용법: resolveurl [파일명]

옵션:
  -h, --help       이 도움말을 표시합니다.

설명:
  이 프로그램은 각 줄에 하나씩 포함된 URL에서 호스트명을 추출하고,
  해당 호스트의 IP 주소를 출력합니다.

  파일명을 생략하면 표준 입력(stdin)에서 URL을 입력받습니다.

예시:
  resolveurl urls.txt
  cat urls.txt | resolveurl
`)
}
