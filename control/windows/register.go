package windows

import . "monitor-server-backend/runtime/definition"

var (
	Handlers = HandlerMap{
		1001: startKeyboardRecord,
	}
)
