package global

import "os"

var AMQP_URL string

func LoadSetting() {
	// Get the connection string from the environment variable
	AMQP_URL = os.Getenv("AMQP_URL")

	//If it doesn't exist, use the default connection string.

	if AMQP_URL == "" {
		//Don't do this in production, this is for testing purposes only.
		AMQP_URL = Default_AMQP_URL
	}
}
