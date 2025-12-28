# Event Handling (Like a Kafka?)

- 실무에서 이벤트 핸들링을 직접 구현하는 모듈에 대한 이야기가 나와서 개인적으로 golang으로 구현해보려함. 
- 이벤트 핸들링의 대표적인 브로커인 카프카에 대해서 구성요소와 동작 원리에 대해 알아본 뒤, msa 서비스간 간단한 이벤트를 주고받을 수 있는 이벤트 핸들러를 구현해볼 예정.


## Requirements(Publisher)
- key를 사용하여 이벤트를 발행할 수 있어야 함 (ordering, partitioning)
- 메시지는 Topic 단위로 발행하며, Consumer Group에 대해 알지 않아야 함
- 브로커 연결 정보 등 정적 설정은 설정 파일에서 관리
- 동기/비동기 전송 방식 지원
- 네트워크 또는 브로커 오류 시 Retry/Backoff 지원
- 메시지 직렬화 포맷(JSON/Proto/Avro) 지원

## Requirements(Subscriber(Consumer))
- 구독 대상 Topic에 새로운 메시지가 존재할 때 이를 Polling 방식으로 가져와 처리
- Consumer Group 단위로 메시지 분배 및 파티션 독점
- 메시지 처리 후 offset 커밋 방식 선택(auto/manual) 
- 메시지 처리 후 commit log를 등록할 수 있어야함.
- 실패 시 재시도 또는 DLQ 저장 기능
- 멱등성(idempotency) 보장 가능
- 처리 병렬성 및 스케일아웃 가능
- 메시지 역직렬화 포맷 호환

## Requirements(Broker)
- Publisher에게서 받은 메시지를 Topic/Partition에 저장
- Subscriber가 Polling 방식으로 메시지를 가져가도록 제공
- Consumer Group 기반으로 파티션 단위 메시지 순서를 보장
- 서로 다른 Consumer Group들은 독립적으로 메시지를 소비 가능
- 메시지 retention 정책 설정(기간/용량 기준)
- 전달 보장 모델 설정(at-least-once, exactly-once)
- 장애 시 빠른 복구 및 commit log 기반 내구성 보장

- (고가용성을 위한 replication, leader election 지원)
- (스케일아웃 지원(파티션 확장))

## Architecture
```golang
# total architecture 
Publisher -> Broker Server <- Consumer 의 pub/sub 구조 채용
```
- publisher -> Broker : gRPC로 이벤트 발행 
- Consumer -> Broker : gRPC Streaming 으로 구독한 이벤트 등록시 poll
- Broker는 (메모리/로컬파일/DB저장소)를 기반으로 log 쌓기(server fail 발생시 복구 보장)
    -> 어디까지 처리하였는지 각 이벤트별 리퀘스트 id 필요할 것
    -> 폴링 적용시 롱 폴링 생각해볼것.

```golang 
# serveral architecture(pub -> broker)
// pub: 발행의 주체. 발행만 하면 이후는 신경 안씀. 
// Event catalog파일을 참조하여 이벤트 관련 정보 정의(이벤트명, 타입(순서중시/속도중시), consumer group id, consumer group order id(순서를 맞추기 위해) 등)
// 정의한 파일을 읽어와서 이벤트 발행함.

// broker: 이벤트 관리의 주체
// Event에 대한 상세 정보(Event Catalog)은 신경쓰지 않고 문자(혹은 객체)로만 이벤트 관리
// server failure 발생시 복구를 위해 logging 필요(메모리?로컬텍스트?DB?)
// 서버 관련 설정(워커 몇개나 사용할 것인지...등등)은 일단은 코드상 상수로 관리 예정 (향후 서버 config file 생성 가능성 있음)

(pub) event 발생 -> (broker) req-id 생성 & 로깅 시작 -> (br) 이벤트 타입 파악 -> (br) 타입별 큐 배정(FIFO/goroutine) -> (br) polling 확인 시 큐에서 이벤트 삭제 
```

```golang
# serveral architecture(broker <-> sub)
// sub: 이벤트 수신측
// Long Polling 기법을 사용해서 broker 측에 일정 시간 단위로 리퀘스트를 보냄.
// 대기 중인 큐에 작업이 들어오면 해당 작업을 가져가고 ack 반환. 
// broker 는 ack 을 받으면 큐에서 작업을 삭제
// 만약 대기 중에 큐에 작업이 들어오지 않는다면 재요청

(sub) polling 요청 보냄 -> (br) 타입별 큐 배정 -> (sub) 작업에 들어온 큐를 확인하고 가져감 -> (sub) 브로커에게 ack 발행 -> (br) 큐에서 이벤트 삭제 
(sbu) polling 요청 보냄 -> 대기 중 작업 배정 없음 -> (sub) 재요청
```

## Attributes
아키텍쳐를 구현하기 위해 입력받아야 하는 값들?

```golang
"""
pub/sub 설정 파일 

event-id/ event-type/ consumer-group-ip/ consumer-group-order-id/ 폴링 시간/ 


broekr 설정 

worker 수/ 로깅 정책(사용 유무, 로그 유지 시간)/ 
"""
```

