package aws

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsCfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"

	"github.com/justtrackio/gosoline/pkg/cfg"
)

func GetCredentialsProvider(ctx context.Context, config cfg.Config, settings ClientSettings) (aws.CredentialsProvider, error) {
	if len(settings.AssumeRole) > 0 {
		return GetAssumeRoleCredentialsProvider(ctx, settings.AssumeRole)
	}

	if webIdentitySettings := UnmarshalWebIdentitySettings(config); settings.UseWebIdentity && webIdentitySettings != nil {
		return GetWebIdentityRoleProvider(ctx, webIdentitySettings)
	}

	if creds := UnmarshalCredentials(config); creds != nil {
		return credentials.NewStaticCredentialsProvider(creds.AccessKeyID, creds.SecretAccessKey, creds.SessionToken), nil
	}

	return nil, nil
}

func GetAssumeRoleCredentialsProvider(ctx context.Context, roleArn string) (aws.CredentialsProvider, error) {
	var err error
	var cfg aws.Config

	if cfg, err = awsCfg.LoadDefaultConfig(ctx); err != nil {
		return nil, fmt.Errorf("can not load default aws config: %w", err)
	}

	stsClient := sts.NewFromConfig(cfg)

	return stscreds.NewAssumeRoleProvider(stsClient, roleArn), nil
}

func GetWebIdentityRoleProvider(ctx context.Context, webIdentitySettings *WebIdentitySettings) (aws.CredentialsProvider, error) {
	var (
		err error
		c   aws.Config
	)
	if c, err = awsCfg.LoadDefaultConfig(ctx); err != nil {
		return nil, fmt.Errorf("can not load default aws config: %w", err)
	}
	client := sts.NewFromConfig(c)
	credsCache := aws.NewCredentialsCache(stscreds.NewWebIdentityRoleProvider(
		client,
		webIdentitySettings.RoleARN,
		stscreds.IdentityTokenFile(webIdentitySettings.TokenFilePath),
		func(o *stscreds.WebIdentityRoleOptions) {
			o.RoleSessionName = fmt.Sprintf("gosoline-%s", time.Now().Format("20060102150405"))
		}))

	return credsCache, nil
}

func UnmarshalCredentials(config cfg.Config) *Credentials {
	if !config.IsSet("cloud.aws.credentials") {
		return nil
	}

	creds := &Credentials{}
	config.UnmarshalKey("cloud.aws.credentials", creds)

	return creds
}

func UnmarshalWebIdentitySettings(config cfg.Config) *WebIdentitySettings {
	if !config.IsSet("aws.role_arn") || !config.IsSet("aws.web_identity_token_file") {
		return nil
	}

	settings := &WebIdentitySettings{}
	config.UnmarshalKey("aws", settings)

	return settings
}
