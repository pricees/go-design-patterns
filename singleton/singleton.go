package singleton

type singleton struct{}

var instanceSet bool
var instance singleton

func GetInstance() singleton {
	if !instanceSet {
		instance = singleton{}
	}
	return instance
}
