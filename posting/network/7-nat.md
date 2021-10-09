# :spider_web: 정리하는 네트워크 (7)
> 작성 시작 : 10월 2일 (토)
> 작성 완료 : 10월 9일 (토)
### :computer: NAT ( Network Address Translation )
`IP` 는 42억개의 경우의 수를 가지고있지만 현대에 와서 42억개의 `IP` 는 턱없이 부족합니다. `IP` 주소 고갈 대첵으로써 **단기계획**으로는 `Class` 와 `Subnet` 을 적용하여 대역을 나누고 체계적으로 관리하며 효율성을 증대시켰습니다.
**장기계획**으로는 `IPv6` 로 변환이 있습니다. 

**단기계획** 과 **장기계획** 사이에는 `IP`  고갈을 해결하려고 만든 **중기계획** 이 존재하는데, 이 계획이 바로 `NAT ( Network Address Translation )`  입니다.

`NAT` 기술은 부족한 `IP` 를 `사설 IP` 와 `공인 IP` 로 나누는 것 부터 시작됩니다. 

`NAT` 가 내장된 `Router`나 `L3 Switch`에 `공인 IP` 를 할당하고 그 라우터 밑에 연결된 장비들은  `NAT` 가 제공하는 `사설 IP` 를 부여받습니다.

여기서 `사설 IP` 란 `NAT` 범위 내에서 사용되는 `IP` 이므로 다른 대역에서는 
`IP` 가 겹쳐도 상관이 없습니다.

> **NAT / PAT ( NAPT )**
> 원래 `NAT` 는 네트워크 주소를 1:1 로 변환 해주는 기술입니다. 지금은 `IP` 문제를 해결하기 위해서 위에 설명한 기능을 수행하는 것도 `NAT` 라고 부르지만,
> 원래 1:N 으로 변경해주는 것은 `PAT ( Port Address Translation` 혹은 `NAPT ( Network Address Port Translation`  이라고 부릅니다. 
> 

### :hammer: NAT 의 동작 방식
##### :one: Host-> Gateway ( NAT )
먼저 호스트에서 패킷이 출발하게 됩니다. 출발시에는 패킷 안에 출발지의 `Private IP` 가 담겨있습니다. 
##### :two: Gateway (NAT)
호스트에서 패킷을 받은 공유기는 해당 패킷의 출발지 주소를 공유기가 가지고 있는`Public IP` 로 변경하게 됩니다. 이후에 해당 `Private IP` 와 `Public IP` 를 
`NAT Table` 에 기록합니다.

##### :three: Destination -> Gateway
목적지에 다시 `Gateway` 로 응답 패킷을 보내게 됩니다. 응답 패킷의 목적지 IP `NAT` 장비의 `Public IP`  가 됩니다.

##### :four: Gateway -> Host
`NAT` 장비는 들어온 응답 패킷이 어떤 요청 패킷에 대한 응답 패킷인지 확인하고
`NAT Table` 을 조회하여 최초 응답을 보냈던 `Private IP` 를 목적지로 하고
`NAT` 장비의 `Public IP` 를 출발지로 하여서 패킷을 돌려주게됩니다. 

### SNAT and DNAT
##### SNAT ( Source NAT )
`SNAT` 는 출발지의 `IP` 를 변경하는 `NAT` 입니다. 
##### DNAT (Destination NAT )
`DNAT` 는 목적지의 `IP` 를 변경하는 `NAT` 입니다.


