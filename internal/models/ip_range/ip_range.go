package ip_range

import (
	"net"
	"os"
	"sort"
	"sync"

	"github.com/emirpasic/gods/utils"
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
)

func AddCountry(name string, start uint32, end uint32) {
	contriesRanges = append(contriesRanges, IpRange{Name: name, StartIP: start, EndIp: end})
}

func AddOrganization(name string, start uint32, end uint32) {
	orgRanges = append(orgRanges, IpRange{Name: name, StartIP: start, EndIp: end})
}

func AddCountriesRanges(ranges []IpRange) {
	for _, r := range ranges {
		contriesRanges = append(contriesRanges, r)
	}
}

func AddOrganizationsRanges(ranges []IpRange) {
	for _, r := range ranges {
		orgRanges = append(orgRanges, r)
	}
}

func SortRanges() {
	sort.Slice(contriesRanges, func(i, j int) bool {
		return contriesRanges[i].StartIP < contriesRanges[j].StartIP
	})
	sort.Slice(orgRanges, func(i, j int) bool {
		return orgRanges[i].StartIP < orgRanges[j].StartIP
	})
}

func SearchIpInfo(ipv4 string) (string, string) {
	var country string
	var organization string

	ip := net.ParseIP(ipv4)
	ipInt := utils.Ip2int(ip)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		left, right := uint32(0), uint32(len(contriesRanges)-1)

		for right > left {
			mid := (left + right) / 2
			if contriesRanges[mid].StartIP <= ipInt && contriesRanges[mid].EndIp >= ipInt {
				country = contriesRanges[mid].Name
				break
			}
			if contriesRanges[mid].StartIP > ipInt {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		wg.Done()
	}()

	go func() {
		left, right := uint32(0), uint32(len(contriesRanges)-1)

		for right > left {
			mid := (left + right) / 2
			if contriesRanges[mid].StartIP <= ipInt && contriesRanges[mid].EndIp >= ipInt {
				country = contriesRanges[mid].Name
				break
			}
			if contriesRanges[mid].StartIP > ipInt {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
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
	}

	SortRanges()

	return nil
}
