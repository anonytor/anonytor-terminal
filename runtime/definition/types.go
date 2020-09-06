package definition

type Payload struct {
	Cmd  int    `json:"cmd"`
	UUID string `json:"uuid"`
}

type HandlerFunc func()
type HandlerMap map[int]HandlerFunc
type H map[string]interface{}
