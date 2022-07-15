package load

import "time"

type Config struct {
    Player   string `json:"player"`
    Enemy    string `json:"enemy"`
    Wall     string `json:"wall"`
    Dot      string `json:"dot"`
    Pill     string `json:"pill"`
    Death    string `json:"death"`
    Space    string `json:"space"`
    UseEmoji bool   `json:"use_emoji"`
	EnemyBlue string `json:"enemy_blue"`
	PillDuration time.Duration `json:"pill_duration_secs"`
}

var cfg Config

