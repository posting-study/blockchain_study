# 정리하는 네트워크 (5)

지금까지 공부한 내용을 바탕으로 `Switch` 는 `Collision domain` 을 분리해주는 역할을 주로 한다고 정리했습니다. 만약에 회사 본관에서 **본관 네트워크** 를 이용하다가 별관으로 파견을 나갈 일이 있어서 별관으로 가게된다면
네트워크는 **별관 네트워크** 를 사용해야합니다. 만약에 본관 네트워크와 연결하고싶다면 **본관 네트워크**의 `Switch` 에서 선을 뽑아서 별관 네트워크까지 연결해야합니다. 

이는 굉장한 낭비입니다. **별관 네트워크** 로 간 호스트 1개 때문에 `Switch` 한 대가 낭비되게 됩니다.
만약 **같은 네트워크는 사용해야하는데 물리적으로 떨어져있는 상황**이 한 개가 아니라면? 낭비되는 `Switch` 는 한대로 끝나지 않습니다. 이러한 문제를 해결하기 위해서 `L2 Switch` 에는 `VLAN` 이라는 기능이 존재합니다.

### 🔌 VLAN

`VLAN` 은 물리적인 배치에 구애받지않고 논리적으로 `LAN (Local Area Network` 을 구성할 수 있게 해주는 기능입니다.  기존의 `Switch` 의 기능으로는 `Collision Domain` 을 나누는 일 밖에 하지 못하지만, `VLAN` 기능을 활용하면 같은 인터페이스에 연결되있더라도 `Broadcast Domain` 을 나눌 수 있게됩니다.

### Truck Port
먼저 `VLAN` 을 구성하게 되면 `Switch` 의  포트들이 각각의 `VLAN` 을 따라 나눠지게 됩니다. 
`VLAN` 영역이 분리되면 각각의 **`VLAN` 은 논리적으로는 다른 `Broadcast Domain`** 이기 때문에 같은 `Switch` 에 연결되어있더라도 통신하려면 `Router`가 필요합니다.  이러한 이유 때문에 `VLAN` 기능을 제공하는 `Switch` 는 서로 다른 `VLAN` 요청을 하나의 포트로 받을 수 있는데, 이러한 포트를 `Truck Port` 라고 합니다.

### :label: Tagging
`Truct Port` 를 통해서 이동하는 `VLAN` 의 데이터들을 각각 `Tag` 가 달려있어서 서로 구분됩니다.
> **Native VLAN**
> `Truct Port` 를 이용하는 VLAN 들 중에서 `Tag` 를 가지고 있지않은 `VLAN` 을 `Native VLAN` 이라고 합니다. 

### VTP ( VLAN Truct Protocol )
`VTP` 란 앞서 이야기한 `Truct Port` 의 전용 프로토콜로서 연결된 스위치 사이에서 `VLAN` 의 정보를 공유하는 역할을 합니다.

`VTP` 에는 통 3가지 모드가 존재합니다.

#### :one: Server Mode
`Server Mode` 에서는 각 `Switch` 들이 연결된 구성에서 단 하나만 존재하면서 해당 네트워크의 `VLAN` 정보를 추가 / 삭제 / 수정하는 역할을 담당합니다.
#### :two: Client Mode
`Client Mode` 는 `Server Mode` 와 달리 다른 `VLAN` 에 대한 정보를 추가하거나 삭제할 수 없습니다.  오로지 `Server Mode` 의 `Switch` 가 주는 `VLAN` 에 대한 정보를 받고 다른 `Switch` 들 에게 넘겨줄 수 있을 뿐입니다.
#### :three: Transparent Mode
`Transparent Mode` 는 위에서 설명한 것들과는 조금 다르게 특이합니다. 
자신의 `VLAN` 정보를 수정할 수 있지만 `Server Mode Swtich` 를 포함한  `Switch` 들 에게 해당 정보를 전송하지 않습니다. 하지만 `Switch` 에서 오는 정보들은 또 다른 `Switch` 로 전송하는 역할은 수행합니다.


