// Code generated by "stringer -type TaskEvent"; DO NOT EDIT.

package hub

import "strconv"

const _TaskEvent_name = "NewEventUpdateEventCancelEvent"

var _TaskEvent_index = [...]uint8{0, 8, 19, 30}

func (i TaskEvent) String() string {
	if i < 0 || i >= TaskEvent(len(_TaskEvent_index)-1) {
		return "TaskEvent(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TaskEvent_name[_TaskEvent_index[i]:_TaskEvent_index[i+1]]
}
