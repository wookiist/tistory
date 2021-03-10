# [Blog] Go의 HTTP & REST Client 라이브러리 - Resty

상태: POSTING🖋
작성일시: 2021년 2월 18일 오후 12:38
작성자: Jaewook Oh
최종 편집일시: 2021년 3월 10일 오전 9:59
최종 편집자: Jaewook Oh

# Go의 HTTP & REST Client 라이브러리 - Resty

## API Client

이전 포스팅에서 다뤘던 Echo는 Go의 Web Framework입니다. Echo로 구현한 프로그램은 API Server 등으로 동작할 수 있고, 큰 어려움 없이 Web Server로 사용할 수도 있습니다. 

하지만 프로그램이 API Server가 아닌 API Client로 동작해야 한다면 어떨까요? API Client는 API Server에 Request를 보내는 주체로 Echo와는 정반대의 기능을 수행해야 합니다.

이런 기능을 Go에서 사용할 수 있도록 구현해놓은 패키지가 있습니다. 바로 **Resty** 입니다.

## Resty

Resty는 Ruby의 rest-client에서 영감을 받아 시작된 Go의 HTTP & REST Client 라이브러리 프로젝트입니다. 2015년 9월 15일에 시작해서 2021년 3월 10일 현재, v2.5.0이 가장 최근(2021.02.11)에 릴리즈 됐습니다. Resty는 사용하는 방식도 굉장히 단순합니다.  Resty의 공식 GitHub에 소개된 튜토리얼 코드들로 Resty가 지원하는 다양한 기능들을 알아보겠습니다.

## Install

다른 Go 패키지들처럼 `go get`으로 다운로드하여 설치할 수 있습니다.

```bash
go get github.com/go-resty/resty/v2
```

## Example Code

다음의 코드는 단순 GET을 수행하는 코드입니다.

```go
func main() {
	client := resty.New()
	resp, err := client.R().
			EnableTrace().
			Get("http://example.org")
	fmt.Println(resp)
}
```

이처럼 단순한 작업부터 다양한 Query Parameter, AuthToken 지정, Content-Type 강제, 받아온 Result를 특정 변수에 바인딩하는 작업까지 정말 다양한 기능을 제공하고 있습니다. 만약 위 코드에서 Header를 살펴보고 싶다면 다음을 추가하시면 됩니다.

```go
fmt.Println(resp.StatusCode())
fmt.Println(resp.Status())
fmt.Println(resp.Proto())
fmt.Println(resp.Time())
fmt.Println(resp.ReceivedAt())
```

또 첫번째 코드처럼 `EnableTrace()`를 추가하면, 다음과 같이 Trace Info도 추적해볼 수 있습니다.

```go
ti := resp.Request.TraceInfo()
fmt.Println(ti.DNSLookup)
fmt.Println(ti.ConnTime)
fmt.Println(ti.TCPConnTime)
fmt.Println(ti.TLSHandshake)
fmt.Println(ti.ServerTime)
fmt.Println(ti.ResponseTime)
fmt.Println(ti.TotalTime)
fmt.Println(ti.IsConnReused)
fmt.Println(ti.IsConnWasIdle)
fmt.Println(ti.ConnIdleTime)
fmt.Println(ti.RequestAttempt)
fmt.Println(ti.RemoteAddr.String())
```

## GET

위 코드처럼 단순한 GET 뿐만 아니라, 다양한 기능을 지원하고 있습니다.

### QueryParam

빈번하게 사용되는 Query Parameter도 map 자료를 이용해 간편하게 전달할 수 있습니다. 다음의 예제를 보겠습니다.

```go
resp, err := client.R().
      SetQueryParams(map[string]string{
          "page_no": "1",
          "limit": "20",
          "sort":"name",
          "order": "asc",
          "random":strconv.FormatInt(time.Now().Unix(), 10),
      }).
      SetHeader("Accept", "application/json").
      SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
      Get("/search_result")
```

`SetAuthToken()`은 Token을 전달할 때 사용하게 됩니다만, Token 인증이 필요가 없는 상황이라면 넣을 필요가 없습니다.

### QueryString

위처럼 데이터 하나 하나를 map에 담아 정리하여 보낼 수도 있겠지만, 브라우저 창에서 사용하는 한줄 쓰기로 전달할 수도 있습니다.

```go
resp, err := client.R().
      SetQueryString("productId=232&template=fresh-sample&cat=resty&source=google&kw=buy a lot more").
      SetHeader("Accept", "application/json").
      SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
      Get("/show_product")
```

## POST

POST를 수행할 때도 다양한 조합을 사용해서 작업할 수 있습니다. 다음은 Resty가 지원하는 다양한 Body 형식입니다.

### JSON String

