# Docker

### 도커 설치

```shell
# 패키지 업데이트
$ yum -y update

# 도커 레지스트리와 도커를 설치
$ yum -y install docker docker-registry
```
### 도커 이미지 받아오기

```shell
$ docker pull <image_name>
```
### 도커 이미지 빌드하기

`-t` : 생성될 이미지의 이름을 지정할 수 있다.

`-f` 도커파일의 이름이 `Dockerfile` 이 아닌경우 따로 설정된 이름을 적어주면된다.

```shell
$ docker build -t <image name> <도커파일이 있는 경로>
```
### 도커 이미지 목록 확인
```shell
$ docker images

REPOSITORY TAG IMAGE ID CREATED SIZE

hivv/carrot latest c92f26e27dd1 35 hours ago 1.02 GB

docker.io/node 12 29be39bd6917 7 days ago 918 MB

docker.io/hello-world latest d1165f221234 6 weeks ago 13.3 kB
```
### 도커 이미지 삭제하기
```shell
$ docker rmi [IMAGE_ID | REPOSITORY]
```
### NodeJS 도커 파일 예시
```dockerfile
FROM node:12
# 앱 디렉터리 생성
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm install
COPY . .
# 8080 포트로 실행
EXPOSE 8080
CMD ["npm", "start"]
```
### 도커 이미지 실행

처음에 좀 헷갈렸던 부분인데 도커 안쪽에서 `8080` 포트로 실행이 되는거고 외부적으로 저 컨테이너는 `49160` 포트에서 돌아가는 것이다. 즉 바깥쪽에서 `Nignx` 설정같은 것을 하게되면 반드시 `49160` 외부 포트와 연결해주어야한다.
```shell
$ docker run -p 49160:8080 -d <your username>/<container name>
```
### 컨테이너 로그 확인
```shell
$ docker logs <container id>
```
### 컨테이너에 직접 접속
```shell
$ docker exec -it <container id> /bin/bash
```
### 도커 컨테이너 목록 확인
```shell
$ docker ps

CONTAINER ID IMAGE COMMAND CREATED STATUS PORTS NAMES

f407f746e9f4 hivv/carrot "docker-entrypoint..." 35 hours ago Up 35 hours 3001/tcp, 0.0.0.0:49160->3002/tcp confident_davinci
```


  