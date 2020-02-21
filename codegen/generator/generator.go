package generator

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"github.com/beaconsoftwarellc/gadget/fileutil"
	"github.com/beaconsoftwarellc/gadget/log"
	"github.com/beaconsoftwarellc/gadget/stringutil"
	"github.com/beaconsoftwarellc/quimby/codegen/templates"
)

const (
	fileMode  = os.FileMode(0755)
	basicAuth = "Basic"
	tokenAuth = "Token"

	typeTime       = "time.Time"
	arrayIndicator = "[]"
)

// Generator defines the functions required to be a Code Generator
type Generator interface {
	Run() error
}

// New returns an instatiated Code Generator for Quimby
func New(definitionFileName string) (Generator, error) {
	data, err := ioutil.ReadFile(definitionFileName)
	if nil != err {
		return nil, err
	}

	gen := &generator{
		definition: &definition{},
	}

	err = yaml.Unmarshal(data, gen.definition)
	if nil != err {
		return nil, err
	}
	gen.definition.findCollections()

	return gen, nil
}

type generator struct {
	definition *definition
}

func (gen *generator) Run() error {
	files := map[string]string{ //full output path -> template file
		filepath.Join("controllers", "controller.gen.go"):  "controller.tmpl",
		filepath.Join("controllers", "docs.gen.go"):        "docs.tmpl",
		filepath.Join("models", "models.gen.go"):           "models.tmpl",
		filepath.Join("main.gen.go"):                       "main.tmpl",
		filepath.Join(".gitignore"):                        "gitignore.tmpl",
		filepath.Join("config", "specification.gen.go"):    "config.tmpl",
		filepath.Join("test", "specification_test.gen.go"): "test_config.tmpl",
	}
	for path, template := range files {
		dir, _ := filepath.Split(path)
		if dir != "" {
			_, err := fileutil.EnsureDir(dir, fileMode)
			log.ExitOnError(err)
		}

		gen.writeTemplateToFile(path, template)
	}

	return nil
}

func (gen *generator) writeTemplateToFile(filename, templateName string) {
	fmt.Printf("creating %s (%s) with %s ... ", filename, fileMode, templateName)
	fd, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, fileMode)
	if nil != err {
		fmt.Printf("%s\n", "FAILED")
		log.Error(err)
	}
	defer fd.Close()
	template := templates.GetTemplates()
	err = template.ExecuteTemplate(fd, templateName, gen.definition)
	if nil != err {
		fmt.Printf("%s\n", "FAILED")
		log.ExitOnError(err)
	}
	fmt.Printf("%s\n", "SUCCESS")
}

type definition struct {
	Name          string            `yaml:"Name"`
	BasePath      string            `yaml:"BasePath"`
	Specification specification     `yaml:"Specification"`
	Controllers   []*controller     `yaml:"Controllers"`
	Models        []*model          `yaml:"Models"`
	Collections   []string          `yaml:"-"`
	ModelMap      map[string]*model `yaml:"-"`
}

// specification represents the config that should be generated for this service
type specification struct {
	Imports     []string      `yaml:"Imports"`
	Environment []environment `yaml:"Environment"`
	Services    []service     `yaml:"Services"`
}

type service struct {
	Name        string `yaml:"Name"`
	Type        string `yaml:"Type"`
	Initializer string `yaml:"Initializer"`
}

// environment defines environment variables for a config
type environment struct {
	Name       string      `yaml:"Name"`
	Type       string      `yaml:"Type"`
	Env        string      `yaml:"Env"`
	S3         string      `yaml:"S3"`
	Optional   bool        `yaml:"Optional"`
	Default    interface{} `yaml:"Default"`
	LowerLimit int         `yaml:"LowerLimit"`
	UpperLimit int         `yaml:"UpperLimit"`
}

