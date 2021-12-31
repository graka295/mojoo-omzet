package container

import (
	"mojoo/omzet/helper"

	"gorm.io/gorm"
)

// Container contract
type Container interface {
	GetConnection() *gorm.DB
	GetConfig() helper.Config
	GetLoger() helper.Log
	GetValidate() helper.Validation
	GetResponse() helper.Response
}

// ContainerImpl struct Container
type ContainerImpl struct {
	db       helper.DB
	config   helper.Config
	logger   helper.Log
	validate helper.Validation
	response helper.Response
}

// NewContainer is constructor.
func NewContainer(rep helper.DB, config helper.Config, logger helper.Log, res helper.Response, validate helper.Validation) Container {
	return &ContainerImpl{db: rep, config: config, logger: logger, response: res, validate: validate}
}

// GetConnection get connection
func (x ContainerImpl) GetConnection() *gorm.DB {
	return x.db.Connection()
}

// GetConfig get Config
func (x ContainerImpl) GetConfig() helper.Config {
	return x.config
}

// GetLoger get Loger
func (x ContainerImpl) GetLoger() helper.Log {
	return x.logger
}

// GetValidate get Validate
func (x ContainerImpl) GetValidate() helper.Validation {
	return x.validate
}

// GetResponse get Response
func (x ContainerImpl) GetResponse() helper.Response {
	return x.response
}
