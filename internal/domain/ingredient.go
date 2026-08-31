package domain

type Ingredient struct {
	ID   string
	Name string
}

var DefaultIngredients = []Ingredient{
	{ID: "1", Name: "CAFE"},
	{ID: "2", Name: "LEITE"},
	{ID: "3", Name: "CHOCOLATE"},
	{ID: "4", Name: "ESPUMA"},
	{ID: "5", Name: "CHANTILLY"},
	{ID: "6", Name: "CARAMELO"},
	{ID: "7", Name: "CANELA"},
	{ID: "8", Name: "ACUCAR"},
	{ID: "9", Name: "BAUNILHA"},
	{ID: "10", Name: "CREME"},
	{ID: "11", Name: "AGUA"},
	{ID: "12", Name: "GELO"},
}
