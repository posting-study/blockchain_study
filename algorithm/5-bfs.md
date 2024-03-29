# 네트워크 ( Lv.3 )
##### 문제 설명
네트워크란 컴퓨터 상호 간에 정보를 교환할 수 있도록 연결된 형태를 의미합니다. 예를 들어, 컴퓨터 A와 컴퓨터 B가 직접적으로 연결되어있고, 컴퓨터 B와 컴퓨터 C가 직접적으로 연결되어 있을 때 컴퓨터 A와 컴퓨터 C도 간접적으로 연결되어 정보를 교환할 수 있습니다. 따라서 컴퓨터 A, B, C는 모두 같은 네트워크 상에 있다고 할 수 있습니다.

컴퓨터의 개수 n, 연결에 대한 정보가 담긴 2차원 배열 computers가 매개변수로 주어질 때, 네트워크의 개수를 return 하도록 solution 함수를 작성하시오.

##### 제한사항

-   컴퓨터의 개수 n은 1 이상 200 이하인 자연수입니다.
-   각 컴퓨터는 0부터  `n-1`인 정수로 표현합니다.
-   i번 컴퓨터와 j번 컴퓨터가 연결되어 있으면 `computers[i][j]`를 1로 표현합니다.
-   `computer[i][i]`는 항상 1입니다.

##### 입출력 예
|n|computers|return|
|--|--|--|
|3|[[1, 1, 0], [1, 1, 0], [0, 0, 1]]|2|
|3|[[1, 1, 0], [1, 1, 1], [0, 1, 1]]|1
##### 입출력 예 설명

예제 #1  
아래와 같이 2개의 네트워크가 있습니다.  
![image0.png](https://grepp-programmers.s3.amazonaws.com/files/ybm/5b61d6ca97/cc1e7816-b6d7-4649-98e0-e95ea2007fd7.png)

예제 #2  
아래와 같이 1개의 네트워크가 있습니다.  
![image1.png](https://grepp-programmers.s3.amazonaws.com/files/ybm/7554746da2/edb61632-59f4-4799-9154-de9ca98c9e55.png)

### 문제풀이

예제의 2차원 배열을 좀 더 보기쉽게 정리하면 다음과 같습니다.

||A|B|C
|-|--|--|--|
|**A**|1|1|0|
|**B**|1|1|0|
|**C**|0|0|1|

이 문제에서 `DFS(깊이 우선 탐색` 을 이용하면 중복문제를 해결할 수 없다고 판단되어서 `BFS(넓이 우선 탐색)` 을 사용하기로 했습니다. 

아래부터는 네트워크, 지점을 통합해서 **`노드(Node)`** 라고 하겠습니다. 

우선, 노드의 수가 `N` 개 라고 했을 때 N의 크기를 가지는 `visit` 배열을 생성합니다. 이 배열을 `1`로 초기화하고 해당 배열의 값은 방문했을 때 `0`으로 변경하여 중복탐색을 방지합니다.

이후 1부터 `N`까지 반복문을 돌면서 
```js
if(computers[index][j]  ===  1  && visit[j]  ===  1  )
```
를 통해서 해당 노드에서 **방문한 기록이없고**, **방문가능한 노드** 를 발견하면 이를 `Queue` 에 추가합니다. 

여기서 `DFS` 와 `BFS`의 차이점이 발생하는데 `DFS`에서는 방문가능한 노드를 찾게되면 재귀함수를 통해서 최대로 방문 가능한 노드까지 끝까지 이동하게됩니다.

하지만 `BFS`에서는 방문가능한 노드를 발견하면 바로 이동하지 않고 해당 노드를 `Queue`에 넣어두었다가 이동가능한 노드가 종합이 끝나면 먼저 넣어둔 노드부터 다시 탐색을 진행하게됩니다. 

최종 코드는 다음과 같습니다.
```js
function solution(n, computers) {
    let visit = computers.map( v => 1);
    let cnt = 0;
    for(let i=0; i<n; i++){
        if(!visit[i]) continue;
        
        let que = [i];
        while(que.length){
            let index = que.shift();
            for(let j=0; j<n; j++){
                if(index === j) continue;
                
                if(computers[index][j] === 1 && visit[j] === 1 ){
                    visit[j] = 0;
                    que.push(j);
                }
            }
        }
        cnt++;
    }
    
    return cnt;
}
```
노드를 탐색하는 과정에서 큐가 비게되면 해당 노드가 더이상 이어질 곳이 없다는 의미이므로 이어져있는 노드를 실셈하는 `cnt` 변수의 값을 증가시켜줍니다.

 

