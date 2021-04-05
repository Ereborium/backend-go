package models

type Category struct {
	ID string `pg:"name,pk"`
}

type IngredientTag struct {
	IngredientID int `pg:"ingredient_id,pk"`
	TagID        int `pg:"tag_id,pk"`
}

type Ingredient struct {
	ID            int     `pg:"id,pk"`
	Name          string  `pg:"name"`
	Carbohydrates float64 `pg:"carbohydrates"`
	Proteins      float64 `pg:"proteins"`
	Fats          float64 `pg:"fats"`
}

type ItemCategory struct {
	ItemID   int    `pg:"item_id,pk"`
	Category string `pg:"category,pk"`

	CategoryRel *Category `pg:"fk:category"`
	Item        *Item     `pg:"fk:item_id"`
}

type ItemIngredient struct {

	ItemID       int  `pg:"item_id,pk"`
	IngredientID int  `pg:"ingredient_id,pk"`
	Editable     bool `pg:"editable"`

	Ingredient *Ingredient `pg:"fk:ingredient_id"`
	Item       *Item       `pg:"fk:item_id"`
}

type Item struct {
	ID          int     `pg:"id,pk"`
	Price       float64 `pg:"price,use_zero"`
	Description string  `pg:"description"`
	Image       string  `pg:"image"`
}

type Menu struct {
	ID int `pg:"id,pk"`
}

type Replacement struct {
	ItemIngredientsID int             `pg:"item_ingredients_id,pk"`
	IngredientID      int             `pg:"ingredient_id,pk"`
	AdditionalPrice   int             `pg:"additional_price,use_zero"`
	Ingredient        *Ingredient     `pg:"fk:ingredient_id"`
	ItemIngredients   *ItemIngredient `pg:"fk:item_ingredients_id"`
}

type Tag struct {
	ID   int    `pg:"id,pk"`
	Name string `pg:"name"`
}
