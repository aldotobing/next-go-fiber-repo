package models

// ItemPriceDetails ...
type ItemPrice struct {
	ID                 *string `json:"id"`
	ItemID             *string `json:"item_id"`
	ItemCode           *string `json:"item_code"`
	ItemName           *string `json:"item_name"`
	UomID              *string `json:"uom_id"`
	UomName            *string `json:"uom_name"`
	PriceListVersionID *string `json:"price_list_version_id"`
	Price              *string `json:"price"`
}

// ItemPriceParameter ...
type ItemPriceParameter struct {
	ID                 string `json:"id"`
	ItemID             string `json:"item_id"`
	ItemCode           string `json:"item_code"`
	ItemName           string `json:"item_name"`
	UomID              string `json:"uom_id"`
	UomName            string `json:"uom_name"`
	PriceListVersionID string `json:"price_list_version_id"`
	Search             string `json:"search"`
	Page               int    `json:"page"`
	Offset             int    `json:"offset"`
	Limit              int    `json:"limit"`
	By                 string `json:"by"`
	Sort               string `json:"sort"`
}

var (
	// ItemPriceOrderBy ...
	ItemPriceOrderBy = []string{"i.id", "i._name", "i.created_date"}
	// ItemPriceOrderByrByString ...
	ItemPriceOrderByrByString = []string{
		"i._name",
	}

	ItemPriceSelectStatement = `
	SELECT IP.ID AS ID,
		IP.ITEM_ID AS ITEM_ID,
		I.CODE AS ITEM_CODE,
		I._NAME AS ITEM_NAME,
		IP.UOM_ID AS UOM_ID,
		U._NAME AS UOM_NAME,
		IP.PRICE_LIST_VERSION_ID AS PRICE_LIST_VERSION_ID,
		IP.PRICE AS PRICE
	FROM ITEM_PRICE IP
	JOIN ITEM I ON I.ID = IP.ITEM_ID
	JOIN UOM U ON U.ID = IP.UOM_ID
	`

	// ItemPriceWhereStatement ...
	ItemPriceWhereStatement = ` WHERE i.id IS not NULL `
)
