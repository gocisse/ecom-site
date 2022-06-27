package models

import "database/sql"

type DBModels struct {
	DB *sql.DB
}

type Models struct {
	DB DBModels
}

func NewModel(db *sql.DB) Models {
	return Models{
		DB: DBModels{
			DB: db,
		},
	}
}


type Widgets struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	Description string `json:"description"`
	Image string `json:"image"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
}

func(db *DBModels)GetWidgetByID(id int)( *Widgets, error) {
	var widget Widgets
	row := db.DB.QueryRow("SELECT * FROM widgets WHERE id = ?", id)
	err := row.Scan(&widget.ID, 
		&widget.Name, 
		// &widget.Price, 
		// &widget.Description, 
		// &widget.CreatedAt, 
		// &widget.UpdatedAt,
	)

	if err != nil {
		return &widget, err
	}
	return &widget,nil

}
