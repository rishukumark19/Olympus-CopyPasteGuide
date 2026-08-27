package main
import (
	"encoding/json"
	"fmt"
	"strings"
)

type Config struct {
	TierPolicy    string   `json:"tier_policy,omitempty"`
	ConsumerTiers []string `json:"consumer_tiers,omitempty"`
}

type Req struct {
	Config Config `json:"config"`
}

func main() {
	j := `{"config":{"tier_policy":"exclusive_lease","consumer_tiers":["A"]}}`
	var r Req
	dec := json.NewDecoder(strings.NewReader(j))
	dec.DisallowUnknownFields()
	err := dec.Decode(&r)
	fmt.Println("Error:", err)
}
