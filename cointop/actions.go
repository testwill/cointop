package cointop

// ActionsMap returns a map of all the available actions
func ActionsMap() map[string]bool {
	return map[string]bool{
		"first_page":                        true,
		"help":                              true,
		"toggle_show_help":                  true,
		"close_help":                        true,
		"last_page":                         true,
		"move_to_page_first_row":            true,
		"move_to_page_last_row":             true,
		"move_to_page_visible_first_row":    true,
		"move_to_page_visible_last_row":     true,
		"move_to_page_visible_middle_row":   true,
		"move_up":                           true,
		"move_down":                         true,
		"next_page":                         true,
		"open_link":                         true,
		"page_down":                         true,
		"page_up":                           true,
		"previous_page":                     true,
		"quit":                              true,
		"quit_view":                         true,
		"refresh":                           true,
		"sort_column_1h_change":             true,
		"sort_column_24h_change":            true,
		"sort_column_24h_volume":            true,
		"sort_column_7d_change":             true,
		"sort_column_asc":                   true,
		"sort_column_available_supply":      true,
		"sort_column_desc":                  true,
		"sort_column_last_updated":          true,
		"sort_column_market_cap":            true,
		"sort_column_name":                  true,
		"sort_column_price":                 true,
		"sort_column_rank":                  true,
		"sort_column_symbol":                true,
		"sort_column_total_supply":          true,
		"sort_left_column":                  true,
		"sort_right_column":                 true,
		"toggle_row_chart":                  true,
		"open_search":                       true,
		"toggle_favorite":                   true,
		"toggle_show_favorites":             true,
		"previous_chart_range":              true,
		"next_chart_range":                  true,
		"first_chart_range":                 true,
		"last_chart_range":                  true,
		"toggle_show_currency_convert_menu": true,
		"show_currency_convert_menu":        true,
		"hide_currency_convert_menu":        true,
		"toggle_portfolio":                  true,
		"toggle_show_portfolio":             true,
		"enlarge_chart":                     true,
		"shorten_chart":                     true,
	}
}

// ActionExists returns true if action exists
func (ct *Cointop) ActionExists(action string) bool {
	return ct.ActionsMap[action]
}
