package main

import (
	"github.com/twilio-voice/mcp-server/config"
	"github.com/twilio-voice/mcp-server/models"
	tools_voicev1connectionpolicytarget "github.com/twilio-voice/mcp-server/tools/voicev1connectionpolicytarget"
	tools_voicev1archivedcall "github.com/twilio-voice/mcp-server/tools/voicev1archivedcall"
	tools_voicev1iprecord "github.com/twilio-voice/mcp-server/tools/voicev1iprecord"
	tools_voicev1sourceipmapping "github.com/twilio-voice/mcp-server/tools/voicev1sourceipmapping"
	tools_voicev1highriskspecialprefix "github.com/twilio-voice/mcp-server/tools/voicev1highriskspecialprefix"
	tools_voicev1country "github.com/twilio-voice/mcp-server/tools/voicev1country"
	tools_voicev1byoctrunk "github.com/twilio-voice/mcp-server/tools/voicev1byoctrunk"
	tools_voicev1connectionpolicy "github.com/twilio-voice/mcp-server/tools/voicev1connectionpolicy"
	tools_voicev1settings "github.com/twilio-voice/mcp-server/tools/voicev1settings"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_voicev1connectionpolicytarget.CreateDeleteconnectionpolicytargetTool(cfg),
		tools_voicev1connectionpolicytarget.CreateFetchconnectionpolicytargetTool(cfg),
		tools_voicev1archivedcall.CreateDeletearchivedcallTool(cfg),
		tools_voicev1iprecord.CreateListiprecordTool(cfg),
		tools_voicev1sourceipmapping.CreateListsourceipmappingTool(cfg),
		tools_voicev1highriskspecialprefix.CreateListdialingpermissionshrsprefixesTool(cfg),
		tools_voicev1country.CreateListdialingpermissionscountryTool(cfg),
		tools_voicev1byoctrunk.CreateListbyoctrunkTool(cfg),
		tools_voicev1byoctrunk.CreateFetchbyoctrunkTool(cfg),
		tools_voicev1byoctrunk.CreateDeletebyoctrunkTool(cfg),
		tools_voicev1country.CreateFetchdialingpermissionscountryTool(cfg),
		tools_voicev1connectionpolicytarget.CreateListconnectionpolicytargetTool(cfg),
		tools_voicev1connectionpolicy.CreateListconnectionpolicyTool(cfg),
		tools_voicev1connectionpolicy.CreateDeleteconnectionpolicyTool(cfg),
		tools_voicev1connectionpolicy.CreateFetchconnectionpolicyTool(cfg),
		tools_voicev1sourceipmapping.CreateDeletesourceipmappingTool(cfg),
		tools_voicev1sourceipmapping.CreateFetchsourceipmappingTool(cfg),
		tools_voicev1iprecord.CreateDeleteiprecordTool(cfg),
		tools_voicev1iprecord.CreateFetchiprecordTool(cfg),
		tools_voicev1settings.CreateFetchdialingpermissionssettingsTool(cfg),
	}
}
