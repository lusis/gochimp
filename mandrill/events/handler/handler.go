package handler

import (
	"context"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/lusis/gochimp/mandrill/events"
)

// MessageEventFunc is a function that can process an events.MessageEvent
type MessageEventFunc func(context.Context, events.MessageEvent)

// SyncEventFunc is a function that can process an events.SyncEvent
type SyncEventFunc func(context.Context, events.SyncEvent)

// UnknownEventFunc is a function that can process any events that don't match
type UnknownEventFunc func(context.Context, events.WebhookEvent)

// SendEventFunc is a function that can process send events
type SendEventFunc func(context.Context, events.SendEvent)

// DeferralEventFunc is a function that can process deferral events
type DeferralEventFunc func(context.Context, events.DeferralEvent)

// HardBounceEventFunc is a function that can process hard_bounce events
type HardBounceEventFunc func(context.Context, events.HardBounceEvent)

// SoftBounceEventFunc is a function that can process soft_bounce events
type SoftBounceEventFunc func(context.Context, events.SoftBounceEvent)

// OpenEventFunc is a function that can process open events
type OpenEventFunc func(context.Context, events.OpenEvent)

// ClickEventFunc is a function that can process click events
type ClickEventFunc func(context.Context, events.ClickEvent)

// SpamEventFunc is a function that can process spam events
type SpamEventFunc func(context.Context, events.SpamEvent)

// UnsubEventFunc is a function that can process unsub events
type UnsubEventFunc func(context.Context, events.UnsubEvent)

// RejectEventFunc is a function that can process reject events
type RejectEventFunc func(context.Context, events.RejectEvent)

// InboundEventFunc is a function that can process inbound events
type InboundEventFunc func(context.Context, events.InboundMessageEvent)

// BlacklistEventFunc is a function that can process blacklist sync events
type BlacklistEventFunc func(context.Context, events.BlacklistEvent)

// WhitelistEventFunc is a function that can process whitelist sync events
type WhitelistEventFunc func(context.Context, events.WhitelistEvent)

// Option is a functional type for setting options
type Option func(*EventHandler) error

// WithValidation enabled validation on the handler
// requires the exact url use for the webhook entry
// and the key
func WithValidation(url, key string) Option {
	return func(e *EventHandler) error {
		e.url = url
		e.key = key
		e.validate = true
		return nil
	}
}

// WithLogger sets the logger to use
func WithLogger(l *log.Logger) Option {
	return func(e *EventHandler) error {
		e.logger = l
		return nil
	}
}

// WithDebug enables debugging
func WithDebug() Option {
	return func(e *EventHandler) error {
		e.debug = true
		return nil
	}
}

// EventHandler is a http handler for handling Mandrill webhooks
type EventHandler struct {
	url                string
	key                string
	syncHandlers       []SyncEventFunc
	messageHandlers    []MessageEventFunc
	unknownHandlers    []UnknownEventFunc
	sendHandlers       []SendEventFunc
	deferralHandlers   []DeferralEventFunc
	hardBounceHandlers []HardBounceEventFunc
	softBounceHandlers []SoftBounceEventFunc
	openHandlers       []OpenEventFunc
	clickHandlers      []ClickEventFunc
	spamHandlers       []SpamEventFunc
	unsubHandlers      []UnsubEventFunc
	rejectHandlers     []RejectEventFunc
	inboundHandlers    []InboundEventFunc
	blacklistHandlers  []BlacklistEventFunc
	whitelistHandlers  []WhitelistEventFunc
	validate           bool
	lock               sync.RWMutex
	logger             *log.Logger
	debug              bool
}

// NewEventHandler returns a new event handler
func NewEventHandler(opts ...Option) (*EventHandler, error) {
	e := &EventHandler{}
	for _, opt := range opts {
		if err := opt(e); err != nil {
			return nil, err
		}
	}
	if len(e.key) == 0 || len(e.url) == 0 {
		e.validate = false
	}
	if e.logger == nil {
		e.logger = log.New(os.Stdout, "mandrill_events", log.LstdFlags)
	}
	e.messageHandlers = []MessageEventFunc{}
	e.unknownHandlers = []UnknownEventFunc{}
	e.syncHandlers = []SyncEventFunc{}
	e.sendHandlers = []SendEventFunc{}
	e.deferralHandlers = []DeferralEventFunc{}
	e.hardBounceHandlers = []HardBounceEventFunc{}
	e.softBounceHandlers = []SoftBounceEventFunc{}
	e.openHandlers = []OpenEventFunc{}
	e.clickHandlers = []ClickEventFunc{}
	e.spamHandlers = []SpamEventFunc{}
	e.unsubHandlers = []UnsubEventFunc{}
	e.rejectHandlers = []RejectEventFunc{}
	e.inboundHandlers = []InboundEventFunc{}
	e.blacklistHandlers = []BlacklistEventFunc{}
	e.whitelistHandlers = []WhitelistEventFunc{}
	return e, nil
}

