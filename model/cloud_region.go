package model

// CloudRegion : CloudRegion model
type CloudRegion string

// List of possible CloudRegion values
const (
	CloudRegion_AWS_US_EAST_1                   CloudRegion = "AWS_US_EAST_1"
	CloudRegion_AWS_US_EAST_2                   CloudRegion = "AWS_US_EAST_2"
	CloudRegion_AWS_US_WEST_1                   CloudRegion = "AWS_US_WEST_1"
	CloudRegion_AWS_US_WEST_2                   CloudRegion = "AWS_US_WEST_2"
	CloudRegion_AWS_EU_WEST_1                   CloudRegion = "AWS_EU_WEST_1"
	CloudRegion_AWS_EU_CENTRAL_1                CloudRegion = "AWS_EU_CENTRAL_1"
	CloudRegion_AWS_AP_SOUTHEAST_1              CloudRegion = "AWS_AP_SOUTHEAST_1"
	CloudRegion_AWS_AP_SOUTHEAST_2              CloudRegion = "AWS_AP_SOUTHEAST_2"
	CloudRegion_AWS_AP_NORTHEAST_1              CloudRegion = "AWS_AP_NORTHEAST_1"
	CloudRegion_AWS_AP_NORTHEAST_2              CloudRegion = "AWS_AP_NORTHEAST_2"
	CloudRegion_AWS_AP_SOUTH_1                  CloudRegion = "AWS_AP_SOUTH_1"
	CloudRegion_AWS_SA_EAST_1                   CloudRegion = "AWS_SA_EAST_1"
	CloudRegion_AWS_EU_WEST_2                   CloudRegion = "AWS_EU_WEST_2"
	CloudRegion_AWS_EU_WEST_3                   CloudRegion = "AWS_EU_WEST_3"
	CloudRegion_AWS_CA_CENTRAL_1                CloudRegion = "AWS_CA_CENTRAL_1"
	CloudRegion_AWS_EU_NORTH_1                  CloudRegion = "AWS_EU_NORTH_1"
	CloudRegion_GOOGLE_US_CENTRAL_1             CloudRegion = "GOOGLE_US_CENTRAL_1"
	CloudRegion_GOOGLE_US_EAST_1                CloudRegion = "GOOGLE_US_EAST_1"
	CloudRegion_GOOGLE_ASIA_EAST_1              CloudRegion = "GOOGLE_ASIA_EAST_1"
	CloudRegion_GOOGLE_EUROPE_WEST_1            CloudRegion = "GOOGLE_EUROPE_WEST_1"
	CloudRegion_GOOGLE_US_WEST_1                CloudRegion = "GOOGLE_US_WEST_1"
	CloudRegion_GOOGLE_ASIA_EAST_2              CloudRegion = "GOOGLE_ASIA_EAST_2"
	CloudRegion_GOOGLE_ASIA_NORTHEAST_1         CloudRegion = "GOOGLE_ASIA_NORTHEAST_1"
	CloudRegion_GOOGLE_ASIA_SOUTH_1             CloudRegion = "GOOGLE_ASIA_SOUTH_1"
	CloudRegion_GOOGLE_ASIA_SOUTHEAST_1         CloudRegion = "GOOGLE_ASIA_SOUTHEAST_1"
	CloudRegion_GOOGLE_AUSTRALIA_SOUTHEAST_1    CloudRegion = "GOOGLE_AUSTRALIA_SOUTHEAST_1"
	CloudRegion_GOOGLE_EUROPE_NORTH_1           CloudRegion = "GOOGLE_EUROPE_NORTH_1"
	CloudRegion_GOOGLE_EUROPE_WEST_2            CloudRegion = "GOOGLE_EUROPE_WEST_2"
	CloudRegion_GOOGLE_EUROPE_WEST_3            CloudRegion = "GOOGLE_EUROPE_WEST_3"
	CloudRegion_GOOGLE_EUROPE_WEST_4            CloudRegion = "GOOGLE_EUROPE_WEST_4"
	CloudRegion_GOOGLE_NORTHAMERICA_NORTHEAST_1 CloudRegion = "GOOGLE_NORTHAMERICA_NORTHEAST_1"
	CloudRegion_GOOGLE_SOUTHAMERICA_EAST_1      CloudRegion = "GOOGLE_SOUTHAMERICA_EAST_1"
	CloudRegion_GOOGLE_US_EAST_4                CloudRegion = "GOOGLE_US_EAST_4"
	CloudRegion_GOOGLE_US_WEST_2                CloudRegion = "GOOGLE_US_WEST_2"
	CloudRegion_AZURE_ASIA_EAST                 CloudRegion = "AZURE_ASIA_EAST"
	CloudRegion_AZURE_ASIA_SOUTHEAST            CloudRegion = "AZURE_ASIA_SOUTHEAST"
	CloudRegion_AZURE_AUSTRALIA_EAST            CloudRegion = "AZURE_AUSTRALIA_EAST"
	CloudRegion_AZURE_AUSTRALIA_SOUTHEAST       CloudRegion = "AZURE_AUSTRALIA_SOUTHEAST"
	CloudRegion_AZURE_BRAZIL_SOUTH              CloudRegion = "AZURE_BRAZIL_SOUTH"
	CloudRegion_AZURE_CANADA_CENTRAL            CloudRegion = "AZURE_CANADA_CENTRAL"
	CloudRegion_AZURE_EUROPE_NORTH              CloudRegion = "AZURE_EUROPE_NORTH"
	CloudRegion_AZURE_EUROPE_WEST               CloudRegion = "AZURE_EUROPE_WEST"
	CloudRegion_AZURE_FRANCE_CENTRAL            CloudRegion = "AZURE_FRANCE_CENTRAL"
	CloudRegion_AZURE_GERMANY_WESTCENTRAL       CloudRegion = "AZURE_GERMANY_WESTCENTRAL"
	CloudRegion_AZURE_INDIA_CENTRAL             CloudRegion = "AZURE_INDIA_CENTRAL"
	CloudRegion_AZURE_INDIA_SOUTH               CloudRegion = "AZURE_INDIA_SOUTH"
	CloudRegion_AZURE_JAPAN_EAST                CloudRegion = "AZURE_JAPAN_EAST"
	CloudRegion_AZURE_JAPAN_WEST                CloudRegion = "AZURE_JAPAN_WEST"
	CloudRegion_AZURE_KOREA_CENTRAL             CloudRegion = "AZURE_KOREA_CENTRAL"
	CloudRegion_AZURE_UAE_NORTH                 CloudRegion = "AZURE_UAE_NORTH"
	CloudRegion_AZURE_US_CENTRAL                CloudRegion = "AZURE_US_CENTRAL"
	CloudRegion_AZURE_US_EAST                   CloudRegion = "AZURE_US_EAST"
	CloudRegion_AZURE_US_EAST2                  CloudRegion = "AZURE_US_EAST2"
	CloudRegion_AZURE_US_WEST                   CloudRegion = "AZURE_US_WEST"
	CloudRegion_AZURE_US_WEST2                  CloudRegion = "AZURE_US_WEST2"
	CloudRegion_AZURE_US_SOUTH_CENTRAL          CloudRegion = "AZURE_US_SOUTH_CENTRAL"
	CloudRegion_AZURE_US_NORTH_CENTRAL          CloudRegion = "AZURE_US_NORTH_CENTRAL"
	CloudRegion_AZURE_UK_SOUTH                  CloudRegion = "AZURE_UK_SOUTH"
	CloudRegion_AKAMAI_BR_GRU                   CloudRegion = "AKAMAI_BR_GRU"
	CloudRegion_AKAMAI_FR_PAR                   CloudRegion = "AKAMAI_FR_PAR"
	CloudRegion_AKAMAI_JP_OSA                   CloudRegion = "AKAMAI_JP_OSA"
	CloudRegion_AKAMAI_US_SEA                   CloudRegion = "AKAMAI_US_SEA"
	CloudRegion_NORTH_AMERICA                   CloudRegion = "NORTH_AMERICA"
	CloudRegion_SOUTH_AMERICA                   CloudRegion = "SOUTH_AMERICA"
	CloudRegion_EUROPE                          CloudRegion = "EUROPE"
	CloudRegion_AFRICA                          CloudRegion = "AFRICA"
	CloudRegion_ASIA                            CloudRegion = "ASIA"
	CloudRegion_AUSTRALIA                       CloudRegion = "AUSTRALIA"
	CloudRegion_AWS                             CloudRegion = "AWS"
	CloudRegion_GOOGLE                          CloudRegion = "GOOGLE"
	CloudRegion_EXTERNAL                        CloudRegion = "EXTERNAL"
	CloudRegion_AUTO                            CloudRegion = "AUTO"
)
