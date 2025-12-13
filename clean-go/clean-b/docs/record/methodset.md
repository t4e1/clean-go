# Golang MethodSet
## 에러 발생 상황
```golang
// 수정 전 초기 코드
type Cases struct {
    *AUsecase
    *BUsecase
}

func InitUsecases(outAdapter out.OutAdapter) *Usecases {
    return &Usecases{
        AUsecase: NewAUsecase(outAdapter),
        BUsecase: NewBUsecase(outAdapter),
    }
}
```

- 컴파일 시에는 문제가 없지만 런타임 DI(의존성 주입) 과정에서 `"Usecases does not implement InboundXYZPort (missing method X)"` 같은 오류가 발생.
- 원인은 Method Set 문제.
- 임베딩을 통해 포트 인터페이스를 충족해야 하는데, 포인터 임베딩을 해서 메서드 셋 계산이 의도대로 동작하지 않음.

## 문제의 핵심: MethodSet과 Embedding의 관계
- **Go는 struct embedding이 곧바로 “상속”이 아니다**
- 임베딩은 단순히 내부 필드의 메서드를 바깥 struct의 메서드처럼 승격(promote) 시키는 문법적 기능일 뿐이다.
- 그리고 이 승격이 될지 말지는 Method Set 규칙에 좌우됨.

## Method Set 규칙(핵심 요약)
|타입|	포함되는 메서드|
|----|-------|
|T|	value receiver 메서드만 포함|
|*T|	value + pointer receiver 메서드 포함|

## 그래서 문제가 뭐였나?

- 초기 구조는 다음과 같았음:
```golang
type Cases struct {
    *AUsecase
    *BUsecase
}
```
- 여기서 중요한 점은:

    - → “Cases” 자체는 struct type(T)이다.

    - Go는 다음처럼 Method Set을 계산한다:

    - “Cases” 타입의 Method Set = value receiver 메서드만 포함

    - “*AUsecase”를 임베드했지만, “Cases”는 여전히 T (value type)이다.

    - 따라서 다음과 같은 상황이 발생한다:

```
Cases.
  - AUsecase의 value receiver 메서드? OK  
  - AUsecase의 pointer receiver 메서드? ❌ (포함되지 않음)
```

- 즉, AUsecase가 포인터 리시버로 구현한 Delete(), Update() 같은 메서드는 임베딩된 필드에 존재하지만, Cases 입장에서 method set에 들어오지 않음.

> 결과: 인터페이스를 구현해야 하는데 Delete() 메서드가 없다고 판단 → DI 실패

## 수정 후 코드
```golang
type Cases struct {
    AUsecase
    BUsecase
}

func InitUsecases(outAdapter out.OutAdapter) *Usecases {
    return &Usecases{
        AUsecase: *NewAUsecase(outAdapter),
        BUsecase: *NewBUsecase(outAdapter),
    }
}
```

- value embedding(“T를 embed”) 을 사용하기 위해 생성된 포인터를 다시 value로 만들어 넣어줌.
- 이제 Cases의 Method Set에는 다음이 포함된다:
    - AUsecase의 value 메서드  
    - AUsecase의 pointer 메서드  
    - BUsecase의 value 메서드  
    - BUsecase의 pointer 메서드


- 왜?

    - `Cases`의 임베딩 대상이 “T”라서,

    - “Cases” 자체를 포인터로 사용하면 “*Cases”의 Method Set이 사용되고, 이는 value + pointer receiver 전부 포함하기 때문.

    - 이 구조에서 인터페이스 구현에 필요한 모든 메서드가 승격되고 DI가 정상 작동한다.

## 어디를 포인터로 바꿔야 할까?

- Method Set 문제라는 걸 알면 선택지가 두 가지: 
    - 1) `InitUsecases 구조체 필드를 pointer embedding으로 유지`
    - 그러면 Cases 타입 자체를 *Cases 로만 사용해야 한다. (일반적으로는 이렇게 해도 되지만, DI 과정에서 Cases가 value로 취급되는 순간 문제가 날 수 있음.)

    - 2. `Usecase 생성자(*NewXUsecase)에서 value를 반환하도록 변경` (비추천)
        -Usecase 내부 상태가 바뀌는 구조이므로 value return은 권장되지 않음.

    - 실제로는 선택지 1을 쓰는 게 가장 자연스럽지만 DI 프레임워크나 레이어 구조 상 value embedding이 더 안전하게 동작해서 실전에서는 value 임베딩을 많이 사용한다.

## 요약
- 1. Embedding은 상속이 아니라 단순한 “승격 규칙”이다.
- 2. 이 승격이 Method Set 규칙에 의해 제한된다.
- 3. struct T와 *T의 Method Set 차이를 반드시 고려해야 한다.
- 4. 인터페이스 구현 여부는 “임베딩 후의 struct의 method set”으로 판단한다.
- 5. DI 에러는 대부분 이 규칙 때문에 발생한다.
- 임베딩 형태	Method Set 포함 여부
```
type X struct { T } (value 임베딩)	X → value only / *X → value + pointer
type X struct { *T } (pointer 임베딩)	X → value only (T의 pointer receiver는 포함 ❌)
```
