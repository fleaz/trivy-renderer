package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aquasecurity/trivy-operator/pkg/apis/aquasecurity/v1alpha1"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")

	if envconf := os.Getenv("KUBECONFIG"); envconf != "" {
		kubeconfig = envconf
	}

	// use the current context in kubeconfig
	cfg := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig}, nil)

	kubeConfig, err := cfg.ClientConfig()
	if err != nil {
		fmt.Printf("error getting Kubernetes config: %v\n", err)
		os.Exit(1)
	}
	// Create a dynamic client
	dynClient, err := dynamic.NewForConfig(kubeConfig)
	if err != nil {
		panic(fmt.Sprintf("Failed to create dynamic client: %v", err))
	}

	// Define the GVR for VulnerabilityReports
	gvr := schema.GroupVersionResource{
		Group:    "aquasecurity.github.io",
		Version:  "v1alpha1",
		Resource: "vulnerabilityreports",
	}

	// List all VulnerabilityReports across all namespaces
	currentNS, _, err := cfg.Namespace()
	if err != nil {
		currentNS = ""
	}

	list, err := dynClient.Resource(gvr).Namespace(currentNS).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		panic(fmt.Sprintf("Failed to list VulnerabilityReports: %v", err))
	}

	// Iterate and decode each item
	for _, item := range list.Items {
		var report v1alpha1.VulnerabilityReport
		objJSON, err := item.MarshalJSON()
		if err != nil {
			fmt.Printf("Error marshaling item: %v\n", err)
			continue
		}
		if err := json.Unmarshal(objJSON, &report); err != nil {
			fmt.Printf("Error unmarshaling to typed struct: %v\n", err)
			continue
		}

		if len(report.Report.Vulnerabilities) == 0 {
			continue
		}
		fmt.Printf("Name: %s, Namespace: %s\n", report.Name, report.Namespace)
		table := tablewriter.NewTable(os.Stdout)
		table.Header("Resource", "CVE", "Severity", "Installed", "Fixed", "Meta")
		for _, vuln := range report.Report.Vulnerabilities {
			table.Append([]string{vuln.Resource, vuln.VulnerabilityID, colorizedSeverity(vuln.Severity), vuln.InstalledVersion, vuln.FixedVersion, fmt.Sprintf("%.120s", vuln.Title)})
		}
		table.Render()

	}
}

func colorizedSeverity(sev v1alpha1.Severity) string {
	if sev == "CRITICAL" {
		return color.New(color.BgRed, color.FgBlack).Sprintf("CRITICAL")
	} else if sev == "HIGH" {
		return color.New(color.FgRed).Sprintf("HIGH")
	} else if sev == "MEDIUM" {
		return color.New(color.FgYellow).Sprintf("MEDIUM")
	} else {
		return string(sev)
	}
}
