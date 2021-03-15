# gitignore를 자동으로 만들어보자

git으로 버전 관리를 하다보면, 커밋하고 싶지 않았던 파일이나 디렉터리가 포함되어 커밋되는 경우가 왕왕 있습니다. macOS의 `.DS_Store`나 Vscode의 `__debug_bin` 파일 등이 대표적입니다.  때론 문서를 작업하다 정상적으로 종료하지 않아 남은 `.swp` 파일도 이런 경우에 속하겠죠.

우리는 이렇게 원치 않는 파일을 관리하기 위해 `.gitignore`파일을 작성합니다.

## `.gitignore`

`.gitignore` 파일은 단순히 파일 확장자나 파일의 이름, 디렉터리 이름으로 구성된 파일입니다. 특정 확장자의 파일을 모두 무시하고 싶다면 와일드카드 `*` 를 사용하여 나타내면 됩니다. 특정 디렉터리 하위에 위치한 파일을 무시하고 싶다면 `dir/` 를 작성해주면 됩니다. 따라서 프로젝트에서 무시하고자 하는 파일이나 디렉터리가 있다면, 미리 추가해주기만 하면 됩니다. 하지만 프로젝트가 커지고 커지다보면 미처 예상치 못한 파일이 나타나는 경우가 종종 있습니다. 또한 매번 `.gitignore`를 작성하는 건 여간 귀찮은 일이 아닐 수 없습니다. 

## gitignore.io

![image](https://user-images.githubusercontent.com/16011260/110716292-3303a600-824a-11eb-8f60-d2eeb43f1435.png)

gitignore.io는 이런 불편함을 해소해주는 소중한 사이트입니다. 현재 프로젝트에서 사용 중인 언어, OS, IDE 환경을 간단하게 검색하여 `Create` 버튼을 누르기만 하면 됩니다.

### 환경에 맞춰 검색을 하고

![image](https://user-images.githubusercontent.com/16011260/110717490-82e36c80-824c-11eb-861e-d05c7b78d713.png)

### Create 버튼을 누르면 끝

![image](https://user-images.githubusercontent.com/16011260/110717379-457edf00-824c-11eb-8db2-d62fa26ee50b.png)


이제 이 파일을 복사, 붙여넣기 하면 됩니다.

참 쉽죠?

## 마무리

만약 이 글이 도움이 되셨다면 글 좌측 하단의 하트❤를 눌러주시면 감사하겠습니다.

혹시라도 소개된 내용 중에 이상이 있거나, 이해가 가지 않으시는 부분, 또는 추가적으로 궁금하신 내용이 있다면 주저 마시고 댓글💬을 남겨주세요! 빠른 시간 안에 답변을 드리겠습니다 😊

## 참고

- [https://www.toptal.com/developers/gitignore](https://www.toptal.com/developers/gitignore)