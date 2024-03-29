# 피로도 (Lv.2) 
> 완전탐색을 이용해서 푸는 문제

##### 문제설명
XX게임에는 피로도 시스템(0 이상의 정수로 표현합니다)이 있으며, 일정 피로도를 사용해서 던전을 탐험할 수 있습니다. 이때, 각 던전마다 탐험을 시작하기 위해 필요한 "최소 필요 피로도"와 던전 탐험을 마쳤을 때 소모되는 "소모 피로도"가 있습니다. "최소 필요 피로도"는 해당 던전을 탐험하기 위해 가지고 있어야 하는 최소한의 피로도를 나타내며, "소모 피로도"는 던전을 탐험한 후 소모되는 피로도를 나타냅니다. 예를 들어 "최소 필요 피로도"가 80, "소모 피로도"가 20인 던전을 탐험하기 위해서는 유저의 현재 남은 피로도는 80 이상 이어야 하며, 던전을 탐험한 후에는 피로도 20이 소모됩니다.

이 게임에는 하루에 한 번씩 탐험할 수 있는 던전이 여러개 있는데, 한 유저가 오늘 이 던전들을 최대한 많이 탐험하려 합니다. 유저의 현재 피로도 k와 각 던전별 "최소 필요 피로도", "소모 피로도"가 담긴 2차원 배열 dungeons 가 매개변수로 주어질 때, 유저가 탐험할수 있는 최대 던전 수를 return 하도록 solution 함수를 완성해주세요.

##### 제한사항

-   k는 1 이상 5,000 이하인 자연수입니다.
-   dungeons의 세로(행) 길이(즉, 던전의 개수)는 1 이상 8 이하입니다.
    -   dungeons의 가로(열) 길이는 2 입니다.
    -   dungeons의 각 행은 각 던전의 ["최소 필요 피로도", "소모 피로도"] 입니다.
    -   "최소 필요 피로도"는 항상 "소모 피로도"보다 크거나 같습니다.
    -   "최소 필요 피로도"와 "소모 피로도"는 1 이상 1,000 이하인 자연수입니다.
    -   서로 다른 던전의 ["최소 필요 피로도", "소모 피로도"]가 서로 같을 수 있습니다.
##### 입출력 예 설명

현재 피로도는 80입니다.

만약, 첫 번째 → 두 번째 → 세 번째 던전 순서로 탐험한다면

-   현재 피로도는 80이며, 첫 번째 던전을 돌기위해 필요한 "최소 필요 피로도" 또한 80이므로, 첫 번째 던전을 탐험할 수 있습니다. 첫 번째 던전의 "소모 피로도"는 20이므로, 던전을 탐험한 후 남은 피로도는 60입니다.
-   남은 피로도는 60이며, 두 번째 던전을 돌기위해 필요한 "최소 필요 피로도"는 50이므로, 두 번째 던전을 탐험할 수 있습니다. 두 번째 던전의 "소모 피로도"는 40이므로, 던전을 탐험한 후 남은 피로도는 20입니다.
-   남은 피로도는 20이며, 세 번째 던전을 돌기위해 필요한 "최소 필요 피로도"는 30입니다. 따라서 세 번째 던전은 탐험할 수 없습니다.

만약, 첫 번째 → 세 번째 → 두 번째 던전 순서로 탐험한다면

-   현재 피로도는 80이며, 첫 번째 던전을 돌기위해 필요한 "최소 필요 피로도" 또한 80이므로, 첫 번째 던전을 탐험할 수 있습니다. 첫 번째 던전의 "소모 피로도"는 20이므로, 던전을 탐험한 후 남은 피로도는 60입니다.
-   남은 피로도는 60이며, 세 번째 던전을 돌기위해 필요한 "최소 필요 피로도"는 30이므로, 세 번째 던전을 탐험할 수 있습니다. 세 번째 던전의 "소모 피로도"는 10이므로, 던전을 탐험한 후 남은 피로도는 50입니다.
-   남은 피로도는 50이며, 두 번째 던전을 돌기위해 필요한 "최소 필요 피로도"는 50이므로, 두 번째 던전을 탐험할 수 있습니다. 두 번째 던전의 "소모 피로도"는 40이므로, 던전을 탐험한 후 남은 피로도는 10입니다.

따라서 이 경우 세 던전을 모두 탐험할 수 있으며, 유저가 탐험할 수 있는 최대 던전 수는 3입니다.

-----
### 문제 풀이

비슷한 문제상황을 [DFS로 푸는 문제](https://school.programmers.co.kr/learn/courses/30/lessons/42839)를 접해봤기 때문에 비교적 쉽게 설계했습니다.

```js
function solution(k, dungeons) {
    let min = 1000;
    function dfs(p, arr){
        min = Math.min(arr.length, min);
      
        for(let i=0; i<arr.length; i++){
            let temp = [...arr];
            if(temp[i][0] <= p){
                let find = temp.splice(i, 1);
                dfs(p - find[0][1], temp);
            }
            
        }
    }
    
    dfs(k, dungeons);
    
    return dungeons.length - min;
}
```

`Javascript`에서 배열을 함수 인자로 넘길 시 깊은 복사가 안되는 이슈로 
```js 
let temp = [...arr]
```
를 통해서 배열의 값을 복사하고 이후에 인자로 넘겼습니다.

던전의 갯수가 1이상 8이하로 제공되어서 `DFS (깊이 우선 탐색` 으로 해결하였습니다.

`초기 피로도  K` 가 80일 때 입출력 예시에 나오는 던전 배열을 다음과 같습니다.
|A|B|C|
|--|--|--|
80, 20|50, 40|30, 10|

**A** 부터 던전을 돌면 초기 최소 피로도가 80 이기 때문에 가능하고 20을 소모하고 피로도는 60이 됩니다. 이후에 **B** 를 돌게되면 입장은 가능하지만 이후에 피로도 40을 소모하게되어서 최종 피로도 20으로 남은 **C** 던전의 최소조건인 30을 만족하지 못하게 되어서 해당 루트 탐색이 종료됩니다.

![DFS](https://user-images.githubusercontent.com/22852287/182833707-c6c8fffa-192f-4e4b-acef-19b46b6d29ff.png)

 탐색을 하면서 가장 깊게 들어간 깊이를 `min` 배열에 저장해 두었다가 최종적으로 던전 길이에서 빼도록 프로그래밍 했습니다.
**`끝까지 탐색하면 던전 길이 - 0 이므로 최댓값`**
> 출처:  프로그래머스 코딩 테스트 연습[https://school.programmers.co.kr/learn/challenges](https://school.programmers.co.kr/learn/challenges)