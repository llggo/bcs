package business

import (
	"fmt"
)

type PriorityConfig struct {
	PriorityStep             int `json:"priority_step"`
	InternalVipCard          int `json:"internal_vip_card"`
	CustomerVipCard          int `json:"customer_vip_card"`
	PrivilegedCustomer       int `json:"privileged_customer"`
	MovedTicket              int `json:"moved_ticket"`
	BookedTicket             int `json:"booked_ticket"`
	RestoreTicket            int `json:"restore_ticket"`
	MinPriorityRestricted    int `json:"min_priority_restricted"`
	MinPriorityUnorderedCall int `json:"min_priority_unordered_call"`
}

func (c PriorityConfig) String() string {
	return fmt.Sprintf(
		"priority:vi=%d;vc=%d;pc=%d;mo=%d;bo=%d;re=%d;mps=%d;mpuc=%d",
		c.InternalVipCard, c.CustomerVipCard, c.PrivilegedCustomer,
		c.MovedTicket, c.BookedTicket, c.RestoreTicket, c.MinPriorityRestricted, c.MinPriorityUnorderedCall,
	)
}

func (c *PriorityConfig) Check() {
	if c.PriorityStep < 0 {
		c.PriorityStep = 0
	}

	if c.MinPriorityRestricted < 0 {
		// no restriction
		c.MinPriorityRestricted = 1 << 16
	}
	if c.MinPriorityUnorderedCall < 0 {
		// no unordered call
		c.MinPriorityUnorderedCall = 1 << 16
	}
	if c.InternalVipCard < 0 {
		c.InternalVipCard = 0
	}
	if c.CustomerVipCard < 0 {
		c.CustomerVipCard = 0
	}
	if c.MovedTicket < 0 {
		c.MovedTicket = 0
	}
	if c.BookedTicket < 0 {
		c.BookedTicket = 0
	}
	if c.RestoreTicket < 0 {
		c.RestoreTicket = c.MovedTicket
	}
	if c.PrivilegedCustomer < 0 {
		c.PrivilegedCustomer = 0
	}
}
