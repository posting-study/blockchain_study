# Memory

PC에 저장되어있는 프로그램을 실행시키면 `Process`는 운영체제의 `Memory` 위에서 실행됩니다. 하나의 `Process` 가 하나의 `Memory`에서 실행된다면 단순하게 끝날 문제이지만 **시분할 시스템** 에서는 다양한 `Process`가 `Memory`에 동시에 올라와서 실행됩니다.  이때 모든 메모리 관리는 `MMS(Memory Management System)` 이 담당하게됩니다. 이때 `Process`는 `Memory`를 독차지하려고하고 `MMS` 는 관리를 효율적으로 하려고 하기 때문에 서로 충돌이 일어납니다. 이처럼 현대의 메모리 관리는 과거보다 훨씬 복잡해졌습니다.

### Source Code and Compiler

##### 목적코드 변환
코드가 작성되면 일반적으로 `Compiler`가 해당 코드를 `Compile` 하여서 기계어인 `Object Code(목적 코드)` 변환한다. 사람이 읽을 수 있는 소스코드에서 기계어로 컴파일함으로서 비로소 컴퓨터가 읽을 수 있는 프로그램이 됩니다.

##### 링커
여기서 끝나는 것이 아니라 우리가 C언어 코드를 작성할 때 아래와 같이 작성하는 부분은 운영체제가 제공하는 라이브러리나 C언어 라이브러리를 가져오는 과정이 포함되어있습니다. 
```c
#include <stdio.h>
#include <math.h>
```
``printf()`` 와 같은 함수를 쓰게되면 `<stdio.h>` 라이브러리에서 가져와서 목적코드로 변환이됩니다. 이러한 과정은 `Linker`가 수행합니다.
##### 동적라이브러리
라이브러리를 가져오는 과정까지 마치게되면 `Dynamic libary`를 추가하게 됩니다. `Dynamic libary`는 라이브러리에 새로운 기능이 생길 때 일일히 업데이트하지않아도 동적으로 추가되는 기능을 가지고있는 라이브러리입니다.

동적라이브러리까지 추가하게되면 모든 과정이 끝나고 프로그램이 실행되게됩니다.

### Memory Management

메모리관리는 정확히 `MMU(Memory Management Unit` 이라는 하드웨어가 담당하게됩니다. 이 장치는 크게 3가지 일을 수행하게됩니다.

- Fetch
- Placement
- Replacement
#### Fetch
`Fetch` 는 `Process`와`Data`를 `Memory`로 가져오는 작업입니다.
#### Placement
`Placement`는 가져온 `Process`와 `Data`를 `Memory` 의 어느곳에 배치할지 결정하는 작업입니다.  이때 메모리의 크기에 맞게 `Process`를 자르기도 하는데, 
같은 크기로 자르는 것을 `Paging`, 동적으로 자르는 것을 `Segementation` 이라고 합니다.
#### Replacement
`Replacement`는 메모리가 꽉 찾을 때 어떤 데이터를 내보낼지 결정하는 작업입니다. 이는 `Replacement Algorithm`에 따라 결정됩니다.

### Memory Address
컴퓨터를 사용하다보면 `32bit` 프로그램과 `64bit` 프로그램들을 접하게됩니다.
요즘은 거의 `64bit` 프로그램이 대부분이지만 옛날에는 프로그램을 사이트에서 다운받을 때도 두가지 선택권이 있었습니다.

그렇다면 `32bit` `64bit`에는 무슨 차이가 있을까요? 이는 `CPU`의 한번에 처리가능한 데이터의 최대 크기를 의미합니다. 즉 `32bit CPU` 는 한번에 `32bit`를 처리할 수 있고 그 컴퓨터의 `Register`도 `32bit`이고 데이터가 이동하는 `BUS`도 `32bit`가 됩니다. 당연하게도 `ALU`도 `32bit`가 됩니다.

`CPU`가 `32bit`라면 `MAR(Memory Address Register`도 `32bit`이기 때문에 표현할 수 있는 메모리 주소 범위가 **0~2<sup>32</sup>-1 = 2<sup>32</sup>B** 로 약`4GB` 이며 최대 메모리가 `4GB`가 됩니다.

이에 비해서 `64bit`는 **2<sup>64</sup>** 로 약 `16,777,216TB`로 거의 무한대에 가까운 메모리를 가질 수 있습니다.
|구분|32bit CPU|64bit CPU|
|--|--|--|
|주소 범위|0~2<sup>32</sup> -1 번지 | 0~2<sup>64</sup> -1 번지 
|총 크기| 2<sup>32</sup>(약 4GB) | 2<sup>64</sup>(약 16,777,216TB)

당연하게도 `CPU bit` 에 따라서 그에 맞는 운영체제가 설치되어야합니다.

### :vs: Physical vs Logical

![](https://user-images.githubusercontent.com/22852287/147376584-c25fb64d-cb16-4f4a-8d13-ed279520aed5.png)

메모리에는 운영체제를 구동하기위한 메모리가 미리 탑제되어있습니다.

그리고 메모리 저장방식에는 `Physical Address` 와 `Logical Address` 를 활용하는 방법으로 나뉩니다.

절대 메모리 주소는 앞서 탑제되어있는 운영체제 영역을 포함해서 사용자 영역에 있는 데이터의 위치를 표현하는 방식입니다.

### :arrow_left::arrow_right: Swap
`4GB Memory`를 사용하고 있을 때 `8GB Program`을 돌리려고 한다면 프로그램 하나가 메모리 위에 전부 못 올라가는 현상이 벌어집니다. 이러한 상황을 `Memory Overlay` 라고 합니다. 
한정된 메모리 안에서 큰 프로그램을 실행시키기 위해서 프로그램의 일부 모듈만 나눠서 메모리에 올리는 전략을 사용합니다. 예를들어서 포토샵을 실행시키면 포토샵의 모든 기능을 전부 메모리에 올리는 것이 아니라 기본적인 기능만 `Common Module` 로 묶어서 메모리에 올리고 나머지 기능은 사용할 때 메모리로 불러오는 전략을 사용합니다.

이때, 사용하지 않고있는 모듈들을 단순하게 하드디크에 보관하기에는 다시 로딩하는데 시간이 오래걸릴 수 있기때문에 하드 디스크에 특별한 영역에 저장해둔다.

![](https://user-images.githubusercontent.com/22852287/147376720-415b987b-285d-4927-b75d-85e7bb26aaff.png)