type controller struct {
	Name        string          `yaml:"Name"`
	Description string          `yaml:"Description"`
	Auth        *authentication `yaml:"Auth"`
	Routes      []string        `yaml:"Routes"`
	Parameters  []*parameter    `yaml:"Parameters"`
	Actions     []*action       `yaml:"Actions"`
}

// Add generates the line to add a controller to the server
func (c *controller) Add() string {
	return fmt.Sprintf("server.Router.AddController(controllers.New%sController(specification))", c.Name)
}

// ImplName returns the Implementation Name for the Service
func (c *controller) ImplName() string {
	return stringutil.LowerCamelCase(c.Name)
}

// DefaultRoute returns the first route for the Controller
func (c *controller) DefaultRoute() string {
	return c.Routes[0]
}

// HasAuth determines if the controller has Authentication
func (c *controller) HasAuth() bool {
	return nil != c.Auth
}

// Security returns the base struct for the controller to extend based on the Authentication specified
func (c *controller) Security() string {
	if !c.HasAuth() {
		return "qcontrollers.MethodNotAllowedController\n\tqcontrollers.NoAuthenticationController"
	}
	switch c.Auth.Type {
	case basicAuth:
		return "\tqcontrollers.BasicAuthenticatedController"
	case tokenAuth:
		return "\tqcontrollers.TokenAuthenticatedController"
	default:
		return "\tsecurity." + c.Auth.Type + "AuthenticatedController"
	}
}

// AuthHeaders determine which headers are needed for a controller based on the specified Auth
func (c *controller) AuthHeaders() []parameter {
	headers := make([]parameter, 0)
	if !c.HasAuth() {
		return headers
	}
	switch c.Auth.Type {
	case basicAuth:
		headers = append(headers, parameter{Name: "Authorization", Type: "Standard HTTP Basic Auth"})
	case tokenAuth:
		headers = append(headers, parameter{Name: "Authorization", Type: "Standard Bearer Token Authorization Header"})
	default:
		headers = append(headers, parameter{Name: c.Auth.Type, Type: c.Auth.Description})
	}
	return headers
}

// LoadAuth handles getting the Authentication from Security / Request based on type of Authentication
func (c *controller) LoadAuth() string {
	if !c.HasAuth() {
		return ""
	}
	switch c.Auth.Type {
	case basicAuth:
		return `username, password, ok := context.Request.BasicAuth()
			if !ok {
				context.SetError(qerror.NewRestError(qerror.AuthenticationFailed, "", nil), http.StatusUnauthorized)
			}`
	default:
		return "auth := security.GetAuthentication(context)"
	}
}

type authentication struct {
	Type        string `yaml:"Type"`
	Validator   string `yaml:"Validator"`
	Description string `yaml:"Description"`
}

// model represents the meta data for constructing an API model
type model struct {
	Name        string       `yaml:"Name"`
	Description string       `yaml:"Description"`
	Fields      []modelField `yaml:"Fields"`
}

// modelField represents the meta data for constructing a field on an API model
type modelField struct {
	Name string `yaml:"Name"`
	Type string `yaml:"Type"`
}

type parameter struct {
	Name string `yaml:"Name"`
	Type string `yaml:"Type"`
}

type action struct {
	Method              string `yaml:"Method"`
	Description         string `yaml:"Description"`
	Return              string `yaml:"Return"`
	Model               string `yaml:"Model"`
	Parameters          string `yaml:"Parameters"`
	OverwriteStatusCode *int   `yaml:"OverwriteStatusCode,omitempty"`
}

// OverwriteStatus helps return the provided ResponseStatusCode ignoring the Method
func (a *action) OverwriteStatus() (string, int) {
	switch *a.OverwriteStatusCode {
	case http.StatusOK:
		return "http.StatusOK", http.StatusOK
	case http.StatusCreated:
		return "http.StatusCreated", http.StatusCreated
	case http.StatusNoContent:
		return "http.StatusNoContent", http.StatusNoContent
	default:
		panic(fmt.Sprintf("invalid status: %d", *a.OverwriteStatusCode))
	}
}

