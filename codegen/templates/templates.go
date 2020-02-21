package templates

// THIS IS A GENERATED FILE. DO NOT MODIFY

import (
	"os"
	"path"
	"text/template"
)

const (
	// Config name of template from file config.tmpl
	Config = "config.tmpl"
	// Controller name of template from file controller.tmpl
	Controller = "controller.tmpl"
	// Docs name of template from file docs.tmpl
	Docs = "docs.tmpl"
	// Gitignore name of template from file gitignore.tmpl
	Gitignore = "gitignore.tmpl"
	// Main name of template from file main.tmpl
	Main = "main.tmpl"
	// Models name of template from file models.tmpl
	Models = "models.tmpl"
	// TestConfig name of template from file test_config.tmpl
	TestConfig = "test_config.tmpl"
)

// Template for creating a structured file given a context and a path.
type Template struct {
	// Name of the template within the template collection.
	Name string
}

// GetName of this template
func (t *Template) GetName() string {
	return t.Name
}

// Execute this template writing the output data to the passed outputPath which will be joined
// using path.
func (t *Template) Execute(context interface{}, fileMode os.FileMode, outputPath ...string) error {
	outputFileName := path.Join(outputPath...)
	fd, err := os.OpenFile(outputFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.FileMode(fileMode))
	if nil != err {
		return err
	}
	defer fd.Close()
	templates := GetTemplates()
	return templates.ExecuteTemplate(fd, t.Name, context)
}

var (
	// ConfigTemplate from file config.tmpl
	ConfigTemplate = &Template{Name: Config }
	// ControllerTemplate from file controller.tmpl
	ControllerTemplate = &Template{Name: Controller }
	// DocsTemplate from file docs.tmpl
	DocsTemplate = &Template{Name: Docs }
	// GitignoreTemplate from file gitignore.tmpl
	GitignoreTemplate = &Template{Name: Gitignore }
	// MainTemplate from file main.tmpl
	MainTemplate = &Template{Name: Main }
	// ModelsTemplate from file models.tmpl
	ModelsTemplate = &Template{Name: Models }
	// TestConfigTemplate from file test_config.tmpl
	TestConfigTemplate = &Template{Name: TestConfig }
)

