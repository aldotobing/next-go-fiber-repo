package helper

func MailOTPTemplate(otp string) (res string) {

	strEmail := ` <!DOCTYPE html>
					<html>
					<head>
					<meta name="viewport" content="width=device-width, initial-scale=1">
					<style>
					body {
						font-family: Arial, Helvetica, sans-serif;
						background-color: black;
					}
					
					* {
						box-sizing: border-box;
					}
					
					/* Add padding to containers */
					.container {
						padding: 16px;
						background-color: white;
					}

					input[type=text], input[type=password] {
						width: fit-content;
						padding: 15px;
						margin: 5px 0 22px 0;
						display: inline-block;
						border: none;
						background: #f1f1f1;
					}
					</style>
					</head>
						<body>`

	strEmail += ` <div class="container">`
	strEmail += `<p> Hi, </p>`
	strEmail += `<p> We received an OTP request for your Stellar Pass Account. </p>`
	strEmail += `<p> You can enter the following One-Time Password: </p>`
	strEmail += `<p> <input style="font-size: 16px;text-align: center;background: #f1f1f1;border: none;width: 80px; height:35px" type="text" value="` + otp + `" readonly> </p>`
	strEmail += `<p>Please do not share your One-Time Password and have a good sessions at Stellar! </p>`
	strEmail += `</div>`
	strEmail += `</body>
	</html>`

	res = strEmail

	return res
}
