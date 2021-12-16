# 메모리

PC에 저장되어있는 프로그램을 실행시키면 `Process`는 운영체제의 `Memory` 위에서 실행됩니다. 하나의 `Process` 가 하나의 `Memory`에서 실행된다면 단순하게 끝날 문제이지만 **시분할 시스템** 에서는 다양한 `Process`가 `Memory`에 동시에 올라와서 실행됩니다.  이때 모든 메모리 관리는 `MMS(Memory Management System)` 이 담당하게됩니다. 이때 `Process`는 `Memory`를 독차지하려고하고 `MMS` 는 관리를 효율적으로 하려고 하기 때문에 서로 충돌이 일어납니다. 이처럼 현대의 메모리 관리는 과거보다 훨씬 복잡해졌습니다.

### 소스코드와 컴파일

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

### 메모리 관리의 개요

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


