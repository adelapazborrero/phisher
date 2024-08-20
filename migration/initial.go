package migration

const CreateCampaignTable = `
	CREATE TABLE IF NOT EXISTS campaigns(id TEXT PRIMARY KEY, title TEXT);
`

const DropCampaignTable = `
	DROP TABLE IF EXISTS campaigns;
`
