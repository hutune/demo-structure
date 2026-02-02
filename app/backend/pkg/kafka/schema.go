package kafka

// Topic naming convention: domain.event (e.g. campaign.created, device.heartbeat)
const (
	TopicCampaignCreated = "campaign.created"
	TopicCampaignStatus  = "campaign.status_changed"
	TopicDeviceHeartbeat = "device.heartbeat"
	TopicDevicePlayback  = "device.playback_completed"
	TopicBillingCharged  = "billing.charged"
	TopicBillingRevenue  = "billing.revenue_distributed"
)
