# :spider: 정리하는 네트워크 (8)
> 작성 시작 : 10월 9일 (토)

지금까지 네트워크에 대해 기본적인 내용을 정리하면서 `IP` 에 관련한 이야기를 많이 다뤘습니다. 하지만 인터넷의 사용자는 정작 `IP` 에 대해서 잘 몰라도 인터넷을 자유자제로 사용할 수 있습니다. 실제로 우리는 구글에 접속하고싶다면 구글 웹 페이지를 호스팅하는 서버의 `IP` 를 검색창에 넣지 않고 `www.google.com` 을 입력합니다. 이러한 일이 가능한 이유는 `DNS` 가 `www.google.com` 을 구글 서버의 `IP` 로 변경해주기 때문입니다. 

### :book: DNS ( Domain Name System )
`DNS` 는 기본적으로 `IP` 주소와 `Domain Name` 을 매핑해주는 역할을 수행합니다. `DNS` 는 이런 기본적인 작업 이외에도 여러가지 역할을 수행할 수 있는데, 서버에 연동하는 상황을 예로 `DNS` 를 사용하지 않고 `Client` 측에서 `Server Gateway IP` 인 `22.22.22.22` 로 연동했다고 가정합니다.

만약 `Client` 가 `APP` 이라면 `22.22.22.22` 을 `Gateway` 로 지정하여 이미 스토어에 배포가 되어있는 상황입니다. 그런데 이 상황에서 `Server` 를 이전해야한다면 `Server IP` 가 변경이 됩니다. `Server IP` 가 변경될 때 마다 `APP` 을 다시 배포해야하는 번거로운 상황이 발생하는겁니다.

`DNS` 를 사용하여 `22.22.22.22` 을 `api.app.com` 으로 매핑하였다면 `APP` 에서는 `Gateway Server` 를 `api.app.com` 으로 작성하여 배포하면 됩니다.
`Server IP` 가 `22.22.22.22` 에서 `11.11.11.11` 로 변경되더라도 `DNS` 매핑 정보만 수정하면 `APP` 을 다시 배포해야할 필요가 없습니다.

### DNS 구조
|Third-Level Domain|Second-Level Domain|Top-Level Domain (TLD)|Root|
|--|--|--|--|
|www|naver|com|.|

`DNS` 는 `Root` 부터 `Third-Level Domain` 까지 작성순서와는 반대로 찾아갑니다.

### DNS 작동방식
> **Recursive Query** and **Iterative Query**
> DNS의 작동에는 `Recursive Query` 와 `Iterative Query` 가 사용됩니다.
> 
우리가 브라우저 검색창에 **처음** `www.google.com` 을 검색했을 때 일어나는 과정입니다.
##### :one: Query hosts
윈도우 컴퓨터에는 `hosts` 파일이 존재합니다. `hosts` 파일은 컴퓨터의 `DNS` 전화번호부라고 생각하면 됩니다. 이전에 매칭되었던 기록이 `Cache` 로 존재합니다.
처음 `www.google.com` 을 검색한 것이기 때문에 `hosts` 파일에서는 `Google IP` 를 찾을 수 없습니다. 
#####  :two: DNS Server
`Local Cache` 에서 `IP` 정보를 확인할 수 없었기 때문에 사용자 컴퓨터에 지정된 `DNS Server IP`로 요청을 보내게 됩니다.
<center>
<img src='https://www.dazzii.com/wp-content/uploads/2018/12/dns7.jpg' width='400'> 
</center>
`DNS Server IP` 는 위 사진처럼 확인가능합니다. 요청을 받은 `DNS Server` 는 `www.google.com` 에 대한 정보가 없기때문에 `DNS Server` 에 저장되어있는 `Root NS IP` 로 `www.google.com` 에 대한 `IP` 를 요청하는 쿼리를 날리게됩니다. 

##### :three: Root NS 
`Local DNS` 에서 찾을 수 없었기 때문에 사용자 컴퓨터에 지정되어있는 `Root NS IP` 로 `www.google.com` 의 `IP`를 알려달라는 쿼리를 보내게됩니다. 처음 요청을 받은 `Root NS` 는 
`Google` 에 대한 정보는 없지만 `.com` 에 대한 `TLD Server IP` 는 알고 있기때문에 `com` 의
`TLD Server IP` 를 다시 사용자에게 반환합니다.
##### :four: TLD Server
`TLD Server IP` 를 반환받은 `DNS Server` 는 `TLD Server` 에게 다시 한번 `www.google` 에 대해서 요청합니다. `com TLD Server` 는 요청을 확인하고 `Google NS IP` 를 반환합니다. 
##### :five: Second-Level Domain
`Google NS IP` 를 반환받은 `DNS Server` 는 `Google NS` 에게 `IP` 를 요청하고 `Google NS` 는 `DNS Server` 에게 `Google Gateway Server IP` 를 응답해주게 됩니다. `IP` 를 받은 `DNS Server` 는 요청을 한 `Host` 에게 `IP` 를 전달합니다.
##### :six: Host ( User )
`IP` 를 전달받은 `Host` 는 해당 `IP` 로 요청을 날리게되고 비로소 `Google` 의 웹 페이지를 반환받아 웹 브라우저에 띄울 수 있게됩니다.

### 💿 DNS Record
`Domain` 에는 `IP` 매핑을 포함한 다양한 추가 매핑 레코드가 존재합니다.

- **A Record**
가장 기본적인 일을 수행하는 레코드입니다. `Domain` 의 `A Record` 에는 매핑하고자 하는 `IP` 가 들어갑니다.

- **AAAA Record ( IPv6 )**
역할은 `A Record` 와 똑같지만 `IPv6` 를 매핑하기 위한 레코드입니다.
- **CNAME**
`CNAME` 은 `IP` 와 `Domain` 을 매핑하지않고 `Domain` 과 `Domain` 을 매핑합니다.
그 예시로 `www.google.com` 은 `google.com` 으로 매핑이 되어있습니다. `CNAME` 을 사용하면 
`DNS Record` 를 관리하기 편해집니다. `www.google.com` 과`google.com` 을 모두 `IP` 에 매핑하게 된다면 `IP` 가 변경될 때 이 2개를 모두 수정해야하지만 `www.google.com` 은 `google.com` 에 매핑시켜두고 `google.com` 만 `A Record` 를 작성해둔다면 `IP` 가 변경되더라도 `A Record` 하나만 수정한다면 `www.google.com` 을 건드릴 필요가 없어집니다.
- **TXT ( TeXT Record )**
간단한 텍스트를 입력할 수 있는 레코드입니다. 보통 `Domain`에 대한 소유를 확인하는 과정에서 사용되거나 메일의 발신자를 입증하는 `SPF (Sender Policy Framework)` 에 이용됩니다.


