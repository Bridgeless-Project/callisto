package gov

import (
	"encoding/json"
	"fmt"
	"github.com/forbole/bdjuno/v4/types"
	"github.com/rs/zerolog/log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const (
	ipfsGateway  = "https://gateway.pinata.cloud"
	defaultTitle = "Proposal"
	httpTimeout  = 4
)

// TODO: Get uri format from Core
var uriFormat = regexp.MustCompile(`ipfs://(Qm[1-9A-HJ-NP-Za-km-z]{44}|baf[0-9a-z]{56,})`)

func fetchIPFSProposalMetadata(proposal types.Proposal) types.Proposal {
	logDebug := log.Debug().Str("module", "gov").Str("proposal_id",
		fmt.Sprintf("%d", proposal.ProposalID))

	ipfsUri := convertIPFSURI(proposal.Metadata)
	if ipfsUri == nil {
		// if metadata link is provided in incorrect formating use default title
		setDefaultTitle(&proposal)
		logDebug.Msg("proposal metadata field does not match appropriate formatting")
		return proposal
	}
	client := &http.Client{
		Timeout: httpTimeout * time.Second,
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", ipfsGateway, *ipfsUri), nil)
	if err != nil {
		logDebug.Msg(fmt.Sprintf("error creating GET proposal metadata request: %s", err.Error()))
		setDefaultTitle(&proposal)
		return proposal
	}

	resp, err := client.Do(req)
	if err != nil {
		logDebug.Msg(fmt.Sprintf("error fetching proposal metadata: %s", err.Error()))
		setDefaultTitle(&proposal)
		return proposal
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logDebug.Msg(fmt.Sprintf("error fetching proposal metadata, code: %d", resp.StatusCode))
		setDefaultTitle(&proposal)
		return proposal
	}

	var metadata types.ProposalMetadata
	if err := json.NewDecoder(resp.Body).Decode(&metadata); err != nil {
		logDebug.Msg(fmt.Sprintf("error parsing proposal metadata: %s", err.Error()))
		setDefaultTitle(&proposal)
		return proposal
	}

	if strings.TrimSpace(metadata.Title) == "" {
		logDebug.Msg("proposal metadata title field is empty, setting it to default")
		setDefaultTitle(&proposal)
		return proposal
	}

	proposal.ProposalTitle = metadata.Title
	proposal.ProposalDescription = metadata.Description
	return proposal
}

func convertIPFSURI(uri string) *string {
	match := uriFormat.FindStringSubmatch(uri)
	if len(match) > 1 {
		convertedUri := "ipfs/" + match[1]
		return &convertedUri
	}

	return nil
}

func setDefaultTitle(proposal *types.Proposal) {
	proposal.ProposalTitle = fmt.Sprintf("%s %d", defaultTitle, proposal.ProposalID)
}
