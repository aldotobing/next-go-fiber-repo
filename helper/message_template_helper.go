package helper

import (
	"strconv"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/number"
)

func BuildProcessTransactionTemplate(customerOrderHeader models.CustomerOrderHeader, lineData []models.CustomerOrderLine, userData models.UserAccount) (res string) {

	msgbody := `Kepada Yang Terhormat ` + *userData.Name + `\n\nCheckout anda dengan nomor ` + *customerOrderHeader.DocumentNo
	msgbody += ` sedang dalam proses`
	msgbody += `\n\nBerikut merupakan rincian pesanan anda:`

	bayar, _ := strconv.ParseFloat(*customerOrderHeader.NetAmount, 0)
	harga := strings.ReplaceAll(number.FormatCurrency(bayar, "IDR", ".", "", 0), "Rp", "")
	if lineData != nil && len(lineData) > 0 {
		msgbody += `\n`
		for i := range lineData {
			msgbody += `\n ` + *lineData[i].QTY + ` ` + *lineData[i].UomName + ` ` + *lineData[i].ItemName + `\n`

		}
		ordercount := len(lineData)
		msgbody += `\n`
		msgbody += `Total ` + strconv.Itoa(ordercount) + ` item, senilai ` + harga + ` (belum termasuk potongan/diskon bila ada program potongan/diskon) `
		msgbody += `\n`
		msgbody += `\nTerima kasih atas pemesanan anda`
		msgbody += `\n`
		msgbody += `\nSalam Sehat`
		msgbody += `\n`
		msgbody += `\nAutogenerate Whatsapp`
	}

	return msgbody
}

func BuildVoidTransactionTemplate(customerOrderHeader models.CustomerOrderHeader, lineData []models.CustomerOrderLine, userData models.UserAccount) (res string) {

	msgbody := `Kepada Yang Terhormat ` + *userData.Name + `\n\nCheckout anda dengan nomor ` + *customerOrderHeader.DocumentNo
	msgbody += ` telah dibatalkan`
	msgbody += `\n\nBerikut merupakan rincian pesanan anda:`

	bayar, _ := strconv.ParseFloat(*customerOrderHeader.NetAmount, 0)
	harga := strings.ReplaceAll(number.FormatCurrency(bayar, "IDR", ".", "", 0), "Rp", "")
	if lineData != nil && len(lineData) > 0 {
		msgbody += `\n`
		for i := range lineData {
			msgbody += `\n ` + *lineData[i].QTY + ` ` + *lineData[i].UomName + ` ` + *lineData[i].ItemName + `\n`

		}
		ordercount := len(lineData)
		msgbody += `\n`
		msgbody += `Total ` + strconv.Itoa(ordercount) + ` item, senilai ` + harga + ` (belum termasuk potongan/diskon bila ada program potongan/diskon) `
		msgbody += `\n`
		msgbody += `\nTerima kasih atas pemesanan anda`
		msgbody += `\n`
		msgbody += `\nSalam Sehat`
		msgbody += `\n`
		msgbody += `\nAutogenerate Whatsapp`
	}

	return msgbody
}
