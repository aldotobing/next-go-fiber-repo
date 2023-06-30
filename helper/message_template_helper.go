package helper

import (
	"strconv"
	"strings"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/number"
	pkgtime "nextbasis-service-v-0.1/pkg/time"
)

func BuildProcessTransactionTemplate(customerOrderHeader models.CustomerOrderHeader, lineData []models.CustomerOrderLine, userData models.Customer) (res string) {
	dateString := pkgtime.GetDate(*customerOrderHeader.TransactionDate+"T00:00:00Z", "02 - 01 - 2006", "Asia/Jakarta")
	CretaedBy := ` oleh Toko : ` + *userData.CustomerName + `(` + *userData.Code + `)`
	msgbody := `*Kepada Yang Terhormat* \n\n`
	msgbody += `*` + *userData.Code + ` - ` + *userData.CustomerName + `*`
	msgbody += `\n\n*NO ORDERAN ` + *customerOrderHeader.DocumentNo + ` anda pada tanggal ` + dateString + CretaedBy + ` telah diproses*`
	msgbody += `\n\n*Berikut merupakan rincian pesanan anda:*`

	bayar, _ := strconv.ParseFloat(*customerOrderHeader.NetAmount, 0)
	harga := strings.ReplaceAll(number.FormatCurrency(bayar, "IDR", ".", "", 0), "Rp", "")
	if lineData != nil && len(lineData) > 0 {
		msgbody += `\n`
		for i := range lineData {
			msgbody += `\n ` + *lineData[i].QTY + ` ` + *lineData[i].UomName + ` ` + *lineData[i].ItemName + `\n`

		}
		ordercount := len(lineData)
		msgbody += `\n`
		msgbody += `Total ` + strconv.Itoa(ordercount) + ` item, senilai ` + harga
		msgbody += `\n`
		msgbody += `\nTerima kasih atas pemesanan anda`
		msgbody += `\n`
		msgbody += `\nSalam Sehat`
		msgbody += `\n`
		msgbody += `\nNB : Bila ini bukan transaksi dari Toko Bapak/Ibu, silahkan menghubungi Distributor Produk Sido Muncul.`
	}

	return msgbody
}

func BuildProcessSalesOrderTransactionTemplate(customerOrderHeader models.SalesOrderHeader, lineData []models.SalesOrderLine, userData models.Customer, mode int) (res string) {
	// mode 1 to customer 2 to salesman
	dateString := pkgtime.GetDate(*customerOrderHeader.TransactionDate+"T00:00:00Z", "02 - 01 - 2006", "Asia/Jakarta")

	CretaedBy := ``
	if *customerOrderHeader.DocumentNo != "" && strings.Contains(*customerOrderHeader.DocumentNo, "OSO") {
		CretaedBy += ` oleh Toko : ` + *userData.CustomerName + `(` + *userData.Code + `)`
	} else {
		CretaedBy += ` oleh Salesman : ` + *userData.CustomerSalesmanName + ` untuk Toko : ` + *userData.CustomerName + `(` + *userData.Code + `)`
	}
	msgbody := ``
	if mode == 1 {
		msgbody += `*Kepada Yang Terhormat* \n\n`
		msgbody += `*` + *userData.Code + ` - ` + *userData.CustomerName + `*`
	} else if mode == 2 {
		msgbody += `*Kepada Yang Terhormat Salesman* \n\n`
		msgbody += `*` + *userData.CustomerSalesmanName + `*`
	}

	msgbody += `\n\n*NO ORDERAN ` + *customerOrderHeader.DocumentNo + ` anda pada tanggal ` + dateString + CretaedBy + ` telah diproses*`
	msgbody += `\n\n*Berikut merupakan rincian pesanan anda:*`

	bayar, _ := strconv.ParseFloat(*customerOrderHeader.NetAmount, 0)
	harga := strings.ReplaceAll(number.FormatCurrency(bayar, "IDR", ".", "", 0), "Rp", "")
	if lineData != nil && len(lineData) > 0 {
		msgbody += `\n`
		for i := range lineData {
			msgbody += `\n ` + *lineData[i].QTY + ` ` + *lineData[i].UomName + ` ` + *lineData[i].ItemName + `\n`

		}
		ordercount := len(lineData)
		msgbody += `\n`
		msgbody += `Total ` + strconv.Itoa(ordercount) + ` item, senilai ` + harga
		msgbody += `\n`
		msgbody += `\nTerima kasih atas pemesanan anda`
		msgbody += `\n`
		msgbody += `\nSalam Sehat`
		msgbody += `\n`
		msgbody += `\nNB : Bila ini bukan transaksi dari Toko Bapak/Ibu, silahkan menghubungi Distributor Produk Sido Muncul.`
	}

	return msgbody
}

