package openalgo

// Ping checks API connectivity
func (c *Client) Ping() (map[string]interface{}, error) {
	payload := map[string]interface{}{
		"apikey": c.apiKey,
	}
	return c.makeRequest("POST", "ping", payload)
}

// Holidays gets trading holidays for a year
func (c *Client) Holidays(year int) (map[string]interface{}, error) {
	payload := map[string]interface{}{
		"apikey": c.apiKey,
		"year":   year,
	}
	return c.makeRequest("POST", "market/holidays", payload)
}

// Timings gets exchange trading timings for a specific date
func (c *Client) Timings(date string) (map[string]interface{}, error) {
	payload := map[string]interface{}{
		"apikey": c.apiKey,
		"date":   date,
	}
	return c.makeRequest("POST", "market/timings", payload)
}

// Telegram sends a Telegram notification with default priority (5)
func (c *Client) Telegram(username, message string) (map[string]interface{}, error) {
	return c.TelegramWithPriority(username, message, 5)
}

// TelegramWithPriority sends a Telegram notification with custom priority
func (c *Client) TelegramWithPriority(username, message string, priority int) (map[string]interface{}, error) {
	payload := map[string]interface{}{
		"apikey":   c.apiKey,
		"username": username,
		"message":  message,
		"priority": priority,
	}
	return c.makeRequest("POST", "telegram/notify", payload)
}
