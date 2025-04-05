package templates

import "html/template"

var reqTemplate = "tcp \n <Method>{{.method}}</Method> \n <Path>{{.path}}</Path> \n <Body>{{.body}}</Body>"
var resTemplate = "tcp \n <Body>{{.body}}</Body>"

type TcpTemplate struct {
	data map[string]interface{}
	name string
}

func (t *TcpTemplate) GetRequestTemplate() *template.Template {
	return t.getTemplate(reqTemplate)
}

func (t *TcpTemplate) GetResponseTemplate() *template.Template {
	return t.getTemplate(resTemplate)
}

func (t *TcpTemplate) getTemplate(temp string) *template.Template {
	parsed, err := template.New(t.name).Parse(temp)
	if err != nil {
		panic(err)
	}
	return parsed
}
