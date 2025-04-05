package templates

import "text/template"

type RequestTemplate interface {
	GetRequestTemplate() *template.Template
	GetResponseTemplate() *template.Template
}
