# 정리하는 네트워크 (4)

스위치는 학습되지 않은 `MAC` 의 통신을 받을 때는 기본적으로 전체에게 넘기는 
`Flooding` 을 하게됩니다. 

두개의 스위치가 서로연결되어있을 때 서로 `Flooding` 을 주고받게 되면 결국 돌고도는 `Looping` 이 발생하게됩니다. 

이러한 상황을 방지하기 위해서 `Switch`에는 `STP(Spanning Tree Protocol` 이라는 기술이 적용되어있습니다. 

### STP (Spanning Tree Protocol)
`STP` 는 복수개로 연결되어있는 스위치 사이에서 중복되는 연결을 비활성화시키고 있다가 장애가 발생하면 비활성화 되어있던 연결을 부활시키는 방식으로 `Switch` 간의 연결을 제어하는 프로토콜입니다.  

### STP 의 간단한 원리
먼저 `Switch` 의 통신이 기준이 될 `Root Switch` 를 선정하고 나머지 `Switch` 들은 
`Non Root Switch` 가 됩니다.  
`Non Root Switch` 는 `Root Switch` 와 연결되어있는 `Root Port` 와 해당 포트를 제외한 `Deginated Port` 를 선정하게 됩니다. 
이외의 포트는 `Looping` 을 방지하기 위해서 막아버립니다.

### BPDU ( Bridge Protocol Data Unit)
`Switch` 가 서로 소통하며 `Spanning Tree` 를 구성할 때 사용하는 `Ethernet Packet` 입니다.

`BPDU` 에는 `BridgeID` 라는 것이 포함되어있는데, `Bridge ID` 는 `priority` 와 해당 `Switch` 의 `MAC Address` 로 이루어져있습니다.
|2 Byte|6 Byte|
|-|-|
|Priority|MAC Address|

`Switch` 들은 서로 `BPDU` 를 주고받으면서 서열을 결정하게됩니다.

### STP 의 동작원리 1 :: Root Bridge 선출
규칙 1 : **가장 낮은** `Bridge ID` 를 가진 `Switch`가 `Root Bridge` 로 선출된다.

처음 전원이 켜진 `Switch` 들은 다른 `Switch`에 대한 정보가 없기 때문에 자기 자신을 `Root Bridge` 라고 생각합니다. 각각의 `Switch`들은 서로에게 자신이 `RootBridge` 라고 주장하기 위해서 자신의 `Bridge ID` 를 `BPDU` 에 담아서 보내게 됩니다.

서로 `BPDU` 를 교환하고 그 중에 `Bridge ID` 가 제일 낮은 `Switch`가 `Root Bridge` 로 선출되게 됩니다. `Root Bridge` 로 선출되지않은 `Switch`들은 자연스럽게 `Non Root Bridge`가 됩니다.

### STP 의 동작원리 2 :: Root Port 선정

규칙 2: `Non Root Bridge`에서 가장 낮은 `Path Cost`를 가진 포트가 `Root Port`가 된다.

쉽게 말해서 `Non Root Bridge`에서 `Root Bridge`로 가는 비용이 가장 적게 드는 포트가 선정되는 것 입니다. 

### STP 의 동작원리 3 :: Degignated Port 선정

규칙 3: 한 `Segment`에서 `Path Cost`가 가장 낮은 포트가 `Degignated Port`로 선정됩니다.

`Switch`끼리 연결되어있는 부분을 `Segment`라고 하는데 한 `Segment` 영역에는
서로다른 `Switch`의 포트가 각각 1개씩 존재합니다.

### STP 의 상태변화

|상태|BPDU|MAC 학습|데이터 전송||
|-|-|-|-|-|
|**Disabled**|X|X|X|
|**Blocking**|O|X|X
|**Listening**|O|X|X|
|**Learning**|O|O|X|
|**Forwarding**|O|O|O|

- **Disabled** 상태에서는 포트가 고장나서 `Shutdown` 시켜둔 상태입니다. 이때에는 어떠한 통신도 불가능합니다.

- **Blocking** 상태는 맨 처음 스위치를 켰을 때 머무르는 상태입니다. `Root Bridge`, `Root Port`, `Degignated Port` 를 이때 선정합니다.

- **Listening** 상태는 `Degignated Port` 까지 선정이 완료되면 넘어가는 상태입니다. 이 상태에서 새로운 `Switch` 의 연결이 감지되면 다시 `Blocking` 상태로 넘어가게 됩니다.
- **Learning** 은 `Listening State` 에서 `Forwarding Delay (15s)` 만큼을 대기하면 넘어가는 상태입니다. 이때부터 `Switch` 는 `MAC Address` 를 학습하게 됩니다.
- **Forwarding** 은 `Learning` 상태에서`Forwarding Delay (15s)` 를 한번 더 기다리게되면 해당 상태로 넘어가게 됩니다. `Forwarding State` 로 넘어와야만 데이터 전송이 가능해지면서 비로소 통신이 가능하게됩니다.