func BuildVoidTransactionTemplate(customerOrderHeader models.CustomerOrderHeader, lineData []models.CustomerOrderLine, userData models.Customer) (res string) {
	voidReasonText := ``

	if customerOrderHeader.VoidReasonText != nil {
		voidReasonText = *customerOrderHeader.VoidReasonText
	}
	dateString := pkgtime.GetDate(*customerOrderHeader.TransactionDate+"T00:00:00Z", "02 - 01 - 2006", "Asia/Jakarta")
	CretaedBy := ` oleh Toko : ` + *userData.CustomerName + `(` + *userData.Code + `)`
	msgbody := `*Kepada Yang Terhormat* \n\n`
	msgbody += `*` + *userData.Code + ` - ` + *userData.CustomerName + `*`
	msgbody += `\n\n*NO ORDERAN ` + *customerOrderHeader.DocumentNo + ` anda pada tanggal ` + dateString + CretaedBy + ` telah dibatalkan karena ` + voidReasonText + `*`
	msgbody += `\n\n*Berikut merupakan rincian pesanan anda:*`

	bayar, _ := strconv.ParseFloat(*customerOrderHeader.NetAmount, 0)
	harga := strings.ReplaceAll(number.FormatCurrency(bayar, "IDR", ".", "", 0), "Rp", "")
	if lineData != nil && len(lineData) > 0 {
		msgbody += `\n`
		for i := range lineData {
			msgbody += `\n ` + *lineData[i].QTY + ` ` + *lineData[i].UomName + ` ` + *lineData[i].ItemName + `\n`

		}
		ordercount := len(lineData)
		msgbody += `\n`
		msgbody += `Total ` + strconv.Itoa(ordercount) + ` item, senilai ` + harga
		msgbody += `\n`
		msgbody += `\nTerima kasih atas pemesanan anda`
		msgbody += `\n`
		msgbody += `\nSalam Sehat`
		msgbody += `\n`
		msgbody += `\nNB : Bila ini bukan transaksi dari Toko Bapak/Ibu, silahkan menghubungi Distributor Produk Sido Muncul.`
	}

	return msgbody
}

func BuildVoidTransactionTemplateForSalesman(customerOrderHeader models.CustomerOrderHeader, lineData []models.CustomerOrderLine, userData models.Customer, salesman models.Salesman) (res string) {
	voidReasonText := ``

	if customerOrderHeader.VoidReasonText != nil {
		voidReasonText = *customerOrderHeader.VoidReasonText
	}
	dateString := pkgtime.GetDate(*customerOrderHeader.TransactionDate+"T00:00:00Z", "02 - 01 - 2006", "Asia/Jakarta")
	CretaedBy := ` oleh Toko : ` + *userData.CustomerName + `(` + *userData.Code + `)`
	msgbody := `*Kepada Yang Terhormat Salesman* \n\n`
	msgbody += `*` + *salesman.Name + `*`
	msgbody += `\n\n*NO ORDERAN ` + *customerOrderHeader.DocumentNo + ` pada tanggal ` + dateString + CretaedBy + ` telah dibatalkan karena ` + voidReasonText + `*`
	msgbody += `\n\n*Berikut merupakan rincian pesanan anda:*`

	bayar, _ := strconv.ParseFloat(*customerOrderHeader.NetAmount, 0)
	harga := strings.ReplaceAll(number.FormatCurrency(bayar, "IDR", ".", "", 0), "Rp", "")
	if lineData != nil && len(lineData) > 0 {
		msgbody += `\n`
		for i := range lineData {
			msgbody += `\n ` + *lineData[i].QTY + ` ` + *lineData[i].UomName + ` ` + *lineData[i].ItemName + `\n`

		}
		ordercount := len(lineData)
		msgbody += `\n`
		msgbody += `Total ` + strconv.Itoa(ordercount) + ` item, senilai ` + harga
		msgbody += `\n`
		msgbody += `\nTerima kasih atas pemesanan anda`
		msgbody += `\n`
		msgbody += `\nSalam Sehat`
	}

	return msgbody
}

func BuildProcessTransactionTemplateForSalesman(customerOrderHeader models.CustomerOrderHeader, lineData []models.CustomerOrderLine, userData models.Customer, salesman models.Salesman) (res string) {
	dateString := pkgtime.GetDate(*customerOrderHeader.TransactionDate+"T00:00:00Z", "02 - 01 - 2006", "Asia/Jakarta")
	CretaedBy := ` oleh Toko : ` + *userData.CustomerName + `(` + *userData.Code + `)`
	msgbody := `*Kepada Yang Terhormat Salesman* \n\n`
	msgbody += `*` + *salesman.Name + `*`
	msgbody += `\n\n*NO ORDERAN ` + *customerOrderHeader.DocumentNo + ` pada tanggal ` + dateString + CretaedBy + ` telah diproses*`
	msgbody += `\n\n*Berikut merupakan rincian pesanan anda:*`

	bayar, _ := strconv.ParseFloat(*customerOrderHeader.NetAmount, 0)
	harga := strings.ReplaceAll(number.FormatCurrency(bayar, "IDR", ".", "", 0), "Rp", "")
	if lineData != nil && len(lineData) > 0 {
		msgbody += `\n`
		for i := range lineData {
			msgbody += `\n ` + *lineData[i].QTY + ` ` + *lineData[i].UomName + ` ` + *lineData[i].ItemName + `\n`

		}
		ordercount := len(lineData)
		msgbody += `\n`
		msgbody += `Total ` + strconv.Itoa(ordercount) + ` item, senilai ` + harga
		msgbody += `\n`
		msgbody += `\nTerima kasih atas pemesanan anda`
		msgbody += `\n`
		msgbody += `\nSalam Sehat`

	}

	return msgbody
}
