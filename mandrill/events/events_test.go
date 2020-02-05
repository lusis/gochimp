package events

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// nolint: deadcode
var testExampleJSON = `[{"event":"send","msg":{"ts":1365109999,"subject":"This an example webhook message","email":"example.webhook@mandrillapp.com","sender":"example.sender@mandrillapp.com","tags":["webhook-example"],"opens":[],"clicks":[],"state":"sent","metadata":{"user_id":111},"_id":"exampleaaaaaaaaaaaaaaaaaaaaaaaaa","_version":"exampleaaaaaaaaaaaaaaa"},"_id":"exampleaaaaaaaaaaaaaaaaaaaaaaaaa","ts":1580762074},{"event":"deferral","msg":{"ts":1365109999,"subject":"This an example webhook message","email":"example.webhook@mandrillapp.com","sender":"example.sender@mandrillapp.com","tags":["webhook-example"],"opens":[],"clicks":[],"state":"deferred","metadata":{"user_id":111},"_id":"exampleaaaaaaaaaaaaaaaaaaaaaaaaa1","_version":"exampleaaaaaaaaaaaaaaa","smtp_events":[{"destination_ip":"127.0.0.1","diag":"451 4.3.5 Temporarily unavailable, try again later.","source_ip":"127.0.0.1","ts":1365111111,"type":"deferred","size":0}]},"_id":"exampleaaaaaaaaaaaaaaaaaaaaaaaaa1","ts":1580762074},{"event":"hard_bounce","msg":{"ts":1365109999,"subject":"This an example webhook message","email":"example.webhook@mandrillapp.com","sender":"example.sender@mandrillapp.com","tags":["webhook-example"],"state":"bounced","metadata":{"user_id":111},"_id":"exampleaaaaaaaaaaaaaaaaaaaaaaaaa2","_version":"exampleaaaaaaaaaaaaaaa","bounce_description":"bad_mailbox","bgtools_code":10,"diag":"smtp;550 5.1.1 The email account that you tried to reach does not exist. Please try double-checking the recipient's email address for typos or unnecessary spaces."},"_id":"exampleaaaaaaaaaaaaaaaaaaaaaaaaa2","ts":1580762074},{"event":"soft_bounce","msg":{"ts":1365109999,"subject":"This an example webhook message","email":"example.webhook@mandrillapp.com","sender":"example.sender@mandrillapp.com","tags":["webhook-example"],"state":"soft-bounced","metadata":{"user_id":111},"_id":"exampleaaaaaaaaaaaaaaaaaaaaaaaaa3","_version":"exampleaaaaaaaaaaaaaaa","bounce_description":"mailbox_full","bgtools_code":22,"diag":"smtp;552 5.2.2 Over Quota"},"_id":"exampleaaaaaaaaaaaaaaaaaaaaaaaaa3","ts":1580762074},{"event":"open","msg":{"ts":1365109999,"subject":"This an example webhook message","email":"example.webhook@mandrillapp.com","sender":"example.sender@mandrillapp.com","tags":["webhook-example"],"opens":[{"ts":1365111111}],"clicks":[{"ts":1365111111,"url":"http:\/\/mandrill.com"}],"state":"sent","metadata":{"user_id":111},"_id":"exampleaaaaaaaaaaaaaaaaaaaaaaaaa4","_version":"exampleaaaaaaaaaaaaaaa"},"_id":"exampleaaaaaaaaaaaaaaaaaaaaaaaaa4","ip":"127.0.0.1","location":{"country_short":"US","country":"United States","region":"Oklahoma","city":"Oklahoma City","latitude":35.4675598145,"longitude":-97.5164337158,"postal_code":"73101","timezone":"-05:00"},"user_agent":"Mozilla\/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.1.8) Gecko\/20100317 Postbox\/1.1.3","user_agent_parsed":{"type":"Email Client","ua_family":"Postbox","ua_name":"Postbox 1.1.3","ua_version":"1.1.3","ua_url":"http:\/\/www.postbox-inc.com\/","ua_company":"Postbox, Inc.","ua_company_url":"http:\/\/www.postbox-inc.com\/","ua_icon":"http:\/\/cdn.mandrill.com\/img\/email-client-icons\/postbox.png","os_family":"OS X","os_name":"OS X 10.6 Snow Leopard","os_url":"http:\/\/www.apple.com\/osx\/","os_company":"Apple Computer, Inc.","os_company_url":"http:\/\/www.apple.com\/","os_icon":"http:\/\/cdn.mandrill.com\/img\/email-client-icons\/macosx.png","mobile":false},"ts":1580762074},{"event":"click","msg":{"ts":1365109999,"subject":"This an example webhook message","email":"example.webhook@mandrillapp.com","sender":"example.sender@mandrillapp.com","tags":["webhook-example"],"opens":[{"ts":1365111111}],"clicks":[{"ts":1365111111,"url":"http:\/\/mandrill.com"}],"state":"sent","metadata":{"user_id":111},"_id":"exampleaaaaaaaaaaaaaaaaaaaaaaaaa5","_version":"exampleaaaaaaaaaaaaaaa"},"_id":"exampleaaaaaaaaaaaaaaaaaaaaaaaaa5","ip":"127.0.0.1","location":{"country_short":"US","country":"United States","region":"Oklahoma","city":"Oklahoma City","latitude":35.4675598145,"longitude":-97.5164337158,"postal_code":"73101","timezone":"-05:00"},"user_agent":"Mozilla\/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.1.8) Gecko\/20100317 Postbox\/1.1.3","user_agent_parsed":{"type":"Email Client","ua_family":"Postbox","ua_name":"Postbox 1.1.3","ua_version":"1.1.3","ua_url":"http:\/\/www.postbox-inc.com\/","ua_company":"Postbox, Inc.","ua_company_url":"http:\/\/www.postbox-inc.com\/","ua_icon":"http:\/\/cdn.mandrill.com\/img\/email-client-icons\/postbox.png","os_family":"OS X","os_name":"OS X 10.6 Snow Leopard","os_url":"http:\/\/www.apple.com\/osx\/","os_company":"Apple Computer, Inc.","os_company_url":"http:\/\/www.apple.com\/","os_icon":"http:\/\/cdn.mandrill.com\/img\/email-client-icons\/macosx.png","mobile":false},"url":"http:\/\/mandrill.com","ts":1580762074},{"event":"spam","msg":{"ts":1365109999,"subject":"This an example webhook message","email":"example.webhook@mandrillapp.com","sender":"example.sender@mandrillapp.com","tags":["webhook-example"],"opens":[{"ts":1365111111}],"clicks":[{"ts":1365111111,"url":"http:\/\/mandrill.com"}],"state":"sent","metadata":{"user_id":111},"_id":"exampleaaaaaaaaaaaaaaaaaaaaaaaaa6","_version":"exampleaaaaaaaaaaaaaaa"},"_id":"exampleaaaaaaaaaaaaaaaaaaaaaaaaa6","ts":1580762074},{"event":"unsub","msg":{"ts":1365109999,"subject":"This an example webhook message","email":"example.webhook@mandrillapp.com","sender":"example.sender@mandrillapp.com","tags":["webhook-example"],"opens":[{"ts":1365111111}],"clicks":[{"ts":1365111111,"url":"http:\/\/mandrill.com"}],"state":"sent","metadata":{"user_id":111},"_id":"exampleaaaaaaaaaaaaaaaaaaaaaaaaa7","_version":"exampleaaaaaaaaaaaaaaa"},"_id":"exampleaaaaaaaaaaaaaaaaaaaaaaaaa7","ts":1580762074},{"event":"reject","msg":{"ts":1365109999,"subject":"This an example webhook message","email":"example.webhook@mandrillapp.com","sender":"example.sender@mandrillapp.com","tags":["webhook-example"],"opens":[],"clicks":[],"state":"rejected","metadata":{"user_id":111},"_id":"exampleaaaaaaaaaaaaaaaaaaaaaaaaa8","_version":"exampleaaaaaaaaaaaaaaa"},"_id":"exampleaaaaaaaaaaaaaaaaaaaaaaaaa8","ts":1580762074},{"type":"blacklist","action":"add","reject":{"reason":"hard-bounce","detail":"Example detail","last_event_at":"2014-02-01 12:43:56","email":"example.webhook@mandrillapp.com","created_at":"2014-01-15 11:32:19","expires_at":"2020-04-02 12:09:18","expired":false,"subaccount":"example_subaccount","sender":"example.sender@mandrillapp.com"},"ts":1580762074},{"type":"blacklist","action":"change","reject":{"reason":"hard-bounce","detail":"Example detail","last_event_at":"2014-02-01 12:43:56","email":"example.webhook@mandrillapp.com","created_at":"2014-01-15 11:32:19","expires_at":"2020-04-02 12:09:18","expired":false,"subaccount":"example_subaccount","sender":"example.sender@mandrillapp.com"},"ts":1580762074},{"type":"blacklist","action":"remove","reject":{"reason":"hard-bounce","detail":"Example detail","last_event_at":"2014-02-01 12:43:56","email":"example.webhook@mandrillapp.com","created_at":"2014-01-15 11:32:19","expires_at":"2020-04-02 12:09:18","expired":false,"subaccount":"example_subaccount","sender":"example.sender@mandrillapp.com"},"ts":1580762074},{"type":"whitelist","action":"add","entry":{"email":"example.webhook@mandrillapp.com","detail":"example details","created_at":"2014-01-15 12:03:19"},"ts":1580762074},{"type":"whitelist","action":"remove","entry":{"email":"example.webhook@mandrillapp.com","detail":"example details","created_at":"2014-01-15 12:03:19"},"ts":1580762074}]`
var testExampleBody = `mandrill_events=%5B%7B%22event%22%3A%22send%22%2C%22msg%22%3A%7B%22ts%22%3A1365109999%2C%22subject%22%3A%22This+an+example+webhook+message%22%2C%22email%22%3A%22example.webhook%40mandrillapp.com%22%2C%22sender%22%3A%22example.sender%40mandrillapp.com%22%2C%22tags%22%3A%5B%22webhook-example%22%5D%2C%22opens%22%3A%5B%5D%2C%22clicks%22%3A%5B%5D%2C%22state%22%3A%22sent%22%2C%22metadata%22%3A%7B%22user_id%22%3A111%7D%2C%22_id%22%3A%22exampleaaaaaaaaaaaaaaaaaaaaaaaaa%22%2C%22_version%22%3A%22exampleaaaaaaaaaaaaaaa%22%7D%2C%22_id%22%3A%22exampleaaaaaaaaaaaaaaaaaaaaaaaaa%22%2C%22ts%22%3A1580765707%7D%2C%7B%22event%22%3A%22deferral%22%2C%22msg%22%3A%7B%22ts%22%3A1365109999%2C%22subject%22%3A%22This+an+example+webhook+message%22%2C%22email%22%3A%22example.webhook%40mandrillapp.com%22%2C%22sender%22%3A%22example.sender%40mandrillapp.com%22%2C%22tags%22%3A%5B%22webhook-example%22%5D%2C%22opens%22%3A%5B%5D%2C%22clicks%22%3A%5B%5D%2C%22state%22%3A%22deferred%22%2C%22metadata%22%3A%7B%22user_id%22%3A111%7D%2C%22_id%22%3A%22exampleaaaaaaaaaaaaaaaaaaaaaaaaa1%22%2C%22_version%22%3A%22exampleaaaaaaaaaaaaaaa%22%2C%22smtp_events%22%3A%5B%7B%22destination_ip%22%3A%22127.0.0.1%22%2C%22diag%22%3A%22451+4.3.5+Temporarily+unavailable%2C+try+again+later.%22%2C%22source_ip%22%3A%22127.0.0.1%22%2C%22ts%22%3A1365111111%2C%22type%22%3A%22deferred%22%2C%22size%22%3A0%7D%5D%7D%2C%22_id%22%3A%22exampleaaaaaaaaaaaaaaaaaaaaaaaaa1%22%2C%22ts%22%3A1580765707%7D%2C%7B%22event%22%3A%22hard_bounce%22%2C%22msg%22%3A%7B%22ts%22%3A1365109999%2C%22subject%22%3A%22This+an+example+webhook+message%22%2C%22email%22%3A%22example.webhook%40mandrillapp.com%22%2C%22sender%22%3A%22example.sender%40mandrillapp.com%22%2C%22tags%22%3A%5B%22webhook-example%22%5D%2C%22state%22%3A%22bounced%22%2C%22metadata%22%3A%7B%22user_id%22%3A111%7D%2C%22_id%22%3A%22exampleaaaaaaaaaaaaaaaaaaaaaaaaa2%22%2C%22_version%22%3A%22exampleaaaaaaaaaaaaaaa%22%2C%22bounce_description%22%3A%22bad_mailbox%22%2C%22bgtools_code%22%3A10%2C%22diag%22%3A%22smtp%3B550+5.1.1+The+email+account+that+you+tried+to+reach+does+not+exist.+Please+try+double-checking+the+recipient%27s+email+address+for+typos+or+unnecessary+spaces.%22%7D%2C%22_id%22%3A%22exampleaaaaaaaaaaaaaaaaaaaaaaaaa2%22%2C%22ts%22%3A1580765707%7D%2C%7B%22event%22%3A%22soft_bounce%22%2C%22msg%22%3A%7B%22ts%22%3A1365109999%2C%22subject%22%3A%22This+an+example+webhook+message%22%2C%22email%22%3A%22example.webhook%40mandrillapp.com%22%2C%22sender%22%3A%22example.sender%40mandrillapp.com%22%2C%22tags%22%3A%5B%22webhook-example%22%5D%2C%22state%22%3A%22soft-bounced%22%2C%22metadata%22%3A%7B%22user_id%22%3A111%7D%2C%22_id%22%3A%22exampleaaaaaaaaaaaaaaaaaaaaaaaaa3%22%2C%22_version%22%3A%22exampleaaaaaaaaaaaaaaa%22%2C%22bounce_description%22%3A%22mailbox_full%22%2C%22bgtools_code%22%3A22%2C%22diag%22%3A%22smtp%3B552+5.2.2+Over+Quota%22%7D%2C%22_id%22%3A%22exampleaaaaaaaaaaaaaaaaaaaaaaaaa3%22%2C%22ts%22%3A1580765707%7D%2C%7B%22event%22%3A%22open%22%2C%22msg%22%3A%7B%22ts%22%3A1365109999%2C%22subject%22%3A%22This+an+example+webhook+message%22%2C%22email%22%3A%22example.webhook%40mandrillapp.com%22%2C%22sender%22%3A%22example.sender%40mandrillapp.com%22%2C%22tags%22%3A%5B%22webhook-example%22%5D%2C%22opens%22%3A%5B%7B%22ts%22%3A1365111111%7D%5D%2C%22clicks%22%3A%5B%7B%22ts%22%3A1365111111%2C%22url%22%3A%22http%3A%5C%2F%5C%2Fmandrill.com%22%7D%5D%2C%22state%22%3A%22sent%22%2C%22metadata%22%3A%7B%22user_id%22%3A111%7D%2C%22_id%22%3A%22exampleaaaaaaaaaaaaaaaaaaaaaaaaa4%22%2C%22_version%22%3A%22exampleaaaaaaaaaaaaaaa%22%7D%2C%22_id%22%3A%22exampleaaaaaaaaaaaaaaaaaaaaaaaaa4%22%2C%22ip%22%3A%22127.0.0.1%22%2C%22location%22%3A%7B%22country_short%22%3A%22US%22%2C%22country%22%3A%22United+States%22%2C%22region%22%3A%22Oklahoma%22%2C%22city%22%3A%22Oklahoma+City%22%2C%22latitude%22%3A35.4675598145%2C%22longitude%22%3A-97.5164337158%2C%22postal_code%22%3A%2273101%22%2C%22timezone%22%3A%22-05%3A00%22%7D%2C%22user_agent%22%3A%22Mozilla%5C%2F5.0+%28Macintosh%3B+U%3B+Intel+Mac+OS+X+10.6%3B+en-US%3B+rv%3A1.9.1.8%29+Gecko%5C%2F20100317+Postbox%5C%2F1.1.3%22%2C%22user_agent_parsed%22%3A%7B%22type%22%3A%22Email+Client%22%2C%22ua_family%22%3A%22Postbox%22%2C%22ua_name%22%3A%22Postbox+1.1.3%22%2C%22ua_version%22%3A%221.1.3%22%2C%22ua_url%22%3A%22http%3A%5C%2F%5C%2Fwww.postbox-inc.com%5C%2F%22%2C%22ua_company%22%3A%22Postbox%2C+Inc.%22%2C%22ua_company_url%22%3A%22http%3A%5C%2F%5C%2Fwww.postbox-inc.com%5C%2F%22%2C%22ua_icon%22%3A%22http%3A%5C%2F%5C%2Fcdn.mandrill.com%5C%2Fimg%5C%2Femail-client-icons%5C%2Fpostbox.png%22%2C%22os_family%22%3A%22OS+X%22%2C%22os_name%22%3A%22OS+X+10.6+Snow+Leopard%22%2C%22os_url%22%3A%22http%3A%5C%2F%5C%2Fwww.apple.com%5C%2Fosx%5C%2F%22%2C%22os_company%22%3A%22Apple+Computer%2C+Inc.%22%2C%22os_company_url%22%3A%22http%3A%5C%2F%5C%2Fwww.apple.com%5C%2F%22%2C%22os_icon%22%3A%22http%3A%5C%2F%5C%2Fcdn.mandrill.com%5C%2Fimg%5C%2Femail-client-icons%5C%2Fmacosx.png%22%2C%22mobile%22%3Afalse%7D%2C%22ts%22%3A1580765707%7D%2C%7B%22event%22%3A%22click%22%2C%22msg%22%3A%7B%22ts%22%3A1365109999%2C%22subject%22%3A%22This+an+example+webhook+message%22%2C%22email%22%3A%22example.webhook%40mandrillapp.com%22%2C%22sender%22%3A%22example.sender%40mandrillapp.com%22%2C%22tags%22%3A%5B%22webhook-example%22%5D%2C%22opens%22%3A%5B%7B%22ts%22%3A1365111111%7D%5D%2C%22clicks%22%3A%5B%7B%22ts%22%3A1365111111%2C%22url%22%3A%22http%3A%5C%2F%5C%2Fmandrill.com%22%7D%5D%2C%22state%22%3A%22sent%22%2C%22metadata%22%3A%7B%22user_id%22%3A111%7D%2C%22_id%22%3A%22exampleaaaaaaaaaaaaaaaaaaaaaaaaa5%22%2C%22_version%22%3A%22exampleaaaaaaaaaaaaaaa%22%7D%2C%22_id%22%3A%22exampleaaaaaaaaaaaaaaaaaaaaaaaaa5%22%2C%22ip%22%3A%22127.0.0.1%22%2C%22location%22%3A%7B%22country_short%22%3A%22US%22%2C%22country%22%3A%22United+States%22%2C%22region%22%3A%22Oklahoma%22%2C%22city%22%3A%22Oklahoma+City%22%2C%22latitude%22%3A35.4675598145%2C%22longitude%22%3A-97.5164337158%2C%22postal_code%22%3A%2273101%22%2C%22timezone%22%3A%22-05%3A00%22%7D%2C%22user_agent%22%3A%22Mozilla%5C%2F5.0+%28Macintosh%3B+U%3B+Intel+Mac+OS+X+10.6%3B+en-US%3B+rv%3A1.9.1.8%29+Gecko%5C%2F20100317+Postbox%5C%2F1.1.3%22%2C%22user_agent_parsed%22%3A%7B%22type%22%3A%22Email+Client%22%2C%22ua_family%22%3A%22Postbox%22%2C%22ua_name%22%3A%22Postbox+1.1.3%22%2C%22ua_version%22%3A%221.1.3%22%2C%22ua_url%22%3A%22http%3A%5C%2F%5C%2Fwww.postbox-inc.com%5C%2F%22%2C%22ua_company%22%3A%22Postbox%2C+Inc.%22%2C%22ua_company_url%22%3A%22http%3A%5C%2F%5C%2Fwww.postbox-inc.com%5C%2F%22%2C%22ua_icon%22%3A%22http%3A%5C%2F%5C%2Fcdn.mandrill.com%5C%2Fimg%5C%2Femail-client-icons%5C%2Fpostbox.png%22%2C%22os_family%22%3A%22OS+X%22%2C%22os_name%22%3A%22OS+X+10.6+Snow+Leopard%22%2C%22os_url%22%3A%22http%3A%5C%2F%5C%2Fwww.apple.com%5C%2Fosx%5C%2F%22%2C%22os_company%22%3A%22Apple+Computer%2C+Inc.%22%2C%22os_company_url%22%3A%22http%3A%5C%2F%5C%2Fwww.apple.com%5C%2F%22%2C%22os_icon%22%3A%22http%3A%5C%2F%5C%2Fcdn.mandrill.com%5C%2Fimg%5C%2Femail-client-icons%5C%2Fmacosx.png%22%2C%22mobile%22%3Afalse%7D%2C%22url%22%3A%22http%3A%5C%2F%5C%2Fmandrill.com%22%2C%22ts%22%3A1580765707%7D%2C%7B%22event%22%3A%22spam%22%2C%22msg%22%3A%7B%22ts%22%3A1365109999%2C%22subject%22%3A%22This+an+example+webhook+message%22%2C%22email%22%3A%22example.webhook%40mandrillapp.com%22%2C%22sender%22%3A%22example.sender%40mandrillapp.com%22%2C%22tags%22%3A%5B%22webhook-example%22%5D%2C%22opens%22%3A%5B%7B%22ts%22%3A1365111111%7D%5D%2C%22clicks%22%3A%5B%7B%22ts%22%3A1365111111%2C%22url%22%3A%22http%3A%5C%2F%5C%2Fmandrill.com%22%7D%5D%2C%22state%22%3A%22sent%22%2C%22metadata%22%3A%7B%22user_id%22%3A111%7D%2C%22_id%22%3A%22exampleaaaaaaaaaaaaaaaaaaaaaaaaa6%22%2C%22_version%22%3A%22exampleaaaaaaaaaaaaaaa%22%7D%2C%22_id%22%3A%22exampleaaaaaaaaaaaaaaaaaaaaaaaaa6%22%2C%22ts%22%3A1580765707%7D%2C%7B%22event%22%3A%22unsub%22%2C%22msg%22%3A%7B%22ts%22%3A1365109999%2C%22subject%22%3A%22This+an+example+webhook+message%22%2C%22email%22%3A%22example.webhook%40mandrillapp.com%22%2C%22sender%22%3A%22example.sender%40mandrillapp.com%22%2C%22tags%22%3A%5B%22webhook-example%22%5D%2C%22opens%22%3A%5B%7B%22ts%22%3A1365111111%7D%5D%2C%22clicks%22%3A%5B%7B%22ts%22%3A1365111111%2C%22url%22%3A%22http%3A%5C%2F%5C%2Fmandrill.com%22%7D%5D%2C%22state%22%3A%22sent%22%2C%22metadata%22%3A%7B%22user_id%22%3A111%7D%2C%22_id%22%3A%22exampleaaaaaaaaaaaaaaaaaaaaaaaaa7%22%2C%22_version%22%3A%22exampleaaaaaaaaaaaaaaa%22%7D%2C%22_id%22%3A%22exampleaaaaaaaaaaaaaaaaaaaaaaaaa7%22%2C%22ts%22%3A1580765707%7D%2C%7B%22event%22%3A%22reject%22%2C%22msg%22%3A%7B%22ts%22%3A1365109999%2C%22subject%22%3A%22This+an+example+webhook+message%22%2C%22email%22%3A%22example.webhook%40mandrillapp.com%22%2C%22sender%22%3A%22example.sender%40mandrillapp.com%22%2C%22tags%22%3A%5B%22webhook-example%22%5D%2C%22opens%22%3A%5B%5D%2C%22clicks%22%3A%5B%5D%2C%22state%22%3A%22rejected%22%2C%22metadata%22%3A%7B%22user_id%22%3A111%7D%2C%22_id%22%3A%22exampleaaaaaaaaaaaaaaaaaaaaaaaaa8%22%2C%22_version%22%3A%22exampleaaaaaaaaaaaaaaa%22%7D%2C%22_id%22%3A%22exampleaaaaaaaaaaaaaaaaaaaaaaaaa8%22%2C%22ts%22%3A1580765707%7D%2C%7B%22type%22%3A%22blacklist%22%2C%22action%22%3A%22add%22%2C%22reject%22%3A%7B%22reason%22%3A%22hard-bounce%22%2C%22detail%22%3A%22Example+detail%22%2C%22last_event_at%22%3A%222014-02-01+12%3A43%3A56%22%2C%22email%22%3A%22example.webhook%40mandrillapp.com%22%2C%22created_at%22%3A%222014-01-15+11%3A32%3A19%22%2C%22expires_at%22%3A%222020-04-02+12%3A09%3A18%22%2C%22expired%22%3Afalse%2C%22subaccount%22%3A%22example_subaccount%22%2C%22sender%22%3A%22example.sender%40mandrillapp.com%22%7D%2C%22ts%22%3A1580765707%7D%2C%7B%22type%22%3A%22blacklist%22%2C%22action%22%3A%22change%22%2C%22reject%22%3A%7B%22reason%22%3A%22hard-bounce%22%2C%22detail%22%3A%22Example+detail%22%2C%22last_event_at%22%3A%222014-02-01+12%3A43%3A56%22%2C%22email%22%3A%22example.webhook%40mandrillapp.com%22%2C%22created_at%22%3A%222014-01-15+11%3A32%3A19%22%2C%22expires_at%22%3A%222020-04-02+12%3A09%3A18%22%2C%22expired%22%3Afalse%2C%22subaccount%22%3A%22example_subaccount%22%2C%22sender%22%3A%22example.sender%40mandrillapp.com%22%7D%2C%22ts%22%3A1580765707%7D%2C%7B%22type%22%3A%22blacklist%22%2C%22action%22%3A%22remove%22%2C%22reject%22%3A%7B%22reason%22%3A%22hard-bounce%22%2C%22detail%22%3A%22Example+detail%22%2C%22last_event_at%22%3A%222014-02-01+12%3A43%3A56%22%2C%22email%22%3A%22example.webhook%40mandrillapp.com%22%2C%22created_at%22%3A%222014-01-15+11%3A32%3A19%22%2C%22expires_at%22%3A%222020-04-02+12%3A09%3A18%22%2C%22expired%22%3Afalse%2C%22subaccount%22%3A%22example_subaccount%22%2C%22sender%22%3A%22example.sender%40mandrillapp.com%22%7D%2C%22ts%22%3A1580765707%7D%2C%7B%22type%22%3A%22whitelist%22%2C%22action%22%3A%22add%22%2C%22entry%22%3A%7B%22email%22%3A%22example.webhook%40mandrillapp.com%22%2C%22detail%22%3A%22example+details%22%2C%22created_at%22%3A%222014-01-15+12%3A03%3A19%22%7D%2C%22ts%22%3A1580765707%7D%2C%7B%22type%22%3A%22whitelist%22%2C%22action%22%3A%22remove%22%2C%22entry%22%3A%7B%22email%22%3A%22example.webhook%40mandrillapp.com%22%2C%22detail%22%3A%22example+details%22%2C%22created_at%22%3A%222014-01-15+12%3A03%3A19%22%7D%2C%22ts%22%3A1580765707%7D%5D`