String 자료형을 Body에 담아 전송할 수 있습니다. 주의해야할 점은 Backtick (Backquote)를 사용해서 String을 만든 Raw String이라는 점입니다. `SetBody()`를 보시면, `"username": "testuser"`처럼 JSON 내부에서 이미 double quotation mark를 사용하고 있습니다. 따라서 Backtick을 사용해 내부 자료와 구분이되도록 해주어야 합니다. 이 부분에서 실수하기가 쉽습니다.

```go
resp, err := client.R().
      SetHeader("Content-Type", "application/json").
      SetBody(`{"username":"testuser", "password":"testpass"}`).
      SetResult(&AuthSuccess{}).    // or SetResult(AuthSuccess{}).
      Post("https://myapp.com/login")
```

또한 String 자료형을 지원한다면 `[]byte` 형의 array도 당연히 지원하겠죠?

### `[]byte` array

String 자료형에 이어서 `[]byte` array 자료형도 지원하고 있음을 볼 수 있습니다. String을 지원하기 때문에 어찌보면 당연한 지원이라고 생각할 수도 있을 것 같네요.

```go
resp, err := client.R().
      SetHeader("Content-Type", "application/json").
      SetBody([]byte(`{"username":"testuser", "password":"testpass"}`)).
      SetResult(&AuthSuccess{}).    // or SetResult(AuthSuccess{}).
      Post("https://myapp.com/login")
```

### Struct

Struct 타입을 사용하는 경우, 기본적으로 Content-Type이 `application/json`으로 지정됩니다. 따라서 추가로 세팅해주는 작업을 하지 않아도 됩니다. Struct로 어떤 자료가 보관되어 있다면, 해당 자료는 별도의 변환 작업 없이도 전달이 가능해서 편리합니다. `Marshal`, `Unmarshal`도 사용하지 않습니다.

```go
resp, err := client.R().
      SetBody(User{Username: "testuser", Password: "testpass"}).
      SetResult(&AuthSuccess{}).    // or SetResult(AuthSuccess{}).
      SetError(&AuthError{}).       // or SetError(AuthError{}).
      Post("https://myapp.com/login")
```

### Map

Map 타입을 사용하는 경우, 기본적으로 Content-Type이 `application/json`으로 지정됩니다. 따라서 추가로 세팅해주는 작업을 하지 않아도 됩니다.

```go
resp, err := client.R().
      SetBody(map[string]interface{}{"username": "testuser", "password": "testpass"}).
      SetResult(&AuthSuccess{}).    // or SetResult(AuthSuccess{}).
      SetError(&AuthError{}).       // or SetError(AuthError{}).
      Post("https://myapp.com/login")
```

### Raw Bytes (file upload)

