package data_parser

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/hightemp/ip_info_service/internal/config"
	"github.com/hightemp/ip_info_service/internal/models/ip_range"
)

var (
	contriesData [][]string
	orgData      [][]string
)

func Load() error {
	cfg := config.Get()

	if cfg.ContriesDataFilePath != "" {
		bytes, err := os.ReadFile(cfg.ContriesDataFilePath)

		if err != nil {
			return fmt.Errorf("Can't read file with data: '%s' %v", cfg.ContriesDataFilePath, err)
		}

		err = json.Unmarshal(bytes, &contriesData)

		if err != nil {
			return fmt.Errorf("Can't parse data: '%s' %v", cfg.ContriesDataFilePath, err)
		}
	}

	if cfg.OrgDataFilePath != "" {
		bytes, err := os.ReadFile(cfg.OrgDataFilePath)

		if err != nil {
			return fmt.Errorf("Can't read file with data: '%s' %v", cfg.OrgDataFilePath, err)
		}

		err = json.Unmarshal(bytes, &contriesData)

		if err != nil {
			return fmt.Errorf("Can't parse data: '%s' %v", cfg.OrgDataFilePath, err)
		}
	}

	Parse()
	return nil
}

func Parse() {
	for _, i := range contriesData {
		s, _ := strconv.ParseUint(i[2], 10, 32)
		e, _ := strconv.ParseUint(i[3], 10, 32)
		ip_range.AddCountry(i[1], uint32(s), uint32(e))
	}

	for _, i := range orgData {
		s, _ := strconv.ParseUint(i[2], 10, 32)
		e, _ := strconv.ParseUint(i[3], 10, 32)
		ip_range.AddOrganization(i[1], uint32(s), uint32(e))
	}

	contriesData = nil
	orgData = nil

	runtime.GC()
}
