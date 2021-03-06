# :spider: 정리하는 네트워크 (2)

### IP ( Internet Protocol )
`IP` 는 네트워크 통신을 위한 장비들이 서로를 식별하기 위한 고유번호입니다. 각각의 아이피는 장비마다 고유하며 다른 장비와 겹치는 `IP` 가 존재해서는 안됩니다.

### :house: Building an Ip Address
총 32bit 로 이루어져있고 `Octet` 으로 구분하여  8bit 씩총 4개의 구역으로 이루어져있습니다.\
일반적으로 2진수가 아니고 10진수로 표현해서 
`128.10.2.41` 와 같은 형태가 됩니다.

|128| 10 | 2 |21|
|--|--|--|--|
|10000000|00001010|00000010|1000000010101|

### :chicken: Network Area and Host Area
우리나라의 집 전화번호 체계는 앞자리가 지역번호이고 나머지가 고유번호로 구성되어있습니다.

<table>
  <tr>
    <th colspan='1'>지역 번호</th>
    <th colspan='2'>고유 번호</th>
  </tr>
  <tr>
    <td>031</td>
    <td>xxx</td>
    <td>xxxx</td>
  </tr>
</table>

위와같은 개념으로 네트워크 영역과 호스트 영역은 `IP` 가 가지는 또다른 규칙입니다.
**하나의 네트워크 (LAN) 안에서의 호스트들은 서로 같은 네트워크 영역을 가져야합니다. ** 경기도 지역에 포함되는 전화번호는 모두 `031` 로 시작해야하는 것과 같은 개념입니다.  조금 더 자세하게 설명하면, `128.10.2.21` 이라는 `IP` 주소에서
앞의 두 영역인 `128.10` 을 네트워크 영역이라고 한다면
나머지 `2.21` 을 호스트 영역이라고 합니다.
<table>
  <tr>
    <th colspan='2'>Network Area</th>
    <th colspan='2'>Host Area</th>
  </tr>
  <tr>
    <td>128</td>
    <td>10</td>
    <td>2</td>
    <td>21</td>
  </tr>
</table>

위와같은 네트워크 영역안에서는 `128.10.2.21` 과 `128.10.2.22` 는 같은 네트워크 소속입니다. 해당 네트워크 장비들은 같은 네트워크 안에 속해있기 때문에
`Gateway` 장비 없이 서로 네트워크 통신이 가능합니다.

### :door: Network Class
방대한 양의 `IP` 주소를 체계적으로 나눠주기 위해서는 네트워크 영역과 호스트 영역을 좀 더 효율적으로 나눌 수 있는 규칙이 필요합니다.

`IP` 체계는 영역 분리를 위해서 `Class` 라는 개념을 도입했습니다. 
`IP Class` 는 총 5개로 나누어져있고 `A` 부터 `E` 까지 존재합니다.
단, `D Class` 와 `E Class` 는 연구용으로 잘 사용하지않습니다.

#### :a: Class
`IP` 주소의 앞자리가 2진수로 표현하였을 때 0으로 시작하면 `A Class IP` 입니다.

해당조건을 만족하는 `IP` 는 

`0000 0000.0000 0000.0000 0000.0000 0000` 부터
`0111 1111 1111 1111 1111 1111 1111 1111` 까지 입니다.

이것을 10진수로 표현하면
`0.0.0.0` ~ `127.255.255.255` 입니다.

여기서 `0.x.x.x` 은 자체 네트워크를 의미하기 때문에 제외하고 `127.x.x.x` 도
`loopback (자기 자신)` 을 의미하기 때문에 제외하게 됩니다.

정리하면 `A Class` 로 할당가능한 `IP` 범위는 최종적으로
`1.0.0.0 ~ 126.255.255.255` 가 됩니다.

`A Class` 는 앞의 8 bit 가 네트워크 영역이고 뒤의 24bit 가 호스트 영역입니다.

예를 들어 `A Class` 네트워크영역 `34.xx.xx.xx` 를 할당받게 된다면 이 네트워크 영역에 할당할 수 있는 호스트들은 2<sup>24</sup>-2  개 가됩니다.
> **-2 를 하는 이유**  
> `A class` 에서 xx.0.0.0 은 네트워크 영역 자체를 의미하고
> xx.255.255.255 는 `broadcast domain` 으로 사용하기 때문에 제외합니다.
<table >
  <tr>
    <th colspan='1'>Network Area</th>
    <th colspan='3'>Host Area</th>
  </tr>
  <tr>
    <td>34</td>
    <td>10</td>
    <td>2</td>
    <td>21</td>
  </tr>
  <tr>
    <td><b>0</b>010 0010</td>
    <td>0000 1010</td>
    <td>0000 0010</td>
    <td>0001 0101</td>
  </tr>
</table>
  
  #### :b: Class

`IP` 주소의 앞자리가 2진수로 표현하였을 때 **10** 으로 시작하면 `B Class IP` 입니다.

해당조건을 만족하는 `IP` 는 

`1000 0000.0000 0000.0000 0000.0000 0000` 부터
`1011 1111 1111 1111 1111 1111 1111 1111` 까지 입니다.

10진수로 표현하게 되면 
`128.0.0.0` ~ `191.255.255.255` 입니다.

`B Class` 에서는 앞의 16bit 가 네트워크 영역이고 뒤의 16bit 가 호스트 영역으로 이루어져있습니다.

할당할 수 있는 네트워크 영역은 앞의 이진수 `10` 이 고정이기 때문에 2<sup>14</sup> 개 이고 호스트 영역은 `A class` 와 같은 이유로 2<sup>16 </sup> - 2  개 입니다.

<table >
  <tr>
    <th colspan='2'>Network Area</th>
    <th colspan='2'>Host Area</th>
  </tr>
  <tr>
    <td>128</td>
    <td>10</td>
    <td>2</td>
    <td>21</td>
  </tr>
  <tr>
    <td><b>10</b>00 0000</td>
    <td>0000 1010</td>
    <td>0000 0010</td>
    <td>0001 0101</td>
  </tr>
</table>

#### C class 

`IP` 주소의 앞자리가 2진수로 표현하였을 때 **110** 으로 시작하면 `C Class IP` 입니다.

해당조건을 만족하는 `IP` 는 

`1100 0000.0000 0000.0000 0000.0000 0000` 부터
`1101 1111 1111 1111 1111 1111 1111 1111` 까지 입니다.

10진수로 표현하게 되면 
`192.0.0.0` ~ `223.255.255.255` 입니다.

`B Class` 에서는 앞의 24bit 가 네트워크 영역이고 뒤의 8bit 가 호스트 영역으로 이루어져있습니다.

할당할 수 있는 네트워크 영역은 앞의 이진수 `110` 이 고정이기 때문에  2<sup>21</sup> 개 이고 호스트 영역은 `A class` 와 같은 이유로  2<sup>8</sup> 개 입니다.

<table >
  <tr>
    <th colspan='3'>Network Area</th>
    <th colspan='1'>Host Area</th>
  </tr>
  <tr>
    <td>192</td>
    <td>10</td>
    <td>2</td>
    <td>21</td>
  </tr>
  <tr>
    <td><b>110</b>00 0000</td>
    <td>0000 1010</td>
    <td>0000 0010</td>
    <td>0001 0101</td>
  </tr>
</table>

### 정리
`IP` 의 구성과 체계에 대해서 알아봤습니다. (3) 에서는 좀 더 세부적인 분류방식인 `Subnetting` 에 대해서 포스팅 해보겠습니다.
