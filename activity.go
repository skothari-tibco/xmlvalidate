package xmlvalidate

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/lestrrat-go/libxml2"
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
	} else {
		schema, err = xsd.Parse([]byte(xsds))
	}

	defer schema.Free()
	doc, err := libxml2.Parse([]byte(xml), parser.XMLParseRecover)
	if err := schema.Validate(doc); err != nil {
		fmt.Println("Error")
		return false, nil
	}
	return true, nil
}

func isPath(s string) bool {

	return strings.Contains(s, "file://")
}