POST는 File Upload 등에 사용할 수 있는 Raw Bytes 전송도 지원하고 있습니다. 다음처럼 파일을 읽어서 `fileBytes` 변수에 저장한다음,  이를 Body에 담아 전송할 수 있습니다. 아래 예제는 Dropbox에 파일을 업로드하는 예제입니다. 한 가지 유의하면 좋을 점은 Go 1.16 릴리즈부터 ioutil 패키지가 deprecated 될 것이기 때문에 `ioutil.ReadFile()`이 아닌 `os.ReadFile()`을 사용하는 것이 바람직하다는 점입니다. (참고 : [Go 1.16 부터 io/ioutil 패키지가 deprecated 됩니다](https://wookiist.dev/89))

```go
// POST of raw bytes for file upload. For example: upload file to Dropbox
fileBytes, _ := ioutil.ReadFile("/Users/jeeva/mydocument.pdf")

// See we are not setting content-type header, since go-resty automatically detects Content-Type for you
resp, err := client.R().
      SetBody(fileBytes).
      SetContentLength(true).          // Dropbox expects this value
      SetAuthToken("<your-auth-token>").
      SetError(&DropboxError{}).       // or SetError(DropboxError{}).
      Post("https://content.dropboxapi.com/1/files_put/auto/resty/mydocument.pdf") // for upload Dropbox supports PUT too
```

## PUT

```go
resp, err := client.R().
      SetBody(Article{
        Title: "go-resty",
        Content: "This is my article content, oh ya!",
        Author: "Jeevanandam M",
        Tags: []string{"article", "sample", "resty"},
      }).
      SetAuthToken("C6A79608-782F-4ED0-A11D-BD82FAD829CD").
      SetError(&Error{}).       // or SetError(Error{}).
      Put("https://myapp.com/article/1234")
```

## DELETE

DELETE 메서드도 지원하고 있습니다. 크게 두 가지 방식을 살펴볼 수 있는데, 첫 번째는 아래 코드처럼 하나의 리소스를 정확하게 URI에 넣어 전송하는 방식입니다.

```go
resp, err := client.R().
      SetAuthToken("C6A79608-782F-4ED0-A11D-BD82FAD829CD").
      SetError(&Error{}).       // or SetError(Error{}).
      Delete("https://myapp.com/articles/1234")
```

다른 하나는 payload/body에 삭제할 id를 JSON string에 담아 전달하는 방식으로 아래처럼 지원하고 있습니다.

```go
resp, err := client.R().
      SetAuthToken("C6A79608-782F-4ED0-A11D-BD82FAD829CD").
      SetError(&Error{}).       // or SetError(Error{}).
      SetHeader("Content-Type", "application/json").
      SetBody(`{article_ids: [1002, 1006, 1007, 87683, 45432] }`).
      Delete("https://myapp.com/articles")
```

## ETC

위에서 소개하지는 않았지만 `PATCH`, `HEAD`, `OPTIONS` 메서드도 지원하고 있습니다. 또한 사용자의 편의를 위한 다양한 기능이 제공되고 있습니다. 몇 가지 살펴보겠습니다.

### Dynamic Path Parameter

Path Parameter를 삽입할 때, 문자열 더하기로 생성을 하는 경우가 많은데, Resty에서는 map 자료형을 이용해 Path Parameter를 정리된 형태로 깔끔하게 전달할 수 있는 기능을 제공하고 있습니다.

```go
client.R().SetPathParams(map[string]string{
   "userId": "sample@sample.com",
   "subAccountId": "100002",
}).
Get("/v1/users/{userId}/{subAccountId}/details")
```

### HTTP Response를 File로 저장하기

HTTP Response를 파일에 저장하는 기능을 제공하고 있습니다. 아래는 `SetOutputDirector()`를 이용해서 기본적으로 찾아갈 디렉터리를 지정하고, `SetOutput()`에서 저장하기를 바라는 위치를 상대 경로로 지정해주면 됩니다.

```go
client.SetOutputDirectory("/Users/test/Downloads")

_, err := client.R().
          SetOutput("plugin/ReplyWithHeader-v5.1-beta.zip").
          Get("http://bit.ly/1LouEKr")
```

또는 아래처럼  절대 경로를 사용해서 파일 위치를 지정해줄 수도 있습니다.

```go
_, err := client.R().
          SetOutput("/MyDownloads/plugin/ReplyWithHeader-v5.1-beta.zip").
          Get("http://bit.ly/1LouEKr")
```

### Form Data 제출

로그인 데이터나 주소 데이터 등을 전달할 때 Form에 채워서 전달하는 경우가 많습니다. 이러한 경우 아래의 `SetFormData()` 를 사용해주시면 됩니다.

```go
resp, err := client.R().
      SetFormData(map[string]string{
        "username": "jeeva",
        "password": "mypass",
      }).
      Post("http://myapp.com/login")
```

또는 아래 코드처럼 특정 변수에, 전달하고자 하는 데이터를 모두 담아서, `SetFormDataFromValues()`를 이용해 변수 하나만 전달하는 방식도 있습니다.

```go
criteria := url.Values{
  "search_criteria": []string{"book", "glass", "pencil"},
}
resp, err := client.R().
      SetFormDataFromValues(criteria).
      Post("http://myapp.com/search")
```

### FilePath를 전달하여 바로 파일 업로드하기

FilePath를 넘겨 주어 파일을 POST 메서드를 호출하는 방식도 있습니다. 다음과 같이 단일 파일을 전달하는 기능을 제공합니다.

```go
// Single file scenario
resp, err := client.R().
      SetFile("profile_img", "/Users/test/test-img.png").
      Post("http://myapp.com/upload")
```

또한 복수 개의 파일을 전달하는 기능도 제공하고 있습니다.

```go
resp, err := client.R().
      SetFiles(map[string]string{
        "profile_img": "/Users/test/test-img.png",
        "notes": "/Users/test/text-file.txt",
      }).
      Post("http://myapp.com/upload")
```

## TLS 사용하지 않기

디버그 모드나 HTTPS를 지원하지 않는 API Server와 통신을 하기 위해선 TLS 설정을 비활성화 해주어야 하는 경우가 종종 있습니다. 메서드를 호출하기 전에 아래의 코드를 먼저 넣어주시면 정상적으로 동작하는 것을 확인하실 수 있습니다.

```go
client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
```

## 마무리

이처럼 API 클라이언트를 테스팅하거나 구현하는데에 필요한 많은 기능들을 편리하게 쓸 수 있게 도와주는 Resty를 기회가 된다면 꼭 써보시길 바랍니다. 본 글에 소개된 기능 외에도 더 많은 기능들이 제공되고 있다는 점도 확인해주시면 감사하겠습니다 🙂

만약 이 글이 도움이 되셨다면 글 좌측 하단의 하트❤를 눌러주시면 감사하겠습니다.

혹시라도 소개된 내용 중에 이상이 있거나, 이해가 가지 않으시는 부분, 또는 추가적으로 궁금하신 내용이 있다면 주저 마시고 댓글💬을 남겨주세요! 빠른 시간 안에 답변을 드리겠습니다 😊

## 참고

- [https://github.com/go-resty/resty](https://github.com/go-resty/resty)