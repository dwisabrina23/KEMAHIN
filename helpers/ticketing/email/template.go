package email

const TemplateHTML = `
<!DOCTYPE html>

<html lang="en" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:v="urn:schemas-microsoft-com:vml">

<head>
	<title></title>
	<meta charset="utf-8" />
	<meta content="width=device-width, initial-scale=1.0" name="viewport" />
	<!--[if mso]><xml><o:OfficeDocumentSettings><o:PixelsPerInch>96</o:PixelsPerInch><o:AllowPNG/></o:OfficeDocumentSettings></xml><![endif]-->
	<!--[if !mso]><!-->
	<link href="https://fonts.googleapis.com/css?family=Montserrat" rel="stylesheet" type="text/css" />
	<link href="https://fonts.googleapis.com/css?family=Playfair+Display" rel="stylesheet" type="text/css" />
	<link href="https://fonts.googleapis.com/css?family=Bitter" rel="stylesheet" type="text/css" />
	<link href="https://fonts.googleapis.com/css?family=Lato" rel="stylesheet" type="text/css" />
	<link href="https://fonts.googleapis.com/css?family=Oswald" rel="stylesheet" type="text/css" />
	<link href="https://fonts.googleapis.com/css?family=Droid+Serif" rel="stylesheet" type="text/css" />
	<link href="https://fonts.googleapis.com/css?family=Oxygen" rel="stylesheet" type="text/css" />
	<link href="https://fonts.googleapis.com/css?family=Abril+Fatface" rel="stylesheet" type="text/css" />
	<link href="https://fonts.googleapis.com/css?family=Nunito" rel="stylesheet" type="text/css" />
	<!--<![endif]-->
	<style>
		* {
			box-sizing: border-box;
		}

		th.column {
			padding: 0
		}

		a[x-apple-data-detectors] {
			color: inherit !important;
			text-decoration: inherit !important;
		}

		#MessageViewBody a {
			color: inherit;
			text-decoration: none;
		}

		p {
			line-height: inherit
		}

		@media (max-width:660px) {
			.icons-inner {
				text-align: center;
			}

			.icons-inner td {
				margin: 0 auto;
			}

			.row-content {
				width: 100% !important;
			}

			.image_block img.big {
				width: auto !important;
			}

			.stack .column {
				width: 100%;
				display: block;
			}
		}
	</style>
</head>

<body style="background-color: #FFFFFF; margin: 0; padding: 0; -webkit-text-size-adjust: none; text-size-adjust: none;">
	<table border="0" cellpadding="0" cellspacing="0" class="nl-container" role="presentation"
		style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; background-color: #FFFFFF;" width="100%">
		<tbody>
			<tr>
				<td>
					<table align="center" border="0" cellpadding="0" cellspacing="0" class="row row-1"
						role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;" width="100%">
						<tbody>
							<tr>
								<td>
									<table align="center" border="0" cellpadding="0" cellspacing="0"
										class="row-content stack" role="presentation"
										style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;" width="640">
										<tbody>
											<tr>
												<th class="column"
													style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; font-weight: 400; text-align: left; vertical-align: top; padding-left: 20px; padding-right: 20px; padding-top: 60px; padding-bottom: 60px;"
													width="100%">
													<table border="0" cellpadding="0" cellspacing="0" class="text_block"
														role="presentation"
														style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;"
														width="100%">
														<tr>
															<td>
																<div style="font-family: sans-serif">
																	<div
																		style="font-size: 14px; color: #000000; line-height: 1.5; font-family: Oxygen, Trebuchet MS, Lucida Grande, Lucida Sans Unicode, Lucida Sans, Tahoma, sans-serif;">
																		<p
																			style="margin: 0; font-size: 14px; text-align: center; mso-line-height-alt: 45px;">
																			<span style="font-size:30px;"><strong>TIKET {{.Title}}</strong></span></p>
																	</div>
																</div>
															</td>
														</tr>
													</table>
													<table border="0" cellpadding="0" cellspacing="0"
														class="image_block" role="presentation"
														style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"
														width="100%">
														<tr>
															<td
																style="width:100%;padding-right:0px;padding-left:0px;padding-top:20px;">
																<div align="center" style="line-height:10px"><img
																		class="big"
																		src="{{.ImgLink}}"
																		style="display: block; height: auto; border: 0; width: 600px; max-width: 100%;"
																		width="600" /></div>
															</td>
														</tr>
													</table>
													<table border="0" cellpadding="10" cellspacing="0"
														class="divider_block" role="presentation"
														style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"
														width="100%">
														<tr>
															<td>
																<div align="center">
																	<table border="0" cellpadding="0" cellspacing="0"
																		role="presentation"
																		style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"
																		width="100%">
																		<tr>
																			<td class="divider_inner"
																				style="font-size: 1px; line-height: 1px; border-top: 1px solid #BBBBBB;">
																				<span></span></td>
																		</tr>
																	</table>
																</div>
															</td>
														</tr>
													</table>
													<table border="0" cellpadding="10" cellspacing="0"
														class="text_block" role="presentation"
														style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;"
														width="100%">
														<tr>
															<td>
																<div style="font-family: sans-serif">
																	<div
																		style="font-size: 14px; color: #555555; line-height: 1.2; font-family: Oxygen, Trebuchet MS, Lucida Grande, Lucida Sans Unicode, Lucida Sans, Tahoma, sans-serif;">
																		<p style="margin: 0; font-size: 14px;">
																			Halo,<span style="color:#150b4a;"><strong>
																					{{.Name}}!</strong></span></p>
																	</div>
																</div>
															</td>
														</tr>
													</table>
													<table border="0" cellpadding="10" cellspacing="0"
														class="text_block" role="presentation"
														style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;"
														width="100%">
														<tr>
															<td>
																<div style="font-family: sans-serif">
																	<div
																		style="font-size: 14px; color: #555555; line-height: 1.2; font-family: Oxygen, Trebuchet MS, Lucida Grande, Lucida Sans Unicode, Lucida Sans, Tahoma, sans-serif;">
																		<p style="margin: 0; font-size: 14px;">
																			Pembayaran tiket kamu sudah berhasil kami
																			terima! Berikut kami kirimkan e-ticket
																			sejumlah yang anda pesan. Terima kasih! :D
																		</p>
																	</div>
																</div>
															</td>
														</tr>
													</table>
												</th>
											</tr>
										</tbody>
									</table>
								</td>
							</tr>
						</tbody>
					</table>
					<table align="center" border="0" cellpadding="0" cellspacing="0" class="row row-2"
						role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;" width="100%">
						<tbody>
							<tr>
								<td>
									<table align="center" border="0" cellpadding="0" cellspacing="0"
										class="row-content stack" role="presentation"
										style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;" width="640">
										<tbody>
											<tr>
												<th class="column"
													style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; font-weight: 400; text-align: left; vertical-align: top; padding-top: 0px; padding-bottom: 0px;"
													width="100%">
													<table border="0" cellpadding="0" cellspacing="0"
														class="heading_block" role="presentation"
														style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"
														width="100%">
														<tr>
															<td
																style="text-align:center;width:100%;padding-bottom:20px;">
																<h2
																	style="margin: 0; color: #000000; direction: ltr; font-family: 'Oswald', Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 28px; font-weight: normal; letter-spacing: normal; line-height: 120%; text-align: left; margin-top: 0; margin-bottom: 0;">
																	<strong>??TIKET KAMU</strong></h2>
															</td>
														</tr>
													</table>
												</th>
											</tr>
										</tbody>
									</table>
								</td>
							</tr>
						</tbody>
					</table>
					<table align="center" border="0" cellpadding="0" cellspacing="0" class="row row-3"
						role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;" width="100%">
						<tbody>
							<tr>
								<td>
									<table align="center" border="0" cellpadding="0" cellspacing="0"
										class="row-content stack" role="presentation"
										style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;" width="640">
										<tbody>
											<tr>
												<th class="column"
													style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; font-weight: 400; text-align: left; vertical-align: top; background-color: #edf1f9; border-left: 0px solid #FFFFFF; border-right: 0px solid #FFFFFF;"
													width="16.666666666666668%">
													<table border="0" cellpadding="0" cellspacing="0"
														class="divider_block" role="presentation"
														style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"
														width="100%">
														<tr>
															<td>
																<div align="center">
																	<table border="0" cellpadding="0" cellspacing="0"
																		role="presentation"
																		style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"
																		width="100%">
																		<tr>
																			<td class="divider_inner"
																				style="font-size: 1px; line-height: 1px; border-top: 6px solid #5CC1ED;">
																				<span></span></td>
																		</tr>
																	</table>
																</div>
															</td>
														</tr>
													</table>
													<table border="0" cellpadding="0" cellspacing="0" class="text_block"
														role="presentation"
														style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;"
														width="100%">
														<tr>
															<td style="padding-bottom:10px;padding-top:25px;">
																<div style="font-family: sans-serif">
																	<div
																		style="font-size: 14px; font-family: Oxygen, Trebuchet MS, Lucida Grande, Lucida Sans Unicode, Lucida Sans, Tahoma, sans-serif; color: #000000; line-height: 1.2;">
																		<p
																			style="margin: 0; font-size: 34px; text-align: center;">
																			<span style="font-size:34px;">{{.DateEvent}}</span></p>
																		<p
																			style="margin: 0; font-size: 34px; text-align: center;">
																			<span style="font-size:34px;">{{.MonthEvent}}</span></p>
																	</div>
																</div>
															</td>
														</tr>
													</table>
													<table border="0" cellpadding="0" cellspacing="0" class="text_block"
														role="presentation"
														style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;"
														width="100%">
														<tr>
															<td>
																<div style="font-family: sans-serif">
																	<div
																		style="font-size: 14px; color: #5cc1ed; line-height: 1.5; font-family: Oxygen, Trebuchet MS, Lucida Grande, Lucida Sans Unicode, Lucida Sans, Tahoma, sans-serif;">
																		<p
																			style="margin: 0; font-size: 14px; text-align: center;">
																			???</p>
																	</div>
																</div>
															</td>
														</tr>
													</table>
													<table border="0" cellpadding="0" cellspacing="0" class="text_block"
														role="presentation"
														style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;"
														width="100%">
														<tr>
															<td style="padding-top:10px;padding-bottom:30px;">
																<div style="font-family: sans-serif">
																	<div
																		style="font-size: 14px; color: #000000; line-height: 1.2; font-family: Oxygen, Trebuchet MS, Lucida Grande, Lucida Sans Unicode, Lucida Sans, Tahoma, sans-serif;">
																		<p
																			style="margin: 0; font-size: 14px; text-align: center;">
																			<span style="font-size:24px;">{{.TimeEvent}}</span></p>
																	</div>
																</div>
															</td>
														</tr>
													</table>
												</th>
												<th class="column"
													style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; font-weight: 400; text-align: left; vertical-align: top; padding-left: 20px; padding-right: 20px;"
													width="83.33333333333333%">
													<table border="0" cellpadding="0" cellspacing="0"
														class="heading_block" role="presentation"
														style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"
														width="100%">
														<tr>
															<td style="width:100%;text-align:center;padding-top:20px;">
																<h1
																	style="margin: 0; color: #000000; direction: ltr; font-family: 'Oswald', Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 18px; font-weight: normal; letter-spacing: normal; line-height: 120%; text-align: left; margin-top: 0; margin-bottom: 0;">
																	{{.Title}}</h1>
															</td>
														</tr>
													</table>
													<table border="0" cellpadding="0" cellspacing="0" class="text_block"
														role="presentation"
														style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;"
														width="100%">
														<tr>
															<td style="padding-top:10px;">
																<div style="font-family: sans-serif">
																	<div
																		style="font-size: 14px; color: #000000; line-height: 1.5; font-family: Oxygen, Trebuchet MS, Lucida Grande, Lucida Sans Unicode, Lucida Sans, Tahoma, sans-serif;">
																		<p
																			style="margin: 0; font-size: 14px; mso-line-height-alt: 22.5px; letter-spacing: normal;">
																			<span style="font-size:15px;">{{.NIM}} :
																				<strong>190030551</strong></span></p>
																	</div>
																</div>
															</td>
														</tr>
													</table>
													<table border="0" cellpadding="0" cellspacing="0" class="text_block"
														role="presentation"
														style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;"
														width="100%">
														<tr>
															<td style="padding-top:10px;">
																<div style="font-family: sans-serif">
																	<div
																		style="font-size: 14px; color: #000000; line-height: 1.5; font-family: Oxygen, Trebuchet MS, Lucida Grande, Lucida Sans Unicode, Lucida Sans, Tahoma, sans-serif;">
																		<p
																			style="margin: 0; font-size: 14px; mso-line-height-alt: 22.5px; letter-spacing: normal;">
																			<span style="font-size:15px;">Kode tiket :
																				<strong>{{.TicketCode}}</strong></span>
																		</p>
																	</div>
																</div>
															</td>
														</tr>
													</table>
													<table border="0" cellpadding="0" cellspacing="0"
														class="button_block" role="presentation"
														style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"
														width="100%">
														<tr>
															<td style="padding-top:15px;text-align:left;">
																<!--[if mso]><v:roundrect xmlns:v="urn:schemas-microsoft-com:vml" xmlns:w="urn:schemas-microsoft-com:office:word" href="http://www.example.com/" style="height:46px;width:134px;v-text-anchor:middle;" arcsize="131%" stroke="false" fillcolor="#1147a1"><w:anchorlock/><v:textbox inset="0px,0px,0px,0px"><center style="color:#ffffff; font-family:Tahoma, sans-serif; font-size:16px"><![endif]--><a
																	href="{{.QRLink}}"
																	style="text-decoration:none;display:inline-block;color:#ffffff;background-color:#1147a1;border-radius:60px;width:auto;border-top:1px solid #1147a1;border-right:1px solid #1147a1;border-bottom:4px solid #073889;border-left:1px solid #1147a1;padding-top:5px;padding-bottom:5px;font-family:Oxygen, Trebuchet MS, Lucida Grande, Lucida Sans Unicode, Lucida Sans, Tahoma, sans-serif;text-align:center;mso-border-alt:none;word-break:keep-all;"
																	target="_blank"><span
																		style="padding-left:25px;padding-right:25px;font-size:16px;display:inline-block;letter-spacing:2px;"><span
																			style="font-size: 16px; line-height: 2; word-break: break-word; mso-line-height-alt: 32px;"><span
																				data-mce-style="font-size: 16px; line-height: 32px;"
																				style="font-size: 16px; line-height: 32px;">QR
																				CODE</span></span></span></a>
																<!--[if mso]></center></v:textbox></v:roundrect><![endif]-->
															</td>
														</tr>
													</table>
												</th>
											</tr>
										</tbody>
									</table>
								</td>
							</tr>
						</tbody>
					</table>
					<table align="center" border="0" cellpadding="0" cellspacing="0" class="row row-4"
						role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;" width="100%">
						<tbody>
							<tr>
								<td>
									<table align="center" border="0" cellpadding="0" cellspacing="0"
										class="row-content stack" role="presentation"
										style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;" width="640">
										<tbody>
											<tr>
												<th class="column"
													style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; font-weight: 400; text-align: left; vertical-align: top; padding-top: 0px; padding-bottom: 0px;"
													width="100%">
													<div class="spacer_block" style="height:20px;line-height:20px;">???
													</div>
												</th>
											</tr>
										</tbody>
									</table>
								</td>
							</tr>
						</tbody>
					</table>
					<table align="center" border="0" cellpadding="0" cellspacing="0" class="row row-5"
						role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;" width="100%">
						<tbody>
							<tr>
								<td>
									<table align="center" border="0" cellpadding="0" cellspacing="0"
										class="row-content stack" role="presentation"
										style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;" width="640">
										<tbody>
											<tr>
												<th class="column"
													style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; font-weight: 400; text-align: left; vertical-align: top; padding-top: 0px; padding-bottom: 0px;"
													width="100%">
													<div class="spacer_block" style="height:20px;line-height:20px;">???
													</div>
												</th>
											</tr>
										</tbody>
									</table>
								</td>
							</tr>
						</tbody>
					</table>
					<table align="center" border="0" cellpadding="0" cellspacing="0" class="row row-6"
						role="presentation" style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;" width="100%">
						<tbody>
							<tr>
								<td>
									<table align="center" border="0" cellpadding="0" cellspacing="0"
										class="row-content stack" role="presentation"
										style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;" width="640">
										<tbody>
											<tr>
												<th class="column"
													style="mso-table-lspace: 0pt; mso-table-rspace: 0pt; font-weight: 400; text-align: left; vertical-align: top; background-color: #ed740c; padding-left: 20px; padding-right: 20px; padding-top: 10px; padding-bottom: 10px;"
													width="100%">
													<table border="0" cellpadding="0" cellspacing="0"
														class="button_block" role="presentation"
														style="mso-table-lspace: 0pt; mso-table-rspace: 0pt;"
														width="100%">
														<tr>
															<td style="text-align:center;">
																<div align="center">
																	<!--[if mso]><v:roundrect xmlns:v="urn:schemas-microsoft-com:vml" xmlns:w="urn:schemas-microsoft-com:office:word" href="http://www.example.com/" style="height:41px;width:191px;v-text-anchor:middle;" arcsize="0%" stroke="false" fill="false"><w:anchorlock/><v:textbox inset="0px,0px,0px,0px"><center style="color:#ffffff; font-family:Tahoma, sans-serif; font-size:18px"><![endif]--><a
																		href="http://www.example.com/"
																		style="text-decoration:none;display:inline-block;color:#ffffff;background-color:transparent;border-radius:0px;width:auto;border-top:1px solid transparent;border-right:1px solid transparent;border-bottom:1px solid transparent;border-left:1px solid transparent;padding-top:10px;padding-bottom:10px;font-family:Oxygen, Trebuchet MS, Lucida Grande, Lucida Sans Unicode, Lucida Sans, Tahoma, sans-serif;text-align:center;mso-border-alt:none;word-break:keep-all;"
																		target="_blank"><span
																			style="padding-left:10px;padding-right:10px;font-size:18px;display:inline-block;letter-spacing:2px;"><span
																				style="font-size: 16px; line-height: 1.2; word-break: break-word; mso-line-height-alt: 19px;"><span
																					data-mce-style="font-size: 18px; line-height: 21px;"
																					style="font-size: 18px; line-height: 21px;">ASK
																					A QUESTION</span></span></span></a>
																	<!--[if mso]></center></v:textbox></v:roundrect><![endif]-->
																</div>
															</td>
														</tr>
													</table>
												</th>
											</tr>
										</tbody>
									</table>
								</td>
							</tr>
						</tbody>
					</table>
				</td>
			</tr>
		</tbody>
	</table><!-- End -->
</body>

</html>
`