// GetTemplates returns a template that has the all the other templates parsed into it accessible via their filename.
func GetTemplates() *template.Template {
    master := template.New("templatesTemplate")
    
    // Config
    template.Must(master.New(Config).Parse(string("package config\n\n// THIS IS A GENERATED FILE. DO NOT MODIFY\n// config.tmpl\n\nimport (\n\t\"github.com/beaconsoftwarellc/gadget/environment\"\n\t\"github.com/beaconsoftwarellc/gadget/log\"\n)\n\n// Specification details the expected values for the config\ntype Specification struct {\n  Log log.Logger\n  {{ range $env := .Specification.Environment}}\n  {{ $env.Name }} {{ $env.Type }}{{if $env.Env }} `env:\"{{$env.Env}}{{if $env.Optional}},optional{{end}}\"{{if $env.S3}} s3:\"{{$env.S3}}\"{{end}}`{{end}}{{end}}\n  {{ range $service := .Specification.Services }}\n  {{$service.Name}} {{$service.Type}}{{end}}\n}\n\n// New returns a Specification based on the environment\nfunc New() *Specification {\n\treturn NewValues(environment.GetEnvMap())\n}\n\n// NewValues returns a Specification based on the env var map passed in\nfunc NewValues(envVars map[string]string) *Specification {\n\ts := &Specification{ {{ range $env := .Specification.Environment }}\n\t\t{{ if and $env.Optional $env.Default }}{{$env.Name}}: {{$env.Default}},{{end}}{{end}}\n\t}\n\terr := environment.ProcessMap(s, envVars)\n\tif nil != err {\n\t\tpanic(log.Error(err))\n\t}\n\n\ts.Log = log.New(\"{{.Name}}\", log.FunctionFromEnv())\n\n\t{{ range $service := .Specification.Services }}\n  \ts.{{$service.Name}} = {{$service.Initializer}}{{end}}\n\n\treturn s\n}\n")))
    
    // Controller
    template.Must(master.New(Controller).Parse(string("package controllers\n\n// THIS IS A GENERATED FILE. DO NOT MODIFY\n// controller.tmpl\n\nimport (\n    \"fmt\"\n\tqcontrollers \"github.com/beaconsoftwarellc/quimby/controllers\"\n\tqerror \"github.com/beaconsoftwarellc/quimby/error\"\n\tqhttp \"github.com/beaconsoftwarellc/quimby/http\"\n\t\"github.com/beaconsoftwarellc/gadget/errors\"\n\t\"net/http\"\n\n\t\"{{.BasePath}}/config\"\n\t\"{{.BasePath}}/models\"\n)\n\n{{ range $index, $controller := .Controllers }}\n// {{$controller.Name}}Controller {{$controller.Description}}\ntype {{$controller.Name}}Controller interface {\n\tqhttp.Controller\n\t{{ range .Actions }}do{{.Method}}(context *qhttp.Context, {{.ArgumentDefinition $controller}}) {{.ReturnDefinition}}\n\t{{end}}\n}\n\ntype {{$controller.ImplName}}Controller struct {\n\t{{$controller.Security}}\n\tSpecification *config.Specification\n}\n\n// New{{$controller.Name}}Controller returns an initialized {{$controller.Name}}Controller\nfunc New{{$controller.Name}}Controller(spec *config.Specification) {{$controller.Name}}Controller {\n\tcontroller := &{{$controller.ImplName}}Controller{}\n\tcontroller.Specification = spec\n\t{{if $controller.Auth}}controller.Validator = {{$controller.Auth.Validator}}{{end}}\n\treturn controller\n}\n\n// GetRoutes establishes routes for the {{$controller.Name}}Controller\nfunc (controller *{{$controller.ImplName}}Controller) GetRoutes() []string {\n\treturn []string{\n        {{ range $i, $route := $controller.Routes }}\"{{$route}}\",\n        {{end}}\n\t}\n}\n\n{{ range .Actions }}// {{.Method}} {{.Description}}\nfunc (controller *{{$controller.ImplName}}Controller) {{.Method}}(context *qhttp.Context) {\n\t{{ $controller.LoadAuth -}}\n\t{{ .ReadRequestModel -}}\n\t{{ .ReadQueryModel }}\n\t{{ if .Return }}resp, err := controller.do{{.Method}}(context, {{.Arguments $controller}})\n\tif nil != err {\n\t\tqerror.TranslateError(context, err)\n\t\t{{.ErrorLog}}\n\t\treturn\n\t}\n\tcontext.SetResponse(resp, {{.Status}}){{else}}\n\terr := controller.do{{.Method}}(context, {{.Arguments $controller}})\n\tif nil != err {\n\t\tqerror.TranslateError(context, err)\n\t\t{{.ErrorLog}}\n\t\treturn\n\t}\n\tcontext.SetResponse(nil, {{.Status}}){{end}}\n}\n{{end}}\n\n{{ end }}\n")))
    
    // Docs
    template.Must(master.New(Docs).Parse(string("package controllers\n\n// THIS IS A GENERATED FILE. DO NOT MODIFY\n// docs.tmpl\n\nimport (\n\t\"net/http\"\n\n\tqcontrollers \"github.com/beaconsoftwarellc/quimby/controllers\"\n\t\"{{.BasePath}}/config\"\n\t\"{{.BasePath}}/security\"\n\tqhttp \"github.com/beaconsoftwarellc/quimby/http\"\n)\n\n// DocController renders the autogenerated api documentation\ntype DocController struct {\n\tqcontrollers.MethodNotAllowedController\n\tqcontrollers.BasicAuthenticatedController\n\tSpecification *config.Specification\n}\n\n// NewDocController returns an initialized DocController\nfunc NewDocController(spec *config.Specification) *DocController {\n\tcontroller := &DocController{Specification: spec}\n\tcontroller.Validator = security.NewBasicValidator()\n\n\treturn controller\n}\n\n// GetRoutes establishes routes for the RegistrationsController\nfunc (controller *DocController) GetRoutes() []string {\n\treturn []string{\n\t\t\"docs\",\n\t}\n}\n\n// Get renders the documentation HTML\nfunc (controller *DocController) Get(context *qhttp.Context) {\n\tcontext.Response.Header().Add(\"Content-Type\", \"text/html\")\n\tcontext.SetResponse(htmlDocs, http.StatusOK)\n}\n\nconst htmlDocs = `<!DOCTYPE html>\n<html lang=\"en\">\n  <head>\n\t<!-- Required meta tags -->\n    <meta charset=\"utf-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, shrink-to-fit=no\">\n\n    <!-- Bootstrap CSS -->\n\t\t<link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css\" integrity=\"sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u\" crossorigin=\"anonymous\">\n  </head>\n  <body>\n    <div class=\"container\">\n      <h2>Endpoints</h2>\n\t\t{{ $definition := . }}\n\t  {{ range $index, $controller := .Controllers }}\n\t  {{ range .Actions }}\n\t  <div class=\"panel panel-{{ .PanelCode }}\">\n\t\t  <div class=\"panel-heading\">\n\t\t\t<div class=\"container\">\n\t\t\t  <div class=\"col-sm-1\"><b>{{ .Method }}</b></div>\n\t\t\t  <div class=\"col-sm-10\">/{{ $controller.DefaultRoute }}</div>\n\t\t\t  <div class=\"col-sm-1\"><span data-toggle=\"collapse\" data-target=\"#{{$controller.Name}}{{.Method}}Info\"><span class=\"glyphicon glyphicon-{{if $controller.HasAuth}}lock{{else}}unchecked{{end}}\"></span></span></div>\n\t\t\t</div><!-- .container -->\n\t\t  </div>\n\t\t  <div class=\"panel-body collapse\" id=\"{{$controller.Name}}{{.Method}}Info\">\n\t\t\t  <p>{{ .Description }}</p>\n\t\t\t  <h5>Headers</h5>\n\t\t\t  <table class=\"table\">\n\t\t\t\t<thead>\n\t\t\t\t  <tr>\n\t\t\t\t    <th class=\"col-sm-3\">Name</th>\n\t\t\t\t    <th class=\"col-sm-9\">Description</th>\n\t\t\t\t  </tr>\n\t\t\t\t</thead>\n\t\t\t\t<tbody>{{ range $controller.AuthHeaders }}\n\t\t\t\t  <tr>\n\t\t\t\t    <td class=\"col-sm-3\">{{.Name}}</td>\n\t\t\t\t    <td class=\"col-sm-9\">{{.Type}}</td>\n\t\t\t\t  </tr>{{ end }}\n\t\t\t\t</tbody>\n\t\t\t  </table>\n\t\t\t\t{{ if $controller.Parameters -}}\n\t\t\t  <h5>URI Parameters</h5>\n\t\t\t  <table class=\"table\">\n\t\t\t\t<thead>\n\t\t\t\t  <tr>\n\t\t\t\t    <th class=\"col-sm-3\">Name</th>\n\t\t\t\t    <th class=\"col-sm-9\">Description</th>\n\t\t\t\t  </tr>\n\t\t\t\t</thead>\n\t\t\t\t<tbody>{{ range $controller.Parameters }}\n\t\t\t\t  <tr>\n\t\t\t\t    <td class=\"col-sm-3\">{{ .Name }}</td>\n\t\t\t\t    <td class=\"col-sm-9\">{{ .Type }}</td>\n\t\t\t\t  </tr>{{ end }}</tbody></table>{{ end }}\n\t\t\t\t{{ if .Parameters -}}\n\t\t\t  <h5>Query Parameters</h5>\n\t\t\t  <table class=\"table\">\n\t\t\t\t<thead>\n\t\t\t\t  <tr>\n\t\t\t\t    <th class=\"col-sm-3\">Name</th>\n\t\t\t\t    <th class=\"col-sm-9\">Description</th>\n\t\t\t\t  </tr>\n\t\t\t\t</thead>\n\t\t\t\t  {{ range .QueryParameters $definition }}\n  \t\t\t\t  <tr>\n  \t\t\t\t    <td class=\"col-sm-3\">{{ .QueryName }}</td>\n  \t\t\t\t    <td class=\"col-sm-9\">{{ .QueryType }}</td>\n  \t\t\t\t  </tr>{{ end }}</tbody></table>{{ end }}\n\t\t\t\t{{ if .Model -}}\n\t\t\t\t<h5>Body</h5>\n\t\t\t  <table class=\"table\">\n\t\t\t\t<thead>\n\t\t\t\t  <tr>\n\t\t\t\t    <th class=\"col-sm-3\">Name</th>\n\t\t\t\t    <th class=\"col-sm-9\"></th>\n\t\t\t\t  </tr>\n\t\t\t\t</thead>\n\t\t\t\t<tbody>\n\t\t\t\t\t<tr>\n  \t\t\t\t    <td class=\"col-sm-3\"><a href=\"#{{.Model}}Model\">{{ .Model }}</a></td>\n  \t\t\t\t    <td class=\"col-sm-9\"></td>\n\t\t\t\t\t</tr>\n\t\t\t\t</tbody>\n\t\t\t  </table>{{ end }}\n\t\t\t  <h5>Response</h5>\n\t\t\t  <table class=\"table\">\n\t\t\t\t<thead>\n\t\t\t\t  <tr>\n\t\t\t\t    <th class=\"col-sm-3\">Code</th>\n\t\t\t\t    <th class=\"col-sm-9\">Model</th>\n\t\t\t\t  </tr>\n\t\t\t\t</thead>\n\t\t\t\t<tbody>\n\t\t\t\t  <tr>\n\t\t\t\t    <td class=\"col-sm-3\">{{.StatusCode}}</td>\n\t\t\t\t    <td class=\"col-sm-9\"><a href=\"#{{.Return}}Model\" data-toggle=\"collapse\" data-target=\"#{{.Return}}ModelInfo\">{{.Return}}</a></td>\n\t\t\t\t  </tr>\n\t\t\t\t</tbody>\n\t\t\t  </table>\n\t\t  </div><!-- #{{$controller.Name}}{{.Method}}Info -->\n      </div><!-- .panel -->{{ end }}{{ end }}\n    </div><!-- .container -->\n\n    <div class=\"container\">\n\t  <h2>Models</h2>\n\t  {{ range .Models }}\n\t  <div class=\"panel panel-default\" id=\"{{.Name}}Model\">\n\t\t<div class=\"panel-heading\">\n          <div class=\"container\">\n            <div class=\"col-sm-11\">{{.Name}}</div>\n\t\t\t<div class=\"col-sm-1\"><span data-toggle=\"collapse\" data-target=\"#{{.Name}}ModelInfo\"><span class=\"glyphicon glyphicon-sort\"></span></span></div>\n          </div>\n        </div>\n\t\t<div class=\"panel-body collapse\" id=\"{{.Name}}ModelInfo\">\n\t\t  <p>{{ .Description }}</p>\n          <table class=\"table\">\n            <thead>\n              <tr>\n                <th class=\"col-sm-3\">Name</th>\n                <th class=\"col-sm-9\">Type</th>\n              </tr>\n            </thead>\n            <tbody>{{ range .Fields }}\n              <tr>\n                <td class=\"col-sm-3\">{{ .JSON }}</td>\n                <td class=\"col-sm-9\">{{ .JSONType }}</td>\n\t\t\t  </tr>{{end}}\n\t\t\t</tbody>\n          </table>\n\t    </div><!-- #{{.Name}}ModelInfo -->\n\t  </div><!-- #{{.Name}}Model -->{{ end }}\n      {{ range $index, $name := .Collections }}\n\t  <div class=\"panel panel-default\" id=\"{{$name}}CollectionModel\">\n\t\t<div class=\"panel-heading\">\n          <div class=\"container\">\n            <div class=\"col-sm-11\">{{$name}}Collection</div>\n\t\t\t<div class=\"col-sm-1\"><span data-toggle=\"collapse\" data-target=\"#{{$name}}CollectionModelInfo\"><span class=\"glyphicon glyphicon-sort\"></span></span></div>\n          </div>\n        </div>\n\t\t<div class=\"panel-body collapse\" id=\"{{$name}}CollectionModelInfo\">\n\t\t  <p>{{$name}}Collection is a paginated collection of {{$name}} models</p>\n          <table class=\"table\">\n            <thead>\n              <tr>\n                <th class=\"col-sm-3\">Name</th>\n                <th class=\"col-sm-9\">Type</th>\n              </tr>\n            </thead>\n            <tbody>\n              <tr>\n                <td class=\"col-sm-3\">next</td>\n                <td class=\"col-sm-9\">string</td>\n              </tr>\n              <tr>\n                <td class=\"col-sm-3\">previous</td>\n                <td class=\"col-sm-9\">string</td>\n              </tr>\n              <tr>\n                <td class=\"col-sm-3\">uri</td>\n                <td class=\"col-sm-9\">string</td>\n\t\t\t  </tr>\n              <tr>\n                <td class=\"col-sm-3\">items</td>\n                <td class=\"col-sm-9\"><a href=\"#{{$name}}Model\">[]{{$name}}</a></td>\n              </tr>\n\t\t\t</tbody>\n          </table>\n\t    </div><!-- #{{$name}}CollectionModelInfo -->\n\t  </div><!-- #{{$name}}CollectionModel -->{{ end }}\n  </div><!-- .contrainer -->\n\n\t<!-- Optional JavaScript -->\n\t<!-- jQuery first, then Bootstrap JS -->\n\t<script src=\"https://code.jquery.com/jquery-3.2.1.slim.min.js\" integrity=\"sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN\" crossorigin=\"anonymous\"></script>\n\t<script src=\"https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js\" integrity=\"sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa\" crossorigin=\"anonymous\"></script>\n  </body>\n</html>`\n")))
    
    // Gitignore
    template.Must(master.New(Gitignore).Parse(string("# THIS IS A GENERATED FILE. DO NOT MODIFY\n# gitignore.tmpl\n\n{{ .Lower .Name }}\n")))
    
    // Main
    template.Must(master.New(Main).Parse(string("package main\n\n// THIS IS A GENERATED FILE. DO NOT MODIFY\n// main.tmpl\n\nimport (\n    \"fmt\"\n\tqcontrollers \"github.com/beaconsoftwarellc/quimby/controllers\"\n\tqhttp \"github.com/beaconsoftwarellc/quimby/http\"\n\t\"{{.BasePath}}/controllers\"\n\t\"{{.BasePath}}/config\"\n    \"github.com/beaconsoftwarellc/gadget/log\"\n)\n\n//go:generate codegen definition.yaml\n\nfunc main() {\n\tlog.NewGlobal(\"{{.Name}}\", log.FunctionFromEnv())\n\t// Constants are defined on http\n\t// see: https://golang.org/pkg/net/http/#\n\tspecification := config.New()\n\trootController := &qcontrollers.HealthCheckController{}\n\tserver := qhttp.CreateRESTServer(fmt.Sprintf(\":%d\", specification.Port), rootController)\n\tserver.Router.AddController(&qcontrollers.HealthCheckController{})\n\tserver.Router.AddController(controllers.NewDocController(specification))\n\n\t{{ range $index, $controller := .Controllers }}{{$controller.Add}}\n\t{{ end }}\n\tlog.Infof(\"Server starting ... http://localhost:%d/\", specification.Port)\n\tlog.Error(server.ListenAndServe())\n}\n")))
    
    // Models
    template.Must(master.New(Models).Parse(string("package models\n\n// THIS IS A GENERATED FILE. DO NOT MODIFY\n// models.tmpl\n\nimport (\n\t\"time\"\n)\n\n{{range .Models}}{{if not .Fields -}}// {{.Name}} {{.Description}}\ntype {{.Name}} map[string]interface{} {{else}}\n// {{.Name}} {{.Description}}\ntype {{.Name}} struct {\n\t{{- range .Fields}}\n\t{{.Name}} {{.Type}} `json:\"{{.JSON}}\"`\n\t{{- end}}\n}\n{{end}}{{end}}\n\n{{- range $index, $name := .Collections }}// {{$name}}Collection is a paginated collection of {{$name}} models\ntype {{$name}}Collection struct {\n\tCollection\n\tItems []*{{$name}} `json:\"items\"`\n}\n{{- end}}\n")))
    
    // TestConfig
    template.Must(master.New(TestConfig).Parse(string("package test\n\n// THIS IS A GENERATED FILE. DO NOT MODIFY\n// test_config.tmpl\n\nimport (\n\t\"github.com/beaconsoftwarellc/gadget/environment\"\n\t\"github.com/beaconsoftwarellc/gadget/log\"\n\t\"{{.BasePath}}/config\"\n)\n\n// NewMockSpec for use in unit tests.\nfunc NewMockSpec() *config.Specification {\n\tspec := &config.Specification{\n\t\tLog: log.NewStackLogger(),\n\t}\n\treturn spec\n}\n")))
    
    return master
}
