# 💾 Process and Thread
### 🔨 Process 란?
`Program`이 `Memory` 위로 올라와서 실행되면 **`Process`** 형태가 됩니다.
`OS`도 하나의 프로그램이기때문에 컴퓨터가 실행되고 있는 동안에는 `Memory`
위에 `OS Process`도 존재합니다.

현대의 운영체제는 이러한 `Process`를 **시분할 방식**으로 처리합니다.
이러한 `Process`들은 `PCB(Process Control Block` 와 `Program`를 통해서 만들어집니다.

### ⚙️ Process Status

`Process Status` 는 크게 4가지로 분류됩니다.
- **`Create`** : `Process` 가 만들어져서 대기상태로 갈 준비를 마친상태입니다.
- **`Ready`** : `Process`가 `CPU` 점유를 기다리는 상태입니다. `CPU` 점유는 스케줄링기법에 따라서 대기하고 있는 `Process`에게 돌아가면서 할당됩니다. 
-  **`Run`** : 스케줄링 기법에 의하여 `Process`가 `CPU` 를 점유한 상태입니다. 이 상태에서 `CPU`는 점유중인 `Process`가 요구하는 작업을 수행합니다.
- **`terminate`** : `CPU` 점유가 끝나서 `Pr ocess`가 종료되고 `PCB`가 사라진 상태

### PCB(Process Control Block)
`PCB(Process Control Block` 프로세스에 관한 모든 정보를 담고있는 자료구조입니다. 모든 프로세스는 각각 고유한 `PCB` 를 가집니다. `CPU`가`시분할 방식`으로 작업을 처리할 때 작업하고있던 프로세스의 정보를 각각의 프로세스의 `PCB`에 저장합니다. 이후에 다시 `CPU`가 해당 프로세스를 작업할 때 `PCB`에 담겨있는 정보를 바탕으로 작업을 진행합니다. 

`PCB`는 아래와같은 형태로 구성되어있습니다.
<table>
<tr>
<td>Pointer</td>
<td>Process Status</td>
</tr>
<tr>
<td colspan="2">Process Number(PID)</td>
</tr>
<tr>
<td colspan="2">Program Counter</td>
</tr>
<tr>
<td colspan="2">Register</td>
</tr>
<tr>
<td colspan="2">Memory Limits</td>
</tr>
<tr>
<td colspan="2">PPID, CPID</td>
</tr>
<tr>
<td colspan="2">...</td>
</tr>
</table>

### 📚 프로세스 연산

프로세스는 기본적으로 `Stack`, `Heap`, `Data`, `Bss`, `Text` 영역으로  나뉩니다.

- `Stack` : 함수처리를 위한 영역이며 사용자에게는 스택영역이 공개되지않습니다.
- `Heap`   : 동적인 데이터를 처리하기위한 영역입니다.
- `Data`, `Bss` : 변수들을 저장할 때 사용되는 영역이고 초기값 존재여부에 따라 달라집니다.
- `Text`: 코드의 영역입니다.




