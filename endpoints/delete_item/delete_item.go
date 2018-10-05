// Support for the DynamoDB DeleteItem endpoint.
//
// example use:
//
// tests/delete_item-livestest.go
//
package delete_item

import (
	"encoding/json"
	"errors"

	"github.com/smugmug/godynamo/authreq"
	"github.com/smugmug/godynamo/aws_const"
	"github.com/smugmug/godynamo/conf"
	"github.com/smugmug/godynamo/types/attributesresponse"
	"github.com/smugmug/godynamo/types/attributevalue"
	"github.com/smugmug/godynamo/types/aws_strings"
	"github.com/smugmug/godynamo/types/expected"
	"github.com/smugmug/godynamo/types/expressionattributenames"
	"github.com/smugmug/godynamo/types/item"
)

const (
	ENDPOINT_NAME       = "DeleteItem"
	DELETEITEM_ENDPOINT = aws_const.ENDPOINT_PREFIX + ENDPOINT_NAME
	// the permitted ReturnValues flags for this op
	RETVAL_ALL_OLD = aws_strings.RETVAL_ALL_OLD
	RETVAL_NONE    = aws_strings.RETVAL_NONE
)

type DeleteItem struct {
	ConditionExpression         string                                            `json:",omitempty"`
	ConditionalOperator         string                                            `json:",omitempty"`
	Expected                    expected.Expected                                 `json:",omitempty"`
	ExpressionAttributeNames    expressionattributenames.ExpressionAttributeNames `json:",omitempty"`
	ExpressionAttributeValues   attributevalue.AttributeValueMap                  `json:",omitempty"`
	Key                         item.Key
	ReturnConsumedCapacity      string `json:",omitempty"`
	ReturnItemCollectionMetrics string `json:",omitempty"`
	ReturnValues                string `json:",omitempty"`
	TableName                   string
}

func NewDeleteItem() *DeleteItem {
	u := new(DeleteItem)
	u.Expected = expected.NewExpected()
	u.ExpressionAttributeNames = expressionattributenames.NewExpressionAttributeNames()
	u.ExpressionAttributeValues = attributevalue.NewAttributeValueMap()
	u.Key = item.NewKey()
	return u
}

// Delete is an alias for backwards compatibility
type Delete DeleteItem

func NewDelete() *Delete {
	delete_item := NewDeleteItem()
	delete := Delete(*delete_item)
	return &delete
}

type Request DeleteItem

type Response attributesresponse.AttributesResponse

func NewResponse() *Response {
	a := attributesresponse.NewAttributesResponse()
	r := Response(*a)
	return &r
}

// These implementations of EndpointReq use a parameterized conf.

func (delete_item *DeleteItem) EndpointReqWithConf(c *conf.AWS_Conf) ([]byte, int, error) {
	if delete_item == nil {
		return nil, 0, errors.New("delete_item.(DeleteItem)EndpointReqWithConf: receiver is nil")
	}
	if !conf.IsValid(c) {
		return nil, 0, errors.New("delete_item.EndpointReqWithConf: c is not valid")
	}
	// returns resp_body,code,err
	reqJSON, json_err := json.Marshal(delete_item)
	if json_err != nil {
		return nil, 0, json_err
	}
	return authreq.RetryReqJSON_V4WithConf(reqJSON, DELETEITEM_ENDPOINT, c)
}

func (delete *Delete) EndpointReqWithConf(c *conf.AWS_Conf) ([]byte, int, error) {
	if delete == nil {
		return nil, 0, errors.New("delete_item.(Delete)EndpointReqWithConf: receiver is nil")
	}
	delete_item := DeleteItem(*delete)
	return delete_item.EndpointReqWithConf(c)
}

func (req *Request) EndpointReqWithConf(c *conf.AWS_Conf) ([]byte, int, error) {
	if req == nil {
		return nil, 0, errors.New("delete_item.(Request)EndpointReqWithConf: receiver is nil")
	}
	delete_item := DeleteItem(*req)
	return delete_item.EndpointReqWithConf(c)
}

// These implementations of EndpointReq use the global conf.

func (delete_item *DeleteItem) EndpointReq() ([]byte, int, error) {
	if delete_item == nil {
		return nil, 0, errors.New("delete_item.(DeleteItem)EndpointReq: receiver is nil")
	}
	return delete_item.EndpointReqWithConf(&conf.Vals)
}

func (delete *Delete) EndpointReq() ([]byte, int, error) {
	if delete == nil {
		return nil, 0, errors.New("delete_item.(Delete)EndpointReq: receiver is nil")
	}
	delete_item := DeleteItem(*delete)
	return delete_item.EndpointReqWithConf(&conf.Vals)
}

func (req *Request) EndpointReq() ([]byte, int, error) {
	if req == nil {
		return nil, 0, errors.New("delete_item.(Request)EndpointReq: receiver is nil")
	}
	delete_item := DeleteItem(*req)
	return delete_item.EndpointReqWithConf(&conf.Vals)
}
