# SSL

모든 작업은 `root` 권한이 필요함

### Epel, Certbot 설치

```shell
$ yum -y install epel-release certbot
```
### Nginx 중단

```shell
$ systemctl stop nginx
```
### Certbot 으로 SSL 인증서 발급

```shell
$ certbot --standalone -d <domain name> certonly
```
### Nginx Proxy 에 반영
```nginx
server {
	listen 443 ssl http2;
	listen [::]:443 ssl http2;
	server_name <domain name>;
	ssl_certificate /etc/letsencrypt/live/<domain name>/fullchain.pem;
	ssl_certificate_key /etc/letsencrypt/live/<domain name>/privkey.pem;
	location / {
		proxy_redirect off;
		proxy_pass_header Server;
		proxy_set_header Host $http_host;
		proxy_set_header X-Real-IP $remote_addr;
		proxy_set_header X-Scheme $scheme;
		proxy_pass <hosting url>;
	}
}

```
### Nignx 재실행

```shell
$ systemctl start nginx
```

  
  
  
