package events

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseMessageEventHappyPath(t *testing.T) {
	testCases := map[string]struct {
		data string
		err  error
	}{
		"send": {
			data: `{
				"event": "send",
				"msg": {
					"ts": 1365109999,
					"subject": "This an example webhook message",
					"email": "example.webhook@mandrillapp.com",
					"sender": "example.sender@mandrillapp.com",
					"tags": ["webhook-example"],
					"opens": [],
					"clicks": [],
					"state": "sent",
					"metadata": {
						"user_id": 111
					},
					"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa",
					"_version": "exampleaaaaaaaaaaaaaaa"
				},
				"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa",
				"ts": 1580762074
			}`,
			err: nil,
		},
		"deferral": {
			data: `{
					"event": "deferral",
					"msg": {
						"ts": 1365109999,
						"subject": "This an example webhook message",
						"email": "example.webhook@mandrillapp.com",
						"sender": "example.sender@mandrillapp.com",
						"tags": ["webhook-example"],
						"opens": [],
						"clicks": [],
						"state": "deferred",
						"metadata": {
							"user_id": 111
						},
						"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa1",
						"_version": "exampleaaaaaaaaaaaaaaa",
						"smtp_events": [{
							"destination_ip": "127.0.0.1",
							"diag": "451 4.3.5 Temporarily unavailable, try again later.",
							"source_ip": "127.0.0.1",
							"ts": 1365111111,
							"type": "deferred",
							"size": 0
						}]
					},
					"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa1",
					"ts": 1580762074
				}`,
			err: nil,
		},
		"hard_bounce": {
			data: `{
					"event": "hard_bounce",
					"msg": {
						"ts": 1365109999,
						"subject": "This an example webhook message",
						"email": "example.webhook@mandrillapp.com",
						"sender": "example.sender@mandrillapp.com",
						"tags": ["webhook-example"],
						"state": "bounced",
						"metadata": {
							"user_id": 111
						},
						"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa2",
						"_version": "exampleaaaaaaaaaaaaaaa",
						"bounce_description": "bad_mailbox",
						"bgtools_code": 10,
						"diag": "smtp;550 5.1.1 The email account that you tried to reach does not exist. Please try double-checking the recipient's email address for typos or unnecessary spaces."
					},
					"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa2",
					"ts": 1580762074
				}`,
			err: nil,
		},
		"soft_bounce": {
			data: `{
					"event": "soft_bounce",
					"msg": {
						"ts": 1365109999,
						"subject": "This an example webhook message",
						"email": "example.webhook@mandrillapp.com",
						"sender": "example.sender@mandrillapp.com",
						"tags": ["webhook-example"],
						"state": "soft-bounced",
						"metadata": {
							"user_id": 111
						},
						"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa3",
						"_version": "exampleaaaaaaaaaaaaaaa",
						"bounce_description": "mailbox_full",
						"bgtools_code": 22,
						"diag": "smtp;552 5.2.2 Over Quota"
					},
					"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa3",
					"ts": 1580762074
				}`,
			err: nil,
		},
		"open": {
			data: `{
					"event": "open",
					"msg": {
						"ts": 1365109999,
						"subject": "This an example webhook message",
						"email": "example.webhook@mandrillapp.com",
						"sender": "example.sender@mandrillapp.com",
						"tags": ["webhook-example"],
						"opens": [{
							"ts": 1365111111
						}],
						"clicks": [{
							"ts": 1365111111,
							"url": "http:\/\/mandrill.com"
						}],
						"state": "sent",
						"metadata": {
							"user_id": 111
						},
						"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa4",
						"_version": "exampleaaaaaaaaaaaaaaa"
					},
					"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa4",
					"ip": "127.0.0.1",
					"location": {
						"country_short": "US",
						"country": "United States",
						"region": "Oklahoma",
						"city": "Oklahoma City",
						"latitude": 35.4675598145,
						"longitude": -97.5164337158,
						"postal_code": "73101",
						"timezone": "-05:00"
					},
					"user_agent": "Mozilla\/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.1.8) Gecko\/20100317 Postbox\/1.1.3",
					"user_agent_parsed": {
						"type": "Email Client",
						"ua_family": "Postbox",
						"ua_name": "Postbox 1.1.3",
						"ua_version": "1.1.3",
						"ua_url": "http:\/\/www.postbox-inc.com\/",
						"ua_company": "Postbox, Inc.",
						"ua_company_url": "http:\/\/www.postbox-inc.com\/",
						"ua_icon": "http:\/\/cdn.mandrill.com\/img\/email-client-icons\/postbox.png",
						"os_family": "OS X",
						"os_name": "OS X 10.6 Snow Leopard",
						"os_url": "http:\/\/www.apple.com\/osx\/",
						"os_company": "Apple Computer, Inc.",
						"os_company_url": "http:\/\/www.apple.com\/",
						"os_icon": "http:\/\/cdn.mandrill.com\/img\/email-client-icons\/macosx.png",
						"mobile": false
					},
					"ts": 1580762074
				}`,
			err: nil,
		},

		"click": {
			data: `{
					"event": "click",
					"msg": {
						"ts": 1365109999,
						"subject": "This an example webhook message",
						"email": "example.webhook@mandrillapp.com",
						"sender": "example.sender@mandrillapp.com",
						"tags": ["webhook-example"],
						"opens": [{
							"ts": 1365111111
						}],
						"clicks": [{
							"ts": 1365111111,
							"url": "http:\/\/mandrill.com"
						}],
						"state": "sent",
						"metadata": {
							"user_id": 111
						},
						"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa5",
						"_version": "exampleaaaaaaaaaaaaaaa"
					},
					"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa5",
					"ip": "127.0.0.1",
					"location": {
						"country_short": "US",
						"country": "United States",
						"region": "Oklahoma",
						"city": "Oklahoma City",
						"latitude": 35.4675598145,
						"longitude": -97.5164337158,
						"postal_code": "73101",
						"timezone": "-05:00"
					},
					"user_agent": "Mozilla\/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.1.8) Gecko\/20100317 Postbox\/1.1.3",
					"user_agent_parsed": {
						"type": "Email Client",
						"ua_family": "Postbox",
						"ua_name": "Postbox 1.1.3",
						"ua_version": "1.1.3",
						"ua_url": "http:\/\/www.postbox-inc.com\/",
						"ua_company": "Postbox, Inc.",
						"ua_company_url": "http:\/\/www.postbox-inc.com\/",
						"ua_icon": "http:\/\/cdn.mandrill.com\/img\/email-client-icons\/postbox.png",
						"os_family": "OS X",
						"os_name": "OS X 10.6 Snow Leopard",
						"os_url": "http:\/\/www.apple.com\/osx\/",
						"os_company": "Apple Computer, Inc.",
						"os_company_url": "http:\/\/www.apple.com\/",
						"os_icon": "http:\/\/cdn.mandrill.com\/img\/email-client-icons\/macosx.png",
						"mobile": false
					},
					"url": "http:\/\/mandrill.com",
					"ts": 1580762074
				}`,
			err: nil,
		},
		"spam": {
			data: `{
					"event": "spam",
					"msg": {
						"ts": 1365109999,
						"subject": "This an example webhook message",
						"email": "example.webhook@mandrillapp.com",
						"sender": "example.sender@mandrillapp.com",
						"tags": ["webhook-example"],
						"opens": [{
							"ts": 1365111111
						}],
						"clicks": [{
							"ts": 1365111111,
							"url": "http:\/\/mandrill.com"
						}],
						"state": "sent",
						"metadata": {
							"user_id": 111
						},
						"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa6",
						"_version": "exampleaaaaaaaaaaaaaaa"
					},
					"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa6",
					"ts": 1580762074
				}`,
			err: nil,
		},
		"unsub": {
			data: `{
					"event": "unsub",
					"msg": {
						"ts": 1365109999,
						"subject": "This an example webhook message",
						"email": "example.webhook@mandrillapp.com",
						"sender": "example.sender@mandrillapp.com",
						"tags": ["webhook-example"],
						"opens": [{
							"ts": 1365111111
						}],
						"clicks": [{
							"ts": 1365111111,
							"url": "http:\/\/mandrill.com"
						}],
						"state": "sent",
						"metadata": {
							"user_id": 111
						},
						"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa7",
						"_version": "exampleaaaaaaaaaaaaaaa"
					},
					"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa7",
					"ts": 1580762074
				}`,
			err: nil,
		},
		"reject": {
			data: `{
					"event": "reject",
					"msg": {
						"ts": 1365109999,
						"subject": "This an example webhook message",
						"email": "example.webhook@mandrillapp.com",
						"sender": "example.sender@mandrillapp.com",
						"tags": ["webhook-example"],
						"opens": [],
						"clicks": [],
						"state": "rejected",
						"metadata": {
							"user_id": 111
						},
						"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa8",
						"_version": "exampleaaaaaaaaaaaaaaa"
					},
					"_id": "exampleaaaaaaaaaaaaaaaaaaaaaaaaa8",
					"ts": 1580762074
				}`,
			err: nil,
		},
		"inbound": {
			data: `{
				"event": "inbound",
				"msg": {
					"dkim": {
						"signed": true,
						"valid": true
					},
					"email": "foo@foo.com",
					"from_email": "example.sender@mandrillapp.com",
					"headers": {
						"Content-Type": "multipart\/alternative; boundary=\"_av-7r7zDhHxVEAo2yMWasfuFw\"",
						"Date": "Fri, 10 May 2013 19:28:20 0000",
						"Dkim-Signature": ["v=1; a=rsa-sha1; c=relaxed\/relaxed; s=mandrill; d=mail115.us4.mandrillapp.com; h=From:Sender:Subject:List-Unsubscribe:To:Message-Id:Date:MIME-Version:Content-Type; i=example.sender@mail115.us4.mandrillapp.com; bh=d60x72jf42gLILD7IiqBL0OBb40=; b=iJd7eBugdIjzqW84UZ2xynlg1SojANJR6xfz0JDD44h78EpbqJiYVcMIfRG7mkrn741Bd5YaMR6p 9j41OA9A5am 8BE1Ng2kLRGwou5hRInn xXBAQm2NUt5FkoXSpvm4gC4gQSOxPbQcuzlLha9JqxJ 8ZZ89\/20txUrRq9cYxk=", "v=1; a=rsa-sha256; c=relaxed\/relaxed; d=c.mandrillapp.com; i=@c.mandrillapp.com; q=dns\/txt; s=mandrill; t=1368214100; h=From : Sender : Subject : List-Unsubscribe : To : Message-Id : Date : MIME-Version : Content-Type : From : Subject : Date : X-Mandrill-User : List-Unsubscribe; bh=y5Vz RDcKZmWzRc9s0xUJR6k4APvBNktBqy1EhSWM8o=; b=PLAUIuw8zk8kG5tPkmcnSanElxt6I5lp5t32nSvzVQE7R8u0AmIEjeIDZEt31 Q9PWt nY sHHRsXUQ9SZpndT9Bk \/SmyA2ntU\/2AKuqDpPkMZiTqxmGF80Wz4JJgx2aCEB1LeLVmFFwB 5Nr\/LBSlsBlRcjT9QiWw0\/iRvCn74="],
						"Domainkey-Signature": "a=rsa-sha1; c=nofws; q=dns; s=mandrill; d=mail115.us4.mandrillapp.com; b=X6qudHd4oOJvVQZcoAEUCJgB875SwzEO5UKf6NvpfqyCVjdaO79WdDulLlfNVELeuoK2m6alt2yw 5Qhp4TW5NegyFf6Ogr\/Hy0Lt411r\/0lRf0nyaVkqMM\/9g13B6D9CS092v70wshX8 qdyxK8fADw8 kEelbCK2cEl0AGIeAeo=;",
						"From": "<example.sender@mandrillapp.com>",
						"List-Unsubscribe": "<mailto:unsubscribe-md_999.aaaaaaaa.v1-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@mailin1.us2.mcsv.net?subject=unsub>",
						"Message-Id": "<999.20130510192820.aaaaaaaaaaaaaa.aaaaaaaa@mail115.us4.mandrillapp.com>",
						"Mime-Version": "1.0",
						"Received": ["from mail115.us4.mandrillapp.com (mail115.us4.mandrillapp.com [205.201.136.115]) by mail.example.com (Postfix) with ESMTP id AAAAAAAAAAA for <foo@foo.com>; Fri, 10 May 2013 19:28:21 0000 (UTC)", "from localhost (127.0.0.1) by mail115.us4.mandrillapp.com id hhl55a14i282 for <foo@foo.com>; Fri, 10 May 2013 19:28:20 0000 (envelope-from <bounce-md_999.aaaaaaaa.v1-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@mail115.us4.mandrillapp.com>)"],
						"Sender": "<example.sender@mail115.us4.mandrillapp.com>",
						"Subject": "This is an example webhook message",
						"To": "<foo@foo.com>",
						"X-Report-Abuse": "Please forward a copy of this message, including all headers, to abuse@mandrill.com"
					},
					"html": "<p>This is an example inbound message.<\/p><img src=\"http:\/\/mandrillapp.com\/track\/open.php?u=999&id=aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa&tags=_all,_sendexample.sender@mandrillapp.com\" height=\"1\" width=\"1\">\n",
					"raw_msg": "Received: from mail115.us4.mandrillapp.com (mail115.us4.mandrillapp.com [205.201.136.115])\n\tby mail.example.com (Postfix) with ESMTP id AAAAAAAAAAA\n\tfor <foo@foo.com>; Fri, 10 May 2013 19:28:20 0000 (UTC)\nDKIM-Signature: v=1; a=rsa-sha1; c=relaxed\/relaxed; s=mandrill; d=mail115.us4.mandrillapp.com;\n h=From:Sender:Subject:List-Unsubscribe:To:Message-Id:Date:MIME-Version:Content-Type; i=example.sender@mail115.us4.mandrillapp.com;\n bh=d60x72jf42gLILD7IiqBL0OBb40=;\n b=iJd7eBugdIjzqW84UZ2xynlg1SojANJR6xfz0JDD44h78EpbqJiYVcMIfRG7mkrn741Bd5YaMR6p\n 9j41OA9A5am 8BE1Ng2kLRGwou5hRInn xXBAQm2NUt5FkoXSpvm4gC4gQSOxPbQcuzlLha9JqxJ\n 8ZZ89\/20txUrRq9cYxk=\nDomainKey-Signature: a=rsa-sha1; c=nofws; q=dns; s=mandrill; d=mail115.us4.mandrillapp.com;\n b=X6qudHd4oOJvVQZcoAEUCJgB875SwzEO5UKf6NvpfqyCVjdaO79WdDulLlfNVELeuoK2m6alt2yw\n 5Qhp4TW5NegyFf6Ogr\/Hy0Lt411r\/0lRf0nyaVkqMM\/9g13B6D9CS092v70wshX8 qdyxK8fADw8\n kEelbCK2cEl0AGIeAeo=;\nReceived: from localhost (127.0.0.1) by mail115.us4.mandrillapp.com id hhl55a14i282 for <foo@foo.com>; Fri, 10 May 2013 19:28:20 0000 (envelope-from <bounce-md_999.aaaaaaaa.v1-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@mail115.us4.mandrillapp.com>)\nDKIM-Signature: v=1; a=rsa-sha256; c=relaxed\/relaxed; d=c.mandrillapp.com; \n i=@c.mandrillapp.com; q=dns\/txt; s=mandrill; t=1368214100; h=From : \n Sender : Subject : List-Unsubscribe : To : Message-Id : Date : \n MIME-Version : Content-Type : From : Subject : Date : X-Mandrill-User : \n List-Unsubscribe; bh=y5Vz RDcKZmWzRc9s0xUJR6k4APvBNktBqy1EhSWM8o=; \n b=PLAUIuw8zk8kG5tPkmcnSanElxt6I5lp5t32nSvzVQE7R8u0AmIEjeIDZEt31 Q9PWt nY\n sHHRsXUQ9SZpndT9Bk \/SmyA2ntU\/2AKuqDpPkMZiTqxmGF80Wz4JJgx2aCEB1LeLVmFFwB\n 5Nr\/LBSlsBlRcjT9QiWw0\/iRvCn74=\nFrom: <example.sender@mandrillapp.com>\nSender: <example.sender@mail115.us4.mandrillapp.com>\nSubject: This is an example webhook message\nList-Unsubscribe: <mailto:unsubscribe-md_999.aaaaaaaa.v1-aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@mailin1.us2.mcsv.net?subject=unsub>\nTo: <foo@foo.com>\nX-Report-Abuse: Please forward a copy of this message, including all headers, to abuse@mandrill.com\nX-Mandrill-User: md_999\nMessage-Id: <999.20130510192820.aaaaaaaaaaaaaa.aaaaaaaa@mail115.us4.mandrillapp.com>\nDate: Fri, 10 May 2013 19:28:20 0000\nMIME-Version: 1.0\nContent-Type: multipart\/alternative; boundary=\"_av-7r7zDhHxVEAo2yMWasfuFw\"\n\n--_av-7r7zDhHxVEAo2yMWasfuFw\nContent-Type: text\/plain; charset=utf-8\nContent-Transfer-Encoding: 7bit\n\nThis is an example inbound message.\n--_av-7r7zDhHxVEAo2yMWasfuFw\nContent-Type: text\/html; charset=utf-8\nContent-Transfer-Encoding: 7bit\n\n<p>This is an example inbound message.<\/p><img src=\"http:\/\/mandrillapp.com\/track\/open.php?u=999&id=aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa&tags=_all,_sendexample.sender@mandrillapp.com\" height=\"1\" width=\"1\">\n--_av-7r7zDhHxVEAo2yMWasfuFw--",
					"sender": null,
					"spam_report": {
						"matched_rules": [{
							"description": "RBL: Sender listed at http:\/\/www.dnswl.org\/, no",
							"name": "RCVD_IN_DNSWL_NONE",
							"score": 0
						}, {
							"description": null,
							"name": null,
							"score": 0
						}, {
							"description": "in iadb.isipp.com]",
							"name": "listed",
							"score": 0
						}, {
							"description": "RBL: Participates in the IADB system",
							"name": "RCVD_IN_IADB_LISTED",
							"score": -0.4
						}, {
							"description": "RBL: ISIPP IADB lists as vouched-for sender",
							"name": "RCVD_IN_IADB_VOUCHED",
							"score": -2.2
						}, {
							"description": "RBL: IADB: Sender publishes SPF record",
							"name": "RCVD_IN_IADB_SPF",
							"score": 0
						}, {
							"description": "RBL: IADB: Sender publishes Sender ID record",
							"name": "RCVD_IN_IADB_SENDERID",
							"score": 0
						}, {
							"description": "RBL: IADB: Sender publishes Domain Keys record",
							"name": "RCVD_IN_IADB_DK",
							"score": -0.2
						}, {
							"description": "RBL: IADB: Sender has reverse DNS record",
							"name": "RCVD_IN_IADB_RDNS",
							"score": -0.2
						}, {
							"description": "SPF: HELO matches SPF record",
							"name": "SPF_HELO_PASS",
							"score": 0
						}, {
							"description": "BODY: HTML included in message",
							"name": "HTML_MESSAGE",
							"score": 0
						}, {
							"description": "BODY: HTML: images with 0-400 bytes of words",
							"name": "HTML_IMAGE_ONLY_04",
							"score": 0.3
						}, {
							"description": "Message has a DKIM or DK signature, not necessarily valid",
							"name": "DKIM_SIGNED",
							"score": 0.1
						}, {
							"description": "Message has at least one valid DKIM or DK signature",
							"name": "DKIM_VALID",
							"score": -0.1
						}],
						"score": -2.6
					},
					"spf": {
						"detail": "sender SPF authorized",
						"result": "pass"
					},
					"subject": "This is an example webhook message",
					"tags": [],
					"template": null,
					"text": "This is an example inbound message.\n",
					"text_flowed": false,
					"to": [
						["foo@foo.com", null]
					]
				},
				"ts": 1368214102
			}
			`,
			err: nil,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			oe := WebhookEvent{
				Type:           MessageEventType,
				InnerEventType: name,
				raw:            []byte(testCase.data),
			}
			expectedType, ok := messageEventMapping[name]
			require.True(t, ok)
			res, err := parseMessageEvent(oe)
			if testCase.err == nil {
				require.NoError(t, err)
				require.IsType(t, expectedType, res.Data)
				require.NotEmpty(t, res)
				require.NotNil(t, res)
			} else {
				require.Error(t, err)
				require.IsType(t, testCase.err, err)
			}
		})
	}
}

// these are test data that I've seen as actual webhooks from Mandrill
// with data anonymized
func TestSeenMessageEvents(t *testing.T) {
	testCases := map[string]struct {
		data         string
		inner        string
		expectedType interface{}
	}{
		"send_1": {
			data:         `{"event":"send","_id":"abcdefg","msg":{"ts":1580940796,"subject":"Your report is complete","email":"recipient@foo.com","tags":[],"opens":[],"clicks":[],"state":"sent","smtp_events":[],"subaccount":null,"resends":[],"reject":null,"_id":"abcdefg","sender":"foo@foo.com","template":"XXXXXX"},"ts":1580940796}`,
			inner:        SendEventType,
			expectedType: SendEvent{},
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			oe := WebhookEvent{
				Type:           MessageEventType,
				InnerEventType: testCase.inner,
				raw:            []byte(testCase.data),
			}
			_, ok := messageEventMapping[testCase.inner]
			require.True(t, ok)
			res, err := parseMessageEvent(oe)
			require.NoError(t, err)
			require.NotNil(t, res)
			require.NotEmpty(t, res)
			require.IsType(t, testCase.expectedType, res.Data)
		})
	}
}
