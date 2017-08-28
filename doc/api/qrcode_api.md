## Text Data

`POST api/qrcode_api/create`

Body:

```json
{
	"user_id":"usr_m6IAkX3PovtVuVgKDBn7",
	"name":"Text 1",
	"type":"text",
	"data":{
		"content":"sssssssss"
	},
	"template":"text_big",
	"path_img":"",
	"mode":"dynamic"
}
```

## Url Data

`POST api/qrcode_api/create`

Body:

```json
{
	"user_id":"usr_m6IAkX3PovtVuVgKDBn7",
	"name":"Text 1",
	"type":"url",
	"data":{
		"url": "http://tinhte.vn"
	},
	"template":"",
	"path_img":"",
	"mode":"dynamic"
}
```

## Urls Data

`POST api/qrcode_api/create`

Body:

```json
{
	"user_id":"usr_m6IAkX3PovtVuVgKDBn7",
	"name":"Text 1",
	"type":"urls",
	"data":{
		"url":{
			"name":"Tinh Te",
			"url": "http://tinhte.vn"
		},
		"url2":{
			"name":"FB",
			"url": "http://fb.com"
		}
	},
	"template":"default",
	"path_img":"",
	"mode":"dynamic"
}
```

## Social Data

`POST api/qrcode_api/create`

Body:

```json
{
	"user_id":"usr_m6IAkX3PovtVuVgKDBn7",
	"name":"Text 1",
	"type":"social",
	"data":{
		"facebook":{
			"name":"Face Book",
			"url": "http://fb.com"
		},
		"youtube":{
			"name":"Face Book",
			"url": "http://fb.com"
		}
	},
	"template":"fb_template",
	"path_img":"",
	"mode":"dynamic"
}
```

## Image Data

`POST api/qrcode_api/create`

Body:

```json
{
	"user_id":"usr_m6IAkX3PovtVuVgKDBn7",
	"name":"Text 1",
	"type":"image",
	"data":{
		"name":"image example",
		"path":"https://dummyimage.com/600x400/000/fff"
	},
	"template":"image_responsive",
	"path_img":"",
	"mode":"dynamic"
}
```

## PDF Data

`POST api/qrcode_api/create`

Body:

```json
{
	"user_id":"usr_m6IAkX3PovtVuVgKDBn7",
	"name":"Text 1",
	"type":"pdf",
	"data":{
		"name":"pdf example",
		"path":"http://www.pdf995.com/samples/pdf.pdf"
	},
	"template":"pdf_template",
	"path_img":"",
	"mode":"dynamic"
}
```

## Audio Data

`POST api/qrcode_api/create`

Body:

```json
{
	"user_id":"usr_m6IAkX3PovtVuVgKDBn7",
	"name":"Text 1",
	"type":"audio",
	"data":{
		"name":"image example",
		"path":"/static/upload/audio/sl.mp3"
	},
	"template":"mp3",
	"path_img":"",
	"mode":"dynamic"
}
```