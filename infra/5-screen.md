# Screen

SSH 로 작업하다보면 여러 디렉토리에서 작업을 동시에 해야하는 상황이 종종 나온다.
그때마다 `cd` 로 옮겨가는건 불편하다.
```
screen
```

`Ctrl + a` + `Tab` 으로 이동한다.

`Ctrl + a` + `c` 로 `bash` 를 실행시킨다.

#### 현재 돌아가고있는 스크린 확인 하기
```sh
$ screen -ls
```
#### 스크린으로 들어가기
```sh
$ screen -R <screen name | screen id>
```
#### 스크린에서 나가기
`Ctrl + a` + `d`