func TestGenerateWebhookSignature(t *testing.T) {
	validSig := "LoesQYS9AbSuvfAZF3LSSprv9xY="
	exampleRequest, _ := http.NewRequest("POST", "https://enafnh68w1h3o.x.pipedream.net/", strings.NewReader(testExampleBody))
	exampleRequest.Header.Add("X-MANDRILL-SIGNATURE", validSig)
	exampleRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	tests := map[string]struct {
		key string
		url string
		req http.Request
		err error
	}{
		"basic": {
			key: "o0j4nmpKdMLp2-mRejxAog",
			url: "https://enafnh68w1h3o.x.pipedream.net/",
			req: *exampleRequest,
			err: nil,
		},
		"invalid_signature": {
			key: "abcdefg",
			url: "https://enafnh68w1h3o.x.pipedream.net/",
			req: func() http.Request {
				req, _ := http.NewRequest("POST", "https://enafnh68w1h3o.x.pipedream.net/", strings.NewReader(testExampleBody))
				req.Header.Add("X-MANDRILL-SIGNATURE", validSig)
				req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
				return *req
			}(),
			err: &ErrInvalidSignature{},
		},
		"missing_events": {
			key: "o0j4nmpKdMLp2-mRejxAog",
			url: "https://enafnh68w1h3o.x.pipedream.net/",
			req: func() http.Request {
				req, _ := http.NewRequest("POST", "https://enafnh68w1h3o.x.pipedream.net/", strings.NewReader("foo"))
				req.Header.Add("X-MANDRILL-SIGNATURE", validSig)
				return *req
			}(),
			err: &ErrMissingMandrillEvents{},
		},
		"missing_signature": {
			key: "o0j4nmpKdMLp2-mRejxAog",
			url: "https://enafnh68w1h3o.x.pipedream.net/",
			req: func() http.Request {
				req, _ := http.NewRequest("POST", "https://enafnh68w1h3o.x.pipedream.net/", strings.NewReader("foo"))
				return *req
			}(),
			err: &ErrNoSignature{},
		},
		"parse_form_error": {
			key: "o0j4nmpKdMLp2-mRejxAog",
			url: "https://enafnh68w1h3o.x.pipedream.net/",
			req: func() http.Request {
				req, _ := http.NewRequest("POST", "https://enafnh68w1h3o.x.pipedream.net/", nil)
				req.Header.Add("X-MANDRILL-SIGNATURE", validSig)
				return *req
			}(),
			err: &ErrParse{},
		},
		"not_post": {
			key: "o0j4nmpKdMLp2-mRejxAog",
			url: "https://enafnh68w1h3o.x.pipedream.net/",
			req: func() http.Request {
				req, _ := http.NewRequest("GET", "https://enafnh68w1h3o.x.pipedream.net/", nil)
				req.Header.Add("X-MANDRILL-SIGNATURE", validSig)
				return *req
			}(),
			err: &ErrNotPost{},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			err := ValidateWebhookSignature(tc.url, tc.key, &tc.req)
			if tc.err == nil {
				require.NoError(t, err)
			} else {
				require.IsType(t, tc.err, err)
			}
		})
	}
}

func TestParseEvent(t *testing.T) {
	data := []json.RawMessage{}
	err := json.Unmarshal([]byte(testExampleJSON), &data)
	require.NoError(t, err)
	for _, j := range data {
		res, err := ParseEvent(j)
		require.NoError(t, err)
		require.NotNil(t, res)
	}
}