package configuration

import (
	"strconv"
	"strings"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/haproxytech/models"
)

// GetStickRequestRules returns a struct with configuration version and an array of
// configured stick request rules in the specified backend. Returns error on fail.
func (c *LBCTLClient) GetStickRequestRules(backend string, transactionID string) (*models.GetStickRequestRulesOKBody, error) {
	stickReqRulesString, err := c.executeLBCTL("l7-farm-stickreq-dump", transactionID, backend)
	if err != nil {
		return nil, err
	}

	stickReqRules := c.parseStickRequestRules(stickReqRulesString)

	v, err := c.GetVersion()
	if err != nil {
		return nil, err
	}

	return &models.GetStickRequestRulesOKBody{Version: v, Data: stickReqRules}, nil
}

// GetStickRequestRule returns a struct with configuration version and a requested stick request rule
// in the specified backend. Returns error on fail or if stick request rule does not exist.
func (c *LBCTLClient) GetStickRequestRule(id int64, backend string, transactionID string) (*models.GetStickRequestRuleOKBody, error) {
	stickReqRuleStr, err := c.executeLBCTL("l7-farm-stickreq-show", transactionID, backend, strconv.FormatInt(id, 10))
	if err != nil {
		return nil, err
	}
	stickReqRule := &models.StickRequestRule{ID: id}

	c.parseObject(stickReqRuleStr, stickReqRule)

	v, err := c.GetVersion()
	if err != nil {
		return nil, err
	}

	return &models.GetStickRequestRuleOKBody{Version: v, Data: stickReqRule}, nil
}

// DeleteStickRequestRule deletes a stick request rule in configuration. One of version or transactionID is
// mandatory. Returns error on fail, nil on success.
func (c *LBCTLClient) DeleteStickRequestRule(id int64, backend string, transactionID string, version int64) error {
	return c.deleteObject(strconv.FormatInt(id, 10), "stickreq", backend, "farm", transactionID, version)
}

// CreateStickRequestRule creates a stick request rule in configuration. One of version or transactionID is
// mandatory. Returns error on fail, nil on success.
func (c *LBCTLClient) CreateStickRequestRule(backend string, data *models.StickRequestRule, transactionID string, version int64) error {
	validationErr := data.Validate(strfmt.Default)
	if validationErr != nil {
		return NewConfError(ErrValidationError, validationErr.Error())
	}
	return c.createObject(strconv.FormatInt(data.ID, 10), "stickreq", backend, "farm", data, nil, transactionID, version)
}

// EditStickRequestRule edits a stick request rule in configuration. One of version or transactionID is
// mandatory. Returns error on fail, nil on success.
func (c *LBCTLClient) EditStickRequestRule(id int64, backend string, data *models.StickRequestRule, transactionID string, version int64) error {
	validationErr := data.Validate(strfmt.Default)
	if validationErr != nil {
		return NewConfError(ErrValidationError, validationErr.Error())
	}
	ondiskR, err := c.GetStickRequestRule(id, backend, transactionID)
	if err != nil {
		return err
	}

	return c.editObject(strconv.FormatInt(data.ID, 10), "stickreq", backend, "farm", data, ondiskR, nil, transactionID, version)
}

func (c *LBCTLClient) parseStickRequestRules(response string) models.StickRequestRules {
	stickReqRules := make(models.StickRequestRules, 0, 1)
	for _, stickReqRulesStr := range strings.Split(response, "\n\n") {
		if strings.TrimSpace(stickReqRulesStr) == "" {
			continue
		}
		idStr, _ := splitHeaderLine(stickReqRulesStr)
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			id = 0
		}

		stickReqRulesObj := &models.StickRequestRule{ID: id}
		c.parseObject(stickReqRulesStr, stickReqRulesObj)
		stickReqRules = append(stickReqRules, stickReqRulesObj)
	}
	return stickReqRules
}
