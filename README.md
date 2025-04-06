# resolveurl

`resolveurl`은 텍스트 파일 또는 표준 입력으로부터 URL 목록을 받아 각 URL의 호스트 이름에 대한 IP 주소를 조회해 출력하는 간단한 CLI 도구입니다.

## 📦 설치

Go 언어가 설치되어 있다면 아래 명령어 중 하나로 설치할 수 있습니다:

### ✅ go install 사용
```bash
go install github.com/ygpark/resolveurl@latest
```

### ✅ 직접 빌드
```bash
go build -o resolveurl resolveurl.go
```

## 🚀 사용법

### 파일로부터 URL 목록 처리

```bash
./resolveurl urls.txt
```

### 표준 입력(stdin) 사용

```bash
cat urls.txt | ./resolveurl
```
또는 직접 입력:
```bash
./resolveurl
google.com
naver.com
https://example.com
(입력 후 Ctrl+D 로 종료)
```

## 🧩 옵션

| 옵션        | 설명                 |
|-------------|----------------------|
| `-h`, `--help` | 사용법을 출력합니다. |

## 📄 입력 파일 형식

- 텍스트 파일 내에 한 줄에 하나씩 URL을 입력합니다.
- URL 형식이 아니더라도 도메인 형태면 자동으로 `http://`를 붙여 처리합니다.

예시:
```
https://example.com
naver.com
http://google.com
```

## 🖨️ 출력 예시
```
example.com => ***.***.***.***
naver.com => ***.***.***.***
google.com => ***.***.***.***
```
※ 실제 IP 주소는 예시에서 익명 처리되었습니다.

## 📌 참고 사항

- 내부적으로 `net.LookupIP()` 함수를 사용하여 도메인에 대한 IP 주소를 조회합니다.
- 여러 개의 IP 주소가 있을 경우 모두 출력됩니다.

---
