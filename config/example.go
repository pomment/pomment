package config

import (
	"os"
)

type FileInfo struct {
	Path       string `json:"path"`
	Content    []byte `json:"content"`
	Permission uint32 `json:"permission"`
}

func InitExampleConfig(basePath string) (err error) {
	// Create directory
	err = os.Mkdir(basePath+"/template", os.ModePerm)
	if err != nil {
		return err
	}

	err = os.Mkdir(basePath+"/threads", os.ModePerm)
	if err != nil {
		return err
	}

	fileInfos := []FileInfo{
		{Path: "/config.toml", Content: []byte("[system]\nhost = \"0.0.0.0\"\nport = 8080\nurl = \"https://cm.tcdw.net\"\n\n[admin]\nsalt = \"example_salt\"\n\n[[admin.user]]\nname = \"admin\"\nemail = \"admin@example.com\"\npassword = \"12345678abc\"\n\n[[admin.user]]\nname = \"admin2\"\nemail = \"admin2@example.com\"\npassword = \"12345678abc\"\n\n[email]\nenabled = false\nmode = \"mailgun\"\ntitle = \"You got new reply!\"\nsender = '\"example\" <noreply@example.com>'\nsmtpHost = \"smtp.example.com\"\nsmtpPort = 587\nsmtpUsername = \"postmaster@example.com\"\nsmtpPassword = \"12345678abc\"\nmailgunAPIKey = \"12345678\"\nmailgunDomain = \"mg.example.com\"\n\n[reCAPTCHA]\nenabled = false\nsecretKey = \"12345678abc\"\nminimumScore = 0.1\n\n[push]\nenabled = false\ngateway = \"http://localhost:8081\"\nsiteName = \"example_com\"\nsiteKey = \"example_key\"\n\n[redis]\nenabled = false\nhost = \"localhost\"\nport = 6379\npassword = \"\"\ndatabase = 0\n\n[avatar]\nuseSha256 = false\n"), Permission: 0644},
		{Path: "/template/email.html", Content: []byte("<!DOCTYPE html PUBLIC \"-//W3C//DTD XHTML 1.0 Transitional//EN\" \"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd\">\n<html xmlns=\"http://www.w3.org/1999/xhtml\" lang=\"en\">\n<head>\n    <meta http-equiv=\"Content-Type\" content=\"text/html; charset=UTF-8\" />\n    <title>{{ Post.Name }} replied your post!</title>\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"/>\n</head>\n<body>\n    <p>Hey {{ ParentPost.Name }}, your post in <a href=\"{{ Thread.URL }}\">{{ Thread.Title }}</a>, got a new reply from {{ ParentPost.Name }}:</p>\n    <blockquote>{{ Post.Content }}</blockquote>\n    <p>You may want to reply the post <a href=\"{{ Thread.URL }}\">here</a>, or <a href=\"{{ UnsubscribeURL }}\">unsubscribe</a>.</p>\n</body>\n</html>\n"), Permission: 0644},
		{Path: "/template/unsubscribe.html", Content: []byte("<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, minimum-scale=1\">\n    <title>Unsubscribe Email Notification</title>\n    <style>\n        :root {\n            --color-primary: #0089A7;\n            --color-button-hover: hsl(191, 100%, 38%);\n            --color-button-active: hsl(191, 100%, 27%);\n        }\n        * {\n            font-family: -apple-system, \"lucida grande\", \"lucida sans unicode\", \"Helvetica Neue\", Tahoma, \"PingFang SC\", \"Hiragino Sans GB\", \"Source Han Sans CN Normal\", \"Heiti SC\", \"Microsoft YaHei\", sans-serif\n        }\n        body {\n            margin: 1em;\n            background-color: var(--color-primary);\n            background-image: url(\"data:image/svg+xml,%3Csvg width='200' height='200' viewBox='0 0 268 268' xmlns='http://www.w3.org/2000/svg'%3E%3Cg fill='%23ffffff' fill-opacity='0.08' fill-rule='evenodd'%3E%3Ccircle cx='37' cy='37' r='36'/%3E%3Ccircle cx='171' cy='171' r='36'/%3E%3C/g%3E%3C/svg%3E\");\n            background-attachment: fixed;\n        }\n        a, a:visited {\n            color: var(--color-primary);\n            text-decoration: none;\n        }\n        a:hover {\n            color: var(--color-button-hover);\n            text-decoration: underline;\n        }\n        .container {\n            max-width: 700px;\n            margin: 0 auto;\n            background-color: #fff;\n            border-radius: 1em;\n            box-sizing: border-box;\n            padding: 0 1em;\n            box-shadow: rgba(0, 0, 0, .12) .25em .25em 1em .25em;\n        }\n        .container__header {\n            text-align: center;\n            padding: 1em 0;\n            line-height: 2.2em;\n            box-sizing: border-box;\n            border-bottom: 2px dashed #ddd;\n        }\n        .container__header > h1 {\n            margin: 0;\n            font-size: 1.7em;\n        }\n        .container__info > p {\n            margin: 0.6em 0;\n            line-height: 1.8em;\n        }\n        .container__action {\n            display: flex;\n            justify-content: center;\n            padding: 0.5em 0 1.5em;\n        }\n        .confirm-button {\n            -webkit-appearance: none;\n            -moz-appearance: none;\n            appearance: none;\n            background-color: var(--color-primary);\n            border: 0;\n            border-radius: 1.5rem;\n            color: #fff;\n            cursor: pointer;\n            font-size: 1em;\n            height: 3rem;\n            padding: 0 1.25em;\n            transition: 0.2s;\n        }\n        .confirm-button:hover {\n            background-color: var(--color-button-hover);\n        }\n        .confirm-button:active {\n            background-color: var(--color-button-active);\n        }\n    </style>\n</head>\n<body>\n<main class=\"container\">\n    <header class=\"container__header\">\n        <h1>Unsubscribe Email Notification</h1>\n    </header>\n    <section class=\"container__info\">\n        <p>Are you sure you want to unsubscribe email notification for <a href=\"{{ Thread.URL }}#comment__{{ Thread.ID }}\">this</a> post?</p>\n        <!-- 确认取消评论邮件提醒的方式：提交 POST 请求，目标为同一个页面，带一个 FormData 参数 (userConfirmed: true) -->\n        <form action=\"\" method=\"post\" class=\"container__action\">\n            <input type=\"hidden\" name=\"userConfirmed\" value=\"true\">\n            <button class=\"confirm-button\" type=\"submit\">Yes, unsubscribe</button>\n        </form>\n    </section>\n</main>\n</body>\n</html>\n"), Permission: 0644},
		{Path: "/template/unsubscribe_error.html", Content: []byte("<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, minimum-scale=1\">\n    <title>Unsubscribe Email Notification</title>\n    <style>\n        :root {\n            --color-primary: #0089A7;\n            --color-button-hover: hsl(191, 100%, 38%);\n            --color-button-active: hsl(191, 100%, 27%);\n        }\n        * {\n            font-family: -apple-system, \"lucida grande\", \"lucida sans unicode\", \"Helvetica Neue\", Tahoma, \"PingFang SC\", \"Hiragino Sans GB\", \"Source Han Sans CN Normal\", \"Heiti SC\", \"Microsoft YaHei\", sans-serif\n        }\n        body {\n            margin: 1em;\n            background-color: var(--color-primary);\n            background-image: url(\"data:image/svg+xml,%3Csvg width='200' height='200' viewBox='0 0 268 268' xmlns='http://www.w3.org/2000/svg'%3E%3Cg fill='%23ffffff' fill-opacity='0.08' fill-rule='evenodd'%3E%3Ccircle cx='37' cy='37' r='36'/%3E%3Ccircle cx='171' cy='171' r='36'/%3E%3C/g%3E%3C/svg%3E\");\n            background-attachment: fixed;\n        }\n        a, a:visited {\n            color: var(--color-primary);\n            text-decoration: none;\n        }\n        a:hover {\n            color: var(--color-button-hover);\n            text-decoration: underline;\n        }\n        .container {\n            max-width: 700px;\n            margin: 0 auto;\n            background-color: #fff;\n            border-radius: 1em;\n            box-sizing: border-box;\n            padding: 0 1em;\n            box-shadow: rgba(0, 0, 0, .12) .25em .25em 1em .25em;\n        }\n        .container__header {\n            text-align: center;\n            padding: 1em 0;\n            line-height: 2.2em;\n            box-sizing: border-box;\n            border-bottom: 2px dashed #ddd;\n        }\n        .container__header > h1 {\n            margin: 0;\n            font-size: 1.7em;\n        }\n        .container__info {\n            padding: 1em 0;\n        }\n        .container__info > p {\n            margin: -0.4em 0;\n            line-height: 1.8em;\n        }\n    </style>\n</head>\n<body>\n<main class=\"container\">\n    <header class=\"container__header\">\n        <h1>Unsubscribe Email Notification</h1>\n    </header>\n    <section class=\"container__info\">\n        <p>Unable to perform that action.</p>\n    </section>\n</main>\n</body>\n</html>\n"), Permission: 0644},
		{Path: "/template/unsubscribe_success.html", Content: []byte("<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, minimum-scale=1\">\n    <title>Unsubscribe Email Notification</title>\n    <style>\n        :root {\n            --color-primary: #0089A7;\n            --color-button-hover: hsl(191, 100%, 38%);\n            --color-button-active: hsl(191, 100%, 27%);\n        }\n        * {\n            font-family: -apple-system, \"lucida grande\", \"lucida sans unicode\", \"Helvetica Neue\", Tahoma, \"PingFang SC\", \"Hiragino Sans GB\", \"Source Han Sans CN Normal\", \"Heiti SC\", \"Microsoft YaHei\", sans-serif\n        }\n        body {\n            margin: 1em;\n            background-color: var(--color-primary);\n            background-image: url(\"data:image/svg+xml,%3Csvg width='200' height='200' viewBox='0 0 268 268' xmlns='http://www.w3.org/2000/svg'%3E%3Cg fill='%23ffffff' fill-opacity='0.08' fill-rule='evenodd'%3E%3Ccircle cx='37' cy='37' r='36'/%3E%3Ccircle cx='171' cy='171' r='36'/%3E%3C/g%3E%3C/svg%3E\");\n            background-attachment: fixed;\n        }\n        a, a:visited {\n            color: var(--color-primary);\n            text-decoration: none;\n        }\n        a:hover {\n            color: var(--color-button-hover);\n            text-decoration: underline;\n        }\n        .container {\n            max-width: 700px;\n            margin: 0 auto;\n            background-color: #fff;\n            border-radius: 1em;\n            box-sizing: border-box;\n            padding: 0 1em;\n            box-shadow: rgba(0, 0, 0, .12) .25em .25em 1em .25em;\n        }\n        .container__header {\n            text-align: center;\n            padding: 1em 0;\n            line-height: 2.2em;\n            box-sizing: border-box;\n            border-bottom: 2px dashed #ddd;\n        }\n        .container__header > h1 {\n            margin: 0;\n            font-size: 1.7em;\n        }\n        .container__info {\n            padding: 1em 0;\n        }\n        .container__info > p {\n            margin: -0.4em 0;\n            line-height: 1.8em;\n        }\n    </style>\n</head>\n<body>\n<main class=\"container\">\n    <header class=\"container__header\">\n        <h1>Unsubscribe Email Notification</h1>\n    </header>\n    <section class=\"container__info\">\n        <p>Successfully unsubscribed email notification for <a href=\"{{ Thread.URL }}#comment__{{ Thread.ID }}\">this</a> post.</p>\n    </section>\n</main>\n</body>\n</html>\n"), Permission: 0644},
	}

	for _, fileInfo := range fileInfos {
		// 将文件内容写入到文件中
		err := os.WriteFile(basePath+fileInfo.Path, fileInfo.Content, os.FileMode(fileInfo.Permission))
		if err != nil {
			return err
		}
	}

	return nil
}