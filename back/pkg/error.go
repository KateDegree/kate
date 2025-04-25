package pkg

type Error struct {
	Code    int
	Message string
	Cause   error
}
