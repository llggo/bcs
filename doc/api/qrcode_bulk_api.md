## Text Data

`POST api/bulk/create?type=2`

Body:

```json
{
	"type":3,
	"number":2,
	"bulk":
	{
		"name":"Bukl 1 type 3",
		"company":"Miraway"
	},
	"qrcode":
	{
		"name":"Text 1",
		"type":"text",
		"data":{
			"content":"sssssssss"
		},
		"template":"text_big",
		"path_img":"",
		"mode":"dynamic"
	}
}

```