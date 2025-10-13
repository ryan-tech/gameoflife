package models

type GameState struct {
	Grid [][]bool			`json:"grid"`
	Rows int					`json:"rows"`
	Cols int					`json:"cols"`
	Generation int				`json:"generation"`
}

type RequestBody struct {
	Grid []Row				`json:"grid"`
}

type Row struct {
	Row int					`json:"row"`
	Cells []Cell			`json:"cells"`
}

type Cell struct {
	Col int					`json:"col"`
	Alive bool				`json:"alive"`
}