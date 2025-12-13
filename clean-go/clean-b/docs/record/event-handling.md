# Event Handling (Like a Kafka?)

- 실무에서 이벤트 핸들링을 직접 구현하는 모듈에 대한 이야기가 나와서 개인적으로 golang으로 구현해보려함. 
- 이벤트 핸들링의 대표적인 브로커인 카프카에 대해서 구성요소와 동작 원리에 대해 알아본 뒤, 개인적으로 이벤트 핸들러를 구현해볼 예정.


## Requirements(Publisher)
- key를 사용하여 이벤트를 발행할 수 있어야 함 (ordering, partitioning)
- 메시지는 Topic 단위로 발행하며, Consumer Group에 대해 알지 않아야 함
- 브로커 연결 정보 등 정적 설정은 설정 파일에서 관리
- 동기/비동기 전송 방식 지원
- 네트워크 또는 브로커 오류 시 Retry/Backoff 지원
- 메시지 직렬화 포맷(JSON/Proto/Avro) 지원
- 필요 시 파티션 전략 선택 가능

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