package entity

type Entity interface {
	SKey() string
}

var (
	_ Entity = (*User)(nil)
	_ Entity = (*Token)(nil)
	_ Entity = (*Article)(nil)
)
