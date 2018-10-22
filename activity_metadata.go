package xmlvalidate

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `
{
    "name": "xmlvalidate",
    "type": "flogo:activity",
    "ref": "github.com/skothari-tibco/xmlvalidate",
    "version": "0.0.1",
    "title": "XML Validator",
    "description": "Simple XML Validator Activity",
    "homepage": " ",
    "input":[
      {
        "name": "text",
        "type": "string",
        "value": ""
      },
      {
        "name": "path",
        "type":"string",
        "value":""
      }
    ],
    "output": [
      {
        "name": "isValid",
        "type": "bool"
      },
      {
        "name": "log",
        "type": "string"
      }
    ]
  }
  `

func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}
