package bedrock

import (
	"context"
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/Uptycs/cloudquery/utilities"

	"github.com/Uptycs/basequery-go/plugin/table"
	extaws "github.com/Uptycs/cloudquery/extension/aws"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagent"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func ListAgentActionGroupsColumns() []table.ColumnDefinition {
	return []table.ColumnDefinition{
		table.TextColumn("account_id"),
		table.TextColumn("region_code"),
		table.TextColumn("region"),
		table.TextColumn("agent_id"),
		table.TextColumn("action_group_id"),
		table.TextColumn("name"),
		table.TextColumn("state"),
		table.TextColumn("description"),
		table.TextColumn("updated_at"),
	}
}

func ListAgentActionGroupsGenerate(osqCtx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	resultMap := make([]map[string]string, 0)
	if len(utilities.ExtConfiguration.ExtConfAws.Accounts) == 0 && extaws.ShouldProcessAccount("aws_bedrock_agent_action_group", utilities.AwsAccountID) {
		utilities.GetLogger().WithFields(log.Fields{
			"tableName": "aws_bedrock_agent_action_group",
			"account":   "default",
		}).Info("processing account")
		results, err := processAccountListAgentActionGroups(osqCtx, queryContext, nil)
		if err != nil {
			return resultMap, err
		}
		resultMap = append(resultMap, results...)
	} else {
		for _, account := range utilities.ExtConfiguration.ExtConfAws.Accounts {
			if !extaws.ShouldProcessAccount("aws_bedrock_agent_action_group", account.ID) {
				continue
			}
			utilities.GetLogger().WithFields(log.Fields{
				"tableName": "aws_bedrock_agent_action_group",
				"account":   account.ID,
			}).Info("processing account")
			results, err := processAccountListAgentActionGroups(osqCtx, queryContext, &account)
			if err != nil {
				continue
			}
			resultMap = append(resultMap, results...)
		}
	}

	return resultMap, nil
}

func processRegionListAgentActionGroups(osqCtx context.Context, queryContext table.QueryContext, tableConfig *utilities.TableConfig, account *utilities.ExtensionConfigurationAwsAccount, region types.Region) ([]map[string]string, error) {
	resultMap := make([]map[string]string, 0)
	sess, err := extaws.GetAwsConfig(account, *region.RegionName)
	if err != nil {
		return resultMap, err
	}

	accountId := utilities.AwsAccountID
	if account != nil {
		accountId = account.ID
	}

	utilities.GetLogger().WithFields(log.Fields{
		"tableName": "aws_bedrock_agent_action_group",
		"account":   accountId,
		"region":    *region.RegionName,
	}).Debug("processing region")

	svc := bedrockagent.NewFromConfig(*sess)

	agentParams := &bedrockagent.ListAgentsInput{}
	agentPaginator := bedrockagent.NewListAgentsPaginator(svc, agentParams)

	for {
		agentPage, err := agentPaginator.NextPage(osqCtx)
		if err != nil {
			utilities.GetLogger().WithFields(log.Fields{
				"tableName": "aws_bedrock_agent_action_group",
				"account":   accountId,
				"region":    *region.RegionName,
				"task":      "ListAgents",
				"errString": err.Error(),
			}).Error("failed to list agents")
			return resultMap, err
		}

		for _, agent := range agentPage.AgentSummaries {
			targetVersion := agent.LatestAgentVersion
			if targetVersion == nil || *targetVersion == "" {
				draft := "DRAFT"
				targetVersion = &draft
			}
			agParams := &bedrockagent.ListAgentActionGroupsInput{
				AgentId:      agent.AgentId,
				AgentVersion: targetVersion,
			}

			agPaginator := bedrockagent.NewListAgentActionGroupsPaginator(svc, agParams)

			for {
				agPage, err := agPaginator.NextPage(osqCtx)
				if err != nil {
					utilities.GetLogger().WithFields(log.Fields{
						"tableName": "aws_bedrock_agent_action_group",
						"account":   accountId,
						"region":    *region.RegionName,
						"agentId":   *agent.AgentId,
						"task":      "ListAgentActionGroups",
						"errString": err.Error(),
					}).Error("failed to list action groups")
					break
				}

				byteArr, err := json.Marshal(agPage)
				if err != nil {
					utilities.GetLogger().WithFields(log.Fields{
						"tableName": "aws_bedrock_agent_action_group",
						"account":   accountId,
						"region":    *region.RegionName,
						"agentId":   *agent.AgentId,
						"task":      "ListAgentActionGroups",
						"errString": err.Error(),
					}).Error("failed to marshal response")
					break
				}
				table := utilities.NewTable(byteArr, tableConfig)
				for _, row := range table.Rows {
					if !extaws.ShouldProcessRow(osqCtx, queryContext, "aws_bedrock_agent_action_group", accountId, *region.RegionName, row) {
						continue
					}
					result := extaws.RowToMap(row, accountId, *region.RegionName, tableConfig)
					result["agent_id"] = *agent.AgentId
					resultMap = append(resultMap, result)
				}
				if !agPaginator.HasMorePages() {
					break
				}
			}
		}
		if !agentPaginator.HasMorePages() {
			break
		}
	}
	return resultMap, nil
}

func processAccountListAgentActionGroups(osqCtx context.Context, queryContext table.QueryContext, account *utilities.ExtensionConfigurationAwsAccount) ([]map[string]string, error) {
	resultMap := make([]map[string]string, 0)
	awsSession, err := extaws.GetAwsConfig(account, "us-east-1")
	if err != nil {
		return resultMap, err
	}
	regions, err := extaws.FetchRegions(osqCtx, awsSession)
	if err != nil {
		return resultMap, err
	}
	tableConfig, ok := utilities.TableConfigurationMap["aws_bedrock_agent_action_group"]
	if !ok {
		utilities.GetLogger().WithFields(log.Fields{
			"tableName": "aws_bedrock_agent_action_group",
		}).Error("failed to get table configuration")
		return resultMap, fmt.Errorf("table configuration not found")
	}
	for _, region := range regions {
		accountId := utilities.AwsAccountID
		if account != nil {
			accountId = account.ID
		}
		if !extaws.ShouldProcessRegion("aws_bedrock_agent_action_group", accountId, *region.RegionName) {
			continue
		}
		result, err := processRegionListAgentActionGroups(osqCtx, queryContext, tableConfig, account, region)
		if err != nil {
			continue
		}
		resultMap = append(resultMap, result...)
	}
	return resultMap, nil
}
