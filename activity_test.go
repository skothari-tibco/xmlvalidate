package xmlvalidate

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}
func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput("path", `<xs:schema attributeFormDefault="unqualified" elementFormDefault="qualified" xmlns:xs="http://www.w3.org/2001/XMLSchema">
	<xs:element name="html">
	  <xs:complexType>
		<xs:sequence>
		  <xs:element type="xs:string" name="body"/>
		</xs:sequence>
	  </xs:complexType>
	</xs:element>
  </xs:schema>`)

	tc.SetInput("text", `<html>
  <body> Hey </body>
  </html>`)

	val, _ := act.Eval(tc)

	fmt.Println(val)
	assert.Equal(t, val, true)

}
func TestEval2(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	tc.SetInput("path", `<xs:schema attributeFormDefault="unqualified" elementFormDefault="qualified" xmlns:xs="http://www.w3.org/2001/XMLSchema">
	<xs:element name="html">
	  <xs:complexType>
		<xs:sequence>
		  <xs:element type="xs:string" name="body"/>
		</xs:sequence>
	  </xs:complexType>
	</xs:element>
  </xs:schema>`)

	tc.SetInput("text", `<html>
  <xml> Hey </xml>
  </html>`)

	val, _ := act.Eval(tc)
	fmt.Println(val)
	result := tc.GetOutput("isValid")
	assert.Equal(t, result, false)

}
