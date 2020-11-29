package retrievor

type team struct {
	CompetitorID          int    `json:"competitorId"`
	CompetitorWinamaxName string `json:"competitorWinamaxName"`
	CompetitorFTEName     string `json:"competitorFTEName"`
}

var teamMap = map[string]team{
	"1641": {
		CompetitorID:          1641,
		CompetitorWinamaxName: "Marseille",
		CompetitorFTEName:     "Marseille",
	},
	"1642": {
		CompetitorID:          1642,
		CompetitorWinamaxName: "Montpellier",
		CompetitorFTEName:     "Montpellier",
	},
	"1643": {
		CompetitorID:          1643,
		CompetitorWinamaxName: "Lille",
		CompetitorFTEName:     "Lille",
	},
	"1644": {
		CompetitorID:          1644,
		CompetitorWinamaxName: "Paris Saint Germain",
		CompetitorFTEName:     "PSG",
	},
	"1645": {
		CompetitorID:          1645,
		CompetitorWinamaxName: "Bordeaux",
		CompetitorFTEName:     "Bordeaux",
	},
	"1678": {
		CompetitorID:          1678,
		CompetitorWinamaxName: "Saint-Etienne",
		CompetitorFTEName:     "St Étienne",
	},
	"1647": {
		CompetitorID:          1647,
		CompetitorWinamaxName: "Nantes",
		CompetitorFTEName:     "Nantes",
	},
	"1648": {
		CompetitorID:          1648,
		CompetitorWinamaxName: "Lens",
		CompetitorFTEName:     "Lens",
	},
	"1649": {
		CompetitorID:          1649,
		CompetitorWinamaxName: "Lyon",
		CompetitorFTEName:     "Lyon",
	},
	"1682": {
		CompetitorID:          1682,
		CompetitorWinamaxName: "Reims",
		CompetitorFTEName:     "Reims",
	},
	"1651": {
		CompetitorID:          1651,
		CompetitorWinamaxName: "Metz",
		CompetitorFTEName:     "Metz",
	},
	"1684": {
		CompetitorID:          1684,
		CompetitorWinamaxName: "Angers",
		CompetitorFTEName:     "Angers",
	},
	"1653": {
		CompetitorID:          1653,
		CompetitorWinamaxName: "Monaco",
		CompetitorFTEName:     "Monaco",
	},
	"1686": {
		CompetitorID:          1686,
		CompetitorWinamaxName: "Dijon",
		CompetitorFTEName:     "Dijon FCO",
	},
	"1656": {
		CompetitorID:          1656,
		CompetitorWinamaxName: "Lorient",
		CompetitorFTEName:     "Lorient",
	},
	"1715": {
		CompetitorID:          1715,
		CompetitorWinamaxName: "Brest",
		CompetitorFTEName:     "Brest",
	},
	"1658": {
		CompetitorID:          1658,
		CompetitorWinamaxName: "Rennes",
		CompetitorFTEName:     "Rennes",
	},
	"1659": {
		CompetitorID:          1659,
		CompetitorWinamaxName: "Strasbourg",
		CompetitorFTEName:     "Strasbourg",
	},
	"1661": {
		CompetitorID:          1661,
		CompetitorWinamaxName: "Nice",
		CompetitorFTEName:     "Nice",
	},
	"1663": {
		CompetitorID:          1663,
		CompetitorWinamaxName: "N\u00eemes",
		CompetitorFTEName:     "N\u00eemes",
	},
}