// Status returns the proper HTTP Status string for the Method
func (a *action) Status() string {
	if nil != a.OverwriteStatusCode {
		stringStatus, _ := a.OverwriteStatus()
		return stringStatus
	}
	switch strings.ToUpper(a.Method) {
	case http.MethodGet, http.MethodPut, http.MethodPatch:
		return "http.StatusOK"
	case http.MethodPost:
		return "http.StatusCreated"
	case http.MethodDelete:
		return "http.StatusNoContent"
	default:
		panic(fmt.Sprintf("invalid method: %s", a.Method))
	}
}

// StatusCode returns the proper HTTP Status code for the Method
func (a *action) StatusCode() int {
	if nil != a.OverwriteStatusCode {
		_, intStatus := a.OverwriteStatus()
		return intStatus
	}
	switch strings.ToUpper(a.Method) {
	case http.MethodGet, http.MethodPut, http.MethodPatch:
		return http.StatusOK
	case http.MethodPost:
		return http.StatusCreated
	case http.MethodDelete:
		return http.StatusNoContent
	default:
		panic(fmt.Sprintf("invalid method: %s", a.Method))
	}
}

// PanelCode returns the proper bootstrap class for the Method
func (a *action) PanelCode() string {
	switch strings.ToUpper(a.Method) {
	case http.MethodGet:
		return "success"
	case http.MethodPut, http.MethodPatch:
		return "warning"
	case http.MethodPost:
		return "info"
	case http.MethodDelete:
		return "danger"
	default:
		panic(fmt.Sprintf("invalid method: %s", a.Method))
	}
}

func (a *action) ReturnDefinition() string {
	if stringutil.IsEmpty(a.Return) {
		return "errors.TracerError"
	}
	return fmt.Sprintf("(*models.%s, errors.TracerError)", a.Return)
}

const getRequestModelFmt = `request := &models.%s{}
if err := context.ReadObject(request); nil != err {
	context.SetError(&qerror.RestError{Code: qerror.ValidationError, Message: err.Error()}, http.StatusNotAcceptable)
	return
}
`
const getRequestModelArrayFmt = `request := make([]*models.%s, 0)
if err := context.ReadObject(&request); nil != err {
	context.SetError(&qerror.RestError{Code: qerror.ValidationError, Message: err.Error()}, http.StatusNotAcceptable)
	return
}
`

const getQueryModelFmt = `queryParams := &models.%s{}
if err := context.ReadQueryParams(queryParams); nil != err {
	context.SetError(&qerror.RestError{Code: qerror.ValidationError, Message: err.Error()}, http.StatusNotAcceptable)
	return
}
`

// ReadRequestModel generates the controller code to read a request from the body
func (a *action) ReadRequestModel() string {
	if stringutil.IsEmpty(a.Model) {
		return ""
	}
	if strings.HasPrefix(a.Model, arrayIndicator) {
		return fmt.Sprintf(getRequestModelArrayFmt, strings.Replace(a.Model, arrayIndicator, "", 1))
	}
	return fmt.Sprintf(getRequestModelFmt, a.Model)
}

// ReadQueryModel generates the controller code to read a request from the query parameters
func (a *action) ReadQueryModel() string {
	if stringutil.IsEmpty(a.Parameters) {
		return ""
	}
	return fmt.Sprintf(getQueryModelFmt, a.Parameters)
}

const errorLogModelFmt = `controller.Specification.Log.Infof("%%#v", map[string]string{"context": fmt.Sprintf("%%#v", context), %s"error": fmt.Sprintf("%%#v", err)})`

func (a *action) ErrorLog() string {
	requestModel := `"payload": fmt.Sprintf("%#v", request), `
	if stringutil.IsEmpty(a.Model) {
		requestModel = ""
	}
	return fmt.Sprintf(errorLogModelFmt, requestModel)
}

