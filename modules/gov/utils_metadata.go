package gov

import (
	"encoding/json"
	"fmt"
	"github.com/forbole/bdjuno/v4/types"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const (
	IPFSDomain   = "https://gateway.pinata.cloud"
	defaultTitle = "Proposal"
)

var uriFormat = regexp.MustCompile(`ipfs://(Qm[1-9A-HJ-NP-Za-km-z]{44}|baf[0-9a-z]{56,})`)

func FetchIPFSProposalMetadata(proposal *types.Proposal) {
	if len(strings.TrimSpace(proposal.Metadata)) == 0 {
	}

	ipfsUri, ok := convertIPFSURI(proposal.Metadata)
	if !ok {
		// if metadata link is provided in incorrect formating use default title
		defTitle(proposal)
		return
	}
	client := &http.Client{
		Timeout: time.Second * 4,
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", IPFSDomain, ipfsUri), nil)
	if err != nil {
		defTitle(proposal)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		defTitle(proposal)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		defTitle(proposal)
		return
	}

	var metadata types.ProposalMetadata
	if err := json.NewDecoder(resp.Body).Decode(&metadata); err != nil {
		defTitle(proposal)
		return
	}

	if strings.TrimSpace(metadata.Title) == "" {
		defTitle(proposal)
		return
	}

	proposal.ProposalTitle = metadata.Title
	proposal.ProposalDescription = metadata.Description
}

func convertIPFSURI(uri string) (string, bool) {
	match := uriFormat.FindStringSubmatch(uri)
	if len(match) > 1 {
		return "ipfs/" + match[1], true
	}

	return "", false
}

func defTitle(proposal *types.Proposal) {
	proposal.ProposalTitle = fmt.Sprintf("%s %d", defaultTitle, proposal.ProposalID)
}
