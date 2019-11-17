package utils

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	//"k8s.io/client-go/tools/clientcmd"
	log "github.com/sirupsen/logrus"
	//"os"
	//"strings"
	//k8serror "k8s.io/apimachinery/pkg/api/errors"
)

// GetEnvFromExperiment will return ENVList from the Experiment
func GetEnvFromExperiment(appns string, chaosExperiment string, config *rest.Config) map[string]string {
	_, litmusClientSet, err := GenerateClientSets(config)
	if err != nil {
		log.Info(err)
	}
	m := make(map[string]string)
	experimentEnv, err := litmusClientSet.LitmuschaosV1alpha1().ChaosExperiments(appns).Get(chaosExperiment, metav1.GetOptions{})
	envList := experimentEnv.Spec.Definition.ENVList
	for i := range envList {
		key := envList[i].Name
		value := envList[i].Value
		m[key] = value
	}
	return m

}