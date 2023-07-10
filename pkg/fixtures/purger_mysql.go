package fixtures

import (
	"context"
	"fmt"

	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/db"
	"github.com/justtrackio/gosoline/pkg/log"
)

const (
	foreignKeyChecksStatement = "SET FOREIGN_KEY_CHECKS=%d;"
	truncateTableStatement    = "TRUNCATE TABLE %s;"
)

type mysqlPurger struct {
	client    db.Client
	logger    log.Logger
	tableName string
}

func newMysqlPurger(config cfg.Config, logger log.Logger, tableName string) (*mysqlPurger, error) {
	client, err := db.NewClient(config, logger, "default")
	if err != nil {
		return nil, fmt.Errorf("can not create db client: %w", err)
	}

	return &mysqlPurger{client: client, logger: logger, tableName: tableName}, nil
}

func (p *mysqlPurger) purgeMysql(ctx context.Context) error {
	err := p.setForeignKeyChecks(0)
	if err != nil {
		p.logger.WithFields(log.Fields{
			"error": err,
		}).Error("error disabling foreign key checks")

		return err
	}

	defer func() {
		err := p.setForeignKeyChecks(1)
		if err != nil {
			p.logger.WithFields(log.Fields{
				"error": err,
			}).Error("error enabling foreign key checks")
		}
	}()

	_, err = p.client.Exec(ctx, fmt.Sprintf(truncateTableStatement, p.tableName))

	if err != nil {
		p.logger.WithFields(log.Fields{
			"error": err,
		}).Error("error truncating table %s", p.tableName)
		return err
	}

	return nil
}

func (p *mysqlPurger) setForeignKeyChecks(enabled int) error {
	ctx := context.Background()
	_, err := p.client.Exec(ctx, fmt.Sprintf(foreignKeyChecksStatement, enabled))

	return err
}
