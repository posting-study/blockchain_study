# 리눅스 명령어 모음
## 1. 쉘 관리
### 현재 내가 사용하고 있는 Shell 확인
```shell
$ echo  $0
```
## 2. 프로세스 관리
### 현재 돌아가는 포트목록 확인
```sh
$ netstat -tnlp
```
### 사용중인 프로세스 확인
```shell
$ top

top - 00:00:32 up 8 days, 1:32, 2 users, load average: 0.00, 0.01, 0.05

Tasks: 125 total, 2 running, 123 sleeping, 0 stopped, 0 zombie

%Cpu(s): 0.2 us, 0.2 sy, 0.0 ni, 99.6 id, 0.0 wa, 0.0 hi, 0.0 si, 0.0 st

KiB Mem : 7963288 total, 3359516 free, 719192 used, 3884580 buff/cache

KiB Swap: 15624188 total, 15624188 free, 0 used. 6563868 avail Mem
```
### 포트 번호로 PID 알아내기

```sh
$ sudo lsof -i :<PORT>
```
### PID 로 서비스 종료하기
```shell
$ kill -9 <PID>
```

## 3. 유저
### 유저 변경

```shell
$ su <user name>
```
### 유저 목록 확인
```shell
$ cat /etc/passwd
```

## 4. 디랙터리
### 디렉터리 구조 편하게 보기
```sh
$ tree .
.
├── config
│ ├── express.ts
│ └── vars.ts
└── index.ts
``` 
## 5. 정보확인
### Linux 버전 확인
```shell
$ uname -a
$ cat /etc/issue
$ grep . /etc/*-release
```
### 메모리 확인
```shell
$ cat /proc/meminfo
```
### CPU의 전체 정보
```shell
$ cat /proc/cpuinfo
```

### 코어의 수
```sh
$ grep -c processor /proc/cpuinfo
```
### 물리 CPU의 수
```
$ grep "physical id" /proc/cpuinfo | sort -u | wc -l
```
### CPU당 물리 코어 수
```
$ grep "cpu cores" /proc/cpuinfo | tail -1
```