// Handle handles http requests from mandrill webhooks
func (e *EventHandler) Handle(w http.ResponseWriter, r *http.Request) {
	evts, err := events.ParseRequest(r)
	if err != nil {
		e.logger.Printf("error parsing request: %s", err.Error())
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	if e.validate {
		if err := events.ValidateWebhookSignature(e.url, e.key, r); err != nil {
			e.logger.Printf("error validating signature: %s", err.Error())
		}
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	sendParses := []events.SendEvent{}
	deferralParses := []events.DeferralEvent{}
	hardBounceParses := []events.HardBounceEvent{}
	softBounceParses := []events.SoftBounceEvent{}
	openParses := []events.OpenEvent{}
	clickParses := []events.ClickEvent{}
	spamParses := []events.SpamEvent{}
	unsubParses := []events.UnsubEvent{}
	rejectParses := []events.RejectEvent{}
	inboundParses := []events.InboundMessageEvent{}
	blacklistParses := []events.BlacklistEvent{}
	whitelistParses := []events.WhitelistEvent{}

	failedParses := []events.WebhookEvent{}
	messageParses := []events.MessageEvent{}
	syncParses := []events.SyncEvent{}
	unknownParses := []events.WebhookEvent{}

	// Parse all events into buckets
	if e.debug {
		e.logger.Printf("processing %d events", len(evts))
	}
	for _, evt := range evts {
		if e.debug {
			e.logger.Printf("parsing event")
		}
		switch evt.Type {
		case events.MessageEventType:
			me, err := events.ParseInnerEvent(evt)
			if err != nil {
				if e.debug {
					e.logger.Printf("unable to parse MessageEvent %v", evt)
				}
				failedParses = append(failedParses, evt)
			} else {
				i := me.(events.MessageEvent)
				switch fe := i.Data.(type) {
				case events.SendEvent:
					sendParses = append(sendParses, fe)
				case events.DeferralEvent:
					deferralParses = append(deferralParses, fe)
				case events.HardBounceEvent:
					hardBounceParses = append(hardBounceParses, fe)
				case events.SoftBounceEvent:
					softBounceParses = append(softBounceParses, fe)
				case events.OpenEvent:
					openParses = append(openParses, fe)
				case events.ClickEvent:
					clickParses = append(clickParses, fe)
				case events.SpamEvent:
					spamParses = append(spamParses, fe)
				case events.UnsubEvent:
					unsubParses = append(unsubParses, fe)
				case events.RejectEvent:
					rejectParses = append(rejectParses, fe)
				case events.InboundMessageEvent:
					inboundParses = append(inboundParses, fe)
				default:
					messageParses = append(messageParses, i)
				}
			}
		case events.SyncEventType:
			se, err := events.ParseInnerEvent(evt)
			if err != nil {
				if e.debug {
					e.logger.Printf("unable to parse SyncEvent %v", evt)
				}
				failedParses = append(failedParses, evt)
			} else {
				i := se.(events.SyncEvent)
				switch fe := i.Data.(type) {
				case events.BlacklistEvent:
					blacklistParses = append(blacklistParses, fe)
				case events.WhitelistEvent:
					whitelistParses = append(whitelistParses, fe)
				default:
					syncParses = append(syncParses, i)
				}
			}
		default:
			unknownParses = append(unknownParses, evt)
		}
	}

	for _, p := range failedParses {
		for _, f := range e.unknownHandlers {
			e.lock.RLock()
			defer e.lock.RUnlock()
			go f(r.Context(), p)
		}
	}
	for _, p := range messageParses {
		for _, f := range e.messageHandlers {
			e.lock.RLock()
			defer e.lock.RUnlock()
			go f(r.Context(), p)
		}
	}
	for _, p := range syncParses {
		for _, f := range e.syncHandlers {
			e.lock.RLock()
			defer e.lock.RUnlock()
			go f(r.Context(), p)

		}
	}
	w.WriteHeader(http.StatusOK)
}

// AddSendHandler adds a handler to process events.SendEvent
func (e *EventHandler) AddSendHandler(f ...SendEventFunc) {
	if e.debug {
		e.logger.Printf("adding %d SendHandler(s)", len(f))
	}
	e.lock.Lock()
	e.sendHandlers = append(e.sendHandlers, f...)
	e.lock.Unlock()
}

// AddDeferralHandler adds a handler to process events.DeferralEvent
func (e *EventHandler) AddDeferralHandler(f ...DeferralEventFunc) {
	if e.debug {
		e.logger.Printf("adding %d DeferralHandler(s)", len(f))
	}
	e.lock.Lock()
	e.deferralHandlers = append(e.deferralHandlers, f...)
	e.lock.Unlock()
}

// AddHardBounceHandler adds a handler to process events.HardBounceEvent
func (e *EventHandler) AddHardBounceHandler(f ...HardBounceEventFunc) {
	if e.debug {
		e.logger.Printf("adding %d HardBoundHandler(s)", len(f))
	}
	e.lock.Lock()
	e.hardBounceHandlers = append(e.hardBounceHandlers, f...)
	e.lock.Unlock()
}

// AddSoftBounceHandler adds a handler to process events.SoftBounceEvent
func (e *EventHandler) AddSoftBounceHandler(f ...SoftBounceEventFunc) {
	if e.debug {
		e.logger.Printf("adding %d SoftBounceHandler(s)", len(f))
	}
	e.lock.Lock()
	e.softBounceHandlers = append(e.softBounceHandlers, f...)
	e.lock.Unlock()
}

// AddOpenHandler adds a handler to process events.OpenEvent
func (e *EventHandler) AddOpenHandler(f ...OpenEventFunc) {
	if e.debug {
		e.logger.Printf("adding %d OpenHandler(s)", len(f))
	}
	e.lock.Lock()
	e.openHandlers = append(e.openHandlers, f...)
	e.lock.Unlock()
}

// AddClickHandler adds a handler to process events.ClickEvent
func (e *EventHandler) AddClickHandler(f ...ClickEventFunc) {
	if e.debug {
		e.logger.Printf("adding %d ClickHandler(s)", len(f))
	}
	e.lock.Lock()
	e.clickHandlers = append(e.clickHandlers, f...)
	e.lock.Unlock()
}

// AddSpamHandler adds a handler to process events.SpamEvent
func (e *EventHandler) AddSpamHandler(f ...SpamEventFunc) {
	if e.debug {
		e.logger.Printf("adding %d SpamHandler(s)", len(f))
	}
	e.lock.Lock()
	e.spamHandlers = append(e.spamHandlers, f...)
	e.lock.Unlock()
}

// AddUnsubHandler adds a handler to process events.UnsubEvent
func (e *EventHandler) AddUnsubHandler(f ...UnsubEventFunc) {
	if e.debug {
		e.logger.Printf("adding %d UnsubHandler(s)", len(f))
	}
	e.lock.Lock()
	e.unsubHandlers = append(e.unsubHandlers, f...)
	e.lock.Unlock()
}

// AddRejectHandler adds a handler to process events.RejectEvent
func (e *EventHandler) AddRejectHandler(f ...RejectEventFunc) {
	if e.debug {
		e.logger.Printf("adding %d RejectHandler(s)", len(f))
	}
	e.lock.Lock()
	e.rejectHandlers = append(e.rejectHandlers, f...)
	e.lock.Unlock()
}

// AddInboundHandler adds a handler to process events.InboundMessageEvent
func (e *EventHandler) AddInboundHandler(f ...InboundEventFunc) {
	if e.debug {
		e.logger.Printf("adding %d InboundHandler(s)", len(f))
	}
	e.lock.Lock()
	e.inboundHandlers = append(e.inboundHandlers, f...)
	e.lock.Unlock()
}

// AddBlacklistHandler adds a handler to process events.BlacklistEvent
func (e *EventHandler) AddBlacklistHandler(f ...BlacklistEventFunc) {
	if e.debug {
		e.logger.Printf("adding %d BlacklistHandler(s)", len(f))
	}
	e.lock.Lock()
	e.blacklistHandlers = append(e.blacklistHandlers, f...)
	e.lock.Unlock()
}

// AddWhitelistHandler adds a handler to process events.WhitelistEvent
func (e *EventHandler) AddWhitelistHandler(f ...WhitelistEventFunc) {
	if e.debug {
		e.logger.Printf("adding %d WhitelistHandler(s)", len(f))
	}
	e.lock.Lock()
	e.whitelistHandlers = append(e.whitelistHandlers, f...)
	e.lock.Unlock()
}

// AddMessageHandler adds a function to process events.MessageEvent
func (e *EventHandler) AddMessageHandler(f ...MessageEventFunc) {
	e.lock.Lock()
	e.messageHandlers = append(e.messageHandlers, f...)
	e.lock.Unlock()
}

// AddSyncHandler adds a function to process events.SyncEvent
func (e *EventHandler) AddSyncHandler(f ...SyncEventFunc) {
	e.lock.Lock()
	e.syncHandlers = append(e.syncHandlers, f...)
	e.lock.Unlock()
}

// AddUnknownHandler adds a function to process unknown events
func (e *EventHandler) AddUnknownHandler(f ...UnknownEventFunc) {
	e.lock.Lock()
	e.unknownHandlers = append(e.unknownHandlers, f...)
	e.lock.Unlock()
}
