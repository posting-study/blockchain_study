# 다형성과 타입

다형성이란 동일한 메시지를 수신했을 때 객체의 타입에 따라 다르게 응답할 수 있는 능력을 의미한다. 

게임을 개발할 때를 예를 들어보자.

아래 코드에서 `init`은 해당 `State`를 초기화할 때 사용하고,
`render` 는 화면에 그래픽을 그릴 때, `Update` 는 지속적으로 실행되면서 데이터의 변경을 담당한다.
```java
abstract class GameState {
	abstract protected void init();
	abstract protected void render();
	abstract protected update();
}
```

간단한 게임이라면, 게임시작화면, 게임화면, 게임클리어화면, 게임오버화면 정도의 화면구성을 가질것이다.  각각의 화면 구성을 `GameState` 로 정의하고 `State`는 공통된 메소드 구성을 가지고있으므로 `추상클래스` 로 정의해둔다.
```java
class GameStartState extends GameState {
	@Override
	protected void init(){
		
	}
	@Override
	protected void render(){
		/*
			게임 시작화면 그리기
		*/
	}
	@Override
	protected void update(){
		/*
			시작 버튼 누르는거 감지해서 게임 구성화면으로 넘길 준비하기
		*/
	}
}
```
위 코드처럼 `GameStartState` 는 추상클래스인 `GameState` 를 상속받았기 때문에 3개의 메소드를 `Override` 받았다. 이와 같은 방식으로 계획했던 게임 진행화면, 게임 오버화면, 게임 클리어화면도 동일하게 구성할 수 있다.

이제 이 `GameState` 를 관리하는 클래스를 만들어보자.
```java
class GameStateManager {
	private GameState gameState;
	
	public GameStateManager(GameState gameState){
		this.gameState = gameState;
	}
	
	public changeGameState(GameState gameState){
		this.gameState = gameState;
	}
	
	public void init(){
		this.gameState.init();
	}
	public void render(){
		this.gameState.render();
	}
	public void update(){
		this.gameState.update();
	}
}

```

`GameStateManager`클래스를 잘 살펴보면 `gameState` 라는 데이터의 타입이 
`GameState`로 되어있다. 여기서 사용하는 모든 **상태** 들은 모두 `GameState` 라는 추상클래스를 상속받기 때문에 `GameManager` 에서 지금 상태가 어떤 상태이건간에 동일한 메소드를 호출할 수 있게됩니다.

> **구현 상속과 인터페이스 상속**
> 상속을 구현상속과 인터페이스 상속으로 분류할 수 있다. 흔히 구현 상속을 `Subclassing` 이라고 부르고 인터페이스 상속을 `Subtyping` 이라고 부른다. 순수하게 코드를 재사용하기 위한 목적으로 상속을 사용하는 것을 구현 상속이라고 부른다. 다형적인 협력을 위해 부모 클래스와 자식 클래스가 인터페이스를 공유할 수 있도록 상속을 이용하는 것을 인터페이스 상속이라고 부른다.

> 책임의 위치를 결정하기 위해 조건문을 사용하는 것은 협력의 설계 측면에서 대부분의 경우 좋지 않은 선택이다. 항상 예외 케이스를 최소화하고 일관성을 유지할 수 있는 방법을 선택하라.

