// netnews.go.

////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019..2020 by Vault Thirteen.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
////////////////////////////////////////////////////////////////////////////////
//
// Web Site Address:	https://github.com/vault-thirteen.
//
////////////////////////////////////////////////////////////////////////////////

package header

// Netnews Header Field Names.

// Permanent.
const (
	NetnewsHeaderAlsoControl     = "Also-Control"      // [RFC1849][RFC5536]
	NetnewsHeaderApproved        = "Approved"          // [RFC5536]
	NetnewsHeaderArchive         = "Archive"           // [RFC5536]
	NetnewsHeaderArchivedAt      = "Archived-At"       // [RFC5064]
	NetnewsHeaderArticleNames    = "Article-Names"     // [RFC1849][RFC5536]
	NetnewsHeaderArticleUpdates  = "Article-Updates"   // [RFC1849][RFC5536]
	NetnewsHeaderCancelKey       = "Cancel-Key"        // [RFC8315]
	NetnewsHeaderCancelLock      = "Cancel-Lock"       // [RFC8315]
	NetnewsHeaderComments        = "Comments"          // [RFC5536][RFC5322]
	NetnewsHeaderControl         = "Control"           // [RFC5536]
	NetnewsHeaderDate            = "Date"              // [RFC5536][RFC5322]
	NetnewsHeaderDateReceived    = "Date-Received"     // [RFC0850][RFC5536]
	NetnewsHeaderDistribution    = "Distribution"      // [RFC5536]
	NetnewsHeaderExpires         = "Expires"           // [RFC5536]
	NetnewsHeaderFollowupTo      = "Followup-To"       // [RFC5536]
	NetnewsHeaderFrom            = "From"              // [RFC5536][RFC5322]
	NetnewsHeaderInjectionDate   = "Injection-Date"    // [RFC5536]
	NetnewsHeaderInjectionInfo   = "Injection-Info"    // [RFC5536]
	NetnewsHeaderKeywords        = "Keywords"          // [RFC5536][RFC5322]
	NetnewsHeaderLines           = "Lines"             // [RFC5536][RFC3977]
	NetnewsHeaderMessageID       = "Message-ID"        // [RFC5536][RFC5322]
	NetnewsHeaderNewsgroups      = "Newsgroups"        // [RFC5536]
	NetnewsHeaderNNTPPostingDate = "NNTP-Posting-Date" // [RFC5536]
	NetnewsHeaderNNTPPostingHost = "NNTP-Posting-Host" // [RFC2980][RFC5536]
	NetnewsHeaderOrganization    = "Organization"      // [RFC5536]
	NetnewsHeaderOriginalSender  = "Original-Sender"   // [RFC5537]
	NetnewsHeaderPath            = "Path"              // [RFC5536]
	NetnewsHeaderPostingVersion  = "Posting-Version"   // [RFC0850][RFC5536]
	NetnewsHeaderReferences      = "References"        // [RFC5536][RFC5322]
	NetnewsHeaderRelayVersion    = "Relay-Version"     // [RFC0850][RFC5536]
	NetnewsHeaderReplyTo         = "Reply-To"          // [RFC5536][RFC5322]
	NetnewsHeaderSeeAlso         = "See-Also"          // [RFC1849][RFC5536]
	NetnewsHeaderSender          = "Sender"            // [RFC5536][RFC5322]
	NetnewsHeaderSubject         = "Subject"           // [RFC5536][RFC5322]
	NetnewsHeaderSummary         = "Summary"           // [RFC5536]
	NetnewsHeaderSupersedes      = "Supersedes"        // [RFC5536][RFC2156]
	NetnewsHeaderUserAgent       = "User-Agent"        // [RFC5536][RFC2616]
	NetnewsHeaderXref            = "Xref"              // [RFC5536]
)

// Provisional.
const (
	NetnewsHeaderJabberID    = "Jabber-ID"     // [RFC7259]
	NetnewsHeaderXArchivedAt = "X-Archived-At" // [RFC5064]
	NetnewsHeaderXPGPSig     = "X-PGP-Sig"     // [ftp://ftp.isc.org/pub/pgpcontrol/FORMAT][https://ftp.isc.org/pub/pgpcontrol/FORMAT]
)
