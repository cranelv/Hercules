package network
type PeerInterface interface {

}
type PeerFilterInterface interface {
	Filter([]PeerInterface) []PeerInterface
}