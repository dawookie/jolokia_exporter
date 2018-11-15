package jolokia

import "testing"

var expectedConfig = &Config{
	Metrics: []MetricMapping{
		{
			Source: MetricSource{
				Mbean:     "java.lang:type=Memory",
				Attribute: "HeapMemoryUsage",
				Path:      "used",
			},
			Target: "java_memory_heap_memory_usage_used",
		},
		{
			Source: MetricSource{
				Mbean:     "java.lang:type=Memory",
				Attribute: "HeapMemoryUsage",
				Path:      "max",
			},
			Target: "java_memory_max",
		},
		{
			Source: MetricSource{
				Mbean:     "java.lang:type=Threading",
				Attribute: "ThreadCount",
			},
			Target: "java_threading_thread_count",
		},
		{
			Source: MetricSource{
				Mbean: "java.lang:type=OperatingSystem",
			},
			Target: "java_os",
		},
		{
			Source: MetricSource{
				Mbean:     "java.lang:name=ParNew,type=GarbageCollector",
				Attribute: "CollectionCount",
				MetricTarget: &MetricTarget{
					Url: "service:jmx:rmi:///jndi/rmi://127.0.0.1:9119/jmxrmi",
				},
			},
			Target: "java_gc_parnew_collectioncount",
		},
		{
			Source: MetricSource{
				Mbean:     "java.lang:name=ConcurrentMarkSweep,type=GarbageCollector",
				Attribute: "CollectionCount",
				MetricTarget: &MetricTarget{
					Url: "service:jmx:rmi:///jndi/rmi://127.0.0.1:9119/jmxrmi",
				},
			},
			Target: "java_gc_cms_collectioncount",
		},
	},
}

func TestLoadConfigYAML(t *testing.T) {
	file := "./fixtures/config.yaml"

	config, err := LoadConfig(file)
	if err != nil {
		t.Fatal("Error loading config file:", err)
	}

	checkConfig(t, config)
}

func checkConfig(t *testing.T, config *Config) {
	if config == nil {
		t.Fatal("Expected config to be returned, got nil")
	}

	if len(config.Metrics) != 6 {
		t.Errorf("Expected config to contain 5 metrics, but found %d", len(config.Metrics))
		t.Fatalf("Output %s", config.Metrics)
	}

	for index, metric := range expectedConfig.Metrics {
		if metric.Target != config.Metrics[index].Target {
			t.Errorf("Expected Target on index %d to be %s, got %s", index, config.Metrics[index].Target, metric.Target)
		}
		if metric.Source.Mbean != config.Metrics[index].Source.Mbean {
			t.Errorf("Expected Source.Mbean on index %d to be %s, got %s", index, config.Metrics[index].Source.Mbean, metric.Source.Mbean)
		}
		if metric.Source.Attribute != config.Metrics[index].Source.Attribute {
			t.Errorf("Expected Source.Attribute on index %d to be %s, got %s", index, config.Metrics[index].Source.Attribute, metric.Source.Attribute)
		}
		if metric.Source.Path != config.Metrics[index].Source.Path {
			t.Errorf("Expected Source.Path on index %d to be %s, got %s", index, config.Metrics[index].Source.Path, metric.Source.Path)
		}
	}
}
