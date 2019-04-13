package rest

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
  "name": "tibco-rest",
  "type": "flogo:activity",
  "ref": "github.com/bckumar2080/rest",
  "version": "0.0.1",
  "title": "Invoke REST Service",
  "description": "Simple REST Activity",
  "homepage": "https://github.com/bckumar2080/rest",
  "input":[
    {
      "name": "method",
      "type": "string",
      "required": true,
      "allowed" : ["GET", "POST", "PUT", "PATCH", "DELETE"]
    },
    {
      "name": "uri",
      "type": "string",
      "required": true
    },
    {
      "name": "proxy",
      "type": "string",
      "required": false
    },
    {
      "name": "pathParams",
      "type": "params"
    },
    {
      "name": "queryParams",
      "type": "params"
    },
    {
      "name": "header",
      "type": "params"
    },
    {
      "name": "skipSsl",
      "type": "boolean",
      "value": false
    },
    {
      "name": "content",
      "type": "any"
    }
  ],
  "output": [
    {
      "name": "result",
      "type": "any"
    },
    {
      "name": "status",
      "type": "integer"
    }
  ]
}`

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}
