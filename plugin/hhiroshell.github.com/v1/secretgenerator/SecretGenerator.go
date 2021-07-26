package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

func main() {
	execute()
}

var (
	name      string
	namespace string
	literals  []string
)

func execute() {
	rootCmd := &cobra.Command{
		Use:   "kustomize-go-exec-plugin",
		Short: "Kustomize Go Exec Plugin is a sample kustomize exec plugin written in Go",
		Long: `Kustomize Go Exec Plugin is a sample kustomize exec plugin written in Go.
                Kustomize supports two types of plugins (Go / Exec), but Go plugin is "quite annoying" because it is
                build on the Go plugin system.
                So this repository is trying to write a Kustomize exec plugin in Go`,
		RunE: run,
	}

	rootCmd.Flags().StringVar(&name, "name", "", "Name of the generated Secret")
	rootCmd.MarkFlagRequired("name")
	rootCmd.Flags().StringVar(&namespace, "namespace", "default", "Namespace of the generated Secret")
	rootCmd.Flags().StringArrayVar(&literals, "literal", nil, "Literal key-value pairs used as data of the Secret")
	rootCmd.MarkFlagRequired("literal")

	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(_ *cobra.Command, _ []string) error {
	var data = map[string][]byte{}
	for _, l := range literals {
		key := l[:strings.Index(l, "=")]
		value := []byte(l[strings.Index(l, "=")+1:])
		data[key] = value
	}

	cm := corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Secret",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Data: data,
	}

	bytes, err := yaml.Marshal(cm)
	if err != nil {
		return err
	}
	writer := os.Stdout
	fmt.Fprint(writer, string(bytes))

	return nil
}