func (a *action) ArgumentDefinition(c controller) string {
	arguments := make([]string, 0)
	if c.HasAuth() {
		if basicAuth == c.Auth.Type {
			arguments = append(arguments, "username string", "password string")
		} else {
			arguments = append(arguments, "auth *security.Authentication")
		}
	}
	for _, param := range c.Parameters {
		arguments = append(arguments, stringutil.LowerCamelCase(c.Name+param.ParameterName())+" "+param.Type)
	}

	if strings.HasPrefix(a.Model, arrayIndicator) {
		arguments = append(arguments, "request []*models."+strings.Replace(a.Model, arrayIndicator, "", 1))
	} else if !stringutil.IsEmpty(a.Model) {
		arguments = append(arguments, "request *models."+a.Model)
	}
	if !stringutil.IsEmpty(a.Parameters) {
		arguments = append(arguments, "query *models."+a.Parameters)
	}
	return strings.Join(arguments, ", ")
}

func (a *action) Arguments(c controller) string {
	arguments := make([]string, 0)
	if c.HasAuth() {
		if basicAuth == c.Auth.Type {
			arguments = append(arguments, "username", "password")
		} else {
			arguments = append(arguments, "auth")
		}
	}
	for _, param := range c.Parameters {
		arguments = append(arguments, fmt.Sprintf("context.URIParameters[\"%s\"]", param.Name))
	}

	if !stringutil.IsEmpty(a.Model) {
		arguments = append(arguments, "request")
	}
	if !stringutil.IsEmpty(a.Parameters) {
		arguments = append(arguments, "queryParams")
	}
	return strings.Join(arguments, ", ")
}

func (a *action) QueryParameters(d *definition) []modelField {
	if stringutil.IsEmpty(a.Parameters) {
		return []modelField{}
	}
	d.mapModels()

	model, ok := d.ModelMap[a.Parameters]
	if !ok {
		panic(fmt.Sprintf("%s not found in ModelMap", a.Parameters))
	}
	return model.Fields
}

func (p parameter) ParameterName() string {
	if "id" == p.Name {
		return "ID"
	}
	return strings.Title(p.Name)
}

// JSON converts a ModelField.Name into an underscored version for a column
func (df *modelField) JSON() string {
	return stringutil.Underscore(df.Name)
}

const dateTimeType = "datetime"

// JSONType converts a ModelField.Type into the JSON version for a column
func (df *modelField) JSONType() string {
	if strings.HasSuffix(df.Type, typeTime) {
		return dateTimeType
	}
	if strings.HasSuffix(df.Type, "mysql.NullTime") {
		return dateTimeType
	}
	return df.Type
}

// QueryName converts a ModelField.Name into an underscored version for a query parameter
func (df *modelField) QueryName() string {
	if strings.HasPrefix(df.Type, "[]") {
		return df.JSON() + "[]"
	}
	return df.JSON()
}

// QueryType converts a ModelField.Type into the JSON version for a query parameter
func (df *modelField) QueryType() string {
	if strings.HasPrefix(df.Type, "[]") {
		return strings.Replace(df.Type, "[]", "", 1)
	}
	return df.JSONType()
}

// Lower case the passed string. Used inside templates.
func (definition *definition) Lower(s string) string {
	return strings.ToLower(s)
}

func (definition *definition) findCollections() {
	definition.Collections = make([]string, 0)
	collections := make(map[string]bool)

	for _, controller := range definition.Controllers {
		for _, action := range controller.Actions {
			if strings.HasSuffix(action.Return, "Collection") {
				collection := strings.Replace(action.Return, "Collection", "", 1)
				if !collections[collection] {
					definition.Collections = append(definition.Collections, collection)
					collections[collection] = true
				}
			}
		}
	}
}

func (definition *definition) mapModels() {
	definition.ModelMap = make(map[string]*model, len(definition.Models))
	for _, model := range definition.Models {
		definition.ModelMap[model.Name] = model
	}
}
