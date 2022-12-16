package models

// VideoPromote ...
type VideoPromote struct {
	ID          *string `json:"id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	StartDate   *string `json:"start_date"`
	EndDate     *string `json:"end_date"`
	Active      *string `json:"active"`
	Url         *string `json:"url"`
}

// VideoPromoteParameter ...
type VideoPromoteParameter struct {
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
	// VideoPromoteOrderBy ...
	VideoPromoteOrderBy = []string{"def.id", "def.title"}
	// VideoPromoteOrderByrByString ...
	VideoPromoteOrderByrByString = []string{
		"def.title",
	}

	// VideoPromoteSelectStatement ...
	VideoPromoteSelectStatement = `
	SELECT DEF.ID AS ID,
		DEF.TITLE AS TITLE,
		DEF.DESCRIPTION AS DESCRIPTION,
		DEF.START_DATE AS VideoPromote_START_DATE,
		DEF.END_DATE AS VideoPromote_END_DATE,
		DEF.ACTIVE AS ACTIVE,
		DEF.URL AS URL 
	FROM video_promote DEF`

	// VideoPromoteWhereStatement ...
	VideoPromoteWhereStatement = ` where def.id is not null `
)
