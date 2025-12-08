# Golang Embedding 
- 임베딩(embedding)은 한 타입이 다른 타입의 기능을 자신의 메서드 집합에 자동으로 통합하여, 마치 직접 구현한 것처럼 외부에 노출시키는 Go의 문법적 구성 요소임.
- struct embedding // interface embedding 두 종류가 있음. 
- struct embedding : 구조체를 다른 구조체 안에 포함시켜, 포함시킨 구조체의 속성을 promoted method로 외부로 노출시킴
- interface embedding : 포함된 인터페이스의 메서드 시그니처를 모두 합쳐서 새로운 인터페이스를 만듬
- Facade 패턴 구성시 자주 사용함. 
    - 여러 인터페이스를 하나의 진입점 인터페이스로 묶고 싶을 때 쓰기 좋음.
    - 인터페이스가 많아질 것을 대비하여 하나의 인터페이스에서 여러 인터페이스를 파서드 패턴으로 모아 한번에 초기화 하고 싶었음.

## 에러 발생 상황
```golang
// 유즈케이스 initializing 초기 구성(잘못된 구성)
// 인터페이스 안에는 두 가지만 들어갈 수 있음. 
// 1. method's signature
// 2. embedding된 인터페이스 명
// 초기에 사용한 것 처럼 인터페이스 안에 필드를 선언할 수는 없다.(문법 오류)
type InboundApiPort interface {
	ap APort
	bp BPort
}

type APort interface {
	Get(id int) (domain.Test, error)
}

type BPort interface {
	New(info domain.Test) (bool, error)
}
--------------------
// main.go initializing 구성 
	uc := usecase.InitUsecases(out)
	in := in.InitAdapter(uc)
--------------------
// 어댑터 initializing 구성
type RESTInAdapter struct {
	port in.InboundApiPort
}
```

- 이 경우 main.go의 InitAdapter에서 에러가 발생했음. 
- 원인은 인터페이스의 구현 문제. 
    - 어댑터 입장에서는 Get 과 New 를 구현을 기대하는데, 해당 구현을 만족하지 못하고 있었음.(잘못된 인터페이스 문법 사용으로 인한 오류)
        - Embedding 을 사용하지 않은 것으로 간주되어, 필드 기반 의존성 주입을 요구함. 
        - Get / New 필드가 구현되어 있지 않으므로 에러 발생
- 이 문제를 해결하기 위해 임베딩 사용함
```golang
// 유즈케이스 initializing 최종 구성(임베딩)
type InboundApiPort interface {
	APort  
	BPort
}
```
- 각 포트가 구현하고 있는 기능이 마치 InboundApiPort 에서 구현하는 것 처럼 해석되기 때문에, 정상적으로 동작하게 된다.

## Embedding Study
### 구현 예시

```golang

type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// ReadWriter는 Read 와 Write 가 할 수 있는 일들을 전부 할 수 있다.
// 인터페이스 안에는 인터페이스만을 임베딩 할 수 있음
type ReadWriter interface {
    Reader
    Writer
}

```

```golang

type A struct {
}

func (A) Hello() {
    fmt.Println("Hellow world")
}

type B struct {
    A                   // 임베딩은 필드명을 지정하지 않고 구조체를 바로 넣음
    name string         // 일반적인 구조체 필드 정의 형식 
}

---------

b := B{}
// 임베딩의 경우 아래 두 케이스가 모두 가능해짐.
b.Hello()
b.A.Hello() 
```

### 임베딩은 상속이 아니다.
- 승격 같은 느낌임.(네스티드 되어있는 함수를 꺼내 올린다는 느낌)
```golang
b := B{}
b.Hello() // b.A.Hello() 를 호출한 것과 같다. 단지 문법상 코드에서는 A를 생략한 것
```

### 동일 시그니쳐를 가진 함수를 정의 가능(오버라이딩이 아니다) 
- 동일 시그니쳐를 가진 함수를 정의 가능하고 호출 할 수 있지만 java 나 c++의 오버라이딩 과는 다름 
- 정적인 이름 우선 순위 룰에 따라 함수가 호출됨.

```golang
type A struct {
}

func (A) Hello() {
    fmt.Println("Hellow world")
}

type B struct {
    A                   // 임베딩은 필드명을 지정하지 않고 구조체를 바로 넣음
    name string         // 일반적인 구조체 필드 정의 형식 
}

func (B) Hello() {
    fmt.Println("Hellow B World")
}

b := B{}
b.Hello() // "Hellow B World" 가 출력됨
b.A.Hello() // "Hellow world" 가 출력됨
```


# Reference
 - https://go.dev/doc/effective_go
