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
	cmd := &cobra.Command{
		Use:   "SecretGenerator",
		Short: "Kustomize exec plugin that generates a kubernetes secret",
		Long: `Kustomize exec plugin that generates a manifest of kubernetes secret written in yaml format.
You can specify the secret name, namespace and literals(key-value pairs) of the data field, via CLI flags.`,
		RunE: run,
	}

	cmd.Flags().StringVar(&name, "name", "", "Name of the generated Secret")
	cmd.MarkFlagRequired("name")
	cmd.Flags().StringVar(&namespace, "namespace", "default", "Namespace of the generated Secret")
	cmd.Flags().StringArrayVar(&literals, "literal", nil, "Literal key-value pairs used as data of the Secret. (e.g. --literal key=value)")
	cmd.MarkFlagRequired("literal")

	err := cmd.Execute()
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
	fmt.Fprint(os.Stdout, string(bytes))

	return nil
}
