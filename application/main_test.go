package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

// TODO: create common func.
// func DecodeJSON(src string) events.SNSEventRecord {
// 	data, err := ioutil.ReadFile(src)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		os.Exit(1)
// 	}

// 	var sns events.SNSEventRecord
// 	if err := json.Unmarshal(data, &sns); err != nil {
// 		log.Fatal(err)
// 	}
// 	//fmt.Printf("%v \n", sns)
// 	return sns
// }

var permanent events.SNSEventRecord = events.SNSEventRecord{
	EventVersion: "EventVersion",
	SNS:          events.SNSEntity{Signature: "Signature", Message: "{\r\n\"notificationType\": \"Bounce\",\r\n\"bounce\": {\r\n\"bounceType\": \"Permanent\",\r\n\"bounceSubType\": \"General\",\r\n\"bouncedRecipients\": [\r\n{\r\n\"emailAddress\": \"bounce@simulator.amazonses.com\",\r\n\"action\": \"failed\",\r\n\"status\": \"5.1.1\",\r\n\"diagnosticCode\": \"smtp; 550 5.1.1 user unknown\"\r\n}\r\n],\r\n\"timestamp\": \"2020-02-04T01:41:55.872Z\",\r\n\"feedbackId\": \"010001700ddc71e5-88ddb7ba-bb6a-48b4-9322-e346306e1f1c-000000\",\r\n\"remoteMtaIp\": \"18.206.142.166\",\r\n\"reportingMTA\": \"dsn; a48-131.smtp-out.amazonses.com\"\r\n}\r\n}"},
}
var transient events.SNSEventRecord = events.SNSEventRecord{
	EventVersion: "EventVersion",
	SNS:          events.SNSEntity{Signature: "Signature", Message: "{\r\n\"notificationType\": \"Bounce\",\r\n\"bounce\": {\r\n\"bounceType\": \"Transient\",\r\n\"bounceSubType\": \"General\",\r\n\"bouncedRecipients\": [\r\n{\r\n\"emailAddress\": \"success@simulator.amazonses.com\",\r\n\"action\": \"failed\",\r\n\"status\": \"5.1.1\",\r\n\"diagnosticCode\": \"smtp; 550 5.1.1 user unknown\"\r\n}\r\n],\r\n\"timestamp\": \"2020-02-04T01:41:55.872Z\",\r\n\"feedbackId\": \"010001700ddc71e5-88ddb7ba-bb6a-48b4-9322-e346306e1f1c-000000\",\r\n\"remoteMtaIp\": \"18.206.142.166\",\r\n\"reportingMTA\": \"dsn; a48-131.smtp-out.amazonses.com\"\r\n}\r\n}"},
}

var event = events.SNSEvent{Records: []events.SNSEventRecord{permanent}}

func TestFilterEmail(t *testing.T) {
	cases := map[string]struct {
		actual   events.SNSEvent
		expected string
	}{
		"blacklisting": {
			actual:   events.SNSEvent{Records: []events.SNSEventRecord{permanent}},
			expected: "bounce@simulator.amazonses.com",
		},
		"not blacklisting": {
			actual:   events.SNSEvent{Records: []events.SNSEventRecord{transient}},
			expected: "",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if got := FilterEmail(c.actual); got != c.expected {
				t.Errorf("got %s, to %s", got, c.expected)
			}
		})
	}
}
