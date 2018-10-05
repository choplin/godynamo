package update_table

import (
	"encoding/json"
	"testing"
)

func TestNil(t *testing.T) {
	s := NewUpdateTable()
	_, _, err := s.EndpointReqWithConf(nil)
	if err == nil {
		t.Errorf("nil conf should result in error")
	}
}

func TestRequestUnmarshal(t *testing.T) {
	s := []string{`{"TableName":"Thread","ProvisionedThroughput":{"ReadCapacityUnits":10,"WriteCapacityUnits":10}}`}
	for _, v := range s {
		var u UpdateTable
		um_err := json.Unmarshal([]byte(v), &u)
		if um_err != nil {
			t.Errorf("cannot unmarshal\n")
		}
		_, jerr := json.Marshal(u)
		if jerr != nil {
			t.Errorf("cannot marshal\n")
		}
	}
}

func TestResponseUnmarshal(t *testing.T) {
	s := []string{`{"TableDescription":{"AttributeDefinitions":[{"AttributeName":"ForumName","AttributeType":"S"},{"AttributeName":"LastPostDateTime","AttributeType":"S"},{"AttributeName":"Subject","AttributeType":"S"}],"CreationDateTime":1.363801528686E9,"ItemCount":0,"KeySchema":[{"AttributeName":"ForumName","KeyType":"HASH"},{"AttributeName":"Subject","KeyType":"RANGE"}],"LocalSecondaryIndexes":[{"IndexName":"LastPostIndex","IndexSizeBytes":0,"ItemCount":0,"KeySchema":[{"AttributeName":"ForumName","KeyType":"HASH"},{"AttributeName":"LastPostDateTime","KeyType":"RANGE"}],"Projection":{"ProjectionType":"KEYS_ONLY"}}],"ProvisionedThroughput":{"LastIncreaseDateTime":1.363801701282E9,"NumberOfDecreasesToday":0,"ReadCapacityUnits":5,"WriteCapacityUnits":5},"TableName":"Thread","TableSizeBytes":0,"TableStatus":"UPDATING"}}`}
	for _, v := range s {
		var u Response
		um_err := json.Unmarshal([]byte(v), &u)
		if um_err != nil {
			t.Errorf("cannot unmarshal\n")
		}
		_, jerr := json.Marshal(u)
		if jerr != nil {
			t.Errorf("cannot marshal\n")
		}
	}
}
