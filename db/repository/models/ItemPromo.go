package models

// ItemPromoPromo ...
type ItemPromo struct {
	ItemID       *string `json:"item_id"`
	ItemName     *string `json:"item_name"`
	UomName      *string `json:"uom_name"`
	DiscPercent  *string `json:"disc_percent"`
	DiscAmount   *string `json:"disc_amount"`
	MinValue     *string `json:"min_value"`
	MinQty       *string `json:"min_qty"`
	CustMaxQty   *string `json:"cust_max_qty"`
	GlobalMaxQty *string `json:"global_max_qty"`
	Description  *string `json:"description"`
	StartDate    *string `json:"start_date"`
	EndDate      *string `json:"end_date"`
}

// ItemPromoParameter ...

type ItemPromoParameter struct {
	ItemID    string `json:"item_id"`
	ItemName  string `json:"item_name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Search    string `json:"search"`
	Page      int    `json:"page"`
	Offset    int    `json:"offset"`
	Limit     int    `json:"limit"`
	By        string `json:"by"`
	Sort      string `json:"sort"`
	ExceptId  string `json:"except_id"`
}

var (
	// ItemPromoOrderBy ...
	ItemPromoOrderBy = []string{"pil.item_id", "i._name", "i.created_date"}
	// ItemPromoOrderByrByString ...
	ItemPromoOrderByrByString = []string{
		"i._name",
	}

	// ItemPromoSelectStatement ...
	/*
		--UOM LINE RETURNED AS JSON
		--ALL STRING VALUE
		--ItemPromo CONVERSION > 1 (HIGHEST UOM)
	*/
	ItemPromoSelectStatement = `
	SELECT
	pil.item_id          AS item_id,
	u._name              AS uom_name,
	i._name              AS item_name,
	prl.disc_pct         AS disc_percent,
	prl.disc_amt         AS disc_amount,
	prl.minimum_value    AS minimum_value,
	prl.minimum_qty      AS minimum_qty,
	prl.customer_max_qty AS cus_max_qty,
	prl.global_max_qty   AS global_max_qty,
	prl.description      AS description,
	pr.start_date        AS start_date,
	pr.end_date          AS end_date 
FROM
	promo_item_line pil 
		LEFT JOIN promo_line prl 
		ON prl.id = pil.promo_line_id 
			LEFT JOIN promo pr 
			ON pr.id = prl.promo_id 
				LEFT JOIN item i 
				ON i.id = pil.item_id 
					LEFT JOIN uom u 
					ON u.id = pil.uom_id
	`

	// ItemPromoWhereStatement ...
	ItemPromoWhereStatement = ` WHERE i.created_date IS not NULL `
)
