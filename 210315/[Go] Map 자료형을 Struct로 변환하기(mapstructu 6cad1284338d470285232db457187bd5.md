# Map

Python에는 Dictionary라는 강력한 자료형이 있습니다. 굉장히 유연하고, 편리한 자료형입니다. 기본적으로 Python의 Dictionary는 다른 언어의 Map 자료형과 동일합니다. Python Dictionary의 특별한 점이라면, 하나의 Dictionary 변수 안에 각기 다른 자료형의 데이터가 보관될 수 있다는 점입니다. 따라서 Python에서 어떤 데이터를 Dictionary로 보관하는 일은 그리 어려운 일이 아닙니다. 특히, JSON 데이터를 별다른 추가 작업 없이 Dictionary에 대응할 수 있다는 점은 정말 엄청난 기능이라고 생각합니다.

그러나 Go 진영에선 상황이 다릅니다. Go에서 특정 데이터를 Dictionary로 보관하기 위해선, Key-Value가 어떤 자료형인지 명시해주어야 합니다. 그리고 해당 자료형이 아니면, Map에 보관할 수 없습니다(물론, `Value`를 `interface{}` 타입으로 지정해주면 가능하긴 합니다).또한 Map 자료형 데이터를 구조체로 변환하는 작업도 꽤나 까다롭습니다. 마치 JSON 데이터를 구조체로 보관하기 위해 추가 작업을 수행해주어야 했던 것과 비슷합니다.

# map[string]interface{}

기본적으로 JSON 데이터를 `interface{}` 타입으로 받게 되면 `map[string]interface{}` 자료형에 보관됩니다. 하지만 이렇게 `map` 으로 저장된 데이터를 구조체로 변환하는 방법은 기본적으론 없습니다. 따라서 직접 코드를 구현하거나, 먼저 경험하신 분들이 만들어놓으신 패키지를 활용하여야 합니다.

이런 일을 겪게 된 건 여러 타입의 JSON을 하나의 공통 메서드로 처리할 때였습니다. JSON으로 받은 데이터를 구조체로 변환해서 파라미터로 넘겨 주는 방식을 사용하고 싶었지만, 각기 다른 JSON을 처리할 필요가 있어서 최대한 Generic한 메서드를 만들고자 했습니다. 따라서 파라미터의 자료형을 `interface{}` 타입으로 지정하였습니다. Go의 인터페이스는 C의 `void`와 유사한 역할도 수행하기 때문입니다. 이렇게 되니 JSON 데이터는 사라지고, `map[string]interface{}` 타입만 남았습니다. 여기서 문제가 발생했습니다. `map` 자료형을 `struct`로 변환할 방법을 찾아야 했습니다.

# mapstructure

이러한 배경에서 찾게 된 라이브러리가 바로 mapstructure입니다. 

mapstructure는 Generic Map 자료형을 구조체로 변환하거나, 또는 그 반대로 변환할 때 사용할 수 있는 기능을 제공하며, 이에 수반되는 꽤나 유용한 에러 핸들링 기능을 제공합니다. 이 라이브러리는 미지의 데이터 구조를 갖는 JSON이나 Gob 데이터를 디코딩할 때 유용합니다. 라이브러리의 설명은 이렇습니다.

> Go는 JSON과 같은 데이터를 디코딩하는 훌륭한 표준 라이브러리를 이미 제공하고 있습니다. 가장 표준적인 방법은 구조체를 먼저 생성하고, 인코딩된 형식의 바이트로부터 구조체를 채우는 것입니다. 그러나 만약 특정 필드의 값에 따라 인코딩이 변해야 하는 경우가 있을 수 있습니다.
이러한 경우에는 먼저 JSON 데이터를 `map[string]interface{}`구조로 디코딩하고, 변동이 가능한 필드를 먼저 읽은 다음, 본 라이브러리와 같은 도구를 활용하여 적절한 구조체로 디코딩하는 것이 훨씬 간단합니다.

README에서 예를 든 형식은 이렇습니다.

```json
{
	"type": "person",
	"name": "Mitchell"
}
```

만약 "type" 필드에 "person"이 아닌 "animal" 또는 "object" 등이 왔다면, 해당 다음 필드는 "name"이 아닌, "kind" 등이 될 수도 있을 것입니다. 즉, Go의 표준 방법으로는 유동적인 처리가 어려움을 알 수 있습니다.

## Installation

설치하기 위해선 `go get` 커맨드를 활용해주시면 됩니다. 그리고 필요한 파일에 import 해주세요.

```bash
$ go get github.com/mitchellh/mapstructure
```

## Usage

사용 방법은 굉장히 단순합니다.  `map` 자료가 어떤 구조체로 변환될지를 결정해주시면 끝입니다. 그리고 목표가 되는 구조체에는 `json` 주석처럼 `mapstructure` 주석을 달아주시면 됩니다. 다음의 코드를 참고해주세요.

```go
package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type MyStruct struct {
	Name string `mapstructure:"name"`
	Age  int64  `mapstructure:"age"`
}

func main() {
	myData := make(map[string]interface{})
	myData["Name"] = "Wookiist"
	myData["Age"] = int64(27)

	result := &MyStruct{}
	if err := mapstructure.Decode(myData, &result); err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
```

[Playground에서 위 코드 실행해보기](https://play.golang.org/p/OMAUQkwk-lg)

# 마무리

만약 이 글이 도움이 되셨다면 글 좌측 하단의 하트❤를 눌러주시면 감사하겠습니다.

혹시라도 소개된 내용 중에 이상이 있거나, 이해가 가지 않으시는 부분, 또는 추가적으로 궁금하신 내용이 있다면 주저 마시고 댓글💬을 남겨주세요! 빠른 시간 안에 답변을 드리겠습니다 😊

# 참고

- [https://github.com/mitchellh/mapstructure](https://github.com/mitchellh/mapstructure)
- [https://stackoverflow.com/questions/26744873/converting-map-to-struct](https://stackoverflow.com/questions/26744873/converting-map-to-struct)