package models

// News ...
type News struct {
	ID          *string `json:"id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	StartDate   *string `json:"start_date"`
	EndDate     *string `json:"end_date"`
	Active      *string `json:"active"`
	ImageUrl    *string `json:"image_url"`
	Priority    string  `json:"priority"`
}

// NewsParameter ...
type NewsParameter struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	UserId    string `json:"user_id"`
	Search    string `json:"search"`
	Page      int    `json:"page"`
	Offset    int    `json:"offset"`
	Limit     int    `json:"limit"`
	By        string `json:"by"`
	Sort      string `json:"sort"`
}

var (
	// NewsOrderBy ...
	NewsOrderBy = []string{"def.id", "def.title", "def.priority"}
	// NewsOrderByrByString ...
	NewsOrderByrByString = []string{
		"def.title",
	}

	// NewsSelectStatement ...
	NewsSelectStatement = `
	SELECT DEF.ID AS ID,
		DEF.TITLE AS TITLE,
		DEF.DESCRIPTION AS DESCRIPTION,
		DEF.START_DATE AS NEWS_START_DATE,
		DEF.END_DATE AS NEWS_END_DATE,
<<<<<<< HEAD
		DEF.URL AS URL_NEWS 
=======
		DEF.IMAGE_URL,
		DEF.ACTIVE,
		DEF.PRIORITY
>>>>>>> 3c9340e87a37921d12f3bd4efe63b6a83345787c
	FROM NEWS DEF`

	// NewsWhereStatement ...
	NewsWhereStatement = ` where def.id is not null and def.active = 1 `
)
