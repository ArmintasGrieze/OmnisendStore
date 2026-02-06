package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port            string
	DBPath          string
	OmnisendAPIKey  string
	OmnisendSnippet string
	AdminUser       string
	AdminPass       string
}

func Load() (Config, error) {
	c := Config{
		Port:            os.Getenv("PORT"),
		DBPath:          os.Getenv("DB_PATH"),
		OmnisendAPIKey:  "test", //os.Getenv("OMNISEND_API_KEY"),
		OmnisendSnippet: os.Getenv("OMNISEND_SNIPPET"),
		AdminUser:       "test", //os.Getenv("ADMIN_USER"),
		AdminPass:       "test", //os.Getenv("ADMIN_PASS"),
	}

	if c.Port == "" {
		c.Port = "8080"
	}
	if c.DBPath == "" {
		c.DBPath = "./data/store.db"
	}

	// Omnisend key/snippet can be empty during early local dev, but for your MVP you want them set.
	if c.OmnisendAPIKey == "" {
		return Config{}, fmt.Errorf("OMNISEND_API_KEY is required")
	}
	if c.AdminUser == "" || c.AdminPass == "" {
		// We'll use this for admin panel later; set it now so it's not forgotten.
		return Config{}, fmt.Errorf("ADMIN_USER and ADMIN_PASS are required")
	}

	return c, nil
}
