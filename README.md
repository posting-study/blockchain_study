# 블록체인을 해보자

세상에는 내가 모르는 것 참 많다. 모르는 것이 많은 만큼 열심히 공부할 수 있는 것도 넘쳐난다. 공부해야겠다는 생각이 든 순간 마침 군대로 오게 되었다. 복무하는 동안 배워야 할 것들을 정리하고 차근차근 배워갈 생각이다. 

> 군대에서 일 안하고 공부만 한다고 생각할 수 있는데 ~~(나도 처음에는 공부만 하는게 가능할 거라고 생각했는데)~~ 일도 매일하면서 남는 시간에 잠 줄여가면서 공부하는거다..ㅠ

우선 블록체인 공부를 목표로 해보기로 했다. 하지만 막상 책을 사고 펼쳐보니 내가 알아들을 수 있는 말이 정말 단 하나도 없었다. 
사실 그냥 실습 코드를 따라하거나 남이 만든 코드를 베껴서 적으면 어떻게도 만들 수는 있겠지만 어려운 분야인 만큼 그렇게 공부하는 것은 시간 낭비라고 생각되어서
조금 돌아가더라도 기초부터 차근차근 공부해보기로 했다.

### :checkered_flag:  천릿길도 네트워크 부터!
1. [Ehternet과 Bridge](./posting/network/1-ethernet.md) `(21-09-11)`
2. [Ip를 나누는 Class](./posting/network/2-ip.md) `(21-09-18)`
3. [Ip를 나누는 Subnetting](./posting/network/3-subnet.md) `(21-09-20)`
4. [Switch의 Spanning Tree Protocol](./posting/network/4-stp.md) `(21-09-25)`
5. [Switch의 VLan](./posting/network/5-vlan.md) `(21-10-02)`
6. [Router 와 Routing Protocol](./posting/network/6-router.md) `(21-09-25)`
7. [통신을 도와주는 기술 NAT](./posting/network/7-nat.md) `(21-10-02)`
8. [통신을 도와주는 기술 DNS](./posting/network/8-dns.md) `(21-10-09)`

### 🦕공룡이랑 함께하는 OS

1. [컴퓨터 구조](./posting/os/1-computer.md) `(21-11-28)`

### 📦 DB에 넣고 저장하자

### :computer: OS + Network + DB = Infra?

### :pray: 이제 블록체인 공부해도 될까?

#### 1. 블록체인이란?
거래를 블록으로 표현하여 `LinkedList` 형태로 연결한 구조가 되어 각각의 블록들은 `Hash` 를 통해서 다음 블록들을 가리키며, `RSA` 로 암호와되어서 `P2P Network` 를 통해서 관리되는 분산데이터 베이스의 한 형태입니다. 

블록체인은 분산원장 기술이라고 불리기도 합니다. `(DLT : Distributed Ledger Technology `
