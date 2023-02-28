package ports

type GRPC interface {
	Run()
	GetAddition()
	GetSubtraction()
	GetMultiplication()
	GetDivision()
}
