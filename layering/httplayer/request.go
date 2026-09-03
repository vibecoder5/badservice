package httplayer

type Request struct {
	Name        string
	Money       string
	Optional    *string
	Age         int
	Experiments map[string]string
}
