package data

type Data interface {
	GetMaxId() Id
	GetNextId(param IdCreateParam) Id
}
