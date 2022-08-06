# Firewall

### 방화벽 작동확인
```shell
$ systemctl state firewalld
```
### 방화벽 설정 정보 확인
```shell
$ firewall-cmd --zone=public --list-all
```
### 특정 IP 허용하기
```shell
$ firewall-cmd --permanent --add-source=<IP Address>
```
### 특정 Port 허용하기
```shell
$ firewall-cmd --permanent --zone=public --add-port=포트번호/tcp
```

### 방화벽 다시 실행

```shell
$ firewall-cmd --reload
```