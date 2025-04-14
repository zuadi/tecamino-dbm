package models

type Rights string

const (
	Read      Rights = "R"
	Write     Rights = "W"
	ReadWrite Rights = "RW"
)

func (r *Rights) GetRights() Rights {
	if r == nil {
		return ReadWrite
	} else if *r == "" {
		return ReadWrite
	}
	return *r
}
