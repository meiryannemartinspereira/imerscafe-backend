package domain

type CustomerProfile struct {
	Type        CustomerType
	Description string
	Behavior    string
}

var CustomerProfiles = map[CustomerType]CustomerProfile{
	CustomerCalm: {
		Type:        CustomerCalm,
		Description: "Paciente e tranquilo",
		Behavior:    "Fala de forma calma e espera o bartender concluir o atendimento",
	},

	CustomerInRush: {
		Type:        CustomerInRush,
		Description: "Apresenta pressa e impaciência",
		Behavior:    "Cobra rapidez e demonstra preocupação com o tempo",
	},

	CustomerAesthetic: {
		Type:        CustomerAesthetic,
		Description: "Valoriza a apresentação",
		Behavior:    "Observa detalhes visuais e comenta a apresentação da bebida",
	},

	CustomerExecutive: {
		Type:        CustomerExecutive,
		Description: "Objetivo e exigente",
		Behavior:    "Valoriza eficiência, precisão e rapidez",
	},

	CustomerSpecialist: {
		Type:        CustomerSpecialist,
		Description: "Conhecedor de café",
		Behavior:    "Faz perguntas técnicas e observa detalhes da preparação",
	},

	CustomerIndecisive: {
		Type:        CustomerIndecisive,
		Description: "Tem dificuldade para decidir",
		Behavior:    "Pede sugestões, faz perguntas e pode mudar de ideia",
	},
}
