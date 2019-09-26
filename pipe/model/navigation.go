package model


type Navigation struct {
	Model
	Ttile string
	URL string
	IconURL string
	OpenMethod string
	Number int

	BlogID uint64

}

// Navigation open methods.
const (
	NavigationOpenMethodBlank = "_blank"
	NavigationOpenMethodSelf  = "_self"
)