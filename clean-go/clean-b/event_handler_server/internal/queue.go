package internal

import "errors"

type Queue struct {
}

func InitConccurencyQueue() *Queue {
	return &Queue{}
}

var queue []string

func (q *Queue) AddItem(event_id string) ([]string, error) {
	/*

	 */

	if event_id == "" {
		return nil, errors.New("'event_id' is empty")
	}

	work_queue := append(queue, event_id)

	return work_queue, nil
}

// PopItem()은 인수로 제공 받은 queue 슬라이스에서 마지막 슬라이스를 제거하고 돌려줌.
// 이때, 제거되는 이벤트의 정합성을 체크하고, 제거된 이벤트가 불일치시 에러 발생
func (q *Queue) PopItem() {}
