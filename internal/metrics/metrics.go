package metrics

import (
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/armon/go-metrics"
	"github.com/armon/go-metrics/prometheus"
	"github.com/gregorychen3/espresso-controller/internal/log"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"go.uber.org/zap"
)

var (
	metricKeyMemUtilizationRatio = []string{"raspi_mem_utilization_ratio"}
	metricKeyCPUUtilizationRatio = []string{"raspi_cpu_utilization_ratio"}
	metricKeyRaspiGPUTemperature = []string{"raspi_gpu_temperature"}
	metricKeyRaspiCPUTemperature = []string{"raspi_cpu_temperature"}
)

// InitMetrics creates a prometheus sink
func InitMetrics() error {
	log.Info("Configurting prometheus metrics sink")
	promSink, err := prometheus.NewPrometheusSinkFrom(prometheus.PrometheusOpts{Expiration: time.Hour * 24})
	if err != nil {
		return err
	}

	m, err := metrics.NewGlobal(metrics.DefaultConfig("espresso"), promSink)
	if err != nil {
		return err
	}
	m.EnableHostname = false

	return nil
}

// CollectSystemMetrics for collecting system-level metrics on-demand (e.g., at prometheus scrape-time)
func CollectSystemMetrics() {
	memStats, err := mem.VirtualMemory()
	if err != nil {
		log.Error("Failed to collect system memory utilization metrics", zap.Error(err))
	} else {
		memUtilizationRatio := float32(memStats.UsedPercent / 100)
		log.Info("Sampled system memory utilization", zap.Float32("memUtilizationRatio", memUtilizationRatio))
		metrics.SetGauge(metricKeyMemUtilizationRatio, memUtilizationRatio)
	}

	cpuStats, err := cpu.Percent(0, false)
	if err != nil {
		log.Error("Failed to collect system cpu utilization metrics", zap.Error(err))
	} else {
		cpuUtilizationRatio := float32(cpuStats[0] / 100)
		log.Info("Sampled system cpu utilization", zap.Float32("cpuUtilizationRatio", cpuUtilizationRatio))
		metrics.SetGauge(metricKeyCPUUtilizationRatio, cpuUtilizationRatio)
	}

	gpuTemperature, err := sampleRaspiGPUTemperature()
	if err != nil {
		log.Error("Failed to sample gpu temperature", zap.Error(err))
	} else {
		log.Info("Sampled gpu temperature", zap.Float32("gpuTemperature", gpuTemperature))
		metrics.SetGauge(metricKeyRaspiGPUTemperature, gpuTemperature)
	}

	cpuTemperature, err := sampleRaspiCPUTemperature()
	if err != nil {
		log.Error("Failed to sample cpu temperature", zap.Error(err))
	} else {
		log.Info("Sampled cpu temperature", zap.Float32("cpuTemperature", cpuTemperature))
		metrics.SetGauge(metricKeyRaspiCPUTemperature, cpuTemperature)
	}
}

func sampleRaspiGPUTemperature() (float32, error) {
	outBytes, err := exec.Command("/usr/bin/vcgencmd measure_temp").Output()
	if err != nil {
		return 0, err
	}

	temperature64, err := strconv.ParseFloat(string(outBytes)[5:9], 32)
	if err != nil {
		return 0, err
	}

	return float32(temperature64), nil
}

func sampleRaspiCPUTemperature() (float32, error) {
	fileBytes, err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		return 0, err
	}

	temperaturef64, err := strconv.ParseFloat(strings.TrimSpace(string(fileBytes)), 64)
	if err != nil {
		return 0, err
	}

	return float32(temperaturef64 / 1000), nil
}
