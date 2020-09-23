// header.go.

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

// Common Message Header Names for HTTP and Mail Protocols.

// Notes.
//
// This Packages contains all IANA registered Header Names which have an RFC Reference as of 2019-05-26.
// For more Information visit this URL:
// https://www.iana.org/assignments/message-headers/message-headers.xml
// This List may contain some old deprecated and obsolete Header Names.

// Delimiters.
const DelimiterCommaSpace = ", "

// MakeListOfHeaders Function composes a List of Headers delimited by Comma and
// Space.
func MakeListOfHeaders(
	headers []string,
) (headersList string) {

	var iLast int

	iLast = len(headers) - 1
	for i, header := range headers {
		if i != iLast {
			headersList = headersList + header + DelimiterCommaSpace
		} else {
			headersList = headersList + header
		}
	}

	return
}
