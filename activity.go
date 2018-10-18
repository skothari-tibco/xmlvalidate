package xmlvalidate

import (
	"io/ioutil"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/lestrrat-go/libxml2"
	_ "github.com/lestrrat-go/libxml2/clib"
	"github.com/lestrrat-go/libxml2/parser"
	"github.com/lestrrat-go/libxml2/xsd"
)

var activityLog = logger.GetLogger("activity-flogo-xmlactivity")

func init() {
	activityLog.SetLogLevel(logger.InfoLevel)
}

type XmlValidate struct {
	metadata *activity.Metadata
}

func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &XmlValidate{metadata: metadata}
}

func (a *XmlValidate) Metadata() *activity.Metadata {
	return a.metadata
}

func (a *XmlValidate) Eval(ctx activity.Context) (done bool, err error) {

	xml := ctx.GetInput("text").(string)
	xsds := ctx.GetInput("path").(string)

	var schema *xsd.Schema

	if isPath(xsds) {
		byteArray, _ := ioutil.ReadFile(strings.Split(xsds, "://")[1])
		schema, err = xsd.Parse(byteArray)
		if err != nil {
			return true, err
		}
	} else {
		schema, err = xsd.Parse([]byte(xsds))
		if err != nil {
			return true, err
		}
	}

	defer schema.Free()
	doc, err := libxml2.Parse([]byte(xml), parser.XMLParseRecover)
	if err := schema.Validate(doc); err != nil {
		//fmt.Println("Error")
		ctx.SetOutput("isValid", false)
		return true, nil
	}
	ctx.SetOutput("isValid", true)
	return true, nil
}

func isPath(s string) bool {

	return strings.Contains(s, "file://")
}
