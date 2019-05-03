package cloud

import (
	"github.com/applike/gosoline/pkg/cfg"
	"github.com/applike/gosoline/pkg/mon"
	"github.com/aws/aws-sdk-go/aws"
	ssm2 "github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
	"strings"
)

type SsmParameters map[string]string

type SimpleSystemsManager struct {
	logger mon.Logger
	client ssmiface.SSMAPI
}

func NewSimpleSystemsManager(config cfg.Config, logger mon.Logger) *SimpleSystemsManager {
	client := GetSystemsManagerClient(config, logger)

	return &SimpleSystemsManager{
		logger: logger,
		client: client,
	}
}

func (ssm SimpleSystemsManager) GetParameters(path string) (SsmParameters, error) {
	input := &ssm2.GetParametersByPathInput{
		Path:           aws.String(path),
		Recursive:      aws.Bool(true),
		WithDecryption: aws.Bool(true),
	}

	out, err := ssm.client.GetParametersByPath(input)

	if err != nil {
		return SsmParameters{}, err
	}

	params := make(SsmParameters)
	for _, p := range out.Parameters {
		key := strings.Replace(*p.Name, path+"/", "", -1)
		params[key] = *p.Value
	}

	return params, nil
}
