package domain

type CustomerType string

const (
	CustomerCalm       CustomerType = "CALMO"
	CustomerInRush     CustomerType = "APRESSADO"
	CustomerAesthetic  CustomerType = "ESTETICO"
	CustomerExecutive  CustomerType = "EXECUTIVO"
	CustomerSpecialist CustomerType = "ESPECIALISTA"
	CustomerIndecisive CustomerType = "INDECISO"
)
