package convert

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/petrolax/yaml-converter/currency"
	"github.com/prometheus/client_golang/prometheus"
	yaml "gopkg.in/yaml.v2"
)

func YamlToOpenMetrics(path string) []prometheus.Gauge {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	fileb, err := ioutil.ReadAll(file) // just pass the file name
	if err != nil {
		log.Fatalln(err)
	}

	var crs currency.CurrencyRate
	err = yaml.Unmarshal(fileb, &crs)
	if err != nil {
		log.Fatalln(err)
	}

	res := make([]prometheus.Gauge, len(crs.Currencies))
	for i, curr := range crs.Currencies {
		res[i] = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "currency",
				Name:      curr.Name,
			})
		res[i].Set(curr.Value)
	}
	return res
}
