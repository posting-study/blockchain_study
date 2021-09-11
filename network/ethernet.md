# :spider: 정리하는 네트워크

> 영어 공부 좀 할 겸 장비나 용어들은 전부 영어로 표기했습니다.  
> 오타가 좀 많을 수 있습니다. 지적해주시면 바로 수정하겠습니다.

컴퓨터에서 컴퓨터로 데이터를 옮기려면 어떻게 해야 할까? 보통은 귀찮아서 카카오톡 나에게 보내기 기능으로 자료를 옮기거나 구글 드라이브에 올리곤 한다. 이 모든 일이 쉽게 이루어질 수 있는 이유가 무엇일까요?

### :computer: Ethernet

오늘날의 통신방식을 이해하려면 먼저 이더넷의 동작 원리 부터 알아볼 필요가 있다.  
이더넷은 `CSMA/CD (Carrier Sense Multiple Access / Collision Detection)` 방식의 기술 기반으로 이루어져있다.

[##_Image|kage@n7Yjm/btreMEhIy7i/GkStUEDbsIBChSAkui5fs1/img.png|alignCenter|width="100%" data-origin-width="921" data-origin-height="639" data-ke-mobilestyle="widthOrigin"|||_##]

  
`CSMA/CD` 방식은 일종의 눈치게임 입니다. 위의 사진처럼 4대의 컴퓨터가 `Ethernet` 방식으로 연결되어있다고 가정해봅니다. `CSMA/CD` 방식에서는  
**A 컴퓨터** 와 **C 컴퓨터** 가 통신중일 때는 다른 기기들은 보낼 통신이 있어도 대기하게됩니다. 그러다가 통신이 종료되면 그때 전송하게됩니다.

혹여나 동시에 데이터를 전송하게되어서 서로 데이터가 충돌하게 된다면 이것을  
`Collision` 이 발생했다고 합니다. `Collision` 이 발생하게되면  
각 기기들은 랜덤한 시간만큼 대기하였다가 다시 전송을 진행하게됩니다.

해당 방식은 겉보기에는 좀 없어보이는 방식일 수 있습니다. 하지만 가성비가 굉장히 좋기때문에 다른 방식들 보다 오래 살아남게 되었고 오늘날 인터넷 통신방식의 기본이 되었습니다.

### :apple: Network Interface Card

집에 있는 데스크탑을 뜯어보면 여러 장비중에 `LAN Card` 라는 장비가 있을 것 입니다. 우리가 보통 `LAN 선` 이라고 부르는 것을 꽂아두는 컴퓨터 장치중에 하나인데, 이 장치가 `Ethernet` 에서 뿌려주는 데이터를 식별해서 자신의 컴퓨터로 오는 데이터만 골라서 `CPU` 로 보내주는 역할을 합니다.

### :banana: Broadcast And Unicast

통신방식에는 대표적으로 `Broadcast` 와 `Unicast` 가 있습니다.

여러 대의 컴퓨터가 하나의 허브로 연결되어있을 때 `A -> B Unicast` 통신이 일어나면 해당 허브에 연결되어있는 모든 컴퓨터로 데이터가 전송됩니다.  
하지만 통신에는 도착지의 `MAC` 이 동봉되어있기 때문에 이 `MAC` 이 아닌 PC들은  
`NIC` 가 알아서 거르게됩니다. 이를 통해서 `CPU` 의 부담을 줄일 수 있게됩니다.  
거르지 못한다면 다른 PC들이 통신중일 때 통신중이지 않는 엉뚱한 PC에 부하가 걸리게 됩니다. 이러한 방식을 `Unicast` 라고 합니다.

반면 `Broadcast` 방식은 송신자의 기기를 제외하고 `LAN` 에 연결되어있는  
모든 기기에 데이터를 전송하게됩니다. `Broadcast` 통신은 `Unicast` 와 다르게  
`NIC` 가 데이터를 `CPU`에게 전송하게 됩니다. 이러한 이유로 `Broadcast` 통신을 남용하게되면 기기에 무리가 가게됩니다.

### :cloud: Collision Domain

[##_Image|kage@bmA96d/btreJhofyfT/JgJ32vNfKhCcxQForfw040/img.png|alignCenter|width="100%" data-origin-width="1179" data-origin-height="617" data-ke-mobilestyle="widthOrigin"|||_##]

  
위에서 `Ethernet`은 연결된 네트워크 끼리 눈치게임을 진행한다고 비유했었습니다.  
이러한 눈치게임을 진행하는 영역을 `Ehternet Segment` 라고 부르는데,  
`Ehternet Segment` 안에서는 앞서 설명한 `Collision Detect` 가 반복해서 발생하게 됩니다. `Collision Detect` 가 발생하는 영역을 **Collision Domain** 이라고 합니다..

이제부터는 연결된 기기/PC 들은 `Host` 라고 부르겠습니다.

`Ehternet Segment` 에 연결된 `Host` 가 많을수록 `Collision Detect` 가 증가하게되고 `Collision Domain` 도 커지게 됩니다. 이렇게 되면 중간에 패킷이 손실되는 등 정상적인 통신이 이루어지기 힘들게됩니다. 이러한 문제는 어떻게 해결해야할까요?

### :candy: Bridge

`Bridge`는 위에서 말한 `Collision Domain` 을 분리해주는 역할을 합니다. 말 그대로 `Collision Segment` 사이에서 다리의 역할을 하는 것입니다.

`Bridge` 는 5가지 특성을 가집니다.

-   Learning ( MAC 주소 학습 )
-   Flooding ( 전체에게 전달 )
-   Forwarding ( 다른 영역에 전달 )
-   Filtering ( 다른 영역으로 통신 차단 )
-   Aging ( 학습된 주소 삭제 )

`Collision Domain` 이 분리되어있는 상황에서 첫 통신이 일어나게 되면  
`Bridge` 는 통신을 전송한 기기의 `MAC Address` 을 **Learning** 하게됩니다.  
학습한 `MAC Address`는 `Bridge Table` 에 저장됩니다.

첫 통신이기 때문에 `Bridge` 는 이 통신의 목적지가 `Collision Domain` 안에 있는지 밖에 있는지 알지못하기 때문에 **Flooding** 방식으로 모든 `Host` 에 데이터를 전송하게됩니다.

이후 `Bridge Table` 에 기록된 통신이 일어나고 이 통신이 반대쪽 `Collision Domain` 에 `Host` 에게 보내는 통신이라면 `Bridge` 는 이 통신을 반대쪽 `Collision Domain` 으로 **Forwarding** 해주게됩니다.

만약, 해당 통신이 반대쪽 `Collision Domain` 에 보내는 통신이 아니라 송신자측에 `Collision Domain` 에게 전송하는 것이라면 `Bridge` 가 통신이 반대쪽으로 넘어가지 못하도록 **Filtering** 을 해줍니다.

너무 많은 정보가 `Bridge Table` 에 기록되게 되면 `Bridge` 자체에 부담이 가게된다. 이러한 문제를 해결하기 위해서 `Bridge Table` 는  
일정한 시간이 지나면 데이터가 지워지는 **Aging** 기능이 탑제되어있습니다.

사실 지금은 `Bridge` 의 상급장비인 `Switch` 를 더 많이 사용합니다. 역할은 비슷하지만 `Switch` 가 기능이 더 많기때문에 가격이 비교적 비쌉니다.

### 향후 계획

이후에는 `IP` 와 `Router` 에 대해서 자세히 알아보도록 하겠습니다.