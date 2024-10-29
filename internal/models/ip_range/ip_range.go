package ip_range

import (
	"os"
	"sort"
	"sync"

	"github.com/hightemp/ip_info_service/internal/logger"
	"github.com/hightemp/ip_info_service/internal/utils"
	"gopkg.in/yaml.v3"
)

type IpRange struct {
	Name    string
	StartIP uint32
	EndIp   uint32
}

var (
	contriesRanges []IpRange
	orgRanges      []IpRange
	Loaded         = false
)

func AddCountry(name string, start uint32, end uint32) {
	contriesRanges = append(contriesRanges, IpRange{Name: name, StartIP: start, EndIp: end})
	SortCountriesRanges()
}

func AddOrganization(name string, start uint32, end uint32) {
	orgRanges = append(orgRanges, IpRange{Name: name, StartIP: start, EndIp: end})
	SortOrgRanges()
}

func AddCountriesRanges(ranges []IpRange) {
	contriesRanges = append(contriesRanges, ranges...)
	SortCountriesRanges()
}

func AddOrganizationsRanges(ranges []IpRange) {
	orgRanges = append(orgRanges, ranges...)
	SortOrgRanges()
}

func SortCountriesRanges() {
	sort.Slice(contriesRanges, func(i, j int) bool {
		return contriesRanges[i].StartIP < contriesRanges[j].StartIP
	})
}

func SortOrgRanges() {
	sort.Slice(orgRanges, func(i, j int) bool {
		return orgRanges[i].StartIP < orgRanges[j].StartIP
	})
}

func SortRanges() {
	SortCountriesRanges()
	SortOrgRanges()
}

func binarySearch(arr []IpRange, ipInt uint32) string {
	var result = "Unknown"
	left, right := uint32(0), uint32(len(arr)-1)

	if len(arr) == 0 {
		return result
	}

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid].StartIP <= ipInt && arr[mid].EndIp >= ipInt {
			result = arr[mid].Name
			break
		}
		if arr[mid].StartIP > ipInt {
			if mid == 0 {
				break
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}

func SearchIpInfo(ip string) (string, string) {
	var country string
	var organization string

	ipInt := utils.IpStringToInt(ip)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		country = binarySearch(contriesRanges, ipInt)
		wg.Done()
	}()

	go func() {
		organization = binarySearch(orgRanges, ipInt)
		wg.Done()
	}()

	wg.Wait()

	return country, organization
}

const (
	COUNRIES_DATA_YAML_FILE      = "./data/countries.yaml"
	ORGANIZATIONS_DATA_YAML_FILE = "./data/organizations.yaml"
)

func Save() error {
	data, err := yaml.Marshal(contriesRanges)
	if err != nil {
		return err
	}
	if err = os.WriteFile(COUNRIES_DATA_YAML_FILE, data, 0644); err != nil {
		return err
	}

	data, err = yaml.Marshal(orgRanges)
	if err != nil {
		return err
	}
	if err = os.WriteFile(ORGANIZATIONS_DATA_YAML_FILE, data, 0644); err != nil {
		return err
	}

	return nil
}

func Load() error {
	if Loaded {
		logger.LogInfo("Data was loaded from json. There is no need load from yaml")
		return nil
	}

	if _, err := os.Stat(COUNRIES_DATA_YAML_FILE); err == nil {
		data, err := os.ReadFile(COUNRIES_DATA_YAML_FILE)
		if err != nil {
			return err
		}
		var localContriesRanges []IpRange
		if err := yaml.Unmarshal(data, &localContriesRanges); err != nil {
			return err
		}
		AddCountriesRanges(localContriesRanges)
		logger.LogInfo("Loaded %d countries from %s", len(localContriesRanges), COUNRIES_DATA_YAML_FILE)
	}

	if _, err := os.Stat(ORGANIZATIONS_DATA_YAML_FILE); err == nil {
		data, err := os.ReadFile(ORGANIZATIONS_DATA_YAML_FILE)
		if err != nil {
			return err
		}
		var localOrgRanges []IpRange
		if err := yaml.Unmarshal(data, &localOrgRanges); err != nil {
			return err
		}
		AddOrganizationsRanges(localOrgRanges)
		logger.LogInfo("Loaded %d organizations from %s", len(localOrgRanges), ORGANIZATIONS_DATA_YAML_FILE)
	}

	return nil
}
