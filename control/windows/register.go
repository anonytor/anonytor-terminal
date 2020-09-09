package windows

import . "anonytor-terminal/runtime/definition"

var (
	Handlers = HandlerMap{
		1001: startKeyboardRecord,
	}
